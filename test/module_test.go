package main

import (
	"log"
	"testing"

	"github.com/fluffy-melli/krapo"
	"github.com/fluffy-melli/krapo/RDR_CMP"
	"github.com/fluffy-melli/krapo/render"
)

func Test(t *testing.T) {
	urls, err := RDR_CMP.GetImagesURL("", krapo.Time())
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
