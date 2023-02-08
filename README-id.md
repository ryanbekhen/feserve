# Feserve

[![Mentioned in Awesome Fiber](https://awesome.re/mentioned-badge.svg)](https://github.com/gofiber/awesome-fiber#%EF%B8%8F-tools)
[![Go Report Card](https://goreportcard.com/badge/github.com/ryanbekhen/feserve)](https://goreportcard.com/report/github.com/ryanbekhen/feserve)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/ryanbekhen/feserve/release.yml?style=flat-square)
![Release](https://img.shields.io/github/v/release/ryanbekhen/feserve?display_name=release&style=flat-square)
![GitHub all releases](https://img.shields.io/github/downloads/ryanbekhen/feserve/total?style=flat-square)
![GitHub](https://img.shields.io/github/license/ryanbekhen/feserve?style=flat-square)

[English](README.md) | Indonesia

Feserve adalah aplikasi ringan yang dibuat untuk memudahkan para Frontend Developer dalam men-deploy aplikasinya, tanpa harus menggunakan Nginx, Node.js atau sejenisnya yang memakan banyak ruang penyimpanan.

## Fitur

- Sajikan file statis
- Url path khusus ke file
- Load balancer (HTTP)
- Hasilkan sertifikat SSL dan pembaruan otomatis (Let's Encrypt)

## Instalasi

### File Biner

Disini saya menggunakan linux dengan arsitektur amd64 sebagai contoh. Harap sesuaikan dengan OS dan Arsitektur Anda [di sini](https://github.com/ryanbekhen/feserve/releases). Kemudian unduh, verifikasi tanda tangan, dan ekstrak seperti contoh berikut.

```shell
wget https://github.com/ryanbekhen/feserve/releases/download/v0.1.0/feserve_0.1.0_linux_amd64.zip
wget https://github.com/ryanbekhen/feserve/releases/download/v0.1.0/checksums.txt
unzip feserve_0.1.0_linux_amd64.zip 
sha256sum --ignore-missing -c checksums.txt
```

Setelah menjalankan perintah di atas, pindahkan file biner ke `/usr/local/bin` dengan perintah berikut.

```shell
sudo mv feserve /usr/local/bin

# permission for 80/443
sudo setcap 'cap_net_bind_service=+ep' ./usr/local/bin
```

### Melalui `go install`

```shell
go install github.com/ryanbekhen/feserve
```

> **Catatan**: go version go1.19.5 atau lebih

## Mempersiapkan

### Struktur Direktori

```text
root-directory/
|- build/
|- app.yaml
```

### Konfigurasi `app.yaml`

```yaml
version: 1
port: 8000
publicDir: build
```

Dengan konfigurasi di atas, feserve akan berjalan di port `8000` dan `public/` sebagai direktori publiknya. Untuk melihat detail lebih lanjut [di sini](docs/configuration-id.md).

## Penggunaan

### Lokal

Untuk menjalankannya secara lokal, jalankan saja `feserve` di direktori root Anda dengan perintah berikut.

```shell
feserve
```

Kemudian buka browser di <http://localhost:8000>.

### Docker

Sebelum menjalankan perintah di bawah ini, buat dulu file config `app.yaml`.

```shell
docker run --name feserve -d \
  -p 80:80 \
  -p 443:443 \
  -v $(pwd)/certs:/certs \
  -v $(pwd)/app.yaml:/app.yaml \
  ghcr.io/ryanbekhen/feserve:latest
```

### Dockerfile

Untuk menjalankannya di dalam docker, buat file `Dockerfile` seperti contoh berikut.

```Dockerfile
# pembuatan aplikasi
FROM node:16-alpine As build
WORKDIR /app
COPY . .
RUN npm ci 
RUN npm run build
ENV NODE_ENV production

# pembuatan serve
FROM ghcr.io/ryanbekhen/feserve:latest
WORKDIR /app
COPY app.yaml .
COPY --from=build /app/build /app/build
EXPOSE 8000
ENTRYPOINT ["feserve"]
```

Bisa juga dengan cara berikut jika sudah kita build terlebih dahulu.

```Dockerfile
FROM ghcr.io/ryanbekhen/feserve:latest
WORKDIR /app
COPY app.yaml .
COPY build ./build
EXPOSE 8000
ENTRYPOINT ["feserve"]
```

Untuk mencoba menjalankannya cukup dengan perintah berikut.

```shell
docker build -t image-name .
docker run --rm -p 8000:8000 image-name
```

Kemudian buka browser di <http://localhost:8000>.

## Keamanan

Jika Anda menemukan kerentanan keamanan dalam Feserve, silakan kirim email ke ryanbekhen.official@gmail.com.

## Lisensi

Program ini adalah perangkat lunak gratis. Anda dapat mendistribusikan ulang dan/atau memodifikasinya di bawah ketentuan lisensi MIT. Feserve dan kontribusi apapun adalah hak cipta © oleh Achmad Irianto Eka Putra 2023.
