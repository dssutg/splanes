package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Shared player properties.
const (
	MaxPlayerHealth       = 100 // Maximum player health value
	PlayerMaxBombTickTime = 50  // Number of ticks required to reload bombs
)

// EntityType represents the type of an entity.
// Used to dispatch tick and render functions.
type EntityType uint8

const (
	EntityTypeNone       EntityType = iota // No entity type (unused slot)
	EntityTypePlayer                       // Player-controlled fighter plane
	EntityTypeEnemyPlane                   // Enemy aircraft
	EntityTypeBullet                       // Player bullet/projectile
	EntityTypeBomb                         // Player dropped bomb
	EntityTypeIsland                       // Island/obstacle
	EntityTypeExplosion                    // Explosion effect
	EntityTypeShip                         // Enemy ship
	EntityTypeHealer                       // Health power-up
	EntityTypeSubmarine                    // Enemy submarine
)

// Entity is the main game object structure.
// Contains position, rendering, physics, and game-specific properties.
type Entity struct {
	// Common properties
	Pos            sdl.Rect   // Entity position (X, Y) and size (W, H)
	Crop           sdl.Rect   // Sprite crop area in texture atlas
	Rotation       float32    // Rotation angle in degrees
	VelX           int32      // Horizontal velocity in pixels/tick
	VelY           int32      // Vertical velocity in pixels/tick
	Texture        TextureID  // Texture atlas ID to use
	Kind           EntityType // Entity type for dispatch
	Health         int32      // Current health points
	Damage         int32      // Damage dealt when colliding
	InitialFrameNo int32      // Initial animation frame
	Ticks          int32      // Generic timer counter
	State          uint8      // Entity state (varies by type)

	// For EntityTypeBullet: identifies who fired the bullet
	OwnerKind EntityType

	// For EntityTypePlayer: player-specific state
	Score     uint64 // Total score earned
	Distance  uint64 // Distance flown
	DeathTime int32  // Death animation timer
	BombTicks int32  // Bomb cooldown timer
	AccelX    int32  // Horizontal acceleration
	HasBombed bool   // Recently dropped a bomb

	// For EntityTypePlayer, EntityTypeEnemyPlane: attack state
	HasShot bool // Recently fired

	// For EntityTypeBomb: fall timer
	TicksDelta int32 // Ticks until impact
}

// EntityTableEntry defines the tick and render functions for an entity type.
// It also specifies the Z-index for render ordering (higher = drawn on top).
type EntityTableEntry struct {
	Tick   func(e *Entity) // Game logic update function
	Render func(e *Entity) // Rendering function
	ZIndex int32           // Render order (0 = back, 2 = front)
}

// EntityPool is the pre-allocated array of all game entities.
// Using a fixed pool avoids GC pressure during gameplay.
// Unused slots have Kind set to EntityTypeNone.
var (
	EntityPool [100]Entity // All entities in the game
	entityStub Entity      // Default entity when pool is full
)

// player is a pointer to the player entity in the pool.
// Used for quick access to player state.
var player *Entity

// NewEntity finds a free slot in EntityPool and initializes it with the given type.
// Returns a pointer to the new entity, or a stub if the pool is full.
// The caller must then call the specialized constructor function.
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

// Remove clears the entity, returning it to the pool.
// An entity cannot be used after calling this method,
// so entity handlers should return immediately after calling it.
func (e *Entity) Remove() {
	*e = Entity{} // zero entity
}

// Hurt applies damage to the entity if it is damageable.
// Only player, enemy planes, ships, and submarines can take damage.
func (e *Entity) Hurt(damage int32) {
	canHurt := false
	switch e.Kind {
	case
		EntityTypePlayer,
		EntityTypeEnemyPlane,
		EntityTypeShip,
		EntityTypeSubmarine:
		canHurt = true
	}

	if !canHurt {
		return
	}

	if e.Health < damage {
		e.Health = 0
	} else {
		e.Health -= damage
	}

	PlaySound(soundHurt, 100)
}

// RenderSprite renders the entity to the screen.
// It uses either normal or rotated rendering based on the Rotation field.
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

// RemoveAllEntities clears the entire entity pool.
// Used when restarting the game.
func RemoveAllEntities() {
	clear(EntityPool[:])
}

// EntityNoneCallback is a no-op function for entity types that don't need updates.
func EntityNoneCallback(*Entity) {}

// entityTable is the dispatch table mapping entity types to their
// tick/render functions and Z-index for render ordering.
// This allows O(1) lookups for entity updates and rendering.
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
