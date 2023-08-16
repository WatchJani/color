package color

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	testColors := []struct {
		text string
		fn   func(*Color)
		code string
	}{
		{text: "black", fn: Black, code: fgBlack},
		{text: "red", fn: Red, code: fgRed},
		{text: "green", fn: Green, code: fgGreen},
		{text: "yellow", fn: Yellow, code: fgYellow},
		{text: "blue", fn: Blue, code: fgBlue},
		{text: "magent", fn: Magenta, code: fgMagenta},
		{text: "cyan", fn: Cyan, code: fgCyan},
		{text: "white", fn: White, code: fgWhite},
		{text: "hblack", fn: HiBlack, code: fgHiBlack},
		{text: "hred", fn: HiRed, code: fgHiRed},
		{text: "hgreen", fn: HiGreen, code: fgHiGreen},
		{text: "hyellow", fn: HiYellow, code: fgHiYellow},
		{text: "hblue", fn: HiBlue, code: fgHiBlue},
		{text: "hmagent", fn: HiMagenta, code: fgHiMagenta},
		{text: "hcyan", fn: HiCyan, code: fgHiCyan},
		{text: "hwhite", fn: HiWhite, code: fgHiWhite},
	}

	for _, c := range testColors {
		color := New(c.fn)

		escapedForm := fmt.Sprintf("\033[0;%s;49m%s\033[0m", c.code, c.text) //bgDefault ==> 49

		scannedLine := color(c.text)

		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}
}
