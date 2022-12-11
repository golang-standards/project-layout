# Standard Go Project Layout

Translations:

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

## Огляд

Це базовий макет для проєктів Go-додатків. Це **`не є офіційним стандартом, визначеним командою розробників Go`**; тим не менш, це звід паттернів програмування в екосистемі Go, що склалися історично. Деякі з цих паттернів більш популярні та відомі ніж інші. Також в цьому макеті є декілька покращень, в тому числі декілька додаткових дерикторій, що використовуються в будь-якому достатньо великому реальному застосунку.

**`Якщо ви тільки вивчаєте Go або створюєте якийсь демонстраційний чи простий проєкт для себе цей макет буде занадто складним. Почність з чогось дійсно простого (одного файлу `main.go` та `go.mod` буде достатньо).`** Коли проєкт почне рости не забувайте, важливо щоб код залишався добре структурованим, інакше ви отримаєте брудний код з великою кількістю прихованих залежностей та глобальних станів. Коли над проектом працює більше людей, вам знадобиться ще більше структури. Саме тоді важливо запровадити єдиний спосіб управління пакетами/бібліотеками. Коли ви маєте проект з відкритим вихідним кодом або коли ви знаєте, що інші проекти імпортують код з вашого репозиторію проекту, саме тоді важливо мати приватні (`internal`) пакети та код. Клонуйте сховище, зберігайте те, що вам потрібно, а все інше видаляйте! Те, що це є, не означає, що ви повинні використовувати все це. Жодна з цих моделей не використовується в кожному окремому проекті. Навіть паттерн `vendor` не є універсальним.

