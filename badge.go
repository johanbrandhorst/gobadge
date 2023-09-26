package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/gophers"

	qrcode "github.com/skip2/go-qrcode"
)

const (
	WIDTH  = 160
	HEIGHT = 128
)

var pressed uint8

func Badge() {
	setCustomData()
	display.FillScreen(colors[BLACK])

	for {
		logo()
		myNameIs(YourName)
		blinky(YourTitleA1, YourTitleA2)
		blinky(YourTitleB1, YourTitleB2)
		blinky("Talk to me", "about Wasm")
		QR(YourQRText)
	}
}

func myNameIs(name string) {
	display.FillScreen(colors[WHITE])

	var r int16 = 8

	// black corners detail
	display.FillRectangle(0, 0, r, r, colors[BLACK])
	display.FillRectangle(0, HEIGHT-r, r, r, colors[BLACK])
	display.FillRectangle(WIDTH-r, 0, r, r, colors[BLACK])
	display.FillRectangle(WIDTH-r, HEIGHT-r, r, r, colors[BLACK])

	// round corners
	tinydraw.FilledCircle(&display, r, r, r, colors[RED])
	tinydraw.FilledCircle(&display, WIDTH-r-1, r, r, colors[RED])
	tinydraw.FilledCircle(&display, r, HEIGHT-r-1, r, colors[RED])
	tinydraw.FilledCircle(&display, WIDTH-r-1, HEIGHT-r-1, r, colors[RED])

	// top band
	display.FillRectangle(r, 0, WIDTH-2*r-1, r, colors[RED])
	display.FillRectangle(0, r, WIDTH, 26, colors[RED])

	// bottom band
	display.FillRectangle(r, HEIGHT-r-1, WIDTH-2*r-1, r+1, colors[RED])
	display.FillRectangle(0, HEIGHT-2*r-1, WIDTH, r, colors[RED])

	// top text
	const topText = "My name is"
	w32, _ := tinyfont.LineWidth(&freesans.Regular12pt7b, topText)
	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, (WIDTH-int16(w32))/2, 24, topText, colors[WHITE])

	// middle text
	w32, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, name)
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32))/2, 72, name, colors[BLACK])

	// gophers
	tinyfont.WriteLineColors(&display, &gophers.Regular32pt, WIDTH-48, 110, "BE", []color.RGBA{getRainbowRGB(100), getRainbowRGB(200)})

	for i := 0; i < 300; i++ {
		time.Sleep(20 * time.Millisecond)
		pressed, _ = buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			return
		}
	}
}

func blinky(topline, bottomline string) {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text so we could center them later
	w32top, _ := tinyfont.LineWidth(&freesans.Bold12pt7b, topline)
	w32bottom, _ := tinyfont.LineWidth(&freesans.Bold12pt7b, bottomline)
	for i := int16(0); i < 20; i++ {
		// show black text
		tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32top))/2, 50, topline, colors[BLACK])
		tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 100, bottomline, colors[BLACK])

		pressed, _ = buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			break
		}
	}
}

func logo() {
	display.FillRectangleWithBuffer(0, 0, WIDTH, HEIGHT, logoRGBA)
	for i := 0; i < 300; i++ {
		time.Sleep(20 * time.Millisecond)
		pressed, _ = buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			return
		}
	}
}

func QR(msg string) {
	qr, err := qrcode.New(msg, qrcode.Medium)
	if err != nil {
		println(err, 123)
	}

	qrbytes := qr.Bitmap()
	size := int16(len(qrbytes))

	factor := int16(HEIGHT / len(qrbytes))

	bx := (WIDTH - size*factor) / 2
	by := (HEIGHT - size*factor) / 2
	display.FillScreen(color.RGBA{109, 0, 140, 255})
	for y := int16(0); y < size; y++ {
		for x := int16(0); x < size; x++ {
			if qrbytes[y][x] {
				display.FillRectangle(bx+x*factor, by+y*factor, factor, factor, colors[0])
			} else {
				display.FillRectangle(bx+x*factor, by+y*factor, factor, factor, colors[1])
			}
		}
	}

	for i := 0; i < 300; i++ {
		time.Sleep(20 * time.Millisecond)
		pressed, _ = buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			return
		}
	}
}

func setCustomData() {
	if YourName == "" {
		YourName = DefaultName
	}

	if YourTitleA1 == "" {
		YourTitleA1 = DefaultTitleA1
	}

	if YourTitleA2 == "" {
		YourTitleA2 = DefaultTitleA2
	}

	if YourTitleB1 == "" {
		YourTitleB1 = DefaultTitleB1
	}

	if YourTitleB2 == "" {
		YourTitleB2 = DefaultTitleB2
	}

	if YourMarqueeTop == "" {
		YourMarqueeTop = DefaultMarqueeTop
	}

	if YourMarqueeMiddle == "" {
		YourMarqueeMiddle = DefaultMarqueeMiddle
	}

	if YourMarqueeBottom == "" {
		YourMarqueeBottom = DefaultMarqueeBottom
	}

	if YourQRText == "" {
		YourQRText = DefaultQRText
	}
}
