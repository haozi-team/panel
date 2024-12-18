// Code generated by mockery. DO NOT EDIT.

package biz

import (
	biz "github.com/TheTNB/panel/internal/biz"
	api "github.com/TheTNB/panel/pkg/api"

	mock "github.com/stretchr/testify/mock"
)

// AppRepo is an autogenerated mock type for the AppRepo type
type AppRepo struct {
	mock.Mock
}

type AppRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *AppRepo) EXPECT() *AppRepo_Expecter {
	return &AppRepo_Expecter{mock: &_m.Mock}
}

// All provides a mock function with no fields
func (_m *AppRepo) All() api.Apps {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for All")
	}

	var r0 api.Apps
	if rf, ok := ret.Get(0).(func() api.Apps); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.Apps)
		}
	}

	return r0
}

// AppRepo_All_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'All'
type AppRepo_All_Call struct {
	*mock.Call
}

// All is a helper method to define mock.On call
func (_e *AppRepo_Expecter) All() *AppRepo_All_Call {
	return &AppRepo_All_Call{Call: _e.mock.On("All")}
}

func (_c *AppRepo_All_Call) Run(run func()) *AppRepo_All_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AppRepo_All_Call) Return(_a0 api.Apps) *AppRepo_All_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppRepo_All_Call) RunAndReturn(run func() api.Apps) *AppRepo_All_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: slug
func (_m *AppRepo) Get(slug string) (*api.App, error) {
	ret := _m.Called(slug)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *api.App
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*api.App, error)); ok {
		return rf(slug)
	}
	if rf, ok := ret.Get(0).(func(string) *api.App); ok {
		r0 = rf(slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.App)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type AppRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - slug string
func (_e *AppRepo_Expecter) Get(slug interface{}) *AppRepo_Get_Call {
	return &AppRepo_Get_Call{Call: _e.mock.On("Get", slug)}
}

func (_c *AppRepo_Get_Call) Run(run func(slug string)) *AppRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AppRepo_Get_Call) Return(_a0 *api.App, _a1 error) *AppRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AppRepo_Get_Call) RunAndReturn(run func(string) (*api.App, error)) *AppRepo_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetHomeShow provides a mock function with no fields
func (_m *AppRepo) GetHomeShow() ([]map[string]string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetHomeShow")
	}

	var r0 []map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]map[string]string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppRepo_GetHomeShow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHomeShow'
type AppRepo_GetHomeShow_Call struct {
	*mock.Call
}

// GetHomeShow is a helper method to define mock.On call
func (_e *AppRepo_Expecter) GetHomeShow() *AppRepo_GetHomeShow_Call {
	return &AppRepo_GetHomeShow_Call{Call: _e.mock.On("GetHomeShow")}
}

func (_c *AppRepo_GetHomeShow_Call) Run(run func()) *AppRepo_GetHomeShow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AppRepo_GetHomeShow_Call) Return(_a0 []map[string]string, _a1 error) *AppRepo_GetHomeShow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AppRepo_GetHomeShow_Call) RunAndReturn(run func() ([]map[string]string, error)) *AppRepo_GetHomeShow_Call {
	_c.Call.Return(run)
	return _c
}

// GetInstalled provides a mock function with given fields: slug
func (_m *AppRepo) GetInstalled(slug string) (*biz.App, error) {
	ret := _m.Called(slug)

	if len(ret) == 0 {
		panic("no return value specified for GetInstalled")
	}

	var r0 *biz.App
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*biz.App, error)); ok {
		return rf(slug)
	}
	if rf, ok := ret.Get(0).(func(string) *biz.App); ok {
		r0 = rf(slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*biz.App)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppRepo_GetInstalled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInstalled'
type AppRepo_GetInstalled_Call struct {
	*mock.Call
}

// GetInstalled is a helper method to define mock.On call
//   - slug string
func (_e *AppRepo_Expecter) GetInstalled(slug interface{}) *AppRepo_GetInstalled_Call {
	return &AppRepo_GetInstalled_Call{Call: _e.mock.On("GetInstalled", slug)}
}

func (_c *AppRepo_GetInstalled_Call) Run(run func(slug string)) *AppRepo_GetInstalled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AppRepo_GetInstalled_Call) Return(_a0 *biz.App, _a1 error) *AppRepo_GetInstalled_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AppRepo_GetInstalled_Call) RunAndReturn(run func(string) (*biz.App, error)) *AppRepo_GetInstalled_Call {
	_c.Call.Return(run)
	return _c
}

