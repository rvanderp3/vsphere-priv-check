# vsphere-priv-check

## Overview

Validating privileges for a user can be challenging.  This project attempts to provide a quick method for validating that the user account assigned for installation of an OpenShift cluster has the required privileges.  This tool forms the basis of required privileges from [Required vCenter account privileges](https://docs.openshift.com/container-platform/latest/installing/installing_vsphere/installing-vsphere-installer-provisioned.html#installation-vsphere-installer-infra-requirements-account_installing-vsphere-installer-provisioned)

## Building

1. Setup `go` environment
2. Build the binary
~~~
./hack/build.sh
~~~

## Usage

This tool requires that an account with administrator privileges be provided.  This account will verify the privileges of the account defined in `install-config.yaml`.  To define the administrator account:

~~~
export VCENTER_USERNAME=admin@your.domain
export VCENTER_PASSWORD=yourpassword
~~~

The `install-config.yaml` to be used for the installation must be present in the working directory of this tool.  Information such as target datastore, data center, username, and network are all derived from this file.  

To run the tool:

~~~
$ ./bin/vsphere-priv-check
OpenShift vSphere Pre-Flight Permissions Validator

2021/08/10 13:39:41 checking permissions for user test@vsphere.local

2021/08/10 13:39:41 error while validating required privileges:

*** Missing Privileges ***
vSphere object: vSphere vCenter Datacenter
Resource.AssignVMToPool, VApp.Import, VirtualMachine.Config.AddExistingDisk, VirtualMachine.Config.AddNewDisk, VirtualMachine.Config.AddRemoveDevice, VirtualMachine.Config.AdvancedConfig, VirtualMachine.Config.Annotation, VirtualMachine.Config.CPUCount, VirtualMachine.Config.DiskExtend, VirtualMachine.Config.DiskLease, VirtualMachine.Config.EditDevice, VirtualMachine.Config.Memory, VirtualMachine.Config.RemoveDisk, VirtualMachine.Config.Rename, VirtualMachine.Config.ResetGuestInfo, VirtualMachine.Config.Resource, VirtualMachine.Config.Settings, VirtualMachine.Config.UpgradeVirtualHardware, VirtualMachine.Interact.GuestControl, VirtualMachine.Interact.PowerOff, VirtualMachine.Interact.PowerOn, VirtualMachine.Interact.Reset, VirtualMachine.Inventory.Create, VirtualMachine.Inventory.CreateFromExisting, VirtualMachine.Inventory.Delete, VirtualMachine.Provisioning.Clone, Folder.Create, Folder.Delete
~~~

## Missing Checks

- Privilege Propagation
- Virtual Machine Folder
