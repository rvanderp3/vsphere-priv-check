package util

import (
	"context"
	"errors"
	"github.com/davecgh/go-spew/spew"
	pctypes "github.com/rvanderp/vsphere-perm-check/pkg/types"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/ssoadmin"
	"github.com/vmware/govmomi/sts"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	"net/url"
	"os"
)

type Session struct {
	Vim25Client *vim25.Client
	RestClient  *rest.Client
	SsoClient   *ssoadmin.Client
}

// GetSession - creates the Session struct
func GetSession(ctx context.Context, p *pctypes.Platform) (*Session, error) {
	session := &Session{}
	header := soap.Header{
		Security: &sts.Signer{},
	}

	username := os.Getenv("VCENTER_USERNAME")
	if username == "" {
		return nil, errors.New("username with administrator privileges must be provided via `export VCENTER_USERNAME=admin@your.domain`")
	}
	password := os.Getenv("VCENTER_PASSWORD")
	if password == "" {
		return nil, errors.New("password for `VCENTER_USERNAME` must be provided via `export VCENTER_PASSWORD=yourpassword`")
	}
	u, err := soap.ParseURL(p.VCenter)
	if err != nil {
		return nil, err
	}
	u.User = url.UserPassword(username, password)

	if u.User == nil {
		spew.Dump("user nil")
	}

	// false in this method disables insecure
	// We do not allow insecure connections
	govmomiClient, err := govmomi.NewClient(ctx, u, true)
	vimClient, err := vim25.NewClient(ctx, soap.NewClient(u, true))
	stsClient, cerr := sts.NewClient(ctx, vimClient)

	if cerr != nil {
		return nil, cerr
	}
	ssoClient, err := ssoadmin.NewClient(ctx, vimClient)
	if err != nil {
		return nil, err
	}

	req := sts.TokenRequest{
		Certificate: vimClient.Certificate(),
		Userinfo:    u.User,
	}

	header.Security, cerr = stsClient.Issue(ctx, req)
	if cerr != nil {
		return nil, cerr
	}

	if err = ssoClient.Login(ssoClient.WithHeader(ctx, header)); err != nil {
		return nil, err
	}

	session.SsoClient = ssoClient
	session.Vim25Client = govmomiClient.Client

	restClient := rest.NewClient(vimClient)
	err = restClient.Login(ctx, u.User)
	if err != nil {
		return nil, err
	}
	session.RestClient = restClient
	return session, nil
}
