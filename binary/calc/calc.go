package calc

import (
	"math"
	"strings"
)

func populateSlot(slotCount int) []string {
	slot := make([]string, slotCount)

	for i := 0; i < slotCount; i++ {
		slot[i] = "0"
	}

	return slot
}

const BASE int = 2

func getExponent(number int) int {
	if number <= 0 {
		return 0
	}

	logBase2 := math.Log2(float64(number))

	closestExponent := int(logBase2)

	return closestExponent
}

func calcToBinary(slot []string, number int) []string {
	if number != 0 {
		exponent := getExponent(number)

		slot[(len(slot)-1)-exponent] = "1"

		diff := number - int(math.Pow(float64(BASE), float64(exponent)))

		if diff != 0 {
			calcToBinary(slot, diff)
		}
	}

	return slot
}

func calcToDecimal(index int, binary string, total int) int {
	if index >= len(binary) {
		return total
	}

	bit := binary[len(binary)-1-index]
	if bit == '1' {
		total += int(math.Pow(float64(BASE), float64(index)))
	}

	return calcToDecimal(index+1, binary, total)
}

func DecimalConversion(binary string) int {
	total := 0
	index := 0
	return calcToDecimal(index, binary, total)
}

func BinaryConversion(number int) string {
	var slotCount int

	exponent := getExponent(number)

	if exponent <= 3 {
		slotCount = 3
	} else {
		slotCount = exponent + 1
	}

	slot := populateSlot(int(slotCount + 1))

	slot = calcToBinary(slot, number)

	if slot[0] == "0" {
		slot = slot[1:]
	}

	return strings.Join(slot, "")
}
