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
// Source: ./cloudstack/SnapshotService.go
//
// Generated by this command:
//
//	mockgen -destination=./cloudstack/SnapshotService_mock.go -package=cloudstack -copyright_file=header.txt -source=./cloudstack/SnapshotService.go
//

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSnapshotServiceIface is a mock of SnapshotServiceIface interface.
type MockSnapshotServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotServiceIfaceMockRecorder
	isgomock struct{}
}

// MockSnapshotServiceIfaceMockRecorder is the mock recorder for MockSnapshotServiceIface.
type MockSnapshotServiceIfaceMockRecorder struct {
	mock *MockSnapshotServiceIface
}

// NewMockSnapshotServiceIface creates a new mock instance.
func NewMockSnapshotServiceIface(ctrl *gomock.Controller) *MockSnapshotServiceIface {
	mock := &MockSnapshotServiceIface{ctrl: ctrl}
	mock.recorder = &MockSnapshotServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSnapshotServiceIface) EXPECT() *MockSnapshotServiceIfaceMockRecorder {
	return m.recorder
}

// ArchiveSnapshot mocks base method.
func (m *MockSnapshotServiceIface) ArchiveSnapshot(p *ArchiveSnapshotParams) (*ArchiveSnapshotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ArchiveSnapshot", p)
	ret0, _ := ret[0].(*ArchiveSnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArchiveSnapshot indicates an expected call of ArchiveSnapshot.
func (mr *MockSnapshotServiceIfaceMockRecorder) ArchiveSnapshot(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArchiveSnapshot", reflect.TypeOf((*MockSnapshotServiceIface)(nil).ArchiveSnapshot), p)
}

// CopySnapshot mocks base method.
func (m *MockSnapshotServiceIface) CopySnapshot(p *CopySnapshotParams) (*CopySnapshotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopySnapshot", p)
	ret0, _ := ret[0].(*CopySnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CopySnapshot indicates an expected call of CopySnapshot.
func (mr *MockSnapshotServiceIfaceMockRecorder) CopySnapshot(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopySnapshot", reflect.TypeOf((*MockSnapshotServiceIface)(nil).CopySnapshot), p)
}

// CreateSnapshot mocks base method.
func (m *MockSnapshotServiceIface) CreateSnapshot(p *CreateSnapshotParams) (*CreateSnapshotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshot", p)
	ret0, _ := ret[0].(*CreateSnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshot indicates an expected call of CreateSnapshot.
func (mr *MockSnapshotServiceIfaceMockRecorder) CreateSnapshot(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshot", reflect.TypeOf((*MockSnapshotServiceIface)(nil).CreateSnapshot), p)
}

// CreateSnapshotFromVMSnapshot mocks base method.
func (m *MockSnapshotServiceIface) CreateSnapshotFromVMSnapshot(p *CreateSnapshotFromVMSnapshotParams) (*CreateSnapshotFromVMSnapshotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshotFromVMSnapshot", p)
	ret0, _ := ret[0].(*CreateSnapshotFromVMSnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshotFromVMSnapshot indicates an expected call of CreateSnapshotFromVMSnapshot.
func (mr *MockSnapshotServiceIfaceMockRecorder) CreateSnapshotFromVMSnapshot(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshotFromVMSnapshot", reflect.TypeOf((*MockSnapshotServiceIface)(nil).CreateSnapshotFromVMSnapshot), p)
}

// CreateSnapshotPolicy mocks base method.
func (m *MockSnapshotServiceIface) CreateSnapshotPolicy(p *CreateSnapshotPolicyParams) (*CreateSnapshotPolicyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshotPolicy", p)
	ret0, _ := ret[0].(*CreateSnapshotPolicyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshotPolicy indicates an expected call of CreateSnapshotPolicy.
func (mr *MockSnapshotServiceIfaceMockRecorder) CreateSnapshotPolicy(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshotPolicy", reflect.TypeOf((*MockSnapshotServiceIface)(nil).CreateSnapshotPolicy), p)
}

// CreateVMSnapshot mocks base method.
func (m *MockSnapshotServiceIface) CreateVMSnapshot(p *CreateVMSnapshotParams) (*CreateVMSnapshotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVMSnapshot", p)
	ret0, _ := ret[0].(*CreateVMSnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVMSnapshot indicates an expected call of CreateVMSnapshot.
func (mr *MockSnapshotServiceIfaceMockRecorder) CreateVMSnapshot(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVMSnapshot", reflect.TypeOf((*MockSnapshotServiceIface)(nil).CreateVMSnapshot), p)
}

// DeleteSnapshot mocks base method.
func (m *MockSnapshotServiceIface) DeleteSnapshot(p *DeleteSnapshotParams) (*DeleteSnapshotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", p)
	ret0, _ := ret[0].(*DeleteSnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot.
func (mr *MockSnapshotServiceIfaceMockRecorder) DeleteSnapshot(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockSnapshotServiceIface)(nil).DeleteSnapshot), p)
}

// DeleteSnapshotPolicies mocks base method.
func (m *MockSnapshotServiceIface) DeleteSnapshotPolicies(p *DeleteSnapshotPoliciesParams) (*DeleteSnapshotPoliciesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshotPolicies", p)
	ret0, _ := ret[0].(*DeleteSnapshotPoliciesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSnapshotPolicies indicates an expected call of DeleteSnapshotPolicies.
func (mr *MockSnapshotServiceIfaceMockRecorder) DeleteSnapshotPolicies(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshotPolicies", reflect.TypeOf((*MockSnapshotServiceIface)(nil).DeleteSnapshotPolicies), p)
}

// DeleteVMSnapshot mocks base method.
func (m *MockSnapshotServiceIface) DeleteVMSnapshot(p *DeleteVMSnapshotParams) (*DeleteVMSnapshotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVMSnapshot", p)
	ret0, _ := ret[0].(*DeleteVMSnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteVMSnapshot indicates an expected call of DeleteVMSnapshot.
func (mr *MockSnapshotServiceIfaceMockRecorder) DeleteVMSnapshot(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVMSnapshot", reflect.TypeOf((*MockSnapshotServiceIface)(nil).DeleteVMSnapshot), p)
}

// ExtractSnapshot mocks base method.
func (m *MockSnapshotServiceIface) ExtractSnapshot(p *ExtractSnapshotParams) (*ExtractSnapshotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractSnapshot", p)
	ret0, _ := ret[0].(*ExtractSnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtractSnapshot indicates an expected call of ExtractSnapshot.
func (mr *MockSnapshotServiceIfaceMockRecorder) ExtractSnapshot(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractSnapshot", reflect.TypeOf((*MockSnapshotServiceIface)(nil).ExtractSnapshot), p)
}

// GetSnapshotByID mocks base method.
func (m *MockSnapshotServiceIface) GetSnapshotByID(id string, opts ...OptionFunc) (*Snapshot, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSnapshotByID", varargs...)
	ret0, _ := ret[0].(*Snapshot)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSnapshotByID indicates an expected call of GetSnapshotByID.
func (mr *MockSnapshotServiceIfaceMockRecorder) GetSnapshotByID(id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshotByID", reflect.TypeOf((*MockSnapshotServiceIface)(nil).GetSnapshotByID), varargs...)
}

// GetSnapshotByName mocks base method.
func (m *MockSnapshotServiceIface) GetSnapshotByName(name string, opts ...OptionFunc) (*Snapshot, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSnapshotByName", varargs...)
	ret0, _ := ret[0].(*Snapshot)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSnapshotByName indicates an expected call of GetSnapshotByName.
func (mr *MockSnapshotServiceIfaceMockRecorder) GetSnapshotByName(name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshotByName", reflect.TypeOf((*MockSnapshotServiceIface)(nil).GetSnapshotByName), varargs...)
}

// GetSnapshotID mocks base method.
func (m *MockSnapshotServiceIface) GetSnapshotID(name string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSnapshotID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSnapshotID indicates an expected call of GetSnapshotID.
func (mr *MockSnapshotServiceIfaceMockRecorder) GetSnapshotID(name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshotID", reflect.TypeOf((*MockSnapshotServiceIface)(nil).GetSnapshotID), varargs...)
}

// GetSnapshotPolicyByID mocks base method.
func (m *MockSnapshotServiceIface) GetSnapshotPolicyByID(id string, opts ...OptionFunc) (*SnapshotPolicy, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSnapshotPolicyByID", varargs...)
	ret0, _ := ret[0].(*SnapshotPolicy)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSnapshotPolicyByID indicates an expected call of GetSnapshotPolicyByID.
func (mr *MockSnapshotServiceIfaceMockRecorder) GetSnapshotPolicyByID(id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshotPolicyByID", reflect.TypeOf((*MockSnapshotServiceIface)(nil).GetSnapshotPolicyByID), varargs...)
}

// GetVMSnapshotID mocks base method.
func (m *MockSnapshotServiceIface) GetVMSnapshotID(name string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVMSnapshotID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetVMSnapshotID indicates an expected call of GetVMSnapshotID.
func (mr *MockSnapshotServiceIfaceMockRecorder) GetVMSnapshotID(name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVMSnapshotID", reflect.TypeOf((*MockSnapshotServiceIface)(nil).GetVMSnapshotID), varargs...)
}

// ListSnapshotPolicies mocks base method.
func (m *MockSnapshotServiceIface) ListSnapshotPolicies(p *ListSnapshotPoliciesParams) (*ListSnapshotPoliciesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSnapshotPolicies", p)
	ret0, _ := ret[0].(*ListSnapshotPoliciesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSnapshotPolicies indicates an expected call of ListSnapshotPolicies.
func (mr *MockSnapshotServiceIfaceMockRecorder) ListSnapshotPolicies(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSnapshotPolicies", reflect.TypeOf((*MockSnapshotServiceIface)(nil).ListSnapshotPolicies), p)
}

// ListSnapshots mocks base method.
func (m *MockSnapshotServiceIface) ListSnapshots(p *ListSnapshotsParams) (*ListSnapshotsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSnapshots", p)
	ret0, _ := ret[0].(*ListSnapshotsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSnapshots indicates an expected call of ListSnapshots.
func (mr *MockSnapshotServiceIfaceMockRecorder) ListSnapshots(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSnapshots", reflect.TypeOf((*MockSnapshotServiceIface)(nil).ListSnapshots), p)
}

// ListVMSnapshot mocks base method.
func (m *MockSnapshotServiceIface) ListVMSnapshot(p *ListVMSnapshotParams) (*ListVMSnapshotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVMSnapshot", p)
	ret0, _ := ret[0].(*ListVMSnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVMSnapshot indicates an expected call of ListVMSnapshot.
func (mr *MockSnapshotServiceIfaceMockRecorder) ListVMSnapshot(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVMSnapshot", reflect.TypeOf((*MockSnapshotServiceIface)(nil).ListVMSnapshot), p)
}

// NewArchiveSnapshotParams mocks base method.
func (m *MockSnapshotServiceIface) NewArchiveSnapshotParams(id string) *ArchiveSnapshotParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewArchiveSnapshotParams", id)
	ret0, _ := ret[0].(*ArchiveSnapshotParams)
	return ret0
}

// NewArchiveSnapshotParams indicates an expected call of NewArchiveSnapshotParams.
func (mr *MockSnapshotServiceIfaceMockRecorder) NewArchiveSnapshotParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewArchiveSnapshotParams", reflect.TypeOf((*MockSnapshotServiceIface)(nil).NewArchiveSnapshotParams), id)
}

// NewCopySnapshotParams mocks base method.
func (m *MockSnapshotServiceIface) NewCopySnapshotParams(id string) *CopySnapshotParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCopySnapshotParams", id)
	ret0, _ := ret[0].(*CopySnapshotParams)
	return ret0
}

// NewCopySnapshotParams indicates an expected call of NewCopySnapshotParams.
func (mr *MockSnapshotServiceIfaceMockRecorder) NewCopySnapshotParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCopySnapshotParams", reflect.TypeOf((*MockSnapshotServiceIface)(nil).NewCopySnapshotParams), id)
}

// NewCreateSnapshotFromVMSnapshotParams mocks base method.
func (m *MockSnapshotServiceIface) NewCreateSnapshotFromVMSnapshotParams(vmsnapshotid, volumeid string) *CreateSnapshotFromVMSnapshotParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCreateSnapshotFromVMSnapshotParams", vmsnapshotid, volumeid)
	ret0, _ := ret[0].(*CreateSnapshotFromVMSnapshotParams)
	return ret0
}

// NewCreateSnapshotFromVMSnapshotParams indicates an expected call of NewCreateSnapshotFromVMSnapshotParams.
func (mr *MockSnapshotServiceIfaceMockRecorder) NewCreateSnapshotFromVMSnapshotParams(vmsnapshotid, volumeid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCreateSnapshotFromVMSnapshotParams", reflect.TypeOf((*MockSnapshotServiceIface)(nil).NewCreateSnapshotFromVMSnapshotParams), vmsnapshotid, volumeid)
}

// NewCreateSnapshotParams mocks base method.
func (m *MockSnapshotServiceIface) NewCreateSnapshotParams(volumeid string) *CreateSnapshotParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCreateSnapshotParams", volumeid)
	ret0, _ := ret[0].(*CreateSnapshotParams)
	return ret0
}

// NewCreateSnapshotParams indicates an expected call of NewCreateSnapshotParams.
func (mr *MockSnapshotServiceIfaceMockRecorder) NewCreateSnapshotParams(volumeid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCreateSnapshotParams", reflect.TypeOf((*MockSnapshotServiceIface)(nil).NewCreateSnapshotParams), volumeid)
}

// NewCreateSnapshotPolicyParams mocks base method.
func (m *MockSnapshotServiceIface) NewCreateSnapshotPolicyParams(intervaltype string, maxsnaps int, schedule, timezone, volumeid string) *CreateSnapshotPolicyParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCreateSnapshotPolicyParams", intervaltype, maxsnaps, schedule, timezone, volumeid)
	ret0, _ := ret[0].(*CreateSnapshotPolicyParams)
	return ret0
}

// NewCreateSnapshotPolicyParams indicates an expected call of NewCreateSnapshotPolicyParams.
func (mr *MockSnapshotServiceIfaceMockRecorder) NewCreateSnapshotPolicyParams(intervaltype, maxsnaps, schedule, timezone, volumeid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCreateSnapshotPolicyParams", reflect.TypeOf((*MockSnapshotServiceIface)(nil).NewCreateSnapshotPolicyParams), intervaltype, maxsnaps, schedule, timezone, volumeid)
}

// NewCreateVMSnapshotParams mocks base method.
func (m *MockSnapshotServiceIface) NewCreateVMSnapshotParams(virtualmachineid string) *CreateVMSnapshotParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCreateVMSnapshotParams", virtualmachineid)
	ret0, _ := ret[0].(*CreateVMSnapshotParams)
	return ret0
}

// NewCreateVMSnapshotParams indicates an expected call of NewCreateVMSnapshotParams.
func (mr *MockSnapshotServiceIfaceMockRecorder) NewCreateVMSnapshotParams(virtualmachineid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCreateVMSnapshotParams", reflect.TypeOf((*MockSnapshotServiceIface)(nil).NewCreateVMSnapshotParams), virtualmachineid)
}

// NewDeleteSnapshotParams mocks base method.
func (m *MockSnapshotServiceIface) NewDeleteSnapshotParams(id string) *DeleteSnapshotParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteSnapshotParams", id)
	ret0, _ := ret[0].(*DeleteSnapshotParams)
	return ret0
}

// NewDeleteSnapshotParams indicates an expected call of NewDeleteSnapshotParams.
func (mr *MockSnapshotServiceIfaceMockRecorder) NewDeleteSnapshotParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteSnapshotParams", reflect.TypeOf((*MockSnapshotServiceIface)(nil).NewDeleteSnapshotParams), id)
}

// NewDeleteSnapshotPoliciesParams mocks base method.
func (m *MockSnapshotServiceIface) NewDeleteSnapshotPoliciesParams() *DeleteSnapshotPoliciesParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteSnapshotPoliciesParams")
	ret0, _ := ret[0].(*DeleteSnapshotPoliciesParams)
	return ret0
}

// NewDeleteSnapshotPoliciesParams indicates an expected call of NewDeleteSnapshotPoliciesParams.
func (mr *MockSnapshotServiceIfaceMockRecorder) NewDeleteSnapshotPoliciesParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteSnapshotPoliciesParams", reflect.TypeOf((*MockSnapshotServiceIface)(nil).NewDeleteSnapshotPoliciesParams))
}

// NewDeleteVMSnapshotParams mocks base method.
func (m *MockSnapshotServiceIface) NewDeleteVMSnapshotParams(vmsnapshotid string) *DeleteVMSnapshotParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteVMSnapshotParams", vmsnapshotid)
	ret0, _ := ret[0].(*DeleteVMSnapshotParams)
	return ret0
}

// NewDeleteVMSnapshotParams indicates an expected call of NewDeleteVMSnapshotParams.
func (mr *MockSnapshotServiceIfaceMockRecorder) NewDeleteVMSnapshotParams(vmsnapshotid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteVMSnapshotParams", reflect.TypeOf((*MockSnapshotServiceIface)(nil).NewDeleteVMSnapshotParams), vmsnapshotid)
}

// NewExtractSnapshotParams mocks base method.
func (m *MockSnapshotServiceIface) NewExtractSnapshotParams(id, zoneid string) *ExtractSnapshotParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewExtractSnapshotParams", id, zoneid)
	ret0, _ := ret[0].(*ExtractSnapshotParams)
	return ret0
}

// NewExtractSnapshotParams indicates an expected call of NewExtractSnapshotParams.
func (mr *MockSnapshotServiceIfaceMockRecorder) NewExtractSnapshotParams(id, zoneid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewExtractSnapshotParams", reflect.TypeOf((*MockSnapshotServiceIface)(nil).NewExtractSnapshotParams), id, zoneid)
}

// NewListSnapshotPoliciesParams mocks base method.
func (m *MockSnapshotServiceIface) NewListSnapshotPoliciesParams() *ListSnapshotPoliciesParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListSnapshotPoliciesParams")
	ret0, _ := ret[0].(*ListSnapshotPoliciesParams)
	return ret0
}

// NewListSnapshotPoliciesParams indicates an expected call of NewListSnapshotPoliciesParams.
func (mr *MockSnapshotServiceIfaceMockRecorder) NewListSnapshotPoliciesParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListSnapshotPoliciesParams", reflect.TypeOf((*MockSnapshotServiceIface)(nil).NewListSnapshotPoliciesParams))
}

// NewListSnapshotsParams mocks base method.
func (m *MockSnapshotServiceIface) NewListSnapshotsParams() *ListSnapshotsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListSnapshotsParams")
	ret0, _ := ret[0].(*ListSnapshotsParams)
	return ret0
}

// NewListSnapshotsParams indicates an expected call of NewListSnapshotsParams.
func (mr *MockSnapshotServiceIfaceMockRecorder) NewListSnapshotsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListSnapshotsParams", reflect.TypeOf((*MockSnapshotServiceIface)(nil).NewListSnapshotsParams))
}

// NewListVMSnapshotParams mocks base method.
func (m *MockSnapshotServiceIface) NewListVMSnapshotParams() *ListVMSnapshotParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListVMSnapshotParams")
	ret0, _ := ret[0].(*ListVMSnapshotParams)
	return ret0
}

// NewListVMSnapshotParams indicates an expected call of NewListVMSnapshotParams.
func (mr *MockSnapshotServiceIfaceMockRecorder) NewListVMSnapshotParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListVMSnapshotParams", reflect.TypeOf((*MockSnapshotServiceIface)(nil).NewListVMSnapshotParams))
}

// NewRevertSnapshotParams mocks base method.
func (m *MockSnapshotServiceIface) NewRevertSnapshotParams(id string) *RevertSnapshotParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRevertSnapshotParams", id)
	ret0, _ := ret[0].(*RevertSnapshotParams)
	return ret0
}

// NewRevertSnapshotParams indicates an expected call of NewRevertSnapshotParams.
func (mr *MockSnapshotServiceIfaceMockRecorder) NewRevertSnapshotParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRevertSnapshotParams", reflect.TypeOf((*MockSnapshotServiceIface)(nil).NewRevertSnapshotParams), id)
}

// NewRevertToVMSnapshotParams mocks base method.
func (m *MockSnapshotServiceIface) NewRevertToVMSnapshotParams(vmsnapshotid string) *RevertToVMSnapshotParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRevertToVMSnapshotParams", vmsnapshotid)
	ret0, _ := ret[0].(*RevertToVMSnapshotParams)
	return ret0
}

// NewRevertToVMSnapshotParams indicates an expected call of NewRevertToVMSnapshotParams.
func (mr *MockSnapshotServiceIfaceMockRecorder) NewRevertToVMSnapshotParams(vmsnapshotid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRevertToVMSnapshotParams", reflect.TypeOf((*MockSnapshotServiceIface)(nil).NewRevertToVMSnapshotParams), vmsnapshotid)
}

// NewUpdateSnapshotPolicyParams mocks base method.
func (m *MockSnapshotServiceIface) NewUpdateSnapshotPolicyParams() *UpdateSnapshotPolicyParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateSnapshotPolicyParams")
	ret0, _ := ret[0].(*UpdateSnapshotPolicyParams)
	return ret0
}

// NewUpdateSnapshotPolicyParams indicates an expected call of NewUpdateSnapshotPolicyParams.
func (mr *MockSnapshotServiceIfaceMockRecorder) NewUpdateSnapshotPolicyParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateSnapshotPolicyParams", reflect.TypeOf((*MockSnapshotServiceIface)(nil).NewUpdateSnapshotPolicyParams))
}

// RevertSnapshot mocks base method.
func (m *MockSnapshotServiceIface) RevertSnapshot(p *RevertSnapshotParams) (*RevertSnapshotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevertSnapshot", p)
	ret0, _ := ret[0].(*RevertSnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevertSnapshot indicates an expected call of RevertSnapshot.
func (mr *MockSnapshotServiceIfaceMockRecorder) RevertSnapshot(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevertSnapshot", reflect.TypeOf((*MockSnapshotServiceIface)(nil).RevertSnapshot), p)
}

// RevertToVMSnapshot mocks base method.
func (m *MockSnapshotServiceIface) RevertToVMSnapshot(p *RevertToVMSnapshotParams) (*RevertToVMSnapshotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevertToVMSnapshot", p)
	ret0, _ := ret[0].(*RevertToVMSnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevertToVMSnapshot indicates an expected call of RevertToVMSnapshot.
func (mr *MockSnapshotServiceIfaceMockRecorder) RevertToVMSnapshot(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevertToVMSnapshot", reflect.TypeOf((*MockSnapshotServiceIface)(nil).RevertToVMSnapshot), p)
}

// UpdateSnapshotPolicy mocks base method.
func (m *MockSnapshotServiceIface) UpdateSnapshotPolicy(p *UpdateSnapshotPolicyParams) (*UpdateSnapshotPolicyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSnapshotPolicy", p)
	ret0, _ := ret[0].(*UpdateSnapshotPolicyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSnapshotPolicy indicates an expected call of UpdateSnapshotPolicy.
func (mr *MockSnapshotServiceIfaceMockRecorder) UpdateSnapshotPolicy(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSnapshotPolicy", reflect.TypeOf((*MockSnapshotServiceIface)(nil).UpdateSnapshotPolicy), p)
}
