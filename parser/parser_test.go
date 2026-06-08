package parser

import (
	"reflect"
	"strings"
	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		name string
		html string
		want []string
	}{
		{
			name: "single link",
			html: `<a href="https://example.com">x</a>`,
			want: []string{"https://example.com"},
		},
		{
			name: "multiple links",
			html: `<a href="/a">a</a><p>text</p><a href="/b">b</a>`,
			want: []string{"/a", "/b"},
		},
		{
			name: "nested link",
			html: `<div><ul><li><a href="/deep">x</a></li></ul></div>`,
			want: []string{"/deep"},
		},
		{
			name: "no links",
			html: `<p>just text</p>`,
			want: nil,
		},
		{
			name: "a tag without href",
			html: `<a name="anchor">x</a>`,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Extract(strings.NewReader(tt.html))
			if err != nil {
				t.Fatalf("Extract returned error: %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Extract() = %v, want %v", got, tt.want)
			}
		})
	}
}
