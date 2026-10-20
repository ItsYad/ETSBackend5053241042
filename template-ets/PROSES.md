# PROSES

## 1. Persona dan masalahnya

Persona saya adalah admin himpunan (Persona C). Ruang dan alat himpunan tidak boleh langsung terpakai begitu ada yang meminjam, jadi setiap permintaan harus disetujui admin dulu. Fitur khasnya membuat peminjaman mulai dari status diajukan, lalu admin memilih setujui atau tolak (dengan alasan), dan setelah dipakai ditandai selesai.

## 2. Rancangan struktur data

- SumberDaya {id, nama, jenis, stok} disimpan di slice `sumberDaya`.
- Peminjaman {id, sumber_daya_id, peminjam, tanggal, slot_mulai, slot_selesai, jumlah, status, alasan} disimpan di slice `peminjaman`.
- Counter `nextSDID` dan `nextPemID` untuk id berikutnya. Counter hanya naik setelah data berhasil disimpan.
- Semuanya ada di struct `Server` bersama satu `sync.Mutex`.
- Slice dipilih karena data selalu ditambah di belakang, jadi urutan sudah naik menurut id tanpa perlu diurutkan.
- Struct input memakai pointer supaya field yang tidak dikirim (nil) bisa dibedakan dari nilai 0 atau "".

## 3. Satu keputusan desain

(Satu keputusan yang kamu ambil, pilihan lain yang kamu pertimbangkan, dan alasanmu memilih yang ini. Contoh: kenapa cek bentrok per slot, bukan per rentang.)

## 4. Bug atau kendala (minimal dua)

### Bug 1

- Gejala:
- Cara menemukan:
- Penyebab dan solusi:

### Bug 2

- Gejala:
- Cara menemukan:
- Penyebab dan solusi:

## 5. Change request

(Apa yang diminta, bagian kode mana yang kamu ubah, dan satu keputusan desain yang kamu ambil.)
