// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StatusResponse Health and status information of daemon
//
// +k8s:deepcopy-gen=true
//
// swagger:model StatusResponse
type StatusResponse struct {

	// Status of core datapath attachment mode
	AttachMode AttachMode `json:"attach-mode,omitempty"`

	// Status of Mutual Authentication certificate provider
	AuthCertificateProvider *Status `json:"auth-certificate-provider,omitempty"`

	// Status of bandwidth manager
	BandwidthManager *BandwidthManager `json:"bandwidth-manager,omitempty"`

	// Status of BPF maps
	BpfMaps *BPFMapStatus `json:"bpf-maps,omitempty"`

	// Status of Cilium daemon
	Cilium *Status `json:"cilium,omitempty"`

	// When supported by the API, this client ID should be used by the
	// client when making another request to the server.
	// See for example "/cluster/nodes".
	//
	ClientID int64 `json:"client-id,omitempty"`

	// Status of clock source
	ClockSource *ClockSource `json:"clock-source,omitempty"`

	// Status of cluster
	Cluster *ClusterStatus `json:"cluster,omitempty"`

	// Status of ClusterMesh
	ClusterMesh *ClusterMeshStatus `json:"cluster-mesh,omitempty"`

	// Status of CNI chaining
	CniChaining *CNIChainingStatus `json:"cni-chaining,omitempty"`

	// Status of the CNI configuration file
	CniFile *Status `json:"cni-file,omitempty"`

	// Status of configmap settings synchronization
	ConfigSettings *ConfigSettings `json:"config-settings,omitempty"`

	// Status of local container runtime
	ContainerRuntime *Status `json:"container-runtime,omitempty"`

	// Status of all endpoint controllers
	Controllers ControllerStatuses `json:"controllers,omitempty"`

	// Status of datapath mode
	DatapathMode DatapathMode `json:"datapath-mode,omitempty"`

	// Status of transparent encryption
	Encryption *EncryptionStatus `json:"encryption,omitempty"`

	// Status of the host firewall
	HostFirewall *HostFirewall `json:"host-firewall,omitempty"`

	// Status of Hubble server
	Hubble *HubbleStatus `json:"hubble,omitempty"`

	// Status of identity range of the cluster
	IdentityRange *IdentityRange `json:"identity-range,omitempty"`

	// Status of IP address management
	Ipam *IPAMStatus `json:"ipam,omitempty"`

	// Status of IPv4 BIG TCP
	IPV4BigTCP *IPV4BigTCP `json:"ipv4-big-tcp,omitempty"`

	// Status of IPv6 BIG TCP
	IPV6BigTCP *IPV6BigTCP `json:"ipv6-big-tcp,omitempty"`

	// Status of kube-proxy replacement
	KubeProxyReplacement *KubeProxyReplacement `json:"kube-proxy-replacement,omitempty"`

	// Status of Kubernetes integration
	Kubernetes *K8sStatus `json:"kubernetes,omitempty"`

	// Status of key/value datastore
	Kvstore *Status `json:"kvstore,omitempty"`

	// Status of masquerading
	Masquerading *Masquerading `json:"masquerading,omitempty"`

	// Status of the node monitor
	NodeMonitor *MonitorStatus `json:"nodeMonitor,omitempty"`

	// Status of proxy
	Proxy *ProxyStatus `json:"proxy,omitempty"`

	// Status of routing
	Routing *Routing `json:"routing,omitempty"`

	// Status of SRv6
	Srv6 *Srv6 `json:"srv6,omitempty"`

	// List of stale information in the status
	Stale map[string]strfmt.DateTime `json:"stale,omitempty"`
}

