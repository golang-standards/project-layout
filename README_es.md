# Standard Go Project Layout

Traducciones:

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
* [Italiano](README_it.md)
* [Vietnamese](README_vi.md)
* [Українська](README_ua.md)
* [Indonesian](README_id.md)
* [हिन्दी](README_hi.md)
* [Беларуская](README_by.md)

## Resumen

Este es un diseño básico para proyectos de aplicaciones de Go. No es un estándar oficial definido por el equipo de desarrollo principal de Go; sin embargo, es un conjunto de patrones de diseño de proyectos históricos y emergentes comunes en el ecosistema Go. Algunos de estos patrones son más populares que otros. También tiene una serie de pequeñas mejoras junto con varios directorios de soporte comunes a cualquier aplicación del mundo real lo suficientemente grande.

Si está tratando de aprender Go o si está construyendo un PoC o un proyecto de juguete para usted, este diseño del proyecto es excesivo. Comience con algo realmente simple (un solo archivo `main.go` es más que suficiente). A medida que su proyecto crezca, tenga en cuenta que será importante asegurarse de que su código esté bien estructurado, de lo contrario, terminará con un código desordenado con muchas dependencias ocultas y un estado global. Cuando tenga más personas trabajando en el proyecto, necesitará aún más estructura. Ahí es cuando es importante introducir una forma común de administrar paquetes / bibliotecas. Cuando tienes un proyecto de código abierto o cuando sabes que otros proyectos importan el código del repositorio de tu proyecto, es cuando es importante tener paquetes y código privados (también conocidos como `internal`). ¡Clona el repositorio, guarda lo que necesitas y elimina todo lo demás! El hecho de que esté allí no significa que tenga que usarlo todo. Ninguno de estos patrones se utiliza en todos los proyectos. Incluso el patrón `vendor` no es universal.

