## Sample Kube-Operator
The kube operator creates a custom resources and watches for changes.

### Build
```bash
# pull all the libraries needed (this may take a while with all the Kubernetes dependencies)
glide install

# build the go binary
CGO_ENABLED=0 GOOS=linux go build

# build the docker container
docker build -t kube-operator:0.1 .
```

### Start the Operator

```bash
# Create the sample kube-operator
$ kubectl create -f kube-operator.yaml

# Wait for the pod status to be Running
$ kubectl get pod
NAME                              READY     STATUS    RESTARTS   AGE
kube-operator-821691060-m5vqp     1/1       Running   0          3m

# View the samples CRD
$ kubectl get crd
NAME                 AGE
samples.openebs.io   10m
```

### Create the Sample Kube Resource
```bash
# Create the sample
$ kubectl create -f kube-resource.yaml

# See the sample resource
$ kubectl get samples.openebs.io
NAME               AGE
openebs-sample     7m
```

### Modify the Sample Resource
Change the value of the `Hello` property in `kube-resource.yaml`, then apply the new yaml.
```bash
kubectl apply -f kube-resource.yaml
```

### Logs

Notice the added and modified Hello= text in the log below

```bash
$ kubectl logs -l app=sample-operator
2017-06-29 00:11:22.738629 I | sample: Creating the sample resource
2017-06-29 00:11:22.749325 I | OpenEBS: creating sample resource
2017-06-29 00:11:22.794635 I | OpenEBS: did not yet find resource sample at apis/mycompany.io/v1alpha1/samples. the server could not find the requested resource
2017-06-29 00:11:28.797620 I | OpenEBS: did not yet find resource sample at apis/mycompany.io/v1alpha1/samples. the server could not find the requested resource
2017-06-29 00:11:34.797912 I | sample: Managing the sample resource
2017-06-29 00:11:34.797932 I | sample: finding existing samples...
2017-06-29 00:11:34.799276 I | sample: found 0 samples.
2017-06-29 00:11:34.799296 I | OpenEBS: start watching sample resource in namespace default at 27064
2017-06-29 00:12:07.605948 I | sample: Added Sample 'openebs-sample' with Hello=Cloud!
2017-06-29 00:14:29.553035 I | sample: Modified Sample 'openebs-sample' with Hello=Mule can Fly!!
```

### Cleanup
```bash
kubectl delete -f kube-resource.yaml
kubectl delete -f kube-operator.yaml
kubectl delete crd samples.openebs.io
```
