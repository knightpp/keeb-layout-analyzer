package visual

import (
	"image"
	"image/color"
	"image/draw"
	"strings"

	"github.com/knightpp/keeb-layout-analyzer/internal/layout"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func ToImage(keys []layout.Key, keyCapSize int) image.Image {
	handToFingerToLayerToKey := make(map[layout.Side]map[layout.Finger]map[string][]layout.Key)
	handToFingerToLayerToKey[layout.LeftSide] = make(map[layout.Finger]map[string][]layout.Key)
	handToFingerToLayerToKey[layout.RightSide] = make(map[layout.Finger]map[string][]layout.Key)
	for _, key := range keys {
		hand := handToFingerToLayerToKey[key.Side]
		fingers, ok := hand[key.Finger]
		if !ok {
			fingers = make(map[string][]layout.Key)
			hand[key.Finger] = fingers
		}

		layer := "base"
		if len(key.Activation) != 0 {
			layer = strings.Join(key.Activation, "+")
		}

		fingers[layer] = append(fingers[layer], key)
	}

	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.Draw(img, img.Bounds(), image.NewUniform(color.White), image.Point{}, draw.Src)

	ld := layoutDrawer{
		border:     newBorder(image.Rect(0, 0, keyCapSize, keyCapSize), 1),
		font:       basicfont.Face7x13,
		keycapSize: keyCapSize,
	}

	ld.drawHand(img, layout.LeftSide, handToFingerToLayerToKey[layout.LeftSide], image.Point{})
	ld.drawHand(img, layout.RightSide, handToFingerToLayerToKey[layout.RightSide], image.Point{X: 300})

	return img
}

type layoutDrawer struct {
	border     image.Image
	font       font.Face
	keycapSize int
}

func (ld *layoutDrawer) drawHand(
	img draw.Image,
	side layout.Side,
	fingerToLayer map[layout.Finger]map[string][]layout.Key,
	pos image.Point,
) {
	fingers := maps.Keys(fingerToLayer)
	slices.Sort(fingers)
	offset := pos
	for _, finger := range fingers {
		layerToKeys := fingerToLayer[finger]
		ld.drawFinger(img, finger, layerToKeys, offset)

		offset.Y += 5 * ld.keycapSize
	}
}

func (ld *layoutDrawer) drawFinger(
	img draw.Image,
	finger layout.Finger,
	layerToKeys map[string][]layout.Key,
	pos image.Point,
) {
	ld.drawText(img, finger.String(), pos)
	pos.Y += ld.keycapSize

	layers := maps.Keys(layerToKeys)
	slices.Sort(layers)

	for _, layer := range layers {
		keys := layerToKeys[layer]

		ld.drawLayer(img, layer, keys, pos)
		pos.X += 3 * ld.keycapSize
	}
}

func (ld *layoutDrawer) drawLayer(
	img draw.Image,
	name string,
	keys []layout.Key,
	pos image.Point,
) {
	textColor := color.RGBA{B: 255, A: 255}
	fd := font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(textColor),
		Face: ld.font,
	}
	size, _ := fd.BoundString(name)
	fd.Dot = fixed.P(
		pos.X+5,
		pos.Y+size.Max.Y.Ceil()*5,
	)
	fd.DrawString(name)

	pos.Y += 6 * size.Max.Y.Ceil()

	green := color.RGBA{G: 255, A: 255}
	black := color.RGBA{A: 255}
	for _, key := range keys {
		color := black
		if key.IsHome {
			color = green
		}

		label := string(key.Char)
		if key.Char == 0 {
			label = key.ID
		}

		ld.drawKey(img, label, pos.Add(key.Pos.Point()), color)
	}
}

func (ld *layoutDrawer) drawKey(
	img draw.Image,
	label string,
	pos image.Point,
	textColor color.Color,
) {
	border := ld.border.Bounds()

	draw.Draw(img, border.Add(pos), ld.border, image.Point{}, draw.Src)

	fd := font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(textColor),
		Face: ld.font,
	}
	size, _ := fd.BoundString(label)
	fd.Dot = fixed.P(
		pos.X+border.Dx()/2-size.Max.X.Ceil()/2,
		pos.Y+border.Dy()/2+size.Max.Y.Ceil(),
	)

	fd.DrawString(string(label))
}

func (ld *layoutDrawer) drawText(img draw.Image, text string, pos image.Point) {
	textColor := color.RGBA{B: 255, A: 255}
	fd := font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(textColor),
		Face: ld.font,
	}
	size, _ := fd.BoundString(text)
	fd.Dot = fixed.P(
		pos.X+ld.keycapSize,
		pos.Y+size.Max.Y.Ceil()*5,
	)
	fd.DrawString(text)
}

var _ image.Image = border{}

func newBorder(bounds image.Rectangle, width int) image.Image {
	return border{
		bounds: bounds,
		colors: bounds.Inset(width),
	}
}

type border struct {
	bounds image.Rectangle
	colors image.Rectangle
}

// At implements image.Image.
func (b border) At(x int, y int) color.Color {
	return b.colors.At(x, y)
}

// Bounds implements image.Image.
func (b border) Bounds() image.Rectangle {
	return b.bounds
}

// ColorModel implements image.Image.
func (border) ColorModel() color.Model {
	return color.Gray16Model
}
