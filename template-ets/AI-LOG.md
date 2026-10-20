# AI-LOG

Catat setiap penggunaan AI yang menghasilkan kode. Jika tidak memakai AI sama sekali, tulis satu baris: **tidak memakai AI**.

Log yang jujur tidak mengurangi nilai, sebanyak apa pun AI dipakai. Yang mengurangi nilai adalah log yang tidak sesuai dengan kode, atau isi log yang tidak bisa dijelaskan.

## Sumber daya (B1)

- Prompt: Meminta contoh implementasi endpoint sumber daya beserta validasi nama, jenis, dan stok.
- Dipakai: Referensi implementasi validasi, pembuatan data, daftar sumber daya, dan detail berdasarkan ID.
- Diubah: Saya meninjau aturan validasi agar nama tidak kosong, panjang nama dibatasi, jenis hanya menerima ruang atau alat, dan stok minimal satu.

## Peminjaman dan deteksi bentrok (B2, B3, B5)

- Prompt: Meminta bantuan membuat endpoint peminjaman, filter daftar, detail peminjaman, validasi input, dan pemeriksaan bentrok kapasitas per slot.
- Dipakai: Referensi fungsi validasi peminjaman, perhitungan unit terpakai, pemeriksaan bentrok, pembuatan data, daftar, dan detail peminjaman.
- Diubah: Saya memeriksa alur perhitungan kapasitas agar jumlah peminjaman pada sumber daya dan waktu yang sama tidak melebihi stok.
