package main

import (
	"binary/calc"
	"flag"
	"fmt"
	"os"
)

func main() {
	var dConv int
	var bConv string
	var b2d bool
	var d2b bool

	flag.BoolVar(&d2b, "decimalToBinary", false, "convert decimal to binary")
	flag.IntVar(&dConv, "decimalNumber", 0, "value")

	flag.BoolVar(&b2d, "binaryToDecimal", false, "convert binary to decimal")
	flag.StringVar(&bConv, "binaryString", "000", "value")

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if b2d && !d2b {
		decimal := calc.DecimalConversion(bConv)
		fmt.Printf("The decimal equivalent to %s is %d\n", bConv, decimal)
	} else if d2b && !b2d {
		binary := calc.BinaryConversion(dConv)
		fmt.Printf("The binary equivalent to %d is %s\n", dConv, binary)
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
