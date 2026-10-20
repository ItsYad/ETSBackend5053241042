# AI-LOG

Catat penggunaan AI yang benar-benar terjadi selama pengerjaan. Bagian yang belum dikerjakan tidak perlu diisi dengan klaim pengerjaan.

## Kerangka server, helper JSON, dan struct (models.go, helper.go, main.go)

- Prompt: Memberikan dokumen requirement dan kartu persona untuk meminta bantuan membuat kerangka REST API Go.
- Dipakai: Kerangka server, fungsi helper, dan struct yang dihasilkan atau dibantu AI.
- Diubah: Menghapus komentar bawaan dan merapikan format kode menggunakan gofmt.

## Sumber daya (B1)

- Prompt: Meminta referensi implementasi endpoint sumber daya.
- Dipakai: Referensi implementasi validasi, pembuatan data, daftar sumber daya, dan detail berdasarkan ID.
- Diubah: Meninjau aturan validasi sesuai kebutuhan requirement.

## Peminjaman dan deteksi bentrok (B2, B3, B5)

- Prompt: Meminta bantuan terkait endpoint peminjaman, filter, detail, validasi input, dan pemeriksaan kapasitas.
- Dipakai: Referensi implementasi fungsi validasi peminjaman dan pemeriksaan bentrok kapasitas.
- Diubah: Memeriksa kesesuaian perhitungan kapasitas dengan stok dan slot peminjaman.

## Fallback error JSON untuk path dan metode yang salah (B4)

- Prompt: Meminta bantuan agar respons 404 dan 405 menggunakan format JSON.
- Dipakai: Implementasi fallback route untuk metode dan path yang tidak sesuai.
- Diubah: Tidak ada.
