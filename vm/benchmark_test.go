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

package vm

import (
	ck "launchpad.net/gocheck"
)

func (s *suite) BenchmarkVM(c *ck.C) {
	s.vm.LoadMem(
		[]uint16{
			9, 12,
			13, 0,
			9, 15,
			13, 0,
			17, 16,
			14, 0,
			14, 0,
			2, 0,
			8, 0,
			9, 1,
			7, 0,
			12, 0,
			5, 0,
			9, 2,
			7, 0,
			12, 0,
			21, 0,
			18, 0,
		},
	)
	s.vm.ClockN(c.N)
}
