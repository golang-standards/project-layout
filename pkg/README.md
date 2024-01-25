# `/pkg`

Library code that's ok to use by external applications (e.g., `/pkg/mypubliclib`). Other projects will import these libraries expecting them to work, so think twice before you put something here :-) Note that the `internal` directory is a better way to ensure your private packages are not importable because it's enforced by Go. The `/pkg` directory is still a good way to explicitly communicate that the code in that directory is safe for use by others. The [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) blog post by Travis Jeffery provides a good overview of the `pkg` and `internal` directories and when it might make sense to use them.

It's also a way to group Go code in one place when your root directory contains lots of non-Go components and directories making it easier to run various Go tools (as mentioned in these talks: [`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg) from GopherCon EU 2018, [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) and [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)).

Note that this is not a universally accepted pattern and for every popular repo that uses it you can find 10 that don't. It's up to you to decide if you want to use this pattern or not. Regardless of whether or not it's a good pattern more people will know what you mean than not. It might a bit confusing for some of the new Go devs, but it's a pretty simple confusion to resolve and that's one of the goals for this project layout repo.

Ok not to use it if your app project is really small and where an extra level of nesting doesn't add much value (unless you really want to). Think about it when it's getting big enough and your root directory gets pretty busy (especially if you have a lot of non-Go app components).

The `pkg` directory origins: The old Go source code used to use `pkg` for its packages and then various Go projects in the community started copying the pattern (see [`this`](https://twitter.com/bradfitz/status/1039512487538970624) Brad Fitzpatrick's tweet for more context).


Examples:

* https://github.com/containerd/containerd/tree/main/pkg
* https://github.com/slimtoolkit/slim/tree/master/pkg
* https://github.com/telepresenceio/telepresence/tree/release/v2/pkg
* https://github.com/jaegertracing/jaeger/tree/master/pkg
* https://github.com/istio/istio/tree/master/pkg
* https://github.com/GoogleContainerTools/kaniko/tree/master/pkg
* https://github.com/google/gvisor/tree/master/pkg
* https://github.com/google/syzkaller/tree/master/pkg
* https://github.com/perkeep/perkeep/tree/master/pkg
* https://github.com/heptio/ark/tree/master/pkg
* https://github.com/argoproj/argo-workflows/tree/master/pkg
* https://github.com/argoproj/argo-cd/tree/master/pkg
* https://github.com/heptio/sonobuoy/tree/master/pkg
* https://github.com/helm/helm/tree/master/pkg
* https://github.com/k3s-io/k3s/tree/master/pkg
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
* https://github.com/sosedoff/pgweb/tree/master/pkg
* https://github.com/GoogleContainerTools/skaffold/tree/master/pkg
* https://github.com/knative/serving/tree/master/pkg
* https://github.com/grafana/loki/tree/master/pkg
* https://github.com/bloomberg/goldpinger/tree/master/pkg
* https://github.com/Ne0nd0g/merlin/tree/master/pkg
* https://github.com/jenkins-x/jx/tree/master/pkg
* https://github.com/DataDog/datadog-agent/tree/master/pkg
* https://github.com/dapr/dapr/tree/master/pkg
* https://github.com/cortexproject/cortex/tree/master/pkg
* https://github.com/dexidp/dex/tree/master/pkg
* https://github.com/pusher/oauth2_proxy/tree/master/pkg
* https://github.com/pdfcpu/pdfcpu/tree/master/pkg
* https://github.com/weaveworks/kured/tree/master/pkg
* https://github.com/weaveworks/footloose/tree/master/pkg
* https://github.com/weaveworks/ignite/tree/master/pkg
* https://github.com/tmrts/boilr/tree/master/pkg
* https://github.com/kata-containers/runtime/tree/master/pkg
* https://github.com/okteto/okteto/tree/master/pkg
* https://github.com/solo-io/squash/tree/master/pkg
* https://github.com/google/exposure-notifications-server/tree/main/pkg
* https://github.com/spiffe/spire/tree/main/pkg
* https://github.com/rook/rook/tree/master/pkg
* https://github.com/buildpacks/pack/tree/main/pkg
* https://github.com/cilium/cilium/tree/main/pkg
* https://github.com/containernetworking/cni/tree/main/pkg
* https://github.com/crossplane/crossplane/tree/master/pkg
* https://github.com/dragonflyoss/Dragonfly2/tree/main/pkg
* https://github.com/kubeedge/kubeedge/tree/master/pkg
* https://github.com/kubevela/kubevela/tree/master/pkg
* https://github.com/kubevirt/kubevirt/tree/main/pkg
* https://github.com/kyverno/kyverno/tree/main/pkg
* https://github.com/thanos-io/thanos/tree/main/pkg
* https://github.com/cri-o/cri-o/tree/main/pkg
* https://github.com/fluxcd/flux2/tree/main/pkg
* https://github.com/kedacore/keda/tree/main/pkg
* https://github.com/linkerd/linkerd2/tree/main/pkg
* https://github.com/opencost/opencost/tree/develop/pkg
* https://github.com/antrea-io/antrea/tree/main/pkg
* https://github.com/karmada-io/karmada/tree/master/pkg
* https://github.com/kuberhealthy/kuberhealthy/tree/master/pkg
* https://github.com/submariner-io/submariner/tree/devel/pkg
* https://github.com/trickstercache/trickster/tree/main/pkg
* https://github.com/tellerops/teller/tree/master/pkg
* https://github.com/OpenFunction/OpenFunction/tree/main/pkg
* https://github.com/external-secrets/external-secrets/tree/main/pkg
* https://github.com/ko-build/ko/tree/main/pkg
* https://github.com/lima-vm/lima/tree/master/pkg
* https://github.com/clastix/capsule/tree/master/pkg
* https://github.com/carvel-dev/ytt/tree/develop/pkg
* https://github.com/clusternet/clusternet/tree/main/pkg
* https://github.com/fluid-cloudnative/fluid/tree/master/pkg
* https://github.com/inspektor-gadget/inspektor-gadget/tree/main/pkg
* https://github.com/sustainable-computing-io/kepler/tree/main/pkg
* https://github.com/GoogleContainerTools/kpt/tree/main/pkg
* https://github.com/guacsec/guac/tree/main/pkg
* https://github.com/kubeovn/kube-ovn/tree/master/pkg
* https://github.com/kube-vip/kube-vip/tree/main/pkg
* https://github.com/kubescape/kubescape/tree/master/pkg
* https://github.com/kudobuilder/kudo/tree/main/pkg
* https://github.com/kumahq/kuma/tree/master/pkg
* https://github.com/kubereboot/kured/tree/main/pkg
* https://github.com/nocalhost/nocalhost/tree/main/pkg
* https://github.com/openelb/openelb/tree/master/pkg
* https://github.com/openfga/openfga/tree/main/pkg
* https://github.com/openyurtio/openyurt/tree/master/pkg
* https://github.com/getporter/porter/tree/main/pkg
* https://github.com/sealerio/sealer/tree/main/pkg
* https://github.com/werf/werf/tree/main/pkg
  
