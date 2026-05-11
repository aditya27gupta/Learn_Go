package blogposts

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func assertPost(t *testing.T, got, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %+v, Want %+v", got, want)
	}
}

func TestNewBlogPosts(t *testing.T) {
	firstBody := `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
	secondBody := `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L`

	fs := fstest.MapFS{
		"hello_world.md":  {Data: []byte(firstBody)},
		"hello_world2.md": {Data: []byte(secondBody)},
	}

	posts, err := NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	got := posts[0]
	want := Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	}
	assertPost(t, got, want)
}
