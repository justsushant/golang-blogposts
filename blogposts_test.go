package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/justsushant/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go`

		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker`
	)



	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	assertPost(t, posts[0], blogposts.Post{
		Title: "Post 1",
		Description: "Description 1",
		Tags: []string {"tdd", "go"},
	})

	assertPost(t, posts[1], blogposts.Post{
		Title: "Post 2",
		Description: "Description 2",
		Tags: []string {"rust", "borrow-checker"},
	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}