package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
	"unicode/utf8"
)

func tulisJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func tulisError(w http.ResponseWriter, status int, pesan string) {
	tulisJSON(w, status, map[string]string{"error": pesan})
}

func metodeTidakDiizinkan(w http.ResponseWriter, r *http.Request) {
	tulisError(w, http.StatusMethodNotAllowed, "metode tidak diizinkan untuk path ini")
}

func tidakDitemukan(w http.ResponseWriter, r *http.Request) {
	tulisError(w, http.StatusNotFound, "path tidak ditemukan")
}

func ambilID(r *http.Request) (int, bool) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return 0, false
	}
	return id, true
}

func bacaBody(r *http.Request) []byte {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil
	}
	return data
}

func urai(data []byte, tujuan any) bool {
	dec := json.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(tujuan); err != nil {
		return false
	}
	if _, err := dec.Token(); err != io.EOF {
		return false
	}
	return true
}

func panjangTeks(s string) int {
	return utf8.RuneCountInString(s)
}

func tanggalValid(t string) bool {
	if len(t) != 10 {
		return false
	}
	_, err := time.Parse("2006-01-02", t)
	return err == nil
}
