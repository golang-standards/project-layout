# Standar Tata Letak Proyek Go

Terjemahan:

* [한국어 문서](README_ko.md)
* [简体中文](README_zh.md)
* [正體中文](README_zh-TW.md)
* [简体中文](README_zh-CN.md) - ???
* [Français](README_fr.md)
* [日本語](README_ja.md)
* [Portuguese](README_ptBR.md)
* [Español](README_es.md)
* [Română](README_ro.md)
* [Русский](README_ru.md)
* [Türkçe](README_tr.md)
* [Українська](README_ua.md)
* [Indonesian](README_id.md)
* [हिन्दी](README_hi.md)
* [Беларуская](README_by.md)

## Ringkasan

Berikut ini merupakan tata letak dasar untuk proyek aplikasi Go. Ini **`bukan standar resmi yang ditetapkan oleh tim pengembang inti Go`**; namun, ini merupakan sejumlah pola tata letak proyek historis dan terkini yang umumnya digunakan dalam ekosistem Go. Beberapa pola ini lebih populer daripada yang lain. Selain itu, terdapat beberapa pembaharuan kecil bersama dengan beberapa direktori pendukung yang umum ditemukan dalam aplikasi dunia nyata yang cukup besar.

Jika kamu sedang belajar Go atau sedang membangun PoC atau proyek sederhana untuk dirimu sendiri, tata letak proyek ini terlalu berlebihan. Mulailah dengan sesuatu yang sederhana saja (sebuah file `main.go` tunggal dan `go.mod` sudah cukup). Ketika proyekmu berkembang, penting untuk memastikan kodemu terstruktur dengan baik, jika tidak, kamu akan berakhir dengan kode yang berantakan dengan banyak dependensi tersembunyi dan state global. Ketika ada lebih banyak orang yang bekerja pada proyekmu, kamu akan membutuhkan struktur yang lebih terorganisir. Itulah saat yang penting untuk memperkenalkan cara umum dalam mengelola paket/pustaka. Ketika kamu memiliki proyek open source atau ketika kamu tahu proyek lain mengimpor kode dari repositori proyekmu, saat itulah penting untuk memiliki paket dan kode pribadi (dikenal juga sebagai `internal`). Klon repositori tersebut, ambil yang kamu butuhkan, dan hapus sisanya! Hanya karena ada di sana, tidak berarti kamu harus menggunakan semuanya. Tidak satu pun dari pola-pola ini digunakan dalam setiap proyek. Bahkan pola `vendor` tidaklah universal.

