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
	transitionStart int = iota
	transitionStep1
	transitionFinished
)

func (g *game) transitionUpdate() {

	if g.stateState >= transitionFinished && inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.stateFrame = 0
		g.setUpLevel2()
		g.stopMusic()
		infiniteMusic = music2
		g.state = gameInLevel2
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.stateState = transitionFinished
		return
	}

	if g.stateState < transitionFinished {
		g.stateFrame++
		if g.stateFrame >= framesPerText {
			g.stateState++
		}
	}

}

func (g *game) transitionDraw(screen *ebiten.Image) {

	if g.stateState >= transitionStep1 {
		text.Draw(screen, "Leaving the earth was not easy.", theBigFont, 500, 200, color.White)
	}

	text.Draw(screen, "Press ENTER to continue", theBigFont, 1750, 1040, color.White)
}
