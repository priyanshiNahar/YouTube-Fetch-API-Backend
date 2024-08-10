package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getVideos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	limit, _ := strconv.Atoi(vars["limit"])

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	var videos []Video
	offset := (page - 1) * limit

	result := db.Order("published_at desc").Limit(limit).Offset(offset).Find(&videos)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videos)
}
