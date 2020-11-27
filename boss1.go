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
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	boss1Points = 15000
)

func makeBoss1(x, y float64) boss {
	hitb1 := bossHitBox{
		x:     x,
		xrel:  175,
		y:     y,
		yrel:  -100,
		xSize: 300,
		ySize: 175,
		hullShape: []point{
			point{x: -150, y: 10},
			point{x: -75, y: -90},
			point{x: 150, y: -15},
			point{x: 150, y: 85},
			point{x: -70, y: 85},
		},
	}
	hitb1.updateBox()
	hitb2 := bossHitBox{
		x:     x,
		xrel:  175,
		y:     y,
		yrel:  70,
		xSize: 300,
		ySize: 175,
		hullShape: []point{
			point{x: -150, y: 0},
			point{x: -75, y: -83},
			point{x: 150, y: -83},
			point{x: 150, y: 15},
			point{x: -75, y: 85},
		},
	}
	hitb2.updateBox()
	hurtb1 := bossHitBox{
		x:     x,
		xrel:  0,
		y:     y,
		yrel:  -300,
		xSize: 1224,
		ySize: 500,
		hullShape: []point{
			point{x: -600, y: 140},
			point{x: -450, y: 250},
			point{x: 300, y: 140},
			point{x: 0, y: -200},
		},
	}
	hurtb1.updateBox()
	hurtb2 := bossHitBox{
		x:     x,
		xrel:  -10,
		y:     y,
		yrel:  300,
		xSize: 1000,
		ySize: 450,
		hullShape: []point{
			point{x: -490, y: -30},
			point{x: -120, y: 190},
			point{x: 300, y: -200},
			point{x: -30, y: -190},
		},
	}
	hurtb2.updateBox()
	return boss{
		x:         x,
		xSize:     1224,
		y:         y,
		ySize:     1000,
		pv:        350,
		bossType:  boss1,
		points:    boss1Points,
		hitBoxes:  []bossHitBox{hitb1, hitb2},
		hurtBoxes: []bossHitBox{hurtb1, hurtb2},
	}
}

func (b *boss) boss1Update(bs *bulletSet) {
	switch b.phase {
	case 0:
		if b.x > 8*float64(screenWidth)/9 {
			b.x -= 5
		} else {
			b.phase = 1
			for pos := 0; pos < len(b.hitBoxes); pos++ {
				b.hitBoxes[pos].hitable = true
			}
		}
	}
}

func (b *boss) boss1Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.x-b.xSize/2, b.y-b.ySize/2)
	screen.DrawImage(
		boss1Image,
		op,
	)
}
