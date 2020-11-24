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
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	topUIOffset                  = 100
	bottomUIOffset               = topUIOffset
	rightUIOffset                = 50
	leftUIOffset                 = rightUIOffset
	numFramesForDisplayingBossPV = 100
	bossPVWidth                  = 1080
	bossPVHeight                 = 50
	bonusDisplayWidth            = 1080
	bonusDisplayHeight           = 50
	lifeDisplayWidth             = 50
	lifeDisplayHeight            = 50
)

func (p player) drawUI(screen *ebiten.Image) {

	// Draw bonus selector
	var bonusImage *ebiten.Image
	switch p.currentPowerUp {
	case 0:
		bonusImage = noBonusImage
	case 1:
		bonusImage = firstBonusImage
	case 2:
		bonusImage = secondBonusImage
	case 3:
		bonusImage = thirdBonusImage
	case 4:
		bonusImage = fourthBonusImage
	}
	xTranslate := float64(screenWidth-bonusDisplayWidth) / 2
	yTranslate := float64(screenHeight - bottomUIOffset)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(xTranslate, yTranslate)
	screen.DrawImage(
		bonusImage,
		op,
	)

	// Draw lifes
	xTranslate = float64(leftUIOffset)
	yTranslate = topUIOffset - lifeDisplayHeight
	for i := 0; i < p.lives; i++ {
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(xTranslate, yTranslate)
		screen.DrawImage(
			lifeImage,
			op,
		)
		xTranslate += float64(lifeDisplayWidth)
	}

	// Display power ups names
	for puNum := 1; puNum <= pDifferentPowerUps; puNum++ {
		var s string = "-"
		if p.isAppliablePowerUp(puNum) {
			switch puNum {
			case 1:
				s = "Speed Up"
			case 2:
				s = "Better Shot"
			case 3:
				switch p.currentFire {
				case 0:
					s = "Large Shot"
				case 1:
					s = "Laser"
				case 2:
					s = "Base Shot"
				}
			case 4:
				s = "New Option"
			}
		}
		bounds := text.BoundString(theFont, s)
		width := bounds.Max.X - bounds.Min.X
		height := bounds.Max.Y - bounds.Min.Y
		xFTranslate := (screenWidth-bonusDisplayWidth)/2 + (puNum*2-1)*bonusDisplayWidth/8 - width/2
		yFTranslate := screenHeight - bottomUIOffset + bonusDisplayHeight/2 + height/2
		text.Draw(screen, s, theFont, xFTranslate, yFTranslate, color.White)
	}

	// Draw points
	s := fmt.Sprint(p.points, " Pts")
	bounds := text.BoundString(theFont, s)
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	xFTranslate := screenWidth - rightUIOffset - width
	yFTranslate := topUIOffset - height
	text.Draw(screen, s, theFont, xFTranslate, yFTranslate, color.White)

}

func (bs *bossSet) drawUI(screen *ebiten.Image) {
	if bs.numBosses >= 1 {
		var currentPV int
		for _, b := range bs.bosses {
			currentPV += b.pv
		}
		pvPortion := float64(currentPV) / float64(bs.totalPvMax)
		displayUpTo := int(math.Ceil(bossPVWidth * pvPortion))
		if bs.frameSinceBattleStart < numFramesForDisplayingBossPV {
			maxPVPortion := float64(bs.frameSinceBattleStart+1) / float64(numFramesForDisplayingBossPV)
			if maxPVPortion < pvPortion {
				displayUpTo = int(math.Ceil(bossPVWidth * maxPVPortion))
			}
			bs.frameSinceBattleStart++
		}
		xTranslate := float64(screenWidth-bossPVWidth) / 2
		yTranslate := float64(topUIOffset - bossPVHeight)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(xTranslate, yTranslate)
		screen.DrawImage(
			bs.pvBackImage,
			op,
		)
		screen.DrawImage(
			bs.pvImage.SubImage(image.Rect(0, 0, displayUpTo, bossPVHeight)).(*ebiten.Image),
			op,
		)
	}
}
