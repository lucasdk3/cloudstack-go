//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./cloudstack/NetscalerService.go
//
// Generated by this command:
//
//	mockgen -destination=./cloudstack/NetscalerService_mock.go -package=cloudstack -copyright_file=header.txt -source=./cloudstack/NetscalerService.go
//

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockNetscalerServiceIface is a mock of NetscalerServiceIface interface.
type MockNetscalerServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockNetscalerServiceIfaceMockRecorder
	isgomock struct{}
}

// MockNetscalerServiceIfaceMockRecorder is the mock recorder for MockNetscalerServiceIface.
type MockNetscalerServiceIfaceMockRecorder struct {
	mock *MockNetscalerServiceIface
}

// NewMockNetscalerServiceIface creates a new mock instance.
func NewMockNetscalerServiceIface(ctrl *gomock.Controller) *MockNetscalerServiceIface {
	mock := &MockNetscalerServiceIface{ctrl: ctrl}
	mock.recorder = &MockNetscalerServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetscalerServiceIface) EXPECT() *MockNetscalerServiceIfaceMockRecorder {
	return m.recorder
}

// AddNetscalerLoadBalancer mocks base method.
func (m *MockNetscalerServiceIface) AddNetscalerLoadBalancer(p *AddNetscalerLoadBalancerParams) (*AddNetscalerLoadBalancerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNetscalerLoadBalancer", p)
	ret0, _ := ret[0].(*AddNetscalerLoadBalancerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNetscalerLoadBalancer indicates an expected call of AddNetscalerLoadBalancer.
func (mr *MockNetscalerServiceIfaceMockRecorder) AddNetscalerLoadBalancer(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNetscalerLoadBalancer", reflect.TypeOf((*MockNetscalerServiceIface)(nil).AddNetscalerLoadBalancer), p)
}

// ConfigureNetscalerLoadBalancer mocks base method.
func (m *MockNetscalerServiceIface) ConfigureNetscalerLoadBalancer(p *ConfigureNetscalerLoadBalancerParams) (*NetscalerLoadBalancerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureNetscalerLoadBalancer", p)
	ret0, _ := ret[0].(*NetscalerLoadBalancerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigureNetscalerLoadBalancer indicates an expected call of ConfigureNetscalerLoadBalancer.
func (mr *MockNetscalerServiceIfaceMockRecorder) ConfigureNetscalerLoadBalancer(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureNetscalerLoadBalancer", reflect.TypeOf((*MockNetscalerServiceIface)(nil).ConfigureNetscalerLoadBalancer), p)
}

// DeleteNetscalerControlCenter mocks base method.
func (m *MockNetscalerServiceIface) DeleteNetscalerControlCenter(p *DeleteNetscalerControlCenterParams) (*DeleteNetscalerControlCenterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNetscalerControlCenter", p)
	ret0, _ := ret[0].(*DeleteNetscalerControlCenterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNetscalerControlCenter indicates an expected call of DeleteNetscalerControlCenter.
func (mr *MockNetscalerServiceIfaceMockRecorder) DeleteNetscalerControlCenter(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetscalerControlCenter", reflect.TypeOf((*MockNetscalerServiceIface)(nil).DeleteNetscalerControlCenter), p)
}

// DeleteNetscalerLoadBalancer mocks base method.
func (m *MockNetscalerServiceIface) DeleteNetscalerLoadBalancer(p *DeleteNetscalerLoadBalancerParams) (*DeleteNetscalerLoadBalancerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNetscalerLoadBalancer", p)
	ret0, _ := ret[0].(*DeleteNetscalerLoadBalancerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNetscalerLoadBalancer indicates an expected call of DeleteNetscalerLoadBalancer.
func (mr *MockNetscalerServiceIfaceMockRecorder) DeleteNetscalerLoadBalancer(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetscalerLoadBalancer", reflect.TypeOf((*MockNetscalerServiceIface)(nil).DeleteNetscalerLoadBalancer), p)
}

// GetNetscalerLoadBalancerNetworkID mocks base method.
func (m *MockNetscalerServiceIface) GetNetscalerLoadBalancerNetworkID(keyword, lbdeviceid string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{keyword, lbdeviceid}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNetscalerLoadBalancerNetworkID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNetscalerLoadBalancerNetworkID indicates an expected call of GetNetscalerLoadBalancerNetworkID.
func (mr *MockNetscalerServiceIfaceMockRecorder) GetNetscalerLoadBalancerNetworkID(keyword, lbdeviceid any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{keyword, lbdeviceid}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetscalerLoadBalancerNetworkID", reflect.TypeOf((*MockNetscalerServiceIface)(nil).GetNetscalerLoadBalancerNetworkID), varargs...)
}

// ListNetscalerControlCenter mocks base method.
func (m *MockNetscalerServiceIface) ListNetscalerControlCenter(p *ListNetscalerControlCenterParams) (*ListNetscalerControlCenterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNetscalerControlCenter", p)
	ret0, _ := ret[0].(*ListNetscalerControlCenterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNetscalerControlCenter indicates an expected call of ListNetscalerControlCenter.
func (mr *MockNetscalerServiceIfaceMockRecorder) ListNetscalerControlCenter(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNetscalerControlCenter", reflect.TypeOf((*MockNetscalerServiceIface)(nil).ListNetscalerControlCenter), p)
}

// ListNetscalerLoadBalancerNetworks mocks base method.
func (m *MockNetscalerServiceIface) ListNetscalerLoadBalancerNetworks(p *ListNetscalerLoadBalancerNetworksParams) (*ListNetscalerLoadBalancerNetworksResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNetscalerLoadBalancerNetworks", p)
	ret0, _ := ret[0].(*ListNetscalerLoadBalancerNetworksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNetscalerLoadBalancerNetworks indicates an expected call of ListNetscalerLoadBalancerNetworks.
func (mr *MockNetscalerServiceIfaceMockRecorder) ListNetscalerLoadBalancerNetworks(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNetscalerLoadBalancerNetworks", reflect.TypeOf((*MockNetscalerServiceIface)(nil).ListNetscalerLoadBalancerNetworks), p)
}

// ListNetscalerLoadBalancers mocks base method.
func (m *MockNetscalerServiceIface) ListNetscalerLoadBalancers(p *ListNetscalerLoadBalancersParams) (*ListNetscalerLoadBalancersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNetscalerLoadBalancers", p)
	ret0, _ := ret[0].(*ListNetscalerLoadBalancersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNetscalerLoadBalancers indicates an expected call of ListNetscalerLoadBalancers.
func (mr *MockNetscalerServiceIfaceMockRecorder) ListNetscalerLoadBalancers(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNetscalerLoadBalancers", reflect.TypeOf((*MockNetscalerServiceIface)(nil).ListNetscalerLoadBalancers), p)
}

// NewAddNetscalerLoadBalancerParams mocks base method.
func (m *MockNetscalerServiceIface) NewAddNetscalerLoadBalancerParams(networkdevicetype, password, physicalnetworkid, url, username string) *AddNetscalerLoadBalancerParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAddNetscalerLoadBalancerParams", networkdevicetype, password, physicalnetworkid, url, username)
	ret0, _ := ret[0].(*AddNetscalerLoadBalancerParams)
	return ret0
}

// NewAddNetscalerLoadBalancerParams indicates an expected call of NewAddNetscalerLoadBalancerParams.
func (mr *MockNetscalerServiceIfaceMockRecorder) NewAddNetscalerLoadBalancerParams(networkdevicetype, password, physicalnetworkid, url, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAddNetscalerLoadBalancerParams", reflect.TypeOf((*MockNetscalerServiceIface)(nil).NewAddNetscalerLoadBalancerParams), networkdevicetype, password, physicalnetworkid, url, username)
}

// NewConfigureNetscalerLoadBalancerParams mocks base method.
func (m *MockNetscalerServiceIface) NewConfigureNetscalerLoadBalancerParams(lbdeviceid string) *ConfigureNetscalerLoadBalancerParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewConfigureNetscalerLoadBalancerParams", lbdeviceid)
	ret0, _ := ret[0].(*ConfigureNetscalerLoadBalancerParams)
	return ret0
}

// NewConfigureNetscalerLoadBalancerParams indicates an expected call of NewConfigureNetscalerLoadBalancerParams.
func (mr *MockNetscalerServiceIfaceMockRecorder) NewConfigureNetscalerLoadBalancerParams(lbdeviceid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewConfigureNetscalerLoadBalancerParams", reflect.TypeOf((*MockNetscalerServiceIface)(nil).NewConfigureNetscalerLoadBalancerParams), lbdeviceid)
}

// NewDeleteNetscalerControlCenterParams mocks base method.
func (m *MockNetscalerServiceIface) NewDeleteNetscalerControlCenterParams(id string) *DeleteNetscalerControlCenterParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteNetscalerControlCenterParams", id)
	ret0, _ := ret[0].(*DeleteNetscalerControlCenterParams)
	return ret0
}

// NewDeleteNetscalerControlCenterParams indicates an expected call of NewDeleteNetscalerControlCenterParams.
func (mr *MockNetscalerServiceIfaceMockRecorder) NewDeleteNetscalerControlCenterParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteNetscalerControlCenterParams", reflect.TypeOf((*MockNetscalerServiceIface)(nil).NewDeleteNetscalerControlCenterParams), id)
}

// NewDeleteNetscalerLoadBalancerParams mocks base method.
func (m *MockNetscalerServiceIface) NewDeleteNetscalerLoadBalancerParams(lbdeviceid string) *DeleteNetscalerLoadBalancerParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteNetscalerLoadBalancerParams", lbdeviceid)
	ret0, _ := ret[0].(*DeleteNetscalerLoadBalancerParams)
	return ret0
}

// NewDeleteNetscalerLoadBalancerParams indicates an expected call of NewDeleteNetscalerLoadBalancerParams.
func (mr *MockNetscalerServiceIfaceMockRecorder) NewDeleteNetscalerLoadBalancerParams(lbdeviceid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteNetscalerLoadBalancerParams", reflect.TypeOf((*MockNetscalerServiceIface)(nil).NewDeleteNetscalerLoadBalancerParams), lbdeviceid)
}

// NewListNetscalerControlCenterParams mocks base method.
func (m *MockNetscalerServiceIface) NewListNetscalerControlCenterParams() *ListNetscalerControlCenterParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListNetscalerControlCenterParams")
	ret0, _ := ret[0].(*ListNetscalerControlCenterParams)
	return ret0
}

// NewListNetscalerControlCenterParams indicates an expected call of NewListNetscalerControlCenterParams.
func (mr *MockNetscalerServiceIfaceMockRecorder) NewListNetscalerControlCenterParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListNetscalerControlCenterParams", reflect.TypeOf((*MockNetscalerServiceIface)(nil).NewListNetscalerControlCenterParams))
}

// NewListNetscalerLoadBalancerNetworksParams mocks base method.
func (m *MockNetscalerServiceIface) NewListNetscalerLoadBalancerNetworksParams(lbdeviceid string) *ListNetscalerLoadBalancerNetworksParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListNetscalerLoadBalancerNetworksParams", lbdeviceid)
	ret0, _ := ret[0].(*ListNetscalerLoadBalancerNetworksParams)
	return ret0
}

// NewListNetscalerLoadBalancerNetworksParams indicates an expected call of NewListNetscalerLoadBalancerNetworksParams.
func (mr *MockNetscalerServiceIfaceMockRecorder) NewListNetscalerLoadBalancerNetworksParams(lbdeviceid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListNetscalerLoadBalancerNetworksParams", reflect.TypeOf((*MockNetscalerServiceIface)(nil).NewListNetscalerLoadBalancerNetworksParams), lbdeviceid)
}

// NewListNetscalerLoadBalancersParams mocks base method.
func (m *MockNetscalerServiceIface) NewListNetscalerLoadBalancersParams() *ListNetscalerLoadBalancersParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListNetscalerLoadBalancersParams")
	ret0, _ := ret[0].(*ListNetscalerLoadBalancersParams)
	return ret0
}

// NewListNetscalerLoadBalancersParams indicates an expected call of NewListNetscalerLoadBalancersParams.
func (mr *MockNetscalerServiceIfaceMockRecorder) NewListNetscalerLoadBalancersParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListNetscalerLoadBalancersParams", reflect.TypeOf((*MockNetscalerServiceIface)(nil).NewListNetscalerLoadBalancersParams))
}

// NewRegisterNetscalerControlCenterParams mocks base method.
func (m *MockNetscalerServiceIface) NewRegisterNetscalerControlCenterParams(ipaddress string, numretries int, password, username string) *RegisterNetscalerControlCenterParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRegisterNetscalerControlCenterParams", ipaddress, numretries, password, username)
	ret0, _ := ret[0].(*RegisterNetscalerControlCenterParams)
	return ret0
}

// NewRegisterNetscalerControlCenterParams indicates an expected call of NewRegisterNetscalerControlCenterParams.
func (mr *MockNetscalerServiceIfaceMockRecorder) NewRegisterNetscalerControlCenterParams(ipaddress, numretries, password, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRegisterNetscalerControlCenterParams", reflect.TypeOf((*MockNetscalerServiceIface)(nil).NewRegisterNetscalerControlCenterParams), ipaddress, numretries, password, username)
}

// NewRegisterNetscalerServicePackageParams mocks base method.
func (m *MockNetscalerServiceIface) NewRegisterNetscalerServicePackageParams(description, name string) *RegisterNetscalerServicePackageParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRegisterNetscalerServicePackageParams", description, name)
	ret0, _ := ret[0].(*RegisterNetscalerServicePackageParams)
	return ret0
}

// NewRegisterNetscalerServicePackageParams indicates an expected call of NewRegisterNetscalerServicePackageParams.
func (mr *MockNetscalerServiceIfaceMockRecorder) NewRegisterNetscalerServicePackageParams(description, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRegisterNetscalerServicePackageParams", reflect.TypeOf((*MockNetscalerServiceIface)(nil).NewRegisterNetscalerServicePackageParams), description, name)
}

// RegisterNetscalerControlCenter mocks base method.
func (m *MockNetscalerServiceIface) RegisterNetscalerControlCenter(p *RegisterNetscalerControlCenterParams) (*RegisterNetscalerControlCenterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterNetscalerControlCenter", p)
	ret0, _ := ret[0].(*RegisterNetscalerControlCenterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterNetscalerControlCenter indicates an expected call of RegisterNetscalerControlCenter.
func (mr *MockNetscalerServiceIfaceMockRecorder) RegisterNetscalerControlCenter(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterNetscalerControlCenter", reflect.TypeOf((*MockNetscalerServiceIface)(nil).RegisterNetscalerControlCenter), p)
}

// RegisterNetscalerServicePackage mocks base method.
func (m *MockNetscalerServiceIface) RegisterNetscalerServicePackage(p *RegisterNetscalerServicePackageParams) (*RegisterNetscalerServicePackageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterNetscalerServicePackage", p)
	ret0, _ := ret[0].(*RegisterNetscalerServicePackageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterNetscalerServicePackage indicates an expected call of RegisterNetscalerServicePackage.
func (mr *MockNetscalerServiceIfaceMockRecorder) RegisterNetscalerServicePackage(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterNetscalerServicePackage", reflect.TypeOf((*MockNetscalerServiceIface)(nil).RegisterNetscalerServicePackage), p)
}
