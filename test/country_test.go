package model

import (
	"log"
	"sub2sing-box/model"
	"testing"
)

func TestCountry(t *testing.T) {
	log.Println(model.GetContryName("US 节点"))
}
