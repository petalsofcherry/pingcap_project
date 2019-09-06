package model

import (
	"pingcap_project/constdef"
	"unsafe"
)

type OffsetArrayNode struct {
	pos  uint32
	next *OffsetArrayNode
}

type Hash2OffsetArray [5 * constdef.PerGB / unsafe.Sizeof(OffsetArrayNode{}) * 8]OffsetArrayNode
