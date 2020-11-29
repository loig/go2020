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
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	midBoss1Points = 5000
	midBoss1PV     = 200
)

func makeMidBoss1(y float64) boss {
	x := float64(screenWidth - 1)
	hb := bossHitBox{
		x:     x,
		y:     y,
		xSize: 500,
		ySize: 300,
		hullShape: []point{
			point{x: -5, y: -145},
			point{x: -60, y: -72},
			point{x: -60, y: 72},
			point{x: -5, y: 145},
			point{x: 145, y: 145},
			point{x: 200, y: 72},
			point{x: 200, y: -72},
			point{x: 145, y: -145},
		},
	}
	hb.updateBox()
	return boss{
		x:        x,
		xSize:    300,
		y:        y,
		ySize:    300,
		pv:       midBoss1PV,
		bossType: midBoss1,
		points:   midBoss1Points,
		hitBoxes: []bossHitBox{hb},
	}
}

const (
	midBoss1NumBulletPhase1      = 7
	midBoss1FramePerBulletPhase1 = 6 // 4
	midBoss1LengthPhase1         = midBoss1NumBulletPhase1
	midBoss1NumLoopPhase1        = 13
	midBoss1BulletSpeedPhase1    = 3 // 5
)

func (b *boss) midBoss1Update(bs *bulletSet) bool {
	switch b.phase {
	case 0:
		if b.x > 6*float64(screenWidth)/7 {
			b.x -= 5
		} else {
			b.phase = 1
			b.hitBoxes[0].hitable = true
		}
	case 1:
		b.frame++
		var hasFired bool
		numBullet := b.frame/midBoss1FramePerBulletPhase1 - 1
		if numBullet < midBoss1NumBulletPhase1 {
			if b.frame%midBoss1FramePerBulletPhase1 == 0 {
				angleShift := (float64(b.phaseLoop) - float64(midBoss1NumLoopPhase1)/2) / float64(2*midBoss1NumLoopPhase1)
				angle := (math.Pi/2)*float64(numBullet)/float64(midBoss1NumBulletPhase1-1) + 3*math.Pi/4 + angleShift*math.Pi
				vx := midBoss1BulletSpeedPhase1 * math.Cos(angle)
				vy := midBoss1BulletSpeedPhase1 * -math.Sin(angle)
				bs.addBullet(bullet{
					x:     b.x,
					y:     b.hitBoxes[0].ymin() + (float64(numBullet)+0.5)*b.ySize/float64(midBoss1NumBulletPhase1),
					vx:    vx,
					vy:    vy,
					image: enemyBasicBullet,
				})
				hasFired = true
			}
		}
		if b.frame/midBoss1FramePerBulletPhase1 >= midBoss1LengthPhase1 {
			b.frame = 0
			b.phaseLoop++
			if b.phaseLoop >= midBoss1NumLoopPhase1 {
				b.phaseLoop = 0
			}
		}
		return hasFired
	}
	return false
}

func (b *boss) midBoss1Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.x-b.xSize/4, b.y-b.ySize/2)
	screen.DrawImage(
		midBoss1Image,
		op,
	)
}
