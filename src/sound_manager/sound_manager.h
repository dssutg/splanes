#pragma once

#include "../util/util.h"

// SFX and Voices
typedef enum SoundID : u8 {
  SoundHurt,
  SoundExplosion1,
  SoundCount,
} SoundID;

// Music tracks
typedef enum MusicID : u8 {
  MusicBackground0,
  MusicCount,
} MusicID;

void PlaySound(SoundID soundID, i32 volume);
void PlayMusic(MusicID musicID, i32 volume);
void InitSoundManager(void);
