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

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	finishedStep1 int = iota
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

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(transitionScreenImage, op)

	textPos := cutSceneInitTextPos

	if g.stateState >= finishedStep1 {
		displayCutSceneText("We did it.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= finishedStep2 {
		displayCutSceneText("We finally did it.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= finishedStep3 {
		displayCutSceneText("The overmind is down.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= finishedStep4 {
		displayCutSceneText("At the very moment it died, they all died.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= finishedStep5 {
		displayCutSceneText("Now we can rebuild.", textPos, screen)
	}

	textPos += cutSceneTextSep

	if g.stateState >= finishedStep6 {
		displayCutSceneText("Start a new civilisation from our ashes.", textPos, screen)
	}

	textPos += cutSceneTextSep
	textPos += cutSceneTextSep
	textPos += cutSceneTextSep
	textPos += cutSceneTextSep

	if g.stateState >= finishedFinished {
		s := fmt.Sprint("Final score: ", g.player.points)
		displayCutSceneText(s, textPos, screen)
	}
}
