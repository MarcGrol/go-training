package main

import "fmt"

func main() {
	vowels := Vowels{"a", "e", "i"}
	fmt.Println(vowels.clone())

	vowelPronunciation := VowelPronunciation{
		"a": "eɪ",
		"e": "iː",
		"i": "aɪ",
	}

	fmt.Println(vowelPronunciation.clone())

	fmt.Println(cloneAny(vowels))
	fmt.Println(cloneAny(vowelPronunciation))

}

type Vowel string
type Vowels []Vowel

func (v Vowels) clone() Vowels {
	res := make(Vowels, len(v))
	copy(res, v)
	return res
}

type VowelPronunciation map[Vowel]string

func (vp VowelPronunciation) clone() VowelPronunciation {
	res := make(VowelPronunciation, 0)
	for k, v := range vp {
		res[k] = v
	}
	return res
}

type Cloner[T any] interface {
	clone() T
}

func cloneAny[T Cloner[T]](c T) T {
	return c.clone()
}
