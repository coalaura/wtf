package internal

import (
	"github.com/coalaura/wtf/types"
	"github.com/coalaura/wtf/types/custom"
)

func init() {
	types.RegisterFallback(types.DetectFunc(custom.DetectText))
}
