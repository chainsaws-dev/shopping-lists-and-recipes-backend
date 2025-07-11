package gzipwrap

import (
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"
)

var HostURL = "http://localhost:8080/"

func TestNoGzip(t *testing.T) {
	req, err := http.NewRequest("GET", HostURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	MakeGzipHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})(rec, req)

	if rec.Code != 200 {
		t.Fatalf("expected 200 got %d", rec.Code)
	}

	if req.Header.Get("Content-Encoding") != "" {
		t.Fatalf(`expected Content-Encoding: "" got %s`, req.Header.Get("Content-Encoding"))
	}

	if rec.Body.String() != "test" {
		t.Fatalf(`expected "test" go "%s"`, rec.Body.String())
	}

	if testing.Verbose() {
		b, _ := httputil.DumpResponse(rec.Result(), true)
		t.Log("\n" + string(b))
	}
}

func TestGzip(t *testing.T) {
	req, err := http.NewRequest("GET", HostURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Accept-Encoding", "gzip, deflate")

	rec := httptest.NewRecorder()
	MakeGzipHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "4")
		w.Header().Set("Content-Type", "text/test")
		w.Write([]byte("test"))
	})(rec, req)

	if rec.Code != 200 {
		t.Fatalf("expected 200 got %d", rec.Code)
	}

	if req.Header.Get("Content-Encoding") != "gzip" {
		t.Fatalf("expected Content-Encoding: gzip got %s", req.Header.Get("Content-Encoding"))
	}
	if req.Header.Get("Content-Length") != "" {
		t.Fatalf(`expected Content-Length: "" got %s`, req.Header.Get("Content-Length"))
	}
	if req.Header.Get("Content-Type") != "text/test" {
		t.Fatalf(`expected Content-Type: "text/test" got %s`, req.Header.Get("Content-Type"))
	}

	r, err := gzip.NewReader(rec.Body)
	if err != nil {
		t.Fatal(err)
	}

	body, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "test" {
		t.Fatalf(`expected "test" go "%s"`, string(body))
	}

	if testing.Verbose() {
		b, _ := httputil.DumpResponse(rec.Result(), true)
		t.Log("\n" + string(b))
	}
}

func TestNoBody(t *testing.T) {
	req, err := http.NewRequest("GET", HostURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Accept-Encoding", "gzip, deflate")

	rec := httptest.NewRecorder()
	MakeGzipHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected %d got %d", http.StatusNoContent, rec.Code)
	}

	if req.Header.Get("Content-Encoding") != "" {
		t.Fatalf(`expected Content-Encoding: "" got %s`, req.Header.Get("Content-Encoding"))
	}

	if rec.Body.Len() > 0 {
		t.Logf("%q", rec.Body.String())
		t.Fatalf("no body expected for %d bytes", rec.Body.Len())
	}

	if testing.Verbose() {
		b, _ := httputil.DumpResponse(rec.Result(), true)
		t.Log("\n" + string(b))
	}
}

func BenchmarkGzip(b *testing.B) {
	body := []byte("testtesttesttesttesttesttesttesttesttesttesttesttest")

	req, err := http.NewRequest("GET", HostURL, nil)
	if err != nil {
		b.Fatal(err)
	}
	req.Header.Set("Accept-Encoding", "gzip, deflate")

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rec := httptest.NewRecorder()
			MakeGzipHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write(body)
			})(rec, req)

			if rec.Code != http.StatusOK {
				b.Fatalf("expected %d got %d", http.StatusOK, rec.Code)
			}

		}
	})
}
