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
	"fmt"
	"image/color"

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
	vCap             float64
	numOptions       int
	currentPosition  int
	positionHistory  [pMoveRecorded]playerPosition
	options          [pMaxOption]option
	currentPowerUp   int
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
}

type playerPosition struct {
	x float64
	y float64
}

const (
	pInitX               = 100
	pInitY               = screenHeight / 2
	pWidth               = 45
	pHeight              = 15
	pMaxVCap             = 10
	pVStep               = 3
	pVInit               = 4
	pAx                  = 1
	pAy                  = 1
	pBulletInterval      = 15
	pBulletSpeed         = 6
	pMaxShot             = 5  // for basic shot only
	pMinDelay            = 25 // for big shot only
	pMaxDelay            = 50
	pShotDelayStep       = 5
	pMaxShotWidth        = 20 // for laser shot only
	pMinShotWidth        = 5
	pShotWidthIncrease   = 5
	pMaxOption           = 3
	pDifferentPowerUps   = 4
	pMoveRecorded        = 16
	pFrameBetweenOptions = 5
	pInvicibleDuration   = 120
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
}

func (p player) draw(screen *ebiten.Image) {
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
		p.laser.draw(screen, color.RGBA{0, 255, 0, 255})
	}
	p.bullets.draw(screen, color.RGBA{0, 255, 0, 255})
}

func (p player) drawUI(screen *ebiten.Image) {
	var s string
	switch p.currentPowerUp {
	case 0:
		s = "0: no power up"
	case 1:
		s = "1: speed up"
	case 2:
		s = "2: more shots"
	case 3:
		s = "3: fire change"
	case 4:
		s = "4: more options"
	}
	ebitenutil.DebugPrintAt(screen, s, 0, 1030)
	switch p.currentFire {
	case 0:
		s = "0: basic shot"
	case 1:
		s = "1: big shot"
	case 2:
		s = "2: laser"
	}
	ebitenutil.DebugPrintAt(screen, s, 0, 980)
	s = fmt.Sprint("points ", p.points)
	ebitenutil.DebugPrintAt(screen, s, 0, 1050)
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
			point{p.x - p.xSize/2, p.y + p.ySize/2},
			point{p.x + p.xSize/2, p.y + p.ySize/2},
			point{p.x - p.xSize/2, p.y - p.ySize/2},
		}
		p.hullSet = true
	}
	return p.cHull
}

func (p *player) hasCollided() {
	p.collision = true
}

func (p *player) checkCollisions(bs []*bullet, es []*enemy, ps []*powerUp) {
	for oNum := 0; oNum < p.numOptions; oNum++ {
		o := p.options[oNum]
		for _, b := range bs {
			collide(o, b)
		}
		for _, e := range es {
			collide(o, e)
		}
	}
	if p.invincibleFrames <= 0 {
		for _, b := range bs {
			collide(p, b)
		}
		for _, e := range es {
			collide(p, e)
		}
	}
	if p.laserOn {
		for _, e := range es {
			collide(&(p.laser), e)
		}
	}
	for _, b := range p.bullets.bullets {
		for _, e := range es {
			collide(b, e)
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

func (p *player) update() {
	if p.collision {
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
	p.y += p.vy
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
		return
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
				})
			}
		} else {
			p.bullets.addBigBullet(bullet{
				x: p.x, y: p.y,
				vx: pBulletSpeed, vy: 0,
				ax: 0, ay: 0,
			})
		}
		for oNum := 0; oNum < p.numOptions; oNum++ {
			p.bullets.addBullet(bullet{
				x: p.options[oNum].x, y: p.options[oNum].y,
				vx: pBulletSpeed, vy: 0,
				ax: 0, ay: 0,
			})
		}
	}
}

func (p *player) getPowerUp() {
	if !p.allPowerUp {
		start := p.currentPowerUp
		p.currentPowerUp = (p.currentPowerUp + 1) % (pDifferentPowerUps + 1)
		for !p.isAppliablePowerUp() && p.currentPowerUp != start {
			p.currentPowerUp = (p.currentPowerUp + 1) % (pDifferentPowerUps + 1)
		}
		if p.currentPowerUp == start && !p.isAppliablePowerUp() {
			p.allPowerUp = true
		}
	}
}

func (p player) isAppliablePowerUp() bool {
	switch p.currentPowerUp {
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
		}
	case 3:
		p.currentFire++
		p.currentFire = p.currentFire % 3
	case 4:
		p.numOptions++
	}
	p.currentPowerUp = 0
}

func (p *player) managePowerUp() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if !p.allPowerUp && p.isAppliablePowerUp() {
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
