package main

import (
	"os"
	"search/config"
)

func main() {
	path := os.Args[1]
	kind := os.Args[2]

	stage := config.NewStageFromBoard(path, kind)

	if stage == nil {
		panic("Can't support algorithm> " + kind)
	}

	stage.BuildStage()
}
