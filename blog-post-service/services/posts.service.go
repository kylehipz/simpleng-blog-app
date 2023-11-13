package services

import (
	"github.com/kylehipz/simpleng-blog-app/blog-post-service/store"
	"github.com/kylehipz/simpleng-blog-app/blog-post-service/types"
)

type PostsService struct {
	DB store.Database
}

func (s *PostsService) CreatePost(post *types.Blog) *types.Blog {
	newPost, err := s.DB.AddBlogPost(post)
	if err != nil {
		return newPost
	}

	return nil
}

func (s *PostsService) GetPost(id string) *types.Blog {
	post, err := s.DB.FindBlogPostById(id)
	if err != nil {
		return post
	}

	return nil
}
