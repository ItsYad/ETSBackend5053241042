package main

import (
	"net/http"
	"strconv"
	"strings"
)

func menahanSlot(status string) bool {
	return status == StatusDiajukan || status == StatusDisetujui
}

func (s *Server) validasiPeminjaman(in inputPeminjaman) (Peminjaman, string) {
	var p Peminjaman
	if in.SumberDayaID == nil || in.Peminjam == nil || in.Tanggal == nil ||
		in.SlotMulai == nil || in.SlotSelesai == nil || in.Jumlah == nil {
		return p, "semua field wajib diisi: sumber_daya_id, peminjam, tanggal, slot_mulai, slot_selesai, jumlah"
	}
	peminjam := strings.TrimSpace(*in.Peminjam)
	tanggal := strings.TrimSpace(*in.Tanggal)

	if peminjam == "" {
		return p, "peminjam tidak boleh kosong"
	}
	if panjangTeks(peminjam) > 50 {
		return p, "peminjam maksimal 50 karakter"
	}
	if !tanggalValid(tanggal) {
		return p, "tanggal harus berformat YYYY-MM-DD dan benar-benar ada"
	}
	if *in.SlotMulai < 1 || *in.SlotMulai > 8 || *in.SlotSelesai < 1 || *in.SlotSelesai > 8 {
		return p, "slot harus antara 1 sampai 8"
	}
	if *in.SlotMulai > *in.SlotSelesai {
		return p, "slot_mulai tidak boleh lebih besar dari slot_selesai"
	}
	sd, ada := s.cariSumberDaya(*in.SumberDayaID)
	if !ada {
		return p, "sumber_daya_id tidak ditemukan"
	}
	if *in.Jumlah < 1 || *in.Jumlah > sd.Stok {
		return p, "jumlah harus antara 1 dan stok sumber daya"
	}

	p.SumberDayaID = *in.SumberDayaID
	p.Peminjam = peminjam
	p.Tanggal = tanggal
	p.SlotMulai = *in.SlotMulai
	p.SlotSelesai = *in.SlotSelesai
	p.Jumlah = *in.Jumlah
	return p, ""
}

func (s *Server) unitTerpakai(sdID int, tanggal string, slot int, kecualiID int) int {
	total := 0
	for _, p := range s.peminjaman {
		if p.ID == kecualiID || !menahanSlot(p.Status) {
			continue
		}
		if p.SumberDayaID == sdID && p.Tanggal == tanggal &&
			slot >= p.SlotMulai && slot <= p.SlotSelesai {
			total += p.Jumlah
		}
	}
	return total
}

func (s *Server) adaBentrok(p Peminjaman, kecualiID int) bool {
	sd, _ := s.cariSumberDaya(p.SumberDayaID)
	for slot := p.SlotMulai; slot <= p.SlotSelesai; slot++ {
		if s.unitTerpakai(p.SumberDayaID, p.Tanggal, slot, kecualiID)+p.Jumlah > sd.Stok {
			return true
		}
	}
	return false
}

func (s *Server) indeksPeminjaman(id int) int {
	for i, p := range s.peminjaman {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func (s *Server) buatPeminjaman(w http.ResponseWriter, r *http.Request) {
	body := bacaBody(r)
	var in inputPeminjaman
	if !urai(body, &in) {
		tulisError(w, http.StatusBadRequest, "body bukan JSON yang valid")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	p, pesan := s.validasiPeminjaman(in)
	if pesan != "" {
		tulisError(w, http.StatusBadRequest, pesan)
		return
	}
	p.Status = StatusDiajukan
	p.Alasan = ""
	if s.adaBentrok(p, 0) {
		tulisError(w, http.StatusConflict, "kapasitas pada salah satu slot sudah penuh")
		return
	}
	p.ID = s.nextPemID
	s.nextPemID++
	s.peminjaman = append(s.peminjaman, p)
	tulisJSON(w, http.StatusCreated, p)
}

func (s *Server) daftarPeminjaman(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	filterSD := 0
	adaFilterSD := q.Has("sumber_daya_id")
	if adaFilterSD {
		n, err := strconv.Atoi(q.Get("sumber_daya_id"))
		if err != nil {
			tulisError(w, http.StatusBadRequest, "sumber_daya_id harus berupa angka")
			return
		}
		filterSD = n
	}
	adaTanggal := q.Has("tanggal")
	adaStatus := q.Has("status")

	s.mu.Lock()
	defer s.mu.Unlock()
	hasil := []Peminjaman{}
	for _, p := range s.peminjaman {
		if adaFilterSD && p.SumberDayaID != filterSD {
			continue
		}
		if adaTanggal && p.Tanggal != q.Get("tanggal") {
			continue
		}
		if adaStatus && p.Status != q.Get("status") {
			continue
		}
		hasil = append(hasil, p)
	}
	tulisJSON(w, http.StatusOK, hasil)
}

func (s *Server) detailPeminjaman(w http.ResponseWriter, r *http.Request) {
	id, ok := ambilID(r)
	if !ok {
		tulisError(w, http.StatusBadRequest, "id harus berupa angka")
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	i := s.indeksPeminjaman(id)
	if i < 0 {
		tulisError(w, http.StatusNotFound, "peminjaman tidak ditemukan")
		return
	}
	tulisJSON(w, http.StatusOK, s.peminjaman[i])
}
