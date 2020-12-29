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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	joypadUnplugged int = iota
	joypadLeft
	joypadUp
	joypadRight
	joypadDown
	joypadSpace
	joypadControl
	joypadEnter
	joypadConfigured
)

type joypad struct {
	id                                           ebiten.GamepadID
	left, up, right, down, space, control, enter ebiten.GamepadButton
}

func (g *game) joypadSetupUpdate() {

	if g.isEnterJustPressed() {
		g.stateState = welcomeJoypadSetup
		g.state = gameWelcome
		return
	}

	switch g.stateState {
	case joypadUnplugged:
		if g.isJoypadPlugged() {
			g.stateState++
		}
	case joypadLeft, joypadUp, joypadRight, joypadDown, joypadSpace, joypadControl, joypadEnter:
		if g.setUpJoypadButton() {
			g.stateState++
		}
	default:
	}

}

func (g *game) joypadSetupDraw(screen *ebiten.Image) {
	var s string
	switch g.stateState {
	case joypadUnplugged:
		s = "No gamepad detected, plug one (or press enter to quit)"
	case joypadLeft:
		s = "Press LEFT on your gamepad (or press enter to quit)"
	case joypadUp:
		s = "Press UP on your gamepad (or press enter to quit)"
	case joypadRight:
		s = "Press RIGHT on your gamepad (or press enter to quit)"
	case joypadDown:
		s = "Press DOWN on your gamepad (or press enter to quit)"
	case joypadSpace:
		s = "Press a gamepad button to replace SPACE – used to shoot (or press enter to quit)"
	case joypadControl:
		s = "Press a gamepad button to replace CTRL – used to activate power ups (or press enter to quit)"
	case joypadEnter:
		s = "Press a gamepad button to replace ENTER – used to validate things (or press enter to quit)"
	case joypadConfigured:
		s = "Gamepad successfuly configured, press enter to quit"
	}
	bounds := text.BoundString(theFont, s)
	width := bounds.Max.X - bounds.Min.X
	text.Draw(screen, s, theBigFont, (screenWidth-width)/2-width/4, 500, textLightColor)
}

func (g *game) isJoypadPlugged() bool {

	gpIDs := ebiten.GamepadIDs()
	if len(gpIDs) <= 0 {
		return false
	}
	g.joypad.id = gpIDs[0]
	return true
}

func (g *game) setUpJoypadButton() bool {
	maxButton := ebiten.GamepadButton(ebiten.GamepadButtonNum(g.joypad.id))
	for b := ebiten.GamepadButton(g.joypad.id); b < maxButton; b++ {
		if inpututil.IsGamepadButtonJustPressed(g.joypad.id, b) {
			switch g.stateState {
			case joypadLeft:
				g.joypad.left = b
				if isDebug() {
					log.Print("Joypad left: ", b)
				}
			case joypadUp:
				g.joypad.up = b
				if isDebug() {
					log.Print("Joypad up: ", b)
				}
			case joypadRight:
				g.joypad.right = b
				if isDebug() {
					log.Print("Joypad right: ", b)
				}
			case joypadDown:
				g.joypad.down = b
				if isDebug() {
					log.Print("Joypad down: ", b)
				}
			case joypadSpace:
				g.joypad.space = b
				if isDebug() {
					log.Print("Joypad space: ", b)
				}
			case joypadControl:
				g.joypad.control = b
				if isDebug() {
					log.Print("Joypad ctrl: ", b)
				}
			case joypadEnter:
				g.joypad.enter = b
				if isDebug() {
					log.Print("Joypad enter: ", b)
				}
			}
			return true
		}
	}
	return false
}
