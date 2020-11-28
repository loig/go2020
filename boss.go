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

type boss struct {
	x         float64
	xSize     float64
	y         float64
	ySize     float64
	pv        int
	phase     int
	phaseLoop int
	bossType  int
	frame     int
	points    int
	hitBoxes  []bossHitBox
	hurtBoxes []bossHitBox
}

type bossHitBox struct {
	x          float64
	xrel       float64
	xMin       float64
	xMax       float64
	y          float64
	yrel       float64
	yMin       float64
	yMax       float64
	xSize      float64
	ySize      float64
	hullSet    bool
	hullShape  []point
	cHull      []point
	hitable    bool
	collisions int
}

func (b *bossHitBox) updateBox() {
	b.xMin = b.x - b.xSize/2
	b.xMax = b.x + b.xSize/2
	b.yMin = b.y - b.ySize/2
	b.yMax = b.y + b.ySize/2
}

func (b *bossHitBox) xmin() float64 {
	return b.xMin
}

func (b *bossHitBox) xmax() float64 {
	return b.xMax
}

func (b *bossHitBox) ymin() float64 {
	return b.yMin
}

func (b *bossHitBox) ymax() float64 {
	return b.yMax
}

func (b *bossHitBox) convexHull() []point {
	if !b.hullSet {
		b.cHull = make([]point, len(b.hullShape))
		for i, p := range b.hullShape {
			b.cHull[i].x = b.x + p.x
			b.cHull[i].y = b.y + p.y
		}
		b.hullSet = true
	}
	return b.cHull
}

func (b *bossHitBox) hasCollided() {
	if b.hitable {
		b.collisions++
	}
}

func (b *boss) update(g *game) {
	var hasFired bool
	switch b.bossType {
	case midBoss1:
		hasFired = b.midBoss1Update(&(g.bulletSet))
	case boss1:
		hasFired = b.boss1Update(&(g.bulletSet))
	}
	if hasFired {
		g.playSound(enemyShotSound)
	}
	var wasHurt bool
	for pos := 0; pos < len(b.hitBoxes); pos++ {
		b.pv -= b.hitBoxes[pos].collisions
		if b.hitBoxes[pos].collisions > 0 {
			b.hitBoxes[pos].collisions = 0
			wasHurt = true
		}
	}
	if wasHurt {
		g.playSound(bossHurtSound)
	}
	if b.x+b.hitBoxes[0].xrel != b.hitBoxes[0].x || b.y+b.hitBoxes[0].yrel != b.hitBoxes[0].y {
		for pos := 0; pos < len(b.hitBoxes); pos++ {
			b.hitBoxes[pos].x = b.x + b.hitBoxes[pos].xrel
			b.hitBoxes[pos].y = b.y + b.hitBoxes[pos].yrel
			b.hitBoxes[pos].updateBox()
			b.hitBoxes[pos].hullSet = false
		}
		for pos := 0; pos < len(b.hurtBoxes); pos++ {
			b.hurtBoxes[pos].x = b.x + b.hurtBoxes[pos].xrel
			b.hurtBoxes[pos].y = b.y + b.hurtBoxes[pos].yrel
			b.hurtBoxes[pos].updateBox()
			b.hurtBoxes[pos].hullSet = false
		}
	}
}

