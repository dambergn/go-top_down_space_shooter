package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type bulletMover struct {
	container *element
	speed     float64
}

func newBulletMover(container *element, speed float64) *bulletMover {
	return &bulletMover{
		container: container,
		speed:     speed,
	}
}

func (mover *bulletMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *bulletMover) onUpdate() error {
	cont := mover.container

	cont.position.x += bulletSpeed * math.Cos(cont.rotation) * delta
	cont.position.y += bulletSpeed * math.Sin(cont.rotation) * delta

	if cont.position.x > screenWidth || cont.position.x < 0 ||
		cont.position.y > screenHeight || cont.position.y < 0 {
		cont.active = false
	}

	mover.container.collisions[0].center = cont.position

	return nil
}

func (mover *bulletMover) onCollision(other *element) error {
	mover.container.active = false
	return nil
}
