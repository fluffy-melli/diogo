package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/fluffy-melli/krapo"
	"github.com/fluffy-melli/krapo/RDR_CMP"
	"github.com/fluffy-melli/krapo/render"
)

var API_KEY = ""

func TestGIF(t *testing.T) {
	urls, err := RDR_CMP.GetImagesURL(API_KEY, krapo.LTime(1)) // RDR_CMP.GetAllURL(API_KEY)
	if err != nil {
		log.Fatalln(err)
	}
	gif, err := render.GIF(urls, 10, true)
	if err != nil {
		log.Fatalln(err)
	}
	err = render.Write("./test.gif", gif)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestALLURL(t *testing.T) {
	urls, err := RDR_CMP.GetAllURL(API_KEY)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(len(urls))
}
