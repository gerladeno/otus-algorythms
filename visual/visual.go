package visual

import (
	"algorythms/sortingcommon"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"sync"
)

var (
	events []Event
	mx     sync.Mutex
)

const (
	width  = 1920
	height = 1080
)

type Event struct {
	A     []int
	Swaps []Swap
}

type Swap struct {
	I int
	J int
}

func SwapEvent(s sortingcommon.Sorter, swaps ...Swap) {
	for _, swap := range swaps {
		sortingcommon.Swap(s, swap.I, swap.J)
	}
	mx.Lock()
	defer mx.Unlock()
	events = append(events, Event{
		A:     s.GetSlice(),
		Swaps: swaps,
	})
}

type Game struct {
	s     sortingcommon.Sorter
	next  chan struct{}
	pause bool
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.pause = !g.pause
	}
	if !g.pause {
		g.next <- struct{}{}
	}
	return nil
}

const (
	normal int8 = iota
	selected
	swapped
)

func (g *Game) Draw(screen *ebiten.Image) {
	a := g.s.GetSlice()
	var max int
	for _, elem := range a {
		if max < elem {
			max = elem
		}
	}
	l := len(a)
	for i, elem := range a {
		drawElement(screen, i, l, elem, max, normal)
	}
	mx.Lock()
	defer mx.Unlock()
	for _, e := range events {
		for _, swap := range e.Swaps {
			drawElement(screen, swap.I, l, a[swap.I], max, swapped)
			drawElement(screen, swap.J, l, a[swap.J], max, swapped)
		}
	}
	events = events[:0]
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func Run(eventCh chan struct{}, s sortingcommon.Sorter) error {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{next: eventCh, s: s}); err != nil {
		return err
	}
	return nil
}

func drawElement(dst *ebiten.Image, position, length, value, max int, state int8) {
	var clr color.Color = color.White
	switch state {
	case normal:
	case swapped:
		clr = color.RGBA{R: 245, G: 95, B: 95, A: 255}
	}
	h := 0.9 * float64(height*value) / float64(max)
	w := 0.9 * float64(width) / float64(length)
	x := width * float64(position) / float64(length)
	y := height - h - 50
	ebitenutil.DrawRect(dst, x, y, w, h, clr)
}
