package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Shared player properties
const (
	MaxPlayerHealth       = 100
	PlayerMaxBombTickTime = 50
)

// Entity types
type EntityType uint8

const (
	EntityTypeNone EntityType = iota
	EntityTypePlayer
	EntityTypeEnemyPlane
	EntityTypeBullet
	EntityTypeBomb
	EntityTypeIsland
	EntityTypeExplosion
	EntityTypeShip
	EntityTypeHealer
)

// Entity fat struct
type Entity struct {
	// Common
	pos      sdl.Rect
	crop     sdl.Rect
	xa       int32
	ya       int32
	texture  int
	etype    EntityType
	health   int32
	damage   int32
	data     int32
	tickTime int32

	// EntityBullet
	ownerType EntityType // entity type that spawned this bullet

	// EntityPlayer
	score        uint64
	distance     uint64
	hasShot      bool
	hasBombed    bool
	deathTime    int32
	bombTickTime int32
}

// Entity polymorphic dispatch table entry
type EntityTableEntry struct {
	Tick   func(e *Entity)
	Render func(e *Entity)
	ZIndex int32
}

var (
	entityPool [100]Entity // All entities in game
	entityStub Entity      // Dummy entity returned when entity could not be created
)

// Reference to the player in entity list
var player *Entity

func newEntity(etype EntityType) *Entity {
	entity := &entityStub // default to stub
	for i := range entityPool {
		if entityPool[i].etype == EntityTypeNone {
			entity = &entityPool[i]
			break
		}
	}
	*entity = Entity{} // zero entity
	entity.etype = etype
	return entity
}

func removeEntity(e *Entity) {
	*e = Entity{} // zero entity
}

func hurtEntity(e *Entity, damage int32) {
	if e.etype != EntityTypePlayer && e.etype != EntityTypeEnemyPlane && e.etype != EntityTypeShip {
		return
	}

	if e.health < damage {
		e.health = 0
	} else {
		e.health -= damage
	}

	playSound(soundHurt, 100)
}

func renderEntitySprite(e *Entity) {
	renderSprite(e.texture, e.pos, e.crop)
}

func removeAllEntities() {
	clear(entityPool[:])
}

func entityNoneCallback(_ *Entity) {
	// dummy entity handler that does nothing
}

var entityTable = map[EntityType]EntityTableEntry{
	EntityTypeNone:       {Tick: entityNoneCallback, Render: entityNoneCallback, ZIndex: 0},
	EntityTypePlayer:     {Tick: playerTick, Render: playerRender, ZIndex: 2},
	EntityTypeEnemyPlane: {Tick: enemyPlaneTick, Render: enemyPlaneRender, ZIndex: 2},
	EntityTypeBullet:     {Tick: bulletTick, Render: bulletRender, ZIndex: 2},
	EntityTypeBomb:       {Tick: bombTick, Render: bombRender, ZIndex: 2},
	EntityTypeIsland:     {Tick: islandTick, Render: islandRender, ZIndex: 0},
	EntityTypeExplosion:  {Tick: explosionTick, Render: explosionRender, ZIndex: 2},
	EntityTypeShip:       {Tick: shipTick, Render: shipRender, ZIndex: 1},
	EntityTypeHealer:     {Tick: healerTick, Render: healerRender, ZIndex: 2},
}
