package main

// go run *.go

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth                 = 720
	screenHeight                = 1280
	targetTicksPerSecond        = 60
	setFPS               uint32 = targetTicksPerSecond
)

var (
	delta      float64
	frameCount uint32
	timerFPS   uint32
	lastFrame  uint32
	lastTime   uint32
	fps        uint32
)

func calcFPS() {
	lastFrame = sdl.GetTicks()
	if lastFrame >= (lastTime + 1000) {
		lastTime = lastFrame
		fps = frameCount
		frameCount = 0
	}
}

func main() {
	// initilizeWindow()
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

	// sdl.RENDERER_PRESENTVSYNC added so not to use 100% gpu limits to 60fps
	// renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
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

	// for {
	for _ = range time.Tick(time.Microsecond * 1) {
		frameStartTime := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() { // Queue of sdl events
			switch event.(type) { // Allows window to be closed
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255) // White
		renderer.Clear()
		frameCount++
		timerFPS = sdl.GetTicks() - lastFrame
		if timerFPS < (1000 / setFPS) {
			sdl.Delay((1000 / setFPS) - timerFPS)
		}

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

		renderer.Present()
		calcFPS()
		delta = time.Since(frameStartTime).Seconds() * targetTicksPerSecond
		fmt.Print("\033[H\033[2J") // Clears console
		fmt.Println("FPS:", fps)
		// fmt.Printf("%+v\n", elements)
	}
}
