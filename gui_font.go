package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type fontCacheEntry struct {
	font *ttf.Font
	size int
}

var fontCache []fontCacheEntry

func LoadFont(size int) *ttf.Font {
	// Find the font with the required size in the loaded font cache
	for _, entry := range fontCache {
		if entry.size == size {
			return entry.font
		}
	}

	font, err := ttf.OpenFont("assets/fonts/OpenSans-Bold.ttf", size)
	if err != nil {
		log.Fatal("can't open font file:", sdl.GetError())
	}

	// Add the new font to the font cache
	fontCache = append(fontCache, fontCacheEntry{font: font, size: size})

	return font
}
