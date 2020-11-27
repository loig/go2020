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
	"image"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type player struct {
	image            *ebiten.Image
	fireImage        *ebiten.Image
	bigFireImage     *ebiten.Image
	laserImage       *ebiten.Image
	x                float64
	y                float64
	vx               float64
	vy               float64
	xSize            float64
	ySize            float64
	bullets          bulletSet
	lastBullet       int
	currentFire      int
	numShot          int // for basic shot only
	shotDelay        int // for big shot only
	shotWidth        int // for laser shot only
	laserOn          bool
	laser            bullet
	laserLevel       int
	vCap             float64
	numOptions       int
	currentPosition  int
	positionHistory  [pMoveRecorded]playerPosition
	options          [pMaxOption]option
	currentPowerUp   int
	usedPowerUp      int
	allPowerUp       bool
	hullSet          bool
	cHull            []point
	xMin             float64
	yMin             float64
	xMax             float64
	yMax             float64
	collision        bool
	invincibleFrames int
	points           int
	lastLiveWon      int
	lives            int
}

type playerPosition struct {
	x float64
	y float64
}

const (
	pInitX               = 100
	pInitY               = screenHeight / 2
	pWidth               = 100
	pHeight              = 62
	pMaxVCap             = 10
	pVStep               = 2
	pVInit               = 6
	pAx                  = 1
	pAy                  = 1
	pBulletInterval      = 15
	pBulletSpeed         = 12
	pMaxShot             = 5  // for basic shot only
	pMinDelay            = 20 // for big shot only
	pMaxDelay            = 80
	pShotDelayStep       = 20
	pMaxShotWidth        = 18 // for laser shot only
	pMinShotWidth        = 6
	pShotWidthIncrease   = 6
	pMaxOption           = 3
	pDifferentPowerUps   = 4
	pMoveRecorded        = 32
	pFrameBetweenOptions = 10
	pInvicibleDuration   = 120
	pInitLives           = 3
	pMaxLives            = 7
	pPointsForLive       = 25000
	pPointsPerPowerUp    = 500
	laserImageWidth      = 138
	laserImageHeight     = 30
	laserImageOffset     = 30
)

var pOtherBulletSpeed [5]float64 = [5]float64{0, 1, -1, 2, -2}
var optionImage *ebiten.Image
var laserImage1 *ebiten.Image
var laserImage2 *ebiten.Image
var laserImage3 *ebiten.Image

func initPlayer() player {
	var p player = player{
		x: pInitX, y: pInitY,
		xSize: pWidth, ySize: pHeight,
		bullets:         initBulletSet(),
		lastBullet:      pBulletInterval,
		numShot:         1,
		shotDelay:       pMaxDelay,
		shotWidth:       pMinShotWidth,
		vCap:            pVInit,
		positionHistory: makePositionHistory(pInitX, pInitY),
		lives:           pInitLives,
	}
	img, _, err := ebitenutil.NewImageFromFile("assets/Vaisseau.png")
	if err != nil {
		panic(err)
	}
	p.image = img
	img, _, err = ebitenutil.NewImageFromFile("assets/Tir1.png")
	if err != nil {
		panic(err)
	}
	p.fireImage = img
	img, _, err = ebitenutil.NewImageFromFile("assets/Gros-tir.png")
	if err != nil {
		panic(err)
	}
	p.bigFireImage = img
	img, _, err = ebitenutil.NewImageFromFile("assets/Option.png")
	if err != nil {
		panic(err)
	}
	optionImage = img
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
	p.laserImage = laserImage1
	return p
}

func (p *player) reset() {
	p.x = pInitX
	p.y = pInitY
	p.collision = false
	p.invincibleFrames = pInvicibleDuration
	p.lastBullet = 0
	p.currentFire = 0
	p.numShot = 1
	p.shotDelay = pMaxDelay
	p.shotWidth = pMinShotWidth
	p.vCap = pVInit
	p.numOptions = 0
	p.currentPosition = 0
	p.positionHistory = makePositionHistory(pInitX, pInitY)
	p.currentPowerUp = 0
	p.allPowerUp = false
	p.usedPowerUp = 0
	p.laserImage = laserImage1
	p.laserLevel = 0
}

