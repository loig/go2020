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

type player struct {
	x     float64
	y     float64
	xSize float64
	ySize float64
}

func initPlayer(x, y float64) player {
	return player{
		x: x, y: y,
		xSize: 45, ySize: 15,
	}
}

func (p player) draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.xmin(), p.ymin(), p.xSize, p.ySize, color.RGBA{255, 0, 0, 255})
}

func (p player) xmin() float64 {
	return p.x - p.xSize/2
}

func (p player) xmax() float64 {
	return p.x + p.xSize/2
}

func (p player) ymin() float64 {
	return p.y - p.ySize/2
}

func (p player) ymax() float64 {
	return p.y + p.ySize/2
}

func (p player) hasCollided() {

}

func (p player) checkCollisions(cos []*bullet) {
	for _, o := range cos {
		collide(p, o)
	}
}
