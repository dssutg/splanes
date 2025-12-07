#include <SDL2/SDL_mixer.h>

#include "sound_manager.h"

#include "../util/util.h"

static Mix_Chunk *sounds[SoundCount];
static Mix_Music *musicTracks[MusicCount];

static Mix_Chunk *NewSoundEffect(const char *filename) {
  Mix_Chunk *soundEffect = Mix_LoadWAV(filename);

  if (soundEffect == NULL) {
    Fatalf("can't load %s: %s", filename, SDL_GetError());
  }

  return soundEffect;
}

static Mix_Music *NewMusicTrack(const char *filename) {
  Mix_Music *music = Mix_LoadMUS(filename);

  if (music == NULL) {
    Fatalf("can't load %s: %s", filename, SDL_GetError());
  }

  return music;
}

void PlaySound(i32 trackID, i32 volume, i32 isMusic) {
  if (isMusic) {
    Mix_Music *music = musicTracks[trackID];
    Mix_VolumeMusic(volume);

    if (Mix_PlayMusic(music, -1) == -1) {
      Fatalf("%s", SDL_GetError());
    }
  } else {
    Mix_Chunk *soundEffect = sounds[trackID];
    Mix_VolumeChunk(soundEffect, volume);
    Mix_PlayChannel(-1, soundEffect, 0);
  }
}

void InitSoundManager(void) {
  musicTracks[MusicBackground0] = NewMusicTrack("assets/music/bg_0.ogg");

  sounds[SoundHurt] = NewSoundEffect("assets/sfx/hurt.wav");
  sounds[SoundExplosion1] = NewSoundEffect("assets/sfx/explosion1.wav");
}
