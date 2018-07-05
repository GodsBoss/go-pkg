package typedconf_test

import (
	"github.com/GodsBoss/go-pkg/typedconf"

	"encoding/json"
	"fmt"
	"strings"
)

func Example() {
	input := []byte(
		`
			{
				"sources": [
					{
						"type": "inline",
						"content": "Hello, world!"
					},
					{
						"type": "replace",
            "pattern": "$a $b $c",
            "replace": {
              "$a": "foo",
              "$b": "bar",
              "$c": "baz"
            }
					}
				]
			}
		`,
	)
	cfg := &Config{}
	err := json.Unmarshal(input, cfg)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return
	}
	for _, source := range cfg.Sources {
		txt, _ := source.Text()
		fmt.Println(txt)
	}

	// Output:
	// Hello, world!
	// foo bar baz
}

type Config struct {
	Sources TextSources `json:"sources"`
}

type TextSources []TextSource

func (sources *TextSources) UnmarshalJSON(data []byte) error {
	decoders := typedconf.NewDecoders(json.Unmarshal)
	decoders.Register("inline", func() interface{} {
		return &InlineSource{}
	})
	decoders.Register("replace", func() interface{} {
		return &ReplacerSource{}
	})
	list := []struct{}{}
	err := json.Unmarshal(data, &list)
	if err != nil {
		return err
	}
	decoded := make([]typedconf.Instance, len(list))
	for index := range list {
		decoded[index] = decoders.Instance()
	}
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		return err
	}
	*sources = make(TextSources, len(list))
	for index, dec := range decoded {
		(*sources)[index] = dec.Value().(TextSource)
	}
	return nil
}

type TextSource interface {
	Text() (string, error)
}

type InlineSource struct {
	Content string `json:"content"`
}

func (source *InlineSource) Text() (string, error) {
	return source.Content, nil
}

type ReplacerSource struct {
	Pattern      string            `json:"pattern"`
	Replacements map[string]string `json:"replace"`
}

func (source ReplacerSource) Text() (string, error) {
	content := source.Pattern
	for o, n := range source.Replacements {
		content = strings.Replace(content, o, n, -1)
	}
	return content, nil
}
