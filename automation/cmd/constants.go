package cmd

const (
	CommandAliasRoot = "auto"
	CommandNameRoot  = "automation"

	BaseURLProxmox                   = "https://192.168.0.69:8006/api2/json"
	CommandAliasDeleteVolume         = "dv"
	CommandAliasDestroyVMs           = "dvm"
	CommandAliasProxmox              = "pm"
	CommandAliasShutdownVMs          = "sdvm"
	CommandAliasStopVMs              = "svm"
	CommandNameDeleteVolume          = "delete-volume"
	CommandNameDestroyVMs            = "destroy-vms"
	CommandNameProxmox               = "proxmox"
	CommandNameShutdownVMs           = "shutdown-vms"
	CommandNameStopVMs               = "stop-vms"
	FlagNameDestroyUnreferencedDisks = "destroy-unreferenced-disks"
	FlagNameNode                     = "node"
	FlagNamePurge                    = "purge"
	FlagNameStorage                  = "storage"
	FlagNameVMIDs                    = "vmids"
	FlagNameVolume                   = "volume"

	BaseURLTailscale          = "https://api.tailscale.com/api/v2"
	CommandAliasDeleteDevices = "dd"
	CommandAliasListDevices   = "ld"
	CommandAliasTailscale     = "ts"
	CommandNameDeleteDevices  = "delete-devices"
	CommandNameListDevices    = "list-devices"
	CommandNameTailscale      = "tailscale"
	FlagNameHostnames         = "hostnames"
	FlagNameIDs               = "ids"
	FlagNameTags              = "tags"
	FlagNameTailnetName       = "tailnet-name"

	FlagNameBaseURL = "base-url"
)
