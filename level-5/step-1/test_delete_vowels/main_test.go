package main

// import "testing"

// func TestDeleteVowels(t *testing.T) {
// 	cases := []struct {
// 		name  string
// 		value string
// 		want  string
// 	}{
// 		{
// 			name:  "lowercase vowels",
// 			value: "aeiou",
// 			want:  "",
// 		},
// 		{
// 			name:  "uppercase vowels",
// 			value: "AEIOU",
// 			want:  "",
// 		},
// 		{
// 			name:  "mixed case",
// 			value: "HeLLo WoRLD",
// 			want:  "HLL WRLD",
// 		},
// 		{
// 			name:  "letter y",
// 			value: "yellow sky rhythm",
// 			want:  "yllw sky rhythm",
// 		},
// 		{
// 			name:  "empty string",
// 			value: "",
// 			want:  "",
// 		},
// 		{
// 			name:  "no vowels",
// 			value: "bcdfg",
// 			want:  "bcdfg",
// 		},
// 		{
// 			name:  "vowels between consonants",
// 			value: "banana",
// 			want:  "bnn",
// 		},
// 		{
// 			name:  "numbers and punctuation",
// 			value: "Hello, 123!",
// 			want:  "Hll, 123!",
// 		},
// 		{
// 			name:  "only y",
// 			value: "Yy",
// 			want:  "Yy",
// 		},
// 	}

// 	for _, tc := range cases {
// 		tc := tc

// 		t.Run(tc.name, func(t *testing.T) {
// 			got := DeleteVowels(tc.value)

// 			if got != tc.want {
// 				t.Errorf("DeleteVowels(%q) = %q; want %q",
// 					tc.value, got, tc.want)
// 			}
// 		})
// 	}
// }
