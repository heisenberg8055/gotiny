package routes

import (
	"net/http"

	handlers "github.com/heisenberg8055/gotiny/internal/api/routes/handler"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func Routes(postClient *pgxpool.Pool, redisClient *redis.Client) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /shorten", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddURL(w, r, postClient, redisClient)
	})

	router.HandleFunc("GET /{shortUrl}", handlers.GetURL)

	return router
}
