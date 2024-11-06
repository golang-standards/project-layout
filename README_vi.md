# Bố cục tiêu chuẩn của một dự án Go

Các bản dịch:

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
* [Vietnamese](README_vi.md)
* [हिन्दी](README_hi.md)
* [Беларуская](README_by.md)

## Tổng quan

Đây là một bố cục cơ bản cho nhiều ứng dụng Go. Dù **`không phải là tiêu chuẩn chính thức được đưa ra từ đội ngũ cốt lõi của Go`**, đây là một tập hợp mẫu bố cục trong hệ sinh thái Go. Những cải tiến nhỏ cùng với một số thư mục hỗ trợ đều có thể được dùng phổ biến cho bất kỳ ứng dụng lớn nào trong thực tế.

**`Nếu bạn đang muốn học Go hoặc đang xây dựng một PoC hoặc một dự án cá nhân đơn giản thì bố cục này là thừa. Bạn nên bắt đầu bằng những thứ đơn giản hơn (chẳng hạn một tập tin `main.go`và một tập tin`go.mod` là đủ).`** Hãy nhớ một điều là khi dự án của bạn phát triển, điều quan trọng là dự án của bạn có cấu trúc tốt, nếu không thì bạn sẽ gặp nhiều vấn đề với mã nguồn lộn xộn cùng các phụ thuộc ẩn và trạng thái toàn cục. Khi bạn có nhiều người làm việc trên một dự án, bạn còn cần cấu trúc nhiều hơn. Vì thế, việc quan trọng là giới thiệu cách phổ biến để quản lý các gói/thư viện. Khi bạn có một dự án mã nguồn mở hoặc khi bạn biết các dự án khác, hãy nhập mã từ kho lưu trữ dự án của bạn, đó là lúc điều quan trọng là phải có các gói và mã riêng tư (hay còn gọi là `nội bộ`). Sao y kho lưu trữ, giữ lại những thứ bạn cần và xóa hết phần còn lại! Chỉ vì nó ở đó không có nghĩa là bạn phải sử dụng tất cả. Không có mẫu nào trong số này được sử dụng trong mọi dự án. Ngay cả mô hình `vendor` cũng không phải là phổ biến.

Ở phiên bản Go 1.14, [`Go Modules`](https://go.dev/wiki/Modules) đã sẵn sàng để dùng trên môi trường production. Trừ khi bạn có một lý do cục thể nào đấy, còn không thì bạn nên dùng [`Go Modules`](https://blog.golang.org/using-go-modules) vì bạn sẽ không cần để tâm tới $GOPATH và nơi bạn đặt dự án của mình. Tập tin cơ bản `go.mod` trong repo giả định rằng dự án của bạn được lưu trữ trên GitHub, nhưng nó không phải là một yêu cầu. Đường dẫn module có thể là bất kỳ thứ gì mặc dù phần đầu thành phần đường dẫn module phải có dấu chấm trong tên của nó (phiên bản Go hiện tại không bắt buộc điều này, nhưng nếu bạn đang sử dụng các phiên bản cũ hơn, đừng ngạc nhiên nếu bản dựng của bạn thất bại khi thiếu nó). Xem các lỗi [`37554`](https://github.com/golang/go/issues/37554) và [`32819`](https://github.com/golang/go/issues/32819) nếu bạn muốn tìm hiểu thêm.

Bố cục dự án này có chủ đích chung chung và nó không cố gắng áp đặt một vài trúc gói Go cụ thể.

Dự án này là nỗ lực của cộng đồng. Hãy mở một issue nếu bạn gặp một mẫu thiết kế nào mới hoặc bạn nghĩ những mẫu thiết kế có sẵn cần được cập nhật.

Bắt đầu với [`gofmt`](https://golang.org/cmd/gofmt/) và [`golint`](https://github.com/golang/lint) nếu bạn cần hỗ trợ về cách đặt tên, định dạng và phong cách. Đồng thời, đảm bảo rằng bạn đã đọc các hướng dẫn và khuyến nghị của Go dưới đây:

* https://talks.golang.org/2014/names.slide
* https://golang.org/doc/effective_go.html#names
* https://blog.golang.org/package-names
* https://go.dev/wiki/CodeReviewComments
* [Hướng dẫn thiết kế các gói trong Go](https://rakyll.org/style-packages) (rakyll/JBD)

Đọc thêm [`Mẫu dự án Go`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) để biết thêm thông tin cơ bản.

Tìm hiểu thêm về cách đặt tên và tổ chức các gói cũng như các đề xuất về cấu trúc mã khác:

* [GopherCon EU 2018: Peter Bourgon - Thực tiễn tốt nhất cho lập trình công nghiệp](https://www.youtube.com/watch?v=PTE4VJIdHPg)
* [GopherCon Russia 2018: Ashley McNamara + Brian Ketelsen - Thực tiễn tốt trong Go.](https://www.youtube.com/watch?v=MzTcsI6tn-0)
* [GopherCon 2017: Edward Muller - Các chống mẫu trong Go](https://www.youtube.com/watch?v=ltqV6pDKZD8)
* [GopherCon 2018: Kat Zien - Bạn cấu trúc ứng dụng Go của mình như thế nào](https://www.youtube.com/watch?v=oL6JBUk6tj0)

Một bài đăng của Trung Quốc về hướng dẫn thiết kế theo hướng gói và lớp kiến ​​trúc

* [Thiết kế theo hướng gói và phân lớp kiến ​​trúc](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md)

## Các thư mục trong Go

### `/cmd`

Thư mục chứa các ứng dụng chính cho dự án này.

Tên thư mục cho mỗi ứng dụng phải khớp với tên của tập tin thực thi mà bạn muốn có (ví dụ: `/cmd/myapp`).

Đừng đặt nhiều mã trong thư mục ứng dụng. Nếu bạn nghĩ rằng mã có thể được nhập và sử dụng trong các dự án khác, thì nó sẽ nằm trong thư mục `/pkg`. Nếu mã không thể sử dụng lại được hoặc nếu bạn không muốn người khác sử dụng lại, hãy đặt mã đó vào thư mục `/internal`. Bạn sẽ ngạc nhiên về những gì người khác sẽ làm, vì vậy hãy rõ ràng về ý định của bạn!

Thông thường có một hàm `main` nhỏ nhập và gọi mã từ các thư mục `/internal` và `/pkg` và không có gì khác.

Xem thêm ví dụ ở thư mục [`/cmd`](cmd/README.md).

### `/internal`

Thư mục chứa ứng dụng riêng và mã thư viện. Đây là mã mà bạn không muốn người khác sử dụng trong các ứng dụng hoặc thư viện của họ. Lưu ý, mẫu bố cục này được thực thi bởi chính trình biên dịch Go. Xem Go 1.4 [`ghi chú phát hành`](https://golang.org/doc/go1.4#internalpackages) để biết thêm chi tiết. Lưu ý rằng bạn không bị giới hạn ở thư mục `internal` cấp cao nhất. Bạn có thể có nhiều hơn một thư mục `internal` ở bất kỳ cấp nào trong cây dự án của bạn.

Bạn có thể tùy chọn thêm một chút cấu trúc bổ sung vào các gói bên trong của mình để tách mã nội bộ được chia sẻ và không được chia sẻ. Nó không bắt buộc (đặc biệt đối với các dự án nhỏ), nhưng thật tuyệt khi có manh mối trực quan cho thấy mục đích sử dụng gói dự kiến. Mã ứng dụng thực tế của bạn có thể nằm trong thư mục `/internal/app` (ví dụ: `/internal/app/myapp`) và mã được các ứng dụng đó chia sẻ trong thư mục `/internal/pkg` (ví dụ: `/internal/pkg/myprivlib`).

### `/pkg`

Thư mục chứa code thư viện cho phép các ứng dụng bên ngoài sử dụng (ví dụ: `/pkg/mypubliclib`). Các dự án khác sẽ nhập các thư viện này và kỳ vọng là chúng sẽ hoạt động, vì vậy hãy nghĩ cẩn thận trước khi bạn để thứ gì vào đây :-). Lưu ý rằng thư mục `nội bộ` là cách tốt hơn để đảm bảo các gói riêng tư của bạn không thể nhập được vì nó được Go thực thi. Thư mục `/pkg` vẫn là một cách tốt để thông báo rõ ràng rằng mã trong thư mục đó an toàn cho người khác sử dụng. Bài [`Tôi sẽ dùng pkg thay vì gói nội bộ`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) đăng bởi Travis Jeffery cung cấp một góc nhìn tổng quan về các thư mục `pkg` và `internal` và khi nào nên dùng chúng.

Đó cũng là một cách để nhóm mã Go vào một nơi khi thư mục gốc của bạn chứa nhiều thành phần và thư mục không phải Go, giúp chạy các công cụ Go khác nhau dễ dàng hơn (như đã đề cập trong các bài nói này: [`Thực tiễn tốt nhất cho lập trình công nghiệp`](https://www.youtube.com/watch?v=PTE4VJIdHPg) từ GopherCon EU 2018, [GopherCon 2018: Kat Zien - Làm thế nào để tổ chức các ứng dụng Go](https://www.youtube.com/watch?v=oL6JBUk6tj0) và [GoLab 2018 - Massimiliano Pippi - Mẫu bố cục dự án trong Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)).

Xem thư mục [`/pkg`](pkg/README.md) nếu bạn muốn xem repos Go phổ biến nào dùng bố cục này. Đây là một mẫu bố cục phổ biến, nhưng nó không được chấp nhận rộng rãi và một số người trong cộng đồng Go không khuyến khích nó.

Bạn không nên sử dụng nó nếu dự án ứng dụng của bạn thực sự nhỏ và ở đó mức độ lồng ghép bổ sung không mang lại nhiều giá trị (trừ khi bạn thực sự muốn :-)). Hãy nghĩ về nó khi nó đủ lớn và thư mục gốc của bạn trở nên khá bận rộn (đặc biệt nếu bạn có nhiều thành phần ứng dụng không phải của Go).

Nguồn gốc thư mục `pkg`: Mã nguồn Go cũ dùng `pkg` cho các gói của nó và sau đó các dự án Go khác trong cộng đồng bắt đầu sao chép mẫu này (xem [`Tweet`](https://twitter.com/bradfitz/status/1039512487538970624) của Brad Fitzpatrick để biết thêm ngữ cảnh).

### `/vendor`

Thư mục chứa các phụ thuộc của ứng dụng (được quản lý thủ công hoặc bằng công cụ quản lý phụ thuộc ưa thích của bạn tương tự như tính năng tích hợp mới là [`Go Modules`](https://go.dev/wiki/Modules)). Câu lệnh `go mod vendor` sẽ tạo ra cho bạn một thư mục `/vendor`. Lưu ý rằng bạn có thể sẽ cần phải thêm cờ hiệu `-mod=vendor` cho câu lệnh `go build` nếu bạn không dùng Go 1.14, phiên bản được thêm cờ hiệu mặc định.

Không nên commit các phụ thuộc ứng dụng nếu bạn đang muốn xây dựng một thư viện.

Lưu ý rằng kể từ phiên bản [`1.13`](https://golang.org/doc/go1.13#modules), Go bật tính năng module proxy (mặc định dùng máy chủ module proxy [`https://proxy.golang.org`](https://proxy.golang.org)). Đọc thêm để xem liệu nó có phù hợp với tất cả các yêu cầu và ràng buộc của bạn hay không ở [`đây`](https://blog.golang.org/module-mirror-launch). Nếu có thì bạn không cần tới thư mục `vendor`.

## Thư mục ứng dụng dịch vụ

### `/api`

Thư mục chứa bản mô tả OpenAPI/Swagger, tập tin lược đồ JSON, tập tin định nghĩa giao thức.

Xem thêm ví dụ ở thư mục [`/api`](api/README.md).

## Thư mục ứng dụng Web

### `/web`

Thư mục chứa các thành phần cụ thể của ứng dụng web: tài nguyên web tĩnh, mẫu bên máy chủ và SPAs.

Để các tập mẫu `confd` và `consul-template` ở đây.

### `/init`

Thư mục chứa phần khởi tạo hệ thống (systemd, upstart, sysv) và cấu hình quản lý/giám sát tiến trình (runit, supervisord).

### `/scripts`

Thư mục chứa tập lệnh để thực hiện các hoạt động xây dựng, cài đặt, phân tích ...

Các tập lệnh này làm cho tập Makefile ở cấp cao nhất nhỏ gọn và đơn giản. (Ví dụ: [`https://github.com/hashicorp/terraform/blob/master/Makefile`](https://github.com/hashicorp/terraform/blob/master/Makefile))

Xem ví dụ ở thư mục [`/scripts`](scripts/README.md).

### `/build`

Thư mục chứa gói và tích hợp liên tục.

Đặt các cấu hình và tập lệnh các gói đám mây (AMI), container (Docker), OS (deb, rpm, pkg) của bạn vào thư mục `/build/package`.

Đặt cấu hình và tập lệnh CI (travis, circle, drone) trong thư mục `/build/ci`. Lưu ý rằng một vài công cụ CI (ví dụ: Travis CI) rất kén chọn vị trí của tập tin cấu hình. Thử đặt các tập tin cấu hình ở thư mục `/build/ci`, lên kết chúng với vị trí mà công cụ CI mong đợi (khi có thể).

### `/deployments`

Thư mục chứa IaaS, PaaS, các cấu hình và mẫu triển khai điều phối hệ thống và vùng chứa (docker-compose, kubernetes/helm, mesos, terraform, bosh). Lưu ý rằng trong một số repo (đặc biệt là các ứng dụng được triển khai với kubernetes) thư mục này được gọi là `/deploy`.

### `/test`

Thư mục chứa các ứng dụng thử nghiệm bên ngoài bổ sung và dữ liệu thử nghiệm. Hãy thoải mái cấu trúc thư mục `/test` theo cách bạn muốn. Đối với các dự án lớn hơn, điều hợp lý là có một thư mục con dữ liệu. Ví dụ: bạn có thể có `/test/data` hoặc `/test/testdata` nếu bạn cần Go bỏ qua những gì trong thư mục đó. Lưu ý rằng Go cũng sẽ bỏ qua các thư mục hoặc tệp bắt đầu bằng "." hoặc "_", vì vậy bạn có thể linh hoạt hơn về cách đặt tên cho thư mục dữ liệu thử nghiệm của mình.

Xem ví dụ ở thư mục [`/test`](test/README.md).

## Thư mục khác

### `/docs`

Thư mục chứa tài liệu người dùng và bản thiết kế (bên cạnh tài liệu do godoc tạo ra).

Xem ví dụ ở thư mục [`/docs`](docs/README.md).

### `/tools`

Thư mục chứa công cụ hỗ trợ cho dự án này. Lưu ý rằng các công cụ này có thể nhập mã từ thư mục `/pkg` và `/internal`.

Xem ví dụ ở thư mục [`/tools`](tools/README.md).

### `/examples`

Thư mục chứa mẫu cho ứng dụng và/hoặc các thư viện công cộng của bạn.

Xem ví dụ ở thư mục [`/examples`](examples/README.md)

### `/third_party`

Thư mục chứa các công cụ trợ giúp bên ngoài, mã phân nhánh và các tiện ích bên thứ ba khác (ví dụ: giao diện người dùng Swagger).

### `/githooks`

Thư mục chứa git hooks.

### `/assets`

Các tài sản khác đi cùng với kho lưu trữ của bạn (hình ảnh, logo ...).

### `/website`

Đây là nơi để dữ liệu trang web của bạn nếu bạn không sử dụng các trang của GitHub.

Xem ví dụ ở thư mục [`/website`](website/README.md).

## Thư mục bạn không nên có

### `/src`

Một vài dự án Go có thư mục `src` nhưng nó thường xảy ra khi các nhà phát triển đến từ thế giới Java, nơi đó là một mẫu phổ biến. Bạn không nên học mẫu này từ Java. Bạn thực sự không muốn mã Go hoặc các dự án Go của mình trông giống như Java :-)

Đừng nhầm lẫn giữa thư mục `/src` cấp dự án với thư mục `/src` mà Go sử dụng cho các không gian làm việc của nó như được mô tả trong [`Cách viết mã Go`](https://golang.org/doc/code.html). Biến môi trường `$GOPATH` trỏ tới không gian làm việc hiện tại của bạn (với những hệ thống không phải là Windows, nó mặc định trỏ tới `$HOME/go`). Không gian làm việc này bao gồm các thư mục `/pkg`, `/bin`, `/src` cấp cao nhất. Dự án thực tế của bạn kết thúc là một thư mục con trong `/src`, vì vậy nếu bạn có thư mục `/src` trong dự án của mình, đường dẫn dự án sẽ giống như sau: `/some/path/to/workspace/src/your_project/src/your_code.go`. Lưu ý rằng với Go 1.11, bạn có thể đặt dự án của mình bên ngoài `GOPATH`, nhưng điều đó không có nghĩa là bạn nên sử dụng mẫu bố cục này.

## Danh hiệu

* [Go Report Card](https://goreportcard.com/) - Nó sẽ quét toàn bộ mã của bạn bằng `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` and `misspell`. Thay `github.com/golang-standards/project-layout` bằng tuỳ chọn trong dự án của bạn.

  [![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)

* ~~[GoDoc](http://godoc.org) - Nó sẽ cung cấp phiên bản trực tuyến của tài liệu do GoDoc tự tạo. Đổi đường dẫn để trỏ tới dự án của bạn.~~

  [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)

* [Pkg.go.dev](https://pkg.go.dev) - Pkg.go.dev là điểm đến mới cho khám phá và tài liệu về Go. Bạn có thể tạo huy hiệu bằng [badge generation tool](https://pkg.go.dev/badge).

  [![PkgGoDev](https://pkg.go.dev/badge/github.com/golang-standards/project-layout)](https://pkg.go.dev/github.com/golang-standards/project-layout)

* Bản phát hành - Nó sẽ hiển thị số phát hành mới nhất cho dự án của bạn. Thay đổi liên kết github để trỏ đến dự án của bạn.

  [![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## Ghi chú

Một mẫu dự án "có định kiến" (opinionated) hơn đang được xây dựng với các cấu hình, tập lệnh và mã có thể tái sử dụng.
