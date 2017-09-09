# Standard Go Project Layout

This is a basic layout for Go application projects. It represents the most common directory structure with a number of small enhancements along with several supporting directories common to any real world application.

Clone the repository, keep what you need and delete everything else!

## Go Directories

* `/cmd`

Main applications for this project.

The directory name for each application should match the name of the executable you want to have (e.g., `/cmd/myapp`).

Don't put a lot of code in the application directory unless you think that code can be imported and used in other projects. If this is the case then the code should live in the `/pkg` directory.

It's common to have a small main function that imports and invokes the code from the `/internal` and `/pkg` directories.

* `/internal`

Private application and library code.

Put your actual application code in the `/internal/app` directory (e.g., `/internal/app/myapp`) and the code shared by those apps in the `/internal/pkg` directory (e.g., `/internal/pkg/myprivlib`).

* `/pkg`

Library code that's safe to use by external applications (e.g., `/pkg/mypubliclib`).

Other projects will import these libraries expecting them to work, so think twice before you put something here :-)

* `/vendor`

Application dependencies (managed manually or by your favorite dependency management tool).

Don't commit your application dependencies if you are building a library.

## Service Application Directories

* `/api`

OpenAPI/Swagger specs, JSON schema files, protocol definition files.

## Web Application Directories

* `/web`

Web application specific components: static web assets, server side templates and SPAs.

## Common Application Directories

* `/configs`

Configuration file templates or default configs.

Put your `confd` or `consule-template` template files here.

* `/init`

System init (systemd, upstart, sysv) and process manager/supervisor (runit, supervisord) configs.

* `/scripts`

Scripts to perform various build, install, analysis, etc operations.

These scripts keep the root level Makefile small and simple.

* `/build`

Packaging and Continous Integration.

Put your cloud (AMI), container (Docker), OS (deb, rpm, pkg) package configurations and scripts in the `/build/package` directory.

Put your CI (travis, circle, drone) configurations and scripts in the `/build/ci` directory.

* `/deployments`

IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform).

* `/test`

Additional external test apps and test data.

## Other Directories

* `/docs`

Design and user documents.

* `/tools`

Supporting tools for this project. Note that these tools can import code from the `/pkg` and `/internal` directories.

* `/examples`

Examples for your applications and/or public libraries.

* `/third_party`

External helper tools, forked code and other 3rd party utilities (e.g., Swagger UI).

* `/githooks`

Git hooks.

* `/assets`

Other assets to go along with your repository.





