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
	EntityTypeSubmarine
)

// Entity fat struct
type Entity struct {
	// Common
	Pos            sdl.Rect   // entity position and size
	Crop           sdl.Rect   // current entity sprite
	Rotation       float32    // rotation angle in degrees
	VelX           int32      // entity horizontal velocity
	VelY           int32      // entity vertical velocity
	Texture        TextureID  // current entity texture ID
	Kind           EntityType // entity kind
	Health         int32      // entity health
	Damage         int32      // entity damage (same unit as health)
	InitialFrameNo int32      // initial frame number
	Ticks          int32      // generic timer in ticks
	State          uint8      // entity state

	// For EntityTypeBullet
	OwnerKind EntityType // entity type that spawned this bullet

	// For EntityTypePlayer
	Score     uint64 // player's score
	Distance  uint64 // the distance the player has flown
	DeathTime int32  // death animation time until removal
	BombTicks int32  // bomb delay timer
	AccelX    int32  // player horizontal acceleration
	HasBombed bool   // has the entity recently bombed

	// For EntityTypePlayer, EntityTypeEnemyPlane
	HasShot bool // has the entity recently shot

	// For EntityTypeBomb
	TicksDelta int32 // ticks delta until fallen onto the target
}

// Entity dispatch table entry
type EntityTableEntry struct {
	Tick   func(e *Entity) // update entity logic per frame
	Render func(e *Entity) // render entity
	ZIndex int32           // entity's z-index (render order, higher - rendered above others)
}

var (
	EntityPool [100]Entity // all entities in game
	entityStub Entity      // dummy entity returned when entity could not be created
)

var player *Entity // player reference in entity list

func NewEntity(etype EntityType) *Entity {
	entity := &entityStub // default to stub

	// Find the free position in entity pool,
	// otherwise keep using stub.
	for i := range EntityPool {
		if EntityPool[i].Kind == EntityTypeNone {
			entity = &EntityPool[i]
			break
		}
	}

	// Initialize the found entity.
	// Needs specialized constructors after that.
	*entity = Entity{} // zero entity
	entity.Kind = etype

	return entity
}

func (e *Entity) Remove() {
	*e = Entity{} // zero entity
}

func (e *Entity) Hurt(damage int32) {
	switch e.Kind {
	case EntityTypePlayer:
	case EntityTypeEnemyPlane:
	case EntityTypeShip:
	case EntityTypeSubmarine:
		break
	default:
		return
	}

	if e.Health < damage {
		e.Health = 0
	} else {
		e.Health -= damage
	}

	PlaySound(soundHurt, 100)
}

func (e *Entity) RenderSprite() {
	if e.Rotation == 0 {
		RenderSprite(e.Texture, e.Pos, e.Crop)
		return
	}

	center := sdl.Point{
		X: e.Pos.W / 2,
		Y: e.Pos.H / 2,
	}

	RenderSpriteEx(
		e.Texture,
		e.Pos,
		e.Crop,
		float64(e.Rotation),
		&center,
		0,
	)
}

func RemoveAllEntities() {
	clear(EntityPool[:])
}

// EntityNoneCallback is a function that does nothing.
func EntityNoneCallback(*Entity) {}

var entityTable = map[EntityType]EntityTableEntry{
	EntityTypeNone:       {Tick: EntityNoneCallback, Render: EntityNoneCallback, ZIndex: 0},
	EntityTypePlayer:     {Tick: PlayerTick, Render: PlayerRender, ZIndex: 2},
	EntityTypeEnemyPlane: {Tick: EnemyPlaneTick, Render: EnemyPlaneRender, ZIndex: 2},
	EntityTypeBullet:     {Tick: BulletTick, Render: BulletRender, ZIndex: 2},
	EntityTypeBomb:       {Tick: BombTick, Render: BombRender, ZIndex: 2},
	EntityTypeIsland:     {Tick: IslandTick, Render: IslandRender, ZIndex: 0},
	EntityTypeExplosion:  {Tick: ExplosionTick, Render: ExplosionRender, ZIndex: 2},
	EntityTypeShip:       {Tick: ShipTick, Render: ShipRender, ZIndex: 1},
	EntityTypeHealer:     {Tick: HealerTick, Render: HealerRender, ZIndex: 2},
	EntityTypeSubmarine:  {Tick: SubmarineTick, Render: SubmarineRender, ZIndex: 1},
}
