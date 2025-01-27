// Based on https://github.com/ozdemirburak/morse-decoder
package processors

import (
	"fmt"
	"strings"
)

var moreCodeMap = map[string]string{
	// Latin => https://en.wikipedia.org/wiki/Morse_code
	"A": "01", "B": "1000", "C": "1010", "D": "100", "E": "0", "F": "0010",
	"G": "110", "H": "0000", "I": "00", "J": "0111", "K": "101", "L": "0100",
	"M": "11", "N": "10", "O": "111", "P": "0110", "Q": "1101", "R": "010",
	"S": "000", "T": "1", "U": "001", "V": "0001", "W": "011", "X": "1001",
	"Y": "1011", "Z": "1100",
	// Numbers
	"0": "11111", "1": "01111", "2": "00111", "3": "00011", "4": "00001",
	"5": "00000", "6": "10000", "7": "11000", "8": "11100", "9": "11110",

	// Punctuation
	".": "010101", ",": "110011", "?": "001100", "'": "011110", "!": "101011", "/": "10010",
	"(": "10110", ")": "101101", "&": "01000", ":": "111000", ";": "101010", "=": "10001",
	"+": "01010", "-": "100001", "_": "001101", `"`: "010010", "$": "0001001", "@": "011010",
	"¿": "00101", "¡": "110001",

	// Latin Extended => https://ham.stackexchange.com/questions/1379/international-characters-in-morse-code
	"Ã": "01101", "Á": "01101", "Å": "01101", "À": "01101", "Â": "01101", "Ä": "0101",
	"Ą": "0101", "Æ": "0101", "Ç": "10100", "Ć": "10100", "Ĉ": "10100", "Č": "110",
	"Ð": "00110", "È": "01001", "Ę": "00100", "Ë": "00100", "É": "00100",
	"Ê": "10010", "Ğ": "11010", "Ĝ": "11010", "Ĥ": "1111", "İ": "01001", "Ï": "10011",
	"Ì": "01110", "Ĵ": "01110", "Ł": "01001", "Ń": "11011", "Ñ": "11011", "Ó": "1110",
	"Ò": "1110", "Ö": "1110", "Ô": "1110", "Ø": "1110", "Ś": "0001000", "Ş": "01100",
	"Ș": "1111", "Š": "1111", "Ŝ": "00010", "ß": "000000", "Þ": "01100", "Ü": "0011",
	"Ù": "0011", "Ŭ": "0011", "Ž": "11001", "Ź": "110010", "Ż": "11001",

	// Cyrillic Alphabet => https://en.wikipedia.org/wiki/Russian_Morse_code
	"А": "01", "Б": "1000", "В": "011", "Г": "110", "Д": "100", "Е": "0",
	"Ж": "0001", "З": "1100", "И": "00", "Й": "0111", "К": "101", "Л": "0100",
	"М": "11", "Н": "10", "О": "111", "П": "0110", "Р": "010", "С": "000",
	"Т": "1", "У": "001", "Ф": "0010", "Х": "0000", "Ц": "1010", "Ч": "1110",
	"Ш": "1111", "Щ": "1101", "Ъ": "11011", "Ы": "1011", "Ь": "1001", "Э": "00100",
	"Ю": "0011", "Я": "0101", "Ї": "01110", "Є": "00100", "І": "00", "Ґ": "110",

	// Greek Alphabet => https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets
	"Α": "01", "Β": "1000", "Γ": "110", "Δ": "100", "Ε": "0", "Ζ": "1100",
	"Η": "0000", "Θ": "1010", "Ι": "00", "Κ": "101", "Λ": "0100", "Μ": "11",
	"Ν": "10", "Ξ": "1001", "Ο": "111", "Π": "0110", "Ρ": "010", "Σ": "000",
	"Τ": "1", "Υ": "1011", "Φ": "0010", "Χ": "1111", "Ψ": "1101", "Ω": "011",

	// Hebrew Alphabet => https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets
	"א": "01", "ב": "1000", "ג": "110", "ד": "100", "ה": "111", "ו": "0",
	"ז": "1100", "ח": "0000", "ט": "001", "י": "00", "כ": "101", "ל": "0100",
	"מ": "11", "נ": "10", "ס": "1010", "ע": "0111", "פ": "0110", "צ": "011",
	"ק": "1101", "ר": "010", "ש": "000", "ת": "1",

	// Arabic Alphabet => https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets
	"ا": "01", "ب": "1000", "ت": "1", "ث": "1010", "ج": "0111", "ح": "0000",
	"خ": "111", "د": "100", "ذ": "1100", "ر": "010", "ز": "1110", "س": "000",
	"ش": "1111", "ص": "1001", "ض": "0001", "ط": "001", "ظ": "1011", "ع": "0101",
	"غ": "110", "ف": "0010", "ق": "1101", "ك": "101", "ل": "0100", "م": "11",
	"ن": "10", "ه": "00100", "و": "011", "ي": "00", "ﺀ": "0",

	// Persian Alphabet => https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets
	"پ": "0110", "چ": "1110", "ژ": "110", "ک": "101", "گ": "1101", "ی": "00",

	// Japanese Alphabet => https://ja.wikipedia.org/wiki/%E3%83%A2%E3%83%BC%E3%83%AB%E3%82%B9%E7%AC%A6%E5%8F%B7#%E5%92%8C%E6%96%87%E3%83%A2%E3%83%BC%E3%83%AB%E3%82%B9%E7%AC%A6%E5%8F%B7
	"ア": "11011", "カ": "0100", "サ": "10101", "タ": "10", "ナ": "010", "ハ": "1000",
	"マ": "1001", "ヤ": "011", "ラ": "000", "ワ": "101", "イ": "01", "キ": "10100",
	"シ": "11010", "チ": "0010", "ニ": "1010", "ヒ": "11001", "ミ": "00101", "リ": "110",
	"ヰ": "01001", "ウ": "001", "ク": "0001", "ス": "11101", "ツ": "0110", "ヌ": "0000",
	"フ": "1100", "ム": "1", "ユ": "10011", "ル": "10110", "ン": "01010", "エ": "10111",
	"ケ": "1011", "セ": "01110", "テ": "01011", "ネ": "1101", "ヘ": "0", "メ": "10001",
	"レ": "111", "ヱ": "01100", "オ": "01000", "コ": "1111", "ソ": "1110", "ト": "00100",
	"ノ": "0011", "ホ": "100", "モ": "10010", "ヨ": "11", "ロ": "0101", "ヲ": "0111",
	"゛": "00", "゜": "00110", "。": "010100", "ー": "01101", "、": "010101",
	"（": "101101", "）": "010010",

	// Korean Alphabet => https://en.wikipedia.org/wiki/SKATS
	"ㄱ": "0100", "ㄴ": "0010", "ㄷ": "1000", "ㄹ": "0001", "ㅁ": "11", "ㅂ": "011",
	"ㅅ": "110", "ㅇ": "101", "ㅈ": "0110", "ㅊ": "1010", "ㅋ": "1001", "ㅌ": "1100",
	"ㅍ": "111", "ㅎ": "0111", "ㅏ": "0", "ㅑ": "00", "ㅓ": "1", "ㅕ": "000",
	"ㅗ": "01", "ㅛ": "10", "ㅜ": "0000", "ㅠ": "010", "ㅡ": "100", "ㅣ": "001",

	// Thai Alphabet => https://th.wikipedia.org/wiki/รหัสมอร์ส
	"ก": "110", "ข": "1010", "ค": "101", "ง": "10110", "จ": "10010",
	"ฉ": "1111", "ช": "1001", "ซ": "1100", "ญ": "0111", "ด": "100",
	"ต": "1", "ถ": "10100", "ท": "10011", "น": "10", "บ": "1000",
	"ป": "0110", "ผ": "1101", "ฝ": "10101", "พ": "01100", "ฟ": "0010",
	"ม": "11", "ย": "1011", "ร": "010", "ล": "0100", "ว": "011",
	"ส": "000", "ห": "0000", "อ": "10001", "ฮ": "11011", "ฤ": "01011",
	"ะ": "01000", "า": "01", "ิ": "00100", "ี": "00", "ึ": "00110",
	"ื": "0011", "ุ": "00101", "ู": "1110", "เ": "0", "แ": "0101",
	"ไ": "01001", "โ": "111", "ำ": "00010", "่": "001", "้": "0001",
	"๊": "11000", "๋": "01010", "ั": "01101", "็": "11100", "์": "11001",
	"ๆ": "10111", "ฯ": "11010",
}

