package store

import (
	"context"
	config "socialForumBackend/internal/config"
	db "socialForumBackend/internal/database"
	"socialForumBackend/internal/handler"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

const (
	insertPostQuery   = `INSERT into posts(id,content,author_id,anonymous) values($1,$2,$3,$4)`
	getAllPostQuery   = `SELECT id, content, author_id, anonymous FROM posts`
	deletePostQuery   = `DELETE FROM posts WHERE id = $1 AND author_id = $2`
	liveNewsFeedQuery = `SELECT id, content, author_id, anonymous FROM posts WHERE anonymous = false ORDER BY created_at DESC`
)

type Store interface {
	CreatePost(post *handler.Post) error
	DeletePost(postID uuid.UUID, userID int) error
	ListAllPosts(ctx context.Context) ([]*handler.Post, error)
	LiveNewsFeedQuery(ctx context.Context, userID int) ([]*handler.Post, error)
}

func NewStore() Store {
	return &dbStore{
		db: config.InitDB(),
	}
}

type dbStore struct {
	db *sqlx.DB
}

func (s *dbStore) CreatePost(post *handler.Post) error {
	if err := db.WithTimeout(context.Background(), config.Database.ReadTimeout, func(ctx context.Context) error {
		_, err := db.Get().MustExec(insertPostQuery, post.ID, post.Content, post.AuthorID, post.Anonymous).RowsAffected()
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return errors.Wrapf(err, "[CreatePost] failed to insert into db")
	}
	return nil
}

func (s *dbStore) ListAllPosts(ctx context.Context) ([]*handler.Post, error) {
	var posts []*handler.Post

	if err := db.WithTimeout(ctx, config.Database.ReadTimeout, func(ctx context.Context) error {
		err := db.Get().SelectContext(ctx, &posts, getAllPostQuery)
		return err
	}); err != nil {
		return nil, errors.Wrapf(err, "[ListAllPosts] Failed to get active providers from DB")
	}
	return posts, nil
}

func (s *dbStore) DeletePost(postID uuid.UUID, userID int) error {
	if err := db.WithTimeout(context.Background(), config.Database.ReadTimeout, func(ctx context.Context) error {
		_, err := db.Get().MustExec(deletePostQuery, postID, userID).RowsAffected()
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return errors.Wrapf(err, "[DeletePost] failed to insert into db")
	}
	return nil
}

func (s *dbStore) LiveNewsFeedQuery(ctx context.Context, userID int) ([]*handler.Post, error) {
	var posts []*handler.Post

	if err := db.WithTimeout(ctx, config.Database.ReadTimeout, func(ctx context.Context) error {
		err := db.Get().SelectContext(ctx, &posts, liveNewsFeedQuery, userID)
		return err
	}); err != nil {
		return nil, errors.Wrapf(err, "[LiveNewsFeedQuery] Failed to get active providers from DB")
	}
	return posts, nil
}
