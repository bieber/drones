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

// ColorToInt translates an SDL color struct into a 32-bit integer.
func ColorToInt(c sdl.Color) (color uint32) {
	color = uint32(c.R)
	color <<= 8
	color |= uint32(c.G)
	color <<= 8
	color |= uint32(c.B)
	return
}
