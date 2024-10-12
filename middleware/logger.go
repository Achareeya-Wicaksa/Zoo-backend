package middleware

import (
    "log"
    "net/http"
    "time"
)

// LoggerMiddleware mencatat setiap permintaan ke terminal
func LoggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Mengizinkan permintaan untuk dilanjutkan
        next.ServeHTTP(w, r)

        // Mencatat informasi permintaan
        log.Printf("%s %s took %v", r.Method, r.URL.Path, time.Since(start))
    })
}
