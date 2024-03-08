package http

import (
	"architecture_go/services/article/configs"
	"architecture_go/services/user/internal/useCase"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/justinas/alice"
	"github.com/justinas/nosurf"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type contextKey string

const isAuthenticatedContextKey = contextKey("isAuthenticated")

type UserHTTPDelivery struct {
	userUC         useCase.UserUseCase
	db             *sql.DB
	sessionManager *scs.SessionManager
}

func NewUserHTTP(userUC useCase.UserUseCase, db *sql.DB, sessionManager *scs.SessionManager) *UserHTTPDelivery {
	return &UserHTTPDelivery{userUC: userUC, db: db, sessionManager: sessionManager}
}

func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})

	return csrfHandler
}

func (hd *UserHTTPDelivery) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/user/signup" {
			next.ServeHTTP(w, r)
			return
		}
		if r.URL.Path == "/user/login" {
			next.ServeHTTP(w, r)
			return
		}
		email := hd.sessionManager.GetString(r.Context(), "authenticatedUserEmail")
		if email == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		exists, err := hd.userUC.ReadUser(r.Context(), email)
		if err != nil {
			log.Println("Error reading user:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if exists != nil {
			ctx := context.WithValue(r.Context(), isAuthenticatedContextKey, true)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func (hd *UserHTTPDelivery) requireAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !hd.isAuthenticated(r) {
			log.Println("Error not authenticated user")
			http.Error(w, "Error not authenticated user", http.StatusProxyAuthRequired)
			return
		}

		w.Header().Add("Cache-Control", "no-store")

		next.ServeHTTP(w, r)
	}
}

func (hd *UserHTTPDelivery) isAuthenticated(r *http.Request) bool {
	isAuthenticated, ok := r.Context().Value(isAuthenticatedContextKey).(bool)
	if !ok {
		return false
	}

	return isAuthenticated
}
func (hd *UserHTTPDelivery) userSignup(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	var requestData struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		log.Println("Error decoding request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser, err := hd.userUC.CreateUser(ctx, requestData.FirstName, requestData.LastName, requestData.Email, requestData.Password)
	if err != nil {
		log.Println("Error creating user:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hd.sessionManager.Put(r.Context(), "authenticatedUserEmail", newUser.Email)
	hd.sessionManager.Put(r.Context(), "flash", "Your signup was successful. Please log in.")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func (hd *UserHTTPDelivery) userLogin(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	var requestData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		log.Println("Error decoding request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	email := requestData.Email

	existingUser, err := hd.userUC.ReadUser(ctx, email)
	if err != nil {
		log.Println("Error reading user:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if existingUser == nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = hd.sessionManager.RenewToken(r.Context())
	if err != nil {
		log.Println("Error renew token user:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hd.sessionManager.Put(r.Context(), "authenticatedUserEmail", existingUser.Email)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingUser)
}

func (hd *UserHTTPDelivery) userLogout(w http.ResponseWriter, r *http.Request) {
	err := hd.sessionManager.RenewToken(r.Context())
	if err != nil {
		log.Println("Error renew token user:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hd.sessionManager.Remove(r.Context(), "authenticatedUserEmail")

	hd.sessionManager.Put(r.Context(), "flash", "You've been logged out successfully!")

	log.Println("User logout")
}

func (hd *UserHTTPDelivery) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	var requestData struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		log.Println("Error decoding request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser, err := hd.userUC.CreateUser(ctx, requestData.FirstName, requestData.LastName, requestData.Email, requestData.Password)
	if err != nil {
		log.Println("Error creating user:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func (hd *UserHTTPDelivery) ReadUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	email := r.URL.Query().Get("email")

	existingUser, err := hd.userUC.ReadUser(ctx, email)
	if err != nil {
		log.Println("Error reading user:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingUser)
}

func (hd *UserHTTPDelivery) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	var requestData struct {
		UserID    int    `json:"user_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		log.Println("Error decoding request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := hd.userUC.UpdateUser(ctx, requestData.FirstName, requestData.LastName, requestData.Email, requestData.Password)
	if err != nil {
		log.Println("Error updating user:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (hd *UserHTTPDelivery) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	email := r.URL.Query().Get("email")

	err := hd.userUC.DeleteUser(ctx, email)
	if err != nil {
		log.Println("Error deleting user:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (hd *UserHTTPDelivery) Run(cfg *configs.Config) {
	addr := fmt.Sprintf(":%s", cfg.Port)

	mux := http.NewServeMux()

	authMiddleware := alice.New(hd.sessionManager.LoadAndSave, noSurf, hd.authenticate)

	mux.Handle("/user/create", authMiddleware.ThenFunc(hd.CreateUserHandler))
	mux.Handle("/user/update", authMiddleware.ThenFunc(hd.UpdateUserHandler))
	mux.Handle("/user/delete", authMiddleware.ThenFunc(hd.DeleteUserHandler))

	mux.Handle("/user/signup", authMiddleware.ThenFunc(hd.userSignup))
	mux.Handle("/user/login", authMiddleware.ThenFunc(hd.userLogin))
	mux.HandleFunc("/user/logout", hd.requireAuthentication(hd.userLogout))

	fmt.Println("User Service Delivering on port:", addr)
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
