package must

import (
	"image"
	"log"
	"os"

	// _ "image/gif"
	_ "image/png"
	// _ "image/jpeg"
)

func ReadImage(filename string) image.Image {
	r, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	m, _, err := image.Decode(r)
	if err != nil {
		log.Fatal(err)
	}

	return m
}
