package main

import (
	"encoding/binary"
	"fmt"
	"time"
)

const (
	encoding = "0123456789abcdefghijklmnopqrstuv"
)

func main() {
	var decoder [256]byte
	for i := 0; i < 32; i++ {
		decoder[encoding[i]] = byte(i)
	}

	timeMs1 := time.Now()
	t := timeMs1.Unix() // get time tick
	intTime := uint32(t)
	var timeByte [4]byte
	var encoded [8]byte
	binary.BigEndian.PutUint32(timeByte[:], intTime) // Get bigEndian to Time tick and use it encode
	encoded[7] = encoding[timeByte[3]&31]            // there is encode time and use some time slice
	encoded[6] = encoding[(timeByte[2]>>1)&31]
	encoded[5] = encoding[(timeByte[2]<<4)&31]
	encoded[4] = encoding[(timeByte[2] >> 6)]
	encoded[3] = encoding[timeByte[1]&31]
	encoded[2] = encoding[(timeByte[1]>>3)&31]
	encoded[1] = encoding[timeByte[0]<<1&31]
	encoded[0] = encoding[timeByte[0]>>4]
	s := string(encoded[:]) // s have a encoded time
	fmt.Println(s)
	fmt.Println(timeByte[2:3])
	var gggg byte
	gggg = decoder[encoded[4]]<<6 | decoder[encoded[5]]>>4 | decoder[encoded[6]]<<1 // just de code to timeByte 2. element
	fmt.Println("encoded=", gggg)

	_, _ = t, intTime
}
