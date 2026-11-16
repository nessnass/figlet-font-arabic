// Arabic shaping: each letter takes the form its neighbours call for, and lam
// before alef becomes one ligature. The font draws the forms (Arabic
// Presentation Forms-B), so text is shaped before it is rendered.
package main

const (
	alone = iota
	final
	initial
	medial
)

// joinsNext reports whether a letter joins the letter after it. Letters
// without an initial form (ا د ذ ر ز و and others) never do.
func joinsNext(c rune) bool {
	f, ok := letterForms[c]
	return ok && f[initial] != 0
}

// joinsPrev reports whether a letter can be joined from the letter before
// it. Only the hamza ء can't.
func joinsPrev(c rune) bool {
	f, ok := letterForms[c]
	return ok && f[final] != 0
}

// isMark reports whether c is a vowel mark or tatweel. The font has no
// glyphs for them, so shaping leaves them out.
func isMark(c rune) bool {
	return (c >= 0x064B && c <= 0x065F) || c == 0x0670 || c == 0x0640
}

// shape turns Arabic letters into their presentation forms, in reading
// order. Other characters pass through unchanged.
func shape(text string) string {
	var chars []rune
	for _, c := range text {
		if !isMark(c) {
			chars = append(chars, c)
		}
	}
	out := make([]rune, 0, len(chars))
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		before := i > 0 && joinsNext(chars[i-1]) && joinsPrev(c)

		if lig, ok := lamAlef[peek(chars, i+1)]; c == 'ل' && ok {
			if before {
				out = append(out, lig[1])
			} else {
				out = append(out, lig[0])
			}
			i++
			continue
		}

		forms, ok := letterForms[c]
		if !ok {
			out = append(out, c)
			continue
		}
		after := joinsNext(c) && joinsPrev(peek(chars, i+1))
		form := map[[2]bool]int{
			{false, false}: alone, {false, true}: initial,
			{true, true}: medial, {true, false}: final,
		}[[2]bool{before, after}]
		for forms[form] == 0 { // a letter without this form takes the nearest
			form = map[int]int{medial: final, initial: alone, final: alone}[form]
		}
		out = append(out, forms[form])
	}
	return string(out)
}

func peek(chars []rune, i int) rune {
	if i < len(chars) {
		return chars[i]
	}
	return 0
}

// reverse returns the runes of s in reverse order: the font is right to
// left, so the last letter is drawn first, at the left.
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
