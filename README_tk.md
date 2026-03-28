# Go taslamalaryn standartlary

Terjimeler:

* [한국어 문서](README_ko.md)
* [简体中文](README_zh.md)
* [正體中文](README_zh-TW.md)
* [简体中文](README_zh-CN.md) - ???
* [Français](README_fr.md)
* [日本語](README_ja.md)
* [Português](README_ptBR.md)
* [Español](README_es.md)
* [Română](README_ro.md)
* [Русский](README_ru.md)
* [Türkçe](README_tr.md)
* [Türkmençe](README_tk.md)
* [Italiano](README_it.md)
* [Tiếng Việt](README_vi.md)
* [Українська](README_ua.md)
* [Indonesian](README_id.md)
* [हिन्दी](README_hi.md)
* [فارسی](README_fa.md)
* [Беларуская](README_be.md)

## Gysgaça

Bu Go programma taslamalary üçin esasy gurluşdyr. Bu, diňe umumy gurluşa ünsi jemleýändiginden we taslamanyň içeriğine girmeýändiginden sebäpli, mazmun taýdan esasydyr. Şeýle-de, bu, taslamalaryňyzy nähili has giňişleýin gurluşlandyryp boljakdygy barada ýokary derejede bolan, jikme-jik maglumatlary çuňlaşdyrmadyk bir gurluşdyr. Mysal üçin, bu, "Clean Architecture" ýaly gurluşlary öz içine almaz.

