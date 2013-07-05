/*
 * Copyright 2013, Robert Bieber
 *
 * This file is part of drones.
 *
 * drones is free software: you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * drones is distributed in the hope that it will be useful,
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with drones.  If not, see <http://www.gnu.org/licenses/>.
 */

// Package window defines a window UI element.
package window

import (
	"github.com/bieber/drones/ui"
	"github.com/neagix/Go-SDL/sdl"
)

// Window draws a rectangular area on the screen with a border and
// renders child elements.
type Window struct {
	x           uint
	y           uint
	W           uint
	H           uint
	BorderWidth uint
	BGColor     sdl.Color
	BorderColor sdl.Color
	// Callback to execute if the window intercepts an escape keypress.
	OnEscape func()
}

func New(x, y, w, h, borderWidth uint, bgColor, borderColor sdl.Color) *Window {
	return &Window{
		x:           x,
		y:           y,
		W:           w,
		H:           h,
		BorderWidth: borderWidth,
		BGColor:     bgColor,
		BorderColor: borderColor,
		OnEscape:    nil,
	}
}

func (w *Window) Draw(screen *sdl.Surface) {
	screen.FillRect(
		&sdl.Rect{
			X: int16(w.x),
			Y: int16(w.y),
			W: uint16(w.W),
			H: uint16(w.H),
		},
		ui.ColorToInt(w.BorderColor),
	)
	screen.FillRect(
		&sdl.Rect{
			X: int16(w.x + w.BorderWidth),
			Y: int16(w.y + w.BorderWidth),
			W: uint16(w.W - w.BorderWidth*2),
			H: uint16(w.H - w.BorderWidth*2),
		},
		ui.ColorToInt(w.BGColor),
	)
}

func (w *Window) HandleEvent(event interface{}) bool {
	switch event.(type) {
	case sdl.KeyboardEvent:
		if w.OnEscape != nil && ui.IsKeyDown(event, sdl.K_ESCAPE) {
			w.OnEscape()
			return true
		}
	case sdl.MouseButtonEvent:
		return true
	}
	return false
}

func (w *Window) SetPos(x, y uint) {
	w.x, w.y = x, y
}

func (w *Window) GetPos(x, y uint) {
	x, y = w.x, w.y
	return
}
