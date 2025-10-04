# Exercise: Play around kind and kubectl

The goal this time is to play around and learn the basics.

## 1. Inspect the nodes

Which nodes do kubectl report? which containers is docker running?

### kubectl nodes

```
$ kubectl get nodes -o wide
NAME                             STATUS   ROLES           AGE   VERSION   INTERNAL-IP   EXTERNAL-IP   OS-IMAGE                         KERNEL-VERSION           CONTAINER-RUNTIME
kubedredger-kind-control-plane   Ready    control-plane   67m   v1.34.0   172.18.0.2    <none>        Debian GNU/Linux 12 (bookworm)   6.16.8-200.fc42.x86_64   containerd://2.1.3

```

### docker "nodes"

```
$ docker ps
CONTAINER ID   IMAGE                  COMMAND                  CREATED             STATUS             PORTS                       NAMES
8c160df29ccf   kindest/node:v1.34.0   "/usr/local/bin/entr…"   About an hour ago   Up About an hour   127.0.0.1:44987->6443/tcp   kubedredger-kind-control-plane
```

## 2. Peek at processes

Which process can you see from your host environment? What about the "nodes"?

### processes in the "node"

```
$ docker exec $CONTAINER ps -fauxw
USER         PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root        2024  0.0  0.0   8072  4204 ?        Rs   13:11   0:00 ps -fauxw
root           1  0.0  0.0  21172 12608 ?        Ss   12:02   0:00 /sbin/init
root          97  0.0  0.0  24832 10864 ?        Ss   12:02   0:00 /lib/systemd/systemd-journald
root         112  0.2  0.4 3490888 65612 ?       Ssl  12:02   0:08 /usr/local/bin/containerd
root         306  0.0  0.0 1233812 11232 ?       Sl   12:02   0:00 /usr/local/bin/containerd-shim-runc-v2 -namespace k8s.io -id 3a4b8d1bb1c4dbbfe441a6529c614fdee049afc1529a96f4844bbe5bc496e837 -address /run/containerd/containerd.sock
65535        420  0.0  0.0   1028   752 ?        Ss   12:02   0:00  \_ /pause
root         552  0.1  0.3 1278776 62096 ?       Ssl  12:02   0:06  \_ kube-scheduler --authentication-kubeconfig=/etc/kubernetes/scheduler.conf --authorization-kubeconfig=/etc/kubernetes/scheduler.conf --bind-address=127.0.0.1 --kubeconfig=/etc/kubernetes/scheduler.conf --leader-elect=true
root         332  0.0  0.0 1233812 11208 ?       Sl   12:02   0:00 /usr/local/bin/containerd-shim-runc-v2 -namespace k8s.io -id f3854c21e4e3b42cd8187d3e7ab30543b95d7b098e915d42f41181c897407239 -address /run/containerd/containerd.sock
65535        447  0.0  0.0   1028   752 ?        Ss   12:02   0:00  \_ /pause
root         590  0.3  0.6 1301012 107000 ?      Ssl  12:02   0:13  \_ kube-controller-manager --allocate-node-cidrs=true --authentication-kubeconfig=/etc/kubernetes/controller-manager.conf --authorization-kubeconfig=/etc/kubernetes/controller-manager.conf --bind-address=127.0.0.1 --client-ca-file=/etc/kubernetes/pki/ca.crt --cluster-cidr=10.244.0.0/16 --cluster-name=kubedredger-kind --cluster-signing-cert-file=/etc/kubernetes/pki/ca.crt --cluster-signing-key-file=/etc/kubernetes/pki/ca.key --controllers=*,bootstrapsigner,tokencleaner --enable-hostpath-provisioner=true --kubeconfig=/etc/kubernetes/controller-manager.conf --leader-elect=true --requestheader-client-ca-file=/etc/kubernetes/pki/front-proxy-ca.crt --root-ca-file=/etc/kubernetes/pki/ca.crt --service-account-private-key-file=/etc/kubernetes/pki/sa.key --service-cluster-ip-range=10.96.0.0/16 --use-service-account-credentials=true
root         358  0.0  0.0 1233812 11212 ?       Sl   12:02   0:00 /usr/local/bin/containerd-shim-runc-v2 -namespace k8s.io -id ae037eb0cdd76d22ec88a10b2dcdeda42664f048647779dc79284aa9307bd944 -address /run/containerd/containerd.sock
65535        438  0.0  0.0   1028   748 ?        Ss   12:02   0:00  \_ /pause
root         701  0.4  0.3 11741500 61348 ?      Ssl  12:02   0:17  \_ etcd --advertise-client-urls=https://172.18.0.2:2379 --cert-file=/etc/kubernetes/pki/etcd/server.crt --client-cert-auth=true --data-dir=/var/lib/etcd --feature-gates=InitialCorruptCheck=true --initial-advertise-peer-urls=https://172.18.0.2:2380 --initial-cluster=kubedredger-kind-control-plane=https://172.18.0.2:2380 --key-file=/etc/kubernetes/pki/etcd/server.key --listen-client-urls=https://127.0.0.1:2379,https://172.18.0.2:2379 --listen-metrics-urls=http://127.0.0.1:2381 --listen-peer-urls=https://172.18.0.2:2380 --name=kubedredger-kind-control-plane --peer-cert-file=/etc/kubernetes/pki/etcd/peer.crt --peer-client-cert-auth=true --peer-key-file=/etc/kubernetes/pki/etcd/peer.key --peer-trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt --snapshot-count=10000 --trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt --watch-progress-notify-interval=5s
root         386  0.0  0.0 1233812 11248 ?       Sl   12:02   0:00 /usr/local/bin/containerd-shim-runc-v2 -namespace k8s.io -id a25b45f2fc2b3f2dfc400f96b6157531ea8ecdf67dd39dcb07d3ad0b863d3f93 -address /run/containerd/containerd.sock
65535        430  0.0  0.0   1028   752 ?        Ss   12:02   0:00  \_ /pause
root         597  0.7  1.7 1513976 277128 ?      Ssl  12:02   0:30  \_ kube-apiserver --advertise-address=172.18.0.2 --allow-privileged=true --authorization-mode=Node,RBAC --client-ca-file=/etc/kubernetes/pki/ca.crt --enable-admission-plugins=NodeRestriction --enable-bootstrap-token-auth=true --etcd-cafile=/etc/kubernetes/pki/etcd/ca.crt --etcd-certfile=/etc/kubernetes/pki/apiserver-etcd-client.crt --etcd-keyfile=/etc/kubernetes/pki/apiserver-etcd-client.key --etcd-servers=https://127.0.0.1:2379 --kubelet-client-certificate=/etc/kubernetes/pki/apiserver-kubelet-client.crt --kubelet-client-key=/etc/kubernetes/pki/apiserver-kubelet-client.key --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname --proxy-client-cert-file=/etc/kubernetes/pki/front-proxy-client.crt --proxy-client-key-file=/etc/kubernetes/pki/front-proxy-client.key --requestheader-allowed-names=front-proxy-client --requestheader-client-ca-file=/etc/kubernetes/pki/front-proxy-ca.crt --requestheader-extra-headers-prefix=X-Remote-Extra- --requestheader-group-headers=X-Remote-Group --requestheader-username-headers=X-Remote-User --runtime-config= --secure-port=6443 --service-account-issuer=https://kubernetes.default.svc.cluster.local --service-account-key-file=/etc/kubernetes/pki/sa.pub --service-account-signing-key-file=/etc/kubernetes/pki/sa.key --service-cluster-ip-range=10.96.0.0/16 --tls-cert-file=/etc/kubernetes/pki/apiserver.crt --tls-private-key-file=/etc/kubernetes/pki/apiserver.key
root         778  0.5  0.5 2987368 92280 ?       Ssl  12:02   0:20 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --config=/var/lib/kubelet/config.yaml --node-ip=172.18.0.2 --node-labels= --pod-infra-container-image=registry.k8s.io/pause:3.10.1 --provider-id=kind://docker/kubedredger-kind/kubedredger-kind-control-plane --runtime-cgroups=/system.slice/containerd.service
root         949  0.0  0.0 1233556 11348 ?       Sl   12:02   0:00 /usr/local/bin/containerd-shim-runc-v2 -namespace k8s.io -id d9f96f895f5f875220159adadf69b50bbac471b2aa1507f3d8a7a757cb04237d -address /run/containerd/containerd.sock
65535       1005  0.0  0.0   1028   756 ?        Ss   12:02   0:00  \_ /pause
root        1167  0.0  0.2 1281620 44880 ?       Ssl  12:02   0:00  \_ /bin/kindnetd
root         980  0.0  0.0 1233812 11328 ?       Sl   12:02   0:00 /usr/local/bin/containerd-shim-runc-v2 -namespace k8s.io -id 41d2455281d1eeb68c98101063eecb08e26c0cc982a01768e487e488baa71832 -address /run/containerd/containerd.sock
65535       1012  0.0  0.0   1028   752 ?        Ss   12:02   0:00  \_ /pause
root        1048  0.0  0.3 1274020 47540 ?       Ssl  12:02   0:00  \_ /usr/local/bin/kube-proxy --config=/var/lib/kube-proxy/config.conf --hostname-override=kubedredger-kind-control-plane
root        1399  0.0  0.0 1233812 11232 ?       Sl   12:02   0:00 /usr/local/bin/containerd-shim-runc-v2 -namespace k8s.io -id 4f699b4aeed59c2d85f9e92fb81f09c23db9f13a3faf1e75db722800409cb0c0 -address /run/containerd/containerd.sock
65535       1440  0.0  0.0   1028   752 ?        Ss   12:02   0:00  \_ /pause
65532       1716  0.0  0.4 1300984 63084 ?       Ssl  12:02   0:01  \_ /coredns -conf /etc/coredns/Corefile
root        1432  0.0  0.0 1233556 11344 ?       Sl   12:02   0:00 /usr/local/bin/containerd-shim-runc-v2 -namespace k8s.io -id 496d90da2a96fe1b6ca65799ebbf816f7b43046f3651a6dbc007b160fe88d11e -address /run/containerd/containerd.sock
65535       1494  0.0  0.0   1028   748 ?        Ss   12:02   0:00  \_ /pause
65532       1724  0.0  0.4 1301496 63280 ?       Ssl  12:02   0:01  \_ /coredns -conf /etc/coredns/Corefile
root        1457  0.0  0.0 1233556 11308 ?       Sl   12:02   0:00 /usr/local/bin/containerd-shim-runc-v2 -namespace k8s.io -id 5cd65bd6bab62220fa02fdeb7772534c404900c81eb1da0806fb5d93371dfe58 -address /run/containerd/containerd.sock
65535       1501  0.0  0.0   1028   752 ?        Ss   12:02   0:00  \_ /pause
root        1814  0.0  0.1 1267452 30708 ?       Ssl  12:02   0:00  \_ local-path-provisioner --debug start --helper-image docker.io/kindest/local-path-helper:v20241212-8ac705d0 --config /etc/config/config.json

```

