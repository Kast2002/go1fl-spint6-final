package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(text string) (string, error) {
	copyText := text
	symbols := strings.Split(copyText, "")
	var isMorse bool
	var out string
	for _, v := range symbols {
		if v == "." || v == " " || v == "-" {
			isMorse = true
		} else {
			isMorse = false
		}
	}

	if isMorse {
		out = morse.ToMorse(text)
	} else {
		out = morse.ToText(text)
	}

	return out, nil
}
