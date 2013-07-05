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

// Package mainmenu defines the main menu layer.
package mainmenu

import (
	"github.com/bieber/drones/drones/about"
	"github.com/bieber/drones/fonts"
	"github.com/bieber/drones/ui"
	"github.com/neagix/Go-SDL/sdl"
	"time"
)

var textColor sdl.Color = sdl.Color{255, 255, 255, 0}
var unselectedColor sdl.Color = sdl.Color{180, 180, 180, 0}
var selectedColor sdl.Color = sdl.Color{255, 255, 255, 0}

const menuItemXOffset = 20
const menuItemYOffset = 100
const titleFontSize = 40
const menuItemFontSize = 25

// MainMenu is the bottom-most layer of drones, and renders/manages
// the main menu.
type MainMenu struct {
	titleText        *sdl.Surface
	itemLabels       []string
	menuItemSelected string
	menuItems        map[string]*sdl.Surface
	hoverItems       map[string]*sdl.Surface
	mouseX           uint16
	mouseY           uint16
	newLayers        chan ui.Layer
}

func New(newLayers chan ui.Layer) (m *MainMenu) {
	m = &MainMenu{
		titleText: fonts.ShadedText(
			"Drones",
			"sans_bold",
			titleFontSize,
			textColor,
			sdl.Color{0, 0, 0, 0},
		),
		menuItemSelected: "",
		mouseX:           0,
		mouseY:           0,
		newLayers:        newLayers,
	}

	// List menu items from bottom to top
	m.itemLabels = []string{"About", "Settings", "Play"}
	m.menuItems = make(map[string]*sdl.Surface, len(m.itemLabels))
	m.hoverItems = make(map[string]*sdl.Surface, len(m.itemLabels))
	for _, label := range m.itemLabels {
		m.menuItems[label] = fonts.ShadedText(
			label,
			"sans_bold",
			menuItemFontSize,
			unselectedColor,
			sdl.Color{0, 0, 0, 0},
		)
		m.hoverItems[label] = fonts.ShadedText(
			label,
			"sans_bold",
			menuItemFontSize,
			selectedColor,
			sdl.Color{0, 0, 0, 0},
		)
	}
	return
}

func (m *MainMenu) Draw(screen *sdl.Surface, top bool) {
	w := uint16(m.titleText.W)
	h := uint16(m.titleText.H)
	screen.Blit(
		&sdl.Rect{
			X: int16(screen.W - m.titleText.W - 50),
			Y: 15,
			W: w,
			H: h,
		},
		m.titleText,
		&sdl.Rect{X: 0, Y: 0, W: w, H: h},
	)

	m.menuItemSelected = ""
	y := screen.H - menuItemYOffset
	for _, label := range m.itemLabels {
		item := m.menuItems[label]
		w = uint16(item.W)
		h = uint16(item.H)
		y -= item.H
		if ui.Overlaps(m.mouseX, m.mouseY, menuItemXOffset, uint16(y), w, h) {
			if top {
				item = m.hoverItems[label]
			}
			m.menuItemSelected = label
		}

		screen.Blit(
			&sdl.Rect{
				X: menuItemXOffset,
				Y: int16(y),
				W: w,
				H: h,
			},
			item,
			&sdl.Rect{X: 0, Y: 0, W: w, H: h},
		)
	}
}

func (m *MainMenu) Tick(elapsed time.Duration) bool {
	return true
}

func (m *MainMenu) HandleEvent(event interface{}) bool {
	switch event.(type) {
	case sdl.MouseMotionEvent:
		e := event.(sdl.MouseMotionEvent)
		m.mouseX = e.X
		m.mouseY = e.Y
		return true

	case sdl.MouseButtonEvent:
		e := event.(sdl.MouseButtonEvent)
		if e.Type == sdl.MOUSEBUTTONDOWN && e.Button == 1 {
			m.click()
		}
		return true
	}
	return false
}

func (m *MainMenu) click() {
	switch m.menuItemSelected {
	case "About":
		m.newLayers <- about.New()
	}
}
