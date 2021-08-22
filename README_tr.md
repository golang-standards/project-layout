# Standart Go Projesi Düzeni

Çeviriler:

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

## Genel Bakış:

Bu proje Go projeleri için basit bir yerleşim düzenidir. `Go geliştirici takımının belirlediği, resmi standartları değil`. Go ekosistemindeki, ortak geçmişi ve zaman içinde ortaya çıkmış şablonları içerir. Bu şablonlardan bazıları diğerlerinden daha popülerdir. Bir dizi iyileştirmenin yanı sıra yeterli büyüklükte herhangi bir gerçek dünya uygulamasına özgü birkaç destekli dizin de içerir.

**`Eğer Go'yu öğrenmeye çalışıyorsanız veya PoC (Proof of Concept) ya da kendinize deneme bir proje yapıyorsanız bu klasör düzeni sizin için ideal olmaya bilir. Çok basit bir şeyle başlayın (tek bir` main.go `dosyası ve` go.mod `fazlasıyla yeterli olucaktır).`** Projeniz büyüdükçe projenin iyi yapılandırıldığından emin olmanın önemli olduğunu unutmayın, yoksa bir sürü gizli bağımlılık (dependency) ve bir sürü, her yerden erişilebilen, `global` değişkenlerle dolu dağınık bir kodla karşılaşabilirsiniz. Proje üstünde fazla kişi çalıştığında projeyi çok daha fazla yapılandırmanız gerekebilir. İşte tam da bu durumda paketleri/kütüphaneleri idare edebilmek için ortaya ortak bir yol koymak gerekir. Açık kaynak bir projeniz olduğunda ya da başkalarının sizin projenizi kullandıklarını bildiğinizde projenizde özel paketlerin ve kodların (diğer adıyla, `internal` klasörü) olması önemlidir. Bu projeyi klonlayın, ihtiyacınız olanları bırakın ve diğer bütün her şeyi silin. Bütün klasörlerin burada olması kullanmanız gerektiği anlamına gelmez. `vendor` klasörü bile herkes tarafından kullanılan bir şey değildir.

