package data

var (
	RequiredVcenterPrivileges []string = []string{
		"Cns.Searchable",
		"InventoryService.Tagging.AttachTag",
		"InventoryService.Tagging.CreateCategory",
		"InventoryService.Tagging.CreateTag",
		"InventoryService.Tagging.DeleteCategory",
		"InventoryService.Tagging.DeleteTag",
		"InventoryService.Tagging.EditCategory",
		"InventoryService.Tagging.EditTag",
		"Sessions.ValidateSession",
		"StorageProfile.View",
	}
	RequiredVcenterClusterPrivileges []string = []string{
		"Host.Config.Storage",
		"Resource.AssignVMToPool",
		"VApp.AssignResourcePool",
		"VApp.Import",
		"VirtualMachine.Config.AddNewDisk",
	}
	RequiredDatastorePrivileges []string = []string{
		"Datastore.AllocateSpace",
		"Datastore.Browse",
		"Datastore.FileManagement",
	}
	RequiredPortgroupPrivileges []string = []string{
		"Network.Assign",
	}
	RequiredFolderPrivileges []string = []string{
		"Resource.AssignVMToPool",
		"VApp.Import",
		"VirtualMachine.Config.AddExistingDisk",
		"VirtualMachine.Config.AddNewDisk",
		"VirtualMachine.Config.AddRemoveDevice",
		"VirtualMachine.Config.AdvancedConfig",
		"VirtualMachine.Config.Annotation",
		"VirtualMachine.Config.CPUCount",
		"VirtualMachine.Config.DiskExtend",
		"VirtualMachine.Config.DiskLease",
		"VirtualMachine.Config.EditDevice",
		"VirtualMachine.Config.Memory",
		"VirtualMachine.Config.RemoveDisk",
		"VirtualMachine.Config.Rename",
		"VirtualMachine.Config.ResetGuestInfo",
		"VirtualMachine.Config.Resource",
		"VirtualMachine.Config.Settings",
		"VirtualMachine.Config.UpgradeVirtualHardware",
		"VirtualMachine.Interact.GuestControl",
		"VirtualMachine.Interact.PowerOff",
		"VirtualMachine.Interact.PowerOn",
		"VirtualMachine.Interact.Reset",
		"VirtualMachine.Inventory.Create",
		"VirtualMachine.Inventory.CreateFromExisting",
		"VirtualMachine.Inventory.Delete",
		"VirtualMachine.Provisioning.Clone",
	}
	RequiredDatacenterPrivileges []string = []string{
		"Resource.AssignVMToPool",
		"VApp.Import",
		"VirtualMachine.Config.AddExistingDisk",
		"VirtualMachine.Config.AddNewDisk",
		"VirtualMachine.Config.AddRemoveDevice",
		"VirtualMachine.Config.AdvancedConfig",
		"VirtualMachine.Config.Annotation",
		"VirtualMachine.Config.CPUCount",
		"VirtualMachine.Config.DiskExtend",
		"VirtualMachine.Config.DiskLease",
		"VirtualMachine.Config.EditDevice",
		"VirtualMachine.Config.Memory",
		"VirtualMachine.Config.RemoveDisk",
		"VirtualMachine.Config.Rename",
		"VirtualMachine.Config.ResetGuestInfo",
		"VirtualMachine.Config.Resource",
		"VirtualMachine.Config.Settings",
		"VirtualMachine.Config.UpgradeVirtualHardware",
		"VirtualMachine.Interact.GuestControl",
		"VirtualMachine.Interact.PowerOff",
		"VirtualMachine.Interact.PowerOn",
		"VirtualMachine.Interact.Reset",
		"VirtualMachine.Inventory.Create",
		"VirtualMachine.Inventory.CreateFromExisting",
		"VirtualMachine.Inventory.Delete",
		"VirtualMachine.Provisioning.Clone",
		"VirtualMachine.Provisioning.DeployTemplate",
		"VirtualMachine.Provisioning.MarkAsTemplate",
		"Folder.Create",
		"Folder.Delete",
	}

	RequiredPermissions = map[string]PrivilegeDesc{
		"Datacenter": {
			Privileges: RequiredDatacenterPrivileges,
			Name:       "vSphere vCenter Datacenter",
			Propagates: true,
		},
		"Folder": {
			Privileges: RequiredFolderPrivileges,
			Name:       "Virtual Machine Folder",
			Propagates: false,
		},
		"Datastore": {
			Privileges: RequiredDatastorePrivileges,
			Name:       "vSphere Datastore",
			Propagates: false,
		},
		"Port group": {
			Privileges: RequiredPortgroupPrivileges,
			Name:       "vSphere Port Group",
			Propagates: false,
		},
		"vCenter": {
			Privileges: RequiredVcenterPrivileges,
			Name:       "vSphere vCenter",
			Propagates: false,
		},
		"Cluster": {
			Privileges: RequiredVcenterClusterPrivileges,
			Name:       "vSphere vCenter Cluster",
			Propagates: false,
		},
	}
)
