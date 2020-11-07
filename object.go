/* A game for Game Off 2020
// Copyright (C) 2020 Loïg Jezequel
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>
*/

package main

type collidableObject interface {
	xmin() float64
	xmax() float64
	ymin() float64
	ymax() float64
	hasCollided()
}

func collide(o collidableObject, oo collidableObject) bool {
	collision := !(o.xmin() > oo.xmax() ||
		o.xmax() < oo.xmin() ||
		o.ymin() > oo.ymax() ||
		o.ymax() < oo.ymin())
	if collision {
		o.hasCollided()
		oo.hasCollided()
	}
	return collision
}