Починаючи з Go 1.14 [`Go модулі`](https://github.com/golang/go/wiki/Modules) нарешті готові до використання. Використовуйте [`Go модулі`](https://blog.golang.org/using-go-modules) якщо у вас немає конкретної причини не використовувати їх, а якщо є, то вам не потрібно турбуватися про $GOPATH і про те, куди ви помістили свій проект. Базовий файл `go.mod` в репозиторії передбачає, що ваш проєкт розміщено на GitHub, однак це не є обов'язковою вимогою. Шлях до модуля може бути будь-яким, але перший компонент шляху до модуля повинен містити крапку в назві (поточна версія Go більше не вимагає цього, але якщо ви використовуєте трохи старіші версії, не дивуйтеся, якщо ваші збірки не працюватимуть без цього). Якщо ви хочете дізнатися більше про це, дивіться: [`37554`](https://github.com/golang/go/issues/37554) та [`32819`](https://github.com/golang/go/issues/32819).

Ця схема проекту є навмисно загальною і не намагається нав'язати конкретну структуру пакета Go.

Це зусилля спільноти. Відкрийте тему, якщо ви бачите новий шаблон або якщо ви вважаєте, що один з існуючих шаблонів потребує оновлення.

IЯкщо вам потрібна допомога з іменуванням, форматуванням та стилем, почніть з запуску [`gofmt`](https://golang.org/cmd/gofmt/) та [`golint`](https://github.com/golang/lint). Також обов'язково ознайомтеся з цими керівними принципами та рекомендаціями щодо стилю коду Go:
* https://talks.golang.org/2014/names.slide
* https://golang.org/doc/effective_go.html#names
* https://blog.golang.org/package-names
* https://github.com/golang/go/wiki/CodeReviewComments
* [Style guideline for Go packages](https://rakyll.org/style-packages) (rakyll/JBD)

Дивіться [`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) для додаткової довідкової інформації.

Більше про іменування та організацію пакетів, а також інші рекомендації щодо структури коду:
* [GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)
* [GopherCon Russia 2018: Ashley McNamara + Brian Ketelsen - Go best practices.](https://www.youtube.com/watch?v=MzTcsI6tn-0)
* [GopherCon 2017: Edward Muller - Go Anti-Patterns](https://www.youtube.com/watch?v=ltqV6pDKZD8)
* [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)

Китайська публікація про керівні принципи пакетно-орієнтованого проектування та рівень архітектури
* [面向包的设计和架构分层](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md)

## Go каталоги

### `/cmd`

Головні застосунки цього проєкту

Ім'я каталогу для кожного додатка повинно збігатися з ім'ям виконуваного файлу, який ви хочете мати (наприклад `/cmd/myapp`).

Не розміщуйте багато коду в каталозі програми. Якщо ви вважаєте, що код може бути імпортований і використаний в інших проектах, то він повинен знаходитися в каталозі `/pkg`. Якщо код не може бути використаний повторно або якщо ви не хочете, щоб інші використовували його повторно, помістіть цей код в каталог `/internal`. Ви будете здивовані тим, що будуть робити інші, тому будьте відверті у своїх намірах!

Зазвичай є невелика функція `main`, яка імпортує та викликає код з каталогів `/internal` та `/pkg` і нічого більше.

Дивіться [`/cmd`](cmd/README.md) для прикладів.

### `/internal`

Приватний код додатків та бібліотек. Це код, який ви не хочете, щоб інші імпортували у свої програми або бібліотеки. Зауважте, що цей шаблон компонування забезпечується самим компілятором Go. Дивіться Go 1.4 [`release notes`](https://golang.org/doc/go1.4#internalpackages) для додаткових деталей. Зверніть увагу, що ви не обмежені директорією `internal` верхнього рівня. Ви можете мати більше одного каталогу `internal` на будь-якому рівні дерева вашого проекту.

За бажанням ви можете додати трохи додаткової структури до ваших внутрішніх пакунків, щоб відокремити ваш спільний і не спільний внутрішній код. Це не є обов'язковим (особливо для невеликих проектів), але приємно мати візуальні підказки, що показують передбачуване використання пакунків. Ваш власне код програми може знаходитися у каталозі `/internal/app` (наприклад, `/internal/app/myapp`), а код, який використовується спільно з іншими програмами, у каталозі `/internal/pkg` (наприклад, `/internal/pkg/myprivlib`).

### `/pkg`

Код бібліотеки, який можна використовувати зовнішніми програмами (наприклад, `/pkg/mypubliclib`). Інші проекти імпортуватимуть ці бібліотеки, очікуючи, що вони будуть працювати, тому подумайте двічі, перш ніж щось сюди класти :-) Зауважте, що каталог `internal` є кращим способом гарантувати, що ваші приватні пакунки не будуть імпортовані, оскільки це забезпечується Go. Каталог `/pkg` все ще є гарним способом явно повідомити, що код у цьому каталозі є безпечним для використання іншими. Допис у блозі [`I'll take pkg over internal`] (https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) Тревіса Джеффрі (Travis Jeffery) надає гарний огляд каталогів `pkg` та `internal` і того, коли може мати сенс їх використання.

Це також спосіб згрупувати код Go в одному місці, коли ваш кореневий каталог містить багато не-Go компонентів і каталогів, що полегшує запуск різних інструментів Go (як згадувалося в цих доповідях: [`Best Practices for Industrial Programming`] (https://www.youtube.com/watch?v=PTE4VJIdHPg) з GopherCon EU 2018, [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps] (https://www.youtube.com/watch?v=oL6JBUk6tj0) та [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go] (https://www.youtube.com/watch?v=3gQa1LWwuzk)).

Якщо ви хочете побачити, які популярні репозиторії Go використовують цей шаблон оформлення проектів, зверніться до каталогу [`/pkg`](pkg/README.md). Це загальноприйнятий шаблон, але він не є загальноприйнятим, і дехто у спільноті Go не рекомендує його використовувати.

Можна не використовувати його, якщо ваш проект програми дійсно невеликий і де додатковий рівень вкладеності не додає особливої цінності (якщо тільки ви дійсно цього не хочете :-)). Подумайте про це, коли він стане досить великим і ваш кореневий каталог буде досить зайнятий (особливо якщо у вас багато компонентів програми, які не є Go).

Походження каталогу `pkg`: старий вихідний код Go використовував `pkg` для своїх пакунків, а потім різні проекти Go у спільноті почали копіювати цей шаблон (див. [`це`] (https://twitter.com/bradfitz/status/1039512487538970624) твіт Бреда Фіцпатріка для більш детального контексту).

### `/vendor`

Залежності додатків (управляються вручну або за допомогою вашого улюбленого інструменту управління залежностями, наприклад, нової вбудованої функції [`Go модулі`] (https://github.com/golang/go/wiki/Modules)). Команда `go mod vendor` створить для вас каталог `/vendor`. Зауважте, що вам може знадобитися додати прапорець `-mod=vendor` до команди `go build`, якщо ви не використовуєте Go 1.14, де він увімкнений за замовчуванням.

Не фіксуйте залежності програми, якщо ви створюєте бібліотеку.

Зверніть увагу, що починаючи з [`1.13`](https://golang.org/doc/go1.13#modules) Go також включив функцію модульного проксі (використовуючи [`https://proxy.golang.org`](https://proxy.golang.org) в якості свого модульного проксі-сервера за замовчуванням). Прочитайте більше про нього [`тут`](https://blog.golang.org/module-mirror-launch), щоб дізнатися, чи відповідає він усім вашим вимогам і обмеженням. Якщо так, то каталог `vendor` вам взагалі не знадобиться.

## Каталоги сервісних додатків

### `/api`

Специфікації OpenAPI/Swagger, файли схем JSON, файли визначення протоколів.

Приклади див. у каталозі [`/api`](api/README.md).

## Каталоги веб-додатків

### `/web`

Специфічні компоненти веб-додатків: статичні веб-активи, шаблони на стороні сервера та односторінкові застосунки.

## Загальні директорії додатків

### `/configs`

Шаблони файлів конфігурації або конфігурації за замовчуванням.

Сюди викладіть файли шаблонів `confd` або `consul-template`.

### `/init`

Конфігурації системного запуску (systemd, upstart, sysv) та диспетчера/супервізора процесів (runit, supervisord).

### `/scripts`

Скрипти для виконання різних операцій по збірці, установці, аналізу і т.д.

Ці скрипти роблять Makefile кореневого рівня невеликим і простим (наприклад, [`https://github.com/hashicorp/terraform/blob/master/Makefile`](https://github.com/hashicorp/terraform/blob/master/Makefile)).

Приклади див. у каталозі [`/scripts`](scripts/README.md).

### `/build`

Упаковка та безперервна інтеграція (CI).

Конфігурації хмарних (AMI), контейнерних (Docker), ОС (deb, rpm, pkg) пакетів та скрипти покладіть в каталог `/build/package`.

Помістіть конфігурації та скрипти CI (travis, circle, drone) в каталог `/build/ci`. Зверніть увагу, що деякі інструменти CI (наприклад, Travis CI) дуже прискіпливі до розташування своїх конфігураційних файлів. Спробуйте помістити конфігураційні файли в каталог `/build/ci`, пов'язавши їх з тим місцем, де їх очікують інструменти CI (коли це можливо).

### `/deployments`

IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh). Note that in some repos (especially apps deployed with kubernetes) this directory is called `/deploy`.

### `/test`

Additional external test apps and test data. Feel free to structure the `/test` directory anyway you want. For bigger projects it makes sense to have a data subdirectory. For example, you can have `/test/data` or `/test/testdata` if you need Go to ignore what's in that directory. Note that Go will also ignore directories or files that begin with "." or "_", so you have more flexibility in terms of how you name your test data directory.

See the [`/test`](test/README.md) directory for examples.

## Other Directories

### `/docs`

Design and user documents (in addition to your godoc generated documentation).

See the [`/docs`](docs/README.md) directory for examples.

### `/tools`

Supporting tools for this project. Note that these tools can import code from the `/pkg` and `/internal` directories.

See the [`/tools`](tools/README.md) directory for examples.

### `/examples`

Examples for your applications and/or public libraries.

See the [`/examples`](examples/README.md) directory for examples.

### `/third_party`

External helper tools, forked code and other 3rd party utilities (e.g., Swagger UI).

### `/githooks`

Git hooks.

### `/assets`

Other assets to go along with your repository (images, logos, etc).

### `/website`

This is the place to put your project's website data if you are not using GitHub pages.

See the [`/website`](website/README.md) directory for examples.

## Directories You Shouldn't Have

### `/src`

Some Go projects do have a `src` folder, but it usually happens when the devs came from the Java world where it's a common pattern. If you can help yourself try not to adopt this Java pattern. You really don't want your Go code or Go projects to look like Java :-)

Don't confuse the project level `/src` directory with the `/src` directory Go uses for its workspaces as described in [`How to Write Go Code`](https://golang.org/doc/code.html). The `$GOPATH` environment variable points to your (current) workspace (by default it points to `$HOME/go` on non-windows systems). This workspace includes the top level `/pkg`, `/bin` and `/src` directories. Your actual project ends up being a sub-directory under `/src`, so if you have the `/src` directory in your project the project path will look like this: `/some/path/to/workspace/src/your_project/src/your_code.go`. Note that with Go 1.11 it's possible to have your project outside of your `GOPATH`, but it still doesn't mean it's a good idea to use this layout pattern.


## Badges

* [Go Report Card](https://goreportcard.com/) - It will scan your code with `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` and `misspell`. Replace `github.com/golang-standards/project-layout` with your project reference.

    [![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)

* ~~[GoDoc](http://godoc.org) - It will provide online version of your GoDoc generated documentation. Change the link to point to your project.~~

    [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)

* [Pkg.go.dev](https://pkg.go.dev) - Pkg.go.dev is a new destination for Go discovery & docs. You can create a badge using the [badge generation tool](https://pkg.go.dev/badge).

    [![PkgGoDev](https://pkg.go.dev/badge/github.com/golang-standards/project-layout)](https://pkg.go.dev/github.com/golang-standards/project-layout)

* Release - It will show the latest release number for your project. Change the github link to point to your project.

    [![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## Notes

A more opinionated project template with sample/reusable configs, scripts and code is a WIP.
