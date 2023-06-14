package aesthetic

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/lulzshadowwalker/aesthetic/internal/config"
	"gocv.io/x/gocv"
)

// TODO remove dest and return the converted image instead
func ConvertImage() {
	src, dest := config.Source(), config.Destination()
	image, err := gg.LoadImage(src)

	if err != nil {
		log.Fatalf("error reading the image file %q", err)
	}

	fmt.Println("Converting the image ..")
	im := convertFrame(image)

	fmt.Println("Saving the image ..")

	file, err := os.Create(filepath.Join(dest, "abooba.png"))

	if err != nil {
		fmt.Printf("error saving image to disk %q", err)
	}

	defer file.Close()

	err = png.Encode(file, im)
	if err != nil {
		fmt.Printf("error saving image to disk %q", err)
	}
	fmt.Println("image saved 👀")
}

func convertFrame(image image.Image) image.Image {
	if config.Grayscale() {
		image = imaging.Grayscale(image)
	}

	context := gg.NewContext(image.Bounds().Dx(), image.Bounds().Dy())

	if config.WhiteNoise() {
		context.SetRGB(1, 1, 1)
	} else {
		context.SetRGB(0, 0, 0)
	}
	context.Clear()

	samplingInterval := config.SampleInterval()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for y := 0; y < image.Bounds().Dy(); y += samplingInterval {
		for x := 0; x < image.Bounds().Dx(); x += samplingInterval {
			dotSize := r.Intn(config.DotSize()) + 3
			color := image.At(x, y)
			context.SetColor(color)

			// TODO try using a ✦ character instead of dots
			context.DrawCircle(float64(x), float64(y), float64(dotSize))
			context.Fill()
		}
	}

	return context.Image()
}

// TODO mgirate from gocv 🙃
func ConvertVideo() {
	src, dest := config.Source(), config.Destination()
	vid, err := gocv.VideoCaptureFile(src)

	if err != nil {
		panic(err)
	}

	defer vid.Close()

	frame := gocv.NewMat()
	defer frame.Close()

	tempDir := filepath.Join(dest, ".aes-temp"+time.Now().String())

	err = os.Mkdir(tempDir, os.ModePerm)

	if err != nil {
		log.Fatalf("error creating video %q", err)
	}

	defer os.RemoveAll(tempDir)

	fmt.Println("converting video...")

	for i := 0; vid.Read(&frame); i++ {
		image, _ := frame.ToImage()
		image = convertFrame(image)

		fileName := fmt.Sprintf("abooba%d.png", i)
		file, err := os.Create(filepath.Join(tempDir, fileName))

		if err != nil {
			log.Fatalf("error creating video %q", err)

		}

		defer file.Close()

		err = png.Encode(file, image)
		if err != nil {
			log.Fatalf("error creating video %q", err)
		}
	}

	fmt.Println("video converted\nsaving video")

	frameRate := int(vid.Get(gocv.VideoCaptureFPS))

	cmd := exec.Command(
		"ffmpeg",
		"-framerate", strconv.Itoa(frameRate),
		"-i", tempDir+"/abooba%d.png",
		"-c:v", "libx264",
		"-r", strconv.Itoa(frameRate),
		"-pix_fmt", "yuv420p",
		filepath.Join(dest, time.Now().Local().Format(time.Stamp)+".mp4"),
	)

	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("video saved ✨")
}
