/*
Installation of SDL with Go
sudo apt-get install libsdl2-dev
sudo go get -v github.com/veandco/go-sdl2/sdl
sudo apt-get install libsdl2-ttf-dev
sudo go get -v github.com/veandco/go-sdl2/ttf
*/
package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"log"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

func main() {
	var (
		running  bool
		renderer *sdl.Renderer
		err      error
		window   *sdl.Window
		textFont *ttf.Font
	)

	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		log.Fatal(err)
	}
	defer renderer.Destroy()

	err = ttf.Init()
	if err != nil {
		log.Printf("Failed to initialize TTF: %s\n", err)
	}

	textFont, err = ttf.OpenFont("police.ttf", 80)
	if err != nil {
		log.Printf("Error opening font: %v", err)
	}

	running = true
	for running {
		running = loop(running)

		sometingToScreen(renderer, textFont)
		sdl.Delay(16)
	}

	textFont.Close()
}

func loop(running bool) bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch re := event.(type) {
		case *sdl.QuitEvent:
			log.Println("Quit with window event")
			running = false
			break
		case *sdl.KeyboardEvent:
			if re.Type == sdl.KEYDOWN {
				keyName := re.Keysym.Sym
				switch keyName {
				case sdl.K_ESCAPE:
					log.Println("Quit with ESC keyboard event")
					running = false
					break
				}
			} else if re.Type == sdl.KEYUP {
				// Release the key
			}
		}
	}
	return running
}

func sometingToScreen(renderer *sdl.Renderer, textFont *ttf.Font) {
	var points []sdl.Point
	var rect sdl.Rect
	var rects []sdl.Rect

	renderer.SetDrawColor(1, 1, 0, 255)
	renderer.Clear()

	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.DrawPoint(150, 300)

	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.DrawLine(0, 0, 200, 200)

	points = []sdl.Point{{0, 0}, {100, 300}, {100, 300}, {200, 0}}
	renderer.SetDrawColor(255, 255, 0, 255)
	renderer.DrawLines(points)

	rect = sdl.Rect{300, 0, 200, 200}
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.DrawRect(&rect)

	rects = []sdl.Rect{{400, 400, 100, 100}, {550, 350, 200, 200}}
	renderer.SetDrawColor(0, 255, 255, 255)
	renderer.DrawRects(rects)

	rect = sdl.Rect{250, 250, 200, 200}
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.FillRect(&rect)
		
	rects = []sdl.Rect{{500, 300, 100, 100}, {200, 300, 200, 200}}
	renderer.SetDrawColor(255, 0, 255, 255)
	renderer.FillRects(rects)

	textToRenderer("bonjour les amis", textFont, renderer)

	renderer.Present()
}

func textToRenderer(text string, textFont *ttf.Font, renderer *sdl.Renderer) error {
	surface, err := textFont.RenderUTF8Solid(text, sdl.Color{R:255, G:165, B:0, A:255})
		
	defer surface.Free()
	if err != nil {
		//fmt.Errorf("Error creating score surface: %v", err)
		return err
	}

	texture, err := renderer.CreateTextureFromSurface(surface)
	defer texture.Destroy()

	surfaceRect := &sdl.Rect{X: 100, Y: 80, W: 300, H: 60}

	renderer.Copy(texture, nil, surfaceRect)
	if err != nil {
		//fmt.Errorf("could not copy score texture: %v", err)
		return err
	}
	return nil
}
