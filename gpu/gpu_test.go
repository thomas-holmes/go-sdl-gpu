package gpu

import (
	"log"
	"testing"
	"time"
)

func TestRender(t *testing.T) {
	target := Init(800, 600, 0)

	log.Printf("%T %+v", target, *target)

	color := Color{R: 255, G: 0, B: 255, A: 255}
	doneTime := time.Now().Add(2 * time.Second)
	for time.Now().Before(doneTime) {
		target.Clear()
		target.RectangleFilled(100, 100, 400, 400, color)
		target.Flip()
	}
}
