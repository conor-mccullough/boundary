// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package boundary

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/hashicorp/boundary/api/sessions"
	"github.com/hashicorp/boundary/testing/internal/e2e"
	"github.com/stretchr/testify/require"
)

// WaitForSessionCli waits for a session to appear in the session list and returns the session
// information
func WaitForSessionCli(t testing.TB, ctx context.Context, projectId string) *sessions.Session {
	t.Log("Waiting for session to appear...")
	var session *sessions.Session
	err := backoff.RetryNotify(
		func() error {
			// List sessions
			output := e2e.RunCommand(ctx, "boundary",
				e2e.WithArgs("sessions", "list", "-scope-id", projectId, "-include-terminated", "-format", "json"),
			)
			if output.Err != nil {
				return backoff.Permanent(errors.New(string(output.Stderr)))
			}
			var sessionListResult sessions.SessionListResult
			err := json.Unmarshal(output.Stdout, &sessionListResult)
			if err != nil {
				return backoff.Permanent(err)
			}

			// Check if there is one session
			sessionCount := len(sessionListResult.Items)
			if sessionCount == 0 {
				return errors.New("No items are appearing in the session list")
			}

			t.Logf("Found %d session(s)", sessionCount)
			if sessionCount != 1 {
				return backoff.Permanent(errors.New("Only one session was expected to be found"))
			}

			session = sessionListResult.Items[0]
			return nil
		},
		backoff.WithMaxRetries(backoff.NewConstantBackOff(3*time.Second), 5),
		func(err error, td time.Duration) {
			t.Logf("%s. Retrying...", err.Error())
		},
	)
	require.NoError(t, err)

	return session
}

// WaitForSessionStatusCli reads the specified session and waits for its status to match the
// specified status
func WaitForSessionStatusCli(t testing.TB, ctx context.Context, sessionId string, status string) {
	t.Logf("Waiting for session to be %s...", status)
	err := backoff.RetryNotify(
		func() error {
			output := e2e.RunCommand(ctx, "boundary",
				e2e.WithArgs("sessions", "read", "-id", sessionId, "-format", "json"),
			)
			if output.Err != nil {
				return backoff.Permanent(errors.New(string(output.Stderr)))
			}

			var sessionReadResult sessions.SessionReadResult
			err := json.Unmarshal(output.Stdout, &sessionReadResult)
			if err != nil {
				return backoff.Permanent(err)
			}
			s := sessionReadResult.Item

			t.Logf("Id: %s", sessionReadResult.Item.Id)
			t.Logf("Created Time: %s", sessionReadResult.Item.CreatedTime)
			t.Logf("Updated Time: %s", sessionReadResult.Item.UpdatedTime)
			t.Logf("Expiration Time: %s", sessionReadResult.Item.ExpirationTime)
			for _, v := range sessionReadResult.Item.States {
				t.Logf("Status: %s, Start Time: %s, End Time: %s", v.Status, v.StartTime, v.EndTime)
			}
			t.Logf("Termination Reason: %s", sessionReadResult.Item.TerminationReason)
			if s.Status != status {
				return errors.New(fmt.Sprintf("Waiting for session status... Expected: %s, Actual: %s",
					status,
					sessionReadResult.Item.Status,
				))
			}

			return nil
		},
		backoff.WithMaxRetries(backoff.NewConstantBackOff(3*time.Second), 5),
		func(err error, td time.Duration) {
			t.Logf("%s. Retrying...", err.Error())
		},
	)
	require.NoError(t, err)
}
