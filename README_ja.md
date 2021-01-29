# Standard Go Project Layout

翻訳:

* [English](README.md)
* [한국어 문서](README_ko.md)
* [简体中文](README_zh.md)
* [正體中文](README_zh-TW.md)
* [简体中文](README_zh-CN.md) - ???
* [Français](README_fr.md)
* [日本語](README_ja.md)
* [Español](README_es.md)

## 概要

これは、Goアプリケーションプロジェクトの基本的なレイアウトです。これは、コアとなるGo開発チームによって定義された公式の標準ではありませんが、Goエコシステムの中で、歴史的に共通しているプロジェクトのレイアウトパターンのセットとなっています。これらのパターンの中には、他のパターンよりも人気のあるものもあります。また、現実世界の大規模なアプリケーションに共通するいくつかのサポートディレクトリに加えて、いくつかの小さな機能強化が行われています。

Goを学ぼうとしている場合や、自分でPoCやおもちゃのプロジェクトを構築しようとしている場合、このプロジェクトレイアウトはやりすぎです。最初は本当にシンプルなものから始めてください（`main.go`ファイルが1つあれば十分です）。プロジェクトが大きくなるにつれて、コードが適切に構造化されているかどうかが重要になることに注意してください。そうしないと、多くの隠れた依存関係やグローバルな状態を持つ厄介なコードになってしまいます。プロジェクトで作業する人が増えれば、さらに多くの構造が必要になります。そこで、パッケージやライブラリを管理するための共通の方法を導入することが重要になります。オープンソースプロジェクトがある場合や、他のプロジェクトがプロジェクトリポジトリからコードをインポートしていることを知っている場合は、プライベートな(内部的な)パッケージやコードを持つことが重要になります。リポジトリをクローンして、必要なものだけを残し、他のものはすべて削除してください。リポジトリにあるからといって、すべてを使わなければならないわけではありません。これらのパターンはすべてのプロジェクトで使われているわけではありません。`vendor`パターンでさえも、万能ではありません。

