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
)

func (g game) isLeftPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsGamepadButtonPressed(g.joypad.id, g.joypad.left)
}

func (g game) isUpPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsGamepadButtonPressed(g.joypad.id, g.joypad.up)
}

func (g game) isUpJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsGamepadButtonJustPressed(g.joypad.id, g.joypad.up)
}

func (g game) isRightPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsGamepadButtonPressed(g.joypad.id, g.joypad.right)
}

func (g game) isDownPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsGamepadButtonPressed(g.joypad.id, g.joypad.down)
}

func (g game) isDownJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsGamepadButtonJustPressed(g.joypad.id, g.joypad.down)
}

func (g game) isSpacePressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsGamepadButtonPressed(g.joypad.id, g.joypad.space)
}

func (g game) isControlJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyControl) || inpututil.IsGamepadButtonJustPressed(g.joypad.id, g.joypad.control)
}

func (g game) isEnterJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsGamepadButtonJustPressed(g.joypad.id, g.joypad.enter)
}
