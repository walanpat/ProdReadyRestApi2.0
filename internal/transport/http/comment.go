package http

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go-rest-api/internal/comment"
	"log"
	"net/http"
)

type Response struct {
	Message string
}
type PostCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body",validate:"required"`
}
type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	GetComment(ctx context.Context, ID string) (comment.Comment, error)
	UpdateComment(ctx context.Context, ID string, newcmt comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, ID string) error
}

func convertPostCommentRequestToComment(c PostCommentRequest) comment.Comment {
	return comment.Comment{
		Slug:   c.Slug,
		Author: c.Author,
		Body:   c.Body,
	}

}
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var cmt PostCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}
	validate := validator.New()
	err := validate.Struct(cmt)
	if err != nil {
		http.Error(w, "not a valid comment", http.StatusBadRequest)
	}
	convertedComment := convertPostCommentRequestToComment(cmt)
	postedComment, err := h.Service.PostComment(r.Context(), convertedComment)
	if err != nil {
		log.Print(err)
	}

	if err := json.NewEncoder(w).Encode(postedComment); err != nil {
		panic(err)
	}
}
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}

}
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		panic(err)
	}
	cmt, err := h.Service.UpdateComment(r.Context(), id, cmt)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}

}
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := h.Service.DeleteComment(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully Deleted"}); err != nil {
		panic(err)
	}
}
func (h *Handler) GetCommentByNumber(w http.ResponseWriter, r *http.Request) {

}
