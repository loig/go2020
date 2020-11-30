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
	isDead           bool
	deadNumFrames    int
	deadFrame        int
}

type playerPosition struct {
	x float64
	y float64
}

const (
	pInitX                = 100
	pInitY                = screenHeight / 2
	pWidth                = 100
	pHeight               = 62
	pMaxVCap              = 10
	pVStep                = 2
	pVInit                = 6
	pAx                   = 1
	pAy                   = 1
	pBulletInterval       = 15
	pBulletSpeed          = 12
	pMaxShot              = 5  // for basic shot only
	pMinDelay             = 20 // for big shot only
	pMaxDelay             = 80
	pShotDelayStep        = 20
	pMaxShotWidth         = 18 // for laser shot only
	pMinShotWidth         = 6
	pShotWidthIncrease    = 6
	pMaxOption            = 3
	pDifferentPowerUps    = 4
	pMoveRecorded         = 32
	pFrameBetweenOptions  = 10
	pInvicibleDuration    = 150
	pInitLives            = 3
	pMaxLives             = 5
	pPointsForLive        = 50000
	pPointsPerPowerUp     = 500
	laserImageWidth       = 138
	laserImageHeight      = 30
	laserImageOffset      = 30
	playerXBulletShift    = -20
	playerYBulletShift    = -4
	playerXBigBulletShift = -35
	pDeadNumFrames        = 90
)

var pOtherBulletSpeed [5]float64 = [5]float64{0, 1, -1, 2, -2}

func initPlayer() player {
	return player{
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
		deadNumFrames:   pDeadNumFrames,
	}
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
	p.laserLevel = 0
	p.deadFrame = 0
	p.isDead = false
	laserImage = laserImage1
}

func (p *player) initialPosition() {
	p.x = pInitX
	p.y = pInitY
	p.bullets = initBulletSet()
	p.collision = false
	p.lastBullet = 0
	p.currentPosition = 0
	p.positionHistory = makePositionHistory(pInitX, pInitY)
}

func (p player) draw(screen *ebiten.Image) {
	if p.laserOn {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(p.laser.xmin(), p.laser.y-laserImageHeight/2)
		screen.DrawImage(
			laserImage,
			op,
		)
		currentMaxX := p.laser.xmin() + laserImageWidth
		for currentMaxX < screenWidth {
			op2 := &ebiten.DrawImageOptions{}
			op2.GeoM.Translate(currentMaxX, p.laser.y-laserImageHeight/2)
			screen.DrawImage(
				laserImage.SubImage(image.Rect(laserImageOffset, 0, laserImageWidth, laserImageHeight)).(*ebiten.Image),
				op2,
			)
			currentMaxX += laserImageWidth - laserImageOffset
		}
		if isDebug() {
			p.laser.draw(screen, color.RGBA{255, 0, 0, 255})
		}
	}
	p.bullets.draw(screen, color.RGBA{255, 0, 0, 255})
	if !p.isDead {
		op := &ebiten.DrawImageOptions{}
		if p.invincibleFrames <= 0 || (p.invincibleFrames/7)%2 == 0 {
			op.GeoM.Translate(p.xmin(), p.ymin()+8)
			screen.DrawImage(
				playerImage,
				op,
			)
		}
		if isDebug() {
			cHull := p.convexHull()
			hullColor := color.RGBA{255, 0, 0, 255}
			if p.invincibleFrames > 0 {
				hullColor = color.RGBA{255, 255, 0, 255}
			}
			for i := 0; i < len(cHull); i++ {
				ii := (i + 1) % len(cHull)
				ebitenutil.DrawLine(screen, cHull[i].x, cHull[i].y, cHull[ii].x, cHull[ii].y, hullColor)
			}
		}
	}
	for oPos := 0; oPos < p.numOptions; oPos++ {
		p.options[oPos].draw(screen)
	}
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
			point{p.x - p.xSize/3, p.y + p.ySize/4 - 2},
			point{p.x + p.xSize/2 - 5, p.y + p.ySize/4 - 4},
			point{p.x - p.xSize/2 + 6, p.y - p.ySize/3 + 3},
		}
		p.hullSet = true
	}
	return p.cHull
}

