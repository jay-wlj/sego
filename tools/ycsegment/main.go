package main

import (
	"flag"
	"fmt"
	"strings"
	"os"
	"io"
	"bufio"

	"github.com/huichen/sego"
)

func main() {
	flag.Parse()

	var seg sego.Segmenter
	dictReader := strings.NewReader(dict)
	seg.LoadDictionaryFromReader(dictReader)

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		segments := seg.Segment(line)
		tokens := sego.SegmentsToTokens(segments, true)
		for _, t := range tokens {
			fmt.Printf("%v 1\n", t.Text())
		}
	}
}