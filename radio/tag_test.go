package radio

import (
	"log"
	"testing"
)

func TestGetTags(t *testing.T) {
	tt := GetTags()

	for _, t := range tt {
		log.Println(t.Name)

	}

	// log.Println(tt)
}
