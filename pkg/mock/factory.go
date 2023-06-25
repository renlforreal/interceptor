// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package mock

import "github.com/renlforreal/interceptor"

// Factory is a mock Factory for testing.
type Factory struct {
	NewInterceptorFn func(id string) (interceptor.Interceptor, error)
}

// NewInterceptor implements Interceptor
func (f *Factory) NewInterceptor(id string) (interceptor.Interceptor, error) {
	return f.NewInterceptorFn(id)
}