// GetInstalledAll provides a mock function with given fields: query, cond
func (_m *AppRepo) GetInstalledAll(query string, cond ...string) ([]*biz.App, error) {
	_va := make([]interface{}, len(cond))
	for _i := range cond {
		_va[_i] = cond[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetInstalledAll")
	}

	var r0 []*biz.App
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...string) ([]*biz.App, error)); ok {
		return rf(query, cond...)
	}
	if rf, ok := ret.Get(0).(func(string, ...string) []*biz.App); ok {
		r0 = rf(query, cond...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*biz.App)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...string) error); ok {
		r1 = rf(query, cond...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppRepo_GetInstalledAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInstalledAll'
type AppRepo_GetInstalledAll_Call struct {
	*mock.Call
}

// GetInstalledAll is a helper method to define mock.On call
//   - query string
//   - cond ...string
func (_e *AppRepo_Expecter) GetInstalledAll(query interface{}, cond ...interface{}) *AppRepo_GetInstalledAll_Call {
	return &AppRepo_GetInstalledAll_Call{Call: _e.mock.On("GetInstalledAll",
		append([]interface{}{query}, cond...)...)}
}

func (_c *AppRepo_GetInstalledAll_Call) Run(run func(query string, cond ...string)) *AppRepo_GetInstalledAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *AppRepo_GetInstalledAll_Call) Return(_a0 []*biz.App, _a1 error) *AppRepo_GetInstalledAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AppRepo_GetInstalledAll_Call) RunAndReturn(run func(string, ...string) ([]*biz.App, error)) *AppRepo_GetInstalledAll_Call {
	_c.Call.Return(run)
	return _c
}

// Install provides a mock function with given fields: channel, slug
func (_m *AppRepo) Install(channel string, slug string) error {
	ret := _m.Called(channel, slug)

	if len(ret) == 0 {
		panic("no return value specified for Install")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(channel, slug)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppRepo_Install_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Install'
type AppRepo_Install_Call struct {
	*mock.Call
}

// Install is a helper method to define mock.On call
//   - channel string
//   - slug string
func (_e *AppRepo_Expecter) Install(channel interface{}, slug interface{}) *AppRepo_Install_Call {
	return &AppRepo_Install_Call{Call: _e.mock.On("Install", channel, slug)}
}

func (_c *AppRepo_Install_Call) Run(run func(channel string, slug string)) *AppRepo_Install_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *AppRepo_Install_Call) Return(_a0 error) *AppRepo_Install_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppRepo_Install_Call) RunAndReturn(run func(string, string) error) *AppRepo_Install_Call {
	_c.Call.Return(run)
	return _c
}

// Installed provides a mock function with no fields
func (_m *AppRepo) Installed() ([]*biz.App, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Installed")
	}

	var r0 []*biz.App
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*biz.App, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*biz.App); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*biz.App)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppRepo_Installed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Installed'
type AppRepo_Installed_Call struct {
	*mock.Call
}

// Installed is a helper method to define mock.On call
func (_e *AppRepo_Expecter) Installed() *AppRepo_Installed_Call {
	return &AppRepo_Installed_Call{Call: _e.mock.On("Installed")}
}

func (_c *AppRepo_Installed_Call) Run(run func()) *AppRepo_Installed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AppRepo_Installed_Call) Return(_a0 []*biz.App, _a1 error) *AppRepo_Installed_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AppRepo_Installed_Call) RunAndReturn(run func() ([]*biz.App, error)) *AppRepo_Installed_Call {
	_c.Call.Return(run)
	return _c
}

