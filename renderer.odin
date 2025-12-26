package main

import SDL "vendor:sdl2"
import SDL_Image "vendor:sdl2/image"

Window_Delay_Milliseconds :: 50

Window_Scale :: 1

Window_Width :: 800 * Window_Scale
Window_Height :: 600 * Window_Scale
Window_Rect :: SDL.Rect{0, 0, Window_Width, Window_Height}

Tile_Size :: 32

renderer: ^SDL.Renderer
textures: [16]^SDL.Texture
window: ^SDL.Window
window_title := "Splanes"

render_sprite :: proc(texture: i32, dest, src: SDL.Rect) {
	dest_copy := dest
	src_copy := src
	SDL.RenderCopy(renderer, textures[texture], &src_copy, &dest_copy)
}

load_texture :: proc(filename: string) -> ^SDL.Texture {
	bitmap := SDL_Image.Load(cstring(raw_data(filename)))
	if bitmap == nil {
		fatalf("can't load %v: %v", filename, SDL_Image.GetError())
	}
	defer SDL.FreeSurface(bitmap)

	texture := SDL.CreateTextureFromSurface(renderer, bitmap)
	if texture == nil {
		fatalf("can't create texture from %v: %v", filename, SDL.GetError())
	}

	return texture
}

load_textures :: proc() {
	textures[0] = load_texture("assets/sprites/sprites.png")
}

render_rect :: proc(rect: SDL.Rect, color: SDL.Color) {
	rect_copy := rect
	SDL.SetRenderDrawColor(renderer, color.r, color.g, color.b, color.a)
	SDL.RenderFillRect(renderer, &rect_copy)
}
