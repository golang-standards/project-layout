# Layout padrão de projetos em Go

Traduções:

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

## Visão geral

Este é um layout básico para projetos de aplicações em Go. Não é um padrão oficial definido pela equipe de desenvolvimento principal do Go; no entanto, é um conjunto de padrões de layout de projetos históricos e emergentes comuns no ecossistema Go. Alguns desses padrões são mais populares do que outros. Ele também possui uma série de pequenos aprimoramentos junto com vários diretórios de suporte comuns a qualquer aplicativo grande o suficiente do mundo real.

Se você está tentando aprender Go, se está construindo um PoC(Prova de conceito) ou um pequeno projeto pessoal para você, este layout de projeto é um exagero. Comece com algo realmente simples (um único arquivo `main.go` é mais do que suficiente). Conforme seu projeto cresce, lembre-se de que será importante garantir que seu código esteja bem estruturado, caso contrário, você acabará com um código confuso com muitas dependências ocultas e estados globais. Quando você tiver mais pessoas trabalhando no projeto, precisará de ainda mais estrutura. É quando é importante apresentar uma maneira comum de gerenciar pacotes/bibliotecas. Quando você tem um projeto de código aberto ou quando conhece outros projetos, importe o código do seu repositório de projetos, é quando é importante ter pacotes e códigos privados (também conhecidos como `internal`). Clone o repositório, mantenha o que você precisa e exclua todo o resto! Só porque está lá, não significa que você precise usar tudo. Nenhum desses padrões é usado em todos os projetos. Mesmo o padrão `vendor` não é universal.

