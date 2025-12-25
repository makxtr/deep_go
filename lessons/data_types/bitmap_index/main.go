package main

import (
	"fmt"
	"unsafe"
)

// 0000 0001 -> есть кальяны
// 0000 0010 -> можно с животными
// 0000 0100 -> есть виранда
// 0000 1000 -> есть алкоголь
// 0001 0000 -> есть живая музыка

func searchRestaurants(pattern int8, bitmaps []int8) []int {
	var indexes []int
	for idx, bitmap := range bitmaps {
		if bitmap^pattern == 0 {
			indexes = append(indexes, idx)
		}
	}

	return indexes
}

func main() {
	restaurants := []int8{
		0b00001101,
		0b00000010,
		0b00010000,
		0b00011111,
		0b00001001,
	}

	n := 0x01020304 // 00001010 >> 00000101
	// 00001010 << 00010100

	fmt.Printf("%0x", byte(n>>8))

	return

	pattern := int8(0b00011000)
	indexes := searchRestaurants(pattern, restaurants)
	_ = indexes

	//fmt.Printf("%x", ToLittleEndian[uint32](0x01020304))
}

func ToLittleEndian[T uint8 | uint16 | uint32 | uint64](number T) T {
	split := [4]byte{
		byte(number >> 24), // Старший байт (MSB)
		byte(number >> 16),
		byte(number >> 8),
		byte(number), // Младший байт (LSB)
	}

	for i, j := 0, len(split)-1; i < j; i, j = i+1, j-1 {
		split[i], split[j] = split[j], split[i]
	}

	result := uint32(split[0])<<24 | // байт 0 → сдвинуть на 24 бита влево
		uint32(split[1])<<16 | // байт 1 → сдвинуть на 16 бит влево
		uint32(split[2])<<8 | // байт 2 → сдвинуть на 8 бит влево
		uint32(split[3])

	return T(result)
}

func UnsafeSplit(n uint32) [4]byte {
	// Берем адрес числа, приводим его к указателю на массив из 4 байт и разыменовываем
	return *(*[4]byte)(unsafe.Pointer(&n))
}