## 3. Inspect pods and namespaces

Which pods are running? on which namespaces? check their YAMLs

### namespaces

```
$ kubectl get ns
NAME                 STATUS   AGE
default              Active   69m
kube-node-lease      Active   69m
kube-public          Active   69m
kube-system          Active   69m
local-path-storage   Active   69m
```

### pods

```
$ kubectl get pods -A -o wide
NAMESPACE            NAME                                                     READY   STATUS    RESTARTS   AGE   IP           NODE                             NOMINATED NODE   READINESS GATES
kube-system          coredns-66bc5c9577-b7846                                 1/1     Running   0          69m   10.244.0.3   kubedredger-kind-control-plane   <none>           <none>
kube-system          coredns-66bc5c9577-dhszf                                 1/1     Running   0          69m   10.244.0.2   kubedredger-kind-control-plane   <none>           <none>
kube-system          etcd-kubedredger-kind-control-plane                      1/1     Running   0          70m   172.18.0.2   kubedredger-kind-control-plane   <none>           <none>
kube-system          kindnet-fbpf7                                            1/1     Running   0          69m   172.18.0.2   kubedredger-kind-control-plane   <none>           <none>
kube-system          kube-apiserver-kubedredger-kind-control-plane            1/1     Running   0          70m   172.18.0.2   kubedredger-kind-control-plane   <none>           <none>
kube-system          kube-controller-manager-kubedredger-kind-control-plane   1/1     Running   0          70m   172.18.0.2   kubedredger-kind-control-plane   <none>           <none>
kube-system          kube-proxy-2zbmc                                         1/1     Running   0          69m   172.18.0.2   kubedredger-kind-control-plane   <none>           <none>
kube-system          kube-scheduler-kubedredger-kind-control-plane            1/1     Running   0          70m   172.18.0.2   kubedredger-kind-control-plane   <none>           <none>
local-path-storage   local-path-provisioner-7b8c8ddbd6-8r5t8                  1/1     Running   0          69m   10.244.0.4   kubedredger-kind-control-plane   <none>           <none>
```

