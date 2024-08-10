package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// getVideos retrieves video data with pagination and sorting by published datetime.
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

// searchVideos retrieves video data based on a search query, with partial matching in title and description.
func searchVideos(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	var videos []Video
	result := db.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Order("published_at desc").Find(&videos)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videos)
}