func (p player) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.xmin(), p.ymin())
	screen.DrawImage(
		p.image,
		op,
	)
	cHull := p.convexHull()
	hullColor := color.RGBA{0, 255, 0, 255}
	if p.invincibleFrames > 0 {
		hullColor = color.RGBA{0, 255, 255, 255}
	}
	for i := 0; i < len(cHull); i++ {
		ii := (i + 1) % len(cHull)
		ebitenutil.DrawLine(screen, cHull[i].x, cHull[i].y, cHull[ii].x, cHull[ii].y, hullColor)
	}
	for oPos := 0; oPos < p.numOptions; oPos++ {
		p.options[oPos].draw(screen)
	}
	if p.laserOn {
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(p.laser.xmin(), p.laser.y-laserImageHeight/2)
		screen.DrawImage(
			p.laserImage,
			op,
		)
		currentMaxX := p.laser.xmin() + laserImageWidth
		for currentMaxX < screenWidth {
			op2 := &ebiten.DrawImageOptions{}
			op2.GeoM.Translate(currentMaxX, p.laser.y-laserImageHeight/2)
			screen.DrawImage(
				p.laserImage.SubImage(image.Rect(laserImageOffset, 0, laserImageWidth, laserImageHeight)).(*ebiten.Image),
				op2,
			)
			currentMaxX += laserImageWidth - laserImageOffset
		}
		p.laser.draw(screen, color.RGBA{0, 255, 0, 255})
	}
	p.bullets.draw(screen, color.RGBA{0, 255, 0, 255})
}

func (p *player) updateBox() {
	p.xMin = p.x - p.xSize/2
	p.xMax = p.x + p.xSize/2
	p.yMin = p.y - p.ySize/2
	p.yMax = p.y + p.ySize/2
}

func (p *player) xmin() float64 {
	return p.xMin
}

func (p *player) xmax() float64 {
	return p.xMax
}

func (p *player) ymin() float64 {
	return p.yMin
}

func (p *player) ymax() float64 {
	return p.yMax
}

func (p *player) convexHull() []point {
	if !p.hullSet {
		p.cHull = []point{
			point{p.x - p.xSize/3, p.y + p.ySize/4},
			point{p.x + p.xSize/2, p.y + p.ySize/4},
			point{p.x - p.xSize/2, p.y - p.ySize/3},
		}
		p.hullSet = true
	}
	return p.cHull
}

func (p *player) hasCollided() {
	p.collision = true
}

func (p *player) checkCollisions(bs []*bullet, es []*enemy, bbs []*boss, ps []*powerUp) {
	for oNum := 0; oNum < p.numOptions; oNum++ {
		o := p.options[oNum]
		for _, b := range bs {
			collide(o, b)
		}
		for _, e := range es {
			collide(o, e)
		}
		for _, b := range bbs {
			for pos := 0; pos < len(b.hitBoxes); pos++ {
				collide(o, &(b.hitBoxes[pos]))
			}
		}
	}
	if p.invincibleFrames <= 0 {
		for _, b := range bs {
			collide(p, b)
		}
		for _, e := range es {
			collide(p, e)
		}
		for _, b := range bbs {
			for pos := 0; pos < len(b.hitBoxes); pos++ {
				collide(p, &(b.hitBoxes[pos]))
			}
			for pos := 0; pos < len(b.hurtBoxes); pos++ {
				collide(p, &(b.hurtBoxes[pos]))
			}
		}
	}
	if p.laserOn {
		for _, e := range es {
			collide(&(p.laser), e)
		}
		for _, b := range bbs {
			for pos := 0; pos < len(b.hitBoxes); pos++ {
				collide(&(p.laser), &(b.hitBoxes[pos]))
			}
		}
	}
	for _, b := range p.bullets.bullets {
		for _, e := range es {
			collide(b, e)
		}
		for _, bb := range bbs {
			for pos := 0; pos < len(bb.hitBoxes); pos++ {
				collide(b, &(bb.hitBoxes[pos]))
			}
		}
		if b.isBig {
			for _, bb := range bs {
				collideNoHarm(b, bb)
			}
		}
	}
	for _, pu := range ps {
		if collideNoHarm(p, pu) {
			p.getPowerUp()
		}
	}
}

