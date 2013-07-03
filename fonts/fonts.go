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

// Package fonts provides font loading and caching.
package fonts

import (
	"fmt"
	"github.com/neagix/Go-SDL/sdl"
	"github.com/neagix/Go-SDL/ttf"
	"path/filepath"
)

var fontsDir string

var fontFiles map[string]string = map[string]string{
	"sans_bold": "FreeSansBold.ttf",
}

var fonts map[string]*ttf.Font = make(map[string]*ttf.Font)

// SetFontsPath sets the resource directory to search for font files.
func SetFontsPath(path string) {
	fontsDir = path
}

// Size calculates the size of the given text in the given font with a
// possible error code as the third return value.
func Size(text string, font string, size int) (int, int, int) {
	return GetFont(font, size).SizeUTF8(text)
}

// BlendedText returns blended text of the desired font and size
// rendered to an SDL surface.
func BlendedText(
	text string,
	font string,
	size int,
	color sdl.Color,
) *sdl.Surface {
	return ttf.RenderUTF8_Blended(GetFont(font, size), text, color)
}

// ShadedText returns shaded text of the desired font and size
// rendered to an SDL surface.
func ShadedText(
	text string,
	font string,
	size int,
	color sdl.Color,
	bgColor sdl.Color,
) *sdl.Surface {
	return ttf.RenderUTF8_Shaded(GetFont(font, size), text, color, bgColor)
}

// SolidText returns solid text of the desired font and size rendered
// to an SDL surface.
func SolidText(
	text string,
	font string,
	size int,
	color sdl.Color,
) *sdl.Surface {
	return ttf.RenderUTF8_Solid(GetFont(font, size), text, color)
}

// GetFont fetches a font from a file or retrieves it from cache if
// it's already been requested.
func GetFont(name string, size int) *ttf.Font {
	font, ok := fonts[fmt.Sprintf("%s:%d", name, size)]
	if ok {
		return font
	} else {
		fontFile := fontFiles[name]
		font := ttf.OpenFont(filepath.Join(fontsDir, fontFile), size)
		fonts[fmt.Sprintf("%s:%d", name, size)] = font
		return font
	}
}
