package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"booknest-api/models"
	"booknest-api/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	DB *sql.DB
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.PasswordHash = string(hashedPassword)

	query := `INSERT INTO users (username, email, password_hash, role) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	err = h.DB.QueryRow(query, user.Username, user.Email, user.PasswordHash, user.Role).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var user models.User
	query := `SELECT id, username, email, password_hash, role FROM users WHERE email = $1`
	err = h.DB.QueryRow(query, credentials.Email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	if err != nil {
		http.Error(w, "Error logging in", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *UserHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	var user models.User
	query := `SELECT id, username, email, role, created_at, updated_at FROM users WHERE id = $1`
	err := h.DB.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Error fetching user profile", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
} 