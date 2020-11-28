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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	framesPerText       = 178
	cutSceneInitTextPos = 200
	cutSceneTextSep     = 50
	cutSceneXTextPos    = 500
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

	textPos := cutSceneInitTextPos

	if g.stateState >= introStep1 {
		text.Draw(screen, "Generations ago, they arrived.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep2 {
		text.Draw(screen, "Nobody knew where they came from.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep3 {
		text.Draw(screen, "Nobody knew why they came here.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep4 {
		text.Draw(screen, "Nobody knew nothing, but many died.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep5 {
		text.Draw(screen, "For long we believed they would leave.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep6 {
		text.Draw(screen, "After centuries we had to change our mind.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep7 {
		text.Draw(screen, "I took us even more time to strike back.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep8 {
		text.Draw(screen, "Their overmind is on the moon, we know for sure.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep9 {
		text.Draw(screen, "Destroying it will require sacrifices.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= introStep10 {
		text.Draw(screen, "As a first step, we need to leave the ground.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	enterColor := color.Gray16{0x555f}
	s := "Press ENTER to skip"
	if g.stateState >= introFinished {
		enterColor = color.White
		s = "Press ENTER to continue"
	}
	text.Draw(screen, s, theBigFont, 1750, 1040, enterColor)
}
