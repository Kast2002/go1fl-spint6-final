package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(text string) (string, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return "", errors.New("пустая строка")
	}
	symbols := strings.Split(text, "")
	var isMorse bool
	var out string
	for _, v := range symbols {
		if v == "." || v == " " || v == "-" {
			isMorse = true
			break
		} else {
			isMorse = false
		}
	}

	if isMorse {
		out = morse.ToText(text)
	} else {
		out = morse.ToMorse(text)
	}

	return out, nil
}
