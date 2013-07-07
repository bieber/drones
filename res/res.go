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

// Package res contains static resources compiled into Go functions
// with go-bindata.
package res

import (
	"os"
	"path/filepath"
)

var files map[string]func() []byte = map[string]func() []byte{
	"FreeSansBold.ttf":  FreeSansBold,
	"FreeSerifBold.ttf": FreeSerifBold,
}

func WriteResources(dir string) {
	err := os.Mkdir(dir, 0744)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	for path, f := range files {
		file, err := os.Create(filepath.Join(dir, path))
		if err != nil {
			panic(err)
		}
		_, err = file.Write(f())
		if err != nil {
			panic(err)
		}
		file.Close()
	}
}
