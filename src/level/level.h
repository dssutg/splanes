#pragma once

#include "../lib/std.h"

#include "../entity/entity.h"
#include "../entity/entity_player.h"

class Level {
  public:
  // All entities in game
  std::list<Entity *> entities = {};

  // Reference to the player in entity list
  std::optional<Player *> player = std::nullopt;

  Level(int32_t layerWidth, int32_t layerHeight);

  virtual ~Level();

  void destroy();
  void reset();
  void addEntity(Entity &entity);
  void removeAllEntities();
  void tick();
  void render();

  private:
  int32_t layerWidth = 1;
  int32_t layerHeight = 1;
  int32_t layerTileSize = 1;

  bool needZIndexSort = false;

  std::array<int32_t, 2> layers = {};

  void renderLayer(int32_t offsetY);
};
