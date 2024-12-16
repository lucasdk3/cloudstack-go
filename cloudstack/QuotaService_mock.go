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
// Source: ./cloudstack/QuotaService.go
//
// Generated by this command:
//
//	mockgen -destination=./cloudstack/QuotaService_mock.go -package=cloudstack -copyright_file=header.txt -source=./cloudstack/QuotaService.go
//

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockQuotaServiceIface is a mock of QuotaServiceIface interface.
type MockQuotaServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockQuotaServiceIfaceMockRecorder
	isgomock struct{}
}

// MockQuotaServiceIfaceMockRecorder is the mock recorder for MockQuotaServiceIface.
type MockQuotaServiceIfaceMockRecorder struct {
	mock *MockQuotaServiceIface
}

// NewMockQuotaServiceIface creates a new mock instance.
func NewMockQuotaServiceIface(ctrl *gomock.Controller) *MockQuotaServiceIface {
	mock := &MockQuotaServiceIface{ctrl: ctrl}
	mock.recorder = &MockQuotaServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuotaServiceIface) EXPECT() *MockQuotaServiceIfaceMockRecorder {
	return m.recorder
}

// NewQuotaIsEnabledParams mocks base method.
func (m *MockQuotaServiceIface) NewQuotaIsEnabledParams() *QuotaIsEnabledParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewQuotaIsEnabledParams")
	ret0, _ := ret[0].(*QuotaIsEnabledParams)
	return ret0
}

// NewQuotaIsEnabledParams indicates an expected call of NewQuotaIsEnabledParams.
func (mr *MockQuotaServiceIfaceMockRecorder) NewQuotaIsEnabledParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewQuotaIsEnabledParams", reflect.TypeOf((*MockQuotaServiceIface)(nil).NewQuotaIsEnabledParams))
}

// QuotaIsEnabled mocks base method.
func (m *MockQuotaServiceIface) QuotaIsEnabled(p *QuotaIsEnabledParams) (*QuotaIsEnabledResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuotaIsEnabled", p)
	ret0, _ := ret[0].(*QuotaIsEnabledResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuotaIsEnabled indicates an expected call of QuotaIsEnabled.
func (mr *MockQuotaServiceIfaceMockRecorder) QuotaIsEnabled(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuotaIsEnabled", reflect.TypeOf((*MockQuotaServiceIface)(nil).QuotaIsEnabled), p)
}
