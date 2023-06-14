package config

import (
	"errors"
	"flag"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

var (
	src            string
	dest           string
	grayscale      bool
	sampleInterval int
	dotSize        int
	whiteNoise     bool
)

type FileType int

const (
	Image FileType = iota
	Video
)

// parses and validates the flags
func ParseFlags() {
	// TODO add a sampling rate flag
	// todo add dot size flag
	// todo add background flag
	flag.StringVar(&src, "src", "",
		`Source image or video to be converted. *mandatory*
Currently supports most image formats and mp4 for video sources`,
	)

	help := flag.Bool("help", false, "Lists all commands")
	flag.BoolVar(&grayscale, "grayscale", false, "Decides whether the result should be grayscale (black & white) or preserve the original colors of the image")
	flag.IntVar(&dotSize, "dot-size", 4, "the upper bound for the dot size")
	flag.BoolVar(&whiteNoise, "white-noise", false, "uhm.. white noise")

	flag.IntVar(&sampleInterval, "sample-interval", 5,
		`Sampling rate for the given image or video
Dictates how many pixels are going to be read
Essentially corresponds to the gap or noise size (higher sampling rate, less dots or info to represent the image, larger gaps)
`)

	var defaultDest string
	homeDir, err := os.UserHomeDir()
	if err != nil {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("error optaining the default output destination %q", err)
		}
		defaultDest = filepath.Join(cwd, "desktop")

	} else {
		defaultDest = filepath.Join(homeDir, "desktop")
	}
	flag.StringVar(&dest, "dest", defaultDest, "Destination of the converted image or video")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	} else if src == "" {
		log.Fatalf("\"--source\" must be specified")
	}

	_, err = SourceType()

	if err != nil {
		log.Fatal(err)
	}

	destInfo, err := os.Stat(dest)

	if err != nil {
		log.Fatal(err)
	}

	if !destInfo.IsDir() {
		log.Fatalf("%q has to be a directory", dest)
	}
}

func Source() string {
	return src
}

func Destination() string {
	return dest
}

func SourceType() (FileType, error) {
	info, err := os.Stat(src)

	if err != nil {
		log.Fatalf("error reading file information %q", err)
	} else if info.IsDir() {
		log.Fatalf("%q has to be an image or video file not a directory", src)
	}

	switch mt := mime.TypeByExtension(filepath.Ext(src)); {
	case strings.HasPrefix(mt, "image/"):
		return Image, nil
	case strings.HasPrefix(mt, "video/"):
		return Video, nil
	default:
		return -1, errors.New("unsupported file type.\nSource file has to be either an image or a video")
	}
}

func Grayscale() bool {
	return grayscale
}

func SampleInterval() int {
	return sampleInterval
}

func DotSize() int {
	return dotSize
}

func WhiteNoise() bool {
	return whiteNoise
}
