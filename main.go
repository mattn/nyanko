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
	"path/filepath"

	"github.com/mattn/go-sixel"
	"github.com/nfnt/resize"
)

func nyanko() (string, error) {
	resp, err := http.Get("https://aws.random.cat/meow")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var r struct {
		File string `json:"file"`
	}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", err
	}
	return r.File, nil
}

func innu() (string, error) {
	resp, err := http.Get("https://random.dog/woof.json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var r struct {
		URL string `json:"url"`
	}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", err
	}
	return r.URL, nil
}

func shiba() (string, error) {
	resp, err := http.Get("http://shibe.online/api/shibes")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var r []string
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", err
	}
	return r[0], nil
}

func tori() (string, error) {
	resp, err := http.Get("https://some-random-api.ml/img/birb")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var r struct {
		URL string `json:"link"`
	}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", err
	}
	return r.URL, nil
}

func main() {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	var uri string
	switch filepath.Base(exe) {
	default:
		fallthrough
	case "nyanko":
		uri, err = nyanko()
	case "innu":
		uri, err = innu()
	case "shiba":
		uri, err = shiba()
	case "tori":
		uri, err = tori()
	}
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get(uri)
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
	fmt.Println()
	fmt.Println(uri)
}
