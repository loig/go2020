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

const (
	staticEnemy int = iota
	staticExplodingEnemy
	staticFiringEnemy
	staticFiringDownEnemy
	staticFiringUpEnemy
	midBoss1
)

func makeStaticEnemy(x, y float64) enemy {
	var xSize float64 = 50
	var ySize float64 = 50
	return enemy{
		x: x, y: y,
		vx: -firstPlanPxPerFrame, vy: 0,
		xSize: xSize, ySize: ySize,
		pv:           1,
		powerUpProba: 8,
		hullAt00: []point{
			point{-xSize / 2, -ySize / 2},
			point{-xSize / 2, ySize / 2},
			point{xSize / 2, ySize / 2},
			point{xSize / 2, -ySize / 2},
		},
		image: staticEnemyImage,
	}
}

func makeStaticExplodingEnemy(x, y float64) enemy {
	var xSize float64 = 50
	var ySize float64 = 50
	return enemy{
		x: x, y: y,
		vx: -firstPlanPxPerFrame, vy: 0,
		xSize: xSize, ySize: ySize,
		pv:           1,
		powerUpProba: 3,
		hullAt00: []point{
			point{-xSize / 2, -ySize / 2},
			point{-xSize / 2, ySize / 2},
			point{xSize / 2, ySize / 2},
			point{xSize / 2, -ySize / 2},
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
		image: staticExplodingEnemyImage,
	}
}

func makeStaticFiringEnemy(x, y float64) enemy {
	var xSize float64 = 50
	var ySize float64 = 50
	return enemy{
		x: x, y: y,
		vx: -firstPlanPxPerFrame, vy: 0,
		xSize: xSize, ySize: ySize,
		pv:           1,
		powerUpProba: 10,
		hullAt00: []point{
			point{-xSize / 2, -ySize / 2},
			point{-xSize / 2, ySize / 2},
			point{xSize / 2, ySize / 2},
			point{xSize / 2, -ySize / 2},
		},
		bulletSequence: []bulletShot{
			bulletShot{
				bullets: []bullet{
					bullet{vx: -10, vy: 0, ax: 0, ay: 0},
				},
				interval: 50,
			},
		},
		image: staticFiringEnemyImage,
	}
}

func makeStaticFiringUpEnemy(x, y float64) enemy {
	var xSize float64 = 50
	var ySize float64 = 50
	return enemy{
		x: x, y: y,
		vx: -firstPlanPxPerFrame, vy: 0,
		xSize: xSize, ySize: ySize,
		pv:           1,
		powerUpProba: 10,
		hullAt00: []point{
			point{-xSize / 2, -ySize / 2},
			point{-xSize / 2, ySize / 2},
			point{xSize / 2, ySize / 2},
			point{xSize / 2, -ySize / 2},
		},
		bulletSequence: []bulletShot{
			bulletShot{
				bullets: []bullet{
					bullet{vx: -firstPlanPxPerFrame, vy: -5, ax: 0, ay: 0},
				},
				interval: 70,
			},
		},
		image: staticFiringEnemyImage,
	}
}

func makeStaticFiringDownEnemy(x, y float64) enemy {
	var xSize float64 = 50
	var ySize float64 = 50
	return enemy{
		x: x, y: y,
		vx: -firstPlanPxPerFrame, vy: 0,
		xSize: xSize, ySize: ySize,
		pv:           1,
		powerUpProba: 10,
		hullAt00: []point{
			point{-xSize / 2, -ySize / 2},
			point{-xSize / 2, ySize / 2},
			point{xSize / 2, ySize / 2},
			point{xSize / 2, -ySize / 2},
		},
		bulletSequence: []bulletShot{
			bulletShot{
				bullets: []bullet{
					bullet{vx: -firstPlanPxPerFrame, vy: 5, ax: 0, ay: 0},
				},
				interval: 70,
			},
		},
		image: staticFiringEnemyImage,
	}
}