Dengan Go 1.14, [`Go Modules`](https://go.dev/wiki/Modules) akhirnya siap digunakan untuk produksi. Gunakan [`Go Modules`](https://blog.golang.org/using-go-modules) kecuali kamu memiliki alasan khusus untuk tidak menggunakannya. Jika kamu tidak menggunakan Go Modules, maka kamu tidak perlu khawatir tentang `$GOPATH` dan di mana kamu meletakkan proyekmu. File `go.mod` di dalam repositori ini mengasumsikan bahwa proyekmu di-host di GitHub, tetapi itu bukan menjadi persyaratan. Path modul dapat menjadi apa saja, meskipun komponen path modul pertama sebaiknya memiliki tanda titik dalam namanya (versi Go saat ini tidak lagi memaksakannya, tetapi jika kamu menggunakan versi yang sedikit lebih lama, jangan heran jika proses build tidak akan berhasil). Lihat isu [`37554`](https://github.com/golang/go/issues/37554) dan [`32819`](https://github.com/golang/go/issues/32819) jika kamu ingin tahu lebih banyak mengenai hal ini.

Tata letak proyek ini sengaja dirancang secara generik dan tidak mencoba memaksakan struktur paket Go yang spesifik.

Ini merupakan usaha komunitas. Buka sebuah isu jika kamu melihat pola baru atau jika kamu berpikir bahwa salah satu pola yang sudah ada perlu diperbarui.

Jika kamu membutuhkan bantuan dalam hal penamaan, pemformatan, dan gaya penulisan, mulailah dengan menjalankan [`gofmt`](https://golang.org/cmd/gofmt/) dan [`golint`](https://github.com/golang/lint). Pastikan juga untuk membaca panduan dan rekomendasi gaya penulisan kode Go berikut ini:
* https://talks.golang.org/2014/names.slide
* https://golang.org/doc/effective_go.html#names
* https://blog.golang.org/package-names
* https://go.dev/wiki/CodeReviewComments
* [Style guideline for Go packages](https://rakyll.org/style-packages) (rakyll/JBD)

Lihatlah [`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) untuk informasi latar belakang tambahan.

Lebih lanjut tentang penamaan dan pengorganisasian paket serta rekomendasi struktur kode lainnya:
* [GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)
* [GopherCon Russia 2018: Ashley McNamara + Brian Ketelsen - Go best practices.](https://www.youtube.com/watch?v=MzTcsI6tn-0)
* [GopherCon 2017: Edward Muller - Go Anti-Patterns](https://www.youtube.com/watch?v=ltqV6pDKZD8)
* [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)

Sebuah postingan dalam bahasa Cina tentang pedoman Package-Oriented Design dan Architecture layer:
* [面向包的设计和架构分层](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md)

## Direktori Go

### `/cmd`

Aplikasi utama untuk proyek ini.

Nama direktori untuk setiap aplikasi harus sesuai dengan nama file eksekusi yang diinginkan (misalnya, `/cmd/myapp`).

Jangan menempatkan banyak kode di dalam direktori aplikasi. Jika anda berpikir bahwa kode tersebut dapat diimpor dan digunakan dalam proyek lain, maka kode tersebut harus ditempatkan di dalam direktori `/pkg`. Jika kode tersebut tidak dapat digunakan kembali atau jika anda tidak ingin orang lain menggunakannya kembali, letakkan kode tersebut di dalam direktori `/internal`. Anda akan terkejut dengan apa yang orang lain lakukan, jadi tetap jelaskan niat anda! 

Biasanya, terdapat fungsi `main` kecil yang mengimpor dan memanggil kode dari direktori `/internal` dan `/pkg`, dan tidak ada yang lain.

Lihat direktori [`/cmd`](cmd/README.md) untuk contoh-contoh lebih lanjut.

### `/internal`

Kode aplikasi dan library privat. Ini adalah kode anda yang tidak ingin diimpor oleh aplikasi atau library lain. Perlu dicatat bahwa pola tata letak ini dipaksakan atau dijaga oleh kompiler Go itu sendiri. Lihat catatan rilis Go 1.4 [`di sini`](https://golang.org/doc/go1.4#internalpackages) untuk detailnya. Perhatikan bahwa anda tidak dibatasi pada direktori top level `internal` saja. Anda dapat memiliki lebih dari satu direktori `internal` di setiap tingkatan proyek anda.

Secara opsional anda dapat menambahkan struktur tambahan ke paket internal anda, untuk memisahkan kode internal yang bersifat shared dan non-shared. Hal ini tidak diwajibkan (terutama untuk proyek-proyek kecil), tetapi bagus untuk memiliki petunjuk visual yang menunjukkan penggunaan paket yang dimaksudkan. Sebenarnya kode aplikasi dapat ditempatkan di direktori `/internal/app` (misalnya, `/internal/app/myapp`) dan kode yang dibagikan oleh aplikasi tersebut dapat ditempatkan di direktori `/internal/pkg` (misalnya, `/internal/pkg/myprivlib`).

### `/pkg`

Kode library yang boleh digunakan oleh aplikasi eksternal (misalnya, `/pkg/mypubliclib`). Proyek lain akan mengimpor library ini dengan harapan dapat berfungsi, jadi berpikirlah dua kali sebelum anda meletakkan sesuatu di sini :-) Perlu dicatat bahwa direktori `internal` adalah cara yang lebih baik untuk memastikan paket pribadi anda tidak dapat diimpor karena dijaga oleh Go. Namun, direktori `/pkg` tetap cara yang baik untuk mengkomunikasikan secara eksplisit  bahwa kode di dalam direktori tersebut aman digunakan oleh orang lain. Postingan [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) oleh Travis Jeffery memberikan gambaran yang baik tentang direktori `pkg` dan `internal` serta kapan waktu yang tepat untuk menggunakannya.

Hal ini merupakan cara untuk mengelompokkan kode Go di satu tempat ketika direktori root anda berisi banyak komponen dan direktori non-Go, sehingga memudahkan dalam menjalankan berbagai tools Go (seperti yang disebutkan dalam presentasi-presentasi ini: [`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg) dari GopherCon EU 2018, [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0), dan [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)).

Lihat direktori [`/pkg`](pkg/README.md) jika anda ingin melihat repositori Go populer yang menggunakan pola tata letak proyek seperti ini. Ini adalah pola tata letak yang umum digunakan, akan tetapi tidak diterima secara universal dan beberapa anggota komunitas Go tidak merekomendasikannya.

Tidak masalah jika anda tidak menggunakannya, apabila proyek aplikasi anda benar-benar kecil dan tingkatan level tambahan tidak begitu penting (kecuali jika anda benar-benar menginginkannya :-)). Pikirkanlah hal tersebut ketika proyek anda cukup besar dan direktori root anda sudah cukup sibuk (terutama jika anda memiliki banyak komponen aplikasi non-Go).

Asal-usul direktori `pkg`: Source code Go yang lama menggunakan `pkg` untuk paket-paketnya, dan kemudian berbagai proyek Go dalam komunitas mulai meniru pola tersebut (Lihat [`tweet Brad Fitzpatrick ini`](https://twitter.com/bradfitz/status/1039512487538970624) untuk konteks lebih lanjut).

### `/vendor`

Dependensi aplikasi (dikelola secara manual atau dengan dependency management tool favorit anda seperti fitur baru bawaan [`Go Modules`](https://go.dev/wiki/Modules). Perintah `go mod vendor` akan membuat direktori `/vendor `untuk anda. Perlu dicatat bahwa anda mungkin perlu menambahkan flag `-mod=vendor` ke perintah `go build` jika anda tidak menggunakan Go 1.14 di mana fitur tersebut sudah diaktifkan secara default.

Jangan meng-commit dependency aplikasi anda jika anda sedang membangun sebuah library.

Perlu dicatat bahwa sejak versi [`1.13`](https://golang.org/doc/go1.13#modules), Go telah mengaktifkan fitur module proxy (menggunakan [`https://proxy.golang.org`](https://proxy.golang.org) sebagai server proxy modul default). Baca lebih lanjut tentang fitur ini [`di sini`](https://blog.golang.org/module-mirror-launch) untuk melihat apakah sesuai dengan semua persyaratan dan batasan anda. Jika sesuai, maka anda sama sekali tidak perlu menggunakan direktori `vendor`.

## Direktori Servis Aplikasi

### `/api`

Spesifikasi OpenAPI/Swagger, file skema JSON, file definisi protokol.

Lihat direktori [`/api`](api/README.md) untuk contoh-contohnya.

### `/web`

Komponen-komponen spesifik aplikasi web: aset web statis, template di sisi server, dan SPAs.

## Direktori Umum Aplikasi

### `/configs`

Template file konfigurasi atau konfigurasi default.

Letakkan file template anda di `confd` atau `consul-template`..

### `/init`

Konfigurasi sistem init (systemd, upstart, sysv) dan manajer proses/supervisor (runit, supervisord).

### `/scripts`

Skrip-skrip untuk melakukan berbagai operasi seperti build, instalasi, analisis, dll.

Skrip-skrip ini menjaga Makefile tingkat root agar tetap kecil dan sederhana (misalnya, [`https://github.com/hashicorp/terraform/blob/master/Makefile`](https://github.com/hashicorp/terraform/blob/master/Makefile)).

Lihat direktori [`/scripts`](scripts/README.md) untuk melihat contoh-contohnya.

### `/build`

Pemaketan (Packaging) dan Integrasi Berkelanjutan (Continuous Integration).

Letakkan skrip paket dan konfigurasi untuk cloud (AMI), container (Docker), sistem operasi (deb, rpm, pkg) anda di direktori `/build/package`.

Letakkan skrip dan konfigurasi CI (travis, circle, drone) anda di direktori `/build/ci`. Perhatikan bahwa beberapa tools CI (misalnya, Travis CI) sangat ketat mengenai lokasi file konfigurasi mereka. Coba letakkan file konfigurasi di direktori `/build/ci` dan buat tautan ke lokasi yang diharapkan oleh tools CI (jika memungkinkan).

### `/deployments`

Konfigurasi dan template untuk IaaS, PaaS, orkestrasi sistem, dan container (docker-compose, kubernetes/helm, mesos, terraform, bosh). Perhatikan bahwa dalam beberapa repositori (terutama aplikasi yang diimplementasikan dengan kubernetes), direktori ini disebut `/deploy`.

### `/test`

Tambahan eksternal untuk menguji aplikasi dan data. Aturlah struktur direktori `/test` sesuai keinginan anda. Untuk proyek yang lebih besar, disarankan memiliki subdirektori data. Misalnya, anda dapat memiliki `/test/data` atau `/test/testdata` jika anda ingin Go mengabaikan apa yang ada dalam direktori tersebut. Perhatikan bahwa Go juga akan mengabaikan direktori atau file yang dimulai dengan "." atau "_", sehingga anda memiliki fleksibilitas lebih dalam penamaan direktori data pengujian.

Lihat direktori [`/test`](test/README.md) untuk contoh-contohnya.

## Direktori Lainya

### `/docs`

Dokumentasi desain dan pengguna (selain dokumentasi yang dihasilkan oleh godoc).

Lihat direktori [`/docs`](docs/README.md) untuk contoh-contohnya.

### `/tools`

Tools pendukung untuk proyek ini. Perhatikan bahwa tools ini dapat mengimpor kode dari direktori `/pkg` dan `/internal`.

Lihat direktori [`/tools`](tools/README.md) untuk contoh-contohnya.

### `/examples`

Contoh-contoh aplikasi atau library publik anda.

Lihat direktori [`/examples`](examples/README.md) untuk melihat contoh-contohnya.

### `/third_party`

Tools eksternal, kode yang di-fork, dan utilitas pihak ketiga (third party) lainnya (misalnya, Swagger UI).

### `/githooks`

Git hooks.

### `/assets`

Aset lainnya yang ada di repositori anda (gambar, logo, dll).

### `/website`

Tempat untuk meletakkan data situs web proyek anda jika anda tidak menggunakan halaman GitHub.

Lihat direktori [`/website`](website/README.md) untuk contoh-contohnya.

## Direktori yang Sebaiknya Tidak Dimiliki

### `/src`

Beberapa proyek Go memiliki folder `src`, akan tetapi ini biasanya terjadi ketika pengembang berasal dari dunia Java di mana itu adalah pola umum. Jika memungkinkan, hindari mengadopsi pola Java ini. Anda benar-benar tidak ingin kode Go atau proyek Go anda terlihat seperti Java :-)

Jangan bingung antara direktori proyek `/src` dengan direktori `/src` yang digunakan Go untuk workspace-nya seperti yang dijelaskan dalam [`How to Write Go Code`](https://golang.org/doc/code.html). Variabel environtment `$GOPATH` menunjuk ke workspace anda saat ini (secara default menunjuk ke `$HOME/go` pada sistem non-Windows). Workspace ini mencakup direktori top level `/pkg`, `/bin`, dan `/src`. Proyek aktual anda berada di bawah direktori `/src`, jadi jika anda memiliki direktori `/src` di proyek anda, path proyek akan terlihat seperti ini: `/some/path/to/workspace/src/your_project/src/your_code.go`. Perlu diingat bahwa dengan Go 1.11, memungkinkan untuk memiliki proyek di luar `GOPATH`, tetapi bukan berarti itu adalah ide yang baik untuk menggunakan pola tata letak ini.


## Badges

* [Go Report Card](https://goreportcard.com/) - Akan memindai kode anda menggunakan `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` dan `misspell`. Ubah `github.com/golang-standards/project-layout` dengan referensi proyek anda.

    [![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)

* ~~[GoDoc](http://godoc.org) -Ini akan menyediakan versi online dari dokumentasi yang dihasilkan oleh GoDoc anda. Ubah tautan tersebut agar mengarah ke proyek anda.~~

    [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)

* [Pkg.go.dev](https://pkg.go.dev) - Pkg.go.dev adalah tempat baru untuk menemukan dan mendokumentasikan Go. Anda dapat membuat badge menggunakan [badge generation tool](https://pkg.go.dev/badge).

    [![PkgGoDev](https://pkg.go.dev/badge/github.com/golang-standards/project-layout)](https://pkg.go.dev/github.com/golang-standards/project-layout)

* Release - Ini akan menampilkan nomor rilisan terbaru untuk proyek anda. Ubah tautan GitHub menjadi menunjuk ke proyek anda.

    [![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## Catatan

Opsi template lainya disertai contoh konfigurasi, skrip, dan kode yang dapat digunakan kembali sedang dalam pengembangan (WIP).
