# Layout standard di un progetto Go

Traduzioni:

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

## Panoramica

Questa è un'impostazione di base per applicativi Go. **Non è uno standard ufficiale definito dal team principale di Go**; Invece è un insieme di pattern archittetturali provenienti da progetti ben consolidati nell'ecosistema Go. Alcuni di questi pattern sono più popolari di altri. Sono presenti anche diversi piccoli miglioramenti con alcune cartelle di supporto comuni a qualsiasi grande applicativo in contesto reale.

**`Se stai imparando Go o se stai sviluppando una PoC o un semplice progetto personale, questa struttura è una complicazione non necessaria. Invece inizia con qualcosa di veramente semplice (un unico file `main.go` e `go.mod` è abbastanza).`** Con la crescita del tuo progetto tieni a mente che sarà sempre più importante la corretta impostazione del tuo codice, altrimenti finirai con codice disordinato con molte dipendenze nascoste e uno stato globale. Quando più persone lavorano su un progetto, avrai bisogno di un'impostazione ancora più strutturata. Questo è il momento in cui è importante introdurre un modo comune di gestire pacchetti e librerie. Quando hai un progetto open source o quando sai che altri progetti importano il codice del tuo repository, questo è il momento in cui è importante avere pacchetti e codice privato (`internal`). Clona il repository, mantieni ciò di cui hai bisogno e cancella qualsiasi altra cosa! Solo perchè è presente non significa che vada usato. Nessuno di questi pattern sono usati in ogni singolo progetto. Perfino il `vendor` pattern non è universale.

Con Go 1.14 i [`Go Modules`](https://go.dev/wiki/Modules) sono finalmente pronti per la produzione. Usa [`Go Modules`](https://blog.golang.org/using-go-modules) fino a quando hai una specifica ragione per non usarli e se lo farai non dovrai preoccuparti riguardo $GOPATH e dove mettere il tuo progetto. Il file `go.mod` di base nel repo presuppone che il tuo progetto sia pubblicato su GitHub, ma non è obbligatorio. Il path del modulo può essere uno qualsiasi, anche se la prima parte del path del modulo dovrebbe avere un punto nel nome (l'attuale versione di Go non lo forza più, ma se stai usando una delle versioni leggermente più vecchie, non essere sorpreso se le tue builds falliranno). Guarda le Issues [`37554`](https://github.com/golang/go/issues/37554) e [`32819`](https://github.com/golang/go/issues/32819) se vuoi saperne di più a riguardo.

Questa struttura di progetto è intenzionalmente generica e non cerca di imporre una specifica impostazione per i packages Go.

Questo è uno sforzo della community. Apri una issue se vedi un nuovo pattern o se pensi che uno di quelli esistenti andrebbe aggiornato.

Se hai bisogno di aiuto per la nomenclatura, formattazione e lo stile, inizia da utilizzare [`gofmt`](https://golang.org/cmd/gofmt/) e [`golint`](https://github.com/golang/lint).  Assicurati anche di leggere le linee guida di stile e i consigli:
* https://talks.golang.org/2014/names.slide
* https://golang.org/doc/effective_go.html#names
* https://blog.golang.org/package-names
* https://go.dev/wiki/CodeReviewComments
* [Style guideline for Go packages](https://rakyll.org/style-packages) (rakyll/JBD)

Vedi [`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) per l'aggiunta di altre informazioni a contorno.

Per altro riguardo la nomenclatura e l'organizzazione dei pacchetti e altre impostazioni raccomandate:
* [GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)
* [GopherCon Russia 2018: Ashley McNamara + Brian Ketelsen - Go best practices.](https://www.youtube.com/watch?v=MzTcsI6tn-0)
* [GopherCon 2017: Edward Muller - Go Anti-Patterns](https://www.youtube.com/watch?v=ltqV6pDKZD8)
* [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)

Un post Cinese riguardo delle linee guida per il design Package-Oriented e layer archittetturali
* [面向包的设计和架构分层](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md)

## Go Cartelle

### `/cmd`

Applicazioni principali per questo progetto.

Il nome della cartella per ciascun applicativo dovrebbe coincidere con il nome che vuoi avere per l'eseguibile (es: `/cmd/myapp`).

Non mettere molto codice nella cartella dell'applicazione. Se pensi che il codice potrebbe essere importato e usato in altri progetti, allora dovrebbe essere inserito nella cartella `/pkg`. Se il codice non è riutilizzabile o non vuoi che altri lo riutilizzino, metti questo codice nella cartella `/internal`. Sarai sorpreso di vedere cosa faranno gli altri, quindi si esplicito riguardo le tue intenzioni!

E' comune avere una piccola funzione `main` che importa e invoca il codice dalle cartelle `/internal` e `/pkg` e nient'altro.

Vedi cartella [`/cmd`](cmd/README.md) per esempi.

### `/internal`

Applicativo privato e codice di libreria. Quì c'è il codice che non vuoi gli altri importino nei loro progetti o librerie. Nota che questo pattern archittetturale è forzato dallo stesso compilatore Go. Vedi Go 1.4 [`release notes`](https://golang.org/doc/go1.4#internalpackages) per maggiori dettagli. Nota che non sei obbligato ad avere unicamente la cartella padre `internal`. Puoi avere più di una singola cartella `internal` a qualsiasi livello della tua ramificazione di progetto.


Puoi opzionalmente aggiungere una struttura aggiuntiva ai tuoi pacchetti interni, per separare il tuo codice interno condiviso e non condiviso. Non è obbligatorio (specialmente per piccoli progetti), ma è meglio avere indicazioni per mostrare l'utilizzo raccomandato del pacchetto. Il tuo effettivo codice applicativo può essere inserito nella cartella `/internal/app` (es: `/internal/app/myapp`) e il codice condiviso da queste app nella cartella `/internal/pkg` (es: `/internal/pkg/myprivlib`).

### `/pkg`

Codice di libreria che può essere utilizzato da applicazioni esterne (es:, `/pkg/mypubliclib`). Altri progetti importeranno queste librerie aspettandosi che funzionino, quindi pensaci bene prima di metterci dentro qualcosa :-) Nota che la cartella `internal` è un modo migliore di assicurarsi che i tuoi pacchetti privati non sono importabili, perché ciò è forzato in Go. La cartella `/pkg` è anche un buon modo di esplicitare che il codice contenuto al suo interno è utilizzabile da parte di altri. Il post [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) scritto da Travis Jeffery fornisce una buona panoramica delle cartelle `pkg` e `internal` indicando quando abbia senso usarle.

C'è anche un modo di raggruppare il codice Go in unico posto quando la tua cartella di root contiene molti componenti non-Go e cartelle semplificando l'utilizzo di vari strumenti Go (come menzionato in questi talk: [`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg) dal GopherCon EU 2018, [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) e [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)).

Guarda la cartella [`/pkg`](pkg/README.md) se vuoi vedere quali popolari repo utilizzano questa struttura di progetto. Questo è un pattern di layout comune, tuttavia non è universalmente accettato e qualcuno nella community Go non lo raccomanda.

È ok non utilizzarla se il tuo progetto è molto piccolo e aggiungere un ulteriore strato di annidamento non dà un valore aggiunto (a meno che tu non lo voglia davvero molto :-)). Pensa di prevederla quando sta diventando grande abbastanza e la tua cartella root si sta riempendo troppo (specialmente quando hai molti componenti non Go nell'applicativo).

Origini della cartella `pkg`: Il vecchio codice sorgente di Go usava la cartella `pkg` per i suoi pacchetti e così diversi progetti Go nella community hanno iniziato a copiare questo pattern (vedi [`questo`](https://twitter.com/bradfitz/status/1039512487538970624) tweet di Brad Fitzpatrick per avere un contesto più dettagliato).

### `/vendor`

Dipendenze dell'applicativo (gestite manualmente o dal tuo gestore di pacchetti preferito come la nuova feature built-in [`Go Modules`](https://go.dev/wiki/Modules)). Il comando `go mod vendor` creerà  la cartella `/vendor` per te. Tiene presente che potrebbe essere necessario aggiungere il flag `-mod=vendor` al tuo comando `go build` se non stai utilizzando Go 1.14, dove è utilizzato di default.

Non effettuare il commit delle dipendenze del tuo applicativo se stai sviluppando una libreria.

Nota che sin dalla versione [`1.13`](https://golang.org/doc/go1.13#modules) Go ha anche abilitata la caratteristica del module proxy (usando [`https://proxy.golang.org`](https://proxy.golang.org) come server proxy predefinito). Approfondisci [`quì`](https://blog.golang.org/module-mirror-launch) per vedere se questo soddisfa tutti i tuoi requisiti e vincoli. In caso positivo, allora non avrai bisogno della cartella `vendor`.

## Cartelle di Servizio Applicativo

### `/api`

OpenAPI/Swagger specs, JSON schema files, files di definizione del protocollo.

Vedi la cartella [`/api`](api/README.md) per esempi.

## Cartelle Applicativo Web

### `/web`

Componenti specifici per applitavi Web: assets web statici, templates server side e SPAs.

## Cartelle comuni applicativo

### `/configs`

Templates per il file di configurazione o configurazioni di default.

Metti i tuoi templates `confd` o `consul-template` quì.

### `/init`

Inizializzazione del sistema (systemd, upstart, sysv) e configurazioni per process manager/supervisor (runit, supervisord).

### `/scripts`

Script per effettuare varie operazioni per la build, installazione, analisi, ecc...

Questi script mantengono a livello di root un Makefile piccolo e immediato (es: [`https://github.com/hashicorp/terraform/blob/master/Makefile`](https://github.com/hashicorp/terraform/blob/master/Makefile)).

Vedi la cartella [`/scripts`](scripts/README.md) per esempi.

### `/build`

Packaging e Continuous Integration.

Metti le configurazioni dei tuoi pacchetti: cloud (AMI), container (Docker), OS (deb, rpm, pkg) e script nella cartella `/build/package`.

Metti le tue configurazioni della CI (travis, circle, drone) e script nella cartella `/build/ci`. Nota che qualche tool di CI (es: Travis CI) sono molto stringenti riguardo la posizione dei propri file di configurazione. Prova mettendo le configurazioni nella cartella `/build/ci` collegandole al percorso dove gli strumenti di CI se le aspettano (quando possibile).

### `/deployments`

Configurazioni e template per distribuzioni IaaS, PaaS, di sistema e basati su sistemi di orchestrazione (docker-compose, kubernetes/helm, mesos, terraform, bosh). Nota che in alcuni repo (specialmente per gli applicativi pubblicati con kubernetes) questa cartella è chiamata `/deploy`.

### `/test`

Ulteriori app di test esterne e dati di test. Sentiti libero di strutturare la cartella `/test` come preferisci. Per progetti più grandi ha senso avere una sotto cartella data. Per esempio potresti avere  `/test/data` o `/test/testdata` se hai bisogno che Go ignori il contenuto di questa cartella. Nota che Go ignorerà anche le cartelle o file che iniziano con "." o "_", così si ha più flessibilità in termini di come intendi chiamare la cartella dei tuoi dati test.

Vedi la cartella [`/test`](test/README.md) per esempi.

## Altre Cartelle

### `/docs`

Documenti dell'utente e di Design (in aggiunta alla tua documentazione godoc autogenerata).

Vedi la cartella [`/docs`](docs/README.md) per esempi.

### `/tools`

Strumenti di supporto per il progetto. Nota che questi strumenti possono importare codice dalle cartelle `/pkg` e `/internal`.

Vedi la cartella [`/tools`](tools/README.md) per esempi.

### `/examples`

Esempi per il tuo applicativo e/o librerie pubbliche.

Vedi la cartella [`/examples`](examples/README.md) per esempi.

### `/third_party`

Strumenti esterni di aiuto, codice forcato e altre utility di terze parti (es: Swagger UI).

### `/githooks`

Hook di Git.

### `/assets`

Altri asset del tuo repository (immagini, loghi, etc).

### `/website`

Questo è il posto in cui inserire i dati del sito Web del tuo progetto se non stai utilizzando le GitHub pages.

Vedi la cartella [`/website`](website/README.md) per esempi.

## Cartelle che Non Dovresti Avere

### `/src`

Qualche progetto Go ha una cartella `src`, ma comunemente succede quando gli sviluppatori provengono dal mondo Java, dove è una pratica comune. Se vuoi aiutarti prova a non adottare questo pattern Java. Non vuoi davvero che il tuo codice Go o i tuoi progetti Go assomiglino a quelli Java :-)

Non confondere la cartella `/src` a livello di progetto con la cartella `/src` che Go utilizza per i suoi workspace come descritto in [`How to Write Go Code`](https://golang.org/doc/code.html). La variabile di ambiente `$GOPATH` punta al tuo (attuale) workspace (di default punta a `$HOME/go` su sistemi non Windows). Questo workspace include le cartelle di livello superiore `/pkg`, `/bin` e `/src`. Il tuo progetto attuale finisce per essere una sotto cartella di `/src`, quindi se hai la cartella `/src` nel tuo progetto, il tuo path di progetto assomiglierà a questo: `/some/path/to/workspace/src/your_project/src/your_code.go`. Nota che da Go 1.11 è possibile avere il proprio progetto al di fuori del `GOPATH`, ma ciò non significa che sia una buona idea utilizzare questo pattern di layout.


## Badges

* [Go Report Card](https://goreportcard.com/) - Scansione il tuo codice con `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` e `misspell`. Rimpiazza `github.com/golang-standards/project-layout` con la referenza al tuo progetto.

    [![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)

* ~~[GoDoc](http://godoc.org) - Fornisce una verione online della tua documentazione GoDoc generata. Cambia il link per puntare al tuo progetto.~~

    [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)

* [Pkg.go.dev](https://pkg.go.dev) - Pkg.go.dev è una nuova fonte per la scoperta di Go e documentazione. Puoi creare un badge usando lo [badge generation tool](https://pkg.go.dev/badge).

    [![PkgGoDev](https://pkg.go.dev/badge/github.com/golang-standards/project-layout)](https://pkg.go.dev/github.com/golang-standards/project-layout)

* Release - Mostra l'ultima versione per il tuo progetto. Cambia il link github per puntare al tuo progetto.

    [![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## Note

Un template di progetto standardizzato con configurazioni semplici e riutilizzabili, gli aggiornamenti per gli script e il codice sono in corso.
