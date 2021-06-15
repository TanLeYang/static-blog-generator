package parser

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	chromaHTML "github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

var style = styles.GitHub

func ReplaceCodeBlocks(htmlBytes []byte) ([]byte, error) {
	reader := bytes.NewReader(htmlBytes)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("error parsing html: %s", err.Error())
	}

	doc.Find("code[class*=\"language-\"]").Each(func(i int, s *goquery.Selection) {
		class, isLanguageDefined := s.Attr("class")
		language := strings.TrimPrefix(class, "language-")

		lexer := lexers.Get(language)
		if !isLanguageDefined || lexer == nil {
			lexer = lexers.Fallback
		}

		text := s.Text()
		iterator, err := lexer.Tokenise(nil, text)
		if err != nil {
			log.Printf("error when tokenising code block %s: %s \n", s.Text(), err.Error())
		}

		formatter := chromaHTML.New(chromaHTML.WithClasses(true))

		b := bytes.Buffer{}
		err = formatter.Format(&b, style, iterator)
		if err != nil {
			log.Printf("error during syntax highlighting of text %s: %s \n", text, err.Error())
		}

		s.SetHtml(b.String())
	})

	stringResult, err := doc.Html()
	if err != nil {
		return nil, fmt.Errorf("error when getting html result after syntax highlighting: %s", err.Error())
	}

	return []byte(stringResult), nil
}

func GetSyntaxHighlightCSS() (string, error) {
	writer := bytes.Buffer{}
	formatter := chromaHTML.New(chromaHTML.WithClasses(true))
	if err := formatter.WriteCSS(&writer, style); err != nil {
		return "", fmt.Errorf("error writing syntaxhighlighting css into document: %s", err.Error())
	}

	cssString := writer.String()
	return cssString, nil
}
