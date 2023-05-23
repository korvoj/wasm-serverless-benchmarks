package main

import (
	"fmt"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		filename := os.Args[1]
		img, err := imgio.Open(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		inverted := effect.Invert(img)
		resized := transform.Resize(inverted, 800, 800, transform.Linear)
		rotated := transform.Rotate(resized, 45, nil)

		if err := imgio.Save(fmt.Sprintf("%s-output.png", filename), rotated, imgio.PNGEncoder()); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		log.Fatalln("Please pass a file path to the image to be converted as the first argument...")
	}
}
