package kitty

import "strconv"

// Kitty keyboard protocol progressive enhancement flags.
// See https://sw.kovidgoyal.net/kitty/keyboard-protocol/#progressive-enhancement
const (
	DisambiguateEscapeCodes = 1 << iota
	ReportEventTypes
	ReportAlternateKeys
	ReportAllKeys
	ReportAssociatedKeys

	AllFlags = DisambiguateEscapeCodes | ReportEventTypes | ReportAlternateKeys | ReportAllKeys | ReportAssociatedKeys
)

// Request is a sequence to request the terminal Kitty keyboard protocol
// enabled flags.
const Request = "\x1b[?u"

// Enable returns a sequence to enable the given flags.
//
//	CSI > flags u
//
// See https://sw.kovidgoyal.net/kitty/keyboard-protocol/#progressive-enhancement
func Enable(flags int) string {
	var n string
	if flags > 0 {
		n = strconv.Itoa(flags)
	}

	return "\x1b[>" + n + "u"
}

// Disable returns a sequence to disable the given flags.
//
//	CSI < flags u
//
// See https://sw.kovidgoyal.net/kitty/keyboard-protocol/#progressive-enhancement
func Disable(flags int) string {
	var n string
	if flags > 0 {
		n = strconv.Itoa(flags)
	}

	return "\x1b[<" + n + "u"
}