func (p *player) update(ps *powerUpSet) {
	if p.collision {
		p.lives--
		p.releasePowerUps(ps)
		p.reset()
	} else {
		if p.invincibleFrames > 0 {
			p.invincibleFrames--
		}
	}
	p.hullSet = false
	p.cHull = nil
	p.laserOn = false
	p.move()
	p.managePowerUp()
	p.fire()
	p.moveOptions()
	p.bullets.update()
	p.updateBox()
	p.checkLiveWin()
}

func (p *player) move() {
	var hasMovedX bool
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.vx += pAx
		hasMovedX = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if hasMovedX {
			p.vx = 0
			hasMovedX = false
		} else {
			p.vx -= pAx
			hasMovedX = true
		}
	}
	if !hasMovedX {
		if p.vx > pAx {
			p.vx -= pAx
		} else if p.vx < -pAx {
			p.vx += pAx
		} else {
			p.vx = 0
		}
	}
	var hasMovedY bool
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.vy += pAy
		hasMovedY = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if hasMovedY {
			p.vy = 0
			hasMovedY = false
		} else {
			p.vy -= pAy
			hasMovedY = true
		}
	}
	if !hasMovedY {
		if p.vy > pAy {
			p.vy -= pAy
		} else if p.vy < -pAy {
			p.vy += pAy
		} else {
			p.vy = 0
		}
	}
	if p.vx > p.vCap {
		p.vx = p.vCap
	} else if p.vx < -p.vCap {
		p.vx = -p.vCap
	}
	if p.vy > p.vCap {
		p.vy = p.vCap
	} else if p.vy < -p.vCap {
		p.vy = -p.vCap
	}
	p.x += p.vx
	if p.x < 0+p.xSize/2 {
		p.x = 0 + p.xSize/2
	} else if p.x > screenWidth-p.xSize/2 {
		p.x = screenWidth - p.xSize/2
	}
	p.y += p.vy
	if p.y < 0+p.ySize/2 {
		p.y = 0 + p.ySize/2
	} else if p.y > screenHeight-p.ySize/2 {
		p.y = screenHeight - p.ySize/2
	}
	if hasMovedX || hasMovedY {
		p.recordMove()
	}
}

func (p *player) recordMove() {
	p.currentPosition = (p.currentPosition + 1) % pMoveRecorded
	p.positionHistory[p.currentPosition] = playerPosition{x: p.x, y: p.y}
}

func (p *player) moveOptions() {
	posPos := (p.currentPosition + pMoveRecorded - pFrameBetweenOptions) % pMoveRecorded
	for oPos := 0; oPos < p.numOptions; oPos++ {
		p.options[oPos].x = p.positionHistory[posPos].x
		p.options[oPos].y = p.positionHistory[posPos].y
		posPos = (posPos + pMoveRecorded - pFrameBetweenOptions) % pMoveRecorded
	}
}

