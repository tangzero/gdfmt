package gdscript_test

import (
	"context"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/stretchr/testify/assert"
	"github.com/tangzero/gdfmt/gdscript"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("extends Node2D"), gdscript.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(source (extends_statement (type (identifier))))",
		n.String(),
	)
}
