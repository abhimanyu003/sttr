// Based on https://github.com/ozdemirburak/morse-decoder
package processors

import (
	"fmt"
	"strings"
)

var morseCodeLa = map[string]string{
	// Latin => https://en.wikipedia.org/wiki/Morse_code
	"A": ".-", "B": "-...", "C": "-.-.", "D": "-..", "E": ".", "F": "..-.",
	"G": "--.", "H": "....", "I": "..", "J": ".---", "K": "-.-", "L": ".-..",
	"M": "--", "N": "-.", "O": "---", "P": ".--.", "Q": "--.-", "R": ".-.",
	"S": "...", "T": "-", "U": "..-", "V": "...-", "W": ".--", "X": "-..-",
	"Y": "-.--", "Z": "--..",
	// Numbers
	"0": "-----", "1": ".----", "2": "..---", "3": "...--", "4": "....-",
	"5": ".....", "6": "-....", "7": "--...", "8": "---..", "9": "----.",

	// Punctuation
	".": ".-.-.-", ",": "--..--", "?": "..--..", "'": ".----.", "!": "-.-.--", "/": "-..-.",
	"(": "-.--.", ")": "-.--.-", "&": ".-...", ":": "---...", ";": "-.-.-.", "=": "-...-",
	"+": ".-.-.", "-": "-....-", "_": "..--.-", `"`: ".-..-.", "$": "...-..-", "@": ".--.-.",
	"¿": "..-.-", "¡": "--...-",

	// Non-Latin extensions
	"Ã": ".--.-", "Á": ".--.-", "Å": ".--.-", "À": ".--.-", "Â": ".--.-", "Ä": ".-.-",
	"Ą": ".-.-", "Æ": ".-.-", "Ç": "-.-..", "Ć": "-.-..", "Ĉ": "-.-..",
	"Ð": "..--.", "È": ".-..-", "Ę": "..-..", "Ë": "..-..", "É": "..-..",
	"Ê": "-..-.", "Ğ": "--.-.", "Ĝ": "--.-.", "Ĥ": "----", "İ": ".-..-", "Ï": "-..--",
	"Ì": ".---.", "Ĵ": ".---.", "Ł": ".-..-", "Ń": "--.--", "Ñ": "--.--", "Ó": "---.",
	"Ò": "---.", "Ö": "---.", "Ô": "---.", "Ø": "---.", "Ś": "...-...", "Ş": ".--..",
	"Ș": "----", "Š": "----", "Ŝ": "...-.", "ß": "......", "Þ": ".--..", "Ü": "..--",
	"Ù": "..--", "Ŭ": "..--", "Ž": "--..-", "Ź": "--..-.", "Ż": "--..-",
	// Persian Alphabet => https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets
	// "پ": ".--.", "چ": "---.", "ژ": "--.", "ک": "-.-", "گ": "--.-", "ی": "..",

}

var morseCodeRu = map[string]string{
	// Cyrillic Alphabet => https://en.wikipedia.org/wiki/Russian_Morse_code
	"А": ".-", "Б": "-...", "В": ".--", "Г": "--.", "Д": "-..", "Е": ".",
	"Ж": "...-", "З": "--..", "И": "..", "Й": ".---", "К": "-.-", "Л": ".-..",
	"М": "--", "Н": "-.", "О": "---", "П": ".--.", "Р": ".-.", "С": "...",
	"Т": "-", "У": "..-", "Ф": "..-.", "Х": "....", "Ц": "-.-.", "Ч": "---.",
	"Ш": "----", "Щ": "--.-", "Ъ": "--.--", "Ы": "-.--", "Ь": "-..-", "Э": "..-..",
	"Ю": "..--", "Я": ".-.-", "Ї": ".---.", "Є": "..-..", "І": "..", "Ґ": "--.",
}

var morseCodeGr = map[string]string{
	// Greek Alphabet => https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets
	"Α": ".-", "Β": "-...", "Γ": "--.", "Δ": "-..", "Ε": ".", "Ζ": "--..",
	"Η": "....", "Θ": "-.-.", "Ι": "..", "Κ": "-.-", "Λ": ".-..", "Μ": "--",
	"Ν": "-.", "Ξ": "-..-", "Ο": "---", "Π": ".--.", "Ρ": ".-.", "Σ": "...",
	"Τ": "-", "Υ": "-.--", "Φ": "..-.", "Χ": "----", "Ψ": "--.-", "Ω": ".--",
}

var morseCodeHe = map[string]string{
	// Hebrew Alphabet => https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets
	"א": ".-", "ב": "-...", "ג": "--.", "ד": "-..", "ה": "---", "ו": ".",
	"ז": "--..", "ח": "....", "ט": "..-", "י": "..", "כ": "-.-", "ל": ".-..",
	"מ": "--", "נ": "-.", "ס": "-.-.", "ע": ".---", "פ": ".--.", "צ": ".--",
	"ק": "--.-", "ר": ".-.", "ש": "...", "ת": "-",
}

var morseCodeAr = map[string]string{
	// Arabic Alphabet => https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets
	"ا": ".-", "ب": "-...", "ت": "-", "ث": "-.-.", "ج": ".---", "ح": "....",
	"خ": "---", "د": "-..", "ذ": "--..", "ر": ".-.", "ز": "---.", "س": "...",
	"ش": "----", "ص": "-..-", "ض": "...-", "ط": "..-", "ظ": "-.--", "ع": ".-.-",
	"غ": "--.", "ف": "..-.", "ق": "--.-", "ك": "-.-", "ل": ".-..", "م": "--",
	"ن": "-.", "ه": "..-..", "و": ".--", "ي": "..", "ﺀ": ".",
}

