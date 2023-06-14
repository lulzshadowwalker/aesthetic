package config

import (
	"flag"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

var (
	src       string
	dest      string
	retro     bool
	grayscale bool
)

type FileType int

const (
	Image FileType = iota
	Video
)

func ParseFlags() {
	// TODO add a sampling rate flag 
	flag.StringVar(&src, "src", "",
		`Source image or video to be converted. *mandatory*
Currently supports most image formats and mp4 for video sources`,
	)

	flag.BoolVar(&retro, "retro", false, "Applies a *retro* aesthetic to the video by stalling the frames")
	help := flag.Bool("help", false, "List all commands")
	flag.BoolVar(&grayscale, "grayscale", false, "Decides whether the result should be grayscale (black&white) or preserve the original colors of the image")

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

	SourceType()
}

func Src() string {
	return src
}

func Dest() string {
	// TODO assert Dest is a directory
	return dest
}

func SourceType() FileType {
	info, err := os.Stat(src)

	if err != nil {
		log.Fatalf("error reading file information %q", err)
	} else if info.IsDir() {
		log.Fatalf("%q has to be an image or video file not a directory", src)
	}

	switch mt := mime.TypeByExtension(filepath.Ext(src)); {
	case strings.HasPrefix(mt, "image/"):
		return Image
	case strings.HasPrefix(mt, "video/"):
		return Video
	default:
		log.Fatal("Unsupported source type")
		return -1
	}
}

func Retro() bool {
	return retro
}

func Grayscale() bool {
	return grayscale
}