func (p *player) hasCollided() {
	p.collision = true
}

func (g *game) playerUpdate() {
	if g.player.isDead {
		g.player.deadFrame++
		if g.player.deadFrame >= g.player.deadNumFrames {
			if g.player.lives <= 0 {
				disposeLevelImages()
				g.stopMusic()
				infiniteMusic = music1
				if g.state == gameInLevel1 {
					disposeLevel1Enemies()
				} else {
					disposeLevel2Enemies()
				}
				g.state = gameOver
			} else {
				g.player.isDead = false
			}
		}
	} else {
		if g.player.collision {
			g.player.lives--
			g.playSound(playerHurtSound)
			g.player.releasePowerUps(&(g.powerUpSet))
			g.player.bullets = initBulletSet()
			g.player.reset()
			g.player.isDead = true
			g.player.deadFrame = 0
		} else {
			if g.player.invincibleFrames > 0 {
				g.player.invincibleFrames--
			}
		}
	}
	g.player.hullSet = false
	g.player.cHull = nil
	g.player.laserOn = false
	if !g.player.isDead && !g.level.endLevel {
		g.player.move()
		g.player.managePowerUp(g)
		g.player.fire(g)
		g.player.moveOptions()
		g.player.bullets.update()
		g.player.updateBox()
		g.player.checkLiveWin()
	}
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

func (p *player) fire(g *game) {
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
		g.playSound(playerShotSound)
		p.lastBullet = 0
		if p.currentFire == 0 {
			for bNum := 0; bNum < p.numShot; bNum++ {
				p.bullets.addBullet(bullet{
					x: p.x + 30, y: p.y + 10,
					imageXShift: playerXBulletShift, imageYShift: playerYBulletShift,
					vx: pBulletSpeed, vy: pOtherBulletSpeed[bNum],
					ax: 0, ay: 0,
					image: playerBulletImage,
				})
			}
		} else if p.currentFire == 1 {
			p.bullets.addBigBullet(bullet{
				x: p.x, y: p.y,
				imageXShift: playerXBigBulletShift,
				vx:          pBulletSpeed, vy: 0,
				ax: 0, ay: 0,
				image: playerBigBulletImage,
			})
		}
		for oNum := 0; oNum < p.numOptions; oNum++ {
			p.bullets.addBullet(bullet{
				x: p.options[oNum].x, y: p.options[oNum].y,
				vx: pBulletSpeed, vy: 0,
				ax: 0, ay: 0,
				image: playerBulletImage,
			})
		}
	}
}

func (g *game) playerGetPowerUp() {
	g.player.points += pPointsPerPowerUp
	g.playSound(getBonusSound)
	if !g.player.allPowerUp {
		start := g.player.currentPowerUp
		g.player.currentPowerUp = (g.player.currentPowerUp + 1) % (pDifferentPowerUps + 1)
		for !g.player.isAppliablePowerUp(g.player.currentPowerUp) && g.player.currentPowerUp != start {
			g.player.currentPowerUp = (g.player.currentPowerUp + 1) % (pDifferentPowerUps + 1)
		}
		if g.player.currentPowerUp == start && !g.player.isAppliablePowerUp(g.player.currentPowerUp) {
			g.player.allPowerUp = true
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
				laserImage = laserImage1
			case 1:
				laserImage = laserImage2
			case 2:
				laserImage = laserImage3
			}
		}
	case 3:
		p.currentFire++
		p.currentFire = p.currentFire % 2
	case 4:
		p.numOptions++
	}
	p.usedPowerUp++
	p.currentPowerUp = 0
}

func (p *player) managePowerUp(g *game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if !p.allPowerUp && p.isAppliablePowerUp(p.currentPowerUp) {
			p.applyPowerUp()
			g.playSound(useBonusSound)
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
