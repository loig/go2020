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

func (g *game) Update() error {

	if g.enemySet.numEnemies < 1 {
		g.enemySet.addTestEnemy()
	}

	if g.powerUpSet.numPowerUps < 1 {
		g.powerUpSet.addPowerUp(powerUp{
			x: screenWidth / 2, y: screenHeight / 2, vx: -3, vy: 0,
		})
	}

	g.bulletSet.update()
	g.enemySet.update(&(g.bulletSet), &(g.powerUpSet), &(g.player.points))
	g.powerUpSet.update()
	g.player.update()
	g.level.update()
	g.player.checkCollisions(g.bulletSet.bullets, g.enemySet.enemies, g.powerUpSet.powerUps)

	return nil
}
