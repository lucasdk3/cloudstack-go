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
// Source: ./cloudstack/NicService.go
//
// Generated by this command:
//
//	mockgen -destination=./cloudstack/NicService_mock.go -package=cloudstack -copyright_file=header.txt -source=./cloudstack/NicService.go
//

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockNicServiceIface is a mock of NicServiceIface interface.
type MockNicServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockNicServiceIfaceMockRecorder
	isgomock struct{}
}

// MockNicServiceIfaceMockRecorder is the mock recorder for MockNicServiceIface.
type MockNicServiceIfaceMockRecorder struct {
	mock *MockNicServiceIface
}

// NewMockNicServiceIface creates a new mock instance.
func NewMockNicServiceIface(ctrl *gomock.Controller) *MockNicServiceIface {
	mock := &MockNicServiceIface{ctrl: ctrl}
	mock.recorder = &MockNicServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNicServiceIface) EXPECT() *MockNicServiceIfaceMockRecorder {
	return m.recorder
}

// AddIpToNic mocks base method.
func (m *MockNicServiceIface) AddIpToNic(p *AddIpToNicParams) (*AddIpToNicResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddIpToNic", p)
	ret0, _ := ret[0].(*AddIpToNicResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddIpToNic indicates an expected call of AddIpToNic.
func (mr *MockNicServiceIfaceMockRecorder) AddIpToNic(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIpToNic", reflect.TypeOf((*MockNicServiceIface)(nil).AddIpToNic), p)
}

// ListNics mocks base method.
func (m *MockNicServiceIface) ListNics(p *ListNicsParams) (*ListNicsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNics", p)
	ret0, _ := ret[0].(*ListNicsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNics indicates an expected call of ListNics.
func (mr *MockNicServiceIfaceMockRecorder) ListNics(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNics", reflect.TypeOf((*MockNicServiceIface)(nil).ListNics), p)
}

// NewAddIpToNicParams mocks base method.
func (m *MockNicServiceIface) NewAddIpToNicParams(nicid string) *AddIpToNicParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAddIpToNicParams", nicid)
	ret0, _ := ret[0].(*AddIpToNicParams)
	return ret0
}

// NewAddIpToNicParams indicates an expected call of NewAddIpToNicParams.
func (mr *MockNicServiceIfaceMockRecorder) NewAddIpToNicParams(nicid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAddIpToNicParams", reflect.TypeOf((*MockNicServiceIface)(nil).NewAddIpToNicParams), nicid)
}

// NewListNicsParams mocks base method.
func (m *MockNicServiceIface) NewListNicsParams(virtualmachineid string) *ListNicsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListNicsParams", virtualmachineid)
	ret0, _ := ret[0].(*ListNicsParams)
	return ret0
}

// NewListNicsParams indicates an expected call of NewListNicsParams.
func (mr *MockNicServiceIfaceMockRecorder) NewListNicsParams(virtualmachineid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListNicsParams", reflect.TypeOf((*MockNicServiceIface)(nil).NewListNicsParams), virtualmachineid)
}

// NewRemoveIpFromNicParams mocks base method.
func (m *MockNicServiceIface) NewRemoveIpFromNicParams(id string) *RemoveIpFromNicParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRemoveIpFromNicParams", id)
	ret0, _ := ret[0].(*RemoveIpFromNicParams)
	return ret0
}

// NewRemoveIpFromNicParams indicates an expected call of NewRemoveIpFromNicParams.
func (mr *MockNicServiceIfaceMockRecorder) NewRemoveIpFromNicParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRemoveIpFromNicParams", reflect.TypeOf((*MockNicServiceIface)(nil).NewRemoveIpFromNicParams), id)
}

// NewUpdateVmNicIpParams mocks base method.
func (m *MockNicServiceIface) NewUpdateVmNicIpParams(nicid string) *UpdateVmNicIpParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateVmNicIpParams", nicid)
	ret0, _ := ret[0].(*UpdateVmNicIpParams)
	return ret0
}

// NewUpdateVmNicIpParams indicates an expected call of NewUpdateVmNicIpParams.
func (mr *MockNicServiceIfaceMockRecorder) NewUpdateVmNicIpParams(nicid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateVmNicIpParams", reflect.TypeOf((*MockNicServiceIface)(nil).NewUpdateVmNicIpParams), nicid)
}

// RemoveIpFromNic mocks base method.
func (m *MockNicServiceIface) RemoveIpFromNic(p *RemoveIpFromNicParams) (*RemoveIpFromNicResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveIpFromNic", p)
	ret0, _ := ret[0].(*RemoveIpFromNicResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveIpFromNic indicates an expected call of RemoveIpFromNic.
func (mr *MockNicServiceIfaceMockRecorder) RemoveIpFromNic(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIpFromNic", reflect.TypeOf((*MockNicServiceIface)(nil).RemoveIpFromNic), p)
}

// UpdateVmNicIp mocks base method.
func (m *MockNicServiceIface) UpdateVmNicIp(p *UpdateVmNicIpParams) (*UpdateVmNicIpResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVmNicIp", p)
	ret0, _ := ret[0].(*UpdateVmNicIpResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateVmNicIp indicates an expected call of UpdateVmNicIp.
func (mr *MockNicServiceIfaceMockRecorder) UpdateVmNicIp(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVmNicIp", reflect.TypeOf((*MockNicServiceIface)(nil).UpdateVmNicIp), p)
}
