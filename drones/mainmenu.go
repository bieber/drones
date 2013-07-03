/*
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

package main

import (
	"github.com/bieber/drones/fonts"
	"github.com/neagix/Go-SDL/sdl"
	"time"
)

type MainMenu struct {
	cursor sdl.Rect
}

func (m *MainMenu) Draw(screen *sdl.Surface, top bool) {
	screen.FillRect(&m.cursor, 0x0000ff00)
	width, height, _ := fonts.Size("Drones", "sans_bold", 25)
	text := fonts.BlendedText(
		"Drones",
		"sans_bold",
		25,
		sdl.Color{255, 255, 255, 0},
	)
	screen.Blit(
		&sdl.Rect{
			X: int16(int(screen.W) - width - 10),
			Y: 0,
			W: uint16(width),
			H: uint16(height),
		},
		text,
		&sdl.Rect{X: 0, Y: 0, W: uint16(width), H: uint16(height)},
	)
}

func (m *MainMenu) Tick(elapsed time.Duration) bool {
	return true
}

func (m *MainMenu) HandleEvent(event interface{}) bool {
	switch event.(type) {
	case sdl.MouseMotionEvent:
		motion := event.(sdl.MouseMotionEvent)
		m.cursor.X = int16(motion.X) - 5
		m.cursor.Y = int16(motion.Y) - 5
		m.cursor.W = 10
		m.cursor.H = 10
	default:
		return false
	}
	return false
}
