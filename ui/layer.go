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

// Package ui defines some primitive UI elements.
package ui

import (
	"github.com/neagix/Go-SDL/sdl"
	"sort"
	"time"
)

// Layer represents a type that can present a layer on the screen.
type Layer interface {
	// Draw prompts the layer to draw to the screen. The top parameter
	// specifies whether the layer is at the top of the stack.
	Draw(screen *sdl.Surface, top bool)
	// Tick prompts the layer to advance its state by the specified
	// duration.  Returning false will kill the layer.
	Tick(elapsed time.Duration) bool
	// HandleEvent notifies the layer of an event.  Returning false
	// will allow the event to pass further down on the layer stack.
	HandleEvent(event interface{}) bool
}

type LayerStack []Layer

// HandleEvent passes an event through the layer stack, and returns
// true if any of the layers handled it successfully.
func (stack LayerStack) HandleEvent(event interface{}) bool {
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i].HandleEvent(event) {
			return true
		}
	}
	return false
}

// Draw directs each layer, starting from the bottom, to draw to the
// screen, and then updates the necessary parts of the screen.
func (stack LayerStack) Draw(screen *sdl.Surface) {
	screen.FillRect(
		&sdl.Rect{X: 0, Y: 0, W: uint16(screen.W), H: uint16(screen.H)},
		0,
	)
	for i, layer := range stack {
		top := i == len(stack)-1
		layer.Draw(screen, top)
	}
	screen.Flip()
}

// Tick executes a tick on each layer, and returns a list of layer
// indices that have exited.
func (stack LayerStack) Tick(elapsed time.Duration) []int {
	killed := make([]int, 0, 5)
	for i, layer := range stack {
		if !layer.Tick(elapsed) {
			killed = append(killed, i)
		}
	}
	return killed
}

// RemoveLayers removes layers at the specified indices from the
// stack.
func (stack *LayerStack) RemoveLayers(indices []int) {
	sort.Ints(indices)
	newStack := make([]Layer, 0, len(*stack)-len(indices))
	dI := 0
	for k, v := range *stack {
		if k == indices[dI] {
			if dI < len(indices)-1 {
				dI++
			}
			continue
		}
		newStack = append(newStack, v)
	}
	*stack = newStack
}
