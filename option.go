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

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	oSizex = 30
	oSizey = 30
)

type option struct {
	x float64
	y float64
}

func (o option) xmin() float64 {
	return o.x - oSizex/2
}

func (o option) xmax() float64 {
	return o.x + oSizex/2
}

func (o option) ymin() float64 {
	return o.y - oSizey/2
}

func (o option) ymax() float64 {
	return o.y + oSizey/2
}

func (o option) convexHull() []point {
	return []point{
		point{o.xmin(), o.ymin()},
		point{o.xmax(), o.ymin()},
		point{o.xmax(), o.ymax()},
		point{o.xmin(), o.ymax()},
	}
}

func (o option) hasCollided() {}

func (o option) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(o.xmin(), o.ymin())
	screen.DrawImage(
		optionImage,
		op,
	)
	if isDebug() {
		cHull := o.convexHull()
		for i := 0; i < len(cHull); i++ {
			ii := (i + 1) % len(cHull)
			ebitenutil.DrawLine(screen, cHull[i].x, cHull[i].y, cHull[ii].x, cHull[ii].y, color.RGBA{0, 255, 0, 255})
		}
	}
}
