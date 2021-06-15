package parser

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

func ParseMarkdownToHTML(source []byte) ([]byte, error) {
	md := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		return nil, err
	}

	htmlBytes := buf.Bytes()
	result, err := ReplaceCodeBlocks(htmlBytes)
	if err != nil {
		return nil, err
	}

	return result, nil
}
