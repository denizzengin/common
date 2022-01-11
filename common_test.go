package common

import (
	"log"
	"testing"
)

func TestIsDevelopment(t *testing.T){
	dev := IsDevelopment()
	if !dev {
		log.Fatal("Error default env variable must be Dev, incoming:", GetEnv())
	}
}