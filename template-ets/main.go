package main

import (
	"log"
	"net/http"
	"sync"
)

type Server struct {
	mu         sync.Mutex
	sumberDaya []SumberDaya
	peminjaman []Peminjaman
	nextSDID   int
	nextPemID  int
}

func NewHandler() http.Handler {
	s := &Server{nextSDID: 1, nextPemID: 1}
	mux := http.NewServeMux()

	mux.HandleFunc("POST /sumber-daya", s.buatSumberDaya)
	mux.HandleFunc("GET /sumber-daya", s.daftarSumberDaya)
	mux.HandleFunc("GET /sumber-daya/{id}", s.detailSumberDaya)

	mux.HandleFunc("POST /peminjaman", s.buatPeminjaman)
	mux.HandleFunc("GET /peminjaman", s.daftarPeminjaman)
	mux.HandleFunc("GET /peminjaman/{id}", s.detailPeminjaman)
	for _, path := range []string{
		"/sumber-daya", "/sumber-daya/{id}",
		"/peminjaman", "/peminjaman/{id}",
	} {
		mux.HandleFunc(path, metodeTidakDiizinkan)
	}
	mux.HandleFunc("/", tidakDitemukan)

	return mux
}

func main() {
	log.Println("PinjamKampus API berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", NewHandler()))
}
