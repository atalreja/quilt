package mocks

import mock "github.com/stretchr/testify/mock"
import ovsdb "github.com/quilt/quilt/minion/ovsdb"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CreateACL provides a mock function with given fields: lswitch, direction, priority, match, action
func (_m *Client) CreateACL(lswitch string, direction string, priority int, match string, action string) error {
	ret := _m.Called(lswitch, direction, priority, match, action)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int, string, string) error); ok {
		r0 = rf(lswitch, direction, priority, match, action)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateAddressSet provides a mock function with given fields: name, addresses
func (_m *Client) CreateAddressSet(name string, addresses []string) error {
	ret := _m.Called(name, addresses)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(name, addresses)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateLogicalPort provides a mock function with given fields: lswitch, name, mac, ip
func (_m *Client) CreateLogicalPort(lswitch string, name string, mac string, ip string) error {
	ret := _m.Called(lswitch, name, mac, ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(lswitch, name, mac, ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateLogicalSwitch provides a mock function with given fields: lswitch
func (_m *Client) CreateLogicalSwitch(lswitch string) error {
	ret := _m.Called(lswitch)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(lswitch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteACL provides a mock function with given fields: lswitch, ovsdbACL
func (_m *Client) DeleteACL(lswitch string, ovsdbACL ovsdb.ACL) error {
	ret := _m.Called(lswitch, ovsdbACL)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ovsdb.ACL) error); ok {
		r0 = rf(lswitch, ovsdbACL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAddressSet provides a mock function with given fields: name
func (_m *Client) DeleteAddressSet(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteLogicalPort provides a mock function with given fields: lswitch, lport
func (_m *Client) DeleteLogicalPort(lswitch string, lport ovsdb.LPort) error {
	ret := _m.Called(lswitch, lport)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ovsdb.LPort) error); ok {
		r0 = rf(lswitch, lport)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Disconnect provides a mock function with given fields:
func (_m *Client) Disconnect() {
	_m.Called()
}

// ListACLs provides a mock function with given fields:
func (_m *Client) ListACLs() ([]ovsdb.ACL, error) {
	ret := _m.Called()

	var r0 []ovsdb.ACL
	if rf, ok := ret.Get(0).(func() []ovsdb.ACL); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ovsdb.ACL)
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

// ListAddressSets provides a mock function with given fields:
func (_m *Client) ListAddressSets() ([]ovsdb.AddressSet, error) {
	ret := _m.Called()

	var r0 []ovsdb.AddressSet
	if rf, ok := ret.Get(0).(func() []ovsdb.AddressSet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ovsdb.AddressSet)
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

// ListLogicalPorts provides a mock function with given fields:
func (_m *Client) ListLogicalPorts() ([]ovsdb.LPort, error) {
	ret := _m.Called()

	var r0 []ovsdb.LPort
	if rf, ok := ret.Get(0).(func() []ovsdb.LPort); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ovsdb.LPort)
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

// OpenFlowPorts provides a mock function with given fields:
func (_m *Client) OpenFlowPorts() (map[string]int, error) {
	ret := _m.Called()

	var r0 map[string]int
	if rf, ok := ret.Get(0).(func() map[string]int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]int)
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

var _ ovsdb.Client = (*Client)(nil)
