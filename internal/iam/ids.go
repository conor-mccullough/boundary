// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iam

import (
	"fmt"

	"github.com/hashicorp/boundary/globals"
	"github.com/hashicorp/boundary/internal/db"
	"github.com/hashicorp/boundary/internal/errors"
	"github.com/hashicorp/boundary/internal/types/scope"
)

const (
	// RoleGrantPrefix is the prefix for role grants
	RoleGrantPrefix = "rg"
)

func newRoleId() (string, error) {
	id, err := db.NewPublicId(globals.RolePrefix)
	if err != nil {
		return "", errors.WrapDeprecated(err, "iam.newRoleId")
	}
	return id, nil
}

func newUserId() (string, error) {
	id, err := db.NewPublicId(globals.UserPrefix)
	if err != nil {
		return "", errors.WrapDeprecated(err, "iam.newUserId")
	}
	return id, nil
}

func newGroupId() (string, error) {
	id, err := db.NewPublicId(globals.GroupPrefix)
	if err != nil {
		return "", errors.WrapDeprecated(err, "iam.newGroupId")
	}
	return id, nil
}

func newScopeId(scopeType scope.Type) (string, error) {
	const op = "iam.newScopeId"
	if scopeType == scope.Unknown {
		return "", errors.NewDeprecated(errors.InvalidParameter, op, "unknown scope is not supported")
	}
	id, err := db.NewPublicId(scopeType.Prefix())
	if err != nil {
		return "", errors.WrapDeprecated(err, op, errors.WithMsg(fmt.Sprintf("scope type: %s", scopeType.String())))
	}
	return id, nil
}
