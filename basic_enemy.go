package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	enemySpeed  = 0.05
	enemyHeight = 142
	enemyWidth  = 150
)

func newBasicEnemy(renderer *sdl.Renderer, position vector) *element {
	basicEnemy := &element{}

	basicEnemy.position = position
	basicEnemy.rotation = 180
	// basicEnemy.height = 142
	// basicEnemy.width = 150

	sr := newSpriteRenderer(basicEnemy, renderer, "sprites/LA01-1.bmp")
	basicEnemy.addComponent(sr)

	basicEnemy.active = true

	return basicEnemy
}
