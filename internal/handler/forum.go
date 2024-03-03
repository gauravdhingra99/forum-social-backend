package handler

import "socialForumBackend/service"

type ForumHandler struct {
	postService service.PostInterFace
}

func NewForumHandler(postService service.PostInterFace) *ForumHandler {
	return &ForumHandler{
		postService: postService,
	}
}
