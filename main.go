package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"

	"github.com/knightpp/keeb-layout-analyzer/internal/layout"
)

var (
	//go:embed data/english_common_words.txt
	engCommonWords string

	//go:embed data/ukr.txt
	ukrainian string
)

var (
	fWords  string
	fLayout string
)

func init() {
	flag.StringVar(&fWords, "words", "ukr.txt", "specify path to a file or one of built-ins")
	flag.StringVar(&fLayout, "layout", "ukr34", "layout to use")
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var keys []layout.Key
	switch fLayout {
	case "ukr34":
		keys = layout.Ukrainian34
	case "colemakdh":
		keys = layout.ColemakDH34Keys
	default:
		return fmt.Errorf("unknown layout: %q", fLayout)
	}

	var words string
	switch fWords {
	case "ukr.txt":
		words = ukrainian
	case "english_common_words.txt":
		words = engCommonWords
	default:
		return fmt.Errorf("unimplemented")
	}

	l := layout.New(keys, &layout.State{
		Left:  &layout.Hand{},
		Right: &layout.Hand{},
	})

	stats, err := l.Analyze(words)
	if err != nil {
		fmt.Println(err)
		if errs, ok := err.(interface{ Unwrap() []error }); ok {
			fmt.Println("total errors:", len(errs.Unwrap()))
		}
	}

	fmt.Println(stats)
	return nil
}
