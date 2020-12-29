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
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *game) infoUpdate() {
	if g.isEnterJustPressed() {
		g.state = gameWelcome
	}
}

func (g *game) infoDraw(screen *ebiten.Image) {
	s := "A game for Game Off 2020, made by:\n\n          Cécile Dumont (graphics)\n\n                              and\n\n       Loïg Jezequel (programming)\n\n\nSource code is under GPL-3.0 License\n                and can be found at\n      https://github.com/loig/go2020"
	bounds := text.BoundString(theFont, s)
	width := bounds.Max.X - bounds.Min.X
	text.Draw(
		screen,
		s,
		theBigFont, screenWidth/2-3*width/4, 300, color.White,
	)
	text.Draw(screen, "Press ENTER to quit", theBigFont, 1800, 1040, color.White)
}
