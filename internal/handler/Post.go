package handler

import (
	"encoding/json"
	"net/http"
	"socialForumBackend/internal/models"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Post struct {
	ID         uuid.UUID `db:"id"`
	Content    string    `db:"content"`
	AuthorID   int       `db:"author_id"`
	AuthorName string    `db:"author_name"`
	Anonymous  bool      `db:"anonymous"`
}

type Response struct {
	Status   int     `json:"status"`
	Message  string  `json:"message"`
	PostData []*Post `json:"postData"`
}

func (forumHandler *ForumHandler) CreatePost(writer http.ResponseWriter, request *http.Request) {
	messageReq, err := parseMessageRequestBody(request)
	if err != nil {
		resp := models.NewErrorResponse(models.ErrInvalidRequestBody, "failed to unmarshal request body", "invalid request body")
		resp.Write(writer, http.StatusBadRequest)
		return
	}
	id, err := forumHandler.postService.CreatePost(&messageReq)
	if err != nil {
		resp := models.NewErrorResponse(models.ErrInternalServerError, "failed to publish message", "internal server error")
		resp.Write(writer, http.StatusInternalServerError)
		return
	}

	postData := []*Post{{ID: id}}
	data := Response{Status: http.StatusOK, Message: "Post Successfully Created", PostData: postData}
	resp := models.NewSuccessResponse(data)
	resp.Write(writer, http.StatusOK)
	return
}

func (forumHandler *ForumHandler) DeletePost(writer http.ResponseWriter, request *http.Request) {
	messageReq, err := parseMessageRequestBody(request)
	if err != nil {
		resp := models.NewErrorResponse(models.ErrInvalidRequestBody, "failed to unmarshal request body", "invalid request body")
		resp.Write(writer, http.StatusBadRequest)
		return
	}
	err = forumHandler.postService.DeletePost(messageReq.ID, messageReq.AuthorID)
	if err != nil {
		resp := models.NewErrorResponse(models.ErrInternalServerError, "failed to publish message", "internal server error")
		resp.Write(writer, http.StatusInternalServerError)
		return
	}
	data := Response{Status: http.StatusOK, Message: "Post Successfully Deleted"}
	resp := models.NewSuccessResponse(data)
	resp.Write(writer, http.StatusOK)
	return
}

func (forumHandler *ForumHandler) ListAllPosts(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	userID, _ := strconv.Atoi(mux.Vars(request)["user_id"])
	posts, err := forumHandler.postService.ListAllPosts(ctx, userID)
	if err != nil {
		resp := models.NewErrorResponse(models.ErrInternalServerError, "failed to publish message", "internal server error")
		resp.Write(writer, http.StatusInternalServerError)
		return
	}
	data := Response{Status: http.StatusOK, PostData: posts}
	resp := models.NewSuccessResponse(data)
	resp.Write(writer, http.StatusOK)
	return
}

func (forumHandler *ForumHandler) LiveNewsFeedQuery(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	posts, err := forumHandler.postService.LiveNewsFeedQuery(ctx)
	if err != nil {
		resp := models.NewErrorResponse(models.ErrInternalServerError, "failed to publish message", "internal server error")
		resp.Write(writer, http.StatusInternalServerError)
		return
	}
	data := Response{Status: http.StatusOK, PostData: posts}
	resp := models.NewSuccessResponse(data)
	resp.Write(writer, http.StatusOK)
	return
}

func parseMessageRequestBody(request *http.Request) (Post, error) {
	var req Post
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		return Post{}, errors.Wrapf(err, "failed to unmarshal request body")
	}
	return req, nil
}
