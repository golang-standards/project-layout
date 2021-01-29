# 표준 Go 프로젝트 레이아웃

번역:

* [English](README.md)
* [한국어 문서](README_ko.md)
* [简体中文](README_zh.md)
* [正體中文](README_zh-TW.md)
* [简体中文](README_zh-CN.md) - ???
* [Français](README_fr.md)
* [日本語](README_ja.md)
* [Español](README_es.md)

## 개요

이 레포지토리는 Go 애플리케이션 프로젝트의 기본 레이아웃입니다. 이 레포지토리는 코어 Go 개발팀에 의해 정의된 공식적인 표준은 아니지만, Go 생태계에서 역사적이거나 급부상중인 프로젝트 레이아웃 패턴들의 공통된 부분들입니다. 몇몇 패턴들은 다른 패턴들보다 유명합니다. 또한 이는 충분히 큰 현실세계 애플리케이션들에서 공통적으로 사용하는 여러 보조 디렉터리와 함께 다수의 작은 개선 사항들을 포함하고 있습니다.

만약 당신이 Go를 배우려하거나, 혼자 PoC나 토이 프로젝트를 만드는 거라면 이 프로젝트 레이아웃은 과한 것입니다. 아주 간단한 것부터 시작하세요 (`main.go` 파일 하나면 아주 충분합니다). 프로젝트가 성장하면서 당신의 코드가 잘 구조화 되도록 하는 것이 중요합니다. 그렇지 않으면 많은 숨겨진 종속성들과 전역 상태가 있는 지저분한 코드로 끝나게 될 것 입니다. 더 많은 사람들이 프로젝트에 참여할 때 더 많은 구조가 필요합니다. 그 때가 패키지/라이브러리를 관리하는 공통된 방법을 도입하는 것이 중요한 시점입니다. 당신이 오픈 소스 프로젝트를 가지고 있거나 다른 프로젝트가 당신의 프로젝트 레포지토리에서 코드를 임포트 하고있다면 프라이빗 (`internal` 로 일컫어지는) 패키지와 코드를 도입하는게 중요합니다. 이 레포지토리를 Clone하고, 당신에게 필요한것만 남긴다음 나머지 것들을 다 지우세요! 거기에 있다고 해서 다 사용해야하는 것은 아닙니다. 모든 프로젝트에서 다 사용하는 패턴은 없습니다. `vendor`패턴 조차 보편적이지 않습니다.

Go 1.14로 [`Go Modules`](https://github.com/golang/go/wiki/Modules) 은 최종적으로 프로덕션 준비가 완료되었습니다. 특별히 사용하지 않을 사유가 없다면 [`Go Modules`](https://blog.golang.org/using-go-modules) 를 사용하세요. 그러면 $GOPATH와 어디에 내 프로젝트를 놓을지 고민할 필요가 없습니다. 레포지토리에서 기본적인 `go.mod` 파일은 프로젝트가 GitHub에 호스팅되어있다고 가정합니다만, 필수사항은 아닙니다. 모듈 패스는 어느것이든 될 수 있지만 첫번째 모듈 패스 컴포넌트는 그 이름에 점이 있어야 합니다. (현재 버전의 Go는 더이상 이를 강제하지는 않지만, 만약 조금 더 오래된 버전을 사용하고 있다면 점이 없을때 빌드가 실패해도 놀라지 마세요). 이에 대해 더 알아보고 싶으시면 이슈 [`37554`](https://github.com/golang/go/issues/37554) 와 [`32819`](https://github.com/golang/go/issues/32819) 를 읽어보세요.

이 프로젝트 레이아웃은 의도적으로 일반적이며 특정 Go 패키지 구조를 강요하지 않습니다.

이 레포지토리는 커뮤니티의 노력입니다. 새 패턴을 찾았거나, 있는 패턴들 중 하나가 업데이트가 필요하다고 생각된다면 이슈를 열어주세요.

네이밍, 포매팅, 스타일에 대해 도움이 필요하시다면, [`gofmt`](https://golang.org/cmd/gofmt/) 와 [`golint`](https://github.com/golang/lint)를 돌리는 것으로 시작해보세요. 또한 Go 코드 스타일 가이드라인과 권장 사항들을 꼭 읽어보세요:

* https://talks.golang.org/2014/names.slide
* https://golang.org/doc/effective_go.html#names
* https://blog.golang.org/package-names
* https://github.com/golang/go/wiki/CodeReviewComments
* [Style guideline for Go packages](https://rakyll.org/style-packages) (rakyll/JBD)

자세한 배경 지식에 대해서는 [`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) 를 확인하세요.

다른 코드 구조 추천과 네이밍과 패키지 구성에 대해 더 알아보기:

* [GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)
* [GopherCon Russia 2018: Ashley McNamara + Brian Ketelsen - Go best practices.](https://www.youtube.com/watch?v=MzTcsI6tn-0)
* [GopherCon 2017: Edward Muller - Go Anti-Patterns](https://www.youtube.com/watch?v=ltqV6pDKZD8)
* [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)

패키지 지향 디자인 가이드라인과 설계 레이어에 대한 중국어 글
* [面向包的设计和架构分层](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md)

## Go 디렉터리

### `/cmd`

이 프로젝트의 메인 애플리케이션들입니다.

각 애플리케이션의 디렉터리명은 만들고싶은 실행 파일 이름과 같아야 합니다  (e.g., `/cmd/myapp`).

애플리케이션 디렉터리에 많은 코드를 넣지마세요. 다른 프로젝트들에서 임포트되고 사용될 것 같으면, `/pkg` 디렉터리에 있어야합니다. 코드가 재사용성이 없거나 다른사람들이 재사용하지 않기를 바란다면 코드를 `/internal` 디렉터리에 넣으세요. 다른사람들이 무얼 할지 놀라게 될거에요. 그러니 의도를 분명하게 밝혀주세요!

흔히 `/internal`와 `/pkg` 디렉터리 코드를 임포트, 호출만 하는 작은  `main` 함수는 흔히 볼 수 있습니다. 

예시는 [`/cmd`](cmd/README.md)를 보세요.

### `/internal`

개인적인 애플리케이션과 라이브러리 코드. 다른 사람들이 애플리케이션이나 라이브러리에서 임포트 하기를 원하지 않는 코드들입니다. 이 레이아웃 패턴은 Go 컴파일러 자체에 강제됩니다. 더 알아보고 싶다면 Go 1.4 [`release notes`](https://golang.org/doc/go1.4#internalpackages) 를 참고하세요. 이는 최상단 `internal` 디렉터리에 국한되는 것은 아닙니다. 프로젝트 트리의 모든 레벨에서 하나 이상의 `internal` 디렉터리를 가질 수 있습니다. 

Internal 패키지에서 공유되는 내부 코드와 공유되지 않는 내부 코드를 구분하기 위해 부가적인 구조체를 추가할 수 있습니다. 필수사항은 아니지만(특히 작은 프로젝트에서는), 의도된 패키지 사용법을 보여주는 시각적인 단서를 남기는 것이 좋습니다. 실제 에플리케이션 코드는 `/internal/app` 디렉터리에 넣고 (e.g., `/internal/app/myapp`) , 그 앱들에서 공유되는 코드들은 `/internal/pkg` 디렉터리 (e.g., `/internal/pkg/myprivlib`) 에 넣을 수 있습니다.

### `/pkg`

외부 애플리케이션에서 사용되어도 괜찮은 라이브러리 코드입니다 (e.g., `/pkg/mypubliclib`). 다른 프로젝트는 이 라이브러리들이 작동할거라고 예상하고 임포트 할 것 이므로, 여기에 무언가를 넣기 전에 두번 고민하세요 :-) `internal`디렉터리는 개인적인 패키지들이 임포트 불가능하도록 하는 더 좋은 방법인데, 이유는 Go가 이를 강제하기 떄문입니다. `/pkg` 디렉터리는 그 디렉터리 안의 코드가 다른 사람들에 의해 사용되어도 안전하다고 명시적으로 보여주는 좋은 방법입니다. Travis Jeffery의 [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) 블로그 포스트는 `pkg` 와 `internal` 디렉터리와 언제 쓰는게 맞을지에 대해 좋은 개요를 제공합니다. 

또한 루트 디렉터리에 많은 Go가 아닌 컴포넌트와 디렉터리를 포함하고 있다면 Go 코드를 한 곳에 모아서 다양한 Go 툴들을 쉽게 실행할 수 있습니다 (이 발표들에서 언급되었던것 처럼: GopherCon EU 2018의 [`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg),  [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) 와 [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)).

이 프로젝트 레이아웃 패턴을 사용하는 유명한 Go 레포지토리들을 보고 싶다면 [`/pkg`](pkg/README.md)를 보세요. 이는 흔한 레이아웃 패턴이나, 보편적으로 받아드려지는 것은 아니며 Go 커뮤니티의 일부는 이를 추천하지 않습니다.

앱 프로젝트가 정말 작고 추가적인 레벨의 중첩이 많이 효과적이지 않다면 사용하지 않아도 괜찮습니다 (정말로 원하지 않는 한 :-)). 프로젝트가 더 커지고 루트 디렉터리가 꽤 바빠질 때 고려해보세요 (특히 Go가 아닌 앱 컴포넌트가 많이 있다면).

### `/vendor`

애플리캐이션 종속성 (직접 혹은 새 빌트인 [`Go Modules`](https://github.com/golang/go/wiki/Modules) 피쳐와 같은 종속성 관리 도구로 관리되는).  `go mod vendor` 명령어는 `/vendor` 디렉터리를 만들어 줄 것 입니다. `go build` 명령어를 사용할 때 Go 1.14를 사용하지 않는다면 Go 1.14에서 기본으로 켜져있는 `-mod=vendor` 플래그를 추가해야 할 수 있습니다.

만약 당신이 라이브러리를 만들고 있다면 당신의 애플리케이션 종속성을 커밋하지 마세요.

[`1.13`](https://golang.org/doc/go1.13#modules)부터 Go는 모듈 프록시 기능(기본적으로 [`https://proxy.golang.org`](https://proxy.golang.org)를 모듈 프록시 서버로 사용하여)을 활성화하였습니다. 이것이 당신의 요구조건과 제한사항을 모두 만족하는지 보려면 [`여기`](https://blog.golang.org/module-mirror-launch) 를 읽어보세요. 만약 만족한다면, `vendor` 디렉터리가 전혀 필요하지 않을 것입니다.

## 서비스 애플리케이션 디렉터리

### `/api`

OpenAPI/Swagger 스펙들, JSON schema 파일들, 프로토콜 정의 파일들.

예시로  [`/api`](api/README.md) 디렉터리를 확인하세요. 

## 웹 애플리케이션 디렉터리

### `/web`

웹 플리케이션의 특정한 컴포넌트들: 정적 웹 에셋들, 서버 사이드 템플릿과 SPA들.

## 공통 애플리케이션 디렉터리

### `/configs`

설정 파일 템플릿이나 기본 설정들.

`confd` or `consul-template` 템플릿 파일들을 여기에 놓으세요.

### `/init`

시스템 init (systemd, upstart, sysv) 과 프로세스 매니저/슈퍼바이저 (runit, supervisord) 설정들.

### `/scripts`

빌드, 설치, 분석, 기타 작업을 위한 스크립트들.

이 스크립트들은 루트의 Makefile을 작고 간단하게 유지해줍니다 (e.g., [`https://github.com/hashicorp/terraform/blob/master/Makefile`](https://github.com/hashicorp/terraform/blob/master/Makefile)).

예시로  [`/scripts`](scripts/README.md) 디렉터리를 확인하세요.

### `/build`

패키징과 지속적인 통합(CI).

클라우드 (AMI), 컨테이너 (Docker), OS (deb, rpm, pkg) 패키지 설정과 스크립트를 `/build/package` 디렉터리에 넣으세요.

CI (travis, circle, drone) 설정과 스크립트를 `/build/ci`에 넣으세요. 몇몇 CI 도구들은 (e.g., Travis CI) 설정파일들 위치에 대해 굉장히 까다롭다는 것을 알아두세요. 설정 파일들을 `/build/ci` 에 넣고 CI 도구들이 기대하는 위치에 링크해보세요. (가능하다면).

### `/deployments`

IaaS, PaaS, 시스템과 컨테이너 오케스트레이션 배포 설정과 템플릿 (docker-compose, kubernetes/helm, mesos, terraform, bosh). 몇몇 레포지토리 (특히 쿠버네티스로 배포되는 앱들) 에서 이 디렉터리는 `/deploy` 라고 불립니다.

### `/test`

추가적인 외부 테스트 앱들과 테스트 데이터들. `/test` 디렉터리를 자유롭게 구조화하세요. 더 큰 프로젝트에서 데이터 서브디렉터리를 갖는것이 합당합니다. 예시로,  이 디렉터리에 무엇이 있는지 Go가 무시하려면 `/test/data` 또는 `/test/testdata` 를 만들면 됩니다. Go는 "." 나 "_" 로 시작하는 디렉터리들이나 파일들도 무시하니, 테스트 데이터 디렉터리 이름을 더 자유롭게 지을 수 있습니다. 

예시로 [`/test`](test/README.md) 디렉터리를 확인하세요.

## 다른 디렉터리들

### `/docs`

디자인과 사용자 문서들 (godoc이 만든 문서에 더불어).

예시로 [`/docs`](docs/README.md) 디렉터리를 확인하세요.

### `/tools`

이 프로젝트에서 사용하는 도구들. 이 도구들이 `/pkg` 나 `/internal` 디렉터리에서 코드를 임포트 할 수 있음을 알아두세요.

예시로 [`/tools`](tools/README.md) 디렉터리를 확인하세요.

### `/examples`

애플리케이션 혹은 공개된 라이브러리의 예시들.

예시로 [`/examples`](examples/README.md) 디렉터리를 확인하세요.

### `/third_party`

사용하는 외부 도구들, 포크된 코드들과 다른 서드 파티 유틸리티들 (e.g., Swagger UI).

### `/githooks`

Git hooks.

### `/assets`

레포지토리와 함께 사용될 에셋들 (이미지, 로고, 기타).

### `/website`

GitHub pages를 사용하고 있지 않다면 프로젝트의 웹사이트 데이터를 넣는 곳입니다.

예시로 [`/website`](website/README.md) 디렉터리를 확인하세요.

## 있으면 안되는 디렉터리

### `/src`

몇몇 Go 프로젝트들은 분명  `src` 폴더를 가지고 있으나, 이는 개발자들이 해당 패턴이 흔한 Java 세계에서 왔을 때 대부분 발생합니다. 만약 할 수 있다면 이 Java 패턴을 사용하지 않도록 해보세요. Go 코드나 Go 프로젝트가 Java처럼 보이는걸 원하지 않잖아요 :-)

프로젝트 레벨의  `/src` 디렉터리와 [`How to Write Go Code`](https://golang.org/doc/code.html)에서 설명된 것 처럼 Go에서 워크스페이스로 사용하는 `/src` 디렉터리와 혼동하지 마세요. `$GOPATH` 환경 변수는 당신의 (현재) 워크스페이스를 가리킵니다 (기본적으로 윈도우가 아닌 시스템에서는 `$HOME/go`를 가리킵니다). 이 워크스페이스는 최상위 `/pkg`, `/bin` , `/src` 디렉터리를 포함하고 있습니다. 실제 프로젝트는 결국 `/src`의 서브디렉터리이므로, 프로젝트 안에 `/src` 디렉터리가 있다면 프로젝트 주소는 이럴 것입니다: `/some/path/to/workspace/src/your_project/src/your_code.go`. Go 1.11에서는  `GOPATH` 밖에 프로젝트를 만드는게 가능하지만, 이 레이아웃 패턴을 사용하는게 좋다는 뜻은 아님을 알아두세요.


## 뱃지

* [Go Report Card](https://goreportcard.com/) - 이는 당신의 코드를 `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license`, `misspell`로 스캔합니다.  `github.com/golang-standards/project-layout` 를 당신의 프로젝트 주소로 바꾸세요.

* ~~[GoDoc](http://godoc.org) - 온라인 버전의 GoDoc 생성 문서를 제공합니다. 당신의 프로젝트를 가리키도록 링크를 바꾸세요.~~

* [Pkg.go.dev](https://pkg.go.dev) - Pkg.go.dev는 Go 검색 및 문서들을 위한 새로운 사이트입니다. [badge generation tool](https://pkg.go.dev/badge)을 사용해서 새로운 뱃지를 만들 수 있습니다.

* Release - 프로젝트의 최신 릴리즈 넘버를 보여줍니다. 당신의 프로젝트를 가리키도록 깃헙 링크를 바꾸세요.

[![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/golang-standards/project-layout)](https://pkg.go.dev/github.com/golang-standards/project-layout)
[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## 참고

샘플/재사용 가능한 설정, 스크립트와 코드를 포함하는 더 의견이 담긴 프로젝트 템플릿은 작업 중입니다.

