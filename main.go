package main

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

type Sprite struct {
	name     string
	pixels   []int
	lineSums []int
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
			populateLineSums(&sprite)
			sprites = append(sprites, sprite)
		}
	}

	if len(sprites) == 0 {
		log.Fatal("No *.png files found in path directory")
	}

	jackCode := generateJackClass(sprites)

	output := []byte(jackCode)
	err = ioutil.WriteFile(strings.Replace(path, "*", "Sprites.jack", 1), output, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func processFile(path string) Sprite {
	sprite := Sprite{}

	filename := filepath.Base(path)
	sprite.name = strings.Replace(filename, ".png", "", 1)

	file, _ := os.Open(path)
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	if img.Bounds().Max.X != 16 || img.Bounds().Max.Y != 16 {
		log.Fatal(fmt.Sprintf("Images must be 16x16px, %s is %dx%dpx", sprite.name, img.Bounds().Max.X, img.Bounds().Max.Y))
	}

	fmt.Println(sprite.name)
	pixels := []int{}
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if r+g+b > 0 {
				fmt.Printf(" ")
				pixels = append(pixels, 0)
			} else {
				fmt.Printf("█")
				pixels = append(pixels, 1)
			}
		}
		fmt.Println()
	}
	sprite.pixels = pixels

	return sprite
}

func populateLineSums(sprite *Sprite) {
	if len(sprite.pixels) != 256 {
		log.Fatal(fmt.Sprintf("Sprite must have 256 pixels (16x16), %s has %d", sprite.name, len(sprite.pixels)))
	}

	lineSums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i, val := range sprite.pixels {
		y := i / 16
		x := i % 16
		if x < 15 {
			lineSums[y] = lineSums[y] + (int(math.Pow(2, float64(x))) * val)
		} else {
			lineSums[y] = lineSums[y] - (32768 * val)
		}
	}

	sprite.lineSums = lineSums
}

func generateJackClass(sprites []Sprite) string {
	jackCode := "class Sprites {"
	for _, sprite := range sprites {
		jackCode += fmt.Sprintf("\n\n    function void draw%s(int location) {\n        var int memAddress;\n        let memAddress = 16384+location;\n", sprite.name)
		for i, lineSum := range sprite.lineSums {
			value := ""
			if lineSum == -32768 {
				value = "~32767"
			} else {
				value = fmt.Sprint(lineSum)
			}
			jackCode += fmt.Sprintf("        do Memory.poke(memAddress+%d, %s);\n", i*32, value)
		}
		jackCode += "        return;\n    }"
	}
	jackCode += "\n\n}\n\n"
	return jackCode
}
