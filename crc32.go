package main

import (
	"hash"
	"hash/crc32"
)

// crc32Hash implements the hash.Hash interface for CRC32
type crc32Hash struct {
	table *crc32.Table
	crc   uint32
}

func crc32New() hash.Hash {
	return &crc32Hash{table: crc32.MakeTable(crc32.IEEE), crc: 0}
}

func (h *crc32Hash) Write(p []byte) (n int, err error) {
	h.crc = crc32.Update(h.crc, h.table, p)
	return len(p), nil
}

func (h *crc32Hash) Sum(b []byte) []byte {
	s := h.Sum32()
	return append(b, byte(s>>24), byte(s>>16), byte(s>>8), byte(s))
}

func (h *crc32Hash) Reset() {
	h.crc = 0
}

func (h *crc32Hash) Size() int {
	return 4
}

func (h *crc32Hash) BlockSize() int {
	return 1
}

func (h *crc32Hash) Sum32() uint32 {
	return h.crc
}
