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
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	//defer profile.Start().Stop()

	//defer profile.Start(profile.MemProfile).Stop()

	ebiten.SetWindowSize(screenWidth/2, screenHeight/2)
	ebiten.SetWindowTitle("Game Off 2020")
	if err := ebiten.RunGame(initGame()); err != nil {
		log.Fatal(err)
	}
}
