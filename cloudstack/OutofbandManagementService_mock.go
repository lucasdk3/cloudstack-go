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
// Source: ./cloudstack/OutofbandManagementService.go
//
// Generated by this command:
//
//	mockgen -destination=./cloudstack/OutofbandManagementService_mock.go -package=cloudstack -copyright_file=header.txt -source=./cloudstack/OutofbandManagementService.go
//

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockOutofbandManagementServiceIface is a mock of OutofbandManagementServiceIface interface.
type MockOutofbandManagementServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockOutofbandManagementServiceIfaceMockRecorder
	isgomock struct{}
}

// MockOutofbandManagementServiceIfaceMockRecorder is the mock recorder for MockOutofbandManagementServiceIface.
type MockOutofbandManagementServiceIfaceMockRecorder struct {
	mock *MockOutofbandManagementServiceIface
}

// NewMockOutofbandManagementServiceIface creates a new mock instance.
func NewMockOutofbandManagementServiceIface(ctrl *gomock.Controller) *MockOutofbandManagementServiceIface {
	mock := &MockOutofbandManagementServiceIface{ctrl: ctrl}
	mock.recorder = &MockOutofbandManagementServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOutofbandManagementServiceIface) EXPECT() *MockOutofbandManagementServiceIfaceMockRecorder {
	return m.recorder
}

// ChangeOutOfBandManagementPassword mocks base method.
func (m *MockOutofbandManagementServiceIface) ChangeOutOfBandManagementPassword(p *ChangeOutOfBandManagementPasswordParams) (*ChangeOutOfBandManagementPasswordResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeOutOfBandManagementPassword", p)
	ret0, _ := ret[0].(*ChangeOutOfBandManagementPasswordResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeOutOfBandManagementPassword indicates an expected call of ChangeOutOfBandManagementPassword.
func (mr *MockOutofbandManagementServiceIfaceMockRecorder) ChangeOutOfBandManagementPassword(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeOutOfBandManagementPassword", reflect.TypeOf((*MockOutofbandManagementServiceIface)(nil).ChangeOutOfBandManagementPassword), p)
}

// ConfigureOutOfBandManagement mocks base method.
func (m *MockOutofbandManagementServiceIface) ConfigureOutOfBandManagement(p *ConfigureOutOfBandManagementParams) (*OutOfBandManagementResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureOutOfBandManagement", p)
	ret0, _ := ret[0].(*OutOfBandManagementResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigureOutOfBandManagement indicates an expected call of ConfigureOutOfBandManagement.
func (mr *MockOutofbandManagementServiceIfaceMockRecorder) ConfigureOutOfBandManagement(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureOutOfBandManagement", reflect.TypeOf((*MockOutofbandManagementServiceIface)(nil).ConfigureOutOfBandManagement), p)
}

// IssueOutOfBandManagementPowerAction mocks base method.
func (m *MockOutofbandManagementServiceIface) IssueOutOfBandManagementPowerAction(p *IssueOutOfBandManagementPowerActionParams) (*IssueOutOfBandManagementPowerActionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueOutOfBandManagementPowerAction", p)
	ret0, _ := ret[0].(*IssueOutOfBandManagementPowerActionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IssueOutOfBandManagementPowerAction indicates an expected call of IssueOutOfBandManagementPowerAction.
func (mr *MockOutofbandManagementServiceIfaceMockRecorder) IssueOutOfBandManagementPowerAction(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueOutOfBandManagementPowerAction", reflect.TypeOf((*MockOutofbandManagementServiceIface)(nil).IssueOutOfBandManagementPowerAction), p)
}

// NewChangeOutOfBandManagementPasswordParams mocks base method.
func (m *MockOutofbandManagementServiceIface) NewChangeOutOfBandManagementPasswordParams(hostid string) *ChangeOutOfBandManagementPasswordParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewChangeOutOfBandManagementPasswordParams", hostid)
	ret0, _ := ret[0].(*ChangeOutOfBandManagementPasswordParams)
	return ret0
}

// NewChangeOutOfBandManagementPasswordParams indicates an expected call of NewChangeOutOfBandManagementPasswordParams.
func (mr *MockOutofbandManagementServiceIfaceMockRecorder) NewChangeOutOfBandManagementPasswordParams(hostid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewChangeOutOfBandManagementPasswordParams", reflect.TypeOf((*MockOutofbandManagementServiceIface)(nil).NewChangeOutOfBandManagementPasswordParams), hostid)
}

// NewConfigureOutOfBandManagementParams mocks base method.
func (m *MockOutofbandManagementServiceIface) NewConfigureOutOfBandManagementParams(address, driver, hostid, password, port, username string) *ConfigureOutOfBandManagementParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewConfigureOutOfBandManagementParams", address, driver, hostid, password, port, username)
	ret0, _ := ret[0].(*ConfigureOutOfBandManagementParams)
	return ret0
}

// NewConfigureOutOfBandManagementParams indicates an expected call of NewConfigureOutOfBandManagementParams.
func (mr *MockOutofbandManagementServiceIfaceMockRecorder) NewConfigureOutOfBandManagementParams(address, driver, hostid, password, port, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewConfigureOutOfBandManagementParams", reflect.TypeOf((*MockOutofbandManagementServiceIface)(nil).NewConfigureOutOfBandManagementParams), address, driver, hostid, password, port, username)
}

// NewIssueOutOfBandManagementPowerActionParams mocks base method.
func (m *MockOutofbandManagementServiceIface) NewIssueOutOfBandManagementPowerActionParams(action, hostid string) *IssueOutOfBandManagementPowerActionParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewIssueOutOfBandManagementPowerActionParams", action, hostid)
	ret0, _ := ret[0].(*IssueOutOfBandManagementPowerActionParams)
	return ret0
}

// NewIssueOutOfBandManagementPowerActionParams indicates an expected call of NewIssueOutOfBandManagementPowerActionParams.
func (mr *MockOutofbandManagementServiceIfaceMockRecorder) NewIssueOutOfBandManagementPowerActionParams(action, hostid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewIssueOutOfBandManagementPowerActionParams", reflect.TypeOf((*MockOutofbandManagementServiceIface)(nil).NewIssueOutOfBandManagementPowerActionParams), action, hostid)
}
