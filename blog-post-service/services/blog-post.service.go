package services

import (
	"github.com/kylehipz/simpleng-blog-app/blog-post-service/store"
	"github.com/kylehipz/simpleng-blog-app/blog-post-service/types"
)

type BlogPostService struct {
	DB store.Database
}

func (s *BlogPostService) CreatePost(post *types.Blog) *types.Blog {
	newPost, err := s.DB.AddBlogPost(post)
	if err != nil {
		return newPost
	}

	return nil
}

func (s *BlogPostService) GetPost(id string) *types.Blog {
	post, err := s.DB.FindBlogPostById(id)
	if err != nil {
		return post
	}

	return nil
}
