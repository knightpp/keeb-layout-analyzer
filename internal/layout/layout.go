package layout

import (
	"errors"
	"fmt"
	"image"
	"math"
	"slices"
	"strings"
	"unicode"

	"github.com/knightpp/keeb-layout-analyzer/internal/distance"
	"gonum.org/v1/gonum/stat/combin"
)

const (
	_ Finger = iota
	ThumbFinger
	IndexFinger
	MiddleFinger
	RingFinger
	PinkyFinger
)

type Finger byte

func (f Finger) String() string {
	switch f {
	case MiddleFinger:
		return "middle"
	case IndexFinger:
		return "index"
	case ThumbFinger:
		return "thumb"
	case RingFinger:
		return "ring"
	case PinkyFinger:
		return "pinky"
	default:
		return "unknown"
	}
}

const (
	LeftSide  Side = false
	RightSide Side = true
)

type Side bool

type Key struct {
	ID     string
	Char   rune
	Pos    Vec2
	Finger Finger
	Side   Side
	IsHome bool
	// Activation is a list of key ids required to press this key
	Activation []string
}

func (k Key) String() string {
	return k.ID
}

type Layout struct {
	Keys  []Key
	State *State

	charToKey       map[rune]Key
	idToKey         map[string]Key
	fingerToHomeKey map[Finger]Key
}

func New(keys []Key, state *State) *Layout {
	keys = slices.Clone(keys)

	for i, key := range keys {
		if key.ID == "" {
			keys[i].ID = string(key.Char)
		}
	}

	charToKey := make(map[rune]Key, len(keys))
	for _, key := range keys {
		charToKey[key.Char] = key
	}

	idToKey := make(map[string]Key, len(keys))
	for _, key := range keys {
		idToKey[key.ID] = key
	}

	fingerToHomeKey := make(map[Finger]Key, 10)
	for _, key := range keys {
		if !key.IsHome {
			continue
		}

		fingerToHomeKey[key.Finger] = key
	}

	return &Layout{
		Keys:            keys,
		State:           state,
		charToKey:       charToKey,
		idToKey:         idToKey,
		fingerToHomeKey: fingerToHomeKey,
	}
}

func Permutations(keys []Key, state *State, text string) []*Layout {
	keys = slices.Clone(keys)
	layouts := make([]*Layout, 0, len(keys)*len(keys))
	const k = 2
	gen := combin.NewCombinationGenerator(len(keys), k)
	buf := make([]int, k)
	for gen.Next() {
		cmb := gen.Combination(buf)
		swap(keys, cmb[0], cmb[1])

		layouts = append(layouts, New(keys, state.Clone()))
	}

	return layouts
}

func (l *Layout) Analyze(text string) (Stats, error) {
	for _, finger := range []Finger{
		ThumbFinger, IndexFinger, MiddleFinger, RingFinger, PinkyFinger,
	} {
		l.State.Move(l.fingerToHomeKey[finger])
	}

	var (
		stats Stats
		errs  []error
	)
	for _, ch := range text {
		k, ok := l.charToKey[ch]
		if !ok {
			errs = append(errs, fmt.Errorf("unknown char: %c", ch))
			continue
		}

		for _, act := range k.Activation {
			k, ok := l.idToKey[act]
			if !ok {
				errs = append(errs, fmt.Errorf("id not found: %q", act))
				continue
			}

			stats.TotalDistance += l.State.Move(k).Len()
			stats.record(k)
		}

		stats.TotalDistance += l.State.Move(k).Len()
		stats.record(k)

		// back to home key
		stats.TotalDistance += l.State.Move(l.fingerToHomeKey[k.Finger]).Len()
	}

	return stats, errors.Join(errs...)
}

type Stats struct {
	TotalDistance float64
	FingerUsage   struct {
		Left  FingerUsage
		Right FingerUsage
	}
}

func (s Stats) String() string {
	var buf strings.Builder

	fmt.Fprintf(&buf, "total distance=%s\n", distance.Distance(s.TotalDistance))
	fmt.Fprintf(&buf, "finger usage left hand\n\t%s\n", s.FingerUsage.Left)
	fmt.Fprintf(&buf, "finger usage right hand\n\t%s\n", s.FingerUsage.Right)

	return buf.String()
}

type FingerUsage struct {
	Thumb  uint
	Index  uint
	Middle uint
	Ring   uint
	Pinky  uint
}

func (fu FingerUsage) String() string {
	return fmt.Sprintf(
		"thumb=%d index=%d middle=%d ring=%d pinky=%d",
		fu.Thumb, fu.Index, fu.Middle, fu.Ring, fu.Pinky,
	)
}

func (s *Stats) record(k Key) {
	var fu *FingerUsage
	if k.Side == LeftSide {
		fu = &s.FingerUsage.Left
	} else {
		fu = &s.FingerUsage.Right
	}

	switch k.Finger {
	case IndexFinger:
		fu.Index++
	case ThumbFinger:
		fu.Thumb++
	case MiddleFinger:
		fu.Middle++
	case RingFinger:
		fu.Ring++
	case PinkyFinger:
		fu.Pinky++
	}
}

type Vec2 struct {
	X int
	Y int
}

func (v Vec2) Sub(rhs Vec2) Vec2 {
	return Vec2{
		X: v.X - rhs.X,
		Y: v.Y - rhs.Y,
	}
}

func (v Vec2) Len() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func (v Vec2) Point() image.Point {
	return image.Point{
		X: v.X,
		Y: v.Y,
	}
}

type State struct {
	Left  *Hand
	Right *Hand
}

func (s *State) Move(k Key) Vec2 {
	if k.Side {
		return s.Right.Move(k.Finger, k.Pos)
	} else {
		return s.Left.Move(k.Finger, k.Pos)
	}
}

func (s *State) Clone() *State {
	l := *s.Left
	r := *s.Right
	return &State{
		Left:  &l,
		Right: &r,
	}
}

type Hand struct {
	Thumb  Vec2
	Index  Vec2
	Middle Vec2
	Ring   Vec2
	Pinky  Vec2
}

func (h *Hand) Move(finger Finger, pos Vec2) Vec2 {
	var diff Vec2
	switch finger {
	case IndexFinger:
		diff = h.Index.Sub(pos)
		h.Index = pos
	case ThumbFinger:
		diff = h.Thumb.Sub(pos)
		h.Thumb = pos
	case MiddleFinger:
		diff = h.Middle.Sub(pos)
		h.Middle = pos
	case RingFinger:
		diff = h.Ring.Sub(pos)
		h.Ring = pos
	case PinkyFinger:
		diff = h.Pinky.Sub(pos)
		h.Pinky = pos
	default:
		panic("todo")
	}
	return diff
}

const U1 = 20 * distance.MM

func Cluster(side Side, finger Finger, keyss ...[]Key) []Key {
	for i := range keyss {
		for j := range keyss[i] {
			keyss[i][j].Side = side
			keyss[i][j].Finger = finger
		}
	}

	return Flatten(keyss...)
}

func Column(keys ...Key) []Key {
	for i := range keys {
		keys[i].Pos.Y = i * int(U1)
	}
	return keys
}

func ActivationGroup(activation []string, keys ...Key) []Key {
	for i := range keys {
		keys[i].Activation = activation
	}
	return keys
}

func Flatten(keyss ...[]Key) []Key {
	var n int
	for _, keys := range keyss {
		n += len(keys)
	}

	flat := make([]Key, 0, n)
	for _, keys := range keyss {
		flat = append(flat, keys...)
	}
	return flat
}

func AddUppercase(activation []string, keys []Key) []Key {
	keys = slices.Clone(keys)

	for _, key := range keys {
		fmt.Println(string(key.Char))
		if key.Char == 0 ||
			!unicode.IsLetter(key.Char) ||
			unicode.IsUpper(key.Char) {
			continue
		}

		key.ID = ""
		key.Activation = activation
		key.Char = unicode.ToUpper(key.Char)
		key.IsHome = false

		keys = append(keys, key)
	}

	return keys
}

func swap(keys []Key, i, j int) {
	ith := keys[i]
	jth := keys[j]

	keys[i].Char = jth.Char
	keys[i].ID = jth.ID
	keys[j].Char = ith.Char
	keys[j].ID = ith.ID
}
