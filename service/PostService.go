package service

import (
	"context"
	"socialForumBackend/internal/handler"
	"socialForumBackend/internal/store"

	uuid "github.com/satori/go.uuid"
)

type PostInterFace interface {
	CreatePost(post *handler.Post) (uuid.UUID, error)
	DeletePost(postID uuid.UUID, userID int) error
	ListAllPosts(ctx context.Context) ([]*handler.Post, error)
	LiveNewsFeedQuery(ctx context.Context, userID int) ([]*handler.Post, error)
}

type postService struct {
	store store.Store
}

func NewPostService(store store.Store) PostInterFace {
	return &postService{store: store}
}

func (p *postService) CreatePost(post *handler.Post) (uuid.UUID, error) {
	post.ID = uuid.NewV4()
	err := p.store.CreatePost(post)
	if err != nil {
		return post.ID, err
	}
	return post.ID, nil
}

func (p *postService) DeletePost(postID uuid.UUID, userID int) error {
	err := p.store.DeletePost(postID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (p *postService) ListAllPosts(ctx context.Context) ([]*handler.Post, error) {
	posts, err := p.store.ListAllPosts(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *postService) LiveNewsFeedQuery(ctx context.Context, userID int) ([]*handler.Post, error) {
	posts, err := p.store.LiveNewsFeedQuery(ctx, userID)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
