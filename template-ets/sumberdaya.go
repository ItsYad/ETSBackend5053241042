package main

import (
	"net/http"
	"strings"
)

func validasiSumberDaya(in inputSumberDaya) (SumberDaya, string) {
	var sd SumberDaya
	if in.Nama == nil || in.Jenis == nil || in.Stok == nil {
		return sd, "nama, jenis, dan stok wajib diisi"
	}
	nama := strings.TrimSpace(*in.Nama)
	jenis := strings.TrimSpace(*in.Jenis)
	if nama == "" {
		return sd, "nama tidak boleh kosong"
	}
	if panjangTeks(nama) > 50 {
		return sd, "nama maksimal 50 karakter"
	}
	if jenis != "ruang" && jenis != "alat" {
		return sd, "jenis harus 'ruang' atau 'alat'"
	}
	if *in.Stok < 1 {
		return sd, "stok minimal 1"
	}
	sd.Nama = nama
	sd.Jenis = jenis
	sd.Stok = *in.Stok
	return sd, ""
}

func (s *Server) cariSumberDaya(id int) (SumberDaya, bool) {
	for _, sd := range s.sumberDaya {
		if sd.ID == id {
			return sd, true
		}
	}
	return SumberDaya{}, false
}

func (s *Server) buatSumberDaya(w http.ResponseWriter, r *http.Request) {
	body := bacaBody(r)
	var in inputSumberDaya
	if !urai(body, &in) {
		tulisError(w, http.StatusBadRequest, "body bukan JSON yang valid")
		return
	}
	sd, pesan := validasiSumberDaya(in)
	if pesan != "" {
		tulisError(w, http.StatusBadRequest, pesan)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	sd.ID = s.nextSDID
	s.nextSDID++
	s.sumberDaya = append(s.sumberDaya, sd)
	tulisJSON(w, http.StatusCreated, sd)
}

func (s *Server) daftarSumberDaya(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()
	hasil := make([]SumberDaya, len(s.sumberDaya)) // tidak nil, jadi JSON-nya []
	copy(hasil, s.sumberDaya)
	tulisJSON(w, http.StatusOK, hasil)
}

func (s *Server) detailSumberDaya(w http.ResponseWriter, r *http.Request) {
	id, ok := ambilID(r)
	if !ok {
		tulisError(w, http.StatusBadRequest, "id harus berupa angka")
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	sd, ada := s.cariSumberDaya(id)
	if !ada {
		tulisError(w, http.StatusNotFound, "sumber daya tidak ditemukan")
		return
	}
	tulisJSON(w, http.StatusOK, sd)
}
