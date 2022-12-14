package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 10
	bulletSpeed = 10
)

func newBullet(renderer *sdl.Renderer) *element {
	bullet := &element{}

	sr := newSpriteRenderer(bullet, renderer, "sprites/bullet1.bmp")
	// sr := newSpriteRenderer(bullet, renderer, "sprites/M484BulletCollection1.bmp")
	bullet.addComponent(sr)

	mover := newBulletMover(bullet, bulletSpeed)
	bullet.addComponent(mover)

	col := circle{
		center: bullet.position,
		radius: 5,
	}
	bullet.collisions = append(bullet.collisions, col)

	bullet.tag = "bullet"

	bullet.active = false

	return bullet
}

var bulletPool []*element

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bul := newBullet(renderer)
		elements = append(elements, bul)
		bulletPool = append(bulletPool, bul)
	}
}

func bulletFromPool() (*element, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}
	return nil, false
}
