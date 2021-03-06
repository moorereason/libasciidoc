package html5

import (
	"bytes"
	"html/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var blockImageTmpl template.Template
var inlineImageTmpl template.Template

// initializes the templates
func init() {
	blockImageTmpl = newHTMLTemplate("block image", `<div{{ if ne .ID "" }} id="{{ .ID }}"{{ end }} class="imageblock">
<div class="content">
{{ if ne .Link "" }}<a class="image" href="{{ .Link }}">{{ end}}<img src="{{ .Macro.Path }}" alt="{{ .Macro.Alt }}"{{ if .Macro.Width }} width="{{ .Macro.Width }}"{{ end }}{{ if .Macro.Height }} height="{{ .Macro.Height }}"{{ end }}>{{ if ne .Link "" }}</a>{{ end }}
</div>{{ if ne .Title "" }}
<div class="doctitle">{{ .Title }}</div>
{{ else }}
{{ end }}</div>`)
	inlineImageTmpl = newHTMLTemplate("inline image", `<span class="image"><img src="{{.Macro.Path}}" alt="{{.Macro.Alt}}"{{if .Macro.Width}} width="{{.Macro.Width}}"{{end}}{{if .Macro.Height}} height="{{.Macro.Height}}"{{end}}></span>`)
}

func renderBlockImage(ctx *renderer.Context, img types.BlockImage) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	err := blockImageTmpl.Execute(result, struct {
		ID    string
		Title string
		Link  string
		Macro types.ImageMacro
	}{
		ID:    img.ID.Value,
		Title: img.Title.Value,
		Link:  img.Link.Path,
		Macro: img.Macro,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render block image")
	}
	log.Debugf("rendered block image: %s", result.Bytes())
	return result.Bytes(), nil
}

func renderInlineImage(ctx *renderer.Context, img types.InlineImage) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	err := inlineImageTmpl.Execute(result, img)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render inline image")
	}
	log.Debugf("rendered inline image: %s", result.Bytes())
	return result.Bytes(), nil
}
