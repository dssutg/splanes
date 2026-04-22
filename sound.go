package main

import (
	"log"

	"github.com/veandco/go-sdl2/mix"
)

var (
	soundHurt       *mix.Chunk
	soundExplosion1 *mix.Chunk
)

var musicBackground0 *mix.Music

func NewSoundEffect(filename string) *mix.Chunk {
	sound, err := mix.LoadWAV(filename)
	if err != nil {
		log.Fatalf("can't load %v: %v", filename, err)
	}
	return sound
}

func NewMusicTrack(filename string) *mix.Music {
	music, err := mix.LoadMUS(filename)
	if err != nil {
		log.Fatalf("can't load %v: %v", filename, err)
	}
	return music
}

func PlaySound(sound *mix.Chunk, volume int) {
	sound.Volume(volume)
	sound.Play(-1, 0)
}

func PlayMusic(music *mix.Music, volume int) {
	mix.VolumeMusic(volume)
	if err := music.Play(-1); err != nil {
		log.Fatal("can't play music:", err)
	}
}

func InitSoundManager() {
	musicBackground0 = NewMusicTrack("assets/music/bg_0.ogg")

	soundHurt = NewSoundEffect("assets/sfx/hurt.wav")
	soundExplosion1 = NewSoundEffect("assets/sfx/explosion1.wav")
}
