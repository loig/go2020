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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	finishedStart int = iota
	finishedStep1
	finishedStep2
	finishedStep3
	finishedStep4
	finishedStep5
	finishedStep6
	dummyFinishedStep1
	dummyFinishedStep2
	finishedFinished
)

func (g *game) finishedUpdate() {

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.stateState = finishedFinished
		return
	}

	if g.stateState < finishedFinished {
		g.stateFrame++
		if g.stateFrame >= framesPerText {
			g.stateFrame = 0
			g.stateState++
		}
	}

}

func (g *game) finishedDraw(screen *ebiten.Image) {

	textPos := cutSceneInitTextPos

	if g.stateState >= finishedStep1 {
		text.Draw(screen, "We did it.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= finishedStep2 {
		text.Draw(screen, "We finally did it.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= finishedStep3 {
		text.Draw(screen, "The overmind is down.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= finishedStep4 {
		text.Draw(screen, "At the very moment it died, they all died.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= finishedStep5 {
		text.Draw(screen, "Now we can rebuild.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep

	if g.stateState >= finishedStep6 {
		text.Draw(screen, "Start a new civilisation.", theBigFont, cutSceneXTextPos, textPos, color.White)
	}

	textPos += cutSceneTextSep
	textPos += cutSceneTextSep
	textPos += cutSceneTextSep
	textPos += cutSceneTextSep
	textPos += cutSceneTextSep
	textPos += cutSceneTextSep
	textPos += cutSceneTextSep

	if g.stateState >= finishedFinished {
		s := fmt.Sprint("Final score: ", g.player.points)
		text.Draw(screen, s, theBigFont, cutSceneXTextPos+400, textPos, color.White)
	}
}
