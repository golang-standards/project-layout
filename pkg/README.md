# `/pkg`

Library code that's ok to use by external applications (e.g., `/pkg/mypubliclib`). Other projects will import these libraries expecting them to work, so think twice before you put something here :-)

It's also a way to group Go code in one place when your root directory contains lots of non-Go components and directories making it easier to run various Go tool (as mentioned in the [`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg) from GopherCon EU 2018).

Note that this is not a universally accepted pattern and for every popular repo that uses it you can find 10 that don't. It's up to you to decide if you want to use this pattern or not. Regardless of whether or not it's a good pattern more people will know what you mean than not. It is a bit confusing for new Go devs, but it's a pretty simple confusion to resolve and that's one of the goals for this project layout repo.

Examples:

* https://github.com/gogs/gogs/tree/master/pkg
* https://github.com/prometheus/prometheus/tree/master/pkg
* https://github.com/jaegertracing/jaeger/tree/master/pkg
* https://github.com/istio/istio/tree/master/pkg
* https://github.com/google/gvisor/tree/master/pkg
* https://github.com/google/syzkaller/tree/master/pkg
* https://github.com/perkeep/perkeep/tree/master/pkg
* https://github.com/minio/minio/tree/master/pkg
* https://github.com/heptio/ark/tree/master/pkg
* https://github.com/argoproj/argo/tree/master/pkg
* https://github.com/heptio/sonobuoy/tree/master/pkg
* https://github.com/helm/helm/tree/master/pkg
* https://github.com/kubernetes/kubernetes/tree/master/pkg
* https://github.com/kubernetes/kops/tree/master/pkg
* https://github.com/moby/moby/tree/master/pkg
* https://github.com/grafana/grafana/tree/master/pkg
* https://github.com/influxdata/influxdb/tree/master/pkg
* https://github.com/cockroachdb/cockroach/tree/master/pkg
* https://github.com/derekparker/delve/tree/master/pkg
* https://github.com/etcd-io/etcd/tree/master/pkg
* https://github.com/oklog/oklog/tree/master/pkg
* https://github.com/flynn/flynn/tree/master/pkg
* https://github.com/jesseduffield/lazygit/tree/master/pkg
* https://github.com/gopasspw/gopass/tree/master/pkg
* https://github.com/sourcegraph/sourcegraph/tree/master/pkg
* https://github.com/sosedoff/pgweb/tree/master/pkg
* https://github.com/GoogleContainerTools/skaffold/tree/master/pkg
* https://github.com/knative/serving/tree/master/pkg
* https://github.com/grafana/loki/tree/master/pkg
* https://github.com/bloomberg/goldpinger/tree/master/pkg
* https://github.com/crossplaneio/crossplane/tree/master/pkg
* https://github.com/Ne0nd0g/merlin/tree/master/pkg
* https://github.com/jenkins-x/jx/tree/master/pkg
