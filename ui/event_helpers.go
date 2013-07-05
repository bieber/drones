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

func IsKeyDown(event interface{}, sym uint32) bool {
	k, ok := event.(sdl.KeyboardEvent)
	if !ok {
		return false
	}
	return k.Type == sdl.KEYDOWN && k.Keysym.Sym == sym
}

func IsKeyUp(event interface{}, sym uint32) bool {
	k, ok := event.(sdl.KeyboardEvent)
	if !ok {
		return false
	}
	return k.Type == sdl.KEYUP && k.Keysym.Sym == sym
}

func IsMouseDown(event interface{}, button uint8) bool {
	m, ok := event.(sdl.MouseButtonEvent)
	if !ok {
		return false
	}
	return m.Type == sdl.MOUSEBUTTONDOWN && m.Button == button
}

func IsMouseUp(event interface{}, button uint8) bool {
	m, ok := event.(sdl.MouseButtonEvent)
	if !ok {
		return false
	}
	return m.Type == sdl.MOUSEBUTTONUP && m.Button == button
}
