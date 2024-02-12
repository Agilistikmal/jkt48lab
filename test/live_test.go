package test

import (
	"log"
	"testing"

	"github.com/agilistikmal/jkt48lab/app/service"
)

func TestGet(t *testing.T) {
	l := service.GetIDNLives()
	for _, a := range l {
		log.Println(a.Member)
	}
}
