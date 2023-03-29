package main

import "time"

type VmObjectResponse struct {
	Entities   []EntitiesObject       `json:"entities"`
	APIVersion string                 `json:"api_version"`
	Metadata   MetadataObjectResponse `json:"metadata"`
}

type StatusObject struct {
	Name                      string                          `json:"name"`
	State                     string                          `json:"state"`
	AvailabilityZoneReference AvailabilityZoneReferenceObject `json:"availability_zone_reference"`
	MessageList               []MessageListObject             `json:"message_list"`
	ClusterReference          ClusterReferenceObject          `json:"cluster_reference"`
	Resources                 ResourcesSpecObject             `json:"resources"`
	Description               string                          `json:"description"`
}

type MessageListObject struct {
	Message string       `json:"message"`
	Reason  string       `json:"reason"`
	Details DetailsOject `json:"details"`
}

type DetailsOject struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}

type SpecObject struct {
	Name                      string                          `json:"name"`
	AvailabilityZoneReference AvailabilityZoneReferenceObject `json:"availability_zone_reference"`
	Description               string                          `json:"description"`
	Resources                 ResourcesSpecObject             `json:"resources"`
	ClusterReference          ClusterReferenceObject          `json:"cluster_reference"`
}

type AvailabilityZoneReferenceObject struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type ResourcesSpecObject struct {
	NumThreadsPerCore int `json:"num_threads_per_core"`
	GpuList           []struct {
		Vendor   string `json:"vendor"`
		Mode     string `json:"mode"`
		DeviceID int    `json:"device_id"`
	} `json:"gpu_list"`
	MemorySizeMib int `json:"memory_size_mib"`
	BootConfig    struct {
		BootDeviceOrderList []string `json:"boot_device_order_list"`
		BootDevice          struct {
			DiskAddress struct {
				DeviceIndex int    `json:"device_index"`
				AdapterType string `json:"adapter_type"`
			} `json:"disk_address"`
			MacAddress string `json:"mac_address"`
		} `json:"boot_device"`
		DataSourceReference struct {
			URL  string `json:"url"`
			Kind string `json:"kind"`
			UUID string `json:"uuid"`
			Name string `json:"name"`
		} `json:"data_source_reference"`
		BootType string `json:"boot_type"`
	} `json:"boot_config"`
	DiskList []struct {
		UUID          string `json:"uuid"`
		DiskSizeBytes int    `json:"disk_size_bytes"`
		StorageConfig struct {
			FlashMode                 string `json:"flash_mode"`
			StorageContainerReference struct {
				URL  string `json:"url"`
				Kind string `json:"kind"`
				UUID string `json:"uuid"`
				Name string `json:"name"`
			} `json:"storage_container_reference"`
		} `json:"storage_config"`
		DeviceProperties struct {
			DeviceType  string `json:"device_type"`
			DiskAddress struct {
				DeviceIndex int    `json:"device_index"`
				AdapterType string `json:"adapter_type"`
			} `json:"disk_address"`
		} `json:"device_properties"`
		DataSourceReference struct {
			IsDirectAttach bool   `json:"is_direct_attach"`
			URL            string `json:"url"`
			Kind           string `json:"kind"`
			UUID           string `json:"uuid"`
			Name           string `json:"name"`
		} `json:"data_source_reference"`
		DiskSizeMib          int `json:"disk_size_mib"`
		VolumeGroupReference struct {
			URL  string `json:"url"`
			Kind string `json:"kind"`
			UUID string `json:"uuid"`
			Name string `json:"name"`
		} `json:"volume_group_reference"`
	} `json:"disk_list"`
	SerialPortList []struct {
		Index       int  `json:"index"`
		IsConnected bool `json:"is_connected"`
	} `json:"serial_port_list"`
	IsVcpuHardPinned bool `json:"is_vcpu_hard_pinned"`
	GuestTools       struct {
		NutanixGuestTools struct {
			NgtState              string   `json:"ngt_state"`
			IsoMountState         string   `json:"iso_mount_state"`
			State                 string   `json:"state"`
			Version               string   `json:"version"`
			EnabledCapabilityList []string `json:"enabled_capability_list"`
			Credentials           struct {
				Username string `json:"username"`
				Password string `json:"password"`
			} `json:"credentials"`
		} `json:"nutanix_guest_tools"`
	} `json:"guest_tools"`
	PowerState                    string `json:"power_state"`
	NumVcpusPerSocket             int    `json:"num_vcpus_per_socket"`
	NumSockets                    int    `json:"num_sockets"`
	HardwareVirtualizationEnabled bool   `json:"hardware_virtualization_enabled"`
	StorageConfig                 struct {
		QosPolicy struct {
			ThrottledIops int `json:"throttled_iops"`
		} `json:"qos_policy"`
		FlashMode string `json:"flash_mode"`
	} `json:"storage_config"`
	MachineType           string `json:"machine_type"`
	HardwareClockTimezone string `json:"hardware_clock_timezone"`
	GuestCustomization    struct {
		CloudInit struct {
			MetaData        string `json:"meta_data"`
			UserData        string `json:"user_data"`
			CustomKeyValues struct {
				AdditionalProp1 string `json:"additionalProp1"`
				AdditionalProp2 string `json:"additionalProp2"`
				AdditionalProp3 string `json:"additionalProp3"`
			} `json:"custom_key_values"`
		} `json:"cloud_init"`
		IsOverridable bool `json:"is_overridable"`
		Sysprep       struct {
			InstallType     string `json:"install_type"`
			UnattendXML     string `json:"unattend_xml"`
			CustomKeyValues struct {
				AdditionalProp1 string `json:"additionalProp1"`
				AdditionalProp2 string `json:"additionalProp2"`
				AdditionalProp3 string `json:"additionalProp3"`
			} `json:"custom_key_values"`
		} `json:"sysprep"`
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"guest_customization"`
	PowerStateMechanism struct {
		GuestTransitionConfig struct {
			ShouldFailOnScriptFailure bool `json:"should_fail_on_script_failure"`
			EnableScriptExec          bool `json:"enable_script_exec"`
		} `json:"guest_transition_config"`
		Mechanism string `json:"mechanism"`
	} `json:"power_state_mechanism"`
	VgaConsoleEnabled       bool `json:"vga_console_enabled"`
	MemoryOvercommitEnabled bool `json:"memory_overcommit_enabled"`
	VnumaConfig             struct {
		NumVnumaNodes int `json:"num_vnuma_nodes"`
	} `json:"vnuma_config"`
	NicList []struct {
		NicType        string `json:"nic_type"`
		UUID           string `json:"uuid"`
		IPEndpointList []struct {
			IP                 string   `json:"ip"`
			Type               string   `json:"type"`
			GatewayAddressList []string `json:"gateway_address_list"`
			PrefixLength       int      `json:"prefix_length"`
			IPType             string   `json:"ip_type"`
		} `json:"ip_endpoint_list"`
		NumQueues                     int      `json:"num_queues"`
		SecondaryIPAddressList        []string `json:"secondary_ip_address_list"`
		NetworkFunctionNicType        string   `json:"network_function_nic_type"`
		NetworkFunctionChainReference struct {
			Kind string `json:"kind"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"network_function_chain_reference"`
		VlanMode        string `json:"vlan_mode"`
		MacAddress      string `json:"mac_address"`
		SubnetReference struct {
			Kind string `json:"kind"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"subnet_reference"`
		Model           string `json:"model"`
		IsConnected     bool   `json:"is_connected"`
		TrunkedVlanList []int  `json:"trunked_vlan_list"`
	} `json:"nic_list"`
	GuestOsID         string `json:"guest_os_id"`
	IsAgentVM         bool   `json:"is_agent_vm"`
	GpuConsoleEnabled bool   `json:"gpu_console_enabled"`
	VtpmConfig        struct {
		VtpmEnabled         bool `json:"vtpm_enabled"`
		DataSourceReference struct {
			URL  string `json:"url"`
			Kind string `json:"kind"`
			UUID string `json:"uuid"`
			Name string `json:"name"`
		} `json:"data_source_reference"`
	} `json:"vtpm_config"`
	EnableCPUPassthrough bool `json:"enable_cpu_passthrough"`
	ParentReference      struct {
		URL  string `json:"url"`
		Kind string `json:"kind"`
		UUID string `json:"uuid"`
		Name string `json:"name"`
	} `json:"parent_reference"`
	DisableBranding bool `json:"disable_branding"`
}

type ClusterReferenceObject struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type MetadataObject struct {
	LastUpdateTime       time.Time              `json:"last_update_time"`
	UseCategoriesMapping bool                   `json:"use_categories_mapping"`
	Kind                 string                 `json:"kind"`
	UUID                 string                 `json:"uuid"`
	ProjectReference     ProjectReferenceObject `json:"project_reference"`
	CreationTime         time.Time              `json:"creation_time"`
	SpecVersion          int                    `json:"spec_version"`
	SpecHash             string                 `json:"spec_hash"`
	CategoriesMapping    CategoriesMapping      `json:"categories_mapping"`
	ShouldForceTranslate bool                   `json:"should_force_translate"`
	EntityVersion        string                 `json:"entity_version"`
	OwnerReference       OwnerReferenceObject   `json:"owner_reference"`
	Categories           CategoriesObject       `json:"categories"`
	Name                 string                 `json:"name"`
}

type ProjectReferenceObject struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type CategoriesMapping struct {
	AdditionalProp1 []string `json:"additionalProp1"`
	AdditionalProp2 []string `json:"additionalProp2"`
	AdditionalProp3 []string `json:"additionalProp3"`
}

type OwnerReferenceObject struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type CategoriesObject struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}
