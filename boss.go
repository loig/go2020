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
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type boss struct {
	x         float64
	xMin      float64
	xMax      float64
	xSize     float64
	y         float64
	yMin      float64
	yMax      float64
	ySize     float64
	hullSet   bool
	cHull     []point
	pv        int
	phase     int
	phaseLoop int
	bossType  int
	frame     int
	hitable   bool
	points    int
}

func (b *boss) updateBox() {
	b.xMin = b.x - b.xSize/2
	b.xMax = b.x + b.xSize/2
	b.yMin = b.y - b.ySize/2
	b.yMax = b.y + b.ySize/2
}

func (b *boss) xmin() float64 {
	return b.xMin
}

func (b *boss) xmax() float64 {
	return b.xMax
}

func (b *boss) ymin() float64 {
	return b.yMin
}

func (b *boss) ymax() float64 {
	return b.yMax
}

func (b *boss) convexHull() []point {
	if !b.hullSet {
		b.cHull = []point{
			point{b.xmin(), b.ymin()},
			point{b.xmin(), b.ymax()},
			point{b.xmax(), b.ymax()},
			point{b.xmax(), b.ymin()},
		}
		b.hullSet = true
	}
	return b.cHull
}

func (b *boss) hasCollided() {
	if b.hitable {
		b.pv--
	}
}

func (b *boss) update(bs *bulletSet) {
	switch b.bossType {
	case midBoss1:
		b.midBoss1Update(bs)
	}
	b.updateBox()
	b.hullSet = false
}

func (b *boss) draw(screen *ebiten.Image) {
	switch b.bossType {
	case midBoss1:
		b.midBoss1Draw(screen)
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

func (bs *bossSet) update(bbs *bulletSet, ps *powerUpSet, points *int) {
	if bs.numBosses == 0 {
		bs.frameSinceBattleStart = 0
		bs.totalPvMax = 0
	}
	for pos := 0; pos < bs.numBosses; pos++ {
		bs.bosses[pos].update(bbs)
		if bs.bosses[pos].isDead() {
			*points += bs.bosses[pos].points
			bs.numBosses--
			bs.bosses[pos] = bs.bosses[bs.numBosses]
			bs.bosses = bs.bosses[:bs.numBosses]
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
	}
	bs.totalPvMax += b.pv
	bs.bosses = append(bs.bosses, &b)
}
