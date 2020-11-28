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
	transitionStep1 int = iota
	transitionStep2
	transitionStep3
	transitionStep4
	transitionStep5
	transitionStep6
	transitionStep7
	transitionStep8
	transitionStep9
	transitionStep10
	transitionFinished
)

func (g *game) transitionUpdate() {

	if g.stateState >= transitionFinished && inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.stateFrame = 0
		g.state = gameInLevel2
		g.fadeOutMusic(true)
		infiniteMusic = music2
		g.setUpLevel2()
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.stateState = transitionFinished
		return
	}

	if g.stateState < transitionFinished {
		g.stateFrame++
		if g.stateFrame >= framesPerText {
			g.stateFrame = 0
			g.stateState++
		}
	}

}

func (g *game) transitionDraw(screen *ebiten.Image) {

	textPos := cutSceneInitTextPos

	if g.stateState >= transitionStep1 {
		text.Draw(screen, "Leaving our earth was not easy.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep2 {
		text.Draw(screen, "It took us many tries.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep3 {
		text.Draw(screen, "Many years.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep4 {
		text.Draw(screen, "Many generations in fact.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep5 {
		text.Draw(screen, "But we made it.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep6 {
		text.Draw(screen, "Reaching the moon was even more difficult.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep7 {
		text.Draw(screen, "But we made it.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep8 {
		text.Draw(screen, "The overmind is now in our reach.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep9 {
		text.Draw(screen, "Killing it will require greater sacrifices.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep10 {
		text.Draw(screen, "But we will make it.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	enterColor := color.Gray16{0x555f}
	s := "Press ENTER to skip"
	if g.stateState >= transitionFinished {
		enterColor = color.White
		s = "Press ENTER to continue"
	}
	text.Draw(screen, s, theBigFont, 1750, 1040, enterColor)
}
