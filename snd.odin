#+feature dynamic-literals
package main

import SDL "vendor:sdl2"
import SDL_Mixer "vendor:sdl2/mixer"

SoundHurt: ^SDL_Mixer.Chunk
SoundExplosion1: ^SDL_Mixer.Chunk

MusicBackground0: ^SDL_Mixer.Music

new_sound_effect :: proc(filename: string) -> ^SDL_Mixer.Chunk {
	sound := SDL_Mixer.LoadWAV(cstring(raw_data(filename)))
	if sound == nil {
		fatalf("can't load %v: %v", filename, SDL.GetError())
	}
	return sound
}

new_music_track :: proc(filename: string) -> ^SDL_Mixer.Music {
	music := SDL_Mixer.LoadMUS(cstring(raw_data(filename)))
	if music == nil {
		fatalf("can't load %v: %v", filename, SDL.GetError())
	}
	return music
}

play_sound :: proc(sound: ^SDL_Mixer.Chunk, volume: i32) {
	SDL_Mixer.VolumeChunk(sound, volume)
	SDL_Mixer.PlayChannel(-1, sound, 0)
}

play_music :: proc(music: ^SDL_Mixer.Music, volume: i32) {
	SDL_Mixer.VolumeMusic(volume)
	if SDL_Mixer.PlayMusic(music, -1) == -1 {
		fatalf("can't play music: %v", SDL.GetError())
	}
}

init_sound_manager :: proc() {
	MusicBackground0 = new_music_track("assets/music/bg_0.ogg")

	SoundHurt = new_sound_effect("assets/sfx/hurt.wav")
	SoundExplosion1 = new_sound_effect("assets/sfx/explosion1.wav")
}
