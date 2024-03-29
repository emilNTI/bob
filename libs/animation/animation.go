package animationManager

import (
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type AnimatedImage struct {
	frames        []*ebiten.Image
	current_frame uint
	playing_speed uint
	width         uint
	height        uint
	is_playing    bool
}

func (a *AnimatedImage) Init(path string, playing_speed uint, width uint, height uint) {
	a.playing_speed = playing_speed
	a.width = width
	a.height = height
	// load image horizontal
	t, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal("Error loading animated image")
	}
	for i := 0; i < t.Bounds().Dx()/int(a.width); i++ {
		a.frames = append(a.frames, t.SubImage(image.Rect(
			i*int(width), 0, i*int(width)+int(width), 16,
		)).(*ebiten.Image))
	}
}

func (a *AnimatedImage) SetFrame(frame uint) {
	a.current_frame = frame
}

func (a *AnimatedImage) GetImage() *ebiten.Image {
	return a.frames[a.current_frame]
}

func (a *AnimatedImage) SetPlaying(is_playing bool){
	a.is_playing = is_playing
}

func (a *AnimatedImage) PlayLoop() {
	for {
		if !a.is_playing{continue}
		if int(a.current_frame) < len(a.frames)-1 {
			a.current_frame++
		} else {
			a.current_frame = 0
		}
		time.Sleep(time.Duration(a.playing_speed) * time.Millisecond)
	}
}
