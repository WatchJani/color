package main

import "fmt"

func main() {
	reset := "\033[0m" //CLOSE TAG

	// ===========================
	// effect     //  code //  off
	// ===========================
	// Bold	      |   1	  |  21
	// Dim	      |   2	  |  22
	// Underline  |	  4	  |  24
	// Blink	  |   5	  |  25
	// Reverse	  |   7	  |  27
	// Hide	      |   8	  |  28

	// ============================================================================================
	// Color	// Example	//    Text   //	Background	 //   Bright Text	//    Bright Background
	// ============================================================================================
	// Black	| Black	    |    30	    |   40	         |   90	            |    100
	// Red	    | Red	    |    31	    |   41	         |   91	            |    101
	// Green	| Green	    |    32	    |   42	         |   92	            |    102
	// Yellow	| Yellow	|    33	    |   43	         |   93	            |    103
	// Blue	    | Blue	    |    34	    |   44	         |   94	            |    104
	// Magenta	| Magenta	|    35	    |   45	         |   95	            |    105
	// Cyan	    | Cyan	    |    36	    |   46	         |   96	            |    106
	// White	| White	    |    37	    |   47	         |   97	            |    107
	// Default	| none      |    39	    |   49	         |   99             |    109

	red := "\033[31m"
	green := "\033[32m"
	blue := "\033[34m"
	bold := "\033[1;8;2;120;50;255;48;2;150;255;0m"

	fmt.Println(red + "Crveni tekst" + reset)
	fmt.Println(green + "Zeleni tekst" + reset)
	fmt.Println(blue + "Plavi tekst" + reset)
	fmt.Println(bold + "Podebljani tekst" + reset)

	// Nijanse boja
	fmt.Println("\033[38;2;255;100;0mNarandžasta nijansa" + reset)
	fmt.Println("\033[38;1;2;155;255;255mCijan nijansa" + reset)
}
