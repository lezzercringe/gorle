package gorle

import (
	"fmt"
	"strings"
	"unicode"
)

type Opts struct {
	EnableEscapeSeq bool
	EscapeChar      rune
}

type Decoder struct {
	opts Opts
}

func NewDecoder(opts ...func(*Opts)) *Decoder {
	o := Opts{
		EnableEscapeSeq: false,
		EscapeChar:      '#',
	}

	for _, modifier := range opts {
		modifier(&o)
	}

	return &Decoder{
		opts: o,
	}
}

func WithEscapeSeq(enable bool) func(*Opts) {
	return func(o *Opts) {
		o.EnableEscapeSeq = enable
	}
}

func WithEscapeChar(ch rune) func(*Opts) {
	return func(o *Opts) {
		o.EscapeChar = ch
	}
}

func (d *Decoder) DecodeRunes(encoded []rune) (string, error) {
	var b strings.Builder
	b.Grow(len(encoded))

	var unpack rune
	for i := 0; i < len(encoded); {
		r := encoded[i]

		if unicode.IsDigit(r) {
			// NOTE: or repeat?
			value, advance := d.parseMultiplier(encoded, i)
			if value == 0 {
				return "", fmt.Errorf("zero multiplier is not allowed")
			}

			if unpack == 0 {
				return "", fmt.Errorf("misplaced multiplier %d - no chars before, try escaping it", value)
			}
			b.Grow(value)
			for range value {
				b.WriteRune(unpack)
			}
			unpack = 0
			i += advance
			continue
		}

		if unpack != 0 {
			b.WriteRune(unpack)
		}

		if d.opts.EnableEscapeSeq && r == d.opts.EscapeChar {
			if i+1 == len(encoded) {
				return "", fmt.Errorf("escape sequence started of the input end")
			}

			unpack = encoded[i+1]
			if !unicode.IsDigit(unpack) {
				return "", fmt.Errorf("invalid escape sequence: %c%c", d.opts.EscapeChar, unpack)
			}

			i += 2
			continue
		}

		unpack = r
		i++
	}

	if unpack != 0 {
		b.WriteRune(unpack)
	}

	return b.String(), nil
}

// parseMultiplier must be called with encoded[start] being a digit
// and returns the multiplier value and how many runes were consumed.
func (d *Decoder) parseMultiplier(encoded []rune, start int) (value, advance int) {
	value = int(encoded[start] - '0')
	i := start + 1
	for ; i < len(encoded) && unicode.IsDigit(encoded[i]); i++ {
		value *= 10
		value += int(encoded[i] - '0')
	}
	return value, i - start
}
