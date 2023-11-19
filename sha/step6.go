package main

import (
	"strconv"
)

var a uint32 = h0
var b uint32 = h1
var c uint32 = h2
var d uint32 = h3
var e uint32 = h4
var f uint32 = h5
var g uint32 = h6
var h uint32 = h7

func compression() {
	for i := 0; i < 64; i++ {
		wuint32, _ := strconv.ParseUint(w[i], 2, 32)

		s1 := rightRotate(e, 6) ^ rightRotate(e, 11) ^ rightRotate(e, 25)
		ch := (e & f) ^ (^e & g)
		temp1 := h + s1 + ch + k[i] + uint32(wuint32)
		s0 := rightRotate(a, 2) ^ rightRotate(a, 13) ^ rightRotate(a, 22)
		maj := (a & b) ^ (a & c) ^ (b & c)
		temp2 := s0 + maj
		h = g
		g = f
		f = e
		e = d + temp1
		d = c
		c = b
		b = a
		a = temp1 + temp2
	}
}