Com Go 1.14 [`Go Modules`](https://github.com/golang/go/wiki/Modules) estão finalmente prontos para produção. Use [`Go Modules`](https://blog.golang.org/using-go-modules) a menos que você tenha um motivo específico para não usá-los e, se tiver, não precisa se preocupar com $GOPATH e onde você colocou seu projeto. O arquivo `go.mod` básico no reposiório assume que seu projeto está hospedado no GitHub, mas não é um requisito. O caminho do módulo pode ser qualquer coisa, embora o primeiro componente do caminho do módulo deva ter um ponto em seu nome (a versão atual do Go não o impõe mais, mas se você estiver usando versões um pouco mais antigas, não se surpreenda se suas compilações falharem sem isto). Veja as issues [`37554`](https://github.com/golang/go/issues/37554) e [`32819`](https://github.com/golang/go/issues/32819) se você quiser saber mais sobre isso.

Este layout de projeto é intencionalmente genérico e não tenta impor uma estrutura de pacote Go específica.

Este é um esforço da comunidade. Abra uma nova issue se você ver um novo padrão ou se você acha que um dos padrões existentes precisa ser atualizado.

Se precisar de ajuda com nomenclatura, formatação e estilo, comece executando [`gofmt`](https://golang.org/cmd/gofmt/) e [`golint`](https://github.com/golang/lint). Além disso, certifique-se de ler estas diretrizes e recomendações de estilo de código Go:
* https://talks.golang.org/2014/names.slide
* https://golang.org/doc/effective_go.html#names
* https://blog.golang.org/package-names
* https://github.com/golang/go/wiki/CodeReviewComments
* [Style guideline for Go packages](https://rakyll.org/style-packages) (rakyll/JBD)

Veja [`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) para obter informações adicionais.

Mais sobre como nomear e organizar pacotes, bem como outras recomendações de estrutura de código:
* [GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)
* [GopherCon Russia 2018: Ashley McNamara + Brian Ketelsen - Go best practices.](https://www.youtube.com/watch?v=MzTcsI6tn-0)
* [GopherCon 2017: Edward Muller - Go Anti-Patterns](https://www.youtube.com/watch?v=ltqV6pDKZD8)
* [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)

Uma postagem chinesa sobre as diretrizes de design orientado a pacotes e a camada de arquitetura
* [面向包的设计和架构分层](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md)

## Diretórios Go

### `/cmd`

Principais aplicações para este projeto.

O nome do diretório para cada aplicação deve corresponder ao nome do executável que você deseja ter (ex. `/cmd/myapp`).

Não coloque muitos códigos no diretório da aplicação. Se você acha que o código pode ser importado e usado em outros projetos, ele deve estar no diretório `/pkg`. Se o código não for reutilizável ou se você não quiser que outros o reutilizem, coloque esse código no diretório `/internal`. Você ficará surpreso com o que os outros farão, então seja explícito sobre suas intenções!

É comum ter uma pequena função `main` que importa e invoca o código dos diretórios` /internal` e `/pkg` e nada mais.

Veja o diretório [`/cmd`](cmd/README.md) para mais exemplos.

### `/internal`

Aplicação privada e código de bibliotecas. Este é o código que você não quer que outras pessoas importem em suas aplicações ou bibliotecas. Observe que esse padrão de layout é imposto pelo próprio compilador Go. Veja o Go 1.4 [`release notes`](https://golang.org/doc/go1.4#internalpackages) para mais detalhes. Observe que você não está limitado ao diretório `internal` de nível superior. Você pode ter mais de um diretório `internal` em qualquer nível da árvore do seu projeto.

Opcionalmente, você pode adicionar um pouco de estrutura extra aos seus pacotes internos para separar o seu código interno compartilhado e não compartilhado. Não é obrigatório (especialmente para projetos menores), mas é bom ter dicas visuais que mostram o uso pretendido do pacote. Seu atual código da aplicação pode ir para o diretório `/internal/app` (ex. `/internal/app/myapp`) e o código compartilhado por essas aplicações no diretório `/internal/pkg` (ex. `/internal/pkg/myprivlib`).

### `/pkg`

Código de bibliotecas que podem ser usados por aplicativos externos (ex. `/pkg/mypubliclib`). Outros projetos irão importar essas bibliotecas esperando que funcionem, então pense duas vezes antes de colocar algo aqui :-) Observe que o diretório `internal` é a melhor maneira de garantir que seus pacotes privados não sejam importáveis porque é imposto pelo Go. O diretório `/pkg` contudo é uma boa maneira de comunicar explicitamente que o código naquele diretório é seguro para uso. [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) A postagem no blog de Travis Jeffery fornece uma boa visão geral dos diretórios `pkg` e` internal`, e quando pode fazer sentido usá-los.

É também uma forma de agrupar o código Go em um só lugar quando o diretório raiz contém muitos componentes e diretórios não Go, tornando mais fácil executar várias ferramentas Go (conforme mencionado nestas palestras: [`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg) da GopherCon EU 2018, [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) e [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)).

Consulte o diretório [`/pkg`](pkg/README.md) se quiser ver quais repositórios Go populares usam esse padrão de layout de projeto. Este é um padrão de layout comum, mas não é universalmente aceito e alguns na comunidade Go não o recomendam.

Não há problema em não usá-lo se o projeto do seu aplicativo for muito pequeno e onde um nível extra de aninhamento não agrega muito valor (a menos que você realmente queira :-)). Pense nisso quando estiver ficando grande o suficiente e seu diretório raiz ficar muito ocupado (especialmente se você tiver muitos componentes de aplicativos não Go).

### `/vendor`

Dependências de aplicativos (gerenciadas manualmente ou por sua ferramenta de gerenciamento de dependências favorita, como o novo recurso integrado [`Go Modules`](https://github.com/golang/go/wiki/Modules)). O comando `go mod vendor` criará o diretório` /vendor` para você. Note que você pode precisar adicionar a flag `-mod=vendor` ao seu comando` go build` se você não estiver usando Go 1.14 onde ele está ativado por padrão.

Não comprometa as dependências do seu aplicativo se você estiver construindo uma biblioteca.

Observe que desde o Go [`1.13`](https://golang.org/doc/go1.13#modules) também habilitou o recurso de proxy do módulo (usando [`https://proxy.golang.org`](https://proxy.golang.org) como servidor proxy de módulo por padrão). Leia mais sobre isso [`aqui`](https://blog.golang.org/module-mirror-launch) para ver se ele se encaixa em todos os seus requisitos e restrições. Se isso acontecer, então você não precisará do diretório `vendor`.

## Diretórios de aplicativos de serviço

### `/api`

Especificações OpenAPI/Swagger, arquivos de esquema JSON, arquivos de definição de protocolo.

Veja o diretório [`/api`](api/README.md) para mais exemplos.

## Diretórios de aplicativos da web

### `/web`

Componentes específicos de aplicativos da Web: ativos estáticos da Web, modelos do lado do servidor e SPAs.

## Diretórios de aplicativos comuns

### `/configs`

Modelos de arquivo de configuração ou configurações padrão.

Coloque seus arquivos de modelo `confd` ou` consul-template` aqui.

### `/init`

Configurações de inicialização do sistema (systemd, upstart, sysv) e gerenciador/supervisor de processos (runit, supervisord).

### `/scripts`

Scripts para executar várias operações de construção, instalação, análise, etc.

Esses scripts mantêm o Makefile de nível raiz pequeno e simples (ex. [`https://github.com/hashicorp/terraform/blob/master/Makefile`](https://github.com/hashicorp/terraform/blob/master/Makefile)).

Veja o diretório [`/scripts`](scripts/README.md) para mais exemplos.

### `/build`

Empacotamento e integração contínua.

Coloque suas configurações de pacote e scripts em nuvem (AMI), contêiner (Docker), sistema operacional (deb, rpm, pkg) no diretório `/build/package`.

Coloque suas configurações e scripts de CI (travis, circle, drone) no diretório `/build/ci`. Observe que algumas das ferramentas de CI (por exemplo, Travis CI) são muito exigentes quanto à localização de seus arquivos de configuração. Tente colocar os arquivos de configuração no diretório `/build/ci` vinculando-os ao local onde as ferramentas de CI os esperam (quando possível).

### `/deployments`

IaaS, PaaS, configurações e modelos de implantação de orquestração de sistema e contêiner (docker-compose, kubernetes / helm, mesos, terraform, bosh). Observe que em alguns repositórios (especialmente em aplicativos implantados com kubernetes), esse diretório é denominado `/deploy`.

### `/test`

Aplicações de testes externos adicionais e dados de teste. Sinta-se à vontade para estruturar o diretório `/test` da maneira que quiser. Para projetos maiores, faz sentido ter um subdiretório de dados. Por exemplo, você pode ter `/test/data` ou` /test/testdata` se precisar que o Go ignore o que está naquele diretório. Observe que o Go também irá ignorar diretórios ou arquivos que começam com "." ou "_", para que você tenha mais flexibilidade em termos de como nomear seu diretório de dados de teste.

Veja o diretório [`/test`](test/README.md) para mais exemplos.

## Outros diretórios

### `/docs`

Documentos do projeto e do usuário (além da documentação gerada pelo godoc).

Veja o diretório [`/docs`](docs/README.md) para mais exemplos.

### `/tools`

Ferramentas de suporte para este projeto. Observe que essas ferramentas podem importar código dos diretórios `/pkg` e` /internal`.

Veja o diretório [`/tools`](tools/README.md) para mais exemplos.

### `/examples`

Exemplos para seus aplicativos e / ou bibliotecas públicas.

Veja o diretório [`/examples`](examples/README.md) para mais exemplos.

### `/third_party`

Ferramentas auxiliares externas, código bifurcado e outros utilitários de terceiros (por exemplo, Swagger UI).

### `/githooks`

Git hooks.

### `/assets`

Outros recursos para acompanhar seu repositório (imagens, logotipos etc).

### `/website`

Este é o lugar para colocar os dados do site do seu projeto se você não estiver usando as páginas do GitHub.

Veja o diretório [`/website`](website/README.md) para mais exemplos.

## Diretórios que você não deveria ter

### `/src`

Alguns projetos Go têm uma pasta `src`, mas normalmente acontece quando os devs vêm do mundo Java, onde é um padrão comum. Se você puder se ajudar, tente não adotar esse padrão Java. Você realmente não quer que seu código Go ou projetos Go se pareçam com Java :-)

Não confunda o diretório `/src` do nível do projeto com o diretório` /src` que Go usa para seus espaços de trabalho, conforme descrito em [`How to Write Go Code`](https://golang.org/doc/code.html). A variável de ambiente `$GOPATH` aponta para sua área de trabalho(atual) (por padrão, ela aponta para` $HOME/go` em sistemas não Windows). Este espaço de trabalho inclui os diretórios de nível superior `/pkg`,` /bin` e `/src`. Seu projeto atual acaba sendo um subdiretório em `/ src`, então se você tiver o diretório` /src` em seu projeto, o caminho do projeto será parecido com este: `/algum/caminho/para/workspace/src/your_project/src/ your_code.go`. Observe que com Go 1.11 é possível ter seu projeto fora de seu `GOPATH`, mas ainda não significa que é uma boa ideia usar este padrão de layout.


## Distintivos

* [Go Report Card](https://goreportcard.com/) - Ele irá escanear seu código com `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` e `misspell`. Substitua `github.com/golang-standards/project-layout` com sua referência de projeto.

    [![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)

* ~~[GoDoc](http://godoc.org) - Ele fornecerá uma versão online da documentação gerada pelo GoDoc. Mude o link para apontar para seu projeto.~~

    [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)

* [Pkg.go.dev](https://pkg.go.dev) - Pkg.go.dev é um novo destino para Go discovery e docs. Você pode criar um emblema usando o [badge generation tool](https://pkg.go.dev/badge).

    [![PkgGoDev](https://pkg.go.dev/badge/github.com/golang-standards/project-layout)](https://pkg.go.dev/github.com/golang-standards/project-layout)

* Release - Ele mostrará o número da versão mais recente do seu projeto. Altere o link do github para apontar para seu projeto.

    [![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## Notas

Um modelo de projeto mais opinativo com configurações, scripts e código de amostra/reutilizáveis é um WIP.
