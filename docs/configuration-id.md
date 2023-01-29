# Configuration

[English](configuration.md) | Indonesia

```yaml
# contoh konfigurasi
version: 1
host: 0.0.0.0
port: 8000
headers: 
  X-Custom-Header: "hi"
allowOrigins: "https://gofiber.io, https://gofiber.net"
timezone: Asia/Jakarta
publicDir: public
proxyHeader: CF-Connecting-IP
routes:
  - path: /
    file: index.html
  - path: /about
    file: about.html
  - path: /myjs
    file: myjavascript.js
```

## Version

Versi disini bukan versi aplikasi melainkan versi konfigurasi.

## Host

Jika kita mengatur konfigurasi `host` dengan IP tertentu maka aplikasi hanya dapat diakses dengan IP tersebut. Misalnya kita setting `host` dengan IP `192.168.1.1`, maka saat mengaksesnya kita harus menggunakan url <http://192.168.1.1:8000>. Ketika kita mengaksesnya melalui <http://127.0.0.1:8000> maka akan mendapatkan error `ERR_CONNECTION_REFUSED` pada browser anda. Jika tidak kita atur, secara default aplikasi bisa diakses melalui berbagai IP yang ada di komputer Anda.

## Port

Jika kita tidak mengatur konfigurasi `port` maka secara default aplikasi berjalan pada port `8000`. Jika menggunakan Let's Encrypt maka port `80` sebagai default.

## TLS Port

Jika kita tidak mengatur konfigurasi `tlsPort` maka secara default aplikasi berjalan pada port `443`. Konfigurasi ini digunakan saat Let's Encrypt diatur.

## Headers

Konfigurasi ini untuk melakukan kustom header pada response.

## Allow Origins

Untuk mengaktifkan cors, atur konfigurasi dengan origin yang Anda inginkan cukup dengan menulis `allowOrigins: "https://example.com"`. Jika Anda menginginkan beberapa origin, pisahkan secara terpisah menggunakan koma (`,`) seperti `allowOrigins: "https://example.com, https://example.net"` ini. Secara default cors kosong.

## Timezone

Konfigurasi ini untuk membuat format waktu pada log dengan zona waktu yang kita tentukan. Misal kita setting `timezone` dengan timezone `Asia/Jakarta`, maka format lognya akan seperti ini.

```shell
[2023-01-17T17:34:33+07:00] -  - 200 GET / Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36
```

Jika tidak disetel, ini akan ditetapkan secara default ke zona waktu `UTC`.

## Public Directory

Pada konfigurasi tertulis `publicDir`, konfigurasi ini untuk menentukan lokasi direktori `public`. Misalnya, saat kita menggunakan React.js, direktori build diberi nama `build`, kita cukup menyetel `publicDir` dengan nama direktori build. Jika tidak disetel secara default, direktorinya adalah `public`.

## Proxy Header

Konfigurasi ini digunakan untuk mendapatkan IP pengguna jika host berjalan di belakang Load Balancer. Sebagai contoh, jika kita menggunakan Cloudflare, biasanya IP pengguna diset ke header dengan nama `CF-Connecting-IP` kemudian di konfigurasi `proxyHeader` diset ke `proxyHeader: CF-Connecting-IP`. Secara default nilainya kosong.

## Routes

Secara default jika kita tidak mengatur konfigurasi `routes`, Feserve akan menggunakan konfigurasi berikut.

```yaml
routes:
  - path: *
    file: index.html
```

Dengan konfigurasi di atas, setiap url dan parameter akan direspons dengan file `index.html`, cocok jika kita menggunakan SPA.

Anda juga dapat memanipulasi url sebagai berikut.

```yaml
routes:
  - path: /
    file: index.html
  - path: /about
    file: about.html
  - path: /myjs
    file: myjavascript.js
```

Dengan konfigurasi diatas ketika kita mengakses `/about` file yang direspon adalah `about.html` yang sebenarnya kita juga bisa mengaksesnya dengan `/about.html`. Di sini kita juga melihat dari konfigurasi di atas bahwa jika kita mengakses `/myjs` kita akan mendapatkan file respon javascript dari `myjavascript.js`.

Jika ingin feserve menjadi load balancer dengan config seperti berikut:

```yaml
routes:
  - path: "*"
    file: index.html
  - path: /payment
    balancer:
      - http://localhost:3000
      - http://localhost:3001
```

## Let's Encrypt

Untuk menggunakan Let's Encrypt tambahkan konfigurasi seperti contoh berikut:

```yaml
letsencrypt:
  email: feserve@example.com
  domains:
    - example.com
    - www.example.com
  certsPath: certs
```
