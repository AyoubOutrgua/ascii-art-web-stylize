package functions

import (
	"net/http"
	"os"
)

func StyleFunc(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path
	file, err := os.Stat(filePath[1:])
	if err != nil || file.IsDir() {
		HandleError(w, "Not Found!", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath[1:])
}
