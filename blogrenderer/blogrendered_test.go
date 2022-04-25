package blogrenderer_test

import (
	"bytes"
	"github.com/approvals/go-approval-tests"
	"github.com/tsobe/lgwt/blogrenderer"
	"testing"
)

func TestReader(t *testing.T) {
	var (
		post = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("It converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, post)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
