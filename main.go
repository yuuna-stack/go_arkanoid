package main

import (
	"math/rand"
	"path"
	"runtime"
	"time"

	"github.com/yuuna-stack/go_arkanoid/wrapper"

	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

const resourcesDir = "images"

type Point struct {
	x int
	y int
}

func init() { runtime.LockOSThread() }

func fullname(filename string) string {
	return path.Join(resourcesDir, filename)
}

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	resources := wrapper.Resources{}

	const gameWidth = 520
	const gameHeight = 450

	option := uint(window.SfResize | window.SfClose)
	wnd := wrapper.CreateWindow(gameWidth, gameHeight, "Arkanoid!", option, 60)

	sBackground, err := wrapper.FileToSprite(fullname("background.jpg"), &resources)
	if err != nil {
		panic("Couldn't load background.jpg")
	}
	sBall, err := wrapper.FileToSprite(fullname("ball.png"), &resources)
	if err != nil {
		panic("Couldn't load ball.png")
	}
	sPaddle, err := wrapper.FileToSprite(fullname("paddle.png"), &resources)
	if err != nil {
		panic("Couldn't load paddle.png")
	}
	sPaddle.SetPosition(300, 440)

	var block [1000]*wrapper.Sprite

	n := 0
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			blockPiece, err := wrapper.FileToSprite(fullname("block01.png"), &resources)
			if err != nil {
				panic("Couldn't load block01.png")
			}
			block[n] = blockPiece
			block[n].SetPosition(float32(i*43), float32(j*20))
			n++
		}
	}

	var dx float32 = 6.0
	var dy float32 = 5.0
	var x float32 = 300.0
	var y float32 = 300.0

	for wnd.IsOpen() {
		for wnd.Poll_Event() {
			if wnd.Close_Window() {
				return
			}
			if wnd.Key_Pressed() {
				if wnd.Key_Is(window.SfKeyLeft) {
					sPaddle.Move(-6, 0)
				} else if wnd.Key_Is(window.SfKeyRight) {
					sPaddle.Move(6, 0)
				}
			}
		}

		x += dx
		y += dy
		for i := 0; i < n; i++ {
			if block[i].IntersectRect(int(x+3+3), int(y), 6, 6) {
				block[i].SetPosition(-100, 0)
				dx = -dx
				dy = -dy
			}
		}

		if x < 0 || x > 520 {
			dx = -dx
		}
		if y < 0 || y > 450 {
			dy = -dy
		}

		if sPaddle.IntersectRect(int(x), int(y), 12, 12) {
			dy = -(float32(r1.Int()%5 + 2))
		}

		sBall.SetPosition(x, y)

		wnd.Clear_Window(graphics.GetSfBlack())

		sBackground.Draw(wnd.Get_Window())

		sBall.Draw(wnd.Get_Window())

		sPaddle.Draw(wnd.Get_Window())

		for i := 0; i < n; i++ {
			block[i].Draw(wnd.Get_Window())
		}

		graphics.SfRenderWindow_display(wnd.Get_Window())
	}

	resources.Clear()
	wnd.Clear()
}
