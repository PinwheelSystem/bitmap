> ✍ Bitmap font loader library for Go.

This library loads font data from an image ("bitmap").  
It gets loaded into a Go `map` which means glyph data can be retrieved with `font[glyph]`, for example `font["A"]`.  
A glyph is an array of binary strings, where `1` represents an on pixel, and a `0` can represent the background.

So, to draw a glyph iterate through `font[glyph]`, like:
```golang
bm := bitmap.New()
font := bm.Load("font.png")
x := 0
y := 0
	
xx := x
yy := y
ch := font[glyph]
for i := 0; i < ch.Height; i++ {
 	bin := ch.Data[i] // Gets a line: a glyph is 8x8
 	binarr := strings.Split(bin, "")

 	for _, pix := range binarr {
 		if pix == "1" { draw(xx, yy) }
	 	xx += 1
	}
	yy += 1
	xx = x
}
```

# Used by
This library is used in [Pinwheel](https://github.com/PinwheelSystem/Pinwheel) for its pixel font.

*If you also use this for reason, you can make a PR to add your use here.*

# License 
[BSD 3-Clause](LICENSE)