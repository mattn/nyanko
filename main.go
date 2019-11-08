package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/mattn/go-sixel"
	"github.com/nfnt/resize"
)

func main() {
	resp, err := http.Get("https://aws.random.cat/meow")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var r struct {
		File string `json:"file"`
	}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = http.Get(r.File)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	img = resize.Resize(600, 0, img, resize.Lanczos3)
	err = sixel.NewEncoder(os.Stdout).Encode(img)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.File)
}
