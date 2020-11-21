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
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func makeMidBoss1(x, y float64) boss {
	return boss{
		x:        x,
		xSize:    50,
		y:        y,
		ySize:    300,
		pv:       150,
		bossType: midBoss1,
	}
}

const (
	midBoss1NumBulletPhase1      = 7
	midBoss1FramePerBulletPhase1 = 4
	midBoss1LengthPhase1         = midBoss1NumBulletPhase1
	midBoss1NumLoopPhase1        = 13
	midBoss1BulletSpeedPhase1    = 5
	midBoss1NumBulletPhase2      = 7
	midBoss1FramePerBulletPhase2 = 6
	midBoss1LengthPhase2         = midBoss1NumBulletPhase2
	midBoss1BulletSpeedPhase2    = 4
	midBoss1NumLoopPhase2        = 7
)

func (b *boss) midBoss1Update(bs *bulletSet) {
	switch b.phase {
	case 0:
		if b.x > 3*float64(screenWidth)/4 {
			b.x -= 5
		} else {
			b.phase = 1
		}
	case 1:
		b.frame++
		if b.pv <= 75 {
			b.phase = 2
		}
		numBullet := b.frame/midBoss1FramePerBulletPhase1 - 1
		if numBullet < midBoss1NumBulletPhase1 {
			if b.frame%midBoss1FramePerBulletPhase1 == 0 {
				angleShift := (float64(b.phaseLoop) - float64(midBoss1NumLoopPhase1)/2) / float64(2*midBoss1NumLoopPhase1)
				angle := (math.Pi/2)*float64(numBullet)/float64(midBoss1NumBulletPhase1-1) + 3*math.Pi/4 + angleShift*math.Pi
				vx := midBoss1BulletSpeedPhase1 * math.Cos(angle)
				vy := midBoss1BulletSpeedPhase1 * -math.Sin(angle)
				bs.addBullet(bullet{
					x:     b.x,
					y:     b.ymin() + (float64(numBullet)+0.5)*b.ySize/float64(midBoss1NumBulletPhase1),
					vx:    vx,
					vy:    vy,
					image: enemyBasicBullet,
				})
			}
		}
		if b.frame/midBoss1FramePerBulletPhase1 >= midBoss1LengthPhase1 {
			b.frame = 0
			b.phaseLoop++
			if b.phaseLoop >= midBoss1NumLoopPhase1 {
				b.phaseLoop = 0
			}
		}
	case 2:
		b.frame++
		numBullet := b.frame/midBoss1FramePerBulletPhase2 - 1
		if numBullet < midBoss1NumBulletPhase2 {
			if b.frame%midBoss1FramePerBulletPhase2 == 0 {
				angleShift := (float64(b.phaseLoop) - float64(midBoss1NumLoopPhase2)/2) / float64(2*midBoss1NumLoopPhase2)
				angle := (math.Pi/2)*float64(numBullet)/float64(midBoss1NumBulletPhase2-1) + 3*math.Pi/4 + angleShift*math.Pi
				vx := midBoss1BulletSpeedPhase2 * math.Cos(angle)
				vy := midBoss1BulletSpeedPhase2 * -math.Sin(angle)
				bs.addBullet(bullet{
					x:     b.x,
					y:     b.ymin() + (float64(numBullet)+0.5)*b.ySize/float64(midBoss1NumBulletPhase2),
					vx:    vx,
					vy:    vy,
					image: enemyBasicBullet,
				})
				bs.addBullet(bullet{
					x:     b.x,
					y:     b.ymax() - (float64(numBullet)+0.5)*b.ySize/float64(midBoss1NumBulletPhase2),
					vx:    vx,
					vy:    -vy,
					image: enemyBasicBullet,
				})
			}
		}
		if b.frame/midBoss1FramePerBulletPhase2 >= midBoss1LengthPhase2 {
			b.frame = 0
			b.phaseLoop++
			if b.phaseLoop >= midBoss1NumLoopPhase2 {
				b.phaseLoop = 0
			}
		}
	}
}

func (b *boss) midBoss1Draw(screen *ebiten.Image) {
	cHull := b.convexHull()
	hullColor := color.RGBA{0, 255, 0, 255}
	for i := 0; i < len(cHull); i++ {
		ii := (i + 1) % len(cHull)
		ebitenutil.DrawLine(screen, cHull[i].x, cHull[i].y, cHull[ii].x, cHull[ii].y, hullColor)
	}
}
