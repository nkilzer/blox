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
// Source: ./pkg/facade/ecs.go

package mocks

import (
	ecs "github.com/aws/aws-sdk-go/service/ecs"
	gomock "github.com/golang/mock/gomock"
)

// Mock of ECS interface
type MockECS struct {
	ctrl     *gomock.Controller
	recorder *_MockECSRecorder
}

// Recorder for MockECS (not exported)
type _MockECSRecorder struct {
	mock *MockECS
}

func NewMockECS(ctrl *gomock.Controller) *MockECS {
	mock := &MockECS{ctrl: ctrl}
	mock.recorder = &_MockECSRecorder{mock}
	return mock
}

func (_m *MockECS) EXPECT() *_MockECSRecorder {
	return _m.recorder
}

func (_m *MockECS) StartTask(clusterArn string, containerInstances []*string, startedBy string, taskDefinition string) (*ecs.StartTaskOutput, error) {
	ret := _m.ctrl.Call(_m, "StartTask", clusterArn, containerInstances, startedBy, taskDefinition)
	ret0, _ := ret[0].(*ecs.StartTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSRecorder) StartTask(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartTask", arg0, arg1, arg2, arg3)
}

func (_m *MockECS) ListClusters() ([]*string, error) {
	ret := _m.ctrl.Call(_m, "ListClusters")
	ret0, _ := ret[0].([]*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSRecorder) ListClusters() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListClusters")
}

func (_m *MockECS) DescribeCluster(cluster *string) (*ecs.Cluster, error) {
	ret := _m.ctrl.Call(_m, "DescribeCluster", cluster)
	ret0, _ := ret[0].(*ecs.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSRecorder) DescribeCluster(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeCluster", arg0)
}

func (_m *MockECS) DescribeTaskDefinition(taskDefinition *string) (*ecs.TaskDefinition, error) {
	ret := _m.ctrl.Call(_m, "DescribeTaskDefinition", taskDefinition)
	ret0, _ := ret[0].(*ecs.TaskDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSRecorder) DescribeTaskDefinition(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeTaskDefinition", arg0)
}

func (_m *MockECS) ListTasks(cluster string, startedBy string) ([]*string, error) {
	ret := _m.ctrl.Call(_m, "ListTasks", cluster, startedBy)
	ret0, _ := ret[0].([]*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSRecorder) ListTasks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTasks", arg0, arg1)
}

func (_m *MockECS) ListTasksByInstance(cluster string, instanceARN string) ([]*string, error) {
	ret := _m.ctrl.Call(_m, "ListTasksByInstance", cluster, instanceARN)
	ret0, _ := ret[0].([]*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSRecorder) ListTasksByInstance(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTasksByInstance", arg0, arg1)
}

func (_m *MockECS) DescribeTasks(cluster string, tasks []*string) (*ecs.DescribeTasksOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeTasks", cluster, tasks)
	ret0, _ := ret[0].(*ecs.DescribeTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSRecorder) DescribeTasks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeTasks", arg0, arg1)
}

func (_m *MockECS) StopTask(clusterArn string, taskArn string) error {
	ret := _m.ctrl.Call(_m, "StopTask", clusterArn, taskArn)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockECSRecorder) StopTask(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopTask", arg0, arg1)
}
