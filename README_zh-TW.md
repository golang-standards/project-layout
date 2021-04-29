# 標準 Go 專案目錄結構 (Standard Go Project Layout)

其他語言的翻譯：

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

這是 Go 應用程式專案的基本目錄結構。它不是核心 Go 開發團隊定義的官方標準；然而，它是 Go 生態系統中一組常見的老專案和新專案的目錄結構。其中一些目錄結構比其他目錄結構更受歡迎。這個專案目錄結構還有一些細微的改進，可以支援任何大型且實用的應用程式目錄結構。

如果你才剛開始學習 Go 程式語言，或者你只是想建立一個實驗性的玩具專案，這個專案目錄結構就過於複雜了。從一個非常簡單的 `main.go` 檔案開始，其實已經綽綽有餘。但隨著專案增長，你一定要記得，維持一份良好的程式碼結構其實是非常重要的，否則你最終將會得到一堆淩亂的程式碼，這其中肯定也會包含大量隱藏的相依問題與全域狀態。當有越多人參與專案時，你也將需要更多、更好的目錄結構。這時候就是帶入套件/函式庫常見的管理方法最好的時機。當你擁有一個開源專案，或者當你知道有其他專案從你的專案匯入程式碼時，這時候擁有**私有的** (`internal`) 套件和程式碼就很重要。複製這個專案，保留你需要的內容，刪除所有用不到的內容！因為這些目錄結構在這裡並不意味著你全部都要用。這些目錄結構並不是每個專案都會這樣用，甚至連 `vendor` 模式也不是通用的。

Go 1.14 的 [`Go Modules`](https://github.com/golang/go/wiki/Modules) 已經是正式版本，除非你有特定的理由不使用它們，否則請一律使用 [`Go Modules`](https://blog.golang.org/using-go-modules) 來管理套件。如果你用了 Go Modules 之後，就無需擔心 `$GOPATH` 與專案放置的位置。專案中的 `go.mod` 檔案假設你的專案放置在 GitHub 中，但這不是必須的。模組路徑可以是任意位址，然而你的模組路徑名稱至少要有個「點」(`.`) 才是合法路徑 (雖然最新版的 Go 不會強迫你一定要用網址當成模組路徑，但如果你用了早期的 Go 版本的話，遇到建置失敗就不要覺得奇怪)。如果你想知道更多資訊，請參考 Issues [`37554`](https://github.com/golang/go/issues/37554) 和 [`32819`](https://github.com/golang/go/issues/32819)。

此專案目錄結構是通用的，它並非嘗試強迫讓你使用在特定的 Go 套件結構。

這個套件是社群共同努力的成果。如果你看到新的模式，或者你認為一個現有的模式需要更新，請提出一個 issue 出來討論。

如果需要命名、格式與程式碼風格方面的協助，請使用 [`gofmt`](https://golang.org/cmd/gofmt/) 和 [`golint`](https://github.com/golang/lint)。也請確保你閱讀過以下這些 Go 程式碼撰寫風格的指導方針與建議：

* https://talks.golang.org/2014/names.slide
* https://golang.org/doc/effective_go.html#names
* https://blog.golang.org/package-names
* https://github.com/golang/go/wiki/CodeReviewComments

參見 [`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) 了解更多的背景資訊。

更多關於套件的命名與組織方式，以及其他程式碼結構的建議：

* [GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)
* [GopherCon Russia 2018: Ashley McNamara + Brian Ketelsen - Go best practices.](https://www.youtube.com/watch?v=MzTcsI6tn-0)
* [GopherCon 2017: Edward Muller - Go Anti-Patterns](https://www.youtube.com/watch?v=ltqV6pDKZD8)
* [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)

> 這裡有份中文的文件可供參考：[Go 面向包的设计和架构分层](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md)

## Go 目錄結構

### `/cmd`

本專案的主要應用程式。

每個應用程式的目錄名應該與你的執行檔案名稱一致 (例如：`/cmd/myapp`)。

不要在這個目錄下放置太多程式碼。如果你認為某些程式碼也可以從其他應用程式或專案中匯入使用，那麼這些程式碼應該位於 `/pkg` 目錄中。如果程式碼不是可重複利用的，或者你不希望其他人使用它，請將該程式碼放到 `/internal` 目錄下。你未來將會驚訝的發現別人是怎麼使用你的程式碼，所以請現在就明確的表達你的意圖！

通常主要應用程式只會有一個小小的 `main` 函式，然後大部分的程式都是從 `/internal` 和 `/pkg` 匯入呼叫並執行，除此之外應該什麼都沒有！

請查看 [`/cmd`](cmd/README.md) 目錄獲得更多範例。

### `/internal`

**私有應用程式**和**函式庫**的程式碼，是你不希望其他人在其應用程式或函式庫中匯入的程式碼。請注意：這個目錄結構是由 Go 編譯器本身所要求的。有關更多細節，請參閱 Go 1.4 的 [`release notes`](https://golang.org/doc/go1.4#internalpackages)。注意：這個目錄並不侷限於放在專案最上層的 `internal` 目錄。事實上，你在專案目錄下的任何子目錄都可以包含 `internal` 目錄。

你可以選擇性的加入一些額外的目錄結構到你的內部套件(`internal package`)中，用來區分你想「共用」與「非共用」的內部程式碼(internal code)。這不是必要的（尤其是對小型專案來說），但有視覺上的線索來表達套件的共用意圖來說，肯定會更好(nice to have)。你的應用程式程式碼可以放在 `/internal/app` 目錄下 (例如：`/internal/app/myapp`)，而這些應用程式共享的程式碼就可以放在 `/internal/pkg` 目錄下 (例如：`/internal/pkg/myprivlib`)。

### `/pkg`

函式庫的程式碼當然可以讓外部應用程式來使用 (例如：`/pkg/mypubliclib`)，其他專案會匯入這些函式庫，並且期待它們能正常運作，所以要把程式放在這個目錄下請多想個幾遍！:-) 注意：使用 `internal` 目錄可以確保私有套件不會被匯入到其他專案使用，因為它是由 Go 的編譯器強制執行的，所以是比較好的解決方案。使用 `/pkg` 目錄仍然是一種很好的方式，它代表其他專案可以安全地使用這個目錄下的程式碼。由 Travis Jeffery 撰寫的 [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) 文章提供了關於 `pkg` 和 `internal` 目錄很好的概述，以及使用它們的時機點。

當專案的根目錄包含許多不是用 Go 所寫的元件與目錄時，將 Go 程式碼放在一個集中的目錄下也是種不錯的方法，這使得運行各種 Go 工具變得更加容易（正如以下這些演講中提到的那樣：來自 GopherCon EU 2018 的 [`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg)、[GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) 和 [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)）。

如果你想查看哪些知名的 Go 專案使用本專案的目錄結構，請查看 [`/pkg`](pkg/README.md) 目錄。這是一組常見的目錄結構，但並不是所有人都接受它，有些 Go 社群的人也不推薦使用。

如果你的應用程式專案真的很小，或是套用這些資料夾不會對你有太大幫助（除非你真的很想用XD），不使用本專案推薦的目錄結構是完全沒問題的。當你的專案變的越來越大，根目錄將會會變得越來越複雜（尤其是當你有許多不是 Go 所寫的元件時)，你可以考慮參考這個專案所建議的目錄結構來組織你的程式碼。

### `/vendor`

應用程式的相依套件可透過手動管理，或使用你喜歡的相依性套件管理工具，例如內建的 [`Go Modules`](https://github.com/golang/go/wiki/Modules) 特性。使用 `go mod vendor` 命令可以幫你建立一個 `/vendor` 目錄。請注意：如果你不是用 Go 1.14+ 版本的話，你可能需要在執行 `go build` 的時候增加 `-mod=vendor` 命令列參數。從 Go 1.14 開始，這個參數預設就是啟用的。

如果你正在建立一個函式庫套件，那麼請不要將你應用程式的相依套件加入版控！

注意：從 [`Go 1.13`](https://golang.org/doc/go1.13#modules) 開始，Go 預設啟用了**模組的代理伺服器** (module proxy) 功能 (預設使用 [`https://proxy.golang.org`](https://proxy.golang.org) 作為模組的代理伺服器)。你可以從[`這裡`](https://blog.golang.org/module-mirror-launch)查看這功能是否符合你的需求與限制。如果你可以使用 module proxy 的話，那麼你根本不需要 `vendor` 目錄。

## 服務應用程式目錄 (Service Application Directories)

### `/api`

OpenAPI/Swagger 規格、JSON schema 檔案、各種協定定義檔。

請查看 [`/api`](api/README.md) 目錄獲得更多範例。

## Web 應用程式目錄 (Web Application Directories)

### `/web`

Web 應用程式相關的元件：靜態 Web 檔案、伺服器端範本與 SPAs 相關檔案。

## 通用應用程式目錄 (Common Application Directories)

### `/configs`

組態設定的檔案範本或預設設定。

將你的 `confd` 或 `consul-template` 範本檔案放在這裡。

### `/init`

放置 System init (`systemd`, `upstart`, `sysv`) 與 Process manager/supervisor (`runit`, `supervisor`) 相關設定。

### `/scripts`

放置要執行各種建置、安裝、分析等操作的命令腳本！

這些腳本可以讓你在根目錄的 `Makefile` 更小、更簡單（例如：[`https://github.com/hashicorp/terraform/blob/master/Makefile`](https://github.com/hashicorp/terraform/blob/master/Makefile))。

請查看 [`/scripts`](scripts/README.md) 目錄獲得更多範例。

### `/build`

封裝套件與持續整合(CI)。

將你的雲端 (AMI)、容器 (Docker)、OS (deb, rpm, pkg) 套件的組態設定與腳本放在 `/build/package` 目錄下。

將你的 CI (Travis CI, CircleCI, Drone CI) 的組態設定與腳本放在 `/build/ci` 目錄中。請注意：有些 CI 工具 (例如 Travis CI 等)，它們對這些組態設定檔案的位置非常挑剔。如果可能的話，請嘗試將檔案放在 `/build/ci` 目錄中，並**連結** (linking) 這些檔案到 CI 工具期望它們出現的位置。

### `/deployments`

IaaS、PaaS、系統和容器編配部署的組態設定與範本 (docker-compose、kubernetes/helm、mesos、terraform、bosh)。注意：在某些儲存庫中（特別是那些部署在 kubernetes 的應用程式），這個目錄會被命名為 `/deploy`。

### `/test`

額外的外部測試應用程式和測試資料。你可以自在的調整你在 `/test` 目錄中的結構。對於較大的專案來說，通常會有一個 `data` 資料夾也是蠻正常的。例如：如果你需要 Go 忽略這些目錄下的檔案，你可以使用 `/test/data` 或 `/test/testdata` 當作你的目錄名稱。請注意：Go 還會忽略以 `.` 或 `_` 開頭的目錄或檔案，所以你在測試資料的目錄命名上，將擁有更大的彈性。

請查看 [`/test`](test/README.md) 目錄獲得更多範例。

## 其他目錄

### `/docs`

設計和使用者文件 (除了 `godoc` 自動產生的文件之外)。

請查看 [`/docs`](docs/README.md) 目錄獲得更多範例。

### `/tools`

這個專案的支援工具。注意：這些工具可以從 `/pkg` 和 `/internal` 目錄匯入程式碼。

請查看 [`/tools`](tools/README.md) 目錄獲得更多範例。

### `/examples`

放置關於你的應用程式或公用函式庫的使用範例。

請查看 [`/examples`](examples/README.md) 目錄獲得更多範例。

### `/third_party`

外部輔助工具、Forked 程式碼，以及其他第三方工具 (例如：Swagger UI)。

### `/githooks`

Git hooks。

### `/assets`

其他要一併放入儲存庫的相關檔案 (例如圖片、Logo、... 等等)。

### `/website`

如果你不使用 GitHub Pages 的話，這裡可以放置專案的網站相關資料。

請查看 [`/website`](website/README.md) 目錄獲得更多範例。

## 你不應該擁有的目錄

### `/src`

有些 Go 專案確實擁有一個 `src` 資料夾，但這通常發生在開發人員有 Java 背景，這在 Java 的世界很常見。如果你可以嘗試不要採用這種 Java 常見的資料夾的話，你應該不希望你的 Go 程式碼或 Go 專案看起來像 Java 吧！ :-)

不要將專案層級的 `/src` 目錄與 [`How to Write Go Code`](https://golang.org/doc/code.html) 所描述的 `/src` 混為一談。`$GOPATH` 環境變數指向你目前的工作區 (workspace)（非 Windows 的作業環境預設指向 `$HOME/go`），這個工作區包含了最上層的 `/pkg`、`/bin` 和 `/src` 目錄。你的實際專案最終其實是放在 `/src` 下的一個子目錄，所以你的專案路徑大概會長這樣：`/some/path/to/workspace/src/your_project/src/your_code.go`。注意：雖然 Go 1.11 可以將專案放在 `GOPATH` 之外，但這並不意味著使用這種目錄結構模式是一個好主意！

## 徽章 (Badges)

* [Go Report Card](https://goreportcard.com/) - 它會使用 `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` 與 `misspell` 掃描你的程式碼。如下範例請替換 `github.com/golang-standards/project-layout` 為你的專案參考。

    [![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)

* ~~[GoDoc](http://godoc.org) - It will provide online version of your GoDoc generated documentation. Change the link to point to your project.~~

    [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)

* [Pkg.go.dev](https://pkg.go.dev) - Pkg.go.dev 是一個探索 Go 與查閱相關文件的新目標，你可以利用 [徽章建立工具](https://pkg.go.dev/badge) 來替你的套件建立一個徽章。

    [![PkgGoDev](https://pkg.go.dev/badge/github.com/golang-standards/project-layout)](https://pkg.go.dev/github.com/golang-standards/project-layout)

* Release - 它會顯示你專案中的最新發行版號。請變更 GitHub 連結到你的專案！

    [![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## 備註 (Notes)

更多的專案範本，連同範例或可重用的組態設定、腳本命令、程式碼都在持續進行中!
