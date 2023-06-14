package main

import (
	"github.com/lulzshadowwalker/aesthetic/internal/config"
	"github.com/lulzshadowwalker/aesthetic/pkg/aesthetic"
)

func main() {
	config.ParseFlags()

	if st := config.SourceType(); st == config.Image {
		aesthetic.ConvertImage(config.Src(), config.Dest())
	}

}
