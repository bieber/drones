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

// Package button defines a push-button UI element.
package button

import (
	"github.com/bieber/drones/fonts"
	"github.com/bieber/drones/ui"
	"github.com/neagix/Go-SDL/sdl"
)

// Defaults
var bgColor sdl.Color = sdl.Color{200, 200, 200, 0}
var hoverColor sdl.Color = sdl.Color{255, 255, 255, 0}
var borderColor sdl.Color = sdl.Color{0, 0, 0, 0}
var textColor sdl.Color = sdl.Color{0, 0, 0, 0}

const borderWidth = 2
const fontSize = 15

// Button draws a rectangular area on the screen with a label in the
// center, and responds to mouse clicks.
type Button struct {
	hovered   bool
	labelText *sdl.Surface

	// Button position and size.
	X int16
	Y int16
	W uint16
	H uint16

	// Regular and hover background colors.
	BGColor    sdl.Color
	HoverColor sdl.Color

	// Border color and width.
	BorderColor sdl.Color
	BorderWidth uint

	// Callback to trigger on button press.
	OnClick func()
}

func New(x, y int16, w, h uint16, text string) (b *Button) {
	b = &Button{
		X:           x,
		Y:           y,
		W:           w,
		H:           h,
		BGColor:     bgColor,
		HoverColor:  hoverColor,
		BorderColor: borderColor,
		BorderWidth: borderWidth,
	}
	b.refreshText(text, fontSize, textColor)
	return
}

func (b *Button) SetLabel(text string) {
	b.refreshText(text, fontSize, textColor)
}

func (b *Button) Draw(screen *sdl.Surface) {
	screen.FillRect(
		&sdl.Rect{
			X: int16(b.X),
			Y: int16(b.Y),
			W: uint16(b.W),
			H: uint16(b.H),
		},
		ui.ColorToInt(b.BorderColor),
	)
	var fillColor sdl.Color
	if b.hovered {
		fillColor = b.HoverColor
	} else {
		fillColor = b.BGColor
	}
	screen.FillRect(
		&sdl.Rect{
			X: int16(b.X) + int16(b.BorderWidth),
			Y: int16(b.Y) + int16(b.BorderWidth),
			W: uint16(b.W) - 2*uint16(b.BorderWidth),
			H: uint16(b.H) - 2*uint16(b.BorderWidth),
		},
		ui.ColorToInt(fillColor),
	)
	screen.Blit(
		&sdl.Rect{
			X: int16(b.X) + int16((int32(b.W)-b.labelText.W)/2),
			Y: int16(b.Y) + int16((int32(b.H)-b.labelText.H)/2),
			W: uint16(b.labelText.W),
			H: uint16(b.labelText.H),
		},
		b.labelText,
		&sdl.Rect{W: uint16(b.labelText.W), H: uint16(b.labelText.H)},
	)
}

func (b *Button) HandleEvent(event interface{}) bool {
	switch event.(type) {
	case sdl.MouseMotionEvent:
		m := event.(sdl.MouseMotionEvent)
		if ui.Overlaps(int16(m.X), int16(m.Y), b.X, b.Y, b.W, b.H) {
			b.hovered = true
			return true
		} else {
			b.hovered = false
		}
	case sdl.MouseButtonEvent:
		m := event.(sdl.MouseButtonEvent)
		overlap := ui.Overlaps(int16(m.X), int16(m.Y), b.X, b.Y, b.W, b.H)
		if ui.IsMouseDown(m, 1) && overlap && b.OnClick != nil {
			b.OnClick()
			return true
		}
	}
	return false
}

func (b *Button) SetPos(x, y int16) {
	b.X, b.Y = x, y
}

func (b *Button) GetPos() (x, y int16) {
	x, y = b.X, b.Y
	return
}

func (b *Button) refreshText(text string, size int, color sdl.Color) {
	b.labelText = fonts.BlendedText(text, "sans_bold", size, color)
}
