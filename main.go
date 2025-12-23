package main

import (
	"embed"
	"juandserrano/tlt-2d/game"
)

//go:embed assets/*
var embedFS embed.FS

func main() {
	game.Run(&embedFS)
}
