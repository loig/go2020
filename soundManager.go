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
	"io/ioutil"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type soundManager struct {
	audioContext *audio.Context
	musicPlayer  *audio.Player
}

const (
	playerShotSound int = iota
	enemyShotSound
	playerHurtSound
	enemyHurtSound
	bossHurtSound
	menuSound
)

// loop the music
func (g *game) updateMusic() {
	if g.audio.musicPlayer != nil {
		if !g.audio.musicPlayer.IsPlaying() {
			g.audio.musicPlayer.Rewind()
			g.audio.musicPlayer.Play()
		}
	} else {
		var error error
		g.audio.musicPlayer, error = audio.NewPlayer(g.audio.audioContext, infiniteMusic)
		if error != nil {
			log.Panic("Audio problem:", error)
		}
		g.audio.musicPlayer.Play()
	}

}

// stop the music
func (g *game) stopMusic() {
	if g.audio.musicPlayer != nil && g.audio.musicPlayer.IsPlaying() {
		error := g.audio.musicPlayer.Close()
		if error != nil {
			log.Panic("Sound problem:", error)
		}
		g.audio.musicPlayer = nil
	}
}

// play a sound
func (g *game) playSound(sound int) {
	var soundBytes []byte
	switch sound {
	case playerShotSound:
		soundBytes = playerShotBytes
	case enemyShotSound:
		soundBytes = enemyShotBytes
	case playerHurtSound:
		soundBytes = playerHurtBytes
	case enemyHurtSound:
		soundBytes = enemyHurtBytes
	case bossHurtSound:
		soundBytes = bossHurtBytes
	case menuSound:
		soundBytes = menuBytes
	}
	soundPlayer := audio.NewPlayerFromBytes(g.audio.audioContext, soundBytes)
	soundPlayer.Play()
}

// load all audio assets
func (g *game) initAudio() {
	var error error
	g.audio.audioContext = audio.NewContext(44100)

	// music
	soundFile, error := ebitenutil.OpenFile("assets/intro.mp3")
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	sound, error := mp3.Decode(g.audio.audioContext, soundFile)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	tduration, _ := time.ParseDuration("30s")
	duration := tduration.Seconds()
	bytes := int64(math.Round(duration * 4 * float64(44100)))
	music1 = audio.NewInfiniteLoop(sound, bytes)

	soundFile, error = ebitenutil.OpenFile("assets/level.mp3")
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	sound, error = mp3.Decode(g.audio.audioContext, soundFile)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	tduration, _ = time.ParseDuration("30s")
	duration = tduration.Seconds()
	bytes = int64(math.Round(duration * 4 * float64(44100)))
	music2 = audio.NewInfiniteLoop(sound, bytes)
	infiniteMusic = music2

	// sounds
	soundFile, error = ebitenutil.OpenFile("assets/playershot.mp3")
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	sound, error = mp3.Decode(g.audio.audioContext, soundFile)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	playerShotBytes, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	soundFile, error = ebitenutil.OpenFile("assets/enemyshot.mp3")
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	sound, error = mp3.Decode(g.audio.audioContext, soundFile)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	enemyShotBytes, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	soundFile, error = ebitenutil.OpenFile("assets/playerhurt.mp3")
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	sound, error = mp3.Decode(g.audio.audioContext, soundFile)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	playerHurtBytes, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	soundFile, error = ebitenutil.OpenFile("assets/enemyhurt.mp3")
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	sound, error = mp3.Decode(g.audio.audioContext, soundFile)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	enemyHurtBytes, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	soundFile, error = ebitenutil.OpenFile("assets/bosshurt.mp3")
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	sound, error = mp3.Decode(g.audio.audioContext, soundFile)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	bossHurtBytes, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	soundFile, error = ebitenutil.OpenFile("assets/menu.mp3")
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	sound, error = mp3.Decode(g.audio.audioContext, soundFile)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	menuBytes, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
}