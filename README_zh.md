# Go 项目标准布局

- [English](README.md)
- [한국어 문서](README_ko.md)
- [简体中文](README_zh.md)
- [正體中文](README_zh-TW.md)
- [简体中文](README_zh-CN.md) - ???
- [Français](README_fr.md)
- [日本語](README_ja.md)
- [Portuguese](README_ptBR.md)
- [Español](README_es.md)
- [Română](README_ro.md)
- [Русский](README_ru.md)
- [Türkçe](README_tr.md)
- [Italiano](README_it.md)
- [Vietnamese](README_vi.md)
- [Українська](README_ua.md)
- [Indonesian](README_id.md)
- [हिन्दी](README_hi.md)
- [Беларуская](README_be.md)

这是 Go 应用程序项目的基础布局。请注意，之所以说它基础，一方面是就其内容而言，因为它仅关注通用布局，而不涉及内部的具体内容；另一方面，它也是一个非常高层次（high level）的概述，并没有深入探讨如何能更进一步地构建你的项目结构。例如，它并没有试图涵盖像采用“整洁架构”（Clean Architecture）时你会拥有的那种项目结构。

本文档`并非Go核心开发团队定义的官方标准`，而是 Go 生态中一些常见的历史性和新兴的项目布局模式。这其中有一些模式比另外的一些更受欢迎。这些模式中有些更为流行，有些则使用较少。这个布局还包括了一些小的改进以及在足够复杂的真实世界项目中常见的辅助性目录。需要注意的是，Go 核心团队在项目结构方面提供了一些非常好的通用指南，帮助你理解当项目被导入或安装时意味着什么。详情可参考 Go 官方文档中的[组织 Go 模块(Organizing a Go module)](https://go.dev/doc/modules/layout)页面。本页面包含了如 internal 和 cmd 目录模式（在下文中会提到）以及其他有用的信息。

`如果你正准备学习Go、或者只是为自己构建PoC项目或编写简单的玩具项目，那么按照这个项目进行布局就大材小用了。从一些真正简单的事情开始（一个`main.go`文件就足够了）。` 随着项目的发展，确保代码结构的合理是非常重要的，否则最终会出现很多隐藏的依赖关系和全局状态而导致这个项目的代码混乱。当一个项目多人同时进行时，就更需要有清晰的结构，此时引入一种通用的项目包/标准库管理工具就显得尤为重要。当你维护一个开源项目或者有其他项目导入了你的代码，那么引入私有（即`internal`）包就很重要了。你可以克隆这个项目，保留你项目中需要的部分，然后删除其他部分! 通常来说不需要也没必要使用这个项目中的全部内容。这里的所有模式都不是每个项目必须采用的，甚至如`vendor`模式也并非通用的。

自从 Go 1.14 起，[Go Modules](https://go.dev/wiki/Modules)已经可以用于生产环境中稳定使用了。没有什么特殊原因的话，请使用`Go Modules`。使用 Go Module 时，你就再也不用担心`$GOPATH`的配置和项目实际的存放位置，项目想放在哪里就放在哪里。本项目中`go.mod`文件的内容假设你的项目是托管在 GitHub 上的，当然这不是必选项，因为`Module`中的路径可以是任意的，一般`Module`路径的第一部分中应该包含一个点（.）。最新版的 Go 中不再强制要求这一点，但是如果你使用的是稍微旧一些的版本，没有点号可能遇到编译失败的问题）。详情请看 Issues [37554](https://github.com/golang/go/issues/37554)和 [32819](https://github.com/golang/go/issues/32819)。

本项目布局有意设计的更通用一些，而不会尝试去引入一些特定的 Go 包结构。

这是社区共同的努力。如果发现了一种新的模式或者项目中已经存在的某些模式需要更新，请新建一个 issue。

如果需要一些关于命名、格式化或者样式方面的帮助，请先运行[`gofmt`](https://pkg.go.dev/cmd/gofmt)和[`staticcheck`](https://pkg.go.dev/cmd/gofmt)。之前的标准代码检查工具 golint 已经被弃用并停止维护，推荐使用 staticcheck 这类仍在维护的工具。

另外，请务必阅读以下 Go 代码样式指南和建议：

- <https://talks.golang.org/2014/names.slide>
- <https://golang.org/doc/effective_go.html#names>
- <https://blog.golang.org/package-names>
- <https://go.dev/wiki/CodeReviewComments>
- [Style guideline for Go packages (rakyll/JBD)](https://rakyll.org/style-packages/)

更多背景信息请查看[`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2)。

有关命名和项目包组织样式以及其他代码结构的更多推荐文章：

- [GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)
- [GopherCon Russia 2018: Ashley McNamara + Brian Ketelsen - Go best practices](https://www.youtube.com/watch?v=MzTcsI6tn-0)
- [GopherCon 2017: Edward Muller - Go Anti-Patterns](https://www.youtube.com/watch?v=ltqV6pDKZD8)
- [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)

一篇关于面向包的设计（Package-Oriented-Design）指南与架构分层的中文文章/帖子

- [面向包的设计和架构分层](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md)

## Go 目录

### `/cmd`

该目录用于存放本项目主要的应用程序。

每个应用程序的目录名字，应该和你希望生成的项目可执行文件的名字相匹配（例如: `/cmd/myapp`）。

请不要在这个目录中放太多的代码。如果你认为本项目的某些代码可以被其他项目导入并复用，那么应该把这些代码放在`/pkg`目录中。如果目录中的代码不可重用，或者你不希望被他人使用，应该将代码放在`/internal`目录。你会惊讶于别人可能会怎么用你的代码，因此务必要明确表达你的意图！

通常来说，项目都应该拥有一个小的`main`函数，并在`main`函数中导入或者调用`/internal`和`/pkg`目录中的代码。

具体示例请查看[`/cmd`](https://github.com/golang-standards/project-layout/blob/master/cmd/README.md)目录中的例子。

### `/internal`

该目录用于存放私有的应用程序代码库。这些是你不希望被其他人导入的代码。请注意：这种模式是 Go 编译器强制执行的。详细内容情况可以参考 Go 1.4 的[release notes](https://golang.org/doc/go1.4#internalpackages)。再次注意，在项目的目录树中的任意位置都可以有`internal`目录，而不仅仅是在顶级目录中。

你可以选择性地为内部代码包中添加一些额外的结构，来分隔共享和非共享的内部代码。虽然这不是必选项（尤其是对于小项目中），但是有一个直观的包用途是很棒的。应用程序实际的代码可以放在`/internal/app`目录（例如`internal/app/myapp`），而应用程序的共享代码放在`/internal/pkg`目录中（例如`internal/pkg/myprivlib`）。

将包放入`internal`目录的目的是为了将其设为私有包。如果你将某个包放在 internal 中，那么除非其他包与它有共同的祖先目录，否则无法导入它。值得注意的是，这是 Go 官方文档中唯一有特殊说明并受到编译器特别处理的目录。

### `/pkg`

该目录用于存放可以被外部应用程序的库代码（如，`/pkg/mypubliclib`）。其他项目可能会导入这些库，来保证项目可以正常运行，所以在将代码放在这里前，一定要三思而行 :-) 请注意，`internal`目录是一个更好的选择来确保项目私有代码不会被其他人导入，因为这是 Go 强制执行的。使用`/pkg`目录来明确表示代码可以被其他人安全的导入仍然是一个好方式。Travis Jeffery 撰写的关于 [I’ll take pkg over internal](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) 文章很好地概述了`pkg`和`inernal`目录以及何时使用它们。

当您的根目录包含大量非 Go 组件和目录时，这也是一种将 Go 代码分组到一个位置的方法，从而使运行各种 Go 工具更加容易（在如下的文章中都有提到：2018 年 GopherConEU [Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)，[Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) ，Golab 2018 [Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)）。

点击查看`/pkg`就能看到那些使用这个布局模式的流行 Go 代码仓库。这是一种常见的布局模式，但未被普遍接受，并且 Go 社区中的某些人不推荐这样做。

如果项目确实很小，并且嵌套的层次并不会带来多少价值，那么就不要使用它也是完全没有问题的（除非你就是想用它 :-)）。当你的项目变得更大，并且根目录中包含的内容相当繁杂，可以考虑采用这种方式（尤其是有很多非 Go 的组件）。

`pkg`目录的来源：早期的 Go 源码就是用`pkg`目录来组织包的，后来社区中越来越多的项目开始效仿这种模式。（你可以查看 Brad Fitzpatrick 的[这条推文](https://x.com/bradfitz/status/1039512487538970624)了解更多背景）

### `/vendor`

该目录用于存放应用程序的依赖项（可以通过手动管理，也可以使用你喜欢的依赖管理工具，例如新增的内置[Go Modules](https://go.dev/wiki/Modules)特性）。执行`go mod vendor`命令将会在项目中创建`/vendor`目录，注意，如果使用的不是 Go 1.14 版本，在执行`go build`进行编译时，需要添加`-mod=vendor`命令行选项，因为它不是默认选项。

如果你在构建的是一个库项目，建议不要提交应用程序依赖项。

请注意，从[1.13](https://golang.org/doc/go1.13#modules)开始，Go 也启动了模块代理特性（使用`https：//proxy.golang.org`作为默认的模块代理服务器）。点击[这里](https://blog.golang.org/module-mirror-launch)阅读有关它的更多信息，来了解它是否符合所需要求和约束。如果`Go Module`满足需要，那么就不需要`vendor`目录。

## 服务端应用程序的目录

### `/api`

用于存放 OpenAPI/Swagger 规范、JSON 模式文件、协议定义文件。

更多样例查看[`/api`](https://github.com/golang-standards/project-layout/blob/master/api/README.md)目录。

## Web 应用程序的目录

### `/web`

用于存放 Web 应用程序特定的组件：静态 Web 资源（static web assets）、服务器端模板（server-side templates）、单页应用（Single-Page App，SPA）。

## 通用应用程序的目录

### `/configs`

用于存放配置文件模板或默认配置。

将`confd`或者`consul-template`文件放在这里。

### `/init`

用于存放系统初始化（systemd、upstart、sysv）以及进程管理（runit、supervisord）配置文件。

### `/scripts`

用于执行各种构建，安装，分析等操作的脚本。

这些脚本使根级别的 Makefile 变得更小更简单（例如<https://github.com/hashicorp/terraform/blob/main/Makefile>）。

更多样例查看[`/scripts`](https://github.com/golang-standards/project-layout/blob/master/scripts/README.md)。

### `/build`

打包和持续集成（Continuous Integration）。

将云镜像（AMI）、容器（Docker）、操作系统（deb，rpm，pkg）软件包配置和脚本放在`/build/package`目录中。

将 CI（travis、circle、drone）配置文件和就脚本放在`build/ci`目录中。请注意，有一些 CI 工具（如，travis CI）对于配置文件的位置有严格的要求。尝试将配置文件放在`/build/ci`目录，然后链接到 CI 工具想要的位置。

### `/deployments`

用于存放 IaaS，PaaS，系统和容器编排部署配置和模板（docker-compose，kubernetes/helm，mesos，terraform，bosh）。请注意，在某些存储库中（尤其是使用 kubernetes 部署的应用程序），该目录的名字是`/deploy`。

### `/test`

用于存放外部测试应用程序和测试数据。随时根据需要构建`/test`目录。对于较大的项目，有一个数据子目录更好一些。例如，如果需要 Go 忽略目录中的内容，则可以使用`/test/data`或`/test/testdata`这样的目录名字。请注意，Go 还将忽略以“`.`”或“`_`”开头的目录或文件，因此可以更具灵活性的来命名测试数据目录。

更多样例查看[`/test`](https://github.com/golang-standards/project-layout/blob/master/test/README.md)。

## 其他目录

### `/docs`

用于存放设计和用户文档（除了 godoc 生成的文档）。

更多样例查看[`/docs`](https://github.com/golang-standards/project-layout/blob/master/docs/README.md)。

### `/tools`

此项目的支持工具。请注意，这些工具可以从`/pkg`和`/internal`目录导入代码。

更多样例查看[`/tools`](https://github.com/golang-standards/project-layout/blob/master/tools/README.md)。

### `/examples`

应用程序或公共库的示例。

更多样例查看[`/examples`](https://github.com/golang-standards/project-layout/blob/master/examples/README.md)。

### `/third_party`

外部辅助工具，fork 的代码和其他第三方工具（例如 Swagger UI）。

### `/githooks`

Git 的钩子。

### `/assets`

项目中使用的其他资源（图像，Logo 等）。

### `/website`

如果你没有使用 Github pages，则在这里放置项目的网页数据。

更多样例查看[`/website`](https://github.com/golang-standards/project-layout/blob/master/website/README.md)。

## 不应该出现的目录

### `/src`

有一些 Go 项目确实包含`src`文件夹，但通常只有在开发者是从 Java（这是 Java 中一个通用的模式）转过来的情况下才会有。如果可以的话请不要使用这种 Java 模式。你肯定不希望你的 Go 代码和项目看起来像 Java。:-)

不要将项目级别的`/src`目录与 Go 用于其工作空间的`/src`目录混淆，就像[How to Write Go Code](https://golang.org/doc/code.html)中描述的那样。`$GOPATH`环境变量指向当前的工作空间（默认情况下指向非 Windows 系统中的`$HOME/go`）。此工作空间包括顶级`/pkg`，`/bin`和`/src`目录。实际的项目最终变成`/src`下的子目录，因此，如果项目中有`/src`目录，则项目路径将会变成：`/some/path/to/workspace/src/your_project/src/your_code.go`。请注意，使用 Go 1.11，虽然可以将项目放在 GOPATH 之外，但这并不意味着使用此布局模式是个好主意。

## 徽章

- [Go Report Card](https://goreportcard.com/)：它会使用`gofmt`，`vet`，`gocyclo`，`golint`，`ineffassign`，`license`和`mispell`扫描项目中的代码。将`github.com/golang-standards/project-layout`替换为你的项目的引用。

[![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)

- ~~[GoDoc](http://godoc.org/)：提供 GoDoc 生成的文档的在线版本。更改链接以指向你的项目。~~

  [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)

- [pkg.go.dev](https://pkg.go.dev/)：它是 Go 官方提供的用于文档与包发现的新平台。你可以使用其[徽章生成工具](https://pkg.go.dev/badge/)创建项目徽章。

[![PkgGoDev](https://pkg.go.dev/badge/github.com/golang-standards/project-layout)](https://pkg.go.dev/github.com/golang-standards/project-layout)

- Release：它会显示你项目的最新版本号。更改 github 链接以指向你的项目。

[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## 注意

一个更加主观化的项目模板（包含示例/可复用的配置、脚本和代码）正在开发中（WIP：Work In Progress）。
