package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/ronanvirmani/event-management-system/backend/services"
)

// UploadFile handles file uploads to S3
func UploadFile(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(10 << 20) // Max file size: 10MB

    file, handler, err := r.FormFile("file")
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Error retrieving the file"))
        return
    }
    defer file.Close()

    url, err := services.UploadFileToS3(file, handler)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Error uploading file"))
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"url": url})
}
