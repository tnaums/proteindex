package dex

import 	"github.com/tnaums/proteindex/internal/proteinapi"

type Protein struct {
	Name string
	Blast proteinapi.Blastp
}

func NewDex() map[string]Protein {
	return make(map[string]Protein)
}
