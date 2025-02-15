// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package report

import (
	"time"

	"github.com/pion/logging"
)

// SenderOption can be used to configure SenderInterceptor.
type SenderOption func(r *SenderInterceptor) error

// SenderLog sets a logger for the interceptor.
func SenderLog(log logging.LeveledLogger) SenderOption {
	return func(r *SenderInterceptor) error {
		r.log = log
		return nil
	}
}

// SenderNow sets an alternative for the time.Now function.
func SenderNow(f func() time.Time) SenderOption {
	return func(r *SenderInterceptor) error {
		r.now = f
		return nil
	}
}
