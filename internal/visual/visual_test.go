package visual_test

import (
	"bufio"
	"image/jpeg"
	"os"
	"testing"

	"github.com/knightpp/keeb-layout-analyzer/internal/layout"
	"github.com/knightpp/keeb-layout-analyzer/internal/visual"
)

func TestColemak(t *testing.T) {
	img := visual.ToImage(layout.ColemakDH34Keys, 20)

	file, err := os.Create("colemak-dh-34-keys.jpeg")
	if err != nil {
		t.Fatal(err)
		return
	}
	defer file.Close()

	buf := bufio.NewWriter(file)
	defer buf.Flush()

	err = jpeg.Encode(buf, img, &jpeg.Options{
		Quality: 90,
	})
	if err != nil {
		t.Fatal(err)
		return
	}
}
