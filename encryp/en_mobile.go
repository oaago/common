package encryp

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/oaago/component/logx"
)

const phoneKey = "jiklmnopqrstuv+RSTUwxyz3210549876Q/ZDCBAEFGHIJKLMNOPabdcefghVYXW="

var hex = []string{
	"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "0A", "0B", "0C", "0D", "0E", "0F",
	"10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "1A", "1B", "1C", "1D", "1E", "1F",
	"20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "2A", "2B", "2C", "2D", "2E", "2F",
	"30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "3A", "3B", "3C", "3D", "3E", "3F",
	"40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "4A", "4B", "4C", "4D", "4E", "4F",
	"50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "5A", "5B", "5C", "5D", "5E", "5F",
	"60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "6A", "6B", "6C", "6D", "6E", "6F",
	"70", "71", "72", "73", "74", "75", "76", "77", "78", "79", "7A", "7B", "7C", "7D", "7E", "7F",
	"80", "81", "82", "83", "84", "85", "86", "87", "88", "89", "8A", "8B", "8C", "8D", "8E", "8F",
	"90", "91", "92", "93", "94", "95", "96", "97", "98", "99", "9A", "9B", "9C", "9D", "9E", "9F",
	"A0", "A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8", "A9", "AA", "AB", "AC", "AD", "AE", "AF",
	"B0", "B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8", "B9", "BA", "BB", "BC", "BD", "BE", "BF",
	"C0", "C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9", "CA", "CB", "CC", "CD", "CE", "CF",
	"D0", "D1", "D2", "D3", "D4", "D5", "D6", "D7", "D8", "D9", "DA", "DB", "DC", "DD", "DE", "DF",
	"E0", "E1", "E2", "E3", "E4", "E5", "E6", "E7", "E8", "E9", "EA", "EB", "EC", "ED", "EE", "EF",
	"F0", "F1", "F2", "F3", "F4", "F5", "F6", "F7", "F8", "F9", "FA", "FB", "FC", "FD", "FE", "FF",
}

var val = []int{
	0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F,
	0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F,
	0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F,
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F,
	0x3F, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F,
	0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F,
	0x3F, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F,
	0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F,
	0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F,
	0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F,
	0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F,
	0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F,
	0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F,
	0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F,
	0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F,
	0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F, 0x3F,
}

func DecMobile(input string) string {
	if input == "" {
		return ""
	}

	var data = []byte(input)

	var output string = ""
	var chr1, chr2, chr3 int
	var enc1, enc2, enc3, enc4 int
	var i int = 0
	//数据参照，其中字母顺序与加密的参照表中的顺序必须一样，解码出来的数据才能一致
	var keyStr string = phoneKey
	//检测编码合法性必须是 A-Z, a-z, 0-9, +, /, or =
	for index := range data {
		var byteFont = string(data[index])
		if !strings.Contains(keyStr, byteFont) {
			logx.Logger.Error("There were invalid base64 characters in the input text." + "\n" +
				"Valid base64 characters are A-Z, a-z, 0-9, '+', '/', and '='" + "\n" + "Expect errors in decoding.")
			return ""
		}
	}

	for {
		i++
		enc1 = strings.Index(keyStr, input[i-1:i])
		i++
		enc2 = strings.Index(keyStr, input[i-1:i])
		i++
		enc3 = strings.Index(keyStr, input[i-1:i])
		i++
		enc4 = strings.Index(keyStr, input[i-1:i])

		chr1 = (enc1 << 2) | (enc2 >> 4)
		chr2 = ((enc2 & 15) << 4) | (enc3 >> 2)
		chr3 = ((enc3 & 3) << 6) | enc4

		output = output + string(rune(chr1))

		if enc3 != 64 {
			output = output + string(rune(chr2))
		}
		if enc4 != 64 {
			output = output + string(rune(chr3))
		}

		chr1, chr2, chr3 = 0, 0, 0
		enc1, enc2, enc3, enc4 = 0, 0, 0, 0

		if i >= len(input) {
			break
		}

	}

	if unescape(output) == "" {
		fmt.Println(input)
	}
	return unescape(output)

}