### example yaml

```
kubectl get pod -n kube-system -o yaml $SCHEDULER
apiVersion: v1
kind: Pod
metadata:
  annotations:
    kubernetes.io/config.hash: 3c6b016e61a75ae702bba3ce5d5d9430
    kubernetes.io/config.mirror: 3c6b016e61a75ae702bba3ce5d5d9430
    kubernetes.io/config.seen: "2025-10-04T12:02:13.497248759Z"
    kubernetes.io/config.source: file
  creationTimestamp: "2025-10-04T12:02:20Z"
  generation: 1
  labels:
    component: kube-scheduler
    tier: control-plane
  name: kube-scheduler-kubedredger-kind-control-plane
  namespace: kube-system
  ownerReferences:
  - apiVersion: v1
    controller: true
    kind: Node
    name: kubedredger-kind-control-plane
    uid: e4878496-352c-45af-9f9f-d17c0c87673e
  resourceVersion: "368"
  uid: d791d2a0-f22f-48fa-b0ce-9f89d76f7170
spec:
  containers:
  - command:
    - kube-scheduler
    - --authentication-kubeconfig=/etc/kubernetes/scheduler.conf
    - --authorization-kubeconfig=/etc/kubernetes/scheduler.conf
    - --bind-address=127.0.0.1
    - --kubeconfig=/etc/kubernetes/scheduler.conf
    - --leader-elect=true
    image: registry.k8s.io/kube-scheduler:v1.34.0
    imagePullPolicy: IfNotPresent
    livenessProbe:
      failureThreshold: 8
      httpGet:
        host: 127.0.0.1
        path: /livez
        port: probe-port
        scheme: HTTPS
      initialDelaySeconds: 10
      periodSeconds: 10
      successThreshold: 1
      timeoutSeconds: 15
    name: kube-scheduler
    ports:
    - containerPort: 10259
      hostPort: 10259
      name: probe-port
      protocol: TCP
    readinessProbe:
      failureThreshold: 3
      httpGet:
        host: 127.0.0.1
        path: /readyz
        port: probe-port
        scheme: HTTPS
      periodSeconds: 1
      successThreshold: 1
      timeoutSeconds: 15
    resources:
      requests:
        cpu: 100m
    startupProbe:
      failureThreshold: 24
      httpGet:
        host: 127.0.0.1
        path: /livez
        port: probe-port
        scheme: HTTPS
      initialDelaySeconds: 10
      periodSeconds: 10
      successThreshold: 1
      timeoutSeconds: 15
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    volumeMounts:
    - mountPath: /etc/kubernetes/scheduler.conf
      name: kubeconfig
      readOnly: true
  dnsPolicy: ClusterFirst
  enableServiceLinks: true
  hostNetwork: true
  nodeName: kubedredger-kind-control-plane
  preemptionPolicy: PreemptLowerPriority
  priority: 2000001000
  priorityClassName: system-node-critical
  restartPolicy: Always
  schedulerName: default-scheduler
  securityContext:
    seccompProfile:
      type: RuntimeDefault
  terminationGracePeriodSeconds: 30
  tolerations:
  - effect: NoExecute
    operator: Exists
  volumes:
  - hostPath:
      path: /etc/kubernetes/scheduler.conf
      type: FileOrCreate
    name: kubeconfig
status:
  conditions:
  - lastProbeTime: null
    lastTransitionTime: "2025-10-04T12:02:20Z"
    status: "True"
    type: PodReadyToStartContainers
  - lastProbeTime: null
    lastTransitionTime: "2025-10-04T12:02:20Z"
    status: "True"
    type: Initialized
  - lastProbeTime: null
    lastTransitionTime: "2025-10-04T12:02:28Z"
    status: "True"
    type: Ready
  - lastProbeTime: null
    lastTransitionTime: "2025-10-04T12:02:28Z"
    status: "True"
    type: ContainersReady
  - lastProbeTime: null
    lastTransitionTime: "2025-10-04T12:02:20Z"
    status: "True"
    type: PodScheduled
  containerStatuses:
  - allocatedResources:
      cpu: 100m
    containerID: containerd://05c0c4b6d9e87ea589474c0d15224540f06d91f7bd7d72a809a61c97cdf16d84
    image: registry.k8s.io/kube-scheduler-amd64:v1.34.0
    imageID: sha256:46169d968e9203e8b10debaf898210fe11c94b5864c351ea0f6fcf621f659bdc
    lastState: {}
    name: kube-scheduler
    ready: true
    resources:
      requests:
        cpu: 100m
    restartCount: 0
    started: true
    state:
      running:
        startedAt: "2025-10-04T12:02:15Z"
    user:
      linux:
        gid: 0
        supplementalGroups:
        - 0
        uid: 0
  hostIP: 172.18.0.2
  hostIPs:
  - ip: 172.18.0.2
  phase: Running
  podIP: 172.18.0.2
  podIPs:
  - ip: 172.18.0.2
  qosClass: Burstable
  startTime: "2025-10-04T12:02:20Z"
```

