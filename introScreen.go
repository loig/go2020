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

const framesPerText = 250

const (
	introStart int = iota
	introStep1
	introFinished
)

func (g *game) introUpdate() {

	if g.stateState >= introFinished && inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.stateFrame = 0
		g.state = gameInLevel1
		disposeFirstImages()
		g.stopMusic()
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
		}
	}

}

func (g *game) introDraw(screen *ebiten.Image) {

	if g.stateState >= introStep1 {
		text.Draw(screen, "Generations ago, they arrived.", theBigFont, 500, 200, color.White)
	}

	text.Draw(screen, "Press ENTER to continue", theBigFont, 1800, 1040, color.White)
}
