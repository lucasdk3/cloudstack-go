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
// Source: ./cloudstack/ResourceService.go
//
// Generated by this command:
//
//	mockgen -destination=./cloudstack/ResourceService_mock.go -package=cloudstack -copyright_file=header.txt -source=./cloudstack/ResourceService.go
//

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockResourceServiceIface is a mock of ResourceServiceIface interface.
type MockResourceServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockResourceServiceIfaceMockRecorder
	isgomock struct{}
}

// MockResourceServiceIfaceMockRecorder is the mock recorder for MockResourceServiceIface.
type MockResourceServiceIfaceMockRecorder struct {
	mock *MockResourceServiceIface
}

// NewMockResourceServiceIface creates a new mock instance.
func NewMockResourceServiceIface(ctrl *gomock.Controller) *MockResourceServiceIface {
	mock := &MockResourceServiceIface{ctrl: ctrl}
	mock.recorder = &MockResourceServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceServiceIface) EXPECT() *MockResourceServiceIfaceMockRecorder {
	return m.recorder
}

// NewPurgeExpungedResourcesParams mocks base method.
func (m *MockResourceServiceIface) NewPurgeExpungedResourcesParams() *PurgeExpungedResourcesParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewPurgeExpungedResourcesParams")
	ret0, _ := ret[0].(*PurgeExpungedResourcesParams)
	return ret0
}

// NewPurgeExpungedResourcesParams indicates an expected call of NewPurgeExpungedResourcesParams.
func (mr *MockResourceServiceIfaceMockRecorder) NewPurgeExpungedResourcesParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPurgeExpungedResourcesParams", reflect.TypeOf((*MockResourceServiceIface)(nil).NewPurgeExpungedResourcesParams))
}

// PurgeExpungedResources mocks base method.
func (m *MockResourceServiceIface) PurgeExpungedResources(p *PurgeExpungedResourcesParams) (*PurgeExpungedResourcesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PurgeExpungedResources", p)
	ret0, _ := ret[0].(*PurgeExpungedResourcesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PurgeExpungedResources indicates an expected call of PurgeExpungedResources.
func (mr *MockResourceServiceIfaceMockRecorder) PurgeExpungedResources(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurgeExpungedResources", reflect.TypeOf((*MockResourceServiceIface)(nil).PurgeExpungedResources), p)
}
