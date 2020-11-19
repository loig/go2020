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

import "math/rand"

const (
	testEnemy int = iota
	staticEnemy
	staticExplodingEnemy
	staticFiringEnemy
	staticRotatingFireEnemy
)

func makeStaticEnemy(x, y float64) enemy {
	var xSize float64 = 42
	var ySize float64 = 42
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
	}
}

func makeStaticExplodingEnemy(x, y float64) enemy {
	var xSize float64 = 42
	var ySize float64 = 42
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
	}
}

func makeStaticFiringEnemy(x, y float64) enemy {
	var xSize float64 = 42
	var ySize float64 = 42
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
	}
}

func makeStaticRotatingFireEnemy(x, y float64) enemy {
	var xSize float64 = 42
	var ySize float64 = 42
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
					bullet{vx: -10},
				},
				interval: 20,
			},
			bulletShot{
				bullets: []bullet{
					bullet{vx: -7, vy: -7},
				},
				interval: 20,
			},
			bulletShot{
				bullets: []bullet{
					bullet{vy: -10},
				},
				interval: 20,
			},
			bulletShot{
				bullets: []bullet{
					bullet{vx: 7, vy: -7},
				},
				interval: 20,
			},
			bulletShot{
				bullets: []bullet{
					bullet{vx: 10},
				},
				interval: 20,
			},
			bulletShot{
				bullets: []bullet{
					bullet{vx: 7, vy: 7},
				},
				interval: 20,
			},
			bulletShot{
				bullets: []bullet{
					bullet{vy: 10},
				},
				interval: 20,
			},
			bulletShot{
				bullets: []bullet{
					bullet{vx: -7, vy: 7},
				},
				interval: 20,
			},
		},
	}
}

func makeTestEnemy() enemy {
	return enemy{
		points: 10,
		x:      screenWidth - 1, y: float64(rand.Intn(screenHeight-100) + 50),
		vx: -5, vy: 0,
		xSize: 25, ySize: 15,
		pv:           1,
		powerUpProba: 2,
		bulletSequence: []bulletShot{
			bulletShot{
				bullets: []bullet{
					bullet{vx: -10, vy: 0, ax: 0, ay: 0},
				},
				interval: 30,
			},
			/*
				bulletShot{
					bullets: []bullet{
						bullet{vx: -10, vy: 5, ax: 0, ay: 0},
						bullet{vx: -10, vy: -5, ax: 0, ay: 0},
					},
					interval: 5,
				},
			*/
		},
		/*
			accelerationSequence: []acceleration{
				acceleration{ax: 0, ay: 1, interval: 5},
				acceleration{ax: 0, ay: 0, interval: 10},
				acceleration{ax: 0, ay: -1, interval: 10},
				acceleration{ax: 0, ay: 0, interval: 10},
				acceleration{ax: 0, ay: 1, interval: 5},
			},
		*/
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
		hullAt00: []point{
			point{-25 / 2, 15 / 2},
			point{25 / 2, 15 / 2},
			point{25 / 2, -15 / 2},
		},
	}
}
