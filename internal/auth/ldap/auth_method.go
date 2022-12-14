package ldap

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/boundary/internal/auth/ldap/store"
	"github.com/hashicorp/boundary/internal/errors"
	"github.com/hashicorp/boundary/internal/oplog"
	"google.golang.org/protobuf/proto"
)

// authMethodTableName defines an AuthMethod's table name.
const authMethodTableName = "auth_ldap_method"

// AuthMethod contains an LDAP auth method configuration.  It is owned by a
// scope. AuthMethods MUST have at least one Url. AuthMethods MAY one or zero:
// UserEntrySearchConf, a GroupEntrySearchConf, BindCredential. AuthMethods
// may have zero to many: Accounts, Certificates,
type AuthMethod struct {
	*store.AuthMethod
	tableName string
}

// NewAuthMethod creates a new in memory AuthMethod assigned to a scopeId.  The
// new auth method will have an OperationalState of Inactive.
//
// Supports the options: WithName, WithDescription, WithStartTLS,
// WithInsecureTLS, WithDiscoverDN, WithAnonGroupSearch, WithUpnDomain,
// WithUserSearchConf, WithGroupSearchConf, WithCertificates, WithBindCredential
// are the only valid options and all other options are ignored.
func NewAuthMethod(ctx context.Context, scopeId string, urls []*url.URL, opt ...Option) (*AuthMethod, error) {
	const op = "ldap.NewAuthMethod"
	switch {
	case scopeId == "":
		return nil, errors.New(ctx, errors.InvalidParameter, op, "missing scope id")
	case len(urls) == 0:
		return nil, errors.New(ctx, errors.InvalidParameter, op, "missing urls (must have at least one URL)")
	}
	opts, err := getOpts(opt...)
	if err != nil {
		return nil, errors.Wrap(ctx, err, op)
	}

	strUrls := make([]string, 0, len(urls))
	for _, u := range urls {
		switch u.Scheme {
		case "ldap", "ldaps":
		default:
			return nil, errors.New(ctx, errors.InvalidParameter, op, fmt.Sprintf("%s scheme in url %q is not either ldap or ldaps", u.Scheme, u.String()))
		}
		strUrls = append(strUrls, u.String())
	}
	a := &AuthMethod{
		AuthMethod: &store.AuthMethod{
			ScopeId:              scopeId,
			Name:                 opts.withName,
			Description:          opts.withDescription,
			OperationalState:     string(InactiveState), // all new auth methods are initially inactive
			Urls:                 strUrls,
			StartTls:             opts.withStartTls,
			InsecureTls:          opts.withInsecureTls,
			DiscoverDn:           opts.withDiscoverDn,
			AnonGroupSearch:      opts.withAnonGroupSearch,
			UpnDomain:            opts.withUpnDomain,
			UserDn:               opts.withUserDn,
			UserAttr:             opts.withUserAttr,
			UserFilter:           opts.withUserFilter,
			GroupDn:              opts.withGroupDn,
			GroupAttr:            opts.withGroupAttr,
			GroupFilter:          opts.withGroupFilter,
			BindDn:               opts.withBindDn,
			BindPassword:         opts.withBindPassword,
			Certificates:         opts.withCertificates,
			ClientCertificate:    opts.withClientCertificate,
			ClientCertificateKey: opts.withClientCertificateKey,
		},
	}

	return a, nil
}

// allocAuthMethod makes an empty one in memory
func allocAuthMethod() AuthMethod {
	return AuthMethod{
		AuthMethod: &store.AuthMethod{},
	}
}

// clone an AuthMethod
func (am *AuthMethod) clone() *AuthMethod {
	cp := proto.Clone(am.AuthMethod)
	return &AuthMethod{
		AuthMethod: cp.(*store.AuthMethod),
	}
}

// TableName returns the table name (func is required by gorm)
func (am *AuthMethod) TableName() string {
	if am.tableName != "" {
		return am.tableName
	}
	return authMethodTableName
}

// SetTableName sets the table name (func is required by oplog)
func (am *AuthMethod) SetTableName(n string) {
	am.tableName = n
}

// oplog will create oplog metadata for the AuthMethod.
func (am *AuthMethod) oplog(op oplog.OpType) oplog.Metadata {
	metadata := oplog.Metadata{
		"resource-public-id": []string{am.GetPublicId()},
		"resource-type":      []string{"ldap auth method"},
		"op-type":            []string{op.String()},
		"scope-id":           []string{am.ScopeId},
	}
	return metadata
}

type convertedValues struct {
	Urls                 []any
	Certs                []any
	UserEntrySearchConf  any
	GroupEntrySearchConf any
	ClientCertificate    any
	BindCredential       any
}

