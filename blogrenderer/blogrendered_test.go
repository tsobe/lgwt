package blogrenderer_test

import (
	"bytes"
	"github.com/approvals/go-approval-tests"
	"github.com/tsobe/lgwt/blogrenderer"
	"io"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		post = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	renderer, err := blogrenderer.NewRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("It converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		err = renderer.Render(&buf, post)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("It renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := renderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	b.ResetTimer()
	renderer, err := blogrenderer.NewRenderer()
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		renderer.Render(io.Discard, aPost)
	}
}
