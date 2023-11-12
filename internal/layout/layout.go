package layout

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/knightpp/keeb-layout-analyzer/internal/distance"
)

const (
	MiddleFinger Finger = iota + 1
	IndexFinger
	ThumbFinger
	RingFinger
	PinkyFinger
)

type Finger byte

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
	// Activation is a list of key ids required to press this key
	Activation []string
}

func (k Key) String() string {
	return k.ID
}

type Layout struct {
	Keys  []Key
	State *State

	charToKey map[rune]Key
	idToKey   map[string]Key
}

func New(keys []Key, state *State) *Layout {
	charToKey := make(map[rune]Key, len(keys))
	for _, key := range keys {
		charToKey[key.Char] = key
	}

	idToKey := make(map[string]Key, len(keys))
	for _, key := range keys {
		idToKey[key.ID] = key
	}

	return &Layout{
		Keys:      keys,
		State:     state,
		charToKey: charToKey,
		idToKey:   idToKey,
	}
}

func (l *Layout) Analyze(text string) (Stats, error) {
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
