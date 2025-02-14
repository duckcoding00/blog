package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/duckcoding00/blog/internal/model"
	"github.com/duckcoding00/blog/internal/service"
	"github.com/gorilla/mux"
)

type ErrResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type BlogHandler struct {
	service service.Service
}

func (h *BlogHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Cek method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "method not allowed",
		})
		return
	}

	// Decode request body
	var payload model.BlogPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "invalid request body: " + err.Error(),
		})
		return
	}
	defer r.Body.Close()

	// Validasi payload
	if err := payload.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: err.Error(),
		})
		return
	}

	// Proses create blog
	data, err := h.service.Blog.Create(r.Context(), payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "failed to create blog: " + err.Error(),
		})
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "blog created successfully",
		Data:    data,
	})
}

func (h *BlogHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Cek method
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "method not allowed",
		})
		return
	}

	var payload model.BlogUpdatePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "invalid request body: " + err.Error(),
		})
		return
	}
	defer r.Body.Close()

	// Validasi payload
	if err := payload.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: err.Error(),
		})
		return
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: err.Error(),
		})
		return
	}

	// Proses create blog
	data, err := h.service.Blog.Update(r.Context(), int32(id), payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "failed to create blog: " + err.Error(),
		})
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "blog updated successfully",
		Data:    data,
	})
}

func (h *BlogHandler) GetBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Cek method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "method not allowed",
		})
		return
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: err.Error(),
		})
		return
	}

	// Proses create blog
	data, err := h.service.Blog.GetBlog(r.Context(), int32(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "failed to create blog: " + err.Error(),
		})
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "get blog successfully",
		Data:    data,
	})
}

func (h *BlogHandler) GetBlogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Cek method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "method not allowed",
		})
		return
	}

	// Proses create blog
	data, err := h.service.Blog.GetBlogs(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "failed to create blog: " + err.Error(),
		})
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "get blogs successfully",
		Data:    data,
	})
}

func (h *BlogHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Cek method
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "method not allowed",
		})
		return
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: err.Error(),
		})
		return
	}

	// Proses create blog
	if err := h.service.Blog.Delete(r.Context(), int32(id)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "failed to delete blog: " + err.Error(),
		})
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "blog deleted successfully",
	})
}
