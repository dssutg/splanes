#+feature dynamic-literals
package snd

import SDL "vendor:sdl2"
import SDL_Mixer "vendor:sdl2/mixer"

import "../util"

// SFX and Voices
Sound_ID :: enum u8 {
	Hurt,
	Explosion1,
}

// Music tracks
Music_ID :: enum u8 {
	Background0,
}

sounds: map[Sound_ID]^SDL_Mixer.Chunk
music_tracks: map[Music_ID]^SDL_Mixer.Music

new_sound_effect :: proc(filename: string) -> ^SDL_Mixer.Chunk {
	sound := SDL_Mixer.LoadWAV(cstring(raw_data(filename)))
	if sound == nil {
		util.fatalf("can't load %v: %v", filename, SDL.GetError())
	}
	return sound
}

new_music_track :: proc(filename: string) -> ^SDL_Mixer.Music {
	music := SDL_Mixer.LoadMUS(cstring(raw_data(filename)))
	if music == nil {
		util.fatalf("can't load %v: %v", filename, SDL.GetError())
	}
	return music
}

play_sound :: proc(sound_ID: Sound_ID, volume: i32) {
	sound := sounds[sound_ID]
	SDL_Mixer.VolumeChunk(sound, volume)
	SDL_Mixer.PlayChannel(-1, sound, 0)
}

play_music :: proc(music_ID: Music_ID, volume: i32) {
	music := music_tracks[music_ID]
	SDL_Mixer.VolumeMusic(volume)
	if SDL_Mixer.PlayMusic(music, -1) == -1 {
		util.fatalf("can't play music: %v", SDL.GetError())
	}
}

init_sound_manager :: proc() {
	music_tracks = {
		.Background0 = new_music_track("assets/music/bg_0.ogg"),
	}

	sounds = {
		.Hurt       = new_sound_effect("assets/sfx/hurt.wav"),
		.Explosion1 = new_sound_effect("assets/sfx/explosion1.wav"),
	}
}
