package datasource

import (
	"testing"

	"github.com/TanLeYang/go-ssg/directoryhelper"
	"github.com/TanLeYang/go-ssg/models"
	"github.com/stretchr/testify/assert"
)

func TestLocalPostsDataSource_GetAllPosts(t *testing.T) {
	ds := LocalPostsDatasource{
		Directory: directoryhelper.RootDir() + "/test-data/posts",
	}

	expectedMeta := models.PostMeta{
		Title: "First post!",
		Short: "this is the very first post",
		Date:  "9/6/2021",
		Tags:  []string{"personal", "testing"},
	}

	posts, err := ds.GetAllPosts()
	assert.Nil(t, err, "expect no error when retrieving posts")
	assert.True(t, len(posts) == 1, "expect one post to be returned")
	assert.Equal(t, expectedMeta, posts[0].Meta, "expect post's meta to match expected meta")

	// just print out the content and manually inspect for now
	content := string(posts[0].Content)
	t.Logf("Post's content: %s", content)
}
