package main

import (
	"3dGamePractice/config"
	"3dGamePractice/game"
	"3dGamePractice/graphic"
)

func main() {
	windowWidth := config.GetWindowWidth()
	windowHeight := config.GetWindowHeight()
	targetFrame := config.GetTargetFrame()

	graphic.RaylibSet(
		windowWidth,
		windowHeight,
		targetFrame,
	)
	defer graphic.RaylibClose()

	graphic.RaylibLoop(
		game.Logic,
	)
}
