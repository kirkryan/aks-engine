// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package vlabs

const (
	// APIVersion is the version of this API
	APIVersion = "vlabs"
)

// the orchestrators supported by vlabs
const (
	// DCOS is the string constant for DCOS orchestrator type and defaults to DCOS188
	DCOS string = "DCOS"
	// Swarm is the string constant for the Swarm orchestrator type
	Swarm string = "Swarm"
	// Kubernetes is the string constant for the Kubernetes orchestrator type
	Kubernetes string = "Kubernetes"
	// SwarmMode is the string constant for the Swarm Mode orchestrator type
	SwarmMode string = "SwarmMode"
)

// the OSTypes supported by vlabs
const (
	Windows OSType = "Windows"
	Linux   OSType = "Linux"
)

// the LinuxDistros supported by vlabs
const (
	Ubuntu            Distro = "ubuntu"
	Ubuntu1804        Distro = "ubuntu-18.04"
	RHEL              Distro = "rhel"
	CoreOS            Distro = "coreos"
	AKS1604Deprecated Distro = "aks"               // deprecated AKS 16.04 distro. Equivalent to aks-ubuntu-16.04.
	AKS1804Deprecated Distro = "aks-1804"          // deprecated AKS 18.04 distro. Equivalent to aks-ubuntu-18.04.
	AKSDockerEngine   Distro = "aks-docker-engine" // deprecated docker-engine distro.
	AKSUbuntu1604     Distro = "aks-ubuntu-16.04"
	AKSUbuntu1804     Distro = "aks-ubuntu-18.04"
	ACC1604           Distro = "acc-16.04"
)

// validation values
const (
	// MinAgentCount are the minimum number of agents per agent pool
	MinAgentCount = 1
	// MaxAgentCount are the maximum number of agents per agent pool
	MaxAgentCount = 100
	// MinPort specifies the minimum tcp port to open
	MinPort = 1
	// MaxPort specifies the maximum tcp port to open
	MaxPort = 65535
	// MaxDisks specifies the maximum attached disks to add to the cluster
	MaxDisks = 4
	// MinDiskSizeGB specifies the minimum attached disk size
	MinDiskSizeGB = 1
	// MaxDiskSizeGB specifies the maximum attached disk size
	MaxDiskSizeGB = 1023
	// MinIPAddressCount specifies the minimum number of IP addresses per network interface
	MinIPAddressCount = 1
	// MaxIPAddressCount specifies the maximum number of IP addresses per network interface
	MaxIPAddressCount = 256
)

// Availability profiles
const (
	// AvailabilitySet means that the vms are in an availability set
	AvailabilitySet = "AvailabilitySet"
	// VirtualMachineScaleSets means that the vms are in a virtual machine scaleset
	VirtualMachineScaleSets = "VirtualMachineScaleSets"
)

// storage profiles
const (
	// StorageAccount means that the nodes use raw storage accounts for their os and attached volumes
	StorageAccount = "StorageAccount"
	// ManagedDisks means that the nodes use managed disks for their os and attached volumes
	ManagedDisks = "ManagedDisks"
	// Ephemeral means that the node's os disk is ephemeral. This is not compatible with attached volumes.
	Ephemeral = "Ephemeral"
)

// Supported container runtimes
const (
	Docker         = "docker"
	KataContainers = "kata-containers"
	Containerd     = "containerd"
)

var (
	// NetworkPluginValues holds the valid values for network plugin implementation
	NetworkPluginValues = [...]string{"", "kubenet", "azure", NetworkPluginCilium, "flannel"}

	// NetworkPolicyValues holds the valid values for a network policy
	// "azure" and "none" are there for backwards-compatibility
	NetworkPolicyValues = [...]string{"", "calico", NetworkPolicyCilium, "azure", "none"}

	// ContainerRuntimeValues holds the valid values for container runtimes
	ContainerRuntimeValues = [...]string{"", Docker, KataContainers, Containerd}

	// DistroValues holds the valid values for OS distros
	DistroValues = []Distro{"", Ubuntu, Ubuntu1804, RHEL, CoreOS, AKSUbuntu1604, AKSUbuntu1804, ACC1604}

	// DependenciesLocationValues holds the valid values for dependencies location
	DependenciesLocationValues = []DependenciesLocation{"", AzureStackDependenciesLocationPublic, AzureStackDependenciesLocationChina, AzureStackDependenciesLocationGerman, AzureStackDependenciesLocationUSGovernment}
)

// Kubernetes configuration
const (
	// KubernetesMinMaxPods is the minimum valid value for MaxPods, necessary for running kube-system pods
	KubernetesMinMaxPods = 5
)

// vlabs default configuration
const (
	// DefaultNetworkPlugin defines the network plugin to use by default
	DefaultNetworkPlugin = "azure"
	// DefaultNetworkPluginWindows defines the network plugin to use by default for clusters with Windows agent pools
	DefaultNetworkPluginWindows = "azure"
	// DefaultNetworkPolicy defines the network policy to use by default
	DefaultNetworkPolicy = ""
	// NetworkPolicyCilium is the string expression for cilium network policy config option
	NetworkPolicyCilium = "cilium"
	// NetworkPluginCilium is the string expression for cilium network policy config option
	NetworkPluginCilium = NetworkPolicyCilium
)

const (
	// AgentPoolProfileRoleEmpty is the empty role
	AgentPoolProfileRoleEmpty AgentPoolProfileRole = ""
	// AgentPoolProfileRoleInfra is the infra role
	AgentPoolProfileRoleInfra AgentPoolProfileRole = "infra"
)

const (
	// AzureStackCloud is a const string reference identifier for Azure Stack cloud
	AzureStackCloud = "AzureStackCloud"
	// MaxAzureStackManagedDiskSize is max etcd disk size supported on AzureStackCloud
	MaxAzureStackManagedDiskSize = 1023
)

const (
	// AzureADIdentitySystem is a const string reference identifier for Azure AD identity System
	AzureADIdentitySystem = "azure_ad"
	// ADFSIdentitySystem is a const string reference identifier for ADFS identity System
	ADFSIdentitySystem = "adfs"
)

const (
	// AzureStackDependenciesLocationPublic indicates to get dependencies from in AzurePublic cloud
	AzureStackDependenciesLocationPublic = "public"
	// AzureStackDependenciesLocationChina indicates to get dependencies from AzureChina cloud
	AzureStackDependenciesLocationChina = "china"
	// AzureStackDependenciesLocationGerman indicates to get dependencies from AzureGerman cloud
	AzureStackDependenciesLocationGerman = "german"
	// AzureStackDependenciesLocationUSGovernment indicates to get dependencies from AzureUSGovernment cloud
	AzureStackDependenciesLocationUSGovernment = "usgovernment"
)

const (
	// ClientSecretAuthMethod indicates to use client seret for authentication
	ClientSecretAuthMethod = "client_secret"
	// ClientCertificateAuthMethod indicates to use client certificate for authentication
	ClientCertificateAuthMethod = "client_certificate"
)

// StandardLoadBalancerSku is the string const for Azure Standard Load Balancer
const StandardLoadBalancerSku = "Standard"