## 4. Get container logs

How can we get container logs? get logs from a component.

### container logs

```
kubectl logs -n kube-system $SCHEDULER
I1004 12:02:16.518776       1 serving.go:386] Generated self-signed cert in-memory
W1004 12:02:17.793219       1 requestheader_controller.go:204] Unable to get configmap/extension-apiserver-authentication in kube-system.  Usually fixed by 'kubectl create rolebinding -n kube-system ROLEBINDING_NAME --role=extension-apiserver-authentication-reader --serviceaccount=YOUR_NS:YOUR_SA'
W1004 12:02:17.793244       1 authentication.go:397] Error looking up in-cluster authentication configuration: configmaps "extension-apiserver-authentication" is forbidden: User "system:kube-scheduler" cannot get resource "configmaps" in API group "" in the namespace "kube-system"
W1004 12:02:17.793256       1 authentication.go:398] Continuing without authentication configuration. This may treat all requests as anonymous.
W1004 12:02:17.793265       1 authentication.go:399] To require authentication configuration lookup to succeed, set --authentication-tolerate-lookup-failure=false
I1004 12:02:17.804116       1 server.go:175] "Starting Kubernetes Scheduler" version="v1.34.0"
I1004 12:02:17.804129       1 server.go:177] "Golang settings" GOGC="" GOMAXPROCS="" GOTRACEBACK=""
I1004 12:02:17.805247       1 configmap_cafile_content.go:205] "Starting controller" name="client-ca::kube-system::extension-apiserver-authentication::client-ca-file"
I1004 12:02:17.805272       1 shared_informer.go:349] "Waiting for caches to sync" controller="client-ca::kube-system::extension-apiserver-authentication::client-ca-file"
I1004 12:02:17.805409       1 secure_serving.go:211] Serving securely on 127.0.0.1:10259
I1004 12:02:17.805436       1 tlsconfig.go:243] "Starting DynamicServingCertificateController"
E1004 12:02:17.806578       1 reflector.go:205] "Failed to watch" err="failed to list *v1.Node: nodes is forbidden: User \"system:kube-scheduler\" cannot list resource \"nodes\" in API group \"\" at the cluster scope" logger="UnhandledError" reflector="k8s.io/client-go/informers/factory.go:160" type="*v1.Node"
E1004 12:02:17.806682       1 reflector.go:205] "Failed to watch" err="failed to list *v1.ConfigMap: configmaps \"extension-apiserver-authentication\" is forbidden: User \"system:kube-scheduler\" cannot list resource \"configmaps\" in API group \"\" in the namespace \"kube-system\"" logger="UnhandledError" reflector="runtime/asm_amd64.s:1700" type="*v1.ConfigMap"
E1004 12:02:17.806778       1 reflector.go:205] "Failed to watch" err="failed to list *v1.ResourceClaim: resourceclaims.resource.k8s.io is forbidden: User \"system:kube-scheduler\" cannot list resource \"resourceclaims\" in API group \"resource.k8s.io\" at the cluster scope" logger="UnhandledError" reflector="k8s.io/client-go/informers/factory.go:160" type="*v1.ResourceClaim"
E1004 12:02:17.806879       1 reflector.go:205] "Failed to watch" err="failed to list *v1.ReplicaSet: replicasets.apps is forbidden: User \"system:kube-scheduler\" cannot list resource \"replicasets\" in API group \"apps\" at the cluster scope" logger="UnhandledError" reflector="k8s.io/client-go/informers/factory.go:160" type="*v1.ReplicaSet"
E1004 12:02:17.807180       1 reflector.go:205] "Failed to watch" err="failed to list *v1.Pod: pods is forbidden: User \"system:kube-scheduler\" cannot list resource \"pods\" in API group \"\" at the cluster scope" logger="UnhandledError" reflector="k8s.io/client-go/informers/factory.go:160" type="*v1.Pod"
```

## 5. Create an example pod

Create and delete an example pod. Check `pod.yaml`.

```
kubectl create -f pod.yaml
kubectl delete -f pod.yaml
```

## 6. Mutate some object

Try to add custom labels to a node, or to a pod.

```
kubectl edit pod mypod
```
