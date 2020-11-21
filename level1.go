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

func initLevel() level {
	var l level
	img, _, err := ebitenutil.NewImageFromFile("assets/Montagnes-1.png")
	if err != nil {
		panic(err)
	}
	l.firstPlan = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Montagnes-2.png")
	if err != nil {
		panic(err)
	}
	l.secondPlan = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Montagnes-3.png")
	if err != nil {
		panic(err)
	}
	l.thirdPlan = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Lune.png")
	if err != nil {
		panic(err)
	}
	l.fourthPlan = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Etoiles.png")
	if err != nil {
		panic(err)
	}
	l.background = img

	l.spawnSequence = level1SpawnSequence

	return l
}

var level1SpawnSequence []spawn = []spawn{
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: midBoss1, y: float64(screenHeight) / 2},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 4},
		},
		frameDelay: 300,
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
}
