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
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loig/go2020/assets"
)

func (g *game) setUpLevel1() {

	loadLevel1Enemies()
	loadLevel1Background()

	var l level
	l.firstPlanHeight = 394
	l.secondPlanHeight = 336
	l.thirdPlanHeight = 388
	l.fourthPlanYShift = 48
	l.fourthPlanTotalXShift = 1080
	l.levelTotalFrames = 8661
	l.spawnSequence = level1SpawnSequence

	g.level = l
	g.bulletSet = initBulletSet()
	g.player = initPlayer()
	g.enemySet = initEnemySet()
	g.bossSet = initBossSet()
	g.powerUpSet = initPowerUpSet()
}

func loadLevel1Enemies() {
	img, _, err := image.Decode(bytes.NewReader(assets.Ennemi1))
	//img, _, err := ebitenutil.NewImageFromFile("assets/Ennemi1.png")
	if err != nil {
		panic(err)
	}
	staticEnemyImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Ennemi2))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Ennemi2.png")
	if err != nil {
		panic(err)
	}
	staticExplodingEnemyImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Ennemi3))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Ennemi3.png")
	if err != nil {
		panic(err)
	}
	staticFiringEnemyImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Ennemi3b))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Ennemi3b.png")
	if err != nil {
		panic(err)
	}
	staticFiringDownEnemyImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Ennemi3c))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Ennemi3c.png")
	if err != nil {
		panic(err)
	}
	staticFiringUpEnemyImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Boss1))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Boss1.png")
	if err != nil {
		panic(err)
	}
	boss1Image = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Midboss))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Midboss.png")
	if err != nil {
		panic(err)
	}
	midBoss1Images[0] = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Midboss2))
	//	img, _, err = ebitenutil.NewImageFromFile("assets/Midboss2.png")
	if err != nil {
		panic(err)
	}
	midBoss1Images[1] = ebiten.NewImageFromImage(img)
	midBoss1Images[3] = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Midboss3))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Midboss3.png")
	if err != nil {
		panic(err)
	}
	midBoss1Images[2] = ebiten.NewImageFromImage(img)
}

func disposeLevel1Enemies() {
	staticEnemyImage.Dispose()
	staticFiringEnemyImage.Dispose()
	staticFiringUpEnemyImage.Dispose()
	staticFiringDownEnemyImage.Dispose()
	staticExplodingEnemyImage.Dispose()
	boss1Image.Dispose()
	midBoss1Images[0].Dispose()
	midBoss1Images[1].Dispose()
	midBoss1Images[2].Dispose()
	midBoss1Images[3].Dispose()
}

func loadLevel1Background() {
	img, _, err := image.Decode(bytes.NewReader(assets.Montagnes1))
	//img, _, err := ebitenutil.NewImageFromFile("assets/Montagnes-1.png")
	if err != nil {
		panic(err)
	}
	levelFirstPlan = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Montagnes2))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Montagnes-2.png")
	if err != nil {
		panic(err)
	}
	levelSecondPlan = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Montagnes3))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Montagnes-3.png")
	if err != nil {
		panic(err)
	}
	levelThirdPlan = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Lune))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Lune.png")
	if err != nil {
		panic(err)
	}
	levelFourthPlan = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Etoiles))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Etoiles.png")
	if err != nil {
		panic(err)
	}
	levelBackground = ebiten.NewImageFromImage(img)
}

