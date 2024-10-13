
## نمای کلی 

این متن یک طرح اساسی برای پروژه‌های برنامه‌نویسی به زبان Go است. توجه داشته باشید که این طرح از نظر محتوا ساده است، زیرا فقط روی layout کلی تمرکز دارد و نه آنچه در داخل آن قرار دارد. همچنین ساده است هر چند که از نظر محتوی پیشرفته است ولی به جزئیات زیادی در مورد نحوه ساختاردهی بیشتر پروژه شما نمی‌پردازد. برای مثال، سعی نمی‌کند ساختار پروژه‌ای را  با چیزی مانند Clean Architecture  پوشش دهد.

این یک `استاندارد رسمی تعریف شده توسط تیم توسعه اصلی Go نیست`. این مجموعه‌ای از الگوهای layout پروژه‌های رایج با سابقه و نوظهور در اکوسیستم Go است. برخی از این الگوها از بقیه محبوب‌تر هستند. همچنین دارای تعدادی بهبود کوچک همراه با چندین دایرکتوری پشتیبانی مشترک در هر برنامه واقعی به اندازه کافی بزرگ است. توجه داشته باشید که تیم اصلی Go مجموعه خوبی از دستورالعمل‌های عمومی در مورد ساختاردهی پروژه‌های Go و معنای آن برای پروژه شما هنگام وارد شدن و نصب آن ارائه می‌دهد. برای اطلاعات بیشتر، به صفحه [`Organizing a Go module`](https://go.dev/doc/modules/layout) در اسناد رسمی Go مراجعه کنید. این صفحه شامل الگوهای دایرکتوری داخلی و cmd (که در زیر توضیح داده شده است) و سایر اطلاعات مفید است.

اگر در حال یادگیری Go هستید یا یک PoC یا یک پروژه ساده برای خودتان می‌سازید، این طرح پروژه بیش از حد پیچیده است. در عوض، با چیزی بسیار ساده شروع کنید (یک فایل main.go و go.mod بیش از اندازه کافی است). با پیشرفت پروژه، به خاطر داشته باشید که ساختاردهی مناسب کد شما بسیار مهم خواهد بود، در غیر این صورت در نهایت با یک کد نامرتب با وابستگی‌های پنهان (hidden dependencies) و  global state زیادی مواجه خواهید شد. وقتی افراد بیشتری روی پروژه کار می‌کنند، به ساختار بیشتری نیاز خواهید داشت. در این زمان است که معرفی یک روش مشترک برای مدیریت بسته‌ها/کتابخانه‌ها اهمیت دارد. وقتی یک پروژه متن‌باز دارید یا می‌دانید که پروژه‌های دیگر کد را از مخزن پروژه شما وارد می‌کنند، آن زمان است که داشتن بسته‌ها و کدهای خصوصی (معادل internal) اهمیت دارد. repository را کپی کنید، آنچه را که نیاز دارید نگه دارید و بقیه را حذف کنید! فقط به این دلیل که وجود دارد، به این معنی نیست که باید از همه آن استفاده کنید. هیچ یک از این الگوها در هر پروژه‌ای استفاده نمی‌شوند. حتی الگوی  `vendor` نیز universal نیست.

با آمدن Go 1.14 در نهایت [`Go Modules`](https://go.dev/wiki/Modules)  برای production آماده شدند. از [`Go Modules`](https://blog.golang.org/using-go-modules) استفاده کنید، مگر اینکه دلیل خاصی برای استفاده نکردن از آن‌ها داشته باشید و اگر اینطور است، نیازی به نگرانی در مورد $GOPATH و جایی که پروژه خود را قرار می‌دهید ندارید. فایل `go.mod` پایه در مخزن فرض می‌کند که پروژه شما در GitHub میزبانی می‌شود، اما این یک الزام نیست. module path می‌تواند هر چیزی باشد، اگرچه اولین جزء module path باید یک نقطه در نام خود داشته باشد (نسخه فعلی Go دیگر آن را اجباری نمی‌کند، اما اگر از نسخه‌های کمی قدیمی‌تر استفاده می‌کنید، تعجب نکنید اگر ساخت‌های شما بدون آن شکست بخورد). اگر می‌خواهید درباره آن بیشتر بدانید، به Issues  [`37554`](https://github.com/golang/go/issues/37554) و [`32819`](https://github.com/golang/go/issues/32819)مراجعه کنید.

این طرح پروژه عمداً عمومی است و سعی نمی‌کند یک ساختار بسته Go خاص را تحمیل کند.

این یک تلاش مشترک است. اگر الگوی جدیدی مشاهده کردید یا فکر می‌کنید یکی از الگوهای موجود نیاز به به‌روزرسانی دارد، یک issue را باز کنید.

اگر به کمک در مورد نام‌گذاری، قالب‌بندی و style نیاز دارید، ابتدا [`gofmt`](https://golang.org/cmd/gofmt/) و  [`staticcheck`](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck) را اجرا کنید. linter استاندارد قبلی، golint، اکنون منسوخ شده است و نگهداری نمی‌شود؛ استفاده از یک linter در حال توسعه و نگهداری شده مانند staticcheck توصیه می‌شود. همچنین مطمئن شوید که این دستورالعمل‌ها و توصیه‌های Go code style را بخوانید:

- [https://talks.golang.org/2014/names.slide](https://talks.golang.org/2014/names.slide)
- [https://golang.org/doc/effective_go.html#names](https://golang.org/doc/effective_go.html#names)
- [https://blog.golang.org/package-names](https://blog.golang.org/package-names)
- [https://go.dev/wiki/CodeReviewComments](https://go.dev/wiki/CodeReviewComments)
- [Style guideline for Go packages](https://rakyll.org/style-packages) (rakyll/JBD)

برای اطلاعات بیشتر ، [`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) را ببینید.  
اطلاعات بیشتر در مورد نامگذاری و سازماندهی بسته‌ها و همچنین سایر توصیه‌های ساختار کد:

- [GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)
- [GopherCon Russia 2018: Ashley McNamara + Brian Ketelsen - Go best practices.](https://www.youtube.com/watch?v=MzTcsI6tn-0)
- [GopherCon 2017: Edward Muller - Go Anti-Patterns](https://www.youtube.com/watch?v=ltqV6pDKZD8)
- [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)


## دایرکتوری‌ها در Go

#### /cmd

این دایرکتوری شامل برنامه‌های اصلی پروژه شما است. نام هر دایرکتوری فرعی باید با نام برنامه اجرایی مطابقت داشته باشد (برای مثال، `/cmd/myapp`).

از قرار دادن حجم کد زیاد در دایرکتوری برنامه خودداری کنید. اگر فکر می‌کنید این کد‌ها قابلیت وارد شدن و استفاده در پروژه‌های دیگر را دارد، باید در دایرکتوری `/pkg` قرار گیرد. اگر کد قابل استفاده مجدد نیست یا نمی‌خواهید دیگران از آن استفاده مجدد کنند، آن کد را در دایرکتوری  `/internal` قرار دهید. تعجب خواهید کرد که دیگران چه کارهایی انجام می‌دهند، بنابراین در مورد اهداف خود صریح باشید!

معمولاً یک تابع اصلی کوچک وجود دارد که کد را از دایرکتوری‌های `/internal `و `/pkg `وارد کرده و فراخوانی می‌کند و کار دیگری انجام نمی‌دهد.

به‌عنوان‌مثال به دایرکتوری [`/cmd`](https://github.com/golang-standards/project-layout/blob/master/cmd/README.md) مراجعه کنید.

#### /internal

شامل کد Private application و library code است. این کدی است که نمی‌خواهید دیگران آن را در برنامه‌ها یا کتابخانه‌های خود وارد کنند. توجه داشته باشید که این الگوی چیدمان توسط خود کامپایلر Go اعمال می‌شود. برای جزئیات بیشتر،  Go 1.4 [`release notes`](https://golang.org/doc/go1.4#internalpackages)را ببینید. توجه داشته باشید که شما به دایرکتوری  `internal` سطح بالا محدود نیستید. شما می‌توانید در هر سطحی از درخت پروژه خود بیش از یک دایرکتوری  `internal` داشته باشید.

به‌صورت اختیاری می‌توانید برای جدا کردن کد داخلی مشترک و غیرمشترک خود، کمی ساختار اضافی به بسته‌های داخلی (internal packages) خود اضافه کنید. این کار الزامی نیست (به ویژه برای پروژه‌های کوچک‌تر)، اما داشتن نشانه‌های بصری برای نشان دادن نحوه استفاده موردنظر package بسیار مناسب است. کد application واقعی شما می‌تواند در دایرکتوری  `/internal/app` (مثلاً  `/internal/app/myapp`) و کد مشترک بین آن برنامه‌ها در دایرکتوری `/internal/pkg/` (مثلاً , `/internal/pkg/myprivlib`) قرار گیرد.

شما از دایرکتوری‌های internal برای private کردن package‌ها استفاده می‌کنید. اگر یک package را داخل یک  internal directory قرار دهید، بسته‌های دیگر نمی‌توانند آن را وارد کنند مگر اینکه یک جد مشترک (common ancestor) داشته باشند. این تنها دایرکتوری‌ای است که در مستندات Go نام برده شده و نحوه برخورد با آن توسط کامپایلر Go خاص و متفاوت است.

#### /pkg

کد کتابخانه که امکان استفاده توسط برنامه‌های خارجی را دارد (به‌عنوان مثال،  `/pkg/mypubliclib`). سایر پروژه‌ها این کتابخانه‌ها را import می‌کنند و انتظار کارکرد درست آنها را دارند، بنابراین قبل از قرار دادن چیزی در اینجا خوب فکر کنید :-) و توجه داشته باشید که  `internal` directory، راه بهتری برای اطمینان از وارد نشدن private packages شماست، زیرا توسط Go اجرا می‌شود. دایرکتوری `/pkg`همچنان راه خوبی برای بیان صریح این است که کد موجود در آن دایرکتوری برای استفاده توسط دیگران ایمن است.

مقاله وبلاگ « [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/)» توسط Travis Jeffery، نمای کلی خوبی از دایرکتوری‌های pkg و internal و زمان‌هایی که استفاده از آنها منطقی است ارائه می‌دهد.

همچنین این راهی برای گروه‌بندی کد Go در یک مکان است، زمانی که دایرکتوری اصلی شما حاوی بسیاری از اجزا و دایرکتوری‌های غیر Go باشد، این کار اجرای ابزارهای مختلف Go را آسان‌تر می‌کند (همانطور که در این سخنرانی‌ها ذکر شده است: [`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg) از GopherCon EU و [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) و [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)).

اگر می‌خواهید ببینید کدام مخازن محبوب Go از این  layout pattern پروژه استفاده می‌کنند، به دایرکتوری  [`/pkg`](https://github.com/golang-standards/project-layout/blob/master/pkg/README.md) مراجعه کنید. این یک الگوی layout رایج است، اما به‌طور جهانی پذیرفته نشده است و برخی از اعضای جامعه Go آن را توصیه نمی‌کنند.

اگر پروژه برنامه‌ی شما واقعاً کوچک است و جایی که لایه‌ی اضافی تودرتو بودن ارزش زیادی اضافه نمی‌کند، استفاده نکردن از آن اشکالی ندارد (مگر اینکه واقعاً بخواهید :-)). در مورد آن فکر کنید زمانی که پروژه به اندازه کافی بزرگ می شود و دایرکتوری اصلی شما شلوغ می شود (به خصوص اگر اجزای برنامه غیر Go زیادی دارید).

**ریشه‌های دایرکتوری pkg:** کد منبع قدیمی Go برای بسته‌های خود از pkg استفاده می‌کرد و سپس پروژه‌های مختلف Go در جامعه شروع به کپی کردن این الگو کردند (برای درک بهتر به [این](https://twitter.com/bradfitz/status/1039512487538970624) توییت Brad Fitzpatrick مراجعه کنید).


### /vendor


وابستگی‌های برنامه (به‌صورت دستی یا توسط ابزار مدیریت وابستگی مورد علاقه شما مانند ویژگی جدید [`Go Modules`](https://go.dev/wiki/Modules)داخلی مدیریت می‌شود). دستور `go mod vendor` دایرکتوری `vendor/` را برای شما ایجاد می‌کند. توجه داشته باشید که اگر از Go 1.14 استفاده نمی‌کنید که به صورت پیش‌فرض فعال است، ممکن است نیاز به اضافه کردن پرچم `-mod=vendor` به دستور go build خود داشته باشید.

اگر در حال ساخت کتابخانه هستید، وابستگی‌های برنامه خود را commit نکنید.

توجه داشته باشید که از Go [`1.13`](https://golang.org/doc/go1.13#modules) ، قابلیت module proxy نیز در Go فعال شد (که به طور پیش‌فرض از [https://proxy.golang.org](https://proxy.golang.org/) به عنوان سرور پراکسی ماژول خود استفاده می‌کند). برای اینکه ببینید آیا این قابلیت با تمام الزامات و محدودیت‌های شما مطابقت دارد، در [اینجا](https://blog.golang.org/module-mirror-launch) بیشتر در مورد آن بخوانید. اگر اینطور باشد، اصلاً به دایرکتوری `vendor` نیاز نخواهید داشت.

## Service Application Directories

### `/api`
مشخصات OpenAPI/Swagger، فایل‌های طرحواره JSON، فایل‌های تعریف پروتکل.  
  
برای مثال به دایرکتوری  [`/api/`](https://github.com/golang-standards/project-layout/blob/master/api/README.md) مراجعه کنید.

## Web Application Directories

اجزای خاص برنامه وب: static web assets و templateهای سمت سرور و SPAها.

## Common Application Directories

### `/configs`
قالب‌های فایل پیکربندی یا تنظیمات پیش‌فرض.  
  
فایل‌های قالب  `confd` یا  `consul-template` خود را اینجا قرار دهید.
### `/init`
پیکربندی‌های init سیستم (systemd، upstart، sysv) و process manager/supervisor (runit, supervisord).

### `/scripts`
اسکریپت‌هایی برای انجام عملیات‌های مختلف build, install, analysis و غیره.
این اسکریپت‌ها Makefile سطح ریشه را کوچک و ساده نگه می‌دارند (به عنوان مثال،[`https://github.com/hashicorp/terraform/blob/main/Makefile`](https://github.com/hashicorp/terraform/blob/main/Makefile)).
برای مثال به دایرکتوری  [`scripts/`](https://github.com/golang-standards/project-layout/blob/master/scripts/README.md) مراجعه کنید.
### `/build`

‏Packaging و Continuous Integration

- پیکربندی‌ها و اسکریپت‌های package‌های ابری (AMI)، کانتینری (Docker)، سیستم‌عامل (deb، rpm، pkg) را در این دایرکتوری قرار دهید.

### `/build`


- پیکربندی‌ها و اسکریپت‌های CI (travis، circle، drone) را در این دایرکتوری قرار دهید. توجه داشته باشید که برخی از ابزارهای CI (مانند Travis CI) در مورد مکان فایل‌های پیکربندی خود بسیار حساس هستند. سعی کنید فایل‌های پیکربندی را در دایرکتوری  `/build/ci` قرار داده و آن‌ها را به مکانی که ابزارهای CI انتظار دارند (در صورت امکان) لینک کنید.

### `/deployments`


- پیکربندی‌ها و قالب‌های deployment یا استقرار IaaS، PaaS، سیستم و orchestration کانتینر  (dockerCompose, kubernetes/helm, terraform). توجه داشته باشید که در برخی از repoها (به ویژه برنامه‌هایی که با kubernetes استقرار می‌یابند) این دایرکتوری  `deploy/` نامیده می‌شود.

### `/test`

- **برنامه‌های تست خارجی اضافی و داده‌های تست**. می‌توانید دایرکتوری  `test/` را به هر شکلی که می‌خواهید ساختار دهید. برای پروژه‌های بزرگ‌تر، داشتن یک زیردایرکتوری data منطقی است. برای مثال، می‌توانید  `test/testdata/` یا  `test/data/` را داشته باشید اگر نیاز دارید که Go آنچه در آن دایرکتوری است را نادیده بگیرد. توجه داشته باشید که Go همچنین دایرکتوری‌ها یا فایل‌هایی که با "." یا "" شروع می‌شوند را نادیده می‌گیرد، بنابراین در نحوه نام‌گذاری دایرکتوری داده‌های تست خود انعطاف بیشتری دارید.

برای نمونه‌ها به دایرکتوری  [`test/`](https://github.com/golang-standards/project-layout/blob/master/test/README.md)مراجعه کنید.
## دایرکتوری‌های دیگر

اسناد طراحی و کاربر (علاوه بر مستندات ایجاد شده توسط godoc شما).  
  
برای مثال به دایرکتوری  [`docs/`](https://github.com/golang-standards/project-layout/blob/master/docs/README.md) مراجعه کنید.  
  
### `/tools`

ابزارهای پشتیبانی این پروژه توجه داشته باشید که این ابزارها می توانند کد را از دایرکتوری های pkg/ و internal/ وارد کنند.  
  
برای مثال به دایرکتوری tools/ مراجعه کنید.

### `/examples`
نمونه‌هایی برای application و یا کتابخانه‌های public شما.  
  
برای مثال به دایرکتوری  [`examples/`](https://github.com/golang-standards/project-layout/blob/master/examples/README.md) مراجعه کنید.

### `/third_party`

ابزارهای کمکی خارجی، کد fork شده و سایر ابزارهای شخص ثالث (مانند Swagger UI).

### `/githooks`

Git hooks.

### `/assets`
سایر assetها برای همراهی با repository شما (image, logoها و غیره).

### `/website`
اگر از GitHub pages استفاده نمی‌کنید، اینجا مکانی است که می توانید داده‌های وب‌سایت پروژه خود را قرار دهید.  
  
برای مثال به دایرکتوری [`website/`](https://github.com/golang-standards/project-layout/blob/master/website/README.md)مراجعه کنید.

## دایرکتوری‌هایی که نباید داشته باشید

### /src

برخی از پروژه‌های Go دارای یک پوشه  `src` هستند، اما این معمولاً زمانی اتفاق می‌افتد که توسعه‌دهندگان از دنیای جاوا آمده‌اند که در آنجا یک الگوی رایج است. اگر می‌توانید، سعی کنید این الگوی جاوا را نپذیرید. شما واقعاً نمی‌خواهید که کد Go یا پروژه‌های Go شما شبیه جاوا به نظر برسند :-)

دایرکتوری  `/src` در سطح پروژه را با دایرکتوری  `/src` که Go برای کارگاه‌های خود استفاده می‌کند، اشتباه نگیرید که در [`How to Write Go Code`](https://golang.org/doc/code.html) توضیح داده شده است.  `$GOPATH` environment variable به  (current) workspace فعلی شما اشاره می‌کند (به طور پیش‌فرض به `$HOME/go` در سیستم‌های غیر ویندوزی اشاره می‌کند). این workspace شامل دایرکتوری‌های سطح بالا  `/pkg`, `/bin` و  `/src` است. پروژه واقعی شما در نهایت یک زیردایرکتوری زیر  `/src` می‌شود، بنابراین اگر دایرکتوری  `/src` را در پروژه خود دارید، مسیر پروژه به این شکل خواهد بود:  `/some/path/to/workspace/src/your_project/src/your_code.go`. توجه داشته باشید که با  Go 1.11 امکان دارد پروژه خود را خارج از  `GOPATH` خود داشته باشید، اما این هنوز به معنای این نیست که استفاده از این الگوی layout pattern ایده خوبی است.

## Badges


‏ [Go Report Card](https://goreportcard.com/)
کد شما را با gofmt، go vet، gocyclo، golint، ineffassign، مجوز و غلط املایی اسکن می کند. مرجع پروژه خود را جایگزین github.com/golang-standards/project-layout کنید.

[![Go Report Card](https://camo.githubusercontent.com/c1e1c210dea2e0410ecbb861999b969841a526d424818d7f9b816bd9f1364d55/68747470733a2f2f676f7265706f7274636172642e636f6d2f62616467652f6769746875622e636f6d2f676f6c616e672d7374616e64617264732f70726f6a6563742d6c61796f75743f7374796c653d666c61742d737175617265)](https://goreportcard.com/report/github.com/golang-standards/project-layout)

‏ ~~[GoDoc](http://godoc.org/)~~
~~این نسخه آنلاین اسناد تولید شده GoDoc شما را ارائه می‌دهد. link را تغییر دهید تا به پروژه شما اشاره کند.~~

[![Go Doc](https://camo.githubusercontent.com/fe1188b9f0668a1e0a543e1cbcc6fb28d50a52f74d04e99407f8e6405a7132cd/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f676f646f632d7265666572656e63652d626c75652e7376673f7374796c653d666c61742d737175617265)](http://godoc.org/github.com/golang-standards/project-layout)

‏ [Pkg.go.dev](https://pkg.go.dev/)
‏ Pkg.go.dev مقصد جدیدی برای شناسایی و مستندات Go است. می‌توانید با استفاده از [badge generation tool](https://pkg.go.dev/badge)آن را ایجاد کنید.

[![PkgGoDev](https://camo.githubusercontent.com/2f374e52d47edc4ea0e93661bc0eb204743e26718e423a31ae87bac272385081/68747470733a2f2f706b672e676f2e6465762f62616467652f6769746875622e636f6d2f676f6c616e672d7374616e64617264732f70726f6a6563742d6c61796f7574)](https://pkg.go.dev/github.com/golang-standards/project-layout)

در مورد Release - آخرین شماره انتشار پروژه شما را نشان می دهد. لینک github را تغییر دهید تا به پروژه شما اشاره کند.


## نکته‌ها

یک الگوی پروژه با نظر بیشتر با تنظیمات sample/reusable استفاده مجدد، اسکریپت‌ها و کد یک WIP است.