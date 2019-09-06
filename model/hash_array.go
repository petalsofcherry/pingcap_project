package model

import (
	"fmt"
	"pingcap_project/constdef"
	"pingcap_project/util"
)

type HashArray [5 * constdef.PerGB / constdef.HashSize][constdef.HashSize]byte

func (h HashArray) Append(x interface{}) error {
	var byteArray [constdef.HashSize]byte
	var ok bool
	if byteArray, ok = x.([constdef.HashSize]byte); !ok {
		return fmt.Errorf("x is not []byte")
	}

	if len(h) == cap(h) {
		util.AppendFileArrayItem("temp_files/hash_array", string(byteArray[:]))
	}
	h[len(h)] = byteArray


	return nil
}
