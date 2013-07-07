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

// Package textbox defines a text display area UI element.
package textbox

import (
	"github.com/bieber/drones/fonts"
	"github.com/neagix/Go-SDL/sdl"
	"strings"
)

// Font alignments
type Alignment int

const (
	Left Alignment = iota
	Center
	Right
)

// Defaults
const fontSize = 12

// TextBox renders text in the area it occupies.
type TextBox struct {
	lines []*sdl.Surface

	//TextBox position and size.
	X int16
	Y int16
	W uint16
	H uint16

	// Text alignment.
	FontAlignment Alignment
}

func New(
	x, y int16,
	w, h uint16,
	text, font string,
	fontSize int,
	fontColor sdl.Color,
	fontAlignment Alignment,
) (t *TextBox) {
	t = &TextBox{
		X:             x,
		Y:             y,
		W:             w,
		H:             h,
		FontAlignment: fontAlignment,
	}
	t.buildText(text, font, fontSize, fontColor)
	return
}

func (t *TextBox) Draw(screen *sdl.Surface) {
	y := t.Y
	lineSpacing := 0
	for _, line := range t.lines {
		if line == nil {
			y += int16(lineSpacing)
			continue
		} else {
			lineSpacing = int(line.H) / 2
		}

		if uint16(y)+uint16(line.H)-uint16(t.Y) > t.H {
			continue
		}

		x := int16(t.X)
		switch t.FontAlignment {
		case Left:
			x += 0
		case Center:
			x += (int16(t.W) - int16(line.W)) / 2
		case Right:
			x += int16(t.W) - int16(line.W)
		}

		screen.Blit(
			&sdl.Rect{
				X: x,
				Y: y,
				W: uint16(line.W),
				H: uint16(line.H),
			},
			line,
			&sdl.Rect{W: uint16(line.W), H: uint16(line.H)},
		)
		y += int16(line.H)
	}
}

func (t *TextBox) HandleEvent(event interface{}) bool {
	return false
}

func (t *TextBox) SetPos(x, y int16) {
	t.X, t.Y = x, y
}

func (t *TextBox) GetPos() (x, y int16) {
	x, y = t.X, t.Y
	return
}

func (t *TextBox) buildText(
	text, font string,
	fontSize int,
	fontColor sdl.Color,
) {
	for _, line := range t.lines {
		line.Free()
	}
	t.lines = make([]*sdl.Surface, 0, 10)
	lines := strings.Split(strings.Trim(text, "\n"), "\n")
	for _, line := range lines {
		if line == "" {
			// SDL chokes on empty lines, so replace them with nil straight out
			t.lines = append(t.lines, nil)
			continue
		}
		words := strings.Split(line, " ")
		for current, rest := words[:], []string(nil); len(current) != 0; {
			w, _, _ := fonts.Size(strings.Join(current, " "), font, fontSize)
			if w > int(t.W) {
				rest = append(current[len(current)-1:], rest...)
				current = current[:len(current)-1]
			} else {
				t.lines = append(
					t.lines,
					fonts.BlendedText(
						strings.Join(current, " "),
						font,
						fontSize,
						fontColor,
					),
				)
				current = rest
				rest = nil
			}
		}
	}
}
