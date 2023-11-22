package visual_test

import (
	"bufio"
	"image/jpeg"
	"os"
	"testing"

	"github.com/knightpp/keeb-layout-analyzer/internal/layout"
	"github.com/knightpp/keeb-layout-analyzer/internal/visual"
)

func TestToImage(t *testing.T) {
	// t.Skip()

	testCases := []struct {
		name string
		keys []layout.Key
	}{
		{
			name: "colemak",
			keys: layout.ColemakDH34Keys,
		},
		{
			name: "ukrainian йцукен",
			keys: layout.Ukrainian34,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			img := visual.ToImage(tc.keys, 20)

			file, err := os.Create(tc.name + ".jpeg")
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
		})
	}
}
