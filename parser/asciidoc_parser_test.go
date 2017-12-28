package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func verify(t GinkgoTInterface, expectedDocument interface{}, content string, options ...parser.Option) {
	log.Debugf("processing: %s", content)
	reader := strings.NewReader(content)
	result, err := parser.ParseReader("", reader, options...) //, Debug(true))
	if err != nil {
		log.WithError(err).Error("Error found while parsing the document")
	}
	require.Nil(t, err)
	t.Logf("actual document: `%s`", spew.Sdump(result))
	t.Logf("expected document: `%s`", spew.Sdump(expectedDocument))
	assert.EqualValues(t, expectedDocument, result)
}
