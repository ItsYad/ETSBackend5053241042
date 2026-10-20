# PinjamKampus API

- Nama dan NRP: Ziyad Raziq Lahitidra Afey, 5053241042
- Persona: C (Admin himpunan)
- Cara menjalankan: `go run .` lalu buka http://localhost:8080

## Fitur

- [x] B1 Sumber daya
- [x] B2 Buat peminjaman
- [x] B3 Lihat peminjaman (daftar, filter, detail)
- [x] B4 Validasi dan error JSON
- [x] B5 Deteksi bentrok (409)
- [ ] B6 Ubah dan hapus
- [ ] B7 Aman untuk permintaan bersamaan
- [ ] Fitur khas persona
- [ ] Change request

## Yang belum selesai

- B6: endpoint PUT dan DELETE /peminjaman/{id} belum dibuat.
- B7: setiap handler sudah memakai mutex, tetapi belum saya uji dengan banyak request bersamaan, jadi belum saya centang.
- Fitur khas Persona C: status awal peminjaman sudah diajukan, tetapi endpoint setujui, tolak, dan selesaikan belum dibuat.
- Change request belum dikerjakan.
- Filter tanggal di GET /peminjaman dicocokkan sebagai teks persis, formatnya tidak divalidasi.