// MorseCodeEncode encodes string to Morse Code.
type MorseCodeEncode struct{}

func (p MorseCodeEncode) Name() string {
	return "morse-encode"
}

func (p MorseCodeEncode) Alias() []string {
	return []string{"morse-enc", "morse-encode", "morse-code-encode", "morse-code-enc"}
}

func (p MorseCodeEncode) Transform(data []byte, _ ...Flag) (string, error) {
	res := ""
	dash := "-"
	dot := "."
	space := "/"
	letterSeparator := " "
	str := strings.ReplaceAll(string(data), "\n", " ")
	str = strings.ReplaceAll(str, "\t", " ")
	str = strings.ReplaceAll(str, "\r", " ")
	str = strings.ToUpper(str)

	for _, part := range str {
		p := string(part)
		if p == " " {
			if space != "" {
				res += space + letterSeparator
			}
		} else if moreCodeMap[p] != "" {
			res += moreCodeMap[p] + letterSeparator
		}
	}
	res = strings.TrimSpace(res)
	res = strings.ReplaceAll(res, "1", dash)
	res = strings.ReplaceAll(res, "0", dot)
	return res, nil
}

func (p MorseCodeEncode) Flags() []Flag {
	return nil
}

func (p MorseCodeEncode) Title() string {
	title := "Morse Code Encoding"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p MorseCodeEncode) Description() string {
	return "Encode your text to Morse Code"
}

func (p MorseCodeEncode) FilterValue() string {
	return p.Title()
}

// MorseCodeDecode decodes Morse Code to string
type MorseCodeDecode struct{}

func (p MorseCodeDecode) Name() string {
	return "morse-decode"
}

func (p MorseCodeDecode) Alias() []string {
	return []string{"morse-dec", "morse-decode", "morse-code-decode", "morse-code-dec"}
}

func (p MorseCodeDecode) Transform(data []byte, _ ...Flag) (string, error) {
	res := ""
	wordSeparator := "/"
	letterSeparator := " "
	for _, part := range strings.Split(string(data), letterSeparator) {
		found := false
		for key, val := range moreCodeMap {
			if val == part {
				res += key
				found = true
				break
			}
		}
		if part == wordSeparator {
			res += " "
			found = true
		}
		if !found {
			return res, fmt.Errorf("unknown character " + part)
		}
	}

	return res, nil
}

func (p MorseCodeDecode) Flags() []Flag {
	return nil
}

func (p MorseCodeDecode) Title() string {
	title := "Morse Code Decode"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p MorseCodeDecode) Description() string {
	return "Decode Morse Code to text"
}

func (p MorseCodeDecode) FilterValue() string {
	return p.Title()
}
