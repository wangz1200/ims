package orm

import (
	"log"
	"testing"
)

func TestInitOrm(t *testing.T) {
	if err := Init(); err != nil {
		log.Fatal(err)
		return
	}
}
