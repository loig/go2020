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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	welcomeStart int = iota
	welcomeHelp
	welcomeFullScreen
	welcomeInfo
	welcomeNumStates
)

func (g *game) welcomeUpdate() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch g.stateState {
		case welcomeStart:
			g.state = gameIntro
			g.stateState = 0
			g.stopMusic()
			transitionScreenLoadImages()
			infiniteMusic = music1
		case welcomeHelp:
			g.state = gameHelp
		case welcomeFullScreen:
			ebiten.SetFullscreen(!ebiten.IsFullscreen())
		case welcomeInfo:
			g.state = gameInfo
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.stateState = (g.stateState + 1) % welcomeNumStates
		g.playSound(menuSound)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.stateState = (g.stateState + welcomeNumStates - 1) % welcomeNumStates
		g.playSound(menuSound)
	}
}

func (g game) welcomeDraw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(
		titleScreenImage,
		op,
	)

	introColor := veryDarkColor //color.Gray16{0x777f}
	helpColor := introColor
	fullScreenColor := introColor
	infoColor := introColor

	switch g.stateState {
	case welcomeStart:
		introColor = textLightColor //color.White
	case welcomeHelp:
		helpColor = textLightColor //color.White
	case welcomeFullScreen:
		fullScreenColor = textLightColor
	case welcomeInfo:
		infoColor = textLightColor //color.White
	}

	s := "Start Game"
	bounds := text.BoundString(theFont, s)
	width := bounds.Max.X - bounds.Min.X
	text.Draw(screen, s, theBigFont, (screenWidth-width)/2-width/4, 650, introColor)

	s = "How to Play"
	bounds = text.BoundString(theFont, s)
	width = bounds.Max.X - bounds.Min.X
	text.Draw(screen, s, theBigFont, (screenWidth-width)/2-width/4, 700, helpColor)

	s = "Toggle fullscreen"
	bounds = text.BoundString(theFont, s)
	width = bounds.Max.X - bounds.Min.X
	text.Draw(screen, s, theBigFont, (screenWidth-width)/2-width/4, 750, fullScreenColor)

	s = "About"
	bounds = text.BoundString(theFont, s)
	width = bounds.Max.X - bounds.Min.X
	text.Draw(screen, s, theBigFont, (screenWidth-width)/2-width/4, 800, infoColor)

}
