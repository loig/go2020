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

	if g.bulletSet.numBullets < 100 {
		g.bulletSet.addBullet(bullet{
			x: 799, y: 300, vx: -5, vy: 0, ax: 0, ay: 0,
		})
	}

	g.bulletSet.update()
	g.player.checkCollisions(g.bulletSet.bullets)

	return nil
}
