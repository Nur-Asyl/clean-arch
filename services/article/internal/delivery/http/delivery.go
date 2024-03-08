package http

import (
	"architecture_go/services/article/configs"
	"architecture_go/services/article/internal/useCase"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type ArticleHTTPDelivery struct {
	articleUC useCase.ArticleUseCase
	commentUC useCase.CommentUseCase
}

func NewArticleHTTP(articleUC useCase.ArticleUseCase, commentUC useCase.CommentUseCase) *ArticleHTTPDelivery {
	return &ArticleHTTPDelivery{articleUC: articleUC, commentUC: commentUC}
}

func Trace(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next(w, r)

		log.Printf("%s %s %s %v\n", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	}
}

func (hd *ArticleHTTPDelivery) CreateArticleHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	var requestData struct {
		Name string `json:"name"`
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		log.Println("Error decoding request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newContact, err := hd.articleUC.CreateArticle(ctx, requestData.Name, requestData.Text)
	if err != nil {
		log.Println("Error creating article:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newContact)
}

func (hd *ArticleHTTPDelivery) ReadArticleHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	name := r.URL.Query().Get("name")

	existingContact, err := hd.articleUC.ReadArticle(ctx, name)
	if err != nil {
		log.Println("Error reading article:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingContact)
}

func (hd *ArticleHTTPDelivery) UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	var requestData struct {
		ArticleID int    `json:"article_id"`
		Name      string `json:"name"`
		Text      string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		log.Println("Error decoding request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := hd.articleUC.UpdateArticle(ctx, requestData.Name, requestData.Text)
	if err != nil {
		log.Println("Error updating article:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (hd *ArticleHTTPDelivery) DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	name := r.URL.Query().Get("name")

	err := hd.articleUC.DeleteArticle(ctx, name)
	if err != nil {
		log.Println("Error deleting article:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (hd *ArticleHTTPDelivery) CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	var requestData struct {
		UserID    int    `json:"user_id"`
		ArticleID int    `json:"article_id"`
		Text      string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		log.Println("Error decoding request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newGroup, err := hd.commentUC.CreateComment(ctx, requestData.UserID, requestData.ArticleID, requestData.Text)
	if err != nil {
		log.Println("Error creating comment:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newGroup)
}

func (hd *ArticleHTTPDelivery) ReadCommentHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		log.Println("Error parsing comment userID:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	articleID, err := strconv.Atoi(r.URL.Query().Get("article_id"))
	if err != nil {
		log.Println("Error parsing comment articleID:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	existingGroup, err := hd.commentUC.ReadComment(ctx, userID, articleID)
	if err != nil {
		log.Println("Error reading comment:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingGroup)
}

//func (hd *ArticleHTTPDelivery) AddCommentToArticleHandler(w http.ResponseWriter, r *http.Request) {
//	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
//	defer cancel()
//
//	var requestData struct {
//		ContactID int `json:"contact_id"`
//		GroupID   int `json:"group_id"`
//	}
//
//	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
//		log.Println("Error decoding request:", err)
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	err := hd.commentUC.AddContactToGroup(ctx, requestData.GroupID, requestData.ContactID)
//	if err != nil {
//		log.Println("Error adding article to comment:", err)
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusNoContent)
//}

func (hd *ArticleHTTPDelivery) Run(cfg *configs.Config) {
	addr := fmt.Sprintf(":%s", cfg.Port)

	mux := http.NewServeMux()

	mux.HandleFunc("/article/create", Trace(hd.CreateArticleHandler))
	mux.HandleFunc("/article/get", Trace(hd.ReadArticleHandler))
	mux.HandleFunc("/article/update", Trace(hd.UpdateArticleHandler))
	mux.HandleFunc("/article/delete", Trace(hd.DeleteArticleHandler))

	mux.HandleFunc("/comment/create", Trace(hd.CreateCommentHandler))
	mux.HandleFunc("/comment/get", Trace(hd.ReadCommentHandler))
	//mux.HandleFunc("/comment/addArticle", Trace(hd.AddCommentToArticleHandler))

	fmt.Println("Delivering... on port:", addr)
	go func() {
		err := http.ListenAndServe(addr, mux)
		if err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quitCh
	log.Println("Shutting down server...")
}