// Validate validates this status response
func (m *StatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthCertificateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBandwidthManager(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpfMaps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCilium(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClockSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterMesh(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCniChaining(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCniFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerRuntime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateControllers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatapathMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostFirewall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHubble(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentityRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIpam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV4BigTCP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV6BigTCP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubeProxyReplacement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKvstore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMasquerading(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeMonitor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrv6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStale(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusResponse) validateAttachMode(formats strfmt.Registry) error {
	if swag.IsZero(m.AttachMode) { // not required
		return nil
	}

	if err := m.AttachMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attach-mode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("attach-mode")
		}
		return err
	}

	return nil
}

func (m *StatusResponse) validateAuthCertificateProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthCertificateProvider) { // not required
		return nil
	}

	if m.AuthCertificateProvider != nil {
		if err := m.AuthCertificateProvider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth-certificate-provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth-certificate-provider")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateBandwidthManager(formats strfmt.Registry) error {
	if swag.IsZero(m.BandwidthManager) { // not required
		return nil
	}

	if m.BandwidthManager != nil {
		if err := m.BandwidthManager.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bandwidth-manager")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bandwidth-manager")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateBpfMaps(formats strfmt.Registry) error {
	if swag.IsZero(m.BpfMaps) { // not required
		return nil
	}

	if m.BpfMaps != nil {
		if err := m.BpfMaps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bpf-maps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bpf-maps")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateCilium(formats strfmt.Registry) error {
	if swag.IsZero(m.Cilium) { // not required
		return nil
	}

	if m.Cilium != nil {
		if err := m.Cilium.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cilium")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cilium")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateClockSource(formats strfmt.Registry) error {
	if swag.IsZero(m.ClockSource) { // not required
		return nil
	}

	if m.ClockSource != nil {
		if err := m.ClockSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clock-source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clock-source")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateClusterMesh(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterMesh) { // not required
		return nil
	}

	if m.ClusterMesh != nil {
		if err := m.ClusterMesh.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster-mesh")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster-mesh")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateCniChaining(formats strfmt.Registry) error {
	if swag.IsZero(m.CniChaining) { // not required
		return nil
	}

	if m.CniChaining != nil {
		if err := m.CniChaining.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cni-chaining")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cni-chaining")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateCniFile(formats strfmt.Registry) error {
	if swag.IsZero(m.CniFile) { // not required
		return nil
	}

	if m.CniFile != nil {
		if err := m.CniFile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cni-file")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cni-file")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateConfigSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigSettings) { // not required
		return nil
	}

	if m.ConfigSettings != nil {
		if err := m.ConfigSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config-settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config-settings")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateContainerRuntime(formats strfmt.Registry) error {
	if swag.IsZero(m.ContainerRuntime) { // not required
		return nil
	}

	if m.ContainerRuntime != nil {
		if err := m.ContainerRuntime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("container-runtime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("container-runtime")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateControllers(formats strfmt.Registry) error {
	if swag.IsZero(m.Controllers) { // not required
		return nil
	}

	if err := m.Controllers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("controllers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("controllers")
		}
		return err
	}

	return nil
}

func (m *StatusResponse) validateDatapathMode(formats strfmt.Registry) error {
	if swag.IsZero(m.DatapathMode) { // not required
		return nil
	}

	if err := m.DatapathMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("datapath-mode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("datapath-mode")
		}
		return err
	}

	return nil
}

func (m *StatusResponse) validateEncryption(formats strfmt.Registry) error {
	if swag.IsZero(m.Encryption) { // not required
		return nil
	}

	if m.Encryption != nil {
		if err := m.Encryption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("encryption")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateHostFirewall(formats strfmt.Registry) error {
	if swag.IsZero(m.HostFirewall) { // not required
		return nil
	}

	if m.HostFirewall != nil {
		if err := m.HostFirewall.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host-firewall")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host-firewall")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateHubble(formats strfmt.Registry) error {
	if swag.IsZero(m.Hubble) { // not required
		return nil
	}

	if m.Hubble != nil {
		if err := m.Hubble.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hubble")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hubble")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateIdentityRange(formats strfmt.Registry) error {
	if swag.IsZero(m.IdentityRange) { // not required
		return nil
	}

	if m.IdentityRange != nil {
		if err := m.IdentityRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity-range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identity-range")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateIpam(formats strfmt.Registry) error {
	if swag.IsZero(m.Ipam) { // not required
		return nil
	}

	if m.Ipam != nil {
		if err := m.Ipam.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipam")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipam")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateIPV4BigTCP(formats strfmt.Registry) error {
	if swag.IsZero(m.IPV4BigTCP) { // not required
		return nil
	}

	if m.IPV4BigTCP != nil {
		if err := m.IPV4BigTCP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv4-big-tcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipv4-big-tcp")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateIPV6BigTCP(formats strfmt.Registry) error {
	if swag.IsZero(m.IPV6BigTCP) { // not required
		return nil
	}

	if m.IPV6BigTCP != nil {
		if err := m.IPV6BigTCP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv6-big-tcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipv6-big-tcp")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateKubeProxyReplacement(formats strfmt.Registry) error {
	if swag.IsZero(m.KubeProxyReplacement) { // not required
		return nil
	}

	if m.KubeProxyReplacement != nil {
		if err := m.KubeProxyReplacement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kube-proxy-replacement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kube-proxy-replacement")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateKubernetes(formats strfmt.Registry) error {
	if swag.IsZero(m.Kubernetes) { // not required
		return nil
	}

	if m.Kubernetes != nil {
		if err := m.Kubernetes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubernetes")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateKvstore(formats strfmt.Registry) error {
	if swag.IsZero(m.Kvstore) { // not required
		return nil
	}

	if m.Kvstore != nil {
		if err := m.Kvstore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kvstore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kvstore")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateMasquerading(formats strfmt.Registry) error {
	if swag.IsZero(m.Masquerading) { // not required
		return nil
	}

	if m.Masquerading != nil {
		if err := m.Masquerading.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("masquerading")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("masquerading")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateNodeMonitor(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeMonitor) { // not required
		return nil
	}

	if m.NodeMonitor != nil {
		if err := m.NodeMonitor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeMonitor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeMonitor")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateProxy(formats strfmt.Registry) error {
	if swag.IsZero(m.Proxy) { // not required
		return nil
	}

	if m.Proxy != nil {
		if err := m.Proxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateRouting(formats strfmt.Registry) error {
	if swag.IsZero(m.Routing) { // not required
		return nil
	}

	if m.Routing != nil {
		if err := m.Routing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("routing")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateSrv6(formats strfmt.Registry) error {
	if swag.IsZero(m.Srv6) { // not required
		return nil
	}

	if m.Srv6 != nil {
		if err := m.Srv6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("srv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("srv6")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateStale(formats strfmt.Registry) error {
	if swag.IsZero(m.Stale) { // not required
		return nil
	}

	for k := range m.Stale {

		if err := validate.FormatOf("stale"+"."+k, "body", "date-time", m.Stale[k].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validate this status response based on the context it is used
func (m *StatusResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthCertificateProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBandwidthManager(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBpfMaps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCilium(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClockSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterMesh(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCniChaining(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCniFile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfigSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContainerRuntime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateControllers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatapathMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEncryption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostFirewall(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHubble(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentityRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIpam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPV4BigTCP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPV6BigTCP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKubeProxyReplacement(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKubernetes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKvstore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMasquerading(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeMonitor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSrv6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusResponse) contextValidateAttachMode(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.AttachMode) { // not required
		return nil
	}

	if err := m.AttachMode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attach-mode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("attach-mode")
		}
		return err
	}

	return nil
}

func (m *StatusResponse) contextValidateAuthCertificateProvider(ctx context.Context, formats strfmt.Registry) error {

	if m.AuthCertificateProvider != nil {

		if swag.IsZero(m.AuthCertificateProvider) { // not required
			return nil
		}

		if err := m.AuthCertificateProvider.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth-certificate-provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth-certificate-provider")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateBandwidthManager(ctx context.Context, formats strfmt.Registry) error {

	if m.BandwidthManager != nil {

		if swag.IsZero(m.BandwidthManager) { // not required
			return nil
		}

		if err := m.BandwidthManager.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bandwidth-manager")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bandwidth-manager")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateBpfMaps(ctx context.Context, formats strfmt.Registry) error {

	if m.BpfMaps != nil {

		if swag.IsZero(m.BpfMaps) { // not required
			return nil
		}

		if err := m.BpfMaps.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bpf-maps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bpf-maps")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateCilium(ctx context.Context, formats strfmt.Registry) error {

	if m.Cilium != nil {

		if swag.IsZero(m.Cilium) { // not required
			return nil
		}

		if err := m.Cilium.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cilium")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cilium")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateClockSource(ctx context.Context, formats strfmt.Registry) error {

	if m.ClockSource != nil {

		if swag.IsZero(m.ClockSource) { // not required
			return nil
		}

		if err := m.ClockSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clock-source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clock-source")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {

		if swag.IsZero(m.Cluster) { // not required
			return nil
		}

		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateClusterMesh(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterMesh != nil {

		if swag.IsZero(m.ClusterMesh) { // not required
			return nil
		}

		if err := m.ClusterMesh.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster-mesh")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster-mesh")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateCniChaining(ctx context.Context, formats strfmt.Registry) error {

	if m.CniChaining != nil {

		if swag.IsZero(m.CniChaining) { // not required
			return nil
		}

		if err := m.CniChaining.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cni-chaining")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cni-chaining")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateCniFile(ctx context.Context, formats strfmt.Registry) error {

	if m.CniFile != nil {

		if swag.IsZero(m.CniFile) { // not required
			return nil
		}

		if err := m.CniFile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cni-file")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cni-file")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateConfigSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.ConfigSettings != nil {

		if swag.IsZero(m.ConfigSettings) { // not required
			return nil
		}

		if err := m.ConfigSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config-settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config-settings")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateContainerRuntime(ctx context.Context, formats strfmt.Registry) error {

	if m.ContainerRuntime != nil {

		if swag.IsZero(m.ContainerRuntime) { // not required
			return nil
		}

		if err := m.ContainerRuntime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("container-runtime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("container-runtime")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateControllers(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Controllers.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("controllers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("controllers")
		}
		return err
	}

	return nil
}

func (m *StatusResponse) contextValidateDatapathMode(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.DatapathMode) { // not required
		return nil
	}

	if err := m.DatapathMode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("datapath-mode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("datapath-mode")
		}
		return err
	}

	return nil
}

func (m *StatusResponse) contextValidateEncryption(ctx context.Context, formats strfmt.Registry) error {

	if m.Encryption != nil {

		if swag.IsZero(m.Encryption) { // not required
			return nil
		}

		if err := m.Encryption.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("encryption")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateHostFirewall(ctx context.Context, formats strfmt.Registry) error {

	if m.HostFirewall != nil {

		if swag.IsZero(m.HostFirewall) { // not required
			return nil
		}

		if err := m.HostFirewall.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host-firewall")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host-firewall")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateHubble(ctx context.Context, formats strfmt.Registry) error {

	if m.Hubble != nil {

		if swag.IsZero(m.Hubble) { // not required
			return nil
		}

		if err := m.Hubble.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hubble")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hubble")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateIdentityRange(ctx context.Context, formats strfmt.Registry) error {

	if m.IdentityRange != nil {

		if swag.IsZero(m.IdentityRange) { // not required
			return nil
		}

		if err := m.IdentityRange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity-range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identity-range")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateIpam(ctx context.Context, formats strfmt.Registry) error {

	if m.Ipam != nil {

		if swag.IsZero(m.Ipam) { // not required
			return nil
		}

		if err := m.Ipam.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipam")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipam")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateIPV4BigTCP(ctx context.Context, formats strfmt.Registry) error {

	if m.IPV4BigTCP != nil {

		if swag.IsZero(m.IPV4BigTCP) { // not required
			return nil
		}

		if err := m.IPV4BigTCP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv4-big-tcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipv4-big-tcp")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateIPV6BigTCP(ctx context.Context, formats strfmt.Registry) error {

	if m.IPV6BigTCP != nil {

		if swag.IsZero(m.IPV6BigTCP) { // not required
			return nil
		}

		if err := m.IPV6BigTCP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv6-big-tcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipv6-big-tcp")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateKubeProxyReplacement(ctx context.Context, formats strfmt.Registry) error {

	if m.KubeProxyReplacement != nil {

		if swag.IsZero(m.KubeProxyReplacement) { // not required
			return nil
		}

		if err := m.KubeProxyReplacement.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kube-proxy-replacement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kube-proxy-replacement")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateKubernetes(ctx context.Context, formats strfmt.Registry) error {

	if m.Kubernetes != nil {

		if swag.IsZero(m.Kubernetes) { // not required
			return nil
		}

		if err := m.Kubernetes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubernetes")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateKvstore(ctx context.Context, formats strfmt.Registry) error {

	if m.Kvstore != nil {

		if swag.IsZero(m.Kvstore) { // not required
			return nil
		}

		if err := m.Kvstore.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kvstore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kvstore")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateMasquerading(ctx context.Context, formats strfmt.Registry) error {

	if m.Masquerading != nil {

		if swag.IsZero(m.Masquerading) { // not required
			return nil
		}

		if err := m.Masquerading.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("masquerading")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("masquerading")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateNodeMonitor(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeMonitor != nil {

		if swag.IsZero(m.NodeMonitor) { // not required
			return nil
		}

		if err := m.NodeMonitor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeMonitor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeMonitor")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateProxy(ctx context.Context, formats strfmt.Registry) error {

	if m.Proxy != nil {

		if swag.IsZero(m.Proxy) { // not required
			return nil
		}

		if err := m.Proxy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateRouting(ctx context.Context, formats strfmt.Registry) error {

	if m.Routing != nil {

		if swag.IsZero(m.Routing) { // not required
			return nil
		}

		if err := m.Routing.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("routing")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) contextValidateSrv6(ctx context.Context, formats strfmt.Registry) error {

	if m.Srv6 != nil {

		if swag.IsZero(m.Srv6) { // not required
			return nil
		}

		if err := m.Srv6.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("srv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("srv6")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusResponse) UnmarshalBinary(b []byte) error {
	var res StatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
