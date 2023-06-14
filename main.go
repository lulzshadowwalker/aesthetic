package main

import (
	"github.com/lulzshadowwalker/aesthetic/internal/config"
	"github.com/lulzshadowwalker/aesthetic/pkg/aesthetic"
)

func main() {
	config.ParseFlags()

	if st, _ := config.SourceType(); st == config.Image {
		aesthetic.ConvertImage()
	} else if st == config.Video {
		aesthetic.ConvertVideo()
	}
}
