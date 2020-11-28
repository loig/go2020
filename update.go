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

	g.updateMusic()

	switch g.state {
	case gameWelcome:
		g.welcomeUpdate()
	case gameHelp:
		g.helpUpdate()
	case gameInfo:
		g.infoUpdate()
	case gameIntro:
		g.introUpdate()
	case gameInLevel1, gameInLevel2:
		g.bulletSet.update()
		g.enemySetUpdate()
		g.bossSetUpdate()
		g.powerUpSet.update()
		g.playerUpdate()
		g.levelUpdate()
		g.player.checkCollisions(g.bulletSet.bullets, g.enemySet.enemies, g.bossSet.bosses, g.powerUpSet.powerUps)
	case gameTransition:
		g.transitionUpdate()
	case gameFinished:
		g.finishedUpdate()
	case gameOver:
		g.gameOverUpdate()
	}
	return nil
}
