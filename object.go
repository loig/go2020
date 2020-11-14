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
	convexHull() []point
	hasCollided()
}

type point struct {
	x float64
	y float64
}

func collideNoHarm(o collidableObject, oo collidableObject) bool {
	collision := !(o.xmin() > oo.xmax() ||
		o.xmax() < oo.xmin() ||
		o.ymin() > oo.ymax() ||
		o.ymax() < oo.ymin())
	if !collision {
		return false
	}
	collision = intersectHulls(o, oo) && intersectHulls(oo, o)
	if collision {
		oo.hasCollided()
	}
	return collision
}

func collide(o collidableObject, oo collidableObject) bool {
	collision := collideNoHarm(o, oo)
	if collision {
		o.hasCollided()
	}
	return collision
}

func intersectHulls(o collidableObject, oo collidableObject) bool {
	// will not work if o contains only two points
	ocHull := o.convexHull()
	oocHull := oo.convexHull()
	for i := 0; i < len(ocHull); i++ {
		a := ocHull[i]
		b := ocHull[(i+1)%len(ocHull)]
		c := ocHull[(i+2)%len(ocHull)]
		// check on which side of (a,b) is c
		side := (c.x-a.x)*(b.y-a.y) - (c.y-a.y)*(b.x-a.x)
		left := side < 0
		// for each point on oocHull, check if it is on the other side of (a,b)
		allSameSide := true
		for _, p := range oocHull {
			side = (p.x-a.x)*(b.y-a.y) - (p.y-a.y)*(b.x-a.x)
			if left {
				if side < 0 {
					allSameSide = false
					break
				}
			} else {
				if side > 0 {
					allSameSide = false
					break
				}
			}
		}
		if allSameSide {
			return false
		}
	}
	return true
}
