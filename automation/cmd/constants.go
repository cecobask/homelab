package cmd

const (
	CommandAliasRoot = "auto"
	CommandNameRoot  = "automation"

	BaseURLProxmox                   = "https://192.168.0.69:8006/api2/json"
	CommandAliasDeleteVolumes        = "dv"
	CommandAliasDestroyVMs           = "dvm"
	CommandAliasListVMs              = "lvm"
	CommandAliasProxmox              = "pm"
	CommandAliasStopVMs              = "svm"
	CommandNameDeleteVolumes         = "delete-volumes"
	CommandNameDestroyVMs            = "destroy-vms"
	CommandNameListVMs               = "list-vms"
	CommandNameProxmox               = "proxmox"
	CommandNameStopVMs               = "stop-vms"
	FlagNameContentTypes             = "content-types"
	FlagNameDestroyUnreferencedDisks = "destroy-unreferenced-disks"
	FlagNameNodes                    = "nodes"
	FlagNamePurge                    = "purge"
	FlagNameStorages                 = "storages"
	FlagNameVolumes                  = "volumes"

	BaseURLTailscale          = "https://api.tailscale.com/api/v2"
	CommandAliasDeleteDevices = "dd"
	CommandAliasListDevices   = "ld"
	CommandAliasTailscale     = "ts"
	CommandNameDeleteDevices  = "delete-devices"
	CommandNameListDevices    = "list-devices"
	CommandNameTailscale      = "tailscale"
	FlagNameHostnames         = "hostnames"
	FlagNameTailnetName       = "tailnet-name"

	FlagNameBaseURL = "base-url"
	FlagNameIDs     = "ids"
	FlagNameTags    = "tags"
)
