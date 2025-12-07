#pragma once

#include "../util/util.h"

// SFX and Voices
enum {
  SoundHurt,
  SoundExplosion1,
  SoundCount,
};

// Music tracks
enum {
  MusicBackground0,
  MusicCount,
};

void PlaySound(i32 trackID, i32 volume, i32 isMusic);
void InitSoundManager(void);
