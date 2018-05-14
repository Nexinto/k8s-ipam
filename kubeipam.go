package main

import (
	"flag"
	"os"
	"text/tabwriter"

	"fmt"
	ipamv1 "github.com/Nexinto/k8s-ipam/pkg/apis/ipam.nexinto.com/v1"
	ipamclientset "github.com/Nexinto/k8s-ipam/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type config struct {
	kubernetes  kubernetes.Interface
	ipamclient  ipamclientset.Interface
	namespace   string
	description string
	ref         string
}

func main() {

	var kubeconfig, namespace, description, ref string

	flag.StringVar(&kubeconfig, "kubeconfig", "", "location of kubeconfig")
	flag.StringVar(&namespace, "n", "", "namespace; 'all' or empty for all namespaces")
	flag.StringVar(&description, "d", "", "description")
	flag.StringVar(&ref, "ref", "", "address reference")

	flag.Parse()

	if namespace == "" {
		namespace = "default"
	}

	if namespace == "all" {
		namespace = metav1.NamespaceAll
	}

	if c := os.Getenv("KUBECONFIG"); c != "" {
		kubeconfig = c
	}

	if kubeconfig == "" {
		kubeconfig = os.Getenv("HOME") + "/.kube/config"
	}

	clientConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	kube, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		panic(err.Error())
	}

	ipamclient, err := ipamclientset.NewForConfig(clientConfig)
	if err != nil {
		panic(err.Error())
	}

	if kube == nil || ipamclient == nil {
		panic("muh")
	}

	config := config{
		kubernetes:  kube,
		ipamclient:  ipamclient,
		namespace:   namespace,
		description: description,
		ref:         ref,
	}

	switch flag.Arg(0) {
	case "get":
		config.get(flag.Arg(1))
	case "describe":
		config.describe(flag.Arg(1))
	case "set":
		config.set(flag.Arg(1), flag.Arg(2))
	case "create":
		config.create(flag.Arg(1))
	case "delete":
		config.delete(flag.Arg(1))
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Print(`kubeipam: manage k8s-ipam address request objects

Usage: kupeipam [options] VERB ...
       kubeipam get                                      list addresses
       kubeipam get NAME                                 get address
       kubeipam describe NAME                            get address details
       kubeipam set NAME IPADDRESS                       set reserved address
       kubeipam [-ref REF] [-d DESCRIPTION] create NAME  create a new address request
       kubeipam delete ADDRESS                           delete address request

`)
	flag.PrintDefaults()
}

func (c *config) get(arg string) {
	var addresses []ipamv1.IpAddress
	if arg == "" {
		addresslist, err := c.ipamclient.IpamV1().IpAddresses(c.namespace).List(metav1.ListOptions{})
		addresses = addresslist.Items
		if err != nil {
			panic(err)
		}
	} else {
		addr, err := c.ipamclient.IpamV1().IpAddresses(c.namespace).Get(arg, metav1.GetOptions{})
		if err != nil {
			panic(err)
		}
		addresses = []ipamv1.IpAddress{*addr}
	}
	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	if c.namespace == metav1.NamespaceAll {
		fmt.Fprintln(writer, "NAMESPACE\tNAME\tADDRESS\tREF\tDESCRIPTION")
		for _, address := range addresses {
			fmt.Fprintln(writer, address.Namespace+"\t"+address.Name+"\t"+address.Status.Address+"\t"+address.Spec.Ref+"\t"+address.Spec.Description)
		}
	} else {
		fmt.Fprintln(writer, "NAME\tADDRESS\tREF\tDESCRIPTION")
		for _, address := range addresses {
			fmt.Fprintln(writer, address.Name+"\t"+address.Status.Address+"\t"+address.Spec.Ref+"\t"+address.Spec.Description)
		}
	}
	writer.Flush()
}

func (c *config) describe(arg string) {
	addr, err := c.ipamclient.IpamV1().IpAddresses(c.namespace).Get(arg, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Name:      %s\nNamespace: %s\n\nSpec:\n  Name:        %s\n  Ref:         %s\n  Description: %s\nStatus:\n  Address:  %s\n  Name:     %s\n  Provider: %s\n",
		addr.Name,
		addr.Namespace,
		addr.Spec.Name,
		addr.Spec.Ref,
		addr.Spec.Description,
		addr.Status.Address,
		addr.Status.Name,
		addr.Status.Provider)
}

func (c *config) set(address string, ip string) {
	addr, err := c.ipamclient.IpamV1().IpAddresses(c.namespace).Get(address, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	addr.Status.Address = ip
	addr.Status.Provider = "manual (kubeipam)"

	_, err = c.ipamclient.IpamV1().IpAddresses(c.namespace).Update(addr)
	if err != nil {
		panic(err)
	}
}

func (c *config) create(address string) {
	addr := &ipamv1.IpAddress{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: c.namespace,
			Name:      address,
		},
		Spec: ipamv1.IpAddressSpec{
			Description: c.description,
			Ref:         c.ref,
		},
	}

	_, err := c.ipamclient.IpamV1().IpAddresses(c.namespace).Create(addr)
	if err != nil {
		panic(err)
	}
}

func (c *config) delete(address string) {
	err := c.ipamclient.IpamV1().IpAddresses(c.namespace).Delete(address, &metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
}
