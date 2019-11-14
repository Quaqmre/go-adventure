package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type pos struct {
	x, y float32
}

type color struct {
	r, g, b byte
}

type ball struct {
	pos
	xv, yv float32
	radius int
	color
}
type paddle struct {
	pos
	w, h int
	color
}

func setPixel(x, y int, c color, pixels []byte) {
	index := (x + 800*y) * 4
	if index < len(pixels) {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}
}
func clear(pixels []byte) {
	for i := 0; i < 800*600*4; i++ {
		if i%4 == 0 {
			pixels[i] = 0
			pixels[i+1] = 0
			pixels[i+2] = 0
		}
	}
}
func drawCircle(top ball, colr color, pixel []byte) {
	for y := -top.radius; y < top.radius; y++ {
		for x := -top.radius; x < top.radius; x++ {
			if x*x+y*y < top.radius*top.radius {
				setPixel(int(top.x+float32(x)), int(top.y+float32(y)), colr, pixel)
			}

		}
	}
}
func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.Quit()
	window, err := sdl.CreateWindow("myGame", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, 800, 600)
	if err != nil {
		fmt.Println(err)
	}
	defer tex.Destroy()

	top := ball{radius: 100}
	top.x = 200
	top.y = 200

	colr := color{255, 255, 0}
	pixel := make([]byte, 800*600*4)
	drawCircle(top, colr, pixel)

	tex.Update(nil, pixel, 800*4)
	renderer.Copy(tex, nil, nil)
	renderer.Present()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.MouseButtonEvent:
				clear(pixel)
				top.x += 10
				drawCircle(top, colr, pixel)
				tex.Update(nil, pixel, 800*4)
				renderer.Copy(tex, nil, nil)
				renderer.Present()
			}
		}
		sdl.Delay(16)
	}

}
