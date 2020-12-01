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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	framesPerText       = 178
	cutSceneInitTextPos = 250
	cutSceneTextSep     = 50
)

const (
	introStep1 int = iota
	introStep2
	introStep3
	introStep4
	introStep5
	introStep6
	introStep7
	introStep8
	introStep9
	introStep10
	introFinished
)

func (g *game) introUpdate() {

	if g.stateState >= introFinished && inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.stateFrame = 0
		g.state = gameInLevel1
		disposeFirstImages()
		transitionScreenDisposeImages()
		g.fadeOutMusic(true)
		infiniteMusic = music2
		g.setUpLevel1()
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.stateState = introFinished
		return
	}

	if g.stateState < introFinished {
		g.stateFrame++
		if g.stateFrame >= framesPerText {
			g.stateState++
			g.stateFrame = 0
		}
	}

}

func (g *game) introDraw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(transitionScreenImage, op)

	textPos := cutSceneInitTextPos

	if g.stateState >= introStep1 {
		displayCutSceneText("Generations ago, they arrived.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep2 {
		displayCutSceneText("Nobody knew where they came from.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep3 {
		displayCutSceneText("Nobody knew why they came here.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep4 {
		displayCutSceneText("Nobody knew nothing, but many died.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep5 {
		displayCutSceneText("For long we believed they would leave.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep6 {
		displayCutSceneText("After centuries we had to change our mind.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep7 {
		displayCutSceneText("It took us even more time to strike back.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep8 {
		displayCutSceneText("Their overmind is on the moon.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep9 {
		displayCutSceneText("Killing it will require sacrifices.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep10 {
		displayCutSceneText("As a first step, we need to leave the ground.", textPos, screen)
	}

	enterColor := veryDarkColor //color.Gray16{0x555f}
	s := "Press ENTER to skip"
	if g.stateState >= introFinished {
		enterColor = textLightColor
		s = "Press ENTER to continue"
	}
	text.Draw(screen, s, theBigFont, 1750, 1040, enterColor)
}

func displayCutSceneText(s string, ypos int, screen *ebiten.Image) {
	//_, ad := font.BoundString(theFont, s)
	//width := bounds.Max.X.Floor() - bounds.Min.X.Floor()
	//width := ad.Floor()
	//log.Print(width, ad)
	//xpos := (screenWidth - width) / 2
	xpos := 750
	text.Draw(screen, s, theBigFont, xpos, ypos, textLightColor)
}