Go 1.14 ile birlikte [`Go Modules`](https://github.com/golang/go/wiki/Modules) özelliği sonunda `production` ortamında kullanılması için hazır oldu. Eğer karşı bir nedeniniz yoksa [`Go Modules`](https://blog.golang.org/using-go-modules) kullanın. Eğer kullanırsanız `$GOPATH` ile ve projenizin nereye koyulacağıyla alakalı endişeler duymazsınız. Projenizde bulunan basit `go.mod` dosyası projenizin GitHub'da barındırıldığını varsayar ama bu bir gereklilik değildir. Modül yolu her şey olabilir ama modül yolunun ilk parçasının içinde bir nokta olmalıdır (Go'nun şimdiki versiyonu artık bu konuda zorlamıyor ama eğer daha eski versiyonları kullanırsanız ve derleme işleminde hata alırsanız şaşırmayın). Daha fazla bilgi için [`37554`](https://github.com/golang/go/issues/37554) ve [`32819`](https://github.com/golang/go/issues/32819) hata bildirimlerine bakabilirsiniz.

"Bu proje düzeni genel bir düzendir ve belirli bir Go paketi yapısını empoze etmeye çalışmaz".

Bu proje bir topluluk eseridir. Eğer yeni bir şablon ve ya düzen ile karşılaşırsanız ya da olan şablonların, düzenlerin güncellenmesi gerektiğini düşünüyorsanız bir hata bildirimi açınız.

Eğer isimlendirmekle, biçimlendirmeyle ve stilize etmeyle ilgili yardıma ihtiyacınız varsa [`gofmt`](https://golang.org/cmd/gofmt/) ve [`golint`](https://github.com/golang/lint) yazılımlarını çalıştırmayı deneyin. Ayrıca aşağıdaki Go kod stili rehberlerini ve önerilerini okuduğunuzdan emin olun.    
* https://talks.golang.org/2014/names.slide
* https://golang.org/doc/effective_go.html#names
* https://blog.golang.org/package-names
* https://github.com/golang/go/wiki/CodeReviewComments
* [Style guideline for Go packages](https://rakyll.org/style-packages) (rakyll/JBD)

Ek bilgi için [`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) adlı yazıya bakabilirsin.

Paketlerin adlandırılması ve düzenlenmesi ile diğer kod yapısı önerileri hakkında daha fazla bilgi:
* [GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)
* [GopherCon Russia 2018: Ashley McNamara + Brian Ketelsen - Go best practices.](https://www.youtube.com/watch?v=MzTcsI6tn-0)
* [GopherCon 2017: Edward Muller - Go Anti-Patterns](https://www.youtube.com/watch?v=ltqV6pDKZD8)
* [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)

Paket odaklı dizayn rehberi ve Mimari katmanı ile alakalı Çince bir yazı:
* [面向包的设计和架构分层](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md)

## Go Klasörleri

### `/cmd`

Projenin ana uygulamalarını içerir.

Her uygulamanın klasör adı, eklemek istediğiniz uygulamanın adıyla eşleşmelidir. (örneğin, `/cmd/myapp`)

Bu klasöre çok fazla kod koymanız iyi olmaz. Eğer kodların başka projeler tarafından kullanılabileceğini düşünüyorsanız, o kodları `/pkg` klasörüne koymalısınız. Eğer kod yeniden kullanılabilecek değilse ya da diğerlerinin yeniden kullanmasını istemiyorsanız, bu kodları `/internal` klasörüne koymalısınız. Niyetiniz hakkında açık olun, diğer geliştiricilerin kodunuzu ne için kullanabileceğine şaşıracaksınız!

`/internal` ve `/pkg` klasörlerinden kod çağıran küçük bir `main` fonksiyonunuzun olması yaygındır.

Örnekler için [`/cmd`](cmd/README.md) klasörüne bakabilirsiniz.

### `/internal`

Özel uygulama ve kütüphane kodunu içerir. Buradaki kodlar diğer geliştiricilerin kendi projelerinde kullanmasını istemediğiniz kodlardır. Bu yerleşim düzeninin Go derleyicisinin kendisi tarafından uygulandığına dikkat edin. Ayrıntılar için Go 1.4 [`sürüm notları`](https://golang.org/doc/go1.4#internalpackages). Üst katmandaki `internal` klasörü ile sınırlı olmadığınızı unutmayın, projenizin herhangi bir katmanında birden fazla `internal` klasörüne sahip olabilirsiniz.

Paylaşılan ve paylaşılmayan kodunuzu ayırmak için isteğe bağlı olarak ekstra yapılar ekleyebilirsiniz. Bu gerekli değildir (özellikle küçük projeler için) ancak amaçlanan paket kullanımını gösteren görsel ipuçlarına sahip olmak güzel bir şeydir. Asıl uygulama kodunuz `/internal/app` klasöründe yer alabilir (örneğin, `/internal/app/benimuygulamam`) ve bu kodlar `/internal/pkg` klasöründe yer alan kodlar (örneğin, `/internal/pkg/benimözelkütüphanem`) ile paylaşılabilir.

### `/pkg`

Diğerleriyle paylaşmak istediğiniz kodu içerir (örneğin, `/pkg/myopenlibrary`). Diğer projeler bu kütüphaneleri çalışacağını tahmin ederek kullanırlar. Yani buraya bir şey koyarken iki kere düşünün :-) Unutmayın ki `internal` klasörü başka projeler tarafından kullanılmasını istemediğiniz kodlar için daha iyi bir yerdir çünkü bu Go tarafından zorunlu olarak uygulanır. `/pkg` klasörü içindeki kodun başkaları tarafından kullanılmasının güvenli olduğunu açıkça belirtmek için hala iyi bir yoldur. Travis Jeffery tarafından yazılan [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) blog yazısı `pkg` ve `internal` klasörleri ve onları ne zaman kullanmanın mantıklı olacağına dair genel bakış sağlar.

Ayrıca ana klasörünüz çok sayıda Go kodu olmayan dosya ve klasör içerdiğinde Go kodunu tek bir yerde gruplamanın yoludur ve bazı Go araçlarını kullanmayı daha kolay hale getirir (bu konuşmalarda bahsedildiği gibi: [`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg) GopherCon Avrupa 2018 konferansından, [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) ve [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)).

Hangi popüler Go projelerinin bu düzeni kullandığını merak ediyorsanız [`/pkg`](pkg/README.md) klasörüne bakabilirsiniz. Bu genel bir yerleşim düzenidir ama herkes tarafından kabul edilmeyebilir ve bazı Go toplulukları bu yerleşim düzenini tavsiye etmeyebilir.

Eğer projenizde ekstra bir katmana gerek yoksa ya da projeniz çok küçükse burayı kullanmamanızda bir sakınca olmaz (çok istiyorsanız kullanın). Projeniz yeterince büyük olduğunda ve ana klasörünüz karışmaya başladığında (özellikle çok fazla Go kodu olmayan dosyanız olduğunda) bunu kullanmayı düşünebilirsiniz.

### `/vendor`

Uygulama bağımlılıkları (el ile yönetilen ya da yeni bir özellik olan  [`Go Modules`](https://github.com/golang/go/wiki/Modules) gibi favori bağımlılık yönetim aracınızla yönetilen). `go mod vendor` komutu size yeni bir `/vendor` klasörü yaratır. Unutmayın eğer Go 1.14 kullanmıyorsanız `go build` komutuna `-mod=vendor` parametresi vermeniz gerekebilir Go 1.14 bunu varsayılan olarak yapar.

Eğer bir kütüphane yazıyorsanız bağımlılıklarınızı `commit` etmeyin.

Not olarak [`1.13`](https://golang.org/doc/go1.13#modules) versiyonundan itibaren Go modüller için `proxy` özelliğini aktif hale getirdi (varsayılan olarak [`https://proxy.golang.org`](https://proxy.golang.org) adresini kullanır). [`Buradan`](https://blog.golang.org/module-mirror-launch) gereksinimlerinize ve kısıtlamalarınıza uyup uymadığı hakkında daha fazla bilgi edinebilirsiniz. Eğer bu size uyarsa `vendor` klasörüne çokta ihtiyacınız olmayacaktır.

## Servis Uygulaması Klasörleri

### `/api`

OpenAPI/Swagger dosyaları, JSON şema dosyaları, protokol tanımlama dosyaları.

Örnekler için [`/api`](api/README.md) klasörüne bakabilirsiniz.

## Web Uygulaması Klasörleri

### `/web`

Spesifik web uygulaması parçaları: statik web varlıkları, sunucu taraflı şablonlar ve SPA'lar.

## Genel Uygulama Klasörleri

### `/configs`

Konfigürasyon dosya şablonları veya varsayılan konfigürasyonlar.

`confd` veya `consul-template` dosyalarını buraya koyabilirsiniz.

### `/init`

Sistem başlangıcı (systemd, upstart, sysv) ve işlem yöneticisi (runit, supervisord) konfigürasyonları.

### `/scripts`

Çeşitli derleme, yükleme, analiz ve benzeri işlemleri gerçekleştirmek için olan komut dosyaları.

Bu dosyalar ana katmandaki Makefile dosyasını küçük ve basit tutar. (örneğin, [`https://github.com/hashicorp/terraform/blob/master/Makefile`](https://github.com/hashicorp/terraform/blob/master/Makefile)).

Örnekler için [`/scripts`](scripts/README.md) klasörüne bakabilirsiniz.

### `/build`

Paketleme ve Sürekli Entegrasyon.

`/build/package` klasörüne bulut (AMI), konteyner (Docker), işletim sistemi (deb, rpm, pkg) paket konfigürasyonlarını ve komut dosyalarını koyabilirsiniz.

`/build/ci` klasörüne de Sürekli Entegrasyon (travis, circle, drone) konfigürasyonlarını ve komut dosyalarını koyabilirsiniz. Önemli olarak bazı sürekli entegrasyon araçları (örneğin, Travis CI) konfigürasyon dosyalarının klasörleriyle alakalı seçici olabiliyor. Konfigürasyon dosyalarını `/build/ci` klasörüne koymayı deneyin ve dosyaları sürekli entegrasyon araçlarının bekledikleri yere bağlayın (mümkün olduğunda).

### `/deployments`

IaaS, Paas, sistem ve konteyner orkestrasyon yayınlama konfigürasyonlarını ve şablonlarını (docker-compose, kubernetes/helm, mesos, terraform, bosh) bu klasöre koyabilirsiniz. Önemli olarak bazı projelerde (özellikle kubernetes ile yayınlanan projeler) bu klasörün adı `/deploy` olabilir.

### `/test`

Ek harici test uygulamaları ve test verileri. `/test` klasörünün içini istediğiniz gibi yapılandırmada özgür hissedin. Daha büyük projelerde test verileri için alt klasörler açmak mantıklı olabilir. Örneğin, eğer Go'nun test dosyalarını görmezden gelmesini istiyorsanız bu dosyalar için `/test/data` veya `/test/testdata` adlı klasörlere sahip olabilirsiniz. Not olarak Go "." veya "_" ile başlayan klasörleri ve dosyalarıda görmezden gelir, bu sayede test klasörlerinizi ve dosyalarınızı isimlendirmede daha esnek olabilirsiniz.

Örnekler için [`/test`](test/README.md) klasörüne bakabilirsiniz.

## Diğer Klasörler

### `/docs`

Tasarım ve kullanıcı dökümanları (godoc ile oluşturulmuş dökümanlara ek olarak)

Örnekler için [`/docs`](docs/README.md) klasörüne bakabilirsiniz.

### `/tools`

Bu klasöre projen için destekleyici araçları koyabilirsiniz. Unutmayın bu araçlar `/pkg` ve `/internal` klasörlerindeki kodları kullanabilirler.

Örnekler için [`/tools`](tools/README.md) klasörüne bakabilirsiniz.

### `/examples`

Bu klasöre yaptığınız uygulamalar ya da kütüphaneler için örnekleri koyabilirsiniz.

Örnekler için [`/examples`](examples/README.md) klasörüne bakabilirsiniz.

### `/third_party`

Dış yardımcı araçlar, çatallanmış kod (forked code) ve diğer 3. parti yardımcı şeyler. (örneğin, Swagger UI)

### `/githooks`

Git hooks.

### `/assets`

Bu klasöre resim, logo ve benzeri kaynakları koyabilirsiniz.

### `/website`

Eğer GitHub Pages kullanmıyorsanız burası projenin websitesinin dosyalarının konulması için doğru adres.

Örnekler için [`/website`](website/README.md) klasörüne bakabilirsiniz.

## Sahip olmamanız gereken klasörler

### `/src`

Bazı Go projelerinde `src` adlı bir klasör görebilirsiniz, ama bu çoğunlukla `src` klasörü kullanmanın genel bir kalıp olduğu Java dünyasından gelen geliştiricilerde olur. Eğer yapabilirseniz bu Java kalıbını benimsememeye çalışın. Gerçekten Go kodunuzun ya da Go projenizin Java gibi gözükmesini istemezsiniz :-)

Proje seviyesindeki `/src` klasörü ile [`How to Write Go Code`](https://golang.org/doc/code.html) adresinde belirtilen Go'nun çalışma alanları için kullandığı `/src` klasörünü karıştırmayın. `$GOPATH` ortam değişkeni sizin çalışma alanınıza (şu anki) işaret eder (Windows olmayan sistemlerde varsayılan olarak `$HOME/go` klasörüne işaret eder). Bu çalışma alanı içerisinde en üst katmandaki `/pkg`, `/bin` ve `/src` klasörlerini barındırır. Asıl projeniz `/src` klasörü altında bir alt dizin olur. Eğer projende `/src` klasörüne sahipseniz, projenizin dosya yolu büyük ihtimal şöyle gözükecek `/calismaalaniniz/src/projeniz/src/kodunuz.go`. Not olarak Go 1.11 ile `$GOPATH` klasörü dışında projeler açabilirsiniz ama bunu yapabilmeniz bu yerleşim düzeninin kullanılmasının iyi bir fikir olduğu anlamına gelmez.

## Rozetler

* [Go Report Card](https://goreportcard.com/) - Kodunuzu  `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` ve `misspell` yazılımları ile tarar. `github.com/golang-standards/project-layout` linkini kendi proje linkiniz ile değiştirin.

    [![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)

* ~~[GoDoc](http://godoc.org) - Projenizin GoDoc tarafından yaratılmış online bir dökümantasyonunu çıkartır. İşaret ettiği linki kendi proje linkiniz ile değiştirin.~~

    [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)

* [Pkg.go.dev](https://pkg.go.dev) - Paket keşfi ve dökümantasyonları için yeni adres Pkg.go.dev. [Rozet yaratma aracı](https://pkg.go.dev/badge) ile projenize bir rozet yaratabilirsiniz.

    [![PkgGoDev](https://pkg.go.dev/badge/github.com/golang-standards/project-layout)](https://pkg.go.dev/github.com/golang-standards/project-layout)

* Release - Projenizin en son yayınlanmış versiyonunu gösterir. İşaret ettiği linki kendi proje linkiniz ile değiştirin.

    [![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## Notlar

Örnek/yeniden kullanılabilir yapılandırmalar, komut dosyaları ve kodları içeren daha kararlı bir proje şablonu yapım aşamasındadır.