// IsInstalled provides a mock function with given fields: query, cond
func (_m *AppRepo) IsInstalled(query string, cond ...string) (bool, error) {
	_va := make([]interface{}, len(cond))
	for _i := range cond {
		_va[_i] = cond[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for IsInstalled")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...string) (bool, error)); ok {
		return rf(query, cond...)
	}
	if rf, ok := ret.Get(0).(func(string, ...string) bool); ok {
		r0 = rf(query, cond...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, ...string) error); ok {
		r1 = rf(query, cond...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppRepo_IsInstalled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsInstalled'
type AppRepo_IsInstalled_Call struct {
	*mock.Call
}

// IsInstalled is a helper method to define mock.On call
//   - query string
//   - cond ...string
func (_e *AppRepo_Expecter) IsInstalled(query interface{}, cond ...interface{}) *AppRepo_IsInstalled_Call {
	return &AppRepo_IsInstalled_Call{Call: _e.mock.On("IsInstalled",
		append([]interface{}{query}, cond...)...)}
}

func (_c *AppRepo_IsInstalled_Call) Run(run func(query string, cond ...string)) *AppRepo_IsInstalled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *AppRepo_IsInstalled_Call) Return(_a0 bool, _a1 error) *AppRepo_IsInstalled_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AppRepo_IsInstalled_Call) RunAndReturn(run func(string, ...string) (bool, error)) *AppRepo_IsInstalled_Call {
	_c.Call.Return(run)
	return _c
}

// UnInstall provides a mock function with given fields: slug
func (_m *AppRepo) UnInstall(slug string) error {
	ret := _m.Called(slug)

	if len(ret) == 0 {
		panic("no return value specified for UnInstall")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(slug)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppRepo_UnInstall_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnInstall'
type AppRepo_UnInstall_Call struct {
	*mock.Call
}

// UnInstall is a helper method to define mock.On call
//   - slug string
func (_e *AppRepo_Expecter) UnInstall(slug interface{}) *AppRepo_UnInstall_Call {
	return &AppRepo_UnInstall_Call{Call: _e.mock.On("UnInstall", slug)}
}

func (_c *AppRepo_UnInstall_Call) Run(run func(slug string)) *AppRepo_UnInstall_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AppRepo_UnInstall_Call) Return(_a0 error) *AppRepo_UnInstall_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppRepo_UnInstall_Call) RunAndReturn(run func(string) error) *AppRepo_UnInstall_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: slug
func (_m *AppRepo) Update(slug string) error {
	ret := _m.Called(slug)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(slug)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppRepo_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type AppRepo_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - slug string
func (_e *AppRepo_Expecter) Update(slug interface{}) *AppRepo_Update_Call {
	return &AppRepo_Update_Call{Call: _e.mock.On("Update", slug)}
}

func (_c *AppRepo_Update_Call) Run(run func(slug string)) *AppRepo_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AppRepo_Update_Call) Return(_a0 error) *AppRepo_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppRepo_Update_Call) RunAndReturn(run func(string) error) *AppRepo_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateExist provides a mock function with given fields: slug
func (_m *AppRepo) UpdateExist(slug string) bool {
	ret := _m.Called(slug)

	if len(ret) == 0 {
		panic("no return value specified for UpdateExist")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(slug)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// AppRepo_UpdateExist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateExist'
type AppRepo_UpdateExist_Call struct {
	*mock.Call
}

// UpdateExist is a helper method to define mock.On call
//   - slug string
func (_e *AppRepo_Expecter) UpdateExist(slug interface{}) *AppRepo_UpdateExist_Call {
	return &AppRepo_UpdateExist_Call{Call: _e.mock.On("UpdateExist", slug)}
}

func (_c *AppRepo_UpdateExist_Call) Run(run func(slug string)) *AppRepo_UpdateExist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AppRepo_UpdateExist_Call) Return(_a0 bool) *AppRepo_UpdateExist_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppRepo_UpdateExist_Call) RunAndReturn(run func(string) bool) *AppRepo_UpdateExist_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateShow provides a mock function with given fields: slug, show
func (_m *AppRepo) UpdateShow(slug string, show bool) error {
	ret := _m.Called(slug, show)

	if len(ret) == 0 {
		panic("no return value specified for UpdateShow")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool) error); ok {
		r0 = rf(slug, show)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppRepo_UpdateShow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateShow'
type AppRepo_UpdateShow_Call struct {
	*mock.Call
}

// UpdateShow is a helper method to define mock.On call
//   - slug string
//   - show bool
func (_e *AppRepo_Expecter) UpdateShow(slug interface{}, show interface{}) *AppRepo_UpdateShow_Call {
	return &AppRepo_UpdateShow_Call{Call: _e.mock.On("UpdateShow", slug, show)}
}

func (_c *AppRepo_UpdateShow_Call) Run(run func(slug string, show bool)) *AppRepo_UpdateShow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(bool))
	})
	return _c
}

func (_c *AppRepo_UpdateShow_Call) Return(_a0 error) *AppRepo_UpdateShow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppRepo_UpdateShow_Call) RunAndReturn(run func(string, bool) error) *AppRepo_UpdateShow_Call {
	_c.Call.Return(run)
	return _c
}

// NewAppRepo creates a new instance of AppRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAppRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *AppRepo {
	mock := &AppRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}