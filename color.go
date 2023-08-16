package color

import (
	"fmt"
	"strconv"
	"strings"
)

// ===========================
// effect     //  code //  clear
// ===========================
// Bold	      |   1	   |  21
// Dim	      |   2	   |  22
// Underline  |	  4	   |  24
// Blink	  |   5	   |  25
// Reverse	  |   7	   |  27
// Hide	      |   8	   |  28

// ============================================================================================
// Color	// Example	//    Text   //	Background	 //   Bright Text	//    Bright Background
// ============================================================================================
// Black	| Black	    |    30	    |   40	         |   90	            |    100
// Red	   	| Red	    |    31	    |   41	         |   91	            |    101
// Green	| Green	    |    32	    |   42	         |   92	            |    102
// Yellow	| Yellow    |    33	    |   43	         |   93	            |    103
// Blue	    	| Blue	    |    34	    |   44	         |   94	            |    104
// Magenta	| Magenta   |    35	    |   45	         |   95	            |    105
// Cyan	    	| Cyan	    |    36	    |   46	         |   96	            |    106
// White	| White	    |    37	    |   47	         |   97	            |    107
// Default	| none      |    39	    |   49	         |   99             |    109

type OptFunc func(*Color)

type ColorFunc func(string) string

// Base attributes
const (
	reset      = "0"
	bold       = "1"
	faint      = "2"
	italic     = "3"
	underline  = "4"
	blinkSlow  = "5"
	blinkRapid = "6"
	reverse    = "7"
	concealed  = "8"
	crossedOut = "9"
)

// Foreground text colors
const (
	fgBlack   = "30"
	fgRed     = "31"
	fgGreen   = "32"
	fgYellow  = "33"
	fgBlue    = "34"
	fgMagenta = "35"
	fgCyan    = "36"
	fgWhite   = "37"
	fgDefault = "39"
)

// Foreground Hi-Intensity text colors
const (
	fgHiBlack   = "90"
	fgHiRed     = "91"
	fgHiGreen   = "92"
	fgHiYellow  = "93"
	fgHiBlue    = "94"
	fgHiMagenta = "95"
	fgHiCyan    = "96"
	fgHiWhite   = "97"
	fgHiDefault = "99"
)

// Background text colors
const (
	bgBlack   = "40"
	bgRed     = "41"
	bgGreen   = "42"
	bgYellow  = "43"
	bgBlue    = "44"
	bgMagenta = "45"
	bgCyan    = "46"
	bgWhite   = "47"
	bgDefault = "49"
)

// Background Hi-Intensity text colors
const (
	bgHiBlack   = "100"
	bgHiRed     = "101"
	bgHiGreen   = "102"
	bgHiYellow  = "103"
	bgHiBlue    = "104"
	bgHiMagenta = "105"
	bgHiCyan    = "106"
	bgHiWhite   = "107"
	bgHiDefault = "109"
)

type Color struct {
	attribute  []string
	color      string
	background string
}

func defaultColor() Color {
	return Color{
		color:      fgDefault,
		background: bgDefault,
	}
}

func New(fns ...OptFunc) ColorFunc {
	df := defaultColor()

	for _, fn := range fns {
		fn(&df)
	}

	return colorFormat(&df)
}

func (c *Color) CodeColor() string {
	if len(c.attribute) == 0 {
		attributes(c, reset)
	}

	return fmt.Sprintf("%s;%s;%s", moreAttributes(c.attribute), c.color, c.background)
}

func colorFormat(color *Color) ColorFunc {
	code := color.CodeColor()

	return func(s string) string {
		return fmt.Sprintf("\033[%sm%s\033[0m", code, s)
	}
}

func moreAttributes(attributes []string) string {
	return strings.Join(attributes, ";")
}

func ColorRGB(red, green, blue int) OptFunc {
	rgb := rgbConvertor(red, green, blue)

	return func(c *Color) {
		c.color = fmt.Sprintf("38;2;%s;%s;%s", rgb[0], rgb[1], rgb[2])
	}
}

func rgbConvertor(r, g, b int) []string {
	return []string{strconv.Itoa(r), strconv.Itoa(g), strconv.Itoa(b)}
}

//All func Color

// Black Foreground text colors
func Black(c *Color) { c.color = fgBlack }

// Black Foreground Hi-Intensity text colors
func HiBlack(c *Color) { c.color = fgHiBlack }

// Red Foreground text colors
func Red(c *Color) { c.color = fgRed }

// Red Foreground Hi-Intensity text colors
func HiRed(c *Color) { c.color = fgHiRed }

// Green Foreground text colors
func Green(c *Color) { c.color = fgGreen }

// Green Foreground Hi-Intensity text colors
func HiGreen(c *Color) { c.color = fgHiGreen }

// Yellow Foreground text colors
func Yellow(c *Color) { c.color = fgYellow }

// Yellow Foreground Hi-Intensity text colors
func HiYellow(c *Color) { c.color = fgHiYellow }

// Blue Foreground text colors
func Blue(c *Color) { c.color = fgBlue }

// Blue Foreground Hi-Intensity text colors
func HiBlue(c *Color) { c.color = fgHiBlue }

// Magenta Foreground text colors
func Magenta(c *Color) { c.color = fgMagenta }

// Magenta Foreground Hi-Intensity text colors
func HiMagenta(c *Color) { c.color = fgHiMagenta }

// Cyan Foreground text colors
func Cyan(c *Color) { c.color = fgCyan }

// Cyan Foreground Hi-Intensity text colors
func HiCyan(c *Color) { c.color = fgHiCyan }

// White Foreground text colors
func White(c *Color) { c.color = fgWhite }

// White Foreground Hi-Intensity text colors
func HiWhite(c *Color) { c.color = fgHiWhite }

func attributes(c *Color, attribute string) {
	c.attribute = append(c.attribute, attribute)
}

// set text Bold
func Bold(c *Color) { attributes(c, bold) }

// set text Faint
func Faint(c *Color) { attributes(c, faint) }

// set text Italic
func Italic(c *Color) { attributes(c, italic) }

// set text Underline
func Underline(c *Color) { attributes(c, underline) }

// set text Reset
func Reset(c *Color) { attributes(c, reset) }

// set text BlinkSlow
func BlinkSlow(c *Color) { attributes(c, blinkSlow) }

// set text BlinkRapid
func BlinkRapid(c *Color) { attributes(c, blinkRapid) }

// set text ReverseVideo
func Reverse(c *Color) { attributes(c, reverse) }

// set text Concealed
func Concealed(c *Color) { attributes(c, concealed) }

// set text CrossedOut
func CrossedOut(c *Color) { attributes(c, crossedOut) }
