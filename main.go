package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Sprite struct {
	name   string
	pixels []int
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Missing path parameter")
	}
	path := os.Args[1]

	info, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	if info.IsDir() {
		if strings.HasSuffix(path, "/") {
			path = path + "*"
		} else {
			path = path + "/*"
		}
	} else {
		log.Fatal("Path is not a directory")
	}

	files, _ := filepath.Glob(path)

	sprites := []Sprite{}

	for _, file := range files {
		if strings.HasSuffix(file, ".png") {
			sprite := processFile(file)
			sprites = append(sprites, sprite)
		}
	}

	fmt.Println(sprites)
}

func processFile(path string) Sprite {
	sprite := Sprite{name: filepath.Base(path)}
	return sprite
}
