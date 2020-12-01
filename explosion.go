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
	explosionNumFrames = 2
	explosionNumSteps  = 6
	explosionHalfXSize = 25
	explosionHalfYSize = 25
)

type explosion struct {
	x     float64
	y     float64
	frame int
	step  int
	delay int
	isBig bool
}

type explosionSet struct {
	explosions    []explosion
	numExplosions int
}

func (g *game) explosionSetUpdate() {

	for pos := 0; pos < g.explosions.numExplosions; pos++ {
		if !g.level.bossBattle && !g.level.endLevel {
			g.explosions.explosions[pos].x -= firstPlanPxPerFrame
		}
		if g.explosions.explosions[pos].delay >= 0 {
			g.explosions.explosions[pos].delay--
			if g.explosions.explosions[pos].delay == 0 {
				g.playSound(enemyHurtSound)
			}
		} else {
			if g.explosions.explosions[pos].frame >= explosionNumFrames {
				g.explosions.explosions[pos].step++
				g.explosions.explosions[pos].frame = 0
				if g.explosions.explosions[pos].step >= explosionNumSteps {
					g.explosions.numExplosions--
					g.explosions.explosions[pos] = g.explosions.explosions[g.explosions.numExplosions]
					g.explosions.explosions = g.explosions.explosions[:g.explosions.numExplosions]
					pos--
				}
			} else {
				g.explosions.explosions[pos].frame++
			}
		}
	}

}

func (g game) explosionSetDraw(screen *ebiten.Image) {

	for pos := 0; pos < g.explosions.numExplosions; pos++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(g.explosions.explosions[pos].x-explosionHalfXSize, g.explosions.explosions[pos].y-explosionHalfYSize)
		if g.explosions.explosions[pos].isBig {
			screen.DrawImage(bigExplosionImages[g.explosions.explosions[pos].step], op)
		} else {
			screen.DrawImage(explosionImages[g.explosions.explosions[pos].step], op)
		}
	}

}

func (g *game) addExplosion(x, y float64, delay int, isBig bool) {
	g.explosions.explosions = append(g.explosions.explosions, explosion{x: x, y: y, delay: delay, isBig: isBig})
	g.explosions.numExplosions++
}
