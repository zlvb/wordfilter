package wordfilter

import (
	"unicode/utf8"
	"xgame/common/xlog"
)

type NextWord struct {
	nextMap map[rune]*NextWord
}

func NewNextWord() *NextWord {
	return &NextWord{
		nextMap:make(map[rune]*NextWord),
	}
}

var (
	wordsLib = NewNextWord()
)

func FilterDirtyWord(input string) (bool, string) {
	var filtered bool
	nextMap := wordsLib.nextMap
	out := []rune(input)
	var output []rune
	startIdx := 0
	for i := 0; startIdx+i < len(out); i++ {
		ch := out[i+startIdx]
		it, ok := nextMap[ch]
		if ok {
			xlog.Debug("hit {} {}", ch, startIdx)
			nextMap = it.nextMap
			_, ok2 := nextMap[0]
			if ok2 {
				filtered = true
				for j := 0; j < i+1; j++ {
					xlog.Debug("{} -> {}", out[startIdx+j], "*")
					out[startIdx+j] = '*'
				}
				nextMap = wordsLib.nextMap
				startIdx = i + 1 + startIdx
				i = -1
			}
		} else {
			xlog.Debug("jump {} {}", ch, i+startIdx)
			nextMap = wordsLib.nextMap
			startIdx++
			i = -1
		}
	}

	xlog.Debug("finished")
	if filtered {
		for _, ch := range out {
			output = append(output, ch)
			xlog.Debug("{}", ch)
		}
	}

	return filtered, string(output)
}

func AddWord(w string) int {
	nextMap := wordsLib.nextMap
	count := utf8.RuneCountInString(w)
	i := 0
	for _, ch := range w {
		it, ok := nextMap[ch]
		if ok {
			nextMap = it.nextMap
		} else {
			next := NewNextWord()
			nextMap[ch] = next
			nextMap = next.nextMap
		}

		if i == count - 1 {
			nextMap[0] = nil
		}
		i++
	}

	return 0
}
