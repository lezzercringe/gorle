package gorle_test

import (
	"strings"
	"testing"

	"github.com/lezzercringe/gorle"
)

func TestDecode_WithoutEscapeSequences(t *testing.T) {
	cases := []struct {
		input   string
		want    string
		wantErr string
	}{
		{"s10", "ssssssssss", ""},
		{"s#110n", "s" + strings.Repeat("#", 110) + "n", ""},
		{"s#110n#", "s" + strings.Repeat("#", 110) + "n#", ""},
		{"abc3de4fghj1k", "abcccdeeeefghjk", ""},
		{"123", "", "misplaced multiplier 123 - no chars before, try escaping it"},
		{"#1#2#3", "######", ""},
		{"#-5asdf", "#-----asdf", ""},
		{"#a-5asdf", "#a-----asdf", ""},
		{"-0", "", "zero multiplier is not allowed"},
		{"asdf0", "", "zero multiplier is not allowed"},
		{"0", "", "zero multiplier is not allowed"},
		{"#", "#", ""},
	}

	for _, tt := range cases {
		t.Run(tt.input, func(t *testing.T) {
			s, err := gorle.Decode(tt.input)

			if err == nil && tt.wantErr != "" {
				t.Fatalf("expected error containing %s, got nil", tt.wantErr)
			}

			if err != nil {
				if tt.wantErr == "" {
					t.Fatalf("unexpected err: %v", err)
				}

				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("error %v does not contain substring %s", err, tt.wantErr)
				}
			}

			if s != tt.want {
				t.Fatalf("expected output=%v, got=%v", tt.want, s)
			}
		})
	}
}

func TestDecode_WithEscapeSequences(t *testing.T) {
	cases := []struct {
		input   string
		want    string
		wantErr string
	}{
		{"s10", "ssssssssss", ""},
		{"s#110n", "s1111111111n", ""},
		{"s#110n#", "", "escape sequence started of the input end"},
		{"abc3de4fghj1k", "abcccdeeeefghjk", ""},
		{"123", "", "misplaced multiplier 123 - no chars before, try escaping it"},
		{"#1#2#3", "123", ""},
		{"#-5asdf", "", "invalid escape sequence: #-"},
		{"#a-5asdf", "", "invalid escape sequence: #a"},
		{"-0", "", "zero multiplier is not allowed"},
		{"asdf0", "", "zero multiplier is not allowed"},
		{"0", "", "zero multiplier is not allowed"},
		{"#", "", "escape sequence started of the input end"},
		{"", "", ""},
	}

	for _, tt := range cases {
		t.Run(tt.input, func(t *testing.T) {
			s, err := gorle.Decode(tt.input, gorle.WithEscapeChar('#'), gorle.WithEscapeSeq(true))

			if err == nil && tt.wantErr != "" {
				t.Fatalf("expected error containing %s, got nil", tt.wantErr)
			}

			if err != nil {
				if tt.wantErr == "" {
					t.Fatalf("unexpected err: %v", err)
				}

				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("error %v does not contain substring %s", err, tt.wantErr)
				}
			}

			if s != tt.want {
				t.Fatalf("expected output=%v, got=%v", tt.want, s)
			}
		})
	}
}

func TestDecode_Runes_WithEscapeSequences(t *testing.T) {
	decoder := gorle.NewDecoder(gorle.WithEscapeChar('#'), gorle.WithEscapeSeq(true))

	cases := []struct {
		input   string
		want    string
		wantErr string
	}{
		{"s10", "ssssssssss", ""},
		{"s#110n", "s1111111111n", ""},
		{"s#110n#", "", "escape sequence started of the input end"},
		{"abc3de4fghj1k", "abcccdeeeefghjk", ""},
		{"123", "", "misplaced multiplier 123 - no chars before, try escaping it"},
		{"#1#2#3", "123", ""},
		{"#-5asdf", "", "invalid escape sequence: #-"},
		{"#a-5asdf", "", "invalid escape sequence: #a"},
		{"-0", "", "zero multiplier is not allowed"},
		{"asdf0", "", "zero multiplier is not allowed"},
		{"0", "", "zero multiplier is not allowed"},
		{"#", "", "escape sequence started of the input end"},
		{"", "", ""},
	}

	for _, tt := range cases {
		t.Run(tt.input, func(t *testing.T) {
			s, err := decoder.DecodeRunes([]rune(tt.input))

			if err == nil && tt.wantErr != "" {
				t.Fatalf("expected error containing %s, got nil", tt.wantErr)
			}

			if err != nil {
				if tt.wantErr == "" {
					t.Fatalf("unexpected err: %v", err)
				}

				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("error %v does not contain substring %s", err, tt.wantErr)
				}
			}

			if s != tt.want {
				t.Fatalf("expected output=%v, got=%v", tt.want, s)
			}
		})
	}
}

func TestDecode_Runes_WithoutEscapeSequences(t *testing.T) {
	decoder := gorle.NewDecoder(gorle.WithEscapeSeq(false))

	cases := []struct {
		input   string
		want    string
		wantErr string
	}{
		{"s10", "ssssssssss", ""},
		{"s#110n", "s" + strings.Repeat("#", 110) + "n", ""},
		{"s#110n#", "s" + strings.Repeat("#", 110) + "n#", ""},
		{"abc3de4fghj1k", "abcccdeeeefghjk", ""},
		{"123", "", "misplaced multiplier 123 - no chars before, try escaping it"},
		{"#1#2#3", "######", ""},
		{"#-5asdf", "#-----asdf", ""},
		{"#a-5asdf", "#a-----asdf", ""},
		{"-0", "", "zero multiplier is not allowed"},
		{"asdf0", "", "zero multiplier is not allowed"},
		{"0", "", "zero multiplier is not allowed"},
		{"#", "#", ""},
	}

	for _, tt := range cases {
		t.Run(tt.input, func(t *testing.T) {
			s, err := decoder.DecodeRunes([]rune(tt.input))

			if err == nil && tt.wantErr != "" {
				t.Fatalf("expected error containing %s, got nil", tt.wantErr)
			}

			if err != nil {
				if tt.wantErr == "" {
					t.Fatalf("unexpected err: %v", err)
				}

				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("error %v does not contain substring %s", err, tt.wantErr)
				}
			}

			if s != tt.want {
				t.Fatalf("expected output=%v, got=%v", tt.want, s)
			}
		})
	}
}
