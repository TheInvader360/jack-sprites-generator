package main

import (
	"reflect"
	"testing"
)

func TestProcessFile(t *testing.T) {
	cross := processFile("./Test/Valid/Cross.png")
	crossExpected := buildSpriteCross(false)
	if !reflect.DeepEqual(cross, crossExpected) {
		t.Errorf("Problem processing file (Cross) - found: %v, expected: %v", cross, crossExpected)
	}
}

func TestPopulateLineSums(t *testing.T) {
	cross := buildSpriteCross(false)
	crossExpected := buildSpriteCross(true)
	populateLineSums(&cross)
	if !reflect.DeepEqual(cross.lineSums, crossExpected.lineSums) {
		t.Errorf("Problem with line sums (Cross) - found: %v, expected: %v", cross.lineSums, crossExpected.lineSums)
	}

	diagonal := buildSpriteDiagonal(false)
	diagonalExpected := buildSpriteDiagonal(true)
	populateLineSums(&diagonal)
	if !reflect.DeepEqual(diagonal.lineSums, diagonalExpected.lineSums) {
		t.Errorf("Problem with line sums (Diagonal) - found: %v, expected: %v", diagonal.lineSums, diagonalExpected.lineSums)
	}

	square := buildSpriteSquare(false)
	squareExpected := buildSpriteSquare(true)
	populateLineSums(&square)
	if !reflect.DeepEqual(square.lineSums, squareExpected.lineSums) {
		t.Errorf("Problem with line sums (Square) - found: %v, expected: %v", square.lineSums, squareExpected.lineSums)
	}
}

func TestGenerateJackClass(t *testing.T) {
	sprites := []Sprite{buildSpriteCross(true), buildSpriteDiagonal(true), buildSpriteSquare(true)}
	jackCode := generateJackClass(sprites)
	expected := "class Sprites {\n\n    function void drawCross(int location) {\n        var int memAddress;\n        let memAddress = 16384+location;\n        do Memory.poke(memAddress+0, 0);\n        do Memory.poke(memAddress+32, 16386);\n        do Memory.poke(memAddress+64, 8196);\n        do Memory.poke(memAddress+96, 4104);\n        do Memory.poke(memAddress+128, 2064);\n        do Memory.poke(memAddress+160, 1056);\n        do Memory.poke(memAddress+192, 576);\n        do Memory.poke(memAddress+224, 384);\n        do Memory.poke(memAddress+256, 384);\n        do Memory.poke(memAddress+288, 576);\n        do Memory.poke(memAddress+320, 1056);\n        do Memory.poke(memAddress+352, 2064);\n        do Memory.poke(memAddress+384, 4104);\n        do Memory.poke(memAddress+416, 8196);\n        do Memory.poke(memAddress+448, 16386);\n        do Memory.poke(memAddress+480, 0);\n        return;\n    }\n\n    function void drawDiagonal(int location) {\n        var int memAddress;\n        let memAddress = 16384+location;\n        do Memory.poke(memAddress+0, ~32767);\n        do Memory.poke(memAddress+32, 16384);\n        do Memory.poke(memAddress+64, 8192);\n        do Memory.poke(memAddress+96, 4096);\n        do Memory.poke(memAddress+128, 2048);\n        do Memory.poke(memAddress+160, 1024);\n        do Memory.poke(memAddress+192, 512);\n        do Memory.poke(memAddress+224, 256);\n        do Memory.poke(memAddress+256, 128);\n        do Memory.poke(memAddress+288, 64);\n        do Memory.poke(memAddress+320, 32);\n        do Memory.poke(memAddress+352, 16);\n        do Memory.poke(memAddress+384, 8);\n        do Memory.poke(memAddress+416, 4);\n        do Memory.poke(memAddress+448, 2);\n        do Memory.poke(memAddress+480, 1);\n        return;\n    }\n\n    function void drawSquare(int location) {\n        var int memAddress;\n        let memAddress = 16384+location;\n        do Memory.poke(memAddress+0, -1);\n        do Memory.poke(memAddress+32, -32767);\n        do Memory.poke(memAddress+64, -32767);\n        do Memory.poke(memAddress+96, -32767);\n        do Memory.poke(memAddress+128, -32767);\n        do Memory.poke(memAddress+160, -32767);\n        do Memory.poke(memAddress+192, -32767);\n        do Memory.poke(memAddress+224, -32767);\n        do Memory.poke(memAddress+256, -32767);\n        do Memory.poke(memAddress+288, -32767);\n        do Memory.poke(memAddress+320, -32767);\n        do Memory.poke(memAddress+352, -32767);\n        do Memory.poke(memAddress+384, -32767);\n        do Memory.poke(memAddress+416, -32767);\n        do Memory.poke(memAddress+448, -32767);\n        do Memory.poke(memAddress+480, -1);\n        return;\n    }\n\n}\n\n"
	if jackCode != expected {
		t.Errorf("Problem with generated jack class - found: %s, expected: %s", jackCode, expected)
	}
}

func buildSpriteCross(lineSums bool) Sprite {
	s := Sprite{name: "Cross"}
	s.pixels = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if lineSums {
		s.lineSums = []int{0, 16386, 8196, 4104, 2064, 1056, 576, 384, 384, 576, 1056, 2064, 4104, 8196, 16386, 0}
	}
	return s
}

func buildSpriteDiagonal(lineSums bool) Sprite {
	s := Sprite{name: "Diagonal"}
	s.pixels = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if lineSums {
		s.lineSums = []int{-32768, 16384, 8192, 4096, 2048, 1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	}
	return s
}

func buildSpriteSquare(lineSums bool) Sprite {
	s := Sprite{name: "Square"}
	s.pixels = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	if lineSums {
		s.lineSums = []int{-1, -32767, -32767, -32767, -32767, -32767, -32767, -32767, -32767, -32767, -32767, -32767, -32767, -32767, -32767, -1}
	}
	return s
}
