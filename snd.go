package main

import (
	"github.com/veandco/go-sdl2/mix"
)

var (
	soundHurt       *mix.Chunk
	soundExplosion1 *mix.Chunk

	musicBackground0 *mix.Music
)

func newSoundEffect(filename string) *mix.Chunk {
	sound, err := mix.LoadWAV(filename)
	if err != nil {
		fatalf("can't load %v: %v", filename, err)
	}
	return sound
}

func newMusicTrack(filename string) *mix.Music {
	music, err := mix.LoadMUS(filename)
	if err != nil {
		fatalf("can't load %v: %v", filename, err)
	}
	return music
}

func playSound(sound *mix.Chunk, volume int) {
	sound.Volume(volume)
	sound.Play(-1, 0)
}

func playMusic(music *mix.Music, volume int) {
	mix.VolumeMusic(volume)
	if err := music.Play(-1); err != nil {
		fatalf("can't play music: %v", err)
	}
}

func initSoundManager() {
	musicBackground0 = newMusicTrack("assets/music/bg_0.ogg")

	soundHurt = newSoundEffect("assets/sfx/hurt.wav")
	soundExplosion1 = newSoundEffect("assets/sfx/explosion1.wav")
}
