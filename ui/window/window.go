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

// Defaults
var bgColor sdl.Color = sdl.Color{150, 150, 150, 0}
var borderColor sdl.Color = sdl.Color{100, 100, 100, 0}

const borderWidth = 5

// Window draws a rectangular area on the screen with a border and
// renders child elements.
type Window struct {
	children []ui.Widget

	// Window position and size.
	X int16
	Y int16
	W uint16
	H uint16

	// Width of the border around the window.
	BorderWidth uint16

	// Colors of the background and border.
	BGColor     sdl.Color
	BorderColor sdl.Color

	//Callback to execute if the window intercepts an escape keypress.
	OnEscape func()
}

func New(x, y int16, w, h uint16) *Window {
	return &Window{
		X:           x,
		Y:           y,
		W:           w,
		H:           h,
		BorderWidth: borderWidth,
		BGColor:     bgColor,
		BorderColor: borderColor,
		OnEscape:    nil,
		children:    make([]ui.Widget, 0, 5),
	}
}

// Adds a child element to the Window at the desired relative
// coordinates.
func (w *Window) AddChild(child ui.Widget, x, y int16) {
	w.children = append(w.children, child)
	child.SetPos(w.X+x, w.Y+y)
}

func (w *Window) Draw(screen *sdl.Surface) {
	screen.FillRect(
		&sdl.Rect{
			X: w.X,
			Y: w.Y,
			W: w.W,
			H: w.H,
		},
		ui.ColorToInt(w.BorderColor),
	)
	screen.FillRect(
		&sdl.Rect{
			X: w.X + int16(w.BorderWidth),
			Y: w.Y + int16(w.BorderWidth),
			W: w.W - w.BorderWidth*2,
			H: w.H - w.BorderWidth*2,
		},
		ui.ColorToInt(w.BGColor),
	)

	for _, child := range w.children {
		child.Draw(screen)
	}
}

func (w *Window) HandleEvent(event interface{}) bool {
	for _, child := range w.children {
		if child.HandleEvent(event) {
			return true
		}
	}

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

func (w *Window) SetPos(x, y int16) {
	dx, dy := x-w.X, y-w.Y
	w.X, w.Y = x, y
	for _, child := range w.children {
		x, y := child.GetPos()
		x, y = x+dx, y+dy
		child.SetPos(x, y)
	}
}

func (w *Window) GetPos(x, y int16) {
	x, y = w.X, w.Y
	return
}
