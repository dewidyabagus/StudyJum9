# Membuat DSN MySQL
export DSN_MYSQL="mysql://root:root123@tcp(172.17.0.1:3307)/products?charset=utf8&parseTime=True&loc=Local&x-no-lock=true"

# Menjalankan migrasi menggunakan DSN
./migrate -path ./migrations -database ${DSN_MYSQL} up

# Menjalankan migrasi untuk versi N + 1
./migrate -path ./migrations -database ${DSN_MYSQL} up 1

# Menjalankan migrasi untuk menghapus semua versi migrasi
./migrate -path ./migrations -database ${DSN_MYSQL} down

# Menjalankan migrasi untuk menghapus versi N - 1
./migrate -path ./migrations -database ${DSN_MYSQL} down 1

# Fixing dirty migration dengan memasukan nilai versi sebelumnya
./migrate -path ./migrations -database ${DSN_MYSQL} force 2

# Pindah versi migration
./migrate -path ./migrations -database ${DSN_MYSQL} goto 2
