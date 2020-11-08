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

type enemy struct {
	x                           float64
	y                           float64
	vx                          float64
	vy                          float64
	framesSinceLastBullet       int
	nextBullet                  int
	bulletSequence              []bulletShot
	framesSinceLastAcceleration int
	nextAcceleration            int
	accelerationSequence        []acceleration
	xSize                       float64
	ySize                       float64
	pv                          int
}

type bulletShot struct {
	bullets  []bullet
	interval int
}

type acceleration struct {
	ax       float64
	ay       float64
	interval int
}

func (e enemy) draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, e.xmin(), e.ymin(), e.xSize, e.ySize, color.RGBA{155, 0, 0, 255})
}

func (e *enemy) xmin() float64 {
	return e.x - e.xSize/2
}

func (e *enemy) xmax() float64 {
	return e.x + e.xSize/2
}

func (e *enemy) ymin() float64 {
	return e.y - e.ySize/2
}

func (e *enemy) ymax() float64 {
	return e.y + e.ySize/2
}

func (e *enemy) hasCollided() {
	e.pv--
}

func (e enemy) isOut() bool {
	return e.pv <= 0 || e.xmax() < 0 || e.ymax() < 0 || e.xmin() >= screenWidth || e.ymin() >= screenHeight
}

func (e *enemy) update(bs *bulletSet) {
	e.vx += e.accelerationSequence[e.nextAcceleration].ax
	e.vy += e.accelerationSequence[e.nextAcceleration].ay
	e.x += e.vx
	e.y += e.vy
	e.framesSinceLastAcceleration++
	if e.framesSinceLastAcceleration >= e.accelerationSequence[e.nextAcceleration].interval {
		e.framesSinceLastAcceleration = 0
		e.nextAcceleration = (e.nextAcceleration + 1) % len(e.accelerationSequence)
	}
	e.framesSinceLastBullet++
	if e.framesSinceLastBullet >= e.bulletSequence[e.nextBullet].interval {
		for _, b := range e.bulletSequence[e.nextBullet].bullets {
			b.x = e.x
			b.y = e.y
			bs.addBullet(b)
		}
		e.framesSinceLastBullet = 0
		e.nextBullet = (e.nextBullet + 1) % len(e.bulletSequence)
	}
}

type enemySet struct {
	numEnemies int
	enemies    []*enemy
}

func initEnemySet() enemySet {
	return enemySet{
		numEnemies: 0,
		enemies:    make([]*enemy, 0),
	}
}

func (es *enemySet) addTestEnemy() {
	es.numEnemies++
	e := makeTestEnemy()
	es.enemies = append(es.enemies, &e)
}

func (es *enemySet) update(bs *bulletSet) {
	for pos := 0; pos < es.numEnemies; pos++ {
		es.enemies[pos].update(bs)
		if es.enemies[pos].isOut() {
			es.numEnemies--
			es.enemies[pos] = es.enemies[es.numEnemies]
			es.enemies = es.enemies[:es.numEnemies]
		}
	}
}

func (es *enemySet) draw(screen *ebiten.Image) {
	for _, e := range es.enemies {
		e.draw(screen)
	}
}

func makeTestEnemy() enemy {
	return enemy{
		x: 799, y: 300,
		vx: -5, vy: 0,
		xSize: 25, ySize: 15,
		pv: 2,
		bulletSequence: []bulletShot{
			bulletShot{
				bullets: []bullet{
					bullet{vx: -10, vy: 0, ax: 0, ay: 0},
				},
				interval: 30,
			},
			bulletShot{
				bullets: []bullet{
					bullet{vx: -10, vy: 5, ax: 0, ay: 0},
					bullet{vx: -10, vy: -5, ax: 0, ay: 0},
				},
				interval: 5,
			},
		},
		accelerationSequence: []acceleration{
			acceleration{ax: 0, ay: 1, interval: 5},
			acceleration{ax: 0, ay: 0, interval: 10},
			acceleration{ax: 0, ay: -1, interval: 10},
			acceleration{ax: 0, ay: 0, interval: 10},
			acceleration{ax: 0, ay: 1, interval: 5},
		},
	}
}
