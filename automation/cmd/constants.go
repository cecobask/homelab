package cmd

const (
	CommandAliasRoot = "auto"
	CommandNameRoot  = "automation"

	BaseURLProxmox                   = "https://192.168.0.69:8006/api2/json"
	CommandAliasDeleteVolume         = "dv"
	CommandAliasDestroyVM            = "dvm"
	CommandAliasProxmox              = "pm"
	CommandAliasShutdownVM           = "sdvm"
	CommandAliasStopVM               = "svm"
	CommandNameDeleteVolume          = "delete-volume"
	CommandNameDestroyVM             = "destroy-vm"
	CommandNameProxmox               = "proxmox"
	CommandNameShutdownVM            = "shutdown-vm"
	CommandNameStopVM                = "stop-vm"
	FlagNameDestroyUnreferencedDisks = "destroy-unreferenced-disks"
	FlagNameNode                     = "node"
	FlagNamePurge                    = "purge"
	FlagNameStorage                  = "storage"
	FlagNameVMID                     = "vmid"
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
