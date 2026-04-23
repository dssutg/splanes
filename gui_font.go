package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// fontCacheEntry holds a cached font and its size.
type fontCacheEntry struct {
	font *ttf.Font
	size int
}

// fontCache stores loaded fonts to avoid reloading for each render.
var fontCache []fontCacheEntry

// LoadFont returns a font of the given size, loading it if necessary.
// Fonts are cached to avoid the overhead of repeated file reads.
func LoadFont(size int) *ttf.Font {
	// Find the font with the required size in the loaded font cache.
	for _, entry := range fontCache {
		if entry.size == size {
			return entry.font
		}
	}

	font, err := ttf.OpenFont("assets/fonts/OpenSans-Bold.ttf", size)
	if err != nil {
		log.Fatal("can't open font file:", sdl.GetError())
	}

	// Add the new font to the font cache.
	fontCache = append(fontCache, fontCacheEntry{font: font, size: size})

	return font
}
