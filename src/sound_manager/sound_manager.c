#include <SDL2/SDL_mixer.h>

#include "sound_manager.h"

#include "../util/util.h"

static Mix_Chunk *sounds[SoundCount];
static Mix_Music *musicTracks[MusicCount];

static Mix_Chunk *NewSoundEffect(const char *const filename) {
  auto soundEffect = Mix_LoadWAV(filename);

  if (soundEffect == nullptr) {
    Fatalf("can't load %s: %s", filename, SDL_GetError());
  }

  return soundEffect;
}

static Mix_Music *NewMusicTrack(const char *const filename) {
  auto music = Mix_LoadMUS(filename);

  if (music == nullptr) {
    Fatalf("can't load %s: %s", filename, SDL_GetError());
  }

  return music;
}

void PlaySound(SoundID soundID, i32 volume) {
  const auto soundEffect = sounds[soundID];
  Mix_VolumeChunk(soundEffect, volume);
  Mix_PlayChannel(-1, soundEffect, 0);
}

void PlayMusic(MusicID musicID, i32 volume) {
  const auto music = musicTracks[musicID];
  Mix_VolumeMusic(volume);

  if (Mix_PlayMusic(music, -1) == -1) {
    Fatalf("%s", SDL_GetError());
  }
}

void InitSoundManager(void) {
  musicTracks[MusicBackground0] = NewMusicTrack("assets/music/bg_0.ogg");

  sounds[SoundHurt] = NewSoundEffect("assets/sfx/hurt.wav");
  sounds[SoundExplosion1] = NewSoundEffect("assets/sfx/explosion1.wav");
}
