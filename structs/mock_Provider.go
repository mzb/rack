// Code generated by mockery v1.0.0
package structs

import io "io"
import mock "github.com/stretchr/testify/mock"

// MockProvider is an autogenerated mock type for the Provider type
type MockProvider struct {
	mock.Mock
}

// AppCancel provides a mock function with given fields: name
func (_m *MockProvider) AppCancel(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppCreate provides a mock function with given fields: name, opts
func (_m *MockProvider) AppCreate(name string, opts AppCreateOptions) (*App, error) {
	ret := _m.Called(name, opts)

	var r0 *App
	if rf, ok := ret.Get(0).(func(string, AppCreateOptions) *App); ok {
		r0 = rf(name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*App)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, AppCreateOptions) error); ok {
		r1 = rf(name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppDelete provides a mock function with given fields: name
func (_m *MockProvider) AppDelete(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppGet provides a mock function with given fields: name
func (_m *MockProvider) AppGet(name string) (*App, error) {
	ret := _m.Called(name)

	var r0 *App
	if rf, ok := ret.Get(0).(func(string) *App); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*App)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppList provides a mock function with given fields:
func (_m *MockProvider) AppList() (Apps, error) {
	ret := _m.Called()

	var r0 Apps
	if rf, ok := ret.Get(0).(func() Apps); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Apps)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppLogs provides a mock function with given fields: app, opts
func (_m *MockProvider) AppLogs(app string, opts LogsOptions) (io.ReadCloser, error) {
	ret := _m.Called(app, opts)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, LogsOptions) io.ReadCloser); ok {
		r0 = rf(app, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, LogsOptions) error); ok {
		r1 = rf(app, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppUpdate provides a mock function with given fields: app, opts
func (_m *MockProvider) AppUpdate(app string, opts AppUpdateOptions) error {
	ret := _m.Called(app, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, AppUpdateOptions) error); ok {
		r0 = rf(app, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BuildCreate provides a mock function with given fields: app, method, source, opts
func (_m *MockProvider) BuildCreate(app string, method string, source string, opts BuildCreateOptions) (*Build, error) {
	ret := _m.Called(app, method, source, opts)

	var r0 *Build
	if rf, ok := ret.Get(0).(func(string, string, string, BuildCreateOptions) *Build); ok {
		r0 = rf(app, method, source, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, BuildCreateOptions) error); ok {
		r1 = rf(app, method, source, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildExport provides a mock function with given fields: app, id, w
func (_m *MockProvider) BuildExport(app string, id string, w io.Writer) error {
	ret := _m.Called(app, id, w)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, io.Writer) error); ok {
		r0 = rf(app, id, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BuildGet provides a mock function with given fields: app, id
func (_m *MockProvider) BuildGet(app string, id string) (*Build, error) {
	ret := _m.Called(app, id)

	var r0 *Build
	if rf, ok := ret.Get(0).(func(string, string) *Build); ok {
		r0 = rf(app, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildImport provides a mock function with given fields: app, r
func (_m *MockProvider) BuildImport(app string, r io.Reader) (*Build, error) {
	ret := _m.Called(app, r)

	var r0 *Build
	if rf, ok := ret.Get(0).(func(string, io.Reader) *Build); ok {
		r0 = rf(app, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, io.Reader) error); ok {
		r1 = rf(app, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildList provides a mock function with given fields: app, opts
func (_m *MockProvider) BuildList(app string, opts BuildListOptions) (Builds, error) {
	ret := _m.Called(app, opts)

	var r0 Builds
	if rf, ok := ret.Get(0).(func(string, BuildListOptions) Builds); ok {
		r0 = rf(app, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Builds)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, BuildListOptions) error); ok {
		r1 = rf(app, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildLogs provides a mock function with given fields: app, id, opts
func (_m *MockProvider) BuildLogs(app string, id string, opts LogsOptions) (io.ReadCloser, error) {
	ret := _m.Called(app, id, opts)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, string, LogsOptions) io.ReadCloser); ok {
		r0 = rf(app, id, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, LogsOptions) error); ok {
		r1 = rf(app, id, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildUpdate provides a mock function with given fields: app, id, opts
func (_m *MockProvider) BuildUpdate(app string, id string, opts BuildUpdateOptions) (*Build, error) {
	ret := _m.Called(app, id, opts)

	var r0 *Build
	if rf, ok := ret.Get(0).(func(string, string, BuildUpdateOptions) *Build); ok {
		r0 = rf(app, id, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, BuildUpdateOptions) error); ok {
		r1 = rf(app, id, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CapacityGet provides a mock function with given fields:
func (_m *MockProvider) CapacityGet() (*Capacity, error) {
	ret := _m.Called()

	var r0 *Capacity
	if rf, ok := ret.Get(0).(func() *Capacity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Capacity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CertificateApply provides a mock function with given fields: app, service, port, id
func (_m *MockProvider) CertificateApply(app string, service string, port int, id string) error {
	ret := _m.Called(app, service, port, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int, string) error); ok {
		r0 = rf(app, service, port, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CertificateCreate provides a mock function with given fields: pub, key, chain
func (_m *MockProvider) CertificateCreate(pub string, key string, chain string) (*Certificate, error) {
	ret := _m.Called(pub, key, chain)

	var r0 *Certificate
	if rf, ok := ret.Get(0).(func(string, string, string) *Certificate); ok {
		r0 = rf(pub, key, chain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Certificate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(pub, key, chain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CertificateDelete provides a mock function with given fields: id
func (_m *MockProvider) CertificateDelete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CertificateGenerate provides a mock function with given fields: domains
func (_m *MockProvider) CertificateGenerate(domains []string) (*Certificate, error) {
	ret := _m.Called(domains)

	var r0 *Certificate
	if rf, ok := ret.Get(0).(func([]string) *Certificate); ok {
		r0 = rf(domains)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Certificate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(domains)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CertificateList provides a mock function with given fields:
func (_m *MockProvider) CertificateList() (Certificates, error) {
	ret := _m.Called()

	var r0 Certificates
	if rf, ok := ret.Get(0).(func() Certificates); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Certificates)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventSend provides a mock function with given fields: action, opts
func (_m *MockProvider) EventSend(action string, opts EventSendOptions) error {
	ret := _m.Called(action, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, EventSendOptions) error); ok {
		r0 = rf(action, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilesDelete provides a mock function with given fields: app, pid, files
func (_m *MockProvider) FilesDelete(app string, pid string, files []string) error {
	ret := _m.Called(app, pid, files)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []string) error); ok {
		r0 = rf(app, pid, files)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilesUpload provides a mock function with given fields: app, pid, r
func (_m *MockProvider) FilesUpload(app string, pid string, r io.Reader) error {
	ret := _m.Called(app, pid, r)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, io.Reader) error); ok {
		r0 = rf(app, pid, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Initialize provides a mock function with given fields: opts
func (_m *MockProvider) Initialize(opts ProviderOptions) error {
	ret := _m.Called(opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(ProviderOptions) error); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InstanceKeyroll provides a mock function with given fields:
func (_m *MockProvider) InstanceKeyroll() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InstanceList provides a mock function with given fields:
func (_m *MockProvider) InstanceList() (Instances, error) {
	ret := _m.Called()

	var r0 Instances
	if rf, ok := ret.Get(0).(func() Instances); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Instances)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InstanceShell provides a mock function with given fields: id, rw, opts
func (_m *MockProvider) InstanceShell(id string, rw io.ReadWriter, opts InstanceShellOptions) error {
	ret := _m.Called(id, rw, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, io.ReadWriter, InstanceShellOptions) error); ok {
		r0 = rf(id, rw, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InstanceTerminate provides a mock function with given fields: id
func (_m *MockProvider) InstanceTerminate(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ObjectDelete provides a mock function with given fields: app, key
func (_m *MockProvider) ObjectDelete(app string, key string) error {
	ret := _m.Called(app, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(app, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ObjectExists provides a mock function with given fields: app, key
func (_m *MockProvider) ObjectExists(app string, key string) (bool, error) {
	ret := _m.Called(app, key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(app, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ObjectFetch provides a mock function with given fields: app, key
func (_m *MockProvider) ObjectFetch(app string, key string) (io.ReadCloser, error) {
	ret := _m.Called(app, key)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, string) io.ReadCloser); ok {
		r0 = rf(app, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ObjectList provides a mock function with given fields: app, prefix
func (_m *MockProvider) ObjectList(app string, prefix string) ([]string, error) {
	ret := _m.Called(app, prefix)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(app, prefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ObjectStore provides a mock function with given fields: app, key, r, opts
func (_m *MockProvider) ObjectStore(app string, key string, r io.Reader, opts ObjectStoreOptions) (*Object, error) {
	ret := _m.Called(app, key, r, opts)

	var r0 *Object
	if rf, ok := ret.Get(0).(func(string, string, io.Reader, ObjectStoreOptions) *Object); ok {
		r0 = rf(app, key, r, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Object)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, io.Reader, ObjectStoreOptions) error); ok {
		r1 = rf(app, key, r, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessExec provides a mock function with given fields: app, pid, command, opts
func (_m *MockProvider) ProcessExec(app string, pid string, command string, opts ProcessExecOptions) (int, error) {
	ret := _m.Called(app, pid, command, opts)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string, ProcessExecOptions) int); ok {
		r0 = rf(app, pid, command, opts)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, ProcessExecOptions) error); ok {
		r1 = rf(app, pid, command, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessGet provides a mock function with given fields: app, pid
func (_m *MockProvider) ProcessGet(app string, pid string) (*Process, error) {
	ret := _m.Called(app, pid)

	var r0 *Process
	if rf, ok := ret.Get(0).(func(string, string) *Process); ok {
		r0 = rf(app, pid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Process)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, pid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessList provides a mock function with given fields: app, opts
func (_m *MockProvider) ProcessList(app string, opts ProcessListOptions) (Processes, error) {
	ret := _m.Called(app, opts)

	var r0 Processes
	if rf, ok := ret.Get(0).(func(string, ProcessListOptions) Processes); ok {
		r0 = rf(app, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Processes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ProcessListOptions) error); ok {
		r1 = rf(app, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessRun provides a mock function with given fields: app, opts
func (_m *MockProvider) ProcessRun(app string, opts ProcessRunOptions) (string, error) {
	ret := _m.Called(app, opts)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ProcessRunOptions) string); ok {
		r0 = rf(app, opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ProcessRunOptions) error); ok {
		r1 = rf(app, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessStop provides a mock function with given fields: app, pid
func (_m *MockProvider) ProcessStop(app string, pid string) error {
	ret := _m.Called(app, pid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(app, pid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessWait provides a mock function with given fields: app, pid
func (_m *MockProvider) ProcessWait(app string, pid string) (int, error) {
	ret := _m.Called(app, pid)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string) int); ok {
		r0 = rf(app, pid)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, pid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegistryAdd provides a mock function with given fields: server, username, password
func (_m *MockProvider) RegistryAdd(server string, username string, password string) (*Registry, error) {
	ret := _m.Called(server, username, password)

	var r0 *Registry
	if rf, ok := ret.Get(0).(func(string, string, string) *Registry); ok {
		r0 = rf(server, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Registry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(server, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegistryList provides a mock function with given fields:
func (_m *MockProvider) RegistryList() (Registries, error) {
	ret := _m.Called()

	var r0 Registries
	if rf, ok := ret.Get(0).(func() Registries); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Registries)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegistryRemove provides a mock function with given fields: server
func (_m *MockProvider) RegistryRemove(server string) error {
	ret := _m.Called(server)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(server)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReleaseCreate provides a mock function with given fields: app, opts
func (_m *MockProvider) ReleaseCreate(app string, opts ReleaseCreateOptions) (*Release, error) {
	ret := _m.Called(app, opts)

	var r0 *Release
	if rf, ok := ret.Get(0).(func(string, ReleaseCreateOptions) *Release); ok {
		r0 = rf(app, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Release)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ReleaseCreateOptions) error); ok {
		r1 = rf(app, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleaseGet provides a mock function with given fields: app, id
func (_m *MockProvider) ReleaseGet(app string, id string) (*Release, error) {
	ret := _m.Called(app, id)

	var r0 *Release
	if rf, ok := ret.Get(0).(func(string, string) *Release); ok {
		r0 = rf(app, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Release)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleaseList provides a mock function with given fields: app, opts
func (_m *MockProvider) ReleaseList(app string, opts ReleaseListOptions) (Releases, error) {
	ret := _m.Called(app, opts)

	var r0 Releases
	if rf, ok := ret.Get(0).(func(string, ReleaseListOptions) Releases); ok {
		r0 = rf(app, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Releases)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ReleaseListOptions) error); ok {
		r1 = rf(app, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleasePromote provides a mock function with given fields: app, id
func (_m *MockProvider) ReleasePromote(app string, id string) error {
	ret := _m.Called(app, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(app, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResourceCreate provides a mock function with given fields: name, kind, opts
func (_m *MockProvider) ResourceCreate(name string, kind string, opts ResourceCreateOptions) (*Resource, error) {
	ret := _m.Called(name, kind, opts)

	var r0 *Resource
	if rf, ok := ret.Get(0).(func(string, string, ResourceCreateOptions) *Resource); ok {
		r0 = rf(name, kind, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Resource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, ResourceCreateOptions) error); ok {
		r1 = rf(name, kind, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceDelete provides a mock function with given fields: name
func (_m *MockProvider) ResourceDelete(name string) (*Resource, error) {
	ret := _m.Called(name)

	var r0 *Resource
	if rf, ok := ret.Get(0).(func(string) *Resource); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Resource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceGet provides a mock function with given fields: name
func (_m *MockProvider) ResourceGet(name string) (*Resource, error) {
	ret := _m.Called(name)

	var r0 *Resource
	if rf, ok := ret.Get(0).(func(string) *Resource); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Resource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceLink provides a mock function with given fields: name, app, process
func (_m *MockProvider) ResourceLink(name string, app string, process string) (*Resource, error) {
	ret := _m.Called(name, app, process)

	var r0 *Resource
	if rf, ok := ret.Get(0).(func(string, string, string) *Resource); ok {
		r0 = rf(name, app, process)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Resource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(name, app, process)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceList provides a mock function with given fields:
func (_m *MockProvider) ResourceList() (Resources, error) {
	ret := _m.Called()

	var r0 Resources
	if rf, ok := ret.Get(0).(func() Resources); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Resources)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceUnlink provides a mock function with given fields: name, app, process
func (_m *MockProvider) ResourceUnlink(name string, app string, process string) (*Resource, error) {
	ret := _m.Called(name, app, process)

	var r0 *Resource
	if rf, ok := ret.Get(0).(func(string, string, string) *Resource); ok {
		r0 = rf(name, app, process)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Resource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(name, app, process)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceUpdate provides a mock function with given fields: name, params
func (_m *MockProvider) ResourceUpdate(name string, params map[string]string) (*Resource, error) {
	ret := _m.Called(name, params)

	var r0 *Resource
	if rf, ok := ret.Get(0).(func(string, map[string]string) *Resource); ok {
		r0 = rf(name, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Resource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]string) error); ok {
		r1 = rf(name, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceList provides a mock function with given fields: app
func (_m *MockProvider) ServiceList(app string) (Services, error) {
	ret := _m.Called(app)

	var r0 Services
	if rf, ok := ret.Get(0).(func(string) Services); ok {
		r0 = rf(app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Services)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceUpdate provides a mock function with given fields: app, name, opts
func (_m *MockProvider) ServiceUpdate(app string, name string, opts ServiceUpdateOptions) error {
	ret := _m.Called(app, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, ServiceUpdateOptions) error); ok {
		r0 = rf(app, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SettingDelete provides a mock function with given fields: name
func (_m *MockProvider) SettingDelete(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SettingExists provides a mock function with given fields: name
func (_m *MockProvider) SettingExists(name string) (bool, error) {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SettingGet provides a mock function with given fields: name
func (_m *MockProvider) SettingGet(name string) (string, error) {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SettingList provides a mock function with given fields: opts
func (_m *MockProvider) SettingList(opts SettingListOptions) ([]string, error) {
	ret := _m.Called(opts)

	var r0 []string
	if rf, ok := ret.Get(0).(func(SettingListOptions) []string); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(SettingListOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SettingPut provides a mock function with given fields: name, value
func (_m *MockProvider) SettingPut(name string, value string) error {
	ret := _m.Called(name, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(name, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SystemDecrypt provides a mock function with given fields: data
func (_m *MockProvider) SystemDecrypt(data []byte) ([]byte, error) {
	ret := _m.Called(data)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemEncrypt provides a mock function with given fields: data
func (_m *MockProvider) SystemEncrypt(data []byte) ([]byte, error) {
	ret := _m.Called(data)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemGet provides a mock function with given fields:
func (_m *MockProvider) SystemGet() (*System, error) {
	ret := _m.Called()

	var r0 *System
	if rf, ok := ret.Get(0).(func() *System); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*System)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemInstall provides a mock function with given fields: name, opts
func (_m *MockProvider) SystemInstall(name string, opts SystemInstallOptions) (string, error) {
	ret := _m.Called(name, opts)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, SystemInstallOptions) string); ok {
		r0 = rf(name, opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, SystemInstallOptions) error); ok {
		r1 = rf(name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemLogs provides a mock function with given fields: opts
func (_m *MockProvider) SystemLogs(opts LogsOptions) (io.ReadCloser, error) {
	ret := _m.Called(opts)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(LogsOptions) io.ReadCloser); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(LogsOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemProcesses provides a mock function with given fields: opts
func (_m *MockProvider) SystemProcesses(opts SystemProcessesOptions) (Processes, error) {
	ret := _m.Called(opts)

	var r0 Processes
	if rf, ok := ret.Get(0).(func(SystemProcessesOptions) Processes); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Processes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(SystemProcessesOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemReleases provides a mock function with given fields:
func (_m *MockProvider) SystemReleases() (Releases, error) {
	ret := _m.Called()

	var r0 Releases
	if rf, ok := ret.Get(0).(func() Releases); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Releases)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemUpdate provides a mock function with given fields: opts
func (_m *MockProvider) SystemUpdate(opts SystemUpdateOptions) error {
	ret := _m.Called(opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(SystemUpdateOptions) error); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Workers provides a mock function with given fields:
func (_m *MockProvider) Workers() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
