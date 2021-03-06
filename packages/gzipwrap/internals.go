package gzipwrap

import (
	"compress/gzip"
	"net/http"
	"sync"
)

type gzipResponseWriter struct {
	http.ResponseWriter

	w             *gzip.Writer
	statusCode    int
	headerWritten bool
}

var (
	pool = sync.Pool{
		New: func() interface{} {
			w, _ := gzip.NewWriterLevel(nil, gzip.BestCompression)
			return &gzipResponseWriter{
				w: w,
			}
		},
	}
)

func (gzr *gzipResponseWriter) WriteHeader(statusCode int) {
	gzr.statusCode = statusCode
	gzr.headerWritten = true

	if gzr.statusCode != http.StatusNotModified && gzr.statusCode != http.StatusNoContent {
		gzr.ResponseWriter.Header().Del("Content-Length")
		gzr.ResponseWriter.Header().Set("Content-Encoding", "gzip")
	}

	gzr.ResponseWriter.WriteHeader(statusCode)
}

func (gzr *gzipResponseWriter) Write(b []byte) (int, error) {
	if _, ok := gzr.Header()["Content-Type"]; !ok {
		// If no content type, apply sniffing algorithm to un-gzipped body.
		gzr.ResponseWriter.Header().Set("Content-Type", http.DetectContentType(b))
	}

	if !gzr.headerWritten {
		// This is exactly what Go would also do if it hasn't been written yet.
		gzr.WriteHeader(http.StatusOK)
	}

	return gzr.w.Write(b)
}

func (gzr *gzipResponseWriter) Flush() {
	if gzr.w != nil {
		gzr.w.Flush()
	}

	if fw, ok := gzr.ResponseWriter.(http.Flusher); ok {
		fw.Flush()
	}
}
