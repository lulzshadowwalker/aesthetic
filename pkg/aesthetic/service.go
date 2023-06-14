package aesthetic

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/lulzshadowwalker/aesthetic/internal/config"
)

// TODO remove dest and return the converted image instead
func ConvertImage(source, destination string) {
	image, err := gg.LoadImage(source)

	if err != nil {
		log.Fatalf("error reading the image file %q", err)
	}

	fmt.Println("Converting the image ..")
	im := convertFrame(image)

	fmt.Println("Saving the image ..")

	file, err := os.Create(filepath.Join(destination, "abooba.png"))

	if err != nil {
		fmt.Printf("error saving image to disk %q", err)
	}

	defer file.Close()

	err = png.Encode(file, im)
	if err != nil {
		fmt.Printf("error saving image to disk %q", err)
	}
	fmt.Println("Image saved 👀")
}

// TODO migrate from gg to gocv
func convertFrame(image image.Image) image.Image {
	if config.Grayscale() {
		image = imaging.Grayscale(image)
	}

	context := gg.NewContext(image.Bounds().Dx(), image.Bounds().Dy())

	context.SetRGB(0, 0, 0)
	context.Clear()

	const (
		samplingInterval = 50
	)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for y := 0; y < image.Bounds().Dy(); y += samplingInterval {
		for x := 0; x < image.Bounds().Dx(); x += samplingInterval {
			squareSize := r.Intn(16) + 4
			color := image.At(x, y)
			context.SetColor(color)
			context.DrawCircle(float64(x), float64(y), float64(squareSize))
			context.Fill()
		}
	}

	return context.Image()
}
