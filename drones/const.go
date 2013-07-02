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

// ScreenWidth and ScreenHeight determine the screen dimensions.
const (
	ScreenWidth  int = 640
	ScreenHeight int = 480
)

// FPS is the desired framerate.
const FPS int = 30

// ClockRate is the desired CPU cycles/second for VMs.
const ClockRate int = 1000

// ResPath is the directory to install resources to on startup.
const ResPath string = ".resources"
