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
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type level struct {
	firstPlanHeight  int
	secondPlanHeight int
	thirdPlanHeight  int
	spawnSequence    []spawn
	currentSpawn     int
	currentFrame     int
	lastSpawnFrame   int
	bossBattle       bool
	endLevelFrames   int
	endLevel         bool
}

type spawn struct {
	enemies    []enemySpawn
	frameDelay int
}

type enemySpawn struct {
	enemyType int
	y         float64
}

const (
	firstPlanPxPerFrame  = 4
	secondPlanPxPerFrame = 2
	thirdPlanPxPerFrame  = 1
	fourthPlanPxPerFrame = 0.25
	planImageWidth       = 3824
	framesBeforeLevel    = 180
	framesAtEndOfLevel   = 240
)

func (g *game) levelUpdate() {
	if !g.level.bossBattle && !g.level.endLevel {
		if g.level.currentSpawn >= len(g.level.spawnSequence) {
			// level finished
			g.level.endLevel = true
			g.bulletSet.numBullets = 0
			g.enemySet.numEnemies = 0
			g.level.endLevelFrames = 0
			g.audio.musicPlayer.SetVolume(0)
			return
		}
		g.level.currentFrame++
		if g.level.lastSpawnFrame+g.level.spawnSequence[g.level.currentSpawn].frameDelay == g.level.currentFrame {
			for _, e := range g.level.spawnSequence[g.level.currentSpawn].enemies {
				if e.enemyType >= midBoss1 {
					g.bossSet.addBoss(e.enemyType, e.y)
					if !g.level.bossBattle {
						for ePos := 0; ePos < g.enemySet.numEnemies; ePos++ {
							g.enemySet.enemies[ePos].vx += firstPlanPxPerFrame
						}
						for pPos := 0; pPos < g.powerUpSet.numPowerUps; pPos++ {
							g.powerUpSet.powerUps[pPos].vx += firstPlanPxPerFrame
						}
						g.level.bossBattle = true
					}
				} else {
					g.enemySet.addEnemy(e.enemyType, screenWidth-1, e.y)
				}
			}
			g.level.currentSpawn++
			g.level.lastSpawnFrame = g.level.currentFrame
		}
	} else if g.level.bossBattle {
		if g.bossSet.numBosses == 0 {
			g.level.bossBattle = false
			for ePos := 0; ePos < g.enemySet.numEnemies; ePos++ {
				g.enemySet.enemies[ePos].vx -= firstPlanPxPerFrame
			}
			for pPos := 0; pPos < g.powerUpSet.numPowerUps; pPos++ {
				g.powerUpSet.powerUps[pPos].vx -= firstPlanPxPerFrame
			}
		}
	} else if g.level.endLevel {
		g.level.endLevelFrames++
		if g.level.endLevelFrames >= framesAtEndOfLevel {
			disposeLevelImages()
			g.stopMusic()
			g.stateState = 0
			g.stateFrame = 0
			infiniteMusic = music1
			if g.state == gameInLevel1 {
				disposeLevel1Enemies()
				g.state = gameTransition
			} else {
				disposeLevel2Enemies()
				g.state = gameFinished
			}
		}
	}
}

func (l level) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(
		levelBackground,
		op,
	)

	fourthPlanShift := fourthPlanPxPerFrame * float64(l.currentFrame)
	drawFourthPlan(screen, levelFourthPlan, fourthPlanShift)

	thirdPlanStart := (thirdPlanPxPerFrame * l.currentFrame) % planImageWidth
	drawPlan(screen, levelThirdPlan, thirdPlanStart, l.thirdPlanHeight)

	secondPlanStart := (secondPlanPxPerFrame * l.currentFrame) % planImageWidth
	drawPlan(screen, levelSecondPlan, secondPlanStart, l.secondPlanHeight)

	firstPlanStart := (firstPlanPxPerFrame * l.currentFrame) % planImageWidth
	drawPlan(screen, levelFirstPlan, firstPlanStart, l.firstPlanHeight)
}

func drawPlan(screen, plan *ebiten.Image, start, planHeight int) {
	heightTranslation := screenHeight - float64(planHeight)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, heightTranslation)
	screen.DrawImage(
		plan.SubImage(image.Rect(start, 0, start+screenWidth, planHeight)).(*ebiten.Image),
		op,
	)
	if start+screenWidth > planImageWidth {
		missingPx := start + screenWidth - planImageWidth
		op2 := &ebiten.DrawImageOptions{}
		op2.GeoM.Translate(screenWidth-float64(missingPx), heightTranslation)
		screen.DrawImage(
			plan.SubImage(image.Rect(0, 0, missingPx, screenHeight)).(*ebiten.Image),
			op2,
		)
	}
}

func drawFourthPlan(screen, plan *ebiten.Image, shift float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(screenWidth-shift, 0)
	screen.DrawImage(
		plan,
		op,
	)
}

func disposeLevelImages() {
	levelFirstPlan.Dispose()
	levelSecondPlan.Dispose()
	levelThirdPlan.Dispose()
	levelFourthPlan.Dispose()
	levelBackground.Dispose()
}
