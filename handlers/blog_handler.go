package handlers

import (
	"blogApp/middleware"
	"blogApp/models"
	"blogApp/services"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type PostHandler struct {
	Service *services.PostService
}



func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
    var post models.Post
    err := json.NewDecoder(r.Body).Decode(&post)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	
	userID, err := middleware.ExtractUserID(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }
    user_uuid, err := uuid.Parse(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    post.AuthorID = user_uuid


	//call service layer
	err = h.Service.CreatePost(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

type UpdatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req UpdatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedPost, err := h.Service.UpdatePost(id, req.Title, req.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedPost)
}

func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	post, err := h.Service.ViewPostByID(id)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Service.GetAllPosts()
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.Service.DeletePostByID(id)
	if err != nil {
		http.Error(w, "Failed to delete post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) 
}