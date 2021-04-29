# Structură standard pentru aplicațiile Go

Traduceri:

* [English](README.md)
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

## General

Aceasta este o structură de bază pentru aplicațiile Go. Nu este un standard oficial definit de echipa Go; însă reprezintă un set de structuri/modele comune folosite de-a lungul timpului în aplicațiile din ecosistemul Go. Unele sunt mai populare decât altele. Sunt prezente și îmbunătățiri, cu directoare specifice unei aplicații de mari dimensiuni.

Dacă încerci să înveți Go sau să creezi un PoC/proiect hobby aceasta structură este excesivă. Începe cu ceva foarte simplu (un fișier `main.go` e de ajuns). Pe măsură ce proiectul crește ține minte că va fi nevoie să îl structurezi, altfel te vei trezi cu mult cod spagettă și dependințe/state-uri globale greu de întreținut. Când sunt mai mulți oameni care lucrează la proiect, vei avea nevoie de o structură și mai bună. Din acest motiv este important să introduci un mod comun de gestionare a librăriilor/package-urilor. Când ai un proiect open-source, când știi că alte proiecte importă codul din proiectul tău, e important să ai module (packages) private (internal). Clonează acest repo și ține doar ce ai nevoie. Doar pentru că e acolo, nu înseamnă că trebuie să îl folosești. Aceste standarde nu sunt folosite în toate proiectele Go, nici măcar cel comun `vendor` nu este universal.

De la apariția Go 1.14 [`Go Modules`](https://github.com/golang/go/wiki/Modules) sunt gata de producție (production mode). Folosește [`Go Modules`](https://blog.golang.org/using-go-modules) doar dacă nu ai un motiv specific de a nu le folosești. Dacă nu vrei să le folosești atunci nu este nevoie să te gândești la $GOPATH și unde să îți plasezi proiectul. Fișierul `go.mod` din repo-ul tău asumă că proiectul îți este hostat pe Github, însă asta nu e o necesitate. Calea (path) modulelor poate fi orice, însă primul modul component ar trebui să aibă un punct în numele său (versiunea curentă Go nu o mai cere explicit însă dacă folosești o versiune mai veche nu fi surprins dacă primești erori la compilare). Vezi problemele [`37554`](https://github.com/golang/go/issues/37554) și [`32819`](https://github.com/golang/go/issues/32819) dacă vrei să afli mai multe.

Această structură este intenționat generică și nu își dorește să impună un standard în gestiunea modulelor Go.

Această structură este un efort al comunității. Deschide o problemă (issue) dacă observi o nouă structură sau consideri că una existentă ar trebui actualizată.

Dacă ai nevoie de ajutor cu denumirile, formatare, stilare vezi [`gofmt`](https://golang.org/cmd/gofmt/) și [`golint`](https://github.com/golang/lint). Ar fi bine să citești și aceste linii de ghidare în vederea stilării codului Go:
* https://talks.golang.org/2014/names.slide
* https://golang.org/doc/effective_go.html#names
* https://blog.golang.org/package-names
* https://github.com/golang/go/wiki/CodeReviewComments
* [Ghid de stilare pentru modulele Go](https://rakyll.org/style-packages) (rakyll/JBD)

Vezi [`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) pentru informații adiționale.

Mai multe despre numirea și organizarea modulelor + recomandări:
* [GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)
* [GopherCon Russia 2018: Ashley McNamara + Brian Ketelsen - Go best practices.](https://www.youtube.com/watch?v=MzTcsI6tn-0)
* [GopherCon 2017: Edward Muller - Go Anti-Patterns](https://www.youtube.com/watch?v=ltqV6pDKZD8)
* [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)

Un articol chinezeasc despre Package-Oriented-Design și arhitectură:
* [面向包的设计和架构分层](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md)

## Directoarele Go

### `/cmd`

Aplicațiile principale ale acestui proiect.

Numele directorului pentru fiecare aplicație ar trebui să coincidă cu numele executabilului pe care vrei să îl ai (e.g., `/cmd/myapp`).

Nu pune mult cod în directorul aplicației. Dacă el ar trebui importat și folosit în alte proiecte, atunci ar trebui pus în `/pkg`. Dacă nu este reutilizabil sau vrei ca alții să îl reutilizeze, pune codul în `/internal`. Vei fi surprins ce vor face ceilalți, deci fii explicit în intențiile tale!

Este comun să ai o funcție `main` care doar importă și invocă cod din `/internal` și `/pkg`.

Vezi directorul [`/cmd`](cmd/README.md) pentru exemple.

### `/internal`

Cod privat al modulelor/aplicației. Acesta este cod pe care nu vrei alții să îl importe în aplicațiile/modulele lor. Această structură este forțată de compilatorul Go. Vezi Go 1.4 [`release notes`](https://golang.org/doc/go1.4#internalpackages) pentru mai multe detalii. Nu ești însă limitat de un singur director top-level `internal`. Poți avea mai multe directoare `internal` la orice nivel în cadrul proiectului tău.

Poți adăuga o structură adițională modulelor tale interne pentru a separa codul partajat de cel nepartajat. Nu este necesar, dar este bune să ai indicii vizuale pentru a arăta scopul de folosire al modulelor. Codul actual al aplicației poate fi în directorul `/internal/app` (e.g., `/internal/app/myapp`) iar codul partajat de acele aplicații în `/internal/pkg` (e.g., `/internal/pkg/myprivlib`).

### `/pkg`

Cod librărie care poate fi folosit de aplicațiile externe (e.g., `/pkg/mypubliclib`). Alte proiecte vor importa aceste module și se vor aștepta ca ele să funcționeze, așa că gândește-te de 2 ori înainte de a pune ceva aici :) Directorul `internal` este o modalitate mai bună de a fi sigur că modulele tale no sunt importabile deoarece acest fapt este forțat de Go. Directorul `/pkg` este în continuare un mod bun de a comunica explicit codul ca fiind sigur de folosit de către ceilalți. Postarea [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) oferă o prezentare generală a directoarelor `pkg` și `internal`.

Este o metodă bună să grupezi codul Go într-un singur loc atunci când directorul root al aplicației conține multe componente no-Go, (cum este menționat în aceste prezentări: [`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg) from GopherCon EU 2018, [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) and [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)).

Vezi [`/pkg`](pkg/README.md) dacă vrei să vezi care repo-uri populare Go folosesc această structură de proiect. Aceasta este o structură comună de proiect însă nu este universală și nu o recomand în comunitatea Go. 

### `/vendor`

Dependințele aplicației (gestionate manual sau automat). Comanda `go mod vendor` va creea un director `/vendor`. Probabil va trebui să adaugi ca parametru `-mod=vendor` atunci când execuți `go build` dacă nu folosești Go 1.14

Nu da commit la dependințele aplicației dacă construiești un modul.

Odată cu versiunea [`1.13`](https://golang.org/doc/go1.13#modules) Go a implementat funcționalitatea modulelor proxy (folosind [`https://proxy.golang.org`](https://proxy.golang.org) ca server proxy al modulelor implicit). Poți citi mai multe despre el [`aici`](https://blog.golang.org/module-mirror-launch). S-ar putea să nu ai nevoie de directorul `/vendor` dacă poți folosi modulele proxy.

## Directoarele de servicii ale aplicației 

### `/api`

Specificații OpenAPI/Swagger, fișiere JSON schema, fișiere cu definiții de protocoale.

Vezi directorul [`/api`](api/README.md) pentru exemple.

## Directorul aplicațiilor Web

### `/web`

Componente specifice aplicațiilor Web: asset-uri, template-uri, SPAs, etc

## Directoare comune aplicațiilor 

### `/configs`

Șablonuri de configurare, configurări implicite.

Poți pune `confd` sau `consul-template` aici.

### `/init`

Configurări de inițializare system (systemd, upstart, sysv) și gestiune/supervizare a proceselor (runit, supervisord)

### `/scripts`

Scripturi pentru rularea diferitelor operații.

Ele țin nivelul rădăcinii Makefile mic și simplu (e.g., [`https://github.com/hashicorp/terraform/blob/master/Makefile`](https://github.com/hashicorp/terraform/blob/master/Makefile)).

Vezi directorul [`/scripts`](scripts/README.md) pentru examples.

### `/build`

Packaging și Integrare Continuă.

Pune configurări ale modulelor AMI, Docker, OS (deb, rpm, pkg) și scripturi în directorul `/build/package`

Pune configurările CI (travis, circle, drone) și scripturile similare în directorul `/build/ci`. Unele instrumente CI sunt pretențioase când vine vorba de locația configurărilor (e.g., Travis CI). Încearcă să pui fișierele de configurare în directorul `/build/ci` legându-le de locația în care uneltele CI se așteaptă să fie. 

### `/deployments`

Conține sisteme și containere de orchestrare, implementare, șablonare (docker-compose, kubernetes/helm, mesos, terraform, bosh) IaaS, PaaS. În unele repo-uri (în special cele implementate cu kubernetes) acest director e numit direct `/deploy`. 

### `/test`

Softuri și fișiere adiționale de testare. Poți structura directorul `/test` cum dorești. Pentru proiecte mari are sens să conțină sub-directoare. De exemplu, poți avea `/test/data` sau `/test/testdata` dacă vrei ca Go să ignore ce este in acel director. Go va ignora și directoarele sau fișierele care încep cu "." sau "\_", deci ai mai multă flexibilitate cu privire la modul în care îți numești fișierele/directoarele din `/test`

Vezi [directorul](test/README.md) pentru example.

## Alte directoare

### `/docs`

Documentația structurii aplicației

Vezi [`/docs`](docs/README.md) pentru exemple.

### `/tools`

Unelte de suport pentru acest proiect. Ele pot importa și module din directoarele `/pkg` și `/internal`

Vezi [`/tools`](tools/README.md) pentru exemple.

### `/examples`

Exemple pentru aplicația ta și/sau librării publice

Vezi [`/examples`](examples/README.md) pentru exemple.

### `/third_party`

Unelte externe de ajutor, cod forked și alte utilități (e.g., Swagger UI).

### `/githooks`

Git hooks.

### `/assets`

Alte aseturi conținute de repo (images, logos, etc).

### `/website`

Acesta este locul în care îți pui datele website, dacă nu folosești pagini GitHub.

Vezi [`/website`](website/README.md) pentru exemple.

## Directoare pe care NU ar trebui să le ai

### `/src`

Unele proiecte Go au un director `src`, dar se întâmplă de obicei când developer-ii au venit din domeniul Java. Încearcă să nu adopți o astfel de structură, e indicat ca proiectul tău să nu se asemene cu unul tip Java.

Nu confunda directorul `/src` cu cel pe care Go îl folosește ca spațiu de lucru (workspace). Variabila `$GOPATH` arată workspace-ul (implicit el e setat la `$HOME/go` pentru sistemele de operare non-Windows). Acest workspace folosește directoarele `/pkg`, `/bin` și `/src`. Proiectul tău actual ajunge să fie un sub-director sub `/src`, deci dacă ai un director `/src` în proiectul tău, calea va arăta cam așa: `/some/path/to/workspace/src/your_project/src/your_code.go`. Odată cu Go 1.11 este permis să ai proiectul în afara `GOPATH`, însă nu este neapărat o idee bună să adopți o astfel de schemă.

## Adiționale

* [Go Report Card](https://goreportcard.com/) - Îți va scana codul cu `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` și `misspell`. Înlocuiește `github.com/golang-standards/project-layout` cu referința la proiuectul tău.

    [![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)

* ~~[GoDoc](http://godoc.org) - Va furniza versiuni online ale documentației generate local de GoDoc.~~

    [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)

* [Pkg.go.dev](https://pkg.go.dev) - Pkg.go.dev este o nouă destinație pentru descoperiri și documentație Go. Poți creea o insignă (badge) cu [badge generation tool](https://pkg.go.dev/badge).

    [![PkgGoDev](https://pkg.go.dev/badge/github.com/golang-standards/project-layout)](https://pkg.go.dev/github.com/golang-standards/project-layout)

* Release - Va arăta ultimele lansări (versiuni) ale proiectului tău. 

    [![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## Note

Un șablon mai pedant cu configurări probate/reutilizabile, scripturi și cod este un WIP.
