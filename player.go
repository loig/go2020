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
	x          float64
	y          float64
	vx         float64
	vy         float64
	xSize      float64
	ySize      float64
	bullets    bulletSet
	lastBullet int
}

const (
	pWidth          = 45
	pHeight         = 15
	pVCap           = 5
	pAx             = 1
	pAy             = 1
	pBulletInterval = 15
	pBulletSpeed    = 6
)

func initPlayer(x, y float64) player {
	return player{
		x: x, y: y,
		xSize: pWidth, ySize: pHeight,
		bullets:    initBulletSet(),
		lastBullet: pBulletInterval,
	}
}

func (p player) draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.xmin(), p.ymin(), p.xSize, p.ySize, color.RGBA{255, 0, 0, 255})
	p.bullets.draw(screen)
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

func (p *player) update() {
	p.move()
	p.fire()
	p.bullets.update()
}

func (p *player) move() {
	var hasMovedX bool
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.vx += pAx
		hasMovedX = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if hasMovedX {
			p.vx = 0
			hasMovedX = false
		} else {
			p.vx -= pAx
			hasMovedX = true
		}
	}
	if !hasMovedX {
		if p.vx > pAx {
			p.vx -= pAx
		} else if p.vx < -pAx {
			p.vx += pAx
		} else {
			p.vx = 0
		}
	}
	var hasMovedY bool
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.vy += pAy
		hasMovedY = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if hasMovedY {
			p.vy = 0
			hasMovedY = false
		} else {
			p.vy -= pAy
			hasMovedY = true
		}
	}
	if !hasMovedY {
		if p.vy > pAy {
			p.vy -= pAy
		} else if p.vy < -pAy {
			p.vy += pAy
		} else {
			p.vy = 0
		}
	}
	if p.vx > pVCap {
		p.vx = pVCap
	} else if p.vx < -pVCap {
		p.vx = -pVCap
	}
	if p.vy > pVCap {
		p.vy = pVCap
	} else if p.vy < -pVCap {
		p.vy = -pVCap
	}
	p.x += p.vx
	p.y += p.vy
}

func (p *player) fire() {
	p.lastBullet++
	if p.lastBullet >= pBulletInterval &&
		ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.lastBullet = 0
		p.bullets.addBullet(bullet{
			x: p.x, y: p.y,
			vx: pBulletSpeed, vy: 0,
			ax: 0, ay: 0,
		})
	}
}
