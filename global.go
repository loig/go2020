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
	"golang.org/x/image/font"
)

// screen size
const (
	screenWidth  = 2160
	screenHeight = 1080
)

// enemy graphics
var enemyBasicBullet *ebiten.Image
var staticEnemyImage *ebiten.Image
var staticFiringEnemyImage *ebiten.Image
var staticExplodingEnemyImage *ebiten.Image
var movingFiringEnemyImage *ebiten.Image

// ui graphics
var noBonusImage *ebiten.Image
var firstBonusImage *ebiten.Image
var secondBonusImage *ebiten.Image
var thirdBonusImage *ebiten.Image
var fourthBonusImage *ebiten.Image
var lifeImage *ebiten.Image

// font
var theFont font.Face
