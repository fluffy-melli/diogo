package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/fluffy-melli/krapo/MID_FCST"
)

func TestFcst(t *testing.T) {
	info, err := MID_FCST.GetFcst(API_KEY)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(info[0])
}
