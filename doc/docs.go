package doc

import (
	"github.com/canadadry/pml/core"
)

type PrintDoc func(c core.Core) error

var Docs = map[string]PrintDoc{
	"doc1": doc1,
}
