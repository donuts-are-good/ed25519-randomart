package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

func main() {
	pubKey := generatePublicKey()
	frameCount := 2048
	for frame := 0; frame < frameCount; frame++ {
		img := createImage(pubKey, frame, frameCount)
		filename := fmt.Sprintf("img/ed25519_key_%02d.png", frame)
		saveImage(filename, img)
	}
}

func generatePublicKey() ed25519.PublicKey {
	pubKey, privKey, _ := ed25519.GenerateKey(rand.Reader)
	fmt.Printf("Public key: %x\n", pubKey)
	fmt.Printf("Private key: %x\n", privKey)
	return ed25519.PublicKey(pubKey)
}

func createImage(pubKey ed25519.PublicKey, frame, frameCount int) *image.RGBA {
	const size = 1024
	const margin = 0.2
	const scale = 1 - 2*margin
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.Draw(img, img.Bounds(), image.Transparent, image.Point{}, draw.Src)

	numCircles := len(pubKey)
	for i := 0; i < numCircles; i++ {
		shift := math.Pi / 4 * float64(frame) / float64(frameCount)
		x := int(float64(pubKey[i])/255.0*scale*float64(size) + float64(size)*margin + float64(size)*0.1*math.Sin(shift*float64(i+1)))
		y := int(float64(pubKey[(i+1)%numCircles])/255.0*scale*float64(size) + float64(size)*margin + float64(size)*0.1*math.Cos(shift*float64(i+1)))
		radius := int(float64(pubKey[(i+2)%numCircles]) / 255.0 * scale * float64(size) / 8.0)
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
