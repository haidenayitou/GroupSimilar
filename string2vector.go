package groupsimilar

import (
	"fmt"
	"hash"
	"unicode/utf8"
)

var (
	splits = []rune{',', '.', '!', ' ', '，'}
)

type StringVector struct {
	hashfn hash.Hash64
}

func NewStringVector(fn hash.Hash64) *StringVector {
	stringVector := new(StringVector)
	stringVector.hashfn = fn
	return stringVector
}

func toRunes(group string) []rune {
	ret := make([]rune, 0)
	for len(group) > 0 {
		r, size := utf8.DecodeRuneInString(group)
		ret = append(ret, r)
		fmt.Printf("%c, %v\n, %c", r, size, ret)
		group = group[size:]
	}
	fmt.Println("切换成字符串之后:", ret)
	return ret
}

func (sv *StringVector) ToVecotor(group string) []uint64 {
	runes := toRunes(group)
	matrix := make([][]rune, 0)
	temp_matrix := make([]rune, 0)
	start_index := 0
	end_index := 0
	for index, value := range runes {
		is_in := false
		for _, r := range splits {
			fmt.Println(r, value)
			if r == value {
				is_in = true
				break
			}
		}
		if is_in {
			if start_index != end_index {
				temp_matrix = runes[start_index:end_index]
				matrix = append(matrix, temp_matrix)
				temp_matrix = make([]rune, len(runes)-start_index)
			}
			start_index = index + 1
			end_index = index + 1
		} else {
			end_index = end_index + 1
		}
	}
	if start_index != len(runes) {
		matrix = append(matrix, runes[start_index:])
		fmt.Println(start_index, end_index)
	}

	fmt.Println("matrix", len(matrix), matrix)
	ret := make([]uint64, len(matrix))
	for index, mx := range matrix {
		buf := make([]byte, 0)
		for _, x := range mx {
			temp_buf := make([]byte, utf8.RuneLen(x))
			utf8.EncodeRune(temp_buf, x)
			buf = addSlice(buf, temp_buf)
		}
		sv.hashfn.Reset()
		sv.hashfn.Write(buf)
		ret[index] = sv.hashfn.Sum64()
	}
	return ret
}

func addSlice(a, b []byte) []byte {
	c := make([]byte, len(a)+len(b))
	copy(c, a)
	copy(c[len(a):], b)
	return c
}
