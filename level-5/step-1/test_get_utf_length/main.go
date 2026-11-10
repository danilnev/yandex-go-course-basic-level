package main

// import (
// 	"testing"
// )

// func TestGetUTFLength(t *testing.T) {

// 	cases := []struct {
// 		name        string
// 		input       []byte
// 		wantLength  int
// 		wantErr     error
// 	}{
// 		{"English", []byte("Hello"), 5, nil},
// 		{"Russian", []byte("Привет"), 6, nil},
// 		{"Chinese", []byte("你好"), 2, nil},
// 		{"Mixed", []byte("Hello, мир!"), 11, nil},
// 		{"Numbers", []byte("12345"), 5, nil},
// 		{"Empty", []byte(""), 0, nil},
// 		{"Nil", nil, 0, nil},
// 		{"Emoji", []byte("😀🎉"), 2, nil},

// 		{"InvalidFF", []byte{0xFF}, 0, ErrInvalidUTF8},
// 		{"InvalidFE", []byte{0xFE}, 0, ErrInvalidUTF8},
// 		{"InvalidSequence", []byte{0xC3, 0x28}, 0, ErrInvalidUTF8},
// 		{"IncompleteSequence", []byte{0xE2, 0x82}, 0, ErrInvalidUTF8},
// 		{"MixedInvalid", []byte{'A', 0xFF, 'B'}, 0, ErrInvalidUTF8},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			got, err := GetUTFLength(tc.input)
// 			if got != tc.wantLength || err != tc.wantErr {
// 				t.Fatalf("GetUTFLength(%q) = (%d, %v), want (%d, %v)", tc.input, got, err, tc.wantLength, tc.wantErr)
// 			}
// 		})
// 	}

// }