func (p *player) fire() {
	if p.currentFire == 2 && ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.laserOn = true
		xLen := screenWidth - p.x
		x := xLen/2 + p.x
		p.laser = bullet{
			x: x, y: p.y,
			xSize: xLen, ySize: float64(p.shotWidth),
			xMin: p.x, xMax: p.x + xLen,
			yMin: p.y - float64(p.shotWidth)/2, yMax: p.y + float64(p.shotWidth)/2,
		}
	}
	p.lastBullet++
	bulletInterval := pBulletInterval
	if p.currentFire == 1 {
		bulletInterval = p.shotDelay
	}
	if p.lastBullet >= bulletInterval &&
		ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.lastBullet = 0
		if p.currentFire == 0 {
			for bNum := 0; bNum < p.numShot; bNum++ {
				p.bullets.addBullet(bullet{
					x: p.x, y: p.y,
					vx: pBulletSpeed, vy: pOtherBulletSpeed[bNum],
					ax: 0, ay: 0,
					image: p.fireImage,
				})
			}
		} else if p.currentFire == 1 {
			p.bullets.addBigBullet(bullet{
				x: p.x, y: p.y,
				vx: pBulletSpeed, vy: 0,
				ax: 0, ay: 0,
				image: p.bigFireImage,
			})
		}
		for oNum := 0; oNum < p.numOptions; oNum++ {
			p.bullets.addBullet(bullet{
				x: p.options[oNum].x, y: p.options[oNum].y,
				vx: pBulletSpeed, vy: 0,
				ax: 0, ay: 0,
				image: p.fireImage,
			})
		}
	}
}

func (p *player) getPowerUp() {
	p.points += pPointsPerPowerUp
	if !p.allPowerUp {
		start := p.currentPowerUp
		p.currentPowerUp = (p.currentPowerUp + 1) % (pDifferentPowerUps + 1)
		for !p.isAppliablePowerUp(p.currentPowerUp) && p.currentPowerUp != start {
			p.currentPowerUp = (p.currentPowerUp + 1) % (pDifferentPowerUps + 1)
		}
		if p.currentPowerUp == start && !p.isAppliablePowerUp(p.currentPowerUp) {
			p.allPowerUp = true
		}
	}
}

func (p player) isAppliablePowerUp(powerUp int) bool {
	switch powerUp {
	case 0:
		return false
	case 1:
		return p.vCap < pMaxVCap
	case 2:
		switch p.currentFire {
		case 0: // normal shot
			return p.numShot < pMaxShot
		case 1: // big shot
			return p.shotDelay > pMinDelay
		case 2: // laser
			return p.shotWidth < pMaxShotWidth
		}
	case 3:
		return true
	case 4:
		return p.numOptions < pMaxOption
	}
	return false
}

func (p *player) applyPowerUp() {
	switch p.currentPowerUp {
	case 1:
		p.vCap += pVStep
	case 2:
		switch p.currentFire {
		case 0: // normal shot
			p.numShot++
		case 1: // big shot
			p.shotDelay -= pShotDelayStep
		case 2: // laser
			p.shotWidth += pShotWidthIncrease
			p.laserLevel++
			switch p.laserLevel {
			case 0:
				p.laserImage = laserImage1
			case 1:
				p.laserImage = laserImage2
			case 2:
				p.laserImage = laserImage3
			}
		}
	case 3:
		p.currentFire++
		p.currentFire = p.currentFire % 3
	case 4:
		p.numOptions++
	}
	p.usedPowerUp++
	p.currentPowerUp = 0
}

func (p *player) managePowerUp() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if !p.allPowerUp && p.isAppliablePowerUp(p.currentPowerUp) {
			p.applyPowerUp()
		}
	}
}

func makePositionHistory(x, y float64) [pMoveRecorded]playerPosition {
	var moves [pMoveRecorded]playerPosition
	for i := 0; i < pMoveRecorded; i++ {
		moves[i] = playerPosition{x: x, y: y}
	}
	return moves
}

func (p *player) checkLiveWin() {
	if p.points >= p.lastLiveWon+pPointsForLive {
		if p.lives < pMaxLives {
			p.lives++
		}
		p.lastLiveWon += pPointsForLive
	}
}

func (p *player) releasePowerUps(ps *powerUpSet) {
	numToLaunch := p.usedPowerUp / 2
	for i := 0; i < numToLaunch; i++ {
		xShift := rand.Intn(3) - 1
		yShift := rand.Intn(3) - 1
		ps.addPowerUp(powerUp{
			x: p.x + float64(10*i), y: p.y + float64(10*i),
			vx: -firstPlanPxPerFrame + float64(xShift), vy: float64(yShift),
		})
	}
}
