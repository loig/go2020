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
	boss1Points                = 15000
	boss1NumBulletPerShot1     = 16
	boss1NumBulletPerShot2     = 17
	boss1NumBulletPerShot3     = 16
	boss1BulletMultiplier      = 7
	boss1Loop1                 = 6
	boss1NumFramePerShot       = 40
	boss1PVLimit1              = 250
	boss1PV                    = 500
	boss1MoveSpeed             = 2
	boss1Phase2BulletSpeed     = 5
	boss1Phase2NumFramePerShot = 15
	boss1Phase2BulletYSpeed    = 0.5
	boss1NumBulletsPhase2      = 6
	boss1DeathAnimationFrames  = 180
)

func makeBoss1(y float64) boss {
	x := float64(screenWidth - 1 + 1224/2)
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
		x:              x,
		xSize:          1224,
		y:              y,
		ySize:          1000,
		pv:             boss1PV,
		bossType:       boss1,
		points:         boss1Points,
		hitBoxes:       []bossHitBox{hitb1, hitb2},
		hurtBoxes:      []bossHitBox{hurtb1, hurtb2},
		numDeathFrames: boss1DeathAnimationFrames,
	}
}

func (b *boss) boss1Update(bs *bulletSet) bool {
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
	case 1:
		if b.pv < boss1PVLimit1 {
			b.phase = 2
		}
		b.frame++
		if b.frame >= boss1NumFramePerShot {
			b.frame = 0
			numToShoot := boss1NumBulletPerShot1
			if b.phaseLoop >= boss1Loop1 {
				numToShoot = boss1NumBulletPerShot2
			}
			if b.phaseLoop >= 2*boss1Loop1 {
				numToShoot = boss1NumBulletPerShot3
			}
			if b.phaseLoop >= 3*boss1Loop1 {
				numToShoot = boss1NumBulletPerShot2
			}
			b.phaseLoop++
			if b.phaseLoop >= 4*boss1Loop1 {
				b.phaseLoop = 0
			}
			for bNum := 0; bNum < numToShoot; bNum++ {
				increase := (b.phaseLoop % 2) * boss1BulletMultiplier
				angle := (math.Pi/2)*float64(bNum+increase)/float64(numToShoot-1+2*increase) + 3*math.Pi/4
				vx := midBoss1BulletSpeedPhase1 * math.Cos(angle)
				vy := midBoss1BulletSpeedPhase1 * -math.Sin(angle)
				bs.addBullet(bullet{
					x:     b.x + 90,
					y:     b.y - 10,
					vx:    vx,
					vy:    vy,
					image: enemyBasicBullet,
				})
			}
			return true
		}
	case 2:
		b.frame = 0
		b.y -= boss1MoveSpeed
		if b.y < float64(screenHeight)/4 {
			b.phase = 3
		}
	case 3:
		b.y += boss1MoveSpeed
		if b.y > float64(3*screenHeight)/4 {
			b.phase = 4
		}
	case 4:
		b.y -= boss1MoveSpeed
		if b.y < float64(screenHeight)/4 {
			b.phase = 5
		}
		b.frame++
		if b.frame >= boss1Phase2NumFramePerShot {
			b.frame = 0
			b.phase2AddBullets(bs)
		}
	case 5:
		b.y += boss1MoveSpeed
		if b.y > float64(3*screenHeight)/4 {
			b.phase = 4
		}
		b.frame++
		if b.frame >= boss1Phase2NumFramePerShot {
			b.frame = 0
			b.phase2AddBullets(bs)
		}
	}
	return false
}

func (b boss) phase2AddBullets(bs *bulletSet) {
	bs.addBullet(bullet{
		x:     b.x + 90,
		y:     b.y - 10,
		vx:    -boss1Phase2BulletSpeed,
		vy:    0,
		image: enemyBasicBullet,
	})
	for i := 0; i < boss1NumBulletsPhase2; i++ {
		bs.addBullet(bullet{
			x:     b.x + 90,
			y:     b.y - 10,
			vx:    -boss1Phase2BulletSpeed,
			vy:    -float64(i+1) * boss1Phase2BulletYSpeed,
			image: enemyBasicBullet,
		})
		bs.addBullet(bullet{
			x:     b.x + 90,
			y:     b.y - 10,
			vx:    -boss1Phase2BulletSpeed,
			vy:    float64(i+1) * boss1Phase2BulletYSpeed,
			image: enemyBasicBullet,
		})
	}
}

func (b *boss) boss1Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.x-b.xSize/2, b.y-b.ySize/2)
	if b.pv <= 0 {
		alpha := float64(b.deathAnimationFrame) / float64(b.numDeathFrames)
		op.ColorM.Translate(0, 0, 0, -alpha)
	}
	screen.DrawImage(
		boss1Image,
		op,
	)
}
