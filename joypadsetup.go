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
	"math"

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
	leftrightAxis, updownaxis                    int
	useAxis                                      bool
	lastAxisValues                               []float64
	axisUpJustUsed, axisDownJustUsed             bool
}

func (g *game) recordAxis() {
	for axis := 0; axis < ebiten.GamepadAxisNum(g.joypad.id); axis++ {
		if axis >= len(g.joypad.lastAxisValues) {
			g.joypad.lastAxisValues = append(g.joypad.lastAxisValues, ebiten.GamepadAxis(g.joypad.id, axis))
		} else {
			g.joypad.lastAxisValues[axis] = ebiten.GamepadAxis(g.joypad.id, axis)
		}
		if isDebug() {
			log.Print("Axis ", axis, ": ", ebiten.GamepadAxis(g.joypad.id, axis))
		}
	}
}

func (g *game) joypadSetupUpdate() {

	if g.stateFrame == 0 {
		g.recordAxis()
		g.stateFrame++
	}

	if g.isEnterJustPressed() {
		g.stateState = welcomeJoypadSetup
		g.stateFrame = 0
		g.state = gameWelcome
		return
	}

	switch g.stateState {
	case joypadUnplugged:
		if g.isJoypadPlugged() {
			g.stateState++
		}
	case joypadLeft, joypadUp, joypadSpace, joypadControl, joypadEnter:
		if g.setUpJoypadButton() {
			g.stateState++
		}
	case joypadRight, joypadDown:
		if g.joypad.useAxis {
			g.stateState++
			return
		}
		if g.setUpJoypadButton() {
			g.stateState++
		}
	default:
	}

	g.recordAxis()

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
	// check if some axis is used
	maxAxis := 0
	maxAxisValue := 0.0
	for axis := 0; axis < ebiten.GamepadAxisNum(g.joypad.id); axis++ {
		if len(g.joypad.lastAxisValues) > axis &&
			math.Abs(ebiten.GamepadAxis(g.joypad.id, axis)-g.joypad.lastAxisValues[axis]) > 0.1 {
			axisValue := math.Abs(ebiten.GamepadAxis(g.joypad.id, axis))
			if axisValue > maxAxisValue {
				maxAxisValue = axisValue
				maxAxis = axis
			}
		}
	}
	if maxAxisValue >= 0.5 {
		switch g.stateState {
		case joypadLeft:
			g.joypad.useAxis = true
			g.joypad.leftrightAxis = maxAxis
			if isDebug() {
				log.Print("Joypad left/right: axis ", maxAxis)
			}
			return true
		case joypadUp:
			if maxAxis != g.joypad.leftrightAxis {
				g.joypad.updownaxis = maxAxis
				if isDebug() {
					log.Print("Joypad up/down: axis ", maxAxis)
				}
				return true
			}
		}
	}
	return false
}
