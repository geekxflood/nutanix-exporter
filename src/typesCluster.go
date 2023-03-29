package main

import "time"

type ClusterObjectResponse struct {
	Entities []struct {
		Status struct {
			State       string `json:"state"`
			MessageList []struct {
				Message string `json:"message"`
				Reason  string `json:"reason"`
				Details struct {
					AdditionalProp1 string `json:"additionalProp1"`
					AdditionalProp2 string `json:"additionalProp2"`
					AdditionalProp3 string `json:"additionalProp3"`
				} `json:"details"`
			} `json:"message_list"`
			Name      string `json:"name"`
			Resources struct {
				Nodes struct {
					HypervisorServerList []struct {
						IP      string `json:"ip"`
						Version string `json:"version"`
						Type    string `json:"type"`
					} `json:"hypervisor_server_list"`
				} `json:"nodes"`
				Config struct {
					GpuDriverVersion string `json:"gpu_driver_version"`
					ClientAuth       struct {
						Status  string `json:"status"`
						CaChain string `json:"ca_chain"`
						Name    string `json:"name"`
					} `json:"client_auth"`
					AuthorizedPublicKeyList []struct {
						Name string `json:"name"`
						Key  string `json:"key"`
					} `json:"authorized_public_key_list"`
					SoftwareMap struct {
						AdditionalProp1 struct {
							Status       string `json:"status"`
							Version      string `json:"version"`
							SoftwareType string `json:"software_type"`
						} `json:"additionalProp1"`
						AdditionalProp2 struct {
							Status       string `json:"status"`
							Version      string `json:"version"`
							SoftwareType string `json:"software_type"`
						} `json:"additionalProp2"`
						AdditionalProp3 struct {
							Status       string `json:"status"`
							Version      string `json:"version"`
							SoftwareType string `json:"software_type"`
						} `json:"additionalProp3"`
					} `json:"software_map"`
					EncryptionStatus string `json:"encryption_status"`
					SslKey           struct {
						KeyType     string `json:"key_type"`
						KeyName     string `json:"key_name"`
						SigningInfo struct {
							City             string `json:"city"`
							CommonNameSuffix string `json:"common_name_suffix"`
							State            string `json:"state"`
							CountryCode      string `json:"country_code"`
							CommonName       string `json:"common_name"`
							Organization     string `json:"organization"`
							EmailAddress     string `json:"email_address"`
						} `json:"signing_info"`
						ExpireDatetime time.Time `json:"expire_datetime"`
					} `json:"ssl_key"`
					ServiceList                   []string `json:"service_list"`
					SupportedInformationVerbosity string   `json:"supported_information_verbosity"`
					CertificationSigningInfo      struct {
						City             string `json:"city"`
						CommonNameSuffix string `json:"common_name_suffix"`
						State            string `json:"state"`
						CountryCode      string `json:"country_code"`
						CommonName       string `json:"common_name"`
						Organization     string `json:"organization"`
						EmailAddress     string `json:"email_address"`
					} `json:"certification_signing_info"`
					RedundancyFactor       int `json:"redundancy_factor"`
					ExternalConfigurations struct {
						CitrixConnectorConfig struct {
							CitrixVMReferenceList []struct {
								Kind string `json:"kind"`
								Name string `json:"name"`
								UUID string `json:"uuid"`
							} `json:"citrix_vm_reference_list"`
							ClientSecret     string `json:"client_secret"`
							CustomerID       string `json:"customer_id"`
							ClientID         string `json:"client_id"`
							ResourceLocation struct {
								ID   string `json:"id"`
								Name string `json:"name"`
							} `json:"resource_location"`
						} `json:"citrix_connector_config"`
					} `json:"external_configurations"`
					OperationMode     string `json:"operation_mode"`
					CaCertificateList []struct {
						CaName      string `json:"ca_name"`
						Certificate string `json:"certificate"`
					} `json:"ca_certificate_list"`
					DomainAwarenessLevel string   `json:"domain_awareness_level"`
					EnabledFeatureList   []string `json:"enabled_feature_list"`
					IsAvailable          bool     `json:"is_available"`
					Build                struct {
						CommitID          string    `json:"commit_id"`
						FullVersion       string    `json:"full_version"`
						CommitDate        time.Time `json:"commit_date"`
						IsLongTermSupport bool      `json:"is_long_term_support"`
						Version           string    `json:"version"`
						ShortCommitID     string    `json:"short_commit_id"`
						BuildType         string    `json:"build_type"`
					} `json:"build"`
					Timezone                  string `json:"timezone"`
					EnableEfficientMetricSync bool   `json:"enable_efficient_metric_sync"`
					ClusterArch               string `json:"cluster_arch"`
					ManagementServerList      []struct {
						IP         string   `json:"ip"`
						DrsEnabled bool     `json:"drs_enabled"`
						StatusList []string `json:"status_list"`
						Type       string   `json:"type"`
					} `json:"management_server_list"`
				} `json:"config"`
				Network struct {
					MasqueradingIP   string `json:"masquerading_ip"`
					MasqueradingPort int    `json:"masquerading_port"`
					ExternalIP       string `json:"external_ip"`
					HTTPProxyList    []struct {
						Credentials struct {
							Username string `json:"username"`
							Password string `json:"password"`
						} `json:"credentials"`
						Name          string   `json:"name"`
						ProxyTypeList []string `json:"proxy_type_list"`
						Address       struct {
							IP       string `json:"ip"`
							IsBackup bool   `json:"is_backup"`
							Fqdn     string `json:"fqdn"`
							IP6Range string `json:"ip6_range"`
							IPRange  string `json:"ip_range"`
							Ipv6     string `json:"ipv6"`
							Port     int    `json:"port"`
						} `json:"address"`
					} `json:"http_proxy_list"`
					SMTPServer struct {
						Type         string `json:"type"`
						EmailAddress string `json:"email_address"`
						Server       struct {
							Credentials struct {
								Username string `json:"username"`
								Password string `json:"password"`
							} `json:"credentials"`
							Name          string   `json:"name"`
							ProxyTypeList []string `json:"proxy_type_list"`
							Address       struct {
								IP       string `json:"ip"`
								IsBackup bool   `json:"is_backup"`
								Fqdn     string `json:"fqdn"`
								IP6Range string `json:"ip6_range"`
								IPRange  string `json:"ip_range"`
								Ipv6     string `json:"ipv6"`
								Port     int    `json:"port"`
							} `json:"address"`
						} `json:"server"`
					} `json:"smtp_server"`
					NtpServerIPList        []string `json:"ntp_server_ip_list"`
					ExternalSubnet         string   `json:"external_subnet"`
					NfsSubnetWhitelist     []string `json:"nfs_subnet_whitelist"`
					ExternalDataServicesIP string   `json:"external_data_services_ip"`
					DomainServer           struct {
						Nameserver        string `json:"nameserver"`
						Name              string `json:"name"`
						DomainCredentials struct {
							Username string `json:"username"`
							Password string `json:"password"`
						} `json:"domain_credentials"`
					} `json:"domain_server"`
					FullyQualifiedDomainName string   `json:"fully_qualified_domain_name"`
					NameServerIPList         []string `json:"name_server_ip_list"`
					HTTPProxyWhitelist       []struct {
						Target     string `json:"target"`
						TargetType string `json:"target_type"`
					} `json:"http_proxy_whitelist"`
					InternalSubnet       string `json:"internal_subnet"`
					DefaultVswitchConfig struct {
						NicTeamingPolicy string `json:"nic_teaming_policy"`
						UplinkGrouping   string `json:"uplink_grouping"`
					} `json:"default_vswitch_config"`
				} `json:"network"`
				Analysis struct {
					VMEfficiencyMap struct {
						AdditionalProp1 string `json:"additionalProp1"`
						AdditionalProp2 string `json:"additionalProp2"`
						AdditionalProp3 string `json:"additionalProp3"`
					} `json:"vm_efficiency_map"`
				} `json:"analysis"`
				RuntimeStatusList []string `json:"runtime_status_list"`
			} `json:"resources"`
		} `json:"status"`
		Spec struct {
			Name      string `json:"name"`
			Resources struct {
				Config struct {
					GpuDriverVersion string `json:"gpu_driver_version"`
					ClientAuth       struct {
						Status  string `json:"status"`
						CaChain string `json:"ca_chain"`
						Name    string `json:"name"`
					} `json:"client_auth"`
					AuthorizedPublicKeyList []struct {
						Name string `json:"name"`
						Key  string `json:"key"`
					} `json:"authorized_public_key_list"`
					SoftwareMap struct {
						AdditionalProp1 struct {
							Status       string `json:"status"`
							Version      string `json:"version"`
							SoftwareType string `json:"software_type"`
						} `json:"additionalProp1"`
						AdditionalProp2 struct {
							Status       string `json:"status"`
							Version      string `json:"version"`
							SoftwareType string `json:"software_type"`
						} `json:"additionalProp2"`
						AdditionalProp3 struct {
							Status       string `json:"status"`
							Version      string `json:"version"`
							SoftwareType string `json:"software_type"`
						} `json:"additionalProp3"`
					} `json:"software_map"`
					EncryptionStatus         string `json:"encryption_status"`
					RedundancyFactor         int    `json:"redundancy_factor"`
					CertificationSigningInfo struct {
						City             string `json:"city"`
						CommonNameSuffix string `json:"common_name_suffix"`
						State            string `json:"state"`
						CountryCode      string `json:"country_code"`
						CommonName       string `json:"common_name"`
						Organization     string `json:"organization"`
						EmailAddress     string `json:"email_address"`
					} `json:"certification_signing_info"`
					SupportedInformationVerbosity string `json:"supported_information_verbosity"`
					ExternalConfigurations        struct {
						CitrixConnectorConfig struct {
							CitrixVMReferenceList []struct {
								Kind string `json:"kind"`
								Name string `json:"name"`
								UUID string `json:"uuid"`
							} `json:"citrix_vm_reference_list"`
							ClientSecret     string `json:"client_secret"`
							CustomerID       string `json:"customer_id"`
							ClientID         string `json:"client_id"`
							ResourceLocation struct {
								ID   string `json:"id"`
								Name string `json:"name"`
							} `json:"resource_location"`
						} `json:"citrix_connector_config"`
					} `json:"external_configurations"`
					DomainAwarenessLevel      string   `json:"domain_awareness_level"`
					EnabledFeatureList        []string `json:"enabled_feature_list"`
					Timezone                  string   `json:"timezone"`
					EnableEfficientMetricSync bool     `json:"enable_efficient_metric_sync"`
					OperationMode             string   `json:"operation_mode"`
				} `json:"config"`
				Network struct {
					MasqueradingIP   string `json:"masquerading_ip"`
					MasqueradingPort int    `json:"masquerading_port"`
					ExternalIP       string `json:"external_ip"`
					HTTPProxyList    []struct {
						Credentials struct {
							Username string `json:"username"`
							Password string `json:"password"`
						} `json:"credentials"`
						Name          string   `json:"name"`
						ProxyTypeList []string `json:"proxy_type_list"`
						Address       struct {
							IP       string `json:"ip"`
							IsBackup bool   `json:"is_backup"`
							Fqdn     string `json:"fqdn"`
							IP6Range string `json:"ip6_range"`
							IPRange  string `json:"ip_range"`
							Ipv6     string `json:"ipv6"`
							Port     int    `json:"port"`
						} `json:"address"`
					} `json:"http_proxy_list"`
					SMTPServer struct {
						Type         string `json:"type"`
						EmailAddress string `json:"email_address"`
						Server       struct {
							Credentials struct {
								Username string `json:"username"`
								Password string `json:"password"`
							} `json:"credentials"`
							Name          string   `json:"name"`
							ProxyTypeList []string `json:"proxy_type_list"`
							Address       struct {
								IP       string `json:"ip"`
								IsBackup bool   `json:"is_backup"`
								Fqdn     string `json:"fqdn"`
								IP6Range string `json:"ip6_range"`
								IPRange  string `json:"ip_range"`
								Ipv6     string `json:"ipv6"`
								Port     int    `json:"port"`
							} `json:"address"`
						} `json:"server"`
					} `json:"smtp_server"`
					NtpServerIPList        []string `json:"ntp_server_ip_list"`
					ExternalSubnet         string   `json:"external_subnet"`
					NfsSubnetWhitelist     []string `json:"nfs_subnet_whitelist"`
					ExternalDataServicesIP string   `json:"external_data_services_ip"`
					DomainServer           struct {
						Nameserver        string `json:"nameserver"`
						Name              string `json:"name"`
						DomainCredentials struct {
							Username string `json:"username"`
							Password string `json:"password"`
						} `json:"domain_credentials"`
					} `json:"domain_server"`
					FullyQualifiedDomainName string   `json:"fully_qualified_domain_name"`
					NameServerIPList         []string `json:"name_server_ip_list"`
					HTTPProxyWhitelist       []struct {
						Target     string `json:"target"`
						TargetType string `json:"target_type"`
					} `json:"http_proxy_whitelist"`
					InternalSubnet       string `json:"internal_subnet"`
					DefaultVswitchConfig struct {
						NicTeamingPolicy string `json:"nic_teaming_policy"`
						UplinkGrouping   string `json:"uplink_grouping"`
					} `json:"default_vswitch_config"`
				} `json:"network"`
				RuntimeStatusList []string `json:"runtime_status_list"`
			} `json:"resources"`
		} `json:"spec"`
		APIVersion string `json:"api_version"`
		Metadata   struct {
			LastUpdateTime       time.Time `json:"last_update_time"`
			UseCategoriesMapping bool      `json:"use_categories_mapping"`
			Kind                 string    `json:"kind"`
			UUID                 string    `json:"uuid"`
			ProjectReference     struct {
				Kind string `json:"kind"`
				Name string `json:"name"`
				UUID string `json:"uuid"`
			} `json:"project_reference"`
			CreationTime      time.Time `json:"creation_time"`
			SpecVersion       int       `json:"spec_version"`
			SpecHash          string    `json:"spec_hash"`
			CategoriesMapping struct {
				AdditionalProp1 []string `json:"additionalProp1"`
				AdditionalProp2 []string `json:"additionalProp2"`
				AdditionalProp3 []string `json:"additionalProp3"`
			} `json:"categories_mapping"`
			ShouldForceTranslate bool   `json:"should_force_translate"`
			EntityVersion        string `json:"entity_version"`
			OwnerReference       struct {
				Kind string `json:"kind"`
				Name string `json:"name"`
				UUID string `json:"uuid"`
			} `json:"owner_reference"`
			Categories struct {
				AdditionalProp1 string `json:"additionalProp1"`
				AdditionalProp2 string `json:"additionalProp2"`
				AdditionalProp3 string `json:"additionalProp3"`
			} `json:"categories"`
			Name string `json:"name"`
		} `json:"metadata"`
	} `json:"entities"`
	APIVersion string `json:"api_version"`
	Metadata   struct {
		Kind          string `json:"kind"`
		TotalMatches  int    `json:"total_matches"`
		SortAttribute string `json:"sort_attribute"`
		Filter        string `json:"filter"`
		Length        int    `json:"length"`
		SortOrder     string `json:"sort_order"`
		Offset        int    `json:"offset"`
	} `json:"metadata"`
}
