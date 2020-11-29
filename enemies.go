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
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	staticEnemy int = iota
	staticExplodingEnemy
	staticFiringEnemy
	staticFiringDownEnemy
	staticFiringUpEnemy
	movingFiringEnemy
	midBoss1
	boss1
	boss2
)

const (
	staticEnemyPoints          = 250
	staticExplodingEnemyPoints = 500
	staticFiringEnemyPoints    = 500
	movingFiringEnemyPoints    = 750
)

const (
	movingFiringEnemySpeed = -2
	firingEnemyBulletSpeed = -8
	staticEnemyBulletSpeed = 3
)

func makeStaticEnemy(x, y float64) enemy {
	var xSize float64 = 50
	var ySize float64 = 50
	xSpeed := float64(rand.Intn(3)-1) / 4
	ySpeed := float64(rand.Intn(3)-1) / 4
	return enemy{
		x: x + xSize/2, y: y,
		xMin: x, xMax: x + xSize,
		yMin: y - ySize/2, yMax: y + ySize/2,
		vx: -firstPlanPxPerFrame + xSpeed, vy: ySpeed,
		xSize: xSize, ySize: ySize,
		pv:           1,
		powerUpProba: 15,
		hullAt00: []point{
			point{-xSize/2 + 7, -ySize/2 + 7},
			point{-xSize/2 + 7, ySize / 2},
			point{xSize/2 - 2, ySize / 2},
			point{xSize/2 - 2, -ySize/2 + 7},
		},
		image:  staticEnemyImage,
		points: staticEnemyPoints,
	}
}

func makeStaticExplodingEnemy(x, y float64) enemy {
	var xSize float64 = 50
	var ySize float64 = 50
	xSpeed := float64(rand.Intn(3)-1) / 4
	ySpeed := float64(rand.Intn(3)-1) / 4
	return enemy{
		x: x + xSize/2, y: y,
		xMin: x, xMax: x + xSize,
		yMin: y - ySize/2, yMax: y + ySize/2,
		vx: -firstPlanPxPerFrame + xSpeed, vy: ySpeed,
		xSize: xSize, ySize: ySize,
		pv:           1,
		powerUpProba: 8,
		hullAt00: []point{
			point{-xSize/2 + 7, -ySize/2 + 7},
			point{-xSize/2 + 7, ySize / 2},
			point{xSize/2 - 2, ySize / 2},
			point{xSize/2 - 2, -ySize/2 + 7},
		},
		deathBlow: []bullet{
			bullet{vx: -10},
			bullet{vx: 10},
			bullet{vy: -10},
			bullet{vy: 10},
			bullet{vx: 7, vy: 7},
			bullet{vx: 7, vy: -7},
			bullet{vx: -7, vy: 7},
			bullet{vx: -7, vy: -7},
		},
		image:  staticExplodingEnemyImage,
		points: staticExplodingEnemyPoints,
	}
}

func makeStaticFiringEnemy(x, y float64) enemy {
	var xSize float64 = 134
	var ySize float64 = 128
	return enemy{
		x: x + xSize/2, y: y,
		xMin: x, xMax: x + xSize,
		yMin: y - ySize/2, yMax: y + ySize/2,
		vx: -firstPlanPxPerFrame, vy: 0,
		xSize: xSize, ySize: ySize,
		pv:           1,
		powerUpProba: 10,
		hullAt00: []point{
			point{-xSize / 4, -ySize/3 - 3},
			point{-xSize / 4, ySize/3 + 5},
			point{xSize / 2, ySize / 7},
			point{xSize / 2, -ySize / 7},
		},
		bulletSequence: []bulletShot{
			bulletShot{
				bullets: []bullet{
					bullet{vx: -firstPlanPxPerFrame - staticEnemyBulletSpeed},
					bullet{vx: -firstPlanPxPerFrame - staticEnemyBulletSpeed, vy: 1.5},
					bullet{vx: -firstPlanPxPerFrame - staticEnemyBulletSpeed, vy: -1.5},
				},
				interval: 10,
			},
			bulletShot{
				interval: 60,
			},
		},
		image:  staticFiringEnemyImage,
		points: staticFiringEnemyPoints,
	}
}

func makeStaticFiringUpEnemy(x, y float64) enemy {
	var xSize float64 = 128
	var ySize float64 = 134
	return enemy{
		x: x + xSize/2, y: y,
		xMin: x, xMax: x,
		yMin: y - ySize/2, yMax: y + ySize/2,
		vx: -firstPlanPxPerFrame, vy: 0,
		xSize: xSize, ySize: ySize,
		pv:           1,
		powerUpProba: 10,
		hullAt00: []point{
			point{-xSize/3 - 5, -ySize / 4},
			point{-xSize / 7, ySize / 2},
			point{xSize / 7, ySize / 2},
			point{xSize/3 + 3, -ySize / 4},
		},
		bulletSequence: []bulletShot{
			bulletShot{
				bullets: []bullet{
					bullet{vx: -firstPlanPxPerFrame, vy: -staticEnemyBulletSpeed, ax: 0, ay: 0},
					bullet{vx: -firstPlanPxPerFrame - 1.5, vy: -staticEnemyBulletSpeed, ax: 0, ay: 0},
					bullet{vx: -firstPlanPxPerFrame + 1.5, vy: -staticEnemyBulletSpeed, ax: 0, ay: 0},
				},
				interval: 10,
			},
			bulletShot{
				interval: 60,
			},
		},
		image:  staticFiringUpEnemyImage,
		points: staticFiringEnemyPoints,
	}
}

