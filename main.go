package main

// go run *.go

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 720
	screenHeight = 1280
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		// panic(err)
		fmt.Println("initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Gaming in GO",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Initializing Window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	elements = append(elements, newPlayer(renderer))

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidth + (enemyWidth / 2.0)
			y := float64(j)*enemyHeight + (enemyHeight / 2.0)

			enemy := newBasicEnemy(renderer, vector{x, y})
			elements = append(elements, enemy)
		}
	}

	initBulletPool(renderer)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() { // Queue of sdl events
			switch event.(type) { // Allows window to be closed
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255) // White
		renderer.Clear()

		for _, elem := range elements {
			if elem.active {
				err = elem.update()
				if err != nil {
					fmt.Println("updating element:", err)
					return
				}
				err = elem.draw(renderer)
				if err != nil {
					fmt.Println("drawing element:", elem)
					return
				}
			}
		}

		if err := checkCollisions(); err != nil {
			fmt.Println("Checking collisions:", err)
			return
		}

		// for _, bul := range bulletPool {
		// 	bul.draw(renderer)
		// 	bul.update()
		// }

		renderer.Present()
	}
}