Con Go 1.14, [`los módulos Go`](https://go.dev/wiki/Modules) están finalmente listos para la producción. Use [`los módulos Go`](https://blog.golang.org/using-go-modules) a menos que tenga una razón específica para no usarlos y, si lo hace, no debe preocuparse por $ GOPATH y dónde coloque su proyecto. El archivo `go.mod` básico en el repositorio asume que su proyecto está alojado en GitHub, pero no es un requisito. La ruta del módulo puede ser cualquier cosa, aunque el primer componente de la ruta del módulo debe tener un punto en su nombre (la versión actual de Go ya no lo aplica, pero si está utilizando versiones un poco más antiguas, no se sorprenda si sus compilaciones fallan sin eso). Consulte los problemas [`37554`](https://github.com/golang/go/issues/37554) y [`32819`](https://github.com/golang/go/issues/32819) si desea obtener más información al respecto.

Este diseño de proyecto es intencionalmente genérico y no intenta imponer una estructura de paquete Go específica.

Este es un esfuerzo comunitario. Abra un problema si ve un patrón nuevo o si cree que uno de los patrones existentes debe actualizarse.

Si necesita ayuda con el nombre, el formato y el estilo, comience ejecutando `gofmt` y `golint`. También asegúrese de leer estas recomendaciones y pautas de estilo del código Go:
* https://talks.golang.org/2014/names.slide
* https://golang.org/doc/effective_go.html#names
* https://blog.golang.org/package-names
* https://go.dev/wiki/CodeReviewComments
* [Guía de estilo para paquetes Go](https://rakyll.org/style-packages) (rakyll/JBD)

Consulte [`Diseño de proyecto de Go`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) para obtener información adicional sobre los antecedentes.

Más sobre cómo nombrar y organizar paquetes, así como otras recomendaciones de estructura de código:
* [GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)
* [GopherCon Russia 2018: Ashley McNamara + Brian Ketelsen - Go best practices.](https://www.youtube.com/watch?v=MzTcsI6tn-0)
* [GopherCon 2017: Edward Muller - Go Anti-Patterns](https://www.youtube.com/watch?v=ltqV6pDKZD8)
* [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)

Una publicación en idioma chino sobre las pautas de diseño orientado a paquetes y la capa de arquitectura

* [面向包的设计和架构分层](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md)

## Directorios en Go

### `/cmd`

Principales aplicaciones de este proyecto.

El nombre del directorio de cada aplicación debe coincidir con el nombre del ejecutable que desea tener (p.ej, `/cmd/myapp`).

No pongas mucho código en el directorio de la aplicación. Si cree que el código se puede importar y utilizar en otros proyectos, entonces debería residir en el directorio `/pkg`. Si el código no es reutilizable o si no desea que otros lo reutilicen, coloque ese código en el directorio `/internal`. Se sorprenderá de lo que harán los demás, ¡así que sea explícito sobre sus intenciones!

Es común tener una pequeña función `main` que importa e invoca el código de los directorios `/internal` y `/pkg` y nada más.

Consulte el directorio [`/cmd`](cmd/README.md) para ver ejemplos.

### `/internal`

Aplicación privada y código de biblioteca. Este es el código que no desea que otros importen en sus aplicaciones o bibliotecas. Tenga en cuenta que el propio compilador de Go aplica este patrón de diseño. Consulte [`las notas de la versión`](https://golang.org/doc/go1.4#internalpackages) Go 1.4 para obtener más detalles. Tenga en cuenta que no está limitado al directorio `internal` de nivel superior. Puede tener más de un directorio `internal` en cualquier nivel del árbol de su proyecto.

Opcionalmente, puede agregar un poco de estructura adicional a sus paquetes internos para separar su código interno compartido y no compartido. No es necesario (especialmente para proyectos más pequeños), pero es bueno tener pistas visuales que muestren el uso previsto del paquete. Su código de aplicación real puede ir en el directorio `/internal/app` (p.ej, `/internal/app/myapp`) y el código compartido por esas aplicaciones en el directorio `/internal/pkg` (por ejemplo, `/internal/pkg/myprivlib`).

### `/pkg`

Código de biblioteca que puede usar aplicaciones externas (por ejemplo, `/pkg/mypubliclib`). Otros proyectos importarán estas bibliotecas esperando que funcionen, así que piénselo dos veces antes de poner algo aquí :-) Tenga en cuenta que el directorio `internal` es la mejor manera de asegurarse de que sus paquetes privados no se puedan importar porque Go lo aplica. El directorio `/pkg` sigue siendo una buena forma de comunicar explícitamente que el código de ese directorio es seguro para que lo utilicen otros. La publicación del blog de Travis Jeffery, [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/), proporciona una buena descripción de los directorios `pkg` e `internal` y de cuándo podría tener sentido usarlos.

También es una forma de agrupar el código de Go en un solo lugar cuando su directorio raíz contiene muchos componentes y directorios que no son de Go, lo que facilita la ejecución de varias herramientas de Go (como se menciona en estas charlas: [`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg) de GopherCon EU 2018, [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) y [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)).

Consulte el directorio [`/pkg`](pkg/README.md) si desea ver qué repositorios Go populares utilizan este patrón de diseño de proyecto. Este es un patrón de diseño común, pero no se acepta universalmente y algunos miembros de la comunidad de Go no lo recomiendan.

Está bien no usarlo si el proyecto de su aplicación es realmente pequeño y donde un nivel adicional de anidamiento no agrega mucho valor (a menos que realmente quiera :-)). Piénselo cuando sea lo suficientemente grande y su directorio raíz esté bastante ocupado (especialmente si tiene muchos componentes de aplicaciones que no son de Go).

### `/vendor`

Dependencias de aplicaciones (administradas manualmente o mediante su herramienta de administración de dependencias favorita, como la nueva función de [`Módulos Go`](https://go.dev/wiki/Modules) integrada). El comando `go mod vendor` creará el directorio `/vendor` por usted. Tenga en cuenta que es posible que deba agregar la marca `-mod=vendor` a su comando `go build` si no está usando Go 1.14 donde está activado de forma predeterminada.

No adicione las dependencias de su aplicación si está creando una biblioteca.

Tenga en cuenta que desde [`1.13`](https://golang.org/doc/go1.13#modules) Go también habilitó la función de proxy del módulo (utilizando [`https://proxy.golang.org`](https://proxy.golang.org) como su servidor proxy del módulo de forma predeterminada). Lea más sobre esto [`aquí`](https://blog.golang.org/module-mirror-launch) para ver si se ajusta a todos sus requisitos y limitaciones. Si es así, no necesitará el directorio `vendor` en absoluto.

## Directorios de Aplicaciones de Servicio

### `/api`

Especificaciones de OpenAPI/Swagger, archivos de esquema JSON, archivos de definición de protocolo.

Consulte el directorio [`/api`](api/README.md) para ver ejemplos.

## Directorios de Aplicaciones Web

### `/web`

Componentes específicos de la aplicación web: activos web estáticos, plantillas del lado del servidor y SPAs.

## Directorios de aplicaciones comunes

### `/configs`

Plantillas de archivos de configuración o configuraciones predeterminadas.

Ponga aquí sus archivos de plantilla `confd` o `consul-template`.

### `/init`

Configuraciones de inicio del sistema (systemd, upstart, sysv) y administración de procesos (runit, supervisord).

### `/scripts`

Scripts para realizar varias operaciones de construcción, instalación, análisis, etc.

Estos scripts mantienen el Makefile de nivel raíz pequeño y simple (p.ej, [`https://github.com/hashicorp/terraform/blob/main/Makefile`](https://github.com/hashicorp/terraform/blob/main/Makefile)).

Consulte el directorio `/scripts` para ver ejemplos.

### `/build`

Empaquetado e Integración Continua.

Coloque sus configuraciones de paquetes de nube (AMI), contenedor (Docker), SO (deb, rpm, pkg) y scripts en el directorio `/build/package`.

Coloque sus configuraciones y scripts de CI (travis, circle, drone) en el directorio `/build/ci`. Tenga en cuenta que algunas de las herramientas de CI (p.ej, Travis CI) son muy exigentes con la ubicación de sus archivos de configuración. Intente poner los archivos de configuración en el directorio `/build/ci` vinculándolos a la ubicación donde las herramientas de CI los esperan (cuando sea posible).

### `/deployments`

Configuraciones y plantillas de implementación de IaaS, PaaS, sistema y orquestación de contenedores (docker-compose, kubernetes/helm, mesos, terraform, bosh). Tenga en cuenta que en algunos repositorios (especialmente las aplicaciones implementadas con kubernetes) este directorio se llama `/deploy`.

### `/test`

Aplicaciones de prueba externas adicionales y datos de prueba. Siéntase libre de estructurar el directorio `/test` como desee. Para proyectos más grandes, tiene sentido tener un subdirectorio de datos. Por ejemplo, puede tener `/test/data` o `/test/testdata` si necesita Go para ignorar lo que hay en ese directorio. Tenga en cuenta que Go también ignorará los directorios o archivos que comiencen con "." o "_", por lo que tiene más flexibilidad en términos de cómo nombrar su directorio de datos de prueba.

Consulte el directorio [`/test`](test/README.md) para ver ejemplos.

## Otros Directorios

### `/docs`

Diseño y documentos de usuario (además de su documentación generada por godoc).

Consulte el directorio [`/docs`](docs/README.md) para ver ejemplos.

### `/tools`

Herramientas de apoyo para este proyecto. Tenga en cuenta que estas herramientas pueden importar código de los directorios `/pkg` y `/internal`.

Consulte el directorio [`/tools`](tools/README.md) para ver ejemplos.

### `/examples`

Ejemplos para sus aplicaciones y / o bibliotecas públicas.

Consulte el directorio [`/examples`](examples/README.md) para ver ejemplos.

### `/third_party`

Herramientas de ayuda externas, código bifurcado y otras utilidades de terceros (p.ej, Swagger UI).

### `/githooks`

Ganchos de Git.

### `/assets`

Otros activos que acompañan a su repositorio (imágenes, logotipos, etc.).

### `/website`

Este es el lugar para colocar los datos del sitio web de su proyecto si no está utilizando GitHub pages.

Consulte el directorio [`/website`](website/README.md) para ver ejemplos.

## Directorios que No Deberías Tener

### `/src`

Algunos proyectos de Go tienen una carpeta `src`, pero suele ocurrir cuando los desarrolladores vienen del mundo Java, donde es un patrón común. Si puede ayudarse a sí mismo, intente no adoptar este patrón de Java. Realmente no quieres que tu código Go o proyectos Go se parezcan a Java :-)

No confunda el directorio `/src` de nivel de proyecto con el directorio `/src` que usa Go para sus áreas de trabajo, como se describe en [`How to Write Go Code`](https://golang.org/doc/code.html). La variable de entorno `$GOPATH` apunta a su área de trabajo (actual) (de forma predeterminada apunta a `$HOME/go` en sistemas que no son Windows). Este espacio de trabajo incluye los directorios de nivel superior `/pkg`, `/bin` y `/src`. Su proyecto real termina siendo un subdirectorio en `/src`, por lo que si tiene el directorio `/src` en su proyecto, la ruta del proyecto se verá así: `/some/path/to/workspace/src/your_project/src/your_code.go`. Tenga en cuenta que con Go 1.11 es posible tener su proyecto fuera de su `GOPATH`, pero aún así no significa que sea una buena idea utilizar este patrón de diseño.

## Badges

* [Go Report Card](https://goreportcard.com/) - Escaneará su código con `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` y `misspell`. Cambie `github.com/golang-standards/project-layout` por la referencia de su proyecto.

* ~~[GoDoc](http://godoc.org) - Proporcionará una versión en online de su documentación generada por GoDoc. Cambie el enlace para que apunte a su proyecto.~~

* [Pkg.go.dev](https://pkg.go.dev) - Pkg.go.dev es un nuevo destino para Go discovery & docs. Puede crear una insignia con [la herramienta de generación de insignias](https://pkg.go.dev/badge).

* Release - mostrará el número de versión más reciente de su proyecto. Cambie el enlace de github para que apunte a su proyecto.

[![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/golang-standards/project-layout)](https://pkg.go.dev/github.com/golang-standards/project-layout)
[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## Notas

Una plantilla de proyecto con mayor criterio sobre configuraciones, scripts y códigos de ejemplos está en construcción (WIP).
