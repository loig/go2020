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
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type enemy struct {
	x                           float64
	y                           float64
	xMin                        float64
	xMax                        float64
	yMin                        float64
	yMax                        float64
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
	powerUpProba                int
	deathBlow                   []bullet
	points                      int
	hullSet                     bool
	cHull                       []point
	hullAt00                    []point
	image                       *ebiten.Image
	isAnimated                  bool
	framePerImage               int
	framesSinceLastImage        int
	images                      []*ebiten.Image
	currentImage                int
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
	if e.image != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(e.xmin(), e.ymin())
		screen.DrawImage(
			e.image,
			op,
		)
	}
	cHull := e.convexHull()
	for i := 0; i < len(cHull); i++ {
		ii := (i + 1) % len(cHull)
		ebitenutil.DrawLine(screen, cHull[i].x, cHull[i].y, cHull[ii].x, cHull[ii].y, color.RGBA{255, 0, 0, 255})
	}
}

func (e *enemy) xmin() float64 {
	return e.xMin
}

func (e *enemy) xmax() float64 {
	return e.xMax
}

func (e *enemy) ymin() float64 {
	return e.yMin
}

func (e *enemy) ymax() float64 {
	return e.yMax
}

func (e *enemy) convexHull() []point {
	if !e.hullSet {
		e.cHull = make([]point, len(e.hullAt00))
		for i, p := range e.hullAt00 {
			e.cHull[i].x = p.x + e.x
			e.cHull[i].y = p.y + e.y
		}
	}
	return e.cHull
}

func (e *enemy) hasCollided() {
	e.pv--
}

func (e enemy) isOut() bool {
	return e.pv <= 0 || e.xmax() < 0 || e.ymax() < 0 || e.xmin() >= screenWidth || e.ymin() >= screenHeight
}

func (e *enemy) update(g *game) {
	if e.accelerationSequence != nil {
		e.vx += e.accelerationSequence[e.nextAcceleration].ax
		e.vy += e.accelerationSequence[e.nextAcceleration].ay
	}
	e.x += e.vx
	e.y += e.vy
	e.hullSet = false
	e.cHull = nil
	e.xMin = e.x - e.xSize/2
	e.xMax = e.x + e.xSize/2
	e.yMin = e.y - e.ySize/2
	e.yMax = e.y + e.ySize/2
	if e.accelerationSequence != nil {
		e.framesSinceLastAcceleration++
		if e.framesSinceLastAcceleration >= e.accelerationSequence[e.nextAcceleration].interval {
			e.framesSinceLastAcceleration = 0
			e.nextAcceleration = (e.nextAcceleration + 1) % len(e.accelerationSequence)
		}
	}
	if e.bulletSequence != nil {
		e.framesSinceLastBullet++
		if e.framesSinceLastBullet >= e.bulletSequence[e.nextBullet].interval {
			for _, b := range e.bulletSequence[e.nextBullet].bullets {
				b.x += e.x
				b.y += e.y
				b.image = enemyBasicBullet
				g.bulletSet.addBullet(b)
			}
			if len(e.bulletSequence[e.nextBullet].bullets) >= 1 {
				g.playSound(enemyShotSound)
			}
			e.framesSinceLastBullet = 0
			e.nextBullet = (e.nextBullet + 1) % len(e.bulletSequence)
		}
	}
	if e.isAnimated {
		e.framesSinceLastImage++
		if e.framesSinceLastImage >= e.framePerImage {
			e.currentImage = (e.currentImage + 1) % len(e.images)
			e.image = e.images[e.currentImage]
			e.framesSinceLastImage = 0
		}
	}
}

func (e enemy) deathAction(bs *bulletSet, ps *powerUpSet, points *int, bossBattle bool) {
	*points += e.points
	// gen power up
	if rand.Intn(e.powerUpProba) == 0 {
		vx := float64(-firstPlanPxPerFrame)
		if bossBattle {
			vx = 0
		}
		ps.addPowerUp(powerUp{
			x: e.x, y: e.y,
			vx: vx, vy: 0,
		})
	}
	// gen bullets
	for _, b := range e.deathBlow {
		b.x = e.x
		b.y = e.y
		b.image = enemyBasicBullet
		bs.addBullet(b)
	}
	// gen enemies
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

func (g *game) enemySetUpdate() {
	for pos := 0; pos < g.enemySet.numEnemies; pos++ {
		g.enemySet.enemies[pos].update(g)
		if g.enemySet.enemies[pos].isOut() {
			if g.enemySet.enemies[pos].pv <= 0 {
				g.enemySet.enemies[pos].deathAction(&(g.bulletSet), &(g.powerUpSet), &(g.player.points), g.level.bossBattle)
				g.playSound(enemyHurtSound)
			}
			g.enemySet.numEnemies--
			g.enemySet.enemies[pos] = g.enemySet.enemies[g.enemySet.numEnemies]
			g.enemySet.enemies = g.enemySet.enemies[:g.enemySet.numEnemies]
			pos--
		}
	}
}

func (es *enemySet) draw(screen *ebiten.Image) {
	for pos := 0; pos < es.numEnemies; pos++ {
		es.enemies[pos].draw(screen)
	}
}

func (es *enemySet) addEnemy(enemyType int, x, y float64) {
	es.numEnemies++
	var e enemy
	switch enemyType {
	case staticEnemy:
		e = makeStaticEnemy(x, y)
	case staticExplodingEnemy:
		e = makeStaticExplodingEnemy(x, y)
	case staticFiringEnemy:
		e = makeStaticFiringEnemy(x, y)
	case staticFiringUpEnemy:
		e = makeStaticFiringUpEnemy(x, y)
	case staticFiringDownEnemy:
		e = makeStaticFiringDownEnemy(x, y)
	case movingFiringEnemy:
		e = makeMovingFiringEnemy(x, y)
	}
	es.enemies = append(es.enemies, &e)
}
