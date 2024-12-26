package main

import (
	"encoding/json"
	"fmt"
	"go-server-tutorial/internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string
		Email string
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		fmt.Printf("Error parsing JSON: %v", err)
		
		respondWithError(w, http.StatusBadRequest, "Error parsing JSON")

		return
	}

	user, dbErr := app.storage.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		Username: params.Username,
		Email: params.Email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if dbErr != nil {
		fmt.Printf("Couldn't create user: %v", dbErr)

		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")

		return
	}

	respondWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}