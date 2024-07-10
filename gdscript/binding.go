package gdscript

//#include "tree_sitter/parser.h"
//TSLanguage *tree_sitter_gdscript();
import "C"

import (
	"unsafe"

	sitter "github.com/smacker/go-tree-sitter"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_gdscript())
	return sitter.NewLanguage(ptr)
}
