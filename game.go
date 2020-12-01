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
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/loig/go2020/assets"
	"github.com/loig/go2020/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type game struct {
	state      int
	stateState int
	stateFrame int
	bulletSet  bulletSet
	player     player
	enemySet   enemySet
	bossSet    bossSet
	powerUpSet powerUpSet
	level      level
	explosions explosionSet
	audio      soundManager
}

const (
	gameWelcome int = iota
	gameHelp
	gameInfo
	gameIntro
	gameInLevel1
	gameTransition
	gameInLevel2
	gameFinished
	gameOver
)

func initGame() *game {

	loadFirstImages()
	loadDurableImages()

	// font
	tt, err := opentype.Parse(fonts.LiberationSansBold)
	if err != nil {
		panic(err)
	}

	const dpi = 72
	theFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}
	theBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}

	text.CacheGlyphs(theFont, "0123456789-abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	g := game{}
	g.initAudio()

	return &g
}

func disposeFirstImages() {
	titleScreenImage.Dispose()
	helpScreenImage.Dispose()
}

func loadFirstImages() {
	img, _, err := image.Decode(bytes.NewReader(assets.Titre))
	//img, _, err := ebitenutil.NewImageFromFile("assets/Titre.png")
	if err != nil {
		panic(err)
	}
	titleScreenImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Help))
	//img, _, err := ebitenutil.NewImageFromFile("assets/help.png")
	if err != nil {
		panic(err)
	}
	helpScreenImage = ebiten.NewImageFromImage(img)
}

func loadDurableImages() {

	img, _, err := image.Decode(bytes.NewReader(assets.Vaisseau))
	//img, _, err := ebitenutil.NewImageFromFile("assets/Vaisseau.png")
	if err != nil {
		panic(err)
	}
	playerImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Tir1))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Tir1.png")
	if err != nil {
		panic(err)
	}
	playerBulletImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Grostir))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Gros-tir.png")
	if err != nil {
		panic(err)
	}
	playerBigBulletImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Option))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Option.png")
	if err != nil {
		panic(err)
	}
	optionImage = ebiten.NewImageFromImage(img)

	/*
		img, _, err = ebitenutil.NewImageFromFile("assets/Laser1.png")
		if err != nil {
			panic(err)
		}
		laserImage1 = img
		img, _, err = ebitenutil.NewImageFromFile("assets/Laser2.png")
		if err != nil {
			panic(err)
		}
		laserImage2 = img
		img, _, err = ebitenutil.NewImageFromFile("assets/Laser3.png")
		if err != nil {
			panic(err)
		}
		laserImage3 = img
		laserImage = laserImage1
	*/

	img, _, err = image.Decode(bytes.NewReader(assets.Tir2))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Tir2.png")
	if err != nil {
		panic(err)
	}
	enemyBasicBullet = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Bonusaucun))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Bonus-aucun.png")
	if err != nil {
		panic(err)
	}
	noBonusImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Bonus1))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Bonus-1.png")
	if err != nil {
		panic(err)
	}
	firstBonusImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Bonus2))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Bonus-2.png")
	if err != nil {
		panic(err)
	}
	secondBonusImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Bonus3))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Bonus-3.png")
	if err != nil {
		panic(err)
	}
	thirdBonusImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Bonus4))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Bonus-4.png")
	if err != nil {
		panic(err)
	}
	fourthBonusImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Vie))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Vie.png")
	if err != nil {
		panic(err)
	}
	lifeImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Bonus))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Bonus.png")
	if err != nil {
		panic(err)
	}
	bonusImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Explosion1))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Explosion1.png")
	if err != nil {
		panic(err)
	}
	explosionImages[0] = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Explosion2))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Explosion2.png")
	if err != nil {
		panic(err)
	}
	explosionImages[1] = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Explosion3))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Explosion3.png")
	if err != nil {
		panic(err)
	}
	explosionImages[2] = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Explosion4))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Explosion4.png")
	if err != nil {
		panic(err)
	}
	explosionImages[3] = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Explosion5))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Explosion5.png")
	if err != nil {
		panic(err)
	}
	explosionImages[4] = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Explosion6))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Explosion6.png")
	if err != nil {
		panic(err)
	}
	explosionImages[5] = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.GrosseExplosion1))
	//img, _, err = ebitenutil.NewImageFromFile("assets/GrosseExplosion1.png")
	if err != nil {
		panic(err)
	}
	bigExplosionImages[0] = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.GrosseExplosion2))
	//	img, _, err = ebitenutil.NewImageFromFile("assets/GrosseExplosion2.png")
	if err != nil {
		panic(err)
	}
	bigExplosionImages[1] = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.GrosseExplosion3))
	//img, _, err = ebitenutil.NewImageFromFile("assets/GrosseExplosion3.png")
	if err != nil {
		panic(err)
	}
	bigExplosionImages[2] = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.GrosseExplosion4))
	//img, _, err = ebitenutil.NewImageFromFile("assets/GrosseExplosion4.png")
	if err != nil {
		panic(err)
	}
	bigExplosionImages[3] = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.GrosseExplosion5))
	//img, _, err = ebitenutil.NewImageFromFile("assets/GrosseExplosion5.png")
	if err != nil {
		panic(err)
	}
	bigExplosionImages[4] = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.GrosseExplosion6))
	//img, _, err = ebitenutil.NewImageFromFile("assets/GrosseExplosion6.png")
	if err != nil {
		panic(err)
	}
	bigExplosionImages[5] = ebiten.NewImageFromImage(img)
}
