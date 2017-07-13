// Copyright 2017 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package mocks

import context "context"
import devauth "github.com/mendersoftware/deviceauth/devauth"
import mock "github.com/stretchr/testify/mock"
import model "github.com/mendersoftware/deviceauth/model"

// App is an autogenerated mock type for the App type
type App struct {
	mock.Mock
}

// AcceptDeviceAuth provides a mock function with given fields: ctx, dev_id, auth_id
func (_m *App) AcceptDeviceAuth(ctx context.Context, dev_id string, auth_id string) error {
	ret := _m.Called(ctx, dev_id, auth_id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, dev_id, auth_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DecommissionDevice provides a mock function with given fields: ctx, dev_id
func (_m *App) DecommissionDevice(ctx context.Context, dev_id string) error {
	ret := _m.Called(ctx, dev_id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, dev_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDevice provides a mock function with given fields: ctx, dev_id
func (_m *App) GetDevice(ctx context.Context, dev_id string) (*model.Device, error) {
	ret := _m.Called(ctx, dev_id)

	var r0 *model.Device
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Device); ok {
		r0 = rf(ctx, dev_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, dev_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceToken provides a mock function with given fields: ctx, dev_id
func (_m *App) GetDeviceToken(ctx context.Context, dev_id string) (*model.Token, error) {
	ret := _m.Called(ctx, dev_id)

	var r0 *model.Token
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Token); ok {
		r0 = rf(ctx, dev_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, dev_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevices provides a mock function with given fields: ctx, skip, limit
func (_m *App) GetDevices(ctx context.Context, skip uint, limit uint) ([]model.Device, error) {
	ret := _m.Called(ctx, skip, limit)

	var r0 []model.Device
	if rf, ok := ret.Get(0).(func(context.Context, uint, uint) []model.Device); ok {
		r0 = rf(ctx, skip, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, uint) error); ok {
		r1 = rf(ctx, skip, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RejectDeviceAuth provides a mock function with given fields: ctx, dev_id, auth_id
func (_m *App) RejectDeviceAuth(ctx context.Context, dev_id string, auth_id string) error {
	ret := _m.Called(ctx, dev_id, auth_id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, dev_id, auth_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetDeviceAuth provides a mock function with given fields: ctx, dev_id, auth_id
func (_m *App) ResetDeviceAuth(ctx context.Context, dev_id string, auth_id string) error {
	ret := _m.Called(ctx, dev_id, auth_id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, dev_id, auth_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RevokeToken provides a mock function with given fields: ctx, token_id
func (_m *App) RevokeToken(ctx context.Context, token_id string) error {
	ret := _m.Called(ctx, token_id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, token_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubmitAuthRequest provides a mock function with given fields: ctx, r
func (_m *App) SubmitAuthRequest(ctx context.Context, r *model.AuthReq) (string, error) {
	ret := _m.Called(ctx, r)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *model.AuthReq) string); ok {
		r0 = rf(ctx, r)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.AuthReq) error); ok {
		r1 = rf(ctx, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyToken provides a mock function with given fields: ctx, token
func (_m *App) VerifyToken(ctx context.Context, token string) error {
	ret := _m.Called(ctx, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ devauth.App = (*App)(nil)
