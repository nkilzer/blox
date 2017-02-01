// Copyright 2016-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/blox/blox/cluster-state-service/handler/store (interfaces: DataStore)

package mocks

import (
	context "context"

	clientv3 "github.com/coreos/etcd/clientv3"
	concurrency "github.com/coreos/etcd/clientv3/concurrency"
	gomock "github.com/golang/mock/gomock"
)

// Mock of DataStore interface
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *_MockDataStoreRecorder
}

// Recorder for MockDataStore (not exported)
type _MockDataStoreRecorder struct {
	mock *MockDataStore
}

func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &_MockDataStoreRecorder{mock}
	return mock
}

func (_m *MockDataStore) EXPECT() *_MockDataStoreRecorder {
	return _m.recorder
}

func (_m *MockDataStore) Add(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "Add", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDataStoreRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Add", arg0, arg1)
}

func (_m *MockDataStore) Delete(_param0 string) (int64, error) {
	ret := _m.ctrl.Call(_m, "Delete", _param0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDataStoreRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}

func (_m *MockDataStore) Get(_param0 string) (map[string]string, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDataStoreRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockDataStore) GetV3Client() *clientv3.Client {
	ret := _m.ctrl.Call(_m, "GetV3Client")
	ret0, _ := ret[0].(*clientv3.Client)
	return ret0
}

func (_mr *_MockDataStoreRecorder) GetV3Client() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetV3Client")
}

func (_m *MockDataStore) GetWithPrefix(_param0 string) (map[string]string, error) {
	ret := _m.ctrl.Call(_m, "GetWithPrefix", _param0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDataStoreRecorder) GetWithPrefix(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetWithPrefix", arg0)
}

func (_m *MockDataStore) NewSTMRepeatable(_param0 context.Context, _param1 *clientv3.Client, _param2 func(concurrency.STM) error) (*clientv3.TxnResponse, error) {
	ret := _m.ctrl.Call(_m, "NewSTMRepeatable", _param0, _param1, _param2)
	ret0, _ := ret[0].(*clientv3.TxnResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDataStoreRecorder) NewSTMRepeatable(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewSTMRepeatable", arg0, arg1, arg2)
}

func (_m *MockDataStore) StreamWithPrefix(_param0 context.Context, _param1 string) (chan map[string]string, error) {
	ret := _m.ctrl.Call(_m, "StreamWithPrefix", _param0, _param1)
	ret0, _ := ret[0].(chan map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDataStoreRecorder) StreamWithPrefix(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StreamWithPrefix", arg0, arg1)
}
