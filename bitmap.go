package bitmap

import (
	"os"
	"image"
	"strings"
	"io/ioutil"
	"encoding/json"
	"fmt"

	_ "image/png"
)

type BitmapFont struct {
	font map[string][]string
}

type GlyphInfo struct {
	Char string
	Width int
	X, Y, W, H, Ox, Oy int
}

func New() BitmapFont {
	return BitmapFont{font: make(map[string][]string, 94)	}
}

func (b *BitmapFont) Load(filename string) map[string][]string {
	file, _ := os.Open(filename)
	image, _, _ := image.Decode(file)
	fontdescraw, _ := ioutil.ReadFile(strings.Split(filename, ".")[0]+".json")

	var fontdesc []GlyphInfo
	err := json.Unmarshal(fontdescraw, &fontdesc)
	if err != nil {
		panic(err) 
	}

	characters := " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	for idx, c := range strings.Split(characters, "") {
		xoffset := fontdesc[idx].X
		glyph := make([]string, 8)
		binrep := ""
		fmt.Println(c)

		for y := 0; y < 8; y++ {
			for x := 0; x < fontdesc[idx].W; x++ {
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