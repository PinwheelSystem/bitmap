package bitmap

import (
	"os"
	"image"
	"strings"

	_ "image/png"
)

type BitmapFont struct {
	font map[string][]string
}

func New() BitmapFont {
	return BitmapFont{font: make(map[string][]string, 94)	}
}

func (b *BitmapFont) Load(filename string) map[string][]string {
	file, _ := os.Open(filename)
	image, _, _ := image.Decode(file)

	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890~-+!@#$%^&*()_={}[]|\\:;\"'<,>.?/"
	for idx, c := range strings.Split(characters, "") {
		xoffset := idx * 8
		glyph := make([]string, 8)
		binrep := ""

		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				_, _, _, a := image.At(xoffset + x, y).RGBA()
				a = a >> 8
				if a == 255 {
					binrep += "1"
				} else {
					binrep += "0"
				}
			}
			glyph[y] = binrep
			binrep = ""
		}
		b.font[c] = glyph
	}
	return b.font
}