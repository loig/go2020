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
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type bullet struct {
	x  float64
	y  float64
	vx float64
	vy float64
	ax float64
	ay float64
}

func (b *bullet) update() {
	b.x += b.vx
	b.y += b.vy
	b.vx += b.ax
	b.vy += b.ay
}

func (b bullet) draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, b.x-2, b.y-2, 4, 4, color.RGBA{255, 0, 0, 255})
}

func (b bullet) isOut() bool {
	return b.x < 0 || b.y < 0 || b.x >= screenWidth || b.y >= screenHeight
}

type bulletSet struct {
	numBullets int
	bullets    []*bullet
}

func initBulletSet() bulletSet {
	return bulletSet{
		numBullets: 0,
		bullets:    make([]*bullet, 0),
	}
}

func (bs *bulletSet) addBullet(b bullet) {
	bs.numBullets++
	bs.bullets = append(bs.bullets, &b)
}

func (bs *bulletSet) update() {
	for pos := 0; pos < bs.numBullets; pos++ {
		bs.bullets[pos].update()
		if bs.bullets[pos].isOut() {
			bs.numBullets--
			bs.bullets[pos] = bs.bullets[bs.numBullets]
			bs.bullets = bs.bullets[:bs.numBullets]
		}
	}
}

func (bs *bulletSet) draw(screen *ebiten.Image) {
	for _, b := range bs.bullets {
		b.draw(screen)
	}
}
