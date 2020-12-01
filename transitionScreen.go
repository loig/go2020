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
		transitionScreenDisposeImages()
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

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(transitionScreenImage, op)

	textPos := cutSceneInitTextPos

	if g.stateState >= transitionStep1 {
		displayCutSceneText("Leaving our earth was not easy.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep2 {
		displayCutSceneText("It took us many tries.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep3 {
		displayCutSceneText("Many years.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep4 {
		displayCutSceneText("Many generations in fact.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep5 {
		displayCutSceneText("But we made it.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep6 {
		displayCutSceneText("Reaching the moon was even more difficult.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep7 {
		displayCutSceneText("But we made it.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep8 {
		displayCutSceneText("The overmind is now in our reach.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep9 {
		displayCutSceneText("To deliver the final blow will require sacrifice.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= transitionStep10 {
		displayCutSceneText("But we will make it.", textPos, screen)
	}

	enterColor := textDarkColor //color.Gray16{0x555f}
	s := "Press ENTER to skip"
	if g.stateState >= transitionFinished {
		enterColor = textLightColor //color.White
		s = "Press ENTER to continue"
	}
	text.Draw(screen, s, theBigFont, 1750, 1040, enterColor)
}

func transitionScreenLoadImages() {
	img, _, err := ebitenutil.NewImageFromFile("assets/cutscene.png")
	if err != nil {
		panic(err)
	}
	transitionScreenImage = img
}

func transitionScreenDisposeImages() {
	transitionScreenImage.Dispose()
}
