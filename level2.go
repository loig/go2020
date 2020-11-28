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

import "github.com/hajimehoshi/ebiten/v2/ebitenutil"

func (g *game) setUpLevel2() {

	loadLevel2Enemies()
	loadLevel2Background()

	var l level
	l.firstPlanHeight = 396
	l.secondPlanHeight = 342
	l.thirdPlanHeight = 404
	l.spawnSequence = level2SpawnSequence

	g.level = l
	g.bulletSet = initBulletSet()
	g.player.initialPosition()
	g.enemySet = initEnemySet()
	g.bossSet = initBossSet()
	g.powerUpSet = initPowerUpSet()
}

func loadLevel2Enemies() {
	img, _, err := ebitenutil.NewImageFromFile("assets/Ennemi4.png")
	if err != nil {
		panic(err)
	}
	movingFiringEnemyImage = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Ennemi4.2.png")
	if err != nil {
		panic(err)
	}
	movingFiringEnemyImage2 = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Ennemi4.3.png")
	if err != nil {
		panic(err)
	}
	movingFiringEnemyImage3 = img

}

func disposeLevel2Enemies() {
	movingFiringEnemyImage.Dispose()
	movingFiringEnemyImage2.Dispose()
	movingFiringEnemyImage3.Dispose()
}

func loadLevel2Background() {
	img, _, err := ebitenutil.NewImageFromFile("assets/Level2-1.png")
	if err != nil {
		panic(err)
	}
	levelFirstPlan = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Level2-2.png")
	if err != nil {
		panic(err)
	}
	levelSecondPlan = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Level2-3.png")
	if err != nil {
		panic(err)
	}
	levelThirdPlan = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Level2-4.png")
	if err != nil {
		panic(err)
	}
	levelFourthPlan = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Level2-back.png")
	if err != nil {
		panic(err)
	}
	levelBackground = img
}

var level2SpawnSequence []spawn = []spawn{
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(screenHeight) / 2},
		},
		frameDelay: 300,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(3*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(6*screenHeight) / 20},
		},
		frameDelay: 200,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(17*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(14*screenHeight) / 20},
		},
		frameDelay: 150,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
		},
		frameDelay: 200,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(8*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(12*screenHeight) / 20},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(2*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(18*screenHeight) / 20},
		},
		frameDelay: 150,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(7*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(13*screenHeight) / 20},
		},
		frameDelay: 150,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: midBoss1, y: float64(screenHeight) / 5},
			enemySpawn{enemyType: midBoss1, y: float64(4*screenHeight) / 5},
		},
		frameDelay: 450,
	},
}
