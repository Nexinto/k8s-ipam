Custom resource and controller files for IP address management with Kubernetes.

See https://github.com/Nexinto/k8s-ipam-configmap and https://github.com/Nexinto/k8s-ipam-haci

# kubeipam utility

kubeipam is a command line utility to manage k8s-ipam address request objects.

You can create, inspect, list, delete address requests and manually assign IP addresses.

Usage: 

```
kubeipam: manage k8s-ipam address request objects

Usage: kupeipam [options] VERB ...
       kubeipam get                                      list addresses
       kubeipam get NAME                                 get address
       kubeipam describe NAME                            get address details
       kubeipam set NAME IPADDRESS                       set reserved address
       kubeipam [-ref REF] [-d DESCRIPTION] create NAME  create a new address request
       kubeipam delete ADDRESS                           delete address request

  -d string
    	description
  -kubeconfig string
    	location of kubeconfig
  -n string
    	namespace; 'all' or empty for all namespaces
  -ref string
    	address reference
```

List addresses in the default namespace:

```bash
kubeipam get
```

List all addresses in all namespaces:

```bash
kubeipam -n all get
```

List addresses in the `kube-system` namespace:

```bash
kubeipam -n kube-system get
```

Show details for an address:

```bash
kubeipam describe myservice
```

Request a new address:

```bash
kubeipam -d "My great service needs an address" myservice
```

(The `-d` for description is optional.)

Request a new address that is a reference of an existing address:

```bash
kubeipam -ref kubernetes.default.myservice myotherservice
```

(The reference is the `.Status.Name` you get when describing the address.)

Manually set the address (if no controller is running to do this for you):

```bash
kubeipam set myservice 10.10.9.9
```

Delete an address request:

```bash
kubeipam delete myservice
```

# Creating address requests with kubectl

To request an IP address, create `myservice-ip.yaml`:

```yaml
apiVersion: ipam.nexinto.com/v1
kind: IpAddress
metadata:
  name: myservice
spec:
  comment: My great service will be at this address
```

and create it using `kubectl apply -f myservice-ip.yaml`.

Create a reference:

```yaml
apiVersion: ipam.nexinto.com/v1
kind: IpAddress
metadata:
  name: myotherservice
spec:
  ref: kubernetes.default.myservice
```

The Spec supports the following fields:

- *description* (optional) description for this address reservation
- *name* (optional) name how the IP address management internally stores this address. The default is `$TAG.$NAMESPACE.$NAME`.
- *ref* (optional) do not create a new address; instead reuse an existing entry. Use the IPAM name (like `$TAG.$NAMESPACE.$NAME`), not the Kubernetes object name.
