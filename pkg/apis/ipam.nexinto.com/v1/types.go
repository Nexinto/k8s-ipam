package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// An IP address object
type IpAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IpAddressSpec   `json:"spec"`
	Status IpAddressStatus `json:"status"`
}

type IpAddressSpec struct {

	// The IPAM will explicitly use this name instead of the default name of 'cluster.namespace.service'.
	Name string `json:"name"`

	// References an existing address by its IPAM name.
	Ref string `json:"ref"`

	// Description for this address object.
	Description string `json:"description"`
}

type IpAddressStatus struct {

	// The address if reservation was successful
	Address string `json:"address"`

	// Name of the address provider
	Provider string `json:"provider"`

	// The IPAM name for this address object
	Name string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IpAddressList is a list of IpAddress resources
type IpAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []IpAddress `json:"items"`
}