func makeStaticFiringDownEnemy(x, y float64) enemy {
	var xSize float64 = 128
	var ySize float64 = 134
	return enemy{
		x: x + xSize/2, y: y,
		xMin: x, xMax: x + xSize,
		yMin: y - ySize/2, yMax: y + ySize/2,
		vx: -firstPlanPxPerFrame, vy: 0,
		xSize: xSize, ySize: ySize,
		pv:           1,
		powerUpProba: 10,
		hullAt00: []point{
			point{-xSize / 7, -ySize / 2},
			point{-xSize/3 - 3, ySize / 4},
			point{xSize/3 + 5, ySize / 4},
			point{xSize / 7, -ySize / 2},
		},
		bulletSequence: []bulletShot{
			bulletShot{
				bullets: []bullet{
					bullet{vx: -firstPlanPxPerFrame, vy: staticEnemyBulletSpeed, ax: 0, ay: 0},
					bullet{vx: -firstPlanPxPerFrame + 1.5, vy: staticEnemyBulletSpeed, ax: 0, ay: 0},
					bullet{vx: -firstPlanPxPerFrame - 1.5, vy: staticEnemyBulletSpeed, ax: 0, ay: 0},
				},
				interval: 10,
			},
			bulletShot{
				bullets:  []bullet{},
				interval: 60,
			},
		},
		image:  staticFiringDownEnemyImage,
		points: staticFiringEnemyPoints,
	}
}

func makeMovingFiringEnemy(x, y float64) enemy {
	var xSize float64 = 152
	var ySize float64 = 120
	halfXSize := xSize / 2
	halfYSize := ySize / 2
	return enemy{
		x: x + halfXSize, y: y,
		xMin: x, xMax: x + xSize,
		yMin: y - halfYSize, yMax: y + halfYSize,
		vx: -firstPlanPxPerFrame + movingFiringEnemySpeed, vy: 0,
		xSize: xSize, ySize: ySize,
		pv:           1,
		powerUpProba: 10,
		hullAt00: []point{
			point{-halfXSize + 7, -halfYSize + 5},
			point{-halfXSize + 7, halfYSize - 5},
			point{halfXSize - 7, halfYSize - 5},
			point{halfXSize - 7, -halfYSize + 5},
		},
		bulletSequence: []bulletShot{
			bulletShot{
				bullets: []bullet{
					bullet{x: -halfXSize + 10, y: 40, vx: -firstPlanPxPerFrame + firingEnemyBulletSpeed},
					bullet{x: -halfXSize + 10, y: 15, vx: -firstPlanPxPerFrame + firingEnemyBulletSpeed - 2},
					bullet{x: -halfXSize + 10, y: -15, vx: -firstPlanPxPerFrame + firingEnemyBulletSpeed - 2},
					bullet{x: -halfXSize + 10, y: -40, vx: -firstPlanPxPerFrame + firingEnemyBulletSpeed},
				},
				interval: 10,
			},
			bulletShot{
				bullets: []bullet{
					bullet{x: -halfXSize + 10, y: 40, vx: -firstPlanPxPerFrame + firingEnemyBulletSpeed, vy: 2},
					bullet{x: -halfXSize + 10, y: 15, vx: -firstPlanPxPerFrame + firingEnemyBulletSpeed - 1},
					bullet{x: -halfXSize + 10, y: -15, vx: -firstPlanPxPerFrame + firingEnemyBulletSpeed - 1},
					bullet{x: -halfXSize + 10, y: -40, vx: -firstPlanPxPerFrame + firingEnemyBulletSpeed, vy: -2},
				},
				interval: 10,
			},
			bulletShot{
				bullets: []bullet{
					bullet{x: -halfXSize + 10, y: 40, vx: -firstPlanPxPerFrame + firingEnemyBulletSpeed, vy: 5},
					bullet{x: -halfXSize + 10, y: 15, vx: -firstPlanPxPerFrame + firingEnemyBulletSpeed},
					bullet{x: -halfXSize + 10, y: -15, vx: -firstPlanPxPerFrame + firingEnemyBulletSpeed},
					bullet{x: -halfXSize + 10, y: -40, vx: -firstPlanPxPerFrame + firingEnemyBulletSpeed, vy: -5},
				},
				interval: 10,
			},
			bulletShot{
				bullets:  []bullet{},
				interval: 70,
			},
		},
		image:         movingFiringEnemyImage,
		isAnimated:    true,
		framePerImage: 7,
		images:        []*ebiten.Image{movingFiringEnemyImage, movingFiringEnemyImage2, movingFiringEnemyImage3},
	}
}
