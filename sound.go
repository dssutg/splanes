package main

import (
	"log"

	"github.com/veandco/go-sdl2/mix"
)

// Loaded sound effect chunks.
var (
	soundHurt       *mix.Chunk // Player/enemy hurt sound
	soundExplosion1 *mix.Chunk // Explosion sound
)

// Loaded music tracks.
var musicBackground0 *mix.Music // Background music

// NewSoundEffect loads an audio file and returns a chunk for one-shot playback.
func NewSoundEffect(filename string) *mix.Chunk {
	sound, err := mix.LoadWAV(filename)
	if err != nil {
		log.Fatalf("can't load %v: %v", filename, err)
	}
	return sound
}

// NewMusicTrack loads an OGG file and returns a music track for streaming playback.
func NewMusicTrack(filename string) *mix.Music {
	music, err := mix.LoadMUS(filename)
	if err != nil {
		log.Fatalf("can't load %v: %v", filename, err)
	}
	return music
}

// PlaySound plays a sound effect at the specified volume (0-128).
func PlaySound(sound *mix.Chunk, volume int) {
	sound.Volume(volume)
	_, _ = sound.Play(-1, 0)
}

// PlayMusic starts playing a music track at the specified volume (0-128).
// The music loops indefinitely.
func PlayMusic(music *mix.Music, volume int) {
	mix.VolumeMusic(volume)
	_ = music.Play(-1)
}

// InitSoundManager loads all game audio files into memory.
func InitSoundManager() {
	musicBackground0 = NewMusicTrack("assets/music/bg_0.ogg")

	soundHurt = NewSoundEffect("assets/sfx/hurt.wav")
	soundExplosion1 = NewSoundEffect("assets/sfx/explosion1.wav")
}