func unescape(s string) string {
	var build strings.Builder
	i := 1
	len := len(s)
	for {
		if i <= len {
			ch := s[i-1 : i]
			if ch == "+" {
				build.WriteString(" ")
			} else if "A" <= ch && ch <= "Z" { // 'A'..'Z' : as it was
				build.WriteString(ch)
			} else if "a" <= ch && ch <= "z" {
				build.WriteString(ch)
			} else if "0" <= ch && ch <= "9" {
				build.WriteString(ch)
			} else if ch == "-" ||
				ch == "_" ||
				ch == "." ||
				ch == "!" ||
				ch == "~" ||
				ch == "*" ||
				ch == "/" ||
				ch == "(" ||
				ch == ")" {
				build.WriteString(ch)
			} else if ch == "%" {
				cint := 0
				if "u" != s[i:i+1] {
					cint = (cint << 4) | val[int([]rune((s[i : i+1]))[0])]
					cint = (cint << 4) | val[int([]rune((s[i+1 : i+2]))[0])]
					i += 2
				} else {
					cint = (cint << 4) | val[int([]rune((s[i+1 : i+2]))[0])]
					cint = (cint << 4) | val[int([]rune((s[i+2 : i+3]))[0])]
					cint = (cint << 4) | val[int([]rune((s[i+3 : i+4]))[0])]
					cint = (cint << 4) | val[int([]rune((s[i+4 : i+5]))[0])]
					i += 5
				}
				build.WriteString(strconv.Itoa(cint))
				fmt.Println(build.String())
			}

		} else {
			break
		}
		i++
	}
	return build.String()
}

func EncMobile(input string) string {
	var build strings.Builder
	if input == "" {
		return ""
	}
	input = escape(input)
	var chr1, chr2, chr3 int
	var enc1, enc2, enc3, enc4 int
	var i int = 0
	len := len(input)
	var keyStr string = phoneKey
	for {
		i++
		chr1 = int([]rune(input[i-1 : i])[0])
		enc1 = chr1 >> 2

		if i < len {
			i++
			chr2 = int([]rune(input[i-1 : i])[0])
			enc2 = ((chr1 & 3) << 4) | (chr2 >> 4)
			if !isNumber(chr2) {
				enc3, enc4 = 64, 64
			}

			if i < len {
				i++
				chr3 = int([]rune(input[i-1 : i])[0])
				enc3 = ((chr2 & 15) << 2) | (chr3 >> 6)
				enc4 = chr3 & 63
				if !isNumber(chr3) {
					enc4 = 64
				}
				build.WriteString("")
				index := enc1
				build.WriteString(keyStr[index : index+1])
				index = enc2
				build.WriteString(keyStr[index : index+1])
				index = enc3
				build.WriteString(keyStr[index : index+1])
				index = enc4
				build.WriteString(keyStr[index : index+1])
			} else {
				enc4 = (chr2 & 15) << 2
				build.WriteString("")
				index := enc1
				build.WriteString(keyStr[index : index+1])
				index = enc2
				build.WriteString(keyStr[index : index+1])
				index = enc4
				build.WriteString(keyStr[index : index+1])
				build.WriteString("=")
			}
		} else {
			enc4 = (chr1 & 3) << 4
			build.WriteString("")
			index := enc1
			build.WriteString(keyStr[index : index+1])
			index = enc4
			build.WriteString(keyStr[index : index+1])
			build.WriteString("==")
		}

		chr1, chr2, chr3 = 0, 0, 0
		enc1, enc2, enc3, enc4 = 0, 0, 0, 0

		if i >= len {
			break
		}

	}
	return build.String()

}

func escape(s string) string {
	var build strings.Builder
	len := len(s)
	for i := 1; i <= len; i++ {
		ch := s[i-1 : i]
		if ch == " " {
			build.WriteString("+")
		} else if "A" <= ch && ch <= "Z" { // 'A'..'Z' : as it was
			build.WriteString(ch)
		} else if "a" <= ch && ch <= "z" {
			build.WriteString(ch)
		} else if "0" <= ch && ch <= "9" {
			build.WriteString(ch)
		} else if ch == "-" ||
			ch == "_" ||
			ch == "." ||
			ch == "!" ||
			ch == "~" ||
			ch == "*" ||
			ch == "/" ||
			ch == "(" ||
			ch == ")" {
			build.WriteString(ch)
		} else if int([]rune((ch))[0]) <= 0x007F {
			build.WriteString("%")
			build.WriteString(hex[int([]rune((ch))[0])])
		} else { // unicode : map to %uXXXX
			build.WriteString("%")
			build.WriteString("u")
			build.WriteString(hex[uint(int([]rune((ch))[0]))>>8])
			build.WriteString(hex[(0x00FF & int([]rune((ch))[0]))])
		}
	}
	return build.String()

}

func isNumber(r int) bool {
	pattern := "^[-+]?(([0-9]+)([.]([0-9]+))?|([.]([0-9]+))?)$"
	match, _ := regexp.Match(pattern, []byte(strconv.Itoa(r)))
	return match
}
