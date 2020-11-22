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

type game struct {
	bulletSet  bulletSet
	player     player
	enemySet   enemySet
	bossSet    bossSet
	powerUpSet powerUpSet
	level      level
}

func initGame() *game {

	img, _, err := ebitenutil.NewImageFromFile("assets/Tir2.png")
	if err != nil {
		panic(err)
	}
	enemyBasicBullet = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Ennemi1.png")
	if err != nil {
		panic(err)
	}
	staticEnemyImage = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Ennemi2.png")
	if err != nil {
		panic(err)
	}
	staticFiringEnemyImage = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Ennemi3.png")
	if err != nil {
		panic(err)
	}
	staticExplodingEnemyImage = img

	return &game{
		bulletSet:  initBulletSet(),
		player:     initPlayer(),
		enemySet:   initEnemySet(),
		bossSet:    initBossSet(),
		powerUpSet: initPowerUpSet(),
		level:      initLevel(),
	}
}
