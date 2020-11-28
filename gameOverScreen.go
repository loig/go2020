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

func (g *game) gameOverUpdate() {

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		// go to title screen
		loadFirstImages()
		g.stopMusic()
		infiniteMusic = music2
		g.stateFrame = 0
		g.state = gameWelcome
		g.stateState = 0
	}

}

func (g *game) gameOverDraw(screen *ebiten.Image) {

	text.Draw(screen, "Game Over", theBigFont, 500, 200, color.White)

	text.Draw(screen, "Press ENTER to restart", theBigFont, 1800, 1040, color.White)
}
