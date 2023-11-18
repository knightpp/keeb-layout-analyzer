package main

import (
	_ "embed"
	"fmt"

	"github.com/knightpp/keeb-layout-analyzer/internal/layout"
)

//go:embed data/english_common_words.txt
var engCommonWords string

func main() {
	l := layout.New(layout.ColemakDH34Keys, &layout.State{
		Left:  &layout.Hand{},
		Right: &layout.Hand{},
	})

	stats, err := l.Analyze(engCommonWords)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(stats)
}
