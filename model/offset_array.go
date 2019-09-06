package model

import (
	"pingcap_project/constdef"
)

type OffsetArray [5 * constdef.PerGB / 64]*string
