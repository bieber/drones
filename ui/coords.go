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

func Overlaps(x, y, left, top int16, width, height uint16) bool {
	if x < left || x >= left+int16(width) {
		return false
	}
	if y < top || y >= top+int16(height) {
		return false
	}
	return true
}
