package gfx

import "core:fmt"
import "core:strings"

import SDL "vendor:sdl2"
import SDL_TTF "vendor:sdl2/ttf"

import "../util"

Font_Cache_Entry :: struct {
	font: ^SDL_TTF.Font,
	size: i32,
}

font_cache: [dynamic]Font_Cache_Entry

load_font :: proc(size: i32) -> ^SDL_TTF.Font {
	// Find the font with the required size in the loaded font cache
	for entry in font_cache {
		if entry.size == size {
			return entry.font
		}
	}

	font := SDL_TTF.OpenFont("assets/fonts/OpenSans-Bold.ttf", size)
	if font == nil {
		util.fatalf("can't open font file: %v", SDL.GetError())
	}

	// Add the new font to the font cache
	append(&font_cache, Font_Cache_Entry{font = font, size = size})

	return font
}

render_progress_bar :: proc(rect: SDL.Rect, stroke_size, progress: i32) {
	border_color := SDL.Color{0, 0, 0, 0}

	render_rect({rect.x, rect.y, rect.w, stroke_size}, border_color)
	render_rect({rect.x, rect.y + rect.h - stroke_size - 1, rect.w, stroke_size}, border_color)
	render_rect({rect.x, rect.y, stroke_size, rect.h}, border_color)
	render_rect({rect.x + rect.w - stroke_size - 1, rect.y, stroke_size, rect.h}, border_color)

	bar_color := SDL.Color{255, 0, 0, 255}
	switch {
	case progress > 75:
		bar_color = {0, 255, 0, 255}
	case progress > 30:
		bar_color = {255, 255, 0, 255}
	}

	bar_rect := SDL.Rect {
		rect.x + stroke_size,
		rect.y + stroke_size,
		(rect.w - stroke_size * 2) * progress / 100,
		rect.h - stroke_size * 2,
	}

	render_rect(bar_rect, bar_color)
}

text_buffer: strings.Builder

render_string :: proc(
	x, y, size: i32,
	color: SDL.Color,
	relative_to_window_center: bool,
	lineNo: i32,
	format: string,
	args: ..any,
) {
	strings.builder_reset(&text_buffer)

	fmt.sbprintf(&text_buffer, format, ..args)

	text := strings.to_cstring(&text_buffer)

	if len(text) == 0 {
		return
	}

	font := load_font(size)

	font_surface := SDL_TTF.RenderText_Blended(font, text, color)
	if font_surface == nil {
		util.fatalf("can't get font surface: %v", SDL.GetError())
	}
	defer SDL.FreeSurface(font_surface)

	texture := SDL.CreateTextureFromSurface(renderer, font_surface)
	defer SDL.DestroyTexture(texture)

	src := SDL.Rect{0, 0, font_surface.w, font_surface.h}

	dest := SDL.Rect{x, y, font_surface.w, font_surface.h}
	if relative_to_window_center {
		dest.x = (Window_Width - font_surface.w) / 2
		dest.y = (Window_Height + (font_surface.h + 50) * lineNo) / 2
	}

	SDL.RenderCopy(renderer, texture, &src, &dest)
}

render_health_bar :: proc(health: i32) {
	scale: i32 : 2

	x: i32 : 20
	y: i32 : 20

	crop: SDL.Rect = {364, 199, 130, 14}

	render_sprite(0, {x, y, crop.w * scale, crop.h * scale}, crop)

	crop.x = 494
	crop.w = 7

	if health < 0 {
		return
	}

	max_elems: i32 : 18
	elem_num := min(max_elems, health * max_elems / 100)

	for i in 0 ..< elem_num {
		render_sprite(0, {x + i * crop.w * scale, y, crop.w * scale, crop.h * scale}, crop)
	}
}

render_small_logo :: proc() {
	scale: i32 : 1

	crop_w: i32 : 115
	crop_h: i32 : 58

	w := crop_w * scale
	h := crop_h * scale

	render_sprite(0, {Window_Width - w, 0, w, h}, {700, 401, crop_w, crop_h})
}
