package main

import (
	"fmt"
	"html"
	"math"
	"net/url"
	"strconv"
	"strings"
	"unicode/utf8"
)

func dump(params ...interface{}) {
	fmt.Println(params...)
}

func main() {
	dump("")
	dump("Japan 日本")
	dump("\xe6\x97\xa5")
	dump("\\")
	dump("\"")
	dump("\n")
	dump(`\xe6`)
	dump(html.EscapeString("<>"))
	dump(url.PathEscape("A B"))

	dump("Ja" + "pan")

	dump("Japan" == "Japan")
	dump(strings.EqualFold("Japan", "JAPAN"))
	dump("Japan" < "japan")

	dump(len("日"))
	dump(utf8.RuneCountInString("日"))
	dump(utf8.ValidString("日"))

	dump("Japan"[2])
	dump("Japan"[1:3])
	dump("Japan"[:2])
	dump("Japan"[2:])

	for i, ch := range "Japan 日本" {
		fmt.Printf("%d:%q ", i, ch)
	}
	// Output: 0:'J' 1:'a' 2:'p' 3:'a' 4:'n' 5:' ' 6:'日' 9:'本'

	s := "Japan 日本"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%q ", s[i])
	}
	// Output: 'J' 'a' 'p' 'a' 'n' ' ' 'æ' '\u0097' '¥' 'æ' '\u009c' '¬'

	dump(strings.Contains("Japan", "abc"))
	dump(strings.ContainsAny("Japan", "abc"))
	dump(strings.Count("Banana", "ana"))
	dump(strings.HasPrefix("Japan", "Ja"))
	dump(strings.HasSuffix("Japan", "pan"))
	dump(strings.Index("Japan", "abc"))
	dump(strings.IndexAny("Japan", "abc"))
	dump(strings.LastIndex("Japan", "abc"))
	dump(strings.LastIndexAny("Japan", "abc"))

	dump(strings.Replace("foo", "o", ".", 2))
	f := func(r rune) rune {
		return r + 1
	}
	dump(strings.Map(f, "ab"))
	dump(strings.ToUpper("Japan"))
	dump(strings.ToLower("Japan"))
	dump(strings.Title("ja pan"))
	dump(strings.TrimSpace(" foo\n"))
	dump(strings.Trim("foo", "fo"))
	dump(strings.TrimLeft("foo", "f"))
	dump(strings.TrimRight("foo", "o"))
	dump(strings.TrimPrefix("foo", "fo"))
	dump(strings.TrimSuffix("foo", "o"))

	dump(strings.Fields(" a\t b\n"))
	dump(strings.Split("a,b", ","))
	dump(strings.SplitAfter("a,b", ","))

	dump(strings.Join([]string{"a", "b"}, ":"))
	dump(strings.Repeat("da", 2))

	dump(strconv.Itoa(-42))
	dump(strconv.FormatInt(255, 16))
	str := fmt.Sprintf("%.4f", math.Pi) // s == "3.1416"
	dump(str)
}
