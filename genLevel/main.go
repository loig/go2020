package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 120; i++ {
		k1 := rand.Intn(15) + 2
		k2 := rand.Intn(k1-1) + 1
		d := rand.Intn(60) + 10
		typeNum := rand.Intn(4)
		typeString := ""
		switch typeNum {
		case 0:
			typeString = "staticExplodingEnemy"
		default:
			typeString = "staticEnemy"
		}
		fmt.Print(
			"spawn{\n enemies:[]enemySpawn{\nenemySpawn{enemyType: ", typeString, ", y: float64(", k2, "*screenHeight) / ", k1, "},\n},\nframeDelay: ", d, ",\n},\n",
		)
	}
}
