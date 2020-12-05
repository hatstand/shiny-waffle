package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"os/signal"
	"sort"
	"sync"
	"time"

	"github.com/disintegration/gift"
	"github.com/hatstand/shinywaffle/weather"
	"github.com/hatstand/shinywaffle/wirelesstag"
	"github.com/pbnjay/pixfont"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"

	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/experimental/devices/inky"
	"periph.io/x/periph/host"
)

var path = flag.String("template", "template.png", "Path to template image")

func drawLabel(m draw.Image, data string, x, y int) {
	pixfont.DrawString(m, x, y, data, color.White)
}

func drawTime(m draw.Image) {
	t := time.Now().Format("15:04:05 02/01")
	drawLabel(m, fmt.Sprintf("Updated: %s", t), 0, 104-8)
}

func threshold(r, g, b, a float32) (float32, float32, float32, float32) {
	if r == 0 && g == 0 && b == 0 && a != 0 {
		return 0, 0, 0, 1
	}
	return 1, 1, 1, 1
}

func getIcon(name string) (*oksvg.SvgIcon, error) {
	iconsFile, err := zip.OpenReader("weather-icons-master.zip")
	if err != nil {
		log.Fatalf("Failed to open icons zip: %v", err)
	}
	defer iconsFile.Close()
	for _, f := range iconsFile.File {
		if f.FileHeader.Name == fmt.Sprintf("weather-icons-master/svg/%s.svg", name) {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()
			icon, err := oksvg.ReadIconStream(rc)
			if err != nil {
				return nil, err
			}
			return icon, nil
		}
	}
	return nil, fmt.Errorf("Failed to find icon: %s", name)
}

func conditionCodeToName(code int) string {
	if code >= 200 && code < 300 {
		return "thunderstorm"
	}
	if code >= 300 && code < 400 {
		return "sprinkle"
	}
	if code >= 500 && code < 600 {
		return "rain"
	}
	if code >= 600 && code < 700 {
		return "snow"
	}
	if code >= 700 && code < 800 {
		return "fog"
	}
	if code == 800 {
		return "clear"
	}
	if code > 800 && code < 900 {
		return "cloudy"
	}
	return "alien"
}

func drawWeather(m draw.Image) {
	obs, err := weather.FetchCurrentWeather("London")
	log.Printf("%+v", obs)
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	var iconName string
	if now.Before(time.Time(obs.Location.Sunrise)) || now.After(time.Time(obs.Location.Sunset)) {
		iconName = fmt.Sprintf("wi-night-%s", conditionCodeToName(obs.ConditionCode))
	} else {
		iconName = fmt.Sprintf("wi-day-%s", conditionCodeToName(obs.ConditionCode))
		if iconName == "wi-day-clear" {
			iconName = "wi-day-sunny"
		}
	}

	icon, err := getIcon(iconName)
	if err != nil {
		log.Fatalf("Failed to read SVG: %v", err)
	}

	w, h := int(icon.ViewBox.W), int(icon.ViewBox.H)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	scanner := rasterx.NewScannerGV(w, h, img, img.Bounds())
	raster := rasterx.NewDasher(w, h, scanner)
	icon.Draw(raster, 1.0)
	g := gift.New(gift.ColorFunc(threshold), gift.Invert())
	filtered := image.NewRGBA(g.Bounds(img.Bounds()))
	g.Draw(filtered, img)
	draw.Draw(
		m,
		image.Rect(212-w, 0, 212, h),
		filtered,
		image.ZP,
		draw.Over)
	temp := fmt.Sprintf("%.0fC", obs.CurrentTemp)
	drawLabel(m, temp, 212-w/2-pixfont.MeasureString(temp)/2, h)
}

func main() {
	flag.Parse()

	state, err := host.Init()
	if err != nil {
		log.Fatalf("Failed to init periph: %v", err)
	}
	log.Printf("%+v", state)
	log.Printf("%+v", spireg.All())

	port, err := spireg.Open("")
	if err != nil {
		log.Fatalf("Failed to open SPI port: %v", err)
	}
	dc := gpioreg.ByName("22")
	reset := gpioreg.ByName("27")
	busy := gpioreg.ByName("17")

	dev, err := inky.New(port, dc, reset, busy, &inky.Opts{
		Model:       inky.PHAT,
		ModelColor:  inky.Red,
		BorderColor: inky.Black,
	})
	if err != nil {
		log.Fatalf("Failed to open inky: %v", err)
	}
	dev.SetBorder(inky.Black)

	file, err := os.Open(*path)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		log.Fatalf("Failed to decode %s as png: %v", *path, err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var mu sync.Mutex

	go drawStatus(mu, img, dev)

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			go drawStatus(mu, img, dev)
		case <-c:
			return
		}
	}
}

func drawStatus(mu sync.Mutex, img image.Image, dev *inky.Dev) {
	mu.Lock()
	defer mu.Unlock()
	b := img.Bounds()
	m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(m, m.Bounds(), img, b.Min, draw.Src)

	tags, err := wirelesstag.GetTags()
	if err != nil {
		log.Fatalf("Failed to fetch tags: %v", err)
	}
	sort.Slice(tags, func(i, j int) bool { return tags[i].Name < tags[j].Name })
	for i, v := range tags {
		s := fmt.Sprintf("%s: %.1f°C", v.Name, v.Temperature)
		log.Println(s)
		drawLabel(m, s, 0, 16*(i+1))
	}
	drawTime(m)
	drawWeather(m)

	debug, err := os.Create("debug.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(debug, m); err != nil {
		log.Fatal(err)
	}
	if err := debug.Close(); err != nil {
		log.Fatal(err)
	}
	dev.Draw(m.Bounds(), m, image.Point{0, 0})
}
