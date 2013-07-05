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

// Package about defines an about dialog layer.
package about

import (
	"github.com/bieber/drones/ui/window"
	"github.com/neagix/Go-SDL/sdl"
	"time"
)

const windowWidth = 500
const windowHeight = 300
const borderWidth = 5

var bgColor sdl.Color = sdl.Color{150, 150, 150, 0}
var borderColor sdl.Color = sdl.Color{100, 100, 100, 0}

// About is a layer that displays an About dialog box.
type About struct {
	running bool
	window  *window.Window
}

func New() (a *About) {
	a = &About{
		running: true,
		window: window.New(
			0,
			0,
			windowWidth,
			windowHeight,
			borderWidth,
			bgColor,
			borderColor,
		),
	}
	a.window.OnEscape = func() { a.running = false }
	return
}

func (a *About) Draw(screen *sdl.Surface, top bool) {
	a.window.SetPos(
		(uint(screen.W)-uint(a.window.W))/2,
		(uint(screen.H)-uint(a.window.H))/2,
	)
	a.window.Draw(screen)
}

func (a *About) Tick(elapsed time.Duration) bool {
	return a.running
}

func (a *About) HandleEvent(event interface{}) bool {
	return a.window.HandleEvent(event)
}
