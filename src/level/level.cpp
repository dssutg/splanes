#include "../lib/std.h"

#include "level.h"

#include "../entity/entity_enemy_plane.h"
#include "../entity/entity_ship.h"
#include "../entity/entity_island.h"
#include "../entity/entity_healer.h"

#include "../sound_manager/sound_manager.h"
#include "../renderer/renderer.h"

Level::Level(int32_t layerWidth_, int32_t layerHeight_) {
  this->layerWidth = layerWidth_;
  this->layerHeight = layerHeight_;
  this->layerTileSize = TileSize;

  Level::reset();

  PlayMusic(MusicID::Background0, 70);
}

Level::~Level() {
  destroy();
}

void Level::destroy() {
  removeAllEntities();
}

void Level::reset() {
  destroy();

  player = new Player();
  addEntity(**player);

  layers[0] = -layerHeight;
  layers[1] = 0;
}

void Level::addEntity(Entity &entity) {
  entities.push_front(&entity);
  needZIndexSort = true;
}

void Level::removeAllEntities() {
  for (auto &entity : entities) {
    delete entity;
  }
  entities.clear();
  player = std::nullopt;
}

void Level::tick() {
  if (rand() % 20 == 0) {
    addEntity(*new EnemyPlane());
  }
  if (rand() % 80 == 0) {
    addEntity(*new Ship());
  }
  if (rand() % 30 == 0) {
    addEntity(*new Island());
  }
  if (rand() % 100 == 0) {
    addEntity(*new Healer());
  }

  layers[0] += 10;
  layers[1] += 10;
  if (layers[1] >= layerHeight) {
    layers[1] = layers[0] - layerHeight;
    std::swap(layers[0], layers[1]);
  }

  for (auto &entity : entities) {
    entity->tick();
  }

  // Delete all entities marked as removed
  for (auto it = entities.begin(); it != entities.end();) {
    auto entity = *it;
    if (entity->removed) {
      it = entities.erase(it);
      delete entity;
    } else {
      ++it;
    }
  }
}

void Level::renderLayer(int32_t offsetY) {
  const auto tileSize = layerTileSize;

  const auto tileWidth = (layerWidth + tileSize - 1) / tileSize;
  const auto tileHeight = (layerHeight + tileSize - 1) / tileSize;

  for (int32_t tileY = 0; tileY < tileHeight; tileY++) {
    for (int32_t tileX = 0; tileX < tileWidth; tileX++) {
      const auto x = tileX * tileSize;
      const auto y = tileY * tileSize + offsetY;
      RenderSprite(0,
                   {.x = x, .y = y, .w = tileSize, .h = tileSize},
                   {.x = 265, .y = 364, .w = 32, .h = 32});
    }
  }
}

void Level::render() {
  // Render water layers that constantly exchange making the illusion of
  // infinite map scrolling
  renderLayer(layers[0]);
  renderLayer(layers[1]);

  if (needZIndexSort) {
    // Sort entities by their z-index
    entities.sort([](const Entity *a, const Entity *b) {
      return a->getZIndex() < b->getZIndex();
    });
    needZIndexSort = false;
  }

  // Render entitites
  for (auto &entity : entities) {
    entity->render();
  }
}
