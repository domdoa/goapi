package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code int
	Balance int64
}

type Error struct {
	Code int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code: code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

// Wrappers for error handlers
var (
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter){
		writeError(w, "An unexpected error has occurred", http.StatusInternalServerError)
	}
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params CoinBalanceParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} 

	balance := CoinBalanceResponse{
		Code: 200,
		Balance: 100,
	}

	json.NewEncoder(w).Encode(balance)
}
