package config

import (
	"fmt"
	"os"
)

var (
	windowWidth  int
	windowHeight int
	targetFrame  int
)

func GetWindowWidth() int {
	fmt.Print("Enter window width: ")
	if _, err := fmt.Scan(&windowWidth); err != nil {
		fmt.Println("Invalid input for width:", err)
		os.Exit(1)
	}
	return windowWidth
}

func GetWindowHeight() int {
	fmt.Print("Enter window height: ")
	if _, err := fmt.Scan(&windowHeight); err != nil {
		fmt.Println("Invalid input for height:", err)
		os.Exit(1)
	}
	return windowHeight
}

func GetTargetFrame() int {
	fmt.Print("Enter target frame: ")
	if _, err := fmt.Scan(&targetFrame); err != nil {
		fmt.Println("Invalid input for targetFrame:", err)
		os.Exit(1)
	}
	return targetFrame
}
