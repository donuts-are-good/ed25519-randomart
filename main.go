package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

func main() {
	pubKey := generatePublicKey()
	img := createImage(pubKey)
	saveImage("ed25519_key.png", img)
}

func generatePublicKey() ed25519.PublicKey {
	_, pubKey, _ := ed25519.GenerateKey(rand.Reader)
	return ed25519.PublicKey(pubKey)
}

func createImage(pubKey ed25519.PublicKey) *image.RGBA {
	const size = 512
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.Draw(img, img.Bounds(), image.Transparent, image.Point{}, draw.Src)

	numCircles := len(pubKey)
	for i := 0; i < numCircles; i++ {
		x := (int(pubKey[i]) % (size / 2)) + size/4
		y := (int(pubKey[(i+1)%numCircles]) % (size / 2)) + size/4
		radius := int(pubKey[(i+2)%numCircles]) % (size / 8)
		colorValue := uint8(pubKey[(i+3)%numCircles])
		drawCircle(img, x, y, radius, color.RGBA{colorValue, 255 - colorValue, colorValue / 2, 255})
	}

	return img
}

func drawCircle(img *image.RGBA, x, y, radius int, col color.RGBA) {
	for i := -radius; i <= radius; i++ {
		for j := -radius; j <= radius; j++ {
			dist := math.Sqrt(float64(i*i + j*j))
			if dist <= float64(radius) {
				img.Set(x+i, y+j, col)
			}
		}
	}
}

func saveImage(filename string, img *image.RGBA) {
	file, _ := os.Create(filename)
	defer file.Close()
	png.Encode(file, img)
}
