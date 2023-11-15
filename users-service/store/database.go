package store

import "github.com/kylehipz/simpleng-blog-app/blog-post-service/types"

type Database interface {
	AddBlogPost(post *types.Blog) (*types.Blog, error)
	FindBlogPostById(id string) (*types.Blog, error)
}