func (b *boss) draw(screen *ebiten.Image) {
	switch b.bossType {
	case midBoss1:
		b.midBoss1Draw(screen)
	case boss1:
		b.boss1Draw(screen)
	}
	for pos := 0; pos < len(b.hitBoxes); pos++ {
		// draw hitBox
		cHull := b.hitBoxes[pos].convexHull()
		hullColor := color.RGBA{0, 255, 0, 255}
		for i := 0; i < len(cHull); i++ {
			ii := (i + 1) % len(cHull)
			ebitenutil.DrawLine(screen, cHull[i].x, cHull[i].y, cHull[ii].x, cHull[ii].y, hullColor)
		}
		// draw rectangle
		boxColor := color.RGBA{0, 255, 255, 255}
		ebitenutil.DrawLine(screen, b.hitBoxes[pos].xmin(), b.hitBoxes[pos].ymin(), b.hitBoxes[pos].xmax(), b.hitBoxes[pos].ymin(), boxColor)
		ebitenutil.DrawLine(screen, b.hitBoxes[pos].xmin(), b.hitBoxes[pos].ymax(), b.hitBoxes[pos].xmax(), b.hitBoxes[pos].ymax(), boxColor)
		ebitenutil.DrawLine(screen, b.hitBoxes[pos].xmin(), b.hitBoxes[pos].ymin(), b.hitBoxes[pos].xmin(), b.hitBoxes[pos].ymax(), boxColor)
		ebitenutil.DrawLine(screen, b.hitBoxes[pos].xmax(), b.hitBoxes[pos].ymin(), b.hitBoxes[pos].xmax(), b.hitBoxes[pos].ymax(), boxColor)
	}
	for pos := 0; pos < len(b.hurtBoxes); pos++ {
		// draw hitBox
		cHull := b.hurtBoxes[pos].convexHull()
		hullColor := color.RGBA{0, 255, 0, 255}
		for i := 0; i < len(cHull); i++ {
			ii := (i + 1) % len(cHull)
			ebitenutil.DrawLine(screen, cHull[i].x, cHull[i].y, cHull[ii].x, cHull[ii].y, hullColor)
		}
		// draw rectangle
		boxColor := color.RGBA{0, 255, 255, 255}
		ebitenutil.DrawLine(screen, b.hurtBoxes[pos].xmin(), b.hurtBoxes[pos].ymin(), b.hurtBoxes[pos].xmax(), b.hurtBoxes[pos].ymin(), boxColor)
		ebitenutil.DrawLine(screen, b.hurtBoxes[pos].xmin(), b.hurtBoxes[pos].ymax(), b.hurtBoxes[pos].xmax(), b.hurtBoxes[pos].ymax(), boxColor)
		ebitenutil.DrawLine(screen, b.hurtBoxes[pos].xmin(), b.hurtBoxes[pos].ymin(), b.hurtBoxes[pos].xmin(), b.hurtBoxes[pos].ymax(), boxColor)
		ebitenutil.DrawLine(screen, b.hurtBoxes[pos].xmax(), b.hurtBoxes[pos].ymin(), b.hurtBoxes[pos].xmax(), b.hurtBoxes[pos].ymax(), boxColor)
	}
}

func (b boss) isDead() bool {
	return b.pv <= 0
}

type bossSet struct {
	numBosses             int
	bosses                []*boss
	totalPvMax            int
	frameSinceBattleStart int
	pvImage               *ebiten.Image
	pvBackImage           *ebiten.Image
}

func initBossSet() bossSet {
	img1, _, err := ebitenutil.NewImageFromFile("assets/Barre-vie.png")
	if err != nil {
		panic(err)
	}

	img2, _, err := ebitenutil.NewImageFromFile("assets/Barre-vie-fond.png")
	if err != nil {
		panic(err)
	}

	return bossSet{
		numBosses:   0,
		bosses:      make([]*boss, 0),
		pvImage:     img1,
		pvBackImage: img2,
	}
}

func (g *game) bossSetUpdate() {
	if g.bossSet.numBosses == 0 {
		g.bossSet.frameSinceBattleStart = 0
		g.bossSet.totalPvMax = 0
	}
	for pos := 0; pos < g.bossSet.numBosses; pos++ {
		g.bossSet.bosses[pos].update(g)
		if g.bossSet.bosses[pos].isDead() {
			g.player.points += g.bossSet.bosses[pos].points
			g.bossSet.numBosses--
			g.bossSet.bosses[pos] = g.bossSet.bosses[g.bossSet.numBosses]
			g.bossSet.bosses = g.bossSet.bosses[:g.bossSet.numBosses]
			pos--
		}
	}
}

func (bs *bossSet) draw(screen *ebiten.Image) {
	for _, b := range bs.bosses {
		b.draw(screen)
	}
}

func (bs *bossSet) addBoss(bossType int, x, y float64) {
	bs.numBosses++
	var b boss
	switch bossType {
	case midBoss1:
		b = makeMidBoss1(x, y)
	case boss1:
		b = makeBoss1(x, y)
	}
	bs.totalPvMax += b.pv
	bs.bosses = append(bs.bosses, &b)
}