var morseCodeJa = map[string]string{
	// Japanese Alphabet => https://ja.wikipedia.org/wiki/%E3%83%A2%E3%83%BC%E3%83%AB%E3%82%B9%E7%AC%A6%E5%8F%B7#%E5%92%8C%E6%96%87%E3%83%A2%E3%83%BC%E3%83%AB%E3%82%B9%E7%AC%A6%E5%8F%B7
	"ア": "--.--", "カ": ".-..", "サ": "-.-.-", "タ": "-.", "ナ": ".-.", "ハ": "-...",
	"マ": "-..-", "ヤ": ".--", "ラ": "...", "ワ": "-.-", "イ": ".-", "キ": "-.-..",
	"シ": "--.-.", "チ": "..-.", "ニ": "-.-.", "ヒ": "--..-", "ミ": "..-.-", "リ": "--.",
	"ヰ": ".-..-", "ウ": "..-", "ク": "...-", "ス": "---.-", "ツ": ".--.", "ヌ": "....",
	"フ": "--..", "ム": "-", "ユ": "-..--", "ル": "-.--.", "ン": ".-.-.", "エ": "-.---",
	"ケ": "-.--", "セ": ".---.", "テ": ".-.--", "ネ": "--.-", "ヘ": ".", "メ": "-...-",
	"レ": "---", "ヱ": ".--..", "オ": ".-...", "コ": "----", "ソ": "---.", "ト": "..-..",
	"ノ": "..--", "ホ": "-..", "モ": "-..-.", "ヨ": "--", "ロ": ".-.-", "ヲ": ".---",
	"゛": "..", "゜": "..--.", "。": ".-.-..", "ー": ".--.-", "、": ".-.-.-",
	"（": "-.--.-", "）": ".-..-.",
}

var morseCodeKr = map[string]string{
	// Korean Alphabet => https://en.wikipedia.org/wiki/SKATS
	"ㄱ": ".-..", "ㄴ": "..-.", "ㄷ": "-...", "ㄹ": "...-", "ㅁ": "--", "ㅂ": ".--",
	"ㅅ": "--.", "ㅇ": "-.-", "ㅈ": ".--.", "ㅊ": "-.-.", "ㅋ": "-..-", "ㅌ": "--..",
	"ㅍ": "---", "ㅎ": ".---", "ㅏ": ".", "ㅑ": "..", "ㅓ": "-", "ㅕ": "...",
	"ㅗ": ".-", "ㅛ": "-.", "ㅜ": "....", "ㅠ": ".-.", "ㅡ": "-..", "ㅣ": "..-",
}

var morseCodeTh = map[string]string{
	// Thai Alphabet => https://th.wikipedia.org/wiki/รหัสมอร์ส
	"ก": "--.", "ข": "-.-.", "ค": "-.-", "ง": "-.--.", "จ": "-..-.",
	"ฉ": "----", "ช": "-..-", "ซ": "--..", "ญ": ".---", "ด": "-..",
	"ต": "-", "ถ": "-.-..", "ท": "-..--", "น": "-.", "บ": "-...",
	"ป": ".--.", "ผ": "--.-", "ฝ": "-.-.-", "พ": ".--..", "ฟ": "..-.",
	"ม": "--", "ย": "-.--", "ร": ".-.", "ล": ".-..", "ว": ".--",
	"ส": "...", "ห": "....", "อ": "-...-", "ฮ": "--.--", "ฤ": ".-.--",
	"ะ": ".-...", "า": ".-", "ิ": "..-..", "ี": "..", "ึ": "..--.",
	"ื": "..--", "ุ": "..-.-", "ู": "---.", "เ": ".", "แ": ".-.-",
	"ไ": ".-..-", "โ": "---", "ำ": "...-.", "่": "..-", "้": "...-",
	"๊": "--...", "๋": ".-.-.", "ั": ".--.-", "็": "---..", "์": "--..-",
	"ๆ": "-.---", "ฯ": "--.-.",
}

var morseCodeMap = map[string]map[string]string{
	"la": morseCodeLa,
	"ru": morseCodeRu,
	"gr": morseCodeGr,
	"he": morseCodeHe,
	"ar": morseCodeAr,
	"ja": morseCodeJa,
	"kr": morseCodeKr,
	"th": morseCodeTh,
}

var morseLangs = []string{
	"la", "ru", "gr", "he", "ar", "ja", "kr", "th",
}
var morseDecodeLangFlag = Flag{
	Name:  "lang",
	Short: "l",
	Desc:  "Morse code set to decode [la(Latin), ru(Cyrillic), gr(Greek), he(Hebrew), ar(Arabic), ja(Japanese), kr(Korean), th(Thai)]",
	Value: "la",
	Type:  FlagString,
}

func checkDecodeLangFlag(f []Flag) string {
	lang := "la"
	for _, flag := range f {
		if flag.Short == "l" {
			l, ok := flag.Value.(string)
			if ok {
				for _, ml := range morseLangs {
					if ml == l {
						lang = l
						break
					}
				}
			}
		}
	}
	return lang
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
		} else {
			for _, val := range morseCodeMap {
				if val[p] != "" {
					res += val[p] + letterSeparator
					break
				}
			}
		}
	}
	res = strings.TrimSpace(res)
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

func (p MorseCodeDecode) Transform(data []byte, f ...Flag) (string, error) {
	res := ""
	wordSeparator := "/"
	letterSeparator := " "
	codeLang := checkDecodeLangFlag(f)
	for _, part := range strings.Split(string(data), letterSeparator) {
		found := false
		for key, val := range morseCodeMap[codeLang] {
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
			return res, fmt.Errorf("unknown character < %s > in + %s", part, codeLang)
		}
	}

	return res, nil
}

func (p MorseCodeDecode) Flags() []Flag {
	return []Flag{morseDecodeLangFlag}
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
