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
	"github.com/hajimehoshi/ebiten/v2/audio"
	"golang.org/x/image/font"
)

// screen size
const (
	screenWidth  = 2160
	screenHeight = 1080
)

var titleScreenImage *ebiten.Image
var gameOverScreenImage *ebiten.Image
var helpScreenImage *ebiten.Image

// level graphics
var levelFirstPlan *ebiten.Image
var levelSecondPlan *ebiten.Image
var levelThirdPlan *ebiten.Image
var levelFourthPlan *ebiten.Image
var levelBackground *ebiten.Image

// enemy graphics
var enemyBasicBullet *ebiten.Image
var staticEnemyImage *ebiten.Image
var staticFiringEnemyImage *ebiten.Image
var staticFiringDownEnemyImage *ebiten.Image
var staticFiringUpEnemyImage *ebiten.Image
var staticExplodingEnemyImage *ebiten.Image
var movingFiringEnemyImage *ebiten.Image
var movingFiringEnemyImage2 *ebiten.Image
var movingFiringEnemyImage3 *ebiten.Image

// boss graphics
var boss1Image *ebiten.Image
var midBoss1Image *ebiten.Image
var boss2Image *ebiten.Image

// ui graphics
var noBonusImage *ebiten.Image
var firstBonusImage *ebiten.Image
var secondBonusImage *ebiten.Image
var thirdBonusImage *ebiten.Image
var fourthBonusImage *ebiten.Image
var lifeImage *ebiten.Image

// font
var theFont font.Face
var theBigFont font.Face

var textLightColor color.Color = color.RGBA{205, 204, 191, 255}
var textDarkColor color.Color = color.RGBA{196, 192, 172, 255}
var scoreColor color.Color = color.RGBA{65, 71, 115, 255}

// player graphics
var optionImage *ebiten.Image
var laserImage1 *ebiten.Image
var laserImage2 *ebiten.Image
var laserImage3 *ebiten.Image
var laserImage *ebiten.Image
var playerImage *ebiten.Image
var playerBulletImage *ebiten.Image
var playerBigBulletImage *ebiten.Image
var bonusImage *ebiten.Image

// explosion graphics
var explosionImages [explosionNumSteps]*ebiten.Image
var bigExplosionImages [explosionNumSteps]*ebiten.Image

// sounds
var infiniteMusic *audio.InfiniteLoop
var music1 *audio.InfiniteLoop
var music2 *audio.InfiniteLoop
var playerShotBytes []byte
var playerHurtBytes []byte
var enemyHurtBytes []byte
var bossHurtBytes []byte
var menuBytes []byte
var getBonusBytes []byte
var useBonusBytes []byte
