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
	"github.com/bieber/drones/ui/button"
	"github.com/bieber/drones/ui/textbox"
	"github.com/bieber/drones/ui/window"
	"github.com/neagix/Go-SDL/sdl"
	"time"
)

const (
	windowWidth  = 500
	windowHeight = 300
)

const (
	buttonWidth   = 100
	buttonHeight  = 40
	buttonPadding = 10
)

const (
	aboutPadding = 5
	fontSize     = 12
)

var textColor sdl.Color = sdl.Color{255, 255, 255, 0}

// About is a layer that displays an About dialog box.
type About struct {
	running bool
	window  *window.Window
}

func New() (a *About) {
	a = &About{
		running: true,
		window:  window.New(0, 0, windowWidth, windowHeight),
	}
	a.window.OnEscape = func() { a.running = false }

	b := button.New(0, 0, buttonWidth, buttonHeight, "Ok")
	b.OnClick = func() { a.running = false }
	a.window.AddChild(
		b,
		int16(windowWidth-buttonWidth-a.window.BorderWidth-buttonPadding),
		int16(windowHeight-buttonHeight-a.window.BorderWidth-buttonPadding),
	)
	a.window.AddChild(
		textbox.New(
			0,
			0,
			windowWidth-2*a.window.BorderWidth-2*aboutPadding,
			windowHeight-
				2*a.window.BorderWidth-
				2*buttonPadding-buttonHeight-
				2*aboutPadding,
			aboutMessage,
			"serif_bold",
			fontSize,
			textColor,
			textbox.Center,
		),
		int16(a.window.BorderWidth+aboutPadding),
		int16(a.window.BorderWidth+aboutPadding),
	)
	return
}

func (a *About) Draw(screen *sdl.Surface, top bool) {
	a.window.SetPos(
		int16((screen.W-int32(a.window.W))/2),
		int16((screen.H-int32(a.window.H))/2),
	)
	a.window.Draw(screen)
}

func (a *About) Tick(elapsed time.Duration) bool {
	return a.running
}

func (a *About) HandleEvent(event interface{}) bool {
	return a.window.HandleEvent(event)
}
