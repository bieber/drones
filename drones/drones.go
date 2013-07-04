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

// This is the game binary for drones.
package main

import (
	"github.com/bieber/drones/ui"
	"github.com/bieber/drones/fonts"
	"github.com/bieber/drones/res"
	"github.com/neagix/Go-SDL/sdl"
	"github.com/neagix/Go-SDL/ttf"
	"time"
)

func main() {
	res.WriteResources(ResPath)
	fonts.SetFontsPath(ResPath)

	if sdl.Init(sdl.INIT_VIDEO) != 0 {
		panic(sdl.GetError())
	}
	defer sdl.Quit()
	if ttf.Init() != 0 {
		panic("Error initializing SDL TTF")
	}
	defer ttf.Quit()

	screen := sdl.SetVideoMode(ScreenWidth, ScreenHeight, 32, 0)
	if screen == nil {
		panic(sdl.GetError())
	}
	sdl.WM_SetCaption("Drones", "")

	layerStack := ui.LayerStack{
		&MainMenu{cursor: sdl.Rect{X: 0, Y: 0, W: 10, H: 10}},
	}
	frameTime := time.Second / time.Duration(FPS)
	tickTimer := time.NewTicker(frameTime)

	for run := true; run; {
		select {
		case <-tickTimer.C:
			if toRemove := layerStack.Tick(frameTime); len(toRemove) != 0 {
				layerStack.RemoveLayers(toRemove)
			}
			layerStack.Draw(screen)
		case event := <-sdl.Events:
			layerStack.HandleEvent(event)
			switch event.(type) {
			case sdl.QuitEvent:
				run = false
			}
		}
	}
}
