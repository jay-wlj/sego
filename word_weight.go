package sego

import "sort"

type WordWeight struct {
	Word   string
	Weight float64
}
type WordWeights []WordWeight

func (p WordWeights) Len() int { return len(p) }
func (p WordWeights) Less(i, j int) bool {
	return p[i].Weight > p[j].Weight
}
func (p WordWeights) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (x *Segmenter) Extract(s string, topk int) []string {
	words := x.ExtractWithWeight(s, topk)
	texts := []string{}
	for _, v := range words {
		texts = append(texts, v.Word)
	}
	return texts
}

func (x *Segmenter) ExtractWithWeight(s string, topk int) []WordWeight {
	segments := x.Segment([]byte(s))

	ms := make(map[string]float64)

	for _, v := range segments {
		freq := v.Token().Frequency()
		ms[v.Token().Text()] = freq
	}
	var words WordWeights
	for k, v := range ms {
		words = append(words, WordWeight{Word: k, Weight: v})
	}

	sort.Sort(words)
	pos := len(words)
	if pos > topk {
		pos = topk
	}
	return words[:pos]
}
