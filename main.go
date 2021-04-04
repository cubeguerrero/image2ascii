package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"
	"math"
	"os"
	"strings"
)

var chars = []string{}

func init() {
	chars = strings.Split("`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$", "")
}

type pixel struct {
	R uint8
	G uint8
	B uint8
}

func main() {
	reader, err := os.Open("./ascii-pineapple.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bound := m.Bounds()

	for y := bound.Min.Y; y <= bound.Max.Y; y++ {
		for x := bound.Min.X; x <= bound.Max.X; x++ {
			// c := m.At(x, y).(color.CMYK)
			// r, g, b := color.CMYKToRGB(c.C, c.M, c.Y, c.K)
			c := m.At(x, y).(color.YCbCr)
			r, g, b := color.YCbCrToRGB(c.Y, c.Cb, c.Cr)
			br := calculateBrightness(r, g, b)
			fmt.Print(strings.Repeat(chars[br], 3))
		}
		fmt.Print("\n")
	}
}

func calculateBrightness(r, g, b uint8) int {
	sum := float64(r + g + b)
	avg := sum / 3
	br := avg / 255 * float64(len(chars))
	return int(math.Round(br))
}
