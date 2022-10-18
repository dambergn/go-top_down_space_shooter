package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 10
	bulletSpeed = 10
)

// type bullet struct {
// 	tex    *sdl.Texture
// 	x, y   float64
// 	angle  float64
// 	active bool
// }

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

// func newBullet(renderer *sdl.Renderer) (bul bullet) {
// 	bul.tex = textureFromBMP(renderer, "sprites/M484BulletCollection1.bmp")
// 	return bul
// }

// func (bul *bullet) draw(renderer *sdl.Renderer) {
// 	if !bul.active {
// 		return
// 	}

// 	x := bul.x - 8/2.0
// 	y := bul.y - 10/2.0

// 	renderer.Copy(bul.tex,
// 		&sdl.Rect{X: 12, Y: 12, W: 8, H: 10},
// 		&sdl.Rect{X: int32(x), Y: int32(y), W: 8, H: 10})
// }

// func (bul *bullet) update() {
// 	bul.x += bulletSpeed * math.Cos(bul.angle)
// 	bul.y += bulletSpeed * math.Sin(bul.angle)

// 	if bul.x > screenWidth || bul.x < 0 || bul.y > screenHeight || bul.y < 0 {
// 		bul.active = false
// 	}
// }

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
