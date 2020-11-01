package gzipwrap

import (
	"log"
	"net/http"
	"strings"
)

// MakeGzipHandlerFunc - добавляет поддержку gzip сжатия в HandlerFunc
func MakeGzipHandlerFunc(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fn(w, r)
			return
		}

		gzr := pool.Get().(*gzipResponseWriter)
		gzr.statusCode = 0
		gzr.headerWritten = false
		gzr.ResponseWriter = w
		gzr.w.Reset(w)

		defer func() {
			// gzr.w.Close will write a footer even if no data has been written.
			// StatusNotModified and StatusNoContent expect an empty body so don't close it.
			if gzr.statusCode != http.StatusNotModified && gzr.statusCode != http.StatusNoContent {
				if err := gzr.w.Close(); err != nil {
					log.Printf("[MakeGzipHandlerFunc] %v", err)
				}
			}
			pool.Put(gzr)
		}()

		fn(gzr, r)
	}
}

// MakeGzipHandler - добавляет поддержку gzip сжатия в Handler
func MakeGzipHandler(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			h.ServeHTTP(w, r)
			return
		}

		gzr := pool.Get().(*gzipResponseWriter)
		gzr.statusCode = 0
		gzr.headerWritten = false
		gzr.ResponseWriter = w
		gzr.w.Reset(w)

		defer func() {
			// gzr.w.Close will write a footer even if no data has been written.
			// StatusNotModified and StatusNoContent expect an empty body so don't close it.
			if gzr.statusCode != http.StatusNotModified && gzr.statusCode != http.StatusNoContent {
				if err := gzr.w.Close(); err != nil {
					log.Printf("[MakeGzipHandler] %v", err)
				}
			}
			pool.Put(gzr)
		}()

		h.ServeHTTP(gzr, r)
	})

}