Bu **`Go core işläp düzýän topary tarapyndan kesgitlenen resmi bir standart däldir`**. Bu, Go ekosistemasynda ýüze çykýan we taryhy taýdan ulanylan birnäçe umumy taslama gurluş nusgalarynyň setidir. Bu nusgalaryň käbiri beýlekilerden has meşhurdyr. Şeýle-de, bu, kiçi gowşak täzelenmeler we uly taslamalaryň esasy strukturasyna laýyk birnäçe goldaw berýän kataloglardan ybaratdyr. **`Go topary Go taslamalaryny gurluşlandyrmak we taslamanyň import edilende we gurnalanlarynda näme manyny berýändigi barada ajaýyp umumy ýörelgeleri hödürleýär`**. Döwrebap Go dokumentasiýasynda [`Go moduly gurluşlandyryp başlamak`](https://go.dev/doc/modules/layout) sahypasyna göz aýlaň. Ol `internal` we `cmd` katalog görnüşlerini (aşakda düşündirilen) we beýleki peýdaly maglumatlary öz içine alýar.

**`Eger Go öwrenmäge synanyşýan bolsaňyz ýa-da özüňiz üçin PoC ýa-da ýönekeý bir taslama döretýän bolsaňyz, bu taslama gurluşy fazla bolýar. Başlamak üçin gaty ýönekeý bir zat bilen başlaň (bir `main.go` faýly we `go.mod` ýerliklidir).`** Taslamanyňyz ösende, koduňyz düzgünli gurluşlandyrylmalydygyny unutmaly, otherwise koduňyz gaty çydamly bolmaz we gizlin baglylyklar we global ýagdaýlar bilen doly bolar. Birnäçe adam taslamada işlese, has köp gurluşa gerek bolar. Şonda, paketler we kitaplar bilen işleşmegiň umumy usulyny ornaşdyrmak möhümdir. Açyk çeşme taslamasy ýa-da beýleki taslamalar sizin taslama reposyňyzdan kody import edýän bolsa, şonda gizlin (`internal`) paketler we kodlar peýdaly bolar. Repo'nu klonlaň, gerek bolan zatlary saklaň we galanlaryny ýok ediň! Diňe şol ýerde bar diýip, hemmesini ulanmalysyňyz diýip pikir etmeli däl. Bu nusgalaryň hiç biri her bir taslamada ulanylmaz. `vendor` nusgasy hem umumy däl.

Go 1.14 bilen [`Go Modullary`](https://go.dev/wiki/Modules) hasaplaşyldy. [`Go Modullary`](https://blog.golang.org/using-go-modules) ulanyň, ýöne degişli sebäplerden ulanmaly bolsaňyz, onda `$GOPATH` we taslamany nereye goýmalydygyny alada etmäň. Repo'daki esasy `go.mod` faýly taslamanyňyzy GitHub-da ýerleşdirmäge garaşýar, ýöne bu talap däl. Modul ýolunyň ilkinji komponentinde nokat bolmalydyr ( häzirki Go wersiýasy bu talaby güýçlendirmeýär, emma birneme köne wersiýalary ulanyp bilýän bolsaňyz, ýygnanyşmalaryňyz başarsyz bolsa gaty geň bolmaň). Bu barada goşmaça maglumat almak üçin [`37554`](https://github.com/golang/go/issues/37554) we [`32819`](https://github.com/golang/go/issues/32819) meselesine serediň.

Bu taslama gurluşy niýetli generic bolup, Go paket gurluşy boýunça anyk bir hereketi berjaý etmeýär.

Bu jemgyýetçilik tagallasy. Täze nusga görseňiz ýa-da bar bolan nusgalaryň birini täzelemeli hasaplasyňyz, mesele açyň.

Eger atlandyrmak, formatlamak we stil bilen kömege mätäç bolsaňyz, [`gofmt`](https://golang.org/cmd/gofmt/) we [`staticcheck`](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck) ulanyp başlaň. Eski standart lintr, golint, indi işlenmeýär we goýulmaýar; saklanylýan lintr, meselem staticcheck ulanylmaly. Şeýle-de, şu Go kod stil gollanmalaryny we maslahatlaryny okamagy unutmaň:
* https://talks.golang.org/2014/names.slide
* https://golang.org/doc/effective_go.html#names
* https://blog.golang.org/package-names
* https://go.dev/wiki/CodeReviewComments
* [Go paketleri üçin stil gollanmasy](https://rakyll.org/style-packages) (rakyll/JBD)

Go proýektiniň gurluşy barada goşmaça maglumat almak üçin [`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) sahypasyna serediň.

Atlandyrmak we paketleriň gurluşy bilen baglanyşykly goşmaça maglumat we beýleki kod gurluşy maslahatlary:
* [GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)
* [GopherCon Russiýa 2018: Ashley McNamara + Brian Ketelsen - Go iň gowy tejribeler.](https://www.youtube.com/watch?v=MzTcsI6tn-0)
* [GopherCon 2017: Edward Muller - Go Anty-Patronlar](https://www.youtube.com/watch?v=ltqV6pDKZD8)
* [GopherCon 2018: Kat Zien - Go Aplikasiýalaryňyzy Näme Üpjün Etmegiňiz Gerek](https://www.youtube.com/watch?v=oL6JBUk6tj0)

Paketler Gurluşy we Arhitekturasy boýunça Hytaýça ýazgy
* [面向包的设计和架构分层](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md)

## Go Kataloglary

### `/cmd`

Bu taslama üçin esasy programma görnüşleri.

Her bir programma üçin katalog ady ýerine ýetiriji programmanyň ady bilen deň bolmaly (mysal üçin, `/cmd/myapp`).

Ýönekeý `main` funksiýasy bolmagy adatydyr, bu funksiýa diňe `/internal` we `/pkg` kataloglaryndaky kodlary import edýär we hiç zatdan başga bir zady ýerine ýetirýär.

[`/cmd`](cmd/README.md) katalogynda mysallara serediň.

### `/internal`

Özboluşly programma we kitaphana kody. Bu, başgalaryň programma ýa-da kitaphana hökmünde import etmek islemeýän koduňyzdyr. Bu gurluş şekli Go kompilýatory tarapyndan iş ýüzünde berjaý edilýär. Go 1.4 [`çykaryş bellikleri`](https://golang.org/doc/go1.4#internalpackages) üçin goşmaça maglumat üçin serediň.

Kodlaryňyzy bölüşmeýän we bölüşýän kodlaryňyzy ayyrmak üçin `/internal/pkg` ýaly aýratyn gurluş goşmak isleseňiz, ol size peýdaly bolup biler.

[`/internal`](internal/README.md) katalogynda goşmaça mysallara serediň.

### `/pkg`

Go taslamalarynda ulanyş üçin giňişleýin ulanylmaly kitaplar (mysal üçin, `/pkg/mypubliclib`). Başga taslamalar bu kitaphanalary import ederler. Şeýle-de, Go-da `internal` katalogy koduňyzyň başgalara import edilmegini kesgitlemek üçin has gowy usuldyr.

[`/pkg`](pkg/README.md) katalogynda bu taslama gurluşy nusgasyny ulanýan Go taslamalaryna serediň.

### `/vendor`

Programmanyň dogry esaslandyrylýan dogry paketleri (Go 1.14-den başlap).

### `/api`

OpenAPI/Swagger kesgitlemeleri, JSON schemalar.

[`/api`](api/README.md) katalogynda mysallara serediň.

## Web Aplikasiýa Kataloglary

### `/web`

Web programmanyna degişli komponentler: statik web assetleri, server tarapdaky şablonlar we SPAlar.

### `/configs`

Konfigurasiýa faýllary ýa-da adaty konfigurasiýalar.

### `/scripts`

Gurluş, gurnama, analiz işleri üçin skriptler.

### `/build`

Pakette we CI konfigurasiýalary.

### `/deployments`

IaaS, PaaS, sistemanyň we konteýneriň orkestrasiýasy we konfigurasiýalary.

### `/test`

Go-da test barlaglary.

### `/docs`

Dizaýn we ulanyjy dokumentleri.

### `/tools`

Taslama üçin goldaw berýän gurallar.

### `/examples`

Siziň programmalar we/ýa-da açyk çeşme kitaphanalary üçin mysallar.

### `/third_party`

Daşary kömekçi gurallar, fork edilen kitaphanalar.

### `/githooks`

Git hooklary.

### `/assets`

Repositoriňiz bilen bilelikde ulanyljak başga serişdeler (suratlar, logolar we ş.m.).

### `/website`

Eger-de GitHub sahypalaryny ulanmasaňyz, bu ýerde taslamaňyz üçin sahypa maglumatlaryny goýup bilersiňiz.

Mysal üçin [`/website`](website/README.md) katalogyna serediň.

## Ulanmaly Bolmadyk Kataloglar

### `/src`

Go taslamalarynyň käbiri `src` katalogyny ulanýar, emma bu adatça Java dünýäsinden gelen developerleriň arasynda ýüze çykýar, şol ýerde bu adatça bir meňzeş görnüşdir. Eger-de özüňizi kömek edip bilseňiz, bu Java şekilini ulanmamaga synanyşyň. Go kodyňyzyň ýa-da Go taslamalaryňyzyň Java ýaly görünmegini islemeýärsiňiz :-)

Taslama derejesindäki `/src` katalogyny Go-nyň iş ýerleri üçin ulanylýan `/src` katalogy bilen saryşdyrmaň, bu barada [`Go Kodyny Nädip Ýazmaly`](https://golang.org/doc/code.html) kitabyna serediň. `$GOPATH` gurşawynyň üýtgeýji siziň (häzirki) iş ýerleriňize görkezýär (Windows däl ulgamlarda, default ýagdaýda `$HOME/go`ya görkezýär). Bu iş ýeri /pkg, /bin we /src kataloglaryny öz içine alýar. Taslamaňyz asyl ýagdaýda `/src` katalogynyň içinde ýerleşýär, şonuň üçin taslamaňyz /src katalogyny öz içine alýan bolsa, taslama ýoly şeýle görnüşde bolar: `/some/path/to/workspace/src/your_project/src/your_code.go`. Go 1.11 bilen taslamaňyzy `GOPATH`-dan daşarda ýerleşdirmek mümkin, emma bu gurluş görnüşini ulanmak halaýan zat däl.

## Badges

* [Go Report Card](https://goreportcard.com/) - Bu, kodyňyzy `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` we `misspell` bilen barlary. `github.com/golang-standards/project-layout` URL-ni öz taslamaňyz bilen çalyşyň.

  [![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)

* ~~[GoDoc](http://godoc.org) - GoDoc tarapyndan döredilen dokumentasiýanyň onlaýn görnüşini hödürlär. Baglanyşygy öz taslamaňyza ugrukdyryň.~~

  [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)

* [Pkg.go.dev](https://pkg.go.dev) - Pkg.go.dev Go gözleg we dokumentasiýa üçin täze bir ugur. [badge dörediş gurallary](https://pkg.go.dev/badge) arkaly badge döretmäge mümkinçilik berýär.

  [![PkgGoDev](https://pkg.go.dev/badge/github.com/golang-standards/project-layout)](https://pkg.go.dev/github.com/golang-standards/project-layout)

* Release - Bu taslamaňyz üçin iň soňky çykaryş sanyny görkezýär. Github baglanyşygyny öz taslamaňyza ugrukdyryň.

  [![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## Bellikler

Has pikirlenilen taslama şablony, mysal/gaýratan ulanylyp boljak konfigurasiýalar, skriptler we kodlar bilen işlenýär.
