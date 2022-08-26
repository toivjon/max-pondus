package middleware

import (
	"log"
	"net/http"
	"regexp"
	"runtime/debug"
	"strings"

	"github.com/toivjon/max-pondus/backend/internal/server/common/contextkey"
	"github.com/toivjon/max-pondus/backend/internal/server/common/sort"
)

// Recoverer handles graceful logging and handling of panics.
func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				if rvr == http.ErrAbortHandler {
					panic(r)
				}
				stack := strings.Split(string(debug.Stack()), "\n")
				lines := []string{}
				for i := len(stack) - 1; i > 0; i-- {
					lines = append(lines, stack[i])
					if strings.HasPrefix(stack[i], "panic(") {
						lines = lines[0 : len(lines)-2]
						break
					}
				}
				sort.Reverse(lines)
				for i := range lines {
					lines[i] = strings.TrimSpace(lines[i])
					pattern := regexp.MustCompile("(.*).go:[0-9]+ ")
					match := pattern.FindString(lines[i])
					if match != "" {
						lines[i] = "    " + match
					}
				}
				reqID := r.Context().Value(contextkey.RequestID)
				log.Printf("%s panic: %s\n%s", reqID, rvr, strings.Join(lines, "\n"))
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}