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
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *game) Draw(screen *ebiten.Image) {

	switch g.state {
	case gameWelcome:
		g.welcomeDraw(screen)
	case gameHelp:
		g.helpDraw(screen)
	case gameJoypadSetup:
		g.joypadSetupDraw(screen)
	case gameInfo:
		g.infoDraw(screen)
	case gameIntro:
		g.introDraw(screen)
	case gameInLevel1, gameInLevel2, gameInLevel1Paused, gameInLevel2Paused:
		//if g.stateFrame >= framesBeforeLevel {
		g.level.draw(screen)
		g.bulletSet.draw(screen, color.RGBA{255, 0, 0, 255})
		g.enemySet.draw(screen)
		g.bossSet.draw(screen)
		g.powerUpSet.draw(screen)
		g.player.draw(screen)
		g.explosionSetDraw(screen)
		g.bossSet.drawUI(screen)
		g.player.drawUI(screen)
		//}
		if g.stateFrame < framesBeforeLevel {
			alpha := float64(g.stateFrame) / float64(framesBeforeLevel)
			op := &ebiten.DrawImageOptions{}
			op.ColorM.Translate(-1, -1, -1, -alpha)
			screen.DrawImage(levelBackground, op)
		}
		if g.state == gameInLevel1Paused || g.state == gameInLevel2Paused {
			op := &ebiten.DrawImageOptions{}
			op.ColorM.Translate(-1, -1, -1, -0.5)
			screen.DrawImage(levelBackground, op)
			s := "Paused"
			bounds := text.BoundString(theFont, s)
			width := bounds.Max.X - bounds.Min.X
			text.Draw(screen, s, theBigFont, (screenWidth-width)/2-width/4, 500, textLightColor)
		}
	case gameTransition:
		g.transitionDraw(screen)
	case gameFinished:
		g.finishedDraw(screen)
	case gameOver:
		g.gameOverDraw(screen)
	}

	if isDebug() {
		s := fmt.Sprint(ebiten.CurrentTPS()) //, ebiten.CurrentFPS())
		ebitenutil.DebugPrint(screen, s)
	}

}
