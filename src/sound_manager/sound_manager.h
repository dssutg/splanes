#pragma once

#include "../lib/std.h"

// SFX and Voices
enum class SoundID : uint8_t {
  Hurt,
  Explosion1,
};

// Music tracks
enum class MusicID : uint8_t {
  Background0,
};

void PlaySound(SoundID soundID, int32_t volume);
void PlayMusic(MusicID musicID, int32_t volume);
void InitSoundManager();
void FreeSoundManager();
