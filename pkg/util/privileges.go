package util

import (
	"context"
	"errors"
	"fmt"
	permissions "github.com/rvanderp/vsphere-perm-check/pkg/data"
	pctypes "github.com/rvanderp/vsphere-perm-check/pkg/types"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
	"strings"
)

func ComparePrivileges(derived []types.UserPrivilegeResult, required []string) error {
	var missingPrivileges = ""
	for _, neededPrivilege := range required {
		var hasPrivilege = false
		for _, userPrivilege := range derived {
			for _, assignedPrivilege := range userPrivilege.Privileges {
				if assignedPrivilege == neededPrivilege {
					hasPrivilege = true
				}
			}
		}
		if hasPrivilege == false {
			if missingPrivileges != "" {
				missingPrivileges = missingPrivileges + ", "
			}
			missingPrivileges = missingPrivileges + neededPrivilege
		}
	}
	if missingPrivileges != "" {
		return errors.New(missingPrivileges)
	}
	return nil
}

func checkPropagationRequirement(ctx context.Context, authManager *object.AuthorizationManager, reference types.ManagedObjectReference, userName string) error {
	userNameParts := strings.Split(userName, "@")
	if len(userNameParts) < 2 {
		return errors.New("username must contain user@domain")
	}
	var errorList = ""
	principal := strings.ToUpper(userNameParts[1]) + "\\" + strings.ToLower(userNameParts[0])

	roleList, err := authManager.RetrieveAllPermissions(ctx)
	if err != nil {
		return err
	}
	for _, permission := range roleList {
		if permission.Entity.Reference() == reference {
			if permission.Principal == principal {
				if permission.Propagate == false {
					errorList = errorList + "User: " + permission.Principal + " is not configured to propagate to children\n"
				}
			}
		}
	}

	if errorList != "" {
		return errors.New(errorList)
	}
	return nil
}

func ValidatePrivileges(ssn *Session, p *pctypes.Platform, folder string) error {
	ctx := context.TODO()
	var invalidPrivilegeConfiguration = ""

	authManager := object.NewAuthorizationManager(ssn.Vim25Client)

	finder := find.NewFinder(ssn.Vim25Client)

	if val, ok := permissions.RequiredPermissions["Datacenter"]; ok {
		datacenter, err := finder.Datacenter(ctx, p.Datacenter)
		if err != nil {
			return err
		}
		res, err := authManager.FetchUserPrivilegeOnEntities(ctx, []types.ManagedObjectReference{datacenter.Reference()}, p.Username)
		if err != nil {
			return err
		}
		err = ComparePrivileges(res, val.Privileges)
		if err != nil {
			invalidPrivilegeConfiguration = invalidPrivilegeConfiguration + fmt.Sprintf("*** Missing Privileges ***\nvSphere object: %s\n%s\n\n", val.Name, err.Error())
		}
		err = checkPropagationRequirement(ctx, authManager, datacenter.Reference(), p.Username)
		if err != nil {
			invalidPrivilegeConfiguration = invalidPrivilegeConfiguration + fmt.Sprintf("*** Invalid Propagation Configuration ***\n"+
				"Note: Propagation only required if folder is created by openshift-install.\n"+
				"vSphere object: %s\n%s\n\n", val.Name, err.Error())
		}
	}

	if val, ok := permissions.RequiredPermissions["Datastore"]; ok {
		if p.DefaultDatastore == "" {
			ds, err := finder.DefaultDatastore(ctx)
			if err != nil {
				return err
			}
			p.DefaultDatastore = ds.Name()
		}
		datastore, err := finder.Datastore(ctx, p.DefaultDatastore)
		if err != nil {
			return err
		}
		res, err := authManager.FetchUserPrivilegeOnEntities(ctx, []types.ManagedObjectReference{datastore.Reference()}, p.Username)
		if err != nil {
			return err
		}
		err = ComparePrivileges(res, val.Privileges)
		if err != nil {
			invalidPrivilegeConfiguration = invalidPrivilegeConfiguration + fmt.Sprintf("*** Missing Privileges ***\nvSphere object: %s\n%s\n\n", val.Name, err.Error())
		}
	}

	if val, ok := permissions.RequiredPermissions["Port group"]; ok {
		network, err := finder.Network(ctx, p.Network)
		if err != nil {
			return err
		}
		res, err := authManager.FetchUserPrivilegeOnEntities(ctx, []types.ManagedObjectReference{network.Reference()}, p.Username)
		if err != nil {
			return err
		}
		err = ComparePrivileges(res, val.Privileges)
		if err != nil {
			invalidPrivilegeConfiguration = invalidPrivilegeConfiguration + fmt.Sprintf("*** Missing Privileges ***\nvSphere object: %s\n%s\n\n", val.Name, err.Error())
		}
	}

	if val, ok := permissions.RequiredPermissions["Cluster"]; ok {
		cluster, err := finder.ClusterComputeResource(ctx, p.Cluster)
		if err != nil {
			return err
		}
		res, err := authManager.FetchUserPrivilegeOnEntities(ctx, []types.ManagedObjectReference{cluster.Reference()}, p.Username)
		if err != nil {
			return err
		}
		err = ComparePrivileges(res, val.Privileges)
		if err != nil {
			invalidPrivilegeConfiguration = invalidPrivilegeConfiguration + fmt.Sprintf("*** Missing Privileges ***\nvSphere object: %s\n%s\n\n", val.Name, err.Error())
		}
		err = checkPropagationRequirement(ctx, authManager, cluster.Reference(), p.Username)
		if err != nil {
			invalidPrivilegeConfiguration = invalidPrivilegeConfiguration + fmt.Sprintf("*** Invalid Propagation Configuration ***\nvSphere object: %s\n%s\n\n", val.Name, err.Error())
		}
	}
	if val, ok := permissions.RequiredPermissions["vCenter"]; ok {
		rootFolder := object.NewRootFolder(ssn.Vim25Client)
		res, err := authManager.FetchUserPrivilegeOnEntities(ctx, []types.ManagedObjectReference{rootFolder.Reference()}, p.Username)
		if err != nil {
			return err
		}
		err = ComparePrivileges(res, val.Privileges)
		if err != nil {
			invalidPrivilegeConfiguration = invalidPrivilegeConfiguration + fmt.Sprintf("*** Missing Privileges ***\nvSphere object: %s\n%s\n\n", val.Name, err.Error())
		}
	}
	if folder != "" {
		if val, ok := permissions.RequiredPermissions["Folder"]; ok {
			folderObj, err := finder.Folder(ctx, folder)
			if err != nil {
				return err
			}
			res, err := authManager.FetchUserPrivilegeOnEntities(ctx, []types.ManagedObjectReference{folderObj.Reference()}, p.Username)
			if err != nil {
				return err
			}
			err = ComparePrivileges(res, val.Privileges)
			if err != nil {
				invalidPrivilegeConfiguration = invalidPrivilegeConfiguration + fmt.Sprintf("*** Missing Privileges ***\nvSphere object: %s\n%s\n\n", val.Name, err.Error())
			}
			err = checkPropagationRequirement(ctx, authManager, folderObj.Reference(), p.Username)
			if err != nil {
				invalidPrivilegeConfiguration = invalidPrivilegeConfiguration + fmt.Sprintf("*** Invalid Propagation Configuration ***\nvSphere object: %s\n%s\n\n", val.Name, err.Error())
			}
		}
	}
	if invalidPrivilegeConfiguration != "" {
		return errors.New(invalidPrivilegeConfiguration)
	}
	return nil
}
