package groupsimilar

import (
	"hash/fnv"
	"testing"
)

func TestToRunes(t *testing.T) {
	bf := NewStringVector(fnv.New64())
	ret := bf.ToVecotor("新闻历史，新闻历史!中国wenhua!, xsdsljf")
	t.Log(ret)
}
