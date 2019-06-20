package qrchar

import (
	"bytes"
	"github.com/skip2/go-qrcode"
)

const (
	CHARS     = "█▀▄ "
	TRIM_SIZE = 3
)

var RUNES = []rune(CHARS)

func Encode(content string) (string, error) {
	qr, e := qrcode.New(content, qrcode.Medium)
	if e != nil {
		return "", e
	}

	var s bytes.Buffer
	bmp := qr.Bitmap()
	for i := TRIM_SIZE; i < len(bmp)-TRIM_SIZE; i += 2 {
		for j := TRIM_SIZE; j < len(bmp[i])-TRIM_SIZE; j += 1 {
			upper := getBitmap(bmp, i, j)
			lower := getBitmap(bmp, i+1, j)
			s.WriteRune(getChar(upper, lower))
		}
		s.WriteString("\n")
	}

	return s.String(), nil
}

func EncodePNG(content string) ([]byte, error) {
	return qrcode.Encode(content, qrcode.Medium, 256)
}

func getBitmap(bmp [][]bool, i, j int) bool {
	if bmp == nil || i < TRIM_SIZE || i >= len(bmp)-TRIM_SIZE {
		return true
	}
	if bmp[i] == nil || j < TRIM_SIZE || j >= len(bmp[i])-TRIM_SIZE {
		return true
	}
	return bmp[i][j]
}

func getChar(upper, lower bool) rune {
	p := 0
	if upper {
		p += 2
	}
	if lower {
		p += 1
	}
	return RUNES[p]
}
