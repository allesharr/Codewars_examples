package morse

import (
	"log"
	"regexp"
)

type Duration string

// Morse code signal durations
const (
	Dit Duration = "•" // short
	Dah Duration = "−" // long
)

// const (
// 	hz  = 800
// 	wpm = 10

// 	durationShort = 1200 / wpm
// 	durationLong  = durationShort * 3
// 	durationGap   = durationShort * 2
// )

type Code string

const (
	// alphabets
	A Code = Code(Dit + Dah)
	B Code = Code(Dah + Dit + Dit + Dit)
	C Code = Code(Dah + Dit + Dah + Dit)
	D Code = Code(Dah + Dit + Dit)
	E Code = Code(Dit)
	F Code = Code(Dit + Dit + Dah + Dit)
	G Code = Code(Dah + Dah + Dit)
	H Code = Code(Dit + Dit + Dit + Dit)
	I Code = Code(Dit + Dit)
	J Code = Code(Dit + Dah + Dah + Dah)
	K Code = Code(Dah + Dit + Dah)
	L Code = Code(Dit + Dah + Dit + Dit)
	M Code = Code(Dah + Dah)
	N Code = Code(Dah + Dit)
	O Code = Code(Dah + Dah + Dah)
	P Code = Code(Dit + Dah + Dah + Dit)
	Q Code = Code(Dah + Dah + Dit + Dah)
	R Code = Code(Dit + Dah + Dit)
	S Code = Code(Dit + Dit + Dit)
	T Code = Code(Dah)
	U Code = Code(Dit + Dit + Dah)
	V Code = Code(Dit + Dit + Dit + Dah)
	W Code = Code(Dit + Dah + Dah)
	X Code = Code(Dah + Dit + Dit + Dah)
	Y Code = Code(Dah + Dit + Dah + Dah)
	Z Code = Code(Dah + Dah + Dit + Dit)

	// numbers
	One   Code = Code(Dit + Dah + Dah + Dah + Dah)
	Two   Code = Code(Dit + Dit + Dah + Dah + Dah)
	Three Code = Code(Dit + Dit + Dit + Dah + Dah)
	Four  Code = Code(Dit + Dit + Dit + Dit + Dah)
	Five  Code = Code(Dit + Dit + Dit + Dit + Dit)
	Six   Code = Code(Dah + Dit + Dit + Dit + Dit)
	Seven Code = Code(Dah + Dah + Dit + Dit + Dit)
	Eight Code = Code(Dah + Dah + Dah + Dit + Dit)
	Nine  Code = Code(Dah + Dah + Dah + Dah + Dit)
	Zero  Code = Code(Dah + Dah + Dah + Dah + Dah)

	Space Code = " "
	None  Code = ""
)

// map for codes and characters
var charsMap map[Code]rune

// regular expression for non-encodable strings
var regexToEscape *regexp.Regexp
var regexRedundantSpaces *regexp.Regexp

var codesMap = map[rune]Code{
	'a': A,
	'b': B,
	'c': C,
	'd': D,
	'e': E,
	'f': F,
	'g': G,
	'h': H,
	'i': I,
	'j': J,
	'k': K,
	'l': L,
	'm': M,
	'n': N,
	'o': O,
	'p': P,
	'q': Q,
	'r': R,
	's': S,
	't': T,
	'u': U,
	'v': V,
	'w': W,
	'x': X,
	'y': Y,
	'z': Z,

	'1': One,
	'2': Two,
	'3': Three,
	'4': Four,
	'5': Five,
	'6': Six,
	'7': Seven,
	'8': Eight,
	'9': Nine,
	'0': Zero,

	' ': Space,
}

func getMoorseByString(text string) string {
	var moorse string
	var codeArray, err = Encode(text)
	if err != nil {
		log.Fatal("no such encode")
	}
	for _, element := range codeArray {
		moorse = moorse + string(element) + " "
	}
	return moorse
}

func getStringByMoorse(string) string {
	var str string
	return str
}

func init() {

	// regexToEscape, err := regexp.Compile("[^a-zA-Z0-9]+")
	// if err != nil {
	// 	log.Fatal("Cannot resolve RegexToEscape")
	// }

	// regexRedundantSpaces, err = regexp.Compile(" ")

	// if err != nil {
	// 	log.Fatal("Cannot resolve RedudantSpaces")
	// }
	// codes' map

	// characters' map
	charsMap = make(map[Code]rune)
	for k, v := range codesMap {
		charsMap[v] = k
	}
}

func Encode(text string) (codes []Code, err error) {
	codes = []Code{}

	err = nil
	for _, element := range text {
		codes = append(codes, codesMap[element])
	}

	return codes, err

}
func Decode(text string) (codes []Code, err error) {

}
