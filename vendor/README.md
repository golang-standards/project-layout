# `/vendor`

Application dependencies (managed manually or by your favorite dependency management tool like the new built-in, but still experimental, [`modules`](https://github.com/golang/go/wiki/Modules) feature).

Don't commit your application dependencies if you are building a library.

Note that since [`1.13`](https://golang.org/doc/go1.13#modules) Go also enabled the module proxy feature (using `https://proxy.golang.org` as their module proxy server by default). Read more about it [`here`](https://blog.golang.org/module-mirror-launch) to see if it fits all of your requirements and constraints. If it does, then you won't need the 'vendor' directory at all.
