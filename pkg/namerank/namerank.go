package namerank

import (
	"fmt"
	"log"
	"github.com/AntoineAugusti/wordsegmentation/corpus"
	"github.com/AntoineAugusti/wordsegmentation"
)

func TestSegment() {
	//expected := []string{"what", "is", "the", "weather", "like", "today"}
	englishCorpus := corpus.NewEnglishCorpus()
	segment, score := wordsegmentation.Segment(englishCorpus, "WhatIsTheWeatherliketoday? ")
	fmt.Printf("Score is %v and segment is %v", score, segment)
	//assert.Equal(t, segment, expected)
}

func SegmentAndRank(name string, englishCorpus wordsegmentation.Corpus) float64 {
	seg, score := wordsegmentation.Segment(englishCorpus, name)
	log.Printf("Seg of %s is %v and score is ", name, seg, score)
	return score
}