var level1SpawnSequence []spawn = []spawn{
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 4},
		},
		frameDelay: 600,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 4},
		},
		frameDelay: 250,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(6*screenHeight) / 9},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 8},
		},
		frameDelay: 70,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 8},
		},
		frameDelay: 100,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 8},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 8},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(6*screenHeight) / 8},
		},
		frameDelay: 60,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 11},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 11},
		},
		frameDelay: 60,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 10},
		},
		frameDelay: 45,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 2},
		},
		frameDelay: 55,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 4},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 4},
		},
		frameDelay: 60,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 3},
		},
		frameDelay: 45,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 8},
		},
		frameDelay: 33,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 8},
		},
		frameDelay: 70,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 11},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 11},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 9},
		},
		frameDelay: 20,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 5},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 8},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 6},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 5},
		},
		frameDelay: 56,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 3},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 2},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 8},
		},
		frameDelay: 34,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 11},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 11},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 11},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 15},
		},
		frameDelay: 20,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(screenHeight) / 2},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 5},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(14*screenHeight) / 15},
		},
		frameDelay: 35,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 11},
		},
		frameDelay: 40,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(5*screenHeight) / 12},
		},
		frameDelay: 20,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 11},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 13},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 4},
		},
		frameDelay: 20,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(screenHeight) / 2},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 15},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 7},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(2*screenHeight) / 5},
		},
		frameDelay: 20,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 11},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 6},
		},
		frameDelay: 40,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 17},
		},
		frameDelay: 28,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(3*screenHeight) / 7},
		},
		frameDelay: 52,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(11*screenHeight) / 13},
		},
		frameDelay: 20,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(12*screenHeight) / 13},
		},
		frameDelay: 35,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(screenHeight) / 2},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 7},
		},
		frameDelay: 40,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(3*screenHeight) / 4},
		},
		frameDelay: 20,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(screenHeight) / 4},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 5},
		},
		frameDelay: 20,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(3*screenHeight) / 5},
		},
		frameDelay: 29,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 15},
		},
		frameDelay: 43,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(14*screenHeight) / 15},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(screenHeight) / 2},
		},
		frameDelay: 28,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(screenHeight) / 4},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(3*screenHeight) / 4},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 5},
		},
		frameDelay: 37,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: midBoss1, y: float64(screenHeight) / 2},
		},
		frameDelay: 525,
	},
	// Part 2
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 15},
		},
		frameDelay: 300,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 15},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 15},
		},
		frameDelay: 70,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 15},
		},
		frameDelay: 35,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(10*screenHeight) / 15},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(11*screenHeight) / 15},
		},
		frameDelay: 35,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(3*screenHeight) / 15},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(6*screenHeight) / 15},
		},
		frameDelay: 38,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 15},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 15},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 15},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 15},
		},
		frameDelay: 35,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringDownEnemy, y: float64(screenHeight) / 20},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(3*screenHeight) / 15},
		},
		frameDelay: 40,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 15},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 15},
		},
		frameDelay: 70,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 15},
		},
		frameDelay: 35,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringUpEnemy, y: float64(19*screenHeight) / 20},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(6*screenHeight) / 15},
		},
		frameDelay: 38,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(5*screenHeight) / 15},
		},
		frameDelay: 35,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 20},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 15},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 15},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 15},
		},
		frameDelay: 70,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 15},
		},
		frameDelay: 35,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 15},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(7*screenHeight) / 15},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(5*screenHeight) / 15},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 15},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringEnemy, y: float64(screenHeight) / 2},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 8},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 6},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 5},
		},
		frameDelay: 56,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 3},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(screenHeight) / 2},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 8},
		},
		frameDelay: 34,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 11},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 11},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 15},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 15},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringEnemy, y: float64(screenHeight) / 3},
			enemySpawn{enemyType: staticFiringEnemy, y: float64(screenHeight) / 2},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringEnemy, y: float64(2*screenHeight) / 3},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 8},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 6},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(screenHeight) / 5},
		},
		frameDelay: 56,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringEnemy, y: float64(screenHeight) / 4},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 6},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 5},
		},
		frameDelay: 56,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 3},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 2},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(7*screenHeight) / 8},
		},
		frameDelay: 34,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 11},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 11},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 15},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringDownEnemy, y: float64(screenHeight) / 15},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 2},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(7*screenHeight) / 8},
		},
		frameDelay: 34,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(9*screenHeight) / 11},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(screenHeight) / 11},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringDownEnemy, y: float64(screenHeight) / 17},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 8},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 6},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 5},
		},
		frameDelay: 56,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(2*screenHeight) / 3},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 2},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 8},
		},
		frameDelay: 34,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringUpEnemy, y: float64(14*screenHeight) / 15},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(screenHeight) / 11},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 15},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringUpEnemy, y: float64(16*screenHeight) / 17},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 3},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 3},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringUpEnemy, y: float64(12*screenHeight) / 13},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(screenHeight) / 6},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 5},
		},
		frameDelay: 56,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 4},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringEnemy, y: float64(screenHeight) / 6},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 15},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(5*screenHeight) / 15},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(9*screenHeight) / 15},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringEnemy, y: float64(screenHeight) / 2},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 8},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 6},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringUpEnemy, y: float64(14*screenHeight) / 15},
		},
		frameDelay: 56,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringDownEnemy, y: float64(screenHeight) / 16},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(screenHeight) / 2},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 8},
		},
		frameDelay: 34,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 11},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 11},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(5*screenHeight) / 15},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 15},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringEnemy, y: float64(screenHeight) / 4},
			enemySpawn{enemyType: staticFiringEnemy, y: float64(screenHeight) / 3},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringEnemy, y: float64(2*screenHeight) / 3},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringEnemy, y: float64(3*screenHeight) / 4},
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 5},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringEnemy, y: float64(2*screenHeight) / 5},
		},
		frameDelay: 90,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringEnemy, y: float64(screenHeight) / 5},
			enemySpawn{enemyType: staticFiringEnemy, y: float64(6*screenHeight) / 7},
		},
		frameDelay: 70,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringEnemy, y: float64(2*screenHeight) / 3},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringEnemy, y: float64(screenHeight) / 2},
			enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 3},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticFiringEnemy, y: float64(2*screenHeight) / 3},
		},
		frameDelay: 80,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: boss1, y: float64(screenHeight) / 2},
		},
		frameDelay: 600,
	},
}
