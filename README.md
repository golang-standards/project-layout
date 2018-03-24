# Standard Go Project Layout

This is a basic layout for Go application projects. It represents the most common directory structure with a number of small enhancements along with several supporting directories common to any real world application. 

This project layout is intentionally generic and it doesn't try to impose a specific Go package structure.

Clone the repository, keep what you need and delete everything else!

[Go Project Layout](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) - additional background information.

## Go Directories

### `/cmd`

Main applications for this project.

The directory name for each application should match the name of the executable you want to have (e.g., `/cmd/myapp`).

Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the `/pkg` directory. If the code is not reusable or if you don't want others to reuse it, put that code in the `/internal` directory. You'll be surprised what others will do, so be explicit about your intentions!

It's common to have a small `main` function that imports and invokes the code from the `/internal` and `/pkg` directories and nothing else.

### `/internal`

Private application and library code. This is the code you don't want others importing in their applications or libraries.

Put your actual application code in the `/internal/app` directory (e.g., `/internal/app/myapp`) and the code shared by those apps in the `/internal/pkg` directory (e.g., `/internal/pkg/myprivlib`).

### `/pkg`

Library code that's safe to use by external applications (e.g., `/pkg/mypubliclib`).

Other projects will import these libraries expecting them to work, so think twice before you put something here :-)

### `/vendor`

Application dependencies (managed manually or by your favorite dependency management tool).

Don't commit your application dependencies if you are building a library.

## Service Application Directories

### `/api`

OpenAPI/Swagger specs, JSON schema files, protocol definition files.

## Web Application Directories

### `/web`

Web application specific components: static web assets, server side templates and SPAs.

## Common Application Directories

### `/configs`

Configuration file templates or default configs.

Put your `confd` or `consule-template` template files here.

### `/init`

System init (systemd, upstart, sysv) and process manager/supervisor (runit, supervisord) configs.

### `/scripts`

Scripts to perform various build, install, analysis, etc operations.

These scripts keep the root level Makefile small and simple.

### `/build`

Packaging and Continous Integration.

Put your cloud (AMI), container (Docker), OS (deb, rpm, pkg) package configurations and scripts in the `/build/package` directory.

Put your CI (travis, circle, drone) configurations and scripts in the `/build/ci` directory.

### `/deployments`

IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh).

### `/test`

Additional external test apps and test data.

## Other Directories

### `/docs`

Design and user documents (in addition to your godoc generated documentation).

### `/tools`

Supporting tools for this project. Note that these tools can import code from the `/pkg` and `/internal` directories.

### `/examples`

Examples for your applications and/or public libraries.

### `/third_party`

External helper tools, forked code and other 3rd party utilities (e.g., Swagger UI).

### `/githooks`

Git hooks.

### `/assets`

Other assets to go along with your repository.

## Directories You Shouldn't Have

### `/src`

Some Go projects do have a `src` folder, but it usually happens when the devs came from the Java world where it's a common pattern. If you can help yourself try not to adopt this Java pattern. You really don't want your Go code or Go projects to look like Java :-)


## Badges

* [Go Report Card](https://goreportcard.com/) - It will scan your code with `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` and `misspell`. Replace `github.com/golang-standards/project-layout` with your project reference.

* [GoDoc](http://godoc.org) - It will provide online version of your GoDoc generated documentation. Change the link to point to your project.

* Release - It will show the latest release number for your project. Change the github link to point to your project.

[![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)
[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## Notes

A more opinionated project template with sample/reusable configs, scripts and code is a WIP.

