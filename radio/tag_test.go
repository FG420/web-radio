package radio

import (
	"log"
	"testing"
)

func TestGetTags(t *testing.T) {
	at := GetTags()

	for _, t := range at {
		log.Println(t.Name)
	}

}
