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

	switch g.state {
	case gameWelcome:
		g.welcomeUpdate()
	case gameHelp:
	case gameInfo:
	case gameIntro:
	case gameInLevel1, gameInLevel2:
		g.bulletSet.update()
		g.enemySet.update(&(g.bulletSet), &(g.powerUpSet), &(g.player.points), g.level.bossBattle)
		g.bossSet.update(&(g.bulletSet), &(g.powerUpSet), &(g.player.points))
		g.powerUpSet.update()
		g.player.update(&(g.powerUpSet))
		g.level.update(&(g.enemySet), &(g.bossSet), &(g.powerUpSet))
		g.player.checkCollisions(g.bulletSet.bullets, g.enemySet.enemies, g.bossSet.bosses, g.powerUpSet.powerUps)
	case gameTransition:
	case gameFinished:
	case gameOver:
	}
	return nil
}
