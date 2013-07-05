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

package ui

import (
	"github.com/neagix/Go-SDL/sdl"
)

// Widget defines an interface for UI elements that can be rendered
// on-screen and respond to events.
type Widget interface {
	// Draw should render the widget to the screen.
	Draw(screen *sdl.Surface)

	// HandleEvent should process an event from SDL.  Return true to
	// signal that the event was successfully handled, or false to
	// signal that it should be handled by another component.
	HandleEvent(event interface{}) bool

	// SetPos positions the widget in absolute coordinates.
	SetPos(x, y int16)
	// GetPos returns the position of the widget in absolute coordinates.
	GetPos() (x, y int16)
}
