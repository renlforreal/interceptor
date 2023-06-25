// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package report

import (
	"time"

	"github.com/pion/logging"
)

// ReceiverOption can be used to configure ReceiverInterceptor.
type ReceiverOption func(r *ReceiverInterceptor) error

// ReceiverLog sets a logger for the interceptor.
func ReceiverLog(log logging.LeveledLogger) ReceiverOption {
	return func(r *ReceiverInterceptor) error {
		r.log = log
		return nil
	}
}

// ReceiverNow sets an alternative for the time.Now function.
func ReceiverNow(f func() time.Time) ReceiverOption {
	return func(r *ReceiverInterceptor) error {
		r.now = f
		return nil
	}
}