// convertValueObjects converts the embedded value objects. It will return an
// error if the AuthMethod's public id is not set.
func (am *AuthMethod) convertValueObjects(ctx context.Context) (*convertedValues, error) {
	const op = "ldap.(AuthMethod).convertValueObjects"
	if am.PublicId == "" {
		return nil, errors.New(ctx, errors.InvalidPublicId, op, "missing public id")
	}
	var err error
	converted := &convertedValues{}

	if converted.Urls, err = am.convertUrls(ctx); err != nil {
		return nil, errors.Wrap(ctx, err, op)
	}
	if converted.Certs, err = am.convertCertificates(ctx); err != nil {
		return nil, errors.Wrap(ctx, err, op)
	}
	if converted.UserEntrySearchConf, err = am.convertUserEntrySearchConf(ctx); err != nil {
		return nil, errors.Wrap(ctx, err, op)
	}
	if converted.GroupEntrySearchConf, err = am.convertGroupEntrySearchConf(ctx); err != nil {
		return nil, errors.Wrap(ctx, err, op)
	}
	if converted.ClientCertificate, err = am.convertClientCertificate(ctx); err != nil {
		return nil, errors.Wrap(ctx, err, op)
	}
	if converted.BindCredential, err = am.convertBindCredential(ctx); err != nil {
		return nil, errors.Wrap(ctx, err, op)
	}

	return converted, nil
}

// convertCertificates converts any embedded URLs from []string
// to []any where each slice element is a *Url. It will return an error if the
// AuthMethod's public id is not set.
func (am *AuthMethod) convertUrls(ctx context.Context) ([]any, error) {
	const op = "ldap.(AuthMethod).convertUrls"
	if am.PublicId == "" {
		return nil, errors.New(ctx, errors.InvalidPublicId, op, "missing public id")
	}
	newValObjs := make([]any, 0, len(am.Urls))
	for priority, u := range am.Urls {
		parsed, err := url.Parse(u)
		if err != nil {
			return nil, errors.Wrap(ctx, err, op)
		}
		obj, err := NewUrl(ctx, am.PublicId, priority, parsed)
		if err != nil {
			return nil, errors.Wrap(ctx, err, op)
		}
		newValObjs = append(newValObjs, obj)
	}
	return newValObjs, nil
}

// convertCertificates converts any embedded certificates from []string
// to []any where each slice element is a *Certificate. It will return an error
// if the AuthMethod's public id is not set.
func (am *AuthMethod) convertCertificates(ctx context.Context) ([]any, error) {
	const op = "ldap.(AuthMethod).convertCertificates"
	if am.PublicId == "" {
		return nil, errors.New(ctx, errors.InvalidPublicId, op, "missing public id")
	}
	newValObjs := make([]any, 0, len(am.Certificates))
	for _, cert := range am.Certificates {
		obj, err := NewCertificate(ctx, am.PublicId, cert)
		if err != nil {
			return nil, errors.Wrap(ctx, err, op)
		}
		newValObjs = append(newValObjs, obj)
	}
	return newValObjs, nil
}

// convertUserEntrySearchConf converts an embedded user entry search fields
// into an any type.  It will return an error if the AuthMethod's public id is
// not set.
func (am *AuthMethod) convertUserEntrySearchConf(ctx context.Context) (any, error) {
	const op = "ldap.(AuthMethod).convertUserEntrySearchConf"
	if am.PublicId == "" {
		return nil, errors.New(ctx, errors.InvalidPublicId, op, "missing public id")
	}
	c, err := NewUserEntrySearchConf(ctx, am.PublicId, WithUserDn(ctx, am.UserDn), WithUserAttr(ctx, am.UserAttr), WithUserFilter(ctx, am.UserFilter))
	if err != nil {
		return nil, errors.Wrap(ctx, err, op)
	}
	return c, nil
}

// convertGroupEntrySearchConf converts an embedded group entry search fields
// into an any type.  It will return an error if the AuthMethod's public id is
// not set.
func (am *AuthMethod) convertGroupEntrySearchConf(ctx context.Context) (any, error) {
	const op = "ldap.(AuthMethod).convertGroupEntrySearchConf"
	if am.PublicId == "" {
		return nil, errors.New(ctx, errors.InvalidPublicId, op, "missing public id")
	}
	c, err := NewGroupEntrySearchConf(ctx, am.PublicId, WithGroupDn(ctx, am.GroupDn), WithGroupAttr(ctx, am.GroupAttr), WithGroupFilter(ctx, am.GroupFilter))
	if err != nil {
		return nil, errors.Wrap(ctx, err, op)
	}
	return c, nil
}

// convertClientCertificate converts an embedded client certificate entry into
// an any type.  It will return an error if the AuthMethod's public id is not
// set.
func (am *AuthMethod) convertClientCertificate(ctx context.Context) (any, error) {
	const op = "ldap.(AuthMethod).convertClientCertificate"
	if am.PublicId == "" {
		return nil, errors.New(ctx, errors.InvalidPublicId, op, "missing auth method id")
	}
	cc, err := NewClientCertificate(ctx, am.PublicId, am.ClientCertificateKey, am.ClientCertificate)
	if err != nil {
		return nil, errors.Wrap(ctx, err, op)
	}
	return cc, nil
}

// convertBindCredential converts an embedded bind credential entry into
// an any type.  It will return an error if the AuthMethod's public id is not
// set.
func (am *AuthMethod) convertBindCredential(ctx context.Context) (any, error) {
	const op = "ldap.(AuthMethod).convertBindCredentials"
	if am.PublicId == "" {
		return nil, errors.New(ctx, errors.InvalidPublicId, op, "missing auth method id")
	}
	bc, err := NewBindCredential(ctx, am.PublicId, am.BindDn, []byte(am.BindPassword))
	if err != nil {
		return nil, errors.Wrap(ctx, err, op)
	}
	return bc, nil
}
