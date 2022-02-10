package notifile

import (
	"log"
	"testing"
)

// TestCreate is to create a new notification file
func TestCreate(t *testing.T) {
	data := Data{
		Channel:  []string{"gotify"},
		Headline: "J&J Afterialo",
		Message:  "Bitte schau doch mal nach.",
	}
	err := Create(data, "/var/lib/echgo/notification")
	if err != nil {
		log.Fatalln(err)
	}
}
