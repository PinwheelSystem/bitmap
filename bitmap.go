/* 
The bitmap package loads a bitmap font from an image using a descriptor file.

A bitmap font is a font with glyphs that is literally bitmapped, with a transparent pixel being `0`,
and any opaque pixel being `1`.

A descriptor file describes the size (width, height) of each glyph. It is a JSON file, the same name as the font itself.
An example: the font is `fnt.png`, the descriptor will be `fnt.json`, and would look like:
`
[
	{
		"Char": " ",
		"Width": 7,
		"X": 1,
		"Y": 8,
		"W": 2,
		"H": 0,
		"Ox": 0,
		"Oy": 0
	},
	{
		"Char": "!",
		"Width": 5,
		"X": 4,
		"Y": 1,
		"W": 3,
		"H": 7,
		"Ox": 0,
		"Oy": 7
	},

	...
`
*/
package bitmap

import (
	"os"
	"image"
	"strings"
	"io/ioutil"
	"encoding/json"

	_ "image/png"
)

// BitmapFont holds a slice of Glyphs.
type BitmapFont struct {
	font map[string]Glyph
}

// GlyphInfo is the info about a glyph obtained from the descriptor file.
type GlyphInfo struct {
	Char string
	Width int
	X, Y, W, H, Ox, Oy int
}

// Glyph is a single character from the font.
type Glyph struct {
	Data []string
	Width, Height, Y int
}

// Returns a new, empty bitmap font.
func New() BitmapFont {
	return BitmapFont{font: make(map[string]Glyph, 94)	}
}

// Loads a bitmap font from `filename`.
// `filename` must include the extension.
func (b *BitmapFont) Load(filename string) map[string]Glyph {
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
		ch := fontdesc[idx]
		xoffset := ch.X
		glyph := make([]string, ch.H + 1)
		binrep := ""

		for y := 0; y < ch.H + 1; y++ {
			for x := 0; x < ch.W; x++ {
				_, _, _, a := image.At(xoffset + x, ch.Y + y).RGBA()
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
		b.font[c] = Glyph{glyph, ch.W, ch.H + 1, ch.Y}
	}
	return b.font
}