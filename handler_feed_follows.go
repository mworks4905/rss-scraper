package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/mworks4905/rss-scraper/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error creating feed follow: %v", err))
		return
	}

	responseWithJson(w, 201, DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		UserId uuid.UUID `json:"user_id"`
	}

	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollows, err := apiCfg.DB.GetFeedFollow(r.Context(), params.UserId)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error getting feed follows: %v", err))
		return
	}

	responseWithJson(w, 201, DatabaseFeedFollowstoFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdStr := chi.URLParam(r, "feedFollowId")
	if feedFollowIdStr == "" {
		responseWithError(w, 403, "Error: must provide feed follow id.")
	}

	feedFollowId, err := uuid.Parse(feedFollowIdStr)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing id string: %v", err))
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error deleting feed follow: %v", err))
		return
	}

	responseWithJson(w, 200, struct{}{})
}
