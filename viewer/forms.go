// The forms of each letter in Arabic Presentation Forms-B: alone, final,
// initial, medial, 0 where the letter has no such form. Written by flf.py.
package main

var letterForms = map[rune][4]rune{
	'\u0627': {0xFE8D, 0xFE8E, 0, 0},           // ا alef
	'\u0628': {0xFE8F, 0xFE90, 0xFE91, 0xFE92}, // ب beh
	'\u062A': {0xFE95, 0xFE96, 0xFE97, 0xFE98}, // ت teh
	'\u062B': {0xFE99, 0xFE9A, 0xFE9B, 0xFE9C}, // ث theh
	'\u062C': {0xFE9D, 0xFE9E, 0xFE9F, 0xFEA0}, // ج jeem
	'\u062D': {0xFEA1, 0xFEA2, 0xFEA3, 0xFEA4}, // ح hah
	'\u062E': {0xFEA5, 0xFEA6, 0xFEA7, 0xFEA8}, // خ khah
	'\u062F': {0xFEA9, 0xFEAA, 0, 0},           // د dal
	'\u0630': {0xFEAB, 0xFEAC, 0, 0},           // ذ thal
	'\u0631': {0xFEAD, 0xFEAE, 0, 0},           // ر reh
	'\u0632': {0xFEAF, 0xFEB0, 0, 0},           // ز zain
	'\u0633': {0xFEB1, 0xFEB2, 0xFEB3, 0xFEB4}, // س seen
	'\u0634': {0xFEB5, 0xFEB6, 0xFEB7, 0xFEB8}, // ش sheen
	'\u0635': {0xFEB9, 0xFEBA, 0xFEBB, 0xFEBC}, // ص sad
	'\u0636': {0xFEBD, 0xFEBE, 0xFEBF, 0xFEC0}, // ض dad
	'\u0637': {0xFEC1, 0xFEC2, 0xFEC3, 0xFEC4}, // ط tah
	'\u0638': {0xFEC5, 0xFEC6, 0xFEC7, 0xFEC8}, // ظ zah
	'\u0639': {0xFEC9, 0xFECA, 0xFECB, 0xFECC}, // ع ain
	'\u063A': {0xFECD, 0xFECE, 0xFECF, 0xFED0}, // غ ghain
	'\u0641': {0xFED1, 0xFED2, 0xFED3, 0xFED4}, // ف feh
	'\u0642': {0xFED5, 0xFED6, 0xFED7, 0xFED8}, // ق qaf
	'\u0643': {0xFED9, 0xFEDA, 0xFEDB, 0xFEDC}, // ك kaf
	'\u0644': {0xFEDD, 0xFEDE, 0xFEDF, 0xFEE0}, // ل lam
	'\u0645': {0xFEE1, 0xFEE2, 0xFEE3, 0xFEE4}, // م meem
	'\u0646': {0xFEE5, 0xFEE6, 0xFEE7, 0xFEE8}, // ن noon
	'\u0647': {0xFEE9, 0xFEEA, 0xFEEB, 0xFEEC}, // ه heh
	'\u0648': {0xFEED, 0xFEEE, 0, 0},           // و waw
	'\u064A': {0xFEF1, 0xFEF2, 0xFEF3, 0xFEF4}, // ي yeh
	'\u0629': {0xFE93, 0xFE94, 0, 0},           // ة teh marbuta
	'\u0649': {0xFEEF, 0xFEF0, 0, 0},           // ى alef maksura
	'\u0621': {0xFE80, 0, 0, 0},                // ء hamza
	'\u0623': {0xFE83, 0xFE84, 0, 0},           // أ alef with hamza above
	'\u0625': {0xFE87, 0xFE88, 0, 0},           // إ alef with hamza below
	'\u0622': {0xFE81, 0xFE82, 0, 0},           // آ alef with madda above
	'\u0624': {0xFE85, 0xFE86, 0, 0},           // ؤ waw with hamza above
	'\u0626': {0xFE89, 0xFE8A, 0xFE8B, 0xFE8C}, // ئ yeh with hamza above
}

// Lam before alef, with or without hamza or madda, becomes one ligature
var lamAlef = map[rune][2]rune{
	'\u0627': {0xFEFB, 0xFEFC}, // لا
	'\u0623': {0xFEF7, 0xFEF8}, // لأ
	'\u0625': {0xFEF9, 0xFEFA}, // لإ
	'\u0622': {0xFEF5, 0xFEF6}, // لآ
}

// Every letter, in alphabet order, then the others of written Arabic
var alphabet = []string{"ا", "ب", "ت", "ث", "ج", "ح", "خ", "د", "ذ", "ر", "ز", "س", "ش", "ص", "ض", "ط", "ظ", "ع", "غ", "ف", "ق", "ك", "ل", "م", "ن", "ه", "و", "ي", "ة", "ى", "ء", "أ", "إ", "آ", "ؤ", "ئ", "لا", "لأ", "لإ", "لآ"}
