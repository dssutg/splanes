#include "../lib/std.h"

#include "sound_manager.h"

#include "../util/util.h"

static std::unordered_map<SoundID, Mix_Chunk *> sounds;
static std::unordered_map<MusicID, Mix_Music *> musicTracks;

static Mix_Chunk *NewSoundEffect(const std::string &filename) {
  auto soundEffect = Mix_LoadWAV(filename.c_str());

  if (soundEffect == nullptr) {
    Fatal(std::format("can't load {}: {}", filename, SDL_GetError()));
  }

  return soundEffect;
}

static Mix_Music *NewMusicTrack(const std::string filename) {
  auto music = Mix_LoadMUS(filename.c_str());

  if (music == nullptr) {
    Fatal(std::format("can't load {}: {}", filename, SDL_GetError()));
  }

  return music;
}

void PlaySound(SoundID soundID, int32_t volume) {
  const auto soundEffect = sounds[soundID];
  Mix_VolumeChunk(soundEffect, volume);
  Mix_PlayChannel(-1, soundEffect, 0);
}

void PlayMusic(MusicID musicID, int32_t volume) {
  const auto music = musicTracks[musicID];
  Mix_VolumeMusic(volume);

  if (Mix_PlayMusic(music, -1) == -1) {
    Fatal(SDL_GetError());
  }
}

void InitSoundManager() {
  musicTracks[MusicID::Background0] = NewMusicTrack("assets/music/bg_0.ogg");

  sounds[SoundID::Hurt] = NewSoundEffect("assets/sfx/hurt.wav");
  sounds[SoundID::Explosion1] = NewSoundEffect("assets/sfx/explosion1.wav");
}

void FreeSoundManager() {
  for (const auto &[_, sound] : sounds) {
    Mix_FreeChunk(sound);
  }
  sounds.clear();

  for (const auto &[_, musicTrack] : musicTracks) {
    Mix_FreeMusic(musicTrack);
  }
  musicTracks.clear();
}