Go 1.14では、[`Go Modules`](https://github.com/golang/go/wiki/Modules)がついに本番に向けて準備が整いました。使用しない特別な理由がない限り、[`Go Modules`](https://blog.golang.org/using-go-modules) を使用してください。もし使用するのであれば、$GOPATH やプロジェクトをどこに置くかを気にする必要はありません。レポの基本的な `go.mod` ファイルは、プロジェクトが GitHub でホストされていることを前提としていますが、必須ではありません。モジュールパスは何でも構いませんが、最初のモジュールパスコンポーネントの名前にはドットを付けてください (現在の Go のバージョンではもうこれを強制していませんが、少し古いバージョンを使っているのであれば、これを付けなくてもビルドが失敗しても驚かないでください)。これについて詳しく知りたい場合は、Issue [`37554`](https://github.com/golang/go/issues/37554) と [`32819`](https://github.com/golang/go/issues/32819) を参照してください。

このプロジェクトは意図的に一般的なレイアウトを使用しており、特定のGoパッケージを押し付けているわけではありません。

これはコミュニティの取り組みです。 新しいパターンが表示された場合、または既存のパターンの1つを更新する必要があると思われる場合は、issueで起票してください。

名前付け、フォーマット、スタイルについてサポートが必要な場合は、[`gofmt`](https://golang.org/cmd/gofmt/)と[`golint`](https://github.com/golang/lint)を実行することから始めます。また、次のGoコードスタイルのガイドラインと推奨事項も必ずお読みください:
* https://talks.golang.org/2014/names.slide
* https://golang.org/doc/effective_go.html#names
* https://blog.golang.org/package-names
* https://github.com/golang/go/wiki/CodeReviewComments
* [Style guideline for Go packages](https://rakyll.org/style-packages) (rakyll/JBD)

追加の背景情報については、[`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2)を参照してください。

パッケージの命名と整理、およびその他のコード構造の推奨事項の詳細:
* [GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)
* [GopherCon Russia 2018: Ashley McNamara + Brian Ketelsen - Go best practices.](https://www.youtube.com/watch?v=MzTcsI6tn-0)
* [GopherCon 2017: Edward Muller - Go Anti-Patterns](https://www.youtube.com/watch?v=ltqV6pDKZD8)
* [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)

パッケージ指向の設計ガイドラインとアーキテクチャレイヤーに関する中国の投稿
* [パッケージ指向の設計とアーキテクチャの階層化](https://github.com/danceyoung/paper-code/blob/master/package-directiond-design/packagedirectioneddesign.md)

## Goのディレクトリ

### `/cmd`

このプロジェクトの主なアプリケーション。

各アプリケーションのディレクトリ名は、欲しい実行ファイルの名前と一致するようにしてください(例: `/cmd/myapp`)。

アプリケーションディレクトリには多くのコードを入れないようにしましょう。コードをインポートして他のプロジェクトで使えると思うならば、`/pkg`ディレクトリに置くべきです。コードが再利用できない場合や、他の人に再利用してほしくない場合は、そのコードを`/internal`ディレクトリに置いてください。他の人が何をするか驚くでしょうから、自分の意図を明確にしてください。

このような場合は、`/internal`ディレクトリと`/pkg`ディレクトリからコードをインポートして呼び出す小さなメイン関数を持つのが一般的ですが、それ以外は何もしません。

例に関しては、[`/cmd`](cmd/README.md)ディレクトリを参照してください。

### `/internal`

プライベートなアプリケーションやライブラリのコード。これは、他の人が自分のアプリケーションやライブラリにインポートしたくないコードです。このレイアウトパターンは、Goコンパイラによって強制されることに注意してください。詳細については、Go 1.4の[`リリースノート`](https://golang.org/doc/go1.4#internalpackages)を参照してください。トップレベルの内部ディレクトリに限定されないことに注意してください。プロジェクトツリーのどのレベルでも、複数の内部ディレクトリを持つことができます。

オプションで、内部パッケージに少し余分な構造を追加して、共有内部コードと非共有内部コードを分離することができます。(特に小規模なプロジェクトでは) 必須ではありませんが、パッケージの使用目的を示す視覚的な手がかりがあるのは良いことです。実際のアプリケーションコードは `/internal/app` ディレクトリ (例: `/internal/app/myapp`) に、それらのアプリケーションで共有されるコードは `/internal/pkg` ディレクトリ (例: `/internal/pkg/myprivlib`) に置くことができます。

### `/pkg`

外部アプリケーションで使用しても問題ないライブラリコード(例: `/pkg/mypubliclib`)。他のプロジェクトは、これらのライブラリが動作することを期待してインポートしますので、ここに何かを置く前によく考えてください :-)。内部ディレクトリは、プライベートパッケージがインポートできないようにするためのより良い方法であることに注意してください。`/pkg` ディレクトリは、そのディレクトリにあるコードが他の人に使われても安全であることを明示的に伝える良い方法です。[`I'll take pkg over internal blog post by Travis Jeffery`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) は、`pkg` ディレクトリと内部ディレクトリの概要と、それらを使用することが意味のある場合の概要を提供しています。

また、ルートディレクトリにGo以外のコンポーネントやディレクトリが多数含まれている場合は、Goコードを1つの場所にグループ化して、さまざまなGoツールを簡単に実行できるようにする方法でもあります（これらの講演で言及されているように：[`産業用プログラミングのベストプラクティス`](https ：//www.youtube.com/watch？v = PTE4VJIdHPg) from GopherCon EU 2018、[`GopherCon 2018：Kat Zien-How Do You Structure Your Go Apps`](https://www.youtube.com/watch?v= oL6JBUk6tj0) および [`GoLab 2018-Massimiliano Pippi-Goのプロジェクトレイアウトパターン`](https://www.youtube.com/watch?v=3gQa1LWwuzk)）。

このプロジェクトレイアウトパターンを使用している人気のある Go repos を見たい場合は [`/pkg`](pkg/README.md) ディレクトリを参照してください。これは一般的なレイアウトパターンですが、普遍的に受け入れられているわけではありませんし、Goコミュニティの中には推奨していない人もいます。

アプリプロジェクトが非常に小さく、追加レベルのネストがあまり価値をもたらさない場合は、使用しないでください（本当に必要な場合を除く:-)）。 それが十分に大きくなり、ルートディレクトリがかなり肥大化してきたときに考えてください（特に、Go以外のアプリコンポーネントがたくさんある場合）。

### `/vendor`

アプリケーションの依存関係 (手動で管理するか、新しい組み込みの [`Go Modules`](https://github.com/golang/go/wiki/Modules) 機能のようなお気に入りの依存関係管理ツールで管理します)。`go mod vendor`コマンドは、`/vendor`ディレクトリを作成します。デフォルトでオンになっている Go 1.14 を使用していない場合は、`go build` コマンドに `-mod=vendor` フラグを追加する必要があるかもしれないことに注意してください。

ライブラリをビルドしている場合は、アプリケーションの依存関係をコミットしないでください。

[`1.13`](https://golang.org/doc/go1.13#modules) 以降、Go はモジュールプロキシ機能も有効にしています (デフォルトでは [`https://proxy.golang.org`](https://proxy.golang.org) をモジュールプロキシサーバとして使用しています)。[`この機能`](https://blog.golang.org/module-mirror-launch)についての詳細は、ここを読んで、あなたの要件や制約に適合するかどうかを確認してください。そうであれば、ベンダディレクトリは全く必要ありません。

## Service Application Directories

### `/api`

OpenAPI/Swaggerの仕様、JSONスキーマファイル、プロトコル定義ファイル。

例に関しては、[`/api`](api/README.md)ディレクトリを参照してください。

## Web Application Directories

### `/web`

ウェブアプリケーション固有のコンポーネント：静的ウェブアセット、サーバーサイドテンプレート、SPA。

## Common Application Directories

### `/configs`

設定ファイルのテンプレートまたはデフォルトの設定。

`confd` または `consul-template` テンプレートファイルをここに置きます。

### `/init`

システムinit(systemd, upstart, sysv)とプロセスマネージャ/スーパーバイザ(runit, supervisord)の設定。

### `/scripts`

様々なビルド、インストール、解析などの操作を行うためのスクリプトです。

これらのスクリプトはルートレベルの Makefile を小さくシンプルに保ちます (例: [`https://github.com/hashicorp/terraform/blob/master/Makefile`](https://github.com/hashicorp/terraform/blob/master/Makefile))。

例に関しては、[`/scripts`](scripts/README.md)ディレクトリを参照してください。

### `/build`

パッケージングと継続的インテグレーション。

クラウド (AMI)、コンテナ (Docker)、OS (deb、rpm、pkg) パッケージの設定とスクリプトを `/build/package` ディレクトリに置きます。

CI (travis, circle, drone) の設定とスクリプトを `/build/ci` ディレクトリに配置します。CIツールの中には（Travis CIなど）、設定ファイルの場所に非常にこだわるものがあることに注意してください。コンフィグファイルを `/build/ci` ディレクトリに置き、CIツールが期待する場所にリンクしてみてください（可能であれば）。

### `/deployments`

IaaS、PaaS、システム、コンテナオーケストレーションのデプロイメント設定とテンプレート (docker-compose、kubernetes/helm、mesos、terraform、bosh)。
いくつかのリポジトリ (特に kubernetes でデプロイされたアプリ) では、このディレクトリは `/deploy` と呼ばれていることに注意してください。

### `/test`

追加の外部テストアプリとテストデータ。`/test`ディレクトリは自由に構成してください。大規模なプロジェクトでは、データのサブディレクトリを持つことは理にかなっています。例えば、`/test/data` や `/test/testdata` などのディレクトリが必要な場合、そのディレクトリにあるものを無視することができます。Go は "." や "_" で始まるディレクトリやファイルも無視するので、テストデータディレクトリの名前の付け方に柔軟性があることに注意してください。

例に関しては、[`/test`](test/README.md)ディレクトリを参照してください。

## 他のディレクトリ

### `/docs`

デザインドキュメントとユーザードキュメント (godocで生成されたドキュメントに加えて)。

例に関しては、['/docs`](docs/README.md)ディレクトリを参照してください。

### `/tools`

このプロジェクトをサポートするツールです。これらのツールは `/pkg` と `/internal` ディレクトリからコードをインポートできることに注意してください。

例に関しては、['/tools`](tools/README.md)ディレクトリを参照してください。

### `/examples`

あなたのアプリケーション、またはpublic librariesのための例。

例に関しては、['/examples`](examples/README.md)ディレクトリを参照してください。

### `/third_party`

外部ヘルパーツール、フォークされたコード、その他のサードパーティ製ユーティリティ（Swagger UIなど）。

### `/githooks`

Gitフック。

### `/assets`

リポジトリに付随するその他のアセット（画像、ロゴなど）。

### `/website`

Githubページを使用していない場合は、プロジェクトのWebサイトのデータを置く場所です。

例に関しては、['/website`](website/README.md)ディレクトリを参照してください。

## 作成してはいけないディレクトリ

### `/src`

Goプロジェクトの中には `src` フォルダを持っているものもありますが、これは通常、開発者が一般的なパターンであるJavaの世界から来た場合に起こります。可能であれば、このようなJavaのパターンを採用しないようにしてください。あなたのGoコードやGoプロジェクトがJavaのように見えることは本当に避けてください。

プロジェクトレベルの `/src` ディレクトリと、[`Goコードの書き方`](https://golang.org/doc/code.html)で説明されているように、Go がワークスペースに使用する `/src` ディレクトリを混同しないようにしてください。環境変数 `$GOPATH` は、(現在の)ワークスペースを指します (Windows 以外のシステムでは、デフォルトでは `$HOME/go` を指します)。このワークスペースには、トップレベルの `/pkg`, `/bin`, `/src` ディレクトリが含まれています。実際のプロジェクトは `/src` の下のサブディレクトリになりますので、プロジェクト内に `/src` ディレクトリがある場合、プロジェクトのパスは以下のようになります。`/some/path/to/workspace/src/your_project/src/your_code.go` のようになります。Go 1.11 では、プロジェクトを `GOPATH` の外に置くことができますが、このレイアウトパターンを使うのが良いというわけではないことに注意してください。

## Badges

* [Go Report Card](https://goreportcard.com/) - `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license`, `misspell` でコードをスキャンします。`github.com/golang-standards/project-layout` をプロジェクトリファレンスに置き換えてください。

* ~~[GoDoc](http://godoc.org) - GoDocで作成したドキュメントのオンライン版を提供します。リンクを自分のプロジェクトへのリンクに変更してください。

* [Pkg.go.dev](https://pkg.go.dev) - Pkg.go.devは、囲碁の発見とドキュメントの新しい目的地です。[バッジ生成ツール](https://pkg.go.dev/badge)を使ってバッジを作成することができます。

* リリース - あなたのプロジェクトの最新のリリース番号が表示されます。githubのリンクを変更して、あなたのプロジェクトを指すようにしてください。

[![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/golang-standards/project-layout)](https://pkg.go.dev/github.com/golang-standards/project-layout)
[![リリース](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## 注意事項

サンプル/再利用可能なコンフィグ、スクリプト、コードを備えた、より意見の多いプロジェクトテンプレートはWIPです。
