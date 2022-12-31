/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mocks

import (
	reflect "reflect"
	sync "sync"

	consts "github.com/apache/incubator-eventmesh/eventmesh-server-go/runtime/consts"
	consumer "github.com/apache/incubator-eventmesh/eventmesh-server-go/runtime/core/protocol/grpc/consumer"
	pb "github.com/apache/incubator-eventmesh/eventmesh-server-go/runtime/proto/pb"
	gomock "github.com/golang/mock/gomock"
	set "github.com/liyue201/gostl/ds/set"
)

// MockConsumerGroupTopicOption is a mock of ConsumerGroupTopicOption interface.
type MockConsumerGroupTopicOption struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerGroupTopicOptionMockRecorder
}

// MockConsumerGroupTopicOptionMockRecorder is the mock recorder for MockConsumerGroupTopicOption.
type MockConsumerGroupTopicOptionMockRecorder struct {
	mock *MockConsumerGroupTopicOption
}

// NewMockConsumerGroupTopicOption creates a new mock instance.
func NewMockConsumerGroupTopicOption(ctrl *gomock.Controller) *MockConsumerGroupTopicOption {
	mock := &MockConsumerGroupTopicOption{ctrl: ctrl}
	mock.recorder = &MockConsumerGroupTopicOptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumerGroupTopicOption) EXPECT() *MockConsumerGroupTopicOptionMockRecorder {
	return m.recorder
}

// AllEmiters mocks base method.
func (m *MockConsumerGroupTopicOption) AllEmiters() *set.Set {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllEmiters")
	ret0, _ := ret[0].(*set.Set)
	return ret0
}

// AllEmiters indicates an expected call of AllEmiters.
func (mr *MockConsumerGroupTopicOptionMockRecorder) AllEmiters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllEmiters", reflect.TypeOf((*MockConsumerGroupTopicOption)(nil).AllEmiters))
}

// AllURLs mocks base method.
func (m *MockConsumerGroupTopicOption) AllURLs() *set.Set {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllURLs")
	ret0, _ := ret[0].(*set.Set)
	return ret0
}

// AllURLs indicates an expected call of AllURLs.
func (mr *MockConsumerGroupTopicOptionMockRecorder) AllURLs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllURLs", reflect.TypeOf((*MockConsumerGroupTopicOption)(nil).AllURLs))
}

// ConsumerGroup mocks base method.
func (m *MockConsumerGroupTopicOption) ConsumerGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumerGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ConsumerGroup indicates an expected call of ConsumerGroup.
func (mr *MockConsumerGroupTopicOptionMockRecorder) ConsumerGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumerGroup", reflect.TypeOf((*MockConsumerGroupTopicOption)(nil).ConsumerGroup))
}

// DeregisterClient mocks base method.
func (m *MockConsumerGroupTopicOption) DeregisterClient() consumer.DeregisterClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterClient")
	ret0, _ := ret[0].(consumer.DeregisterClient)
	return ret0
}

// DeregisterClient indicates an expected call of DeregisterClient.
func (mr *MockConsumerGroupTopicOptionMockRecorder) DeregisterClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterClient", reflect.TypeOf((*MockConsumerGroupTopicOption)(nil).DeregisterClient))
}

// GRPCType mocks base method.
func (m *MockConsumerGroupTopicOption) GRPCType() consts.GRPCType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GRPCType")
	ret0, _ := ret[0].(consts.GRPCType)
	return ret0
}

// GRPCType indicates an expected call of GRPCType.
func (mr *MockConsumerGroupTopicOptionMockRecorder) GRPCType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GRPCType", reflect.TypeOf((*MockConsumerGroupTopicOption)(nil).GRPCType))
}

// IDCEmiters mocks base method.
func (m *MockConsumerGroupTopicOption) IDCEmiters() *sync.Map {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IDCEmiters")
	ret0, _ := ret[0].(*sync.Map)
	return ret0
}

// IDCEmiters indicates an expected call of IDCEmiters.
func (mr *MockConsumerGroupTopicOptionMockRecorder) IDCEmiters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IDCEmiters", reflect.TypeOf((*MockConsumerGroupTopicOption)(nil).IDCEmiters))
}

// IDCURLs mocks base method.
func (m *MockConsumerGroupTopicOption) IDCURLs() *sync.Map {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IDCURLs")
	ret0, _ := ret[0].(*sync.Map)
	return ret0
}

// IDCURLs indicates an expected call of IDCURLs.
func (mr *MockConsumerGroupTopicOptionMockRecorder) IDCURLs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IDCURLs", reflect.TypeOf((*MockConsumerGroupTopicOption)(nil).IDCURLs))
}

// RegisterClient mocks base method.
func (m *MockConsumerGroupTopicOption) RegisterClient() consumer.RegisterClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterClient")
	ret0, _ := ret[0].(consumer.RegisterClient)
	return ret0
}

// RegisterClient indicates an expected call of RegisterClient.
func (mr *MockConsumerGroupTopicOptionMockRecorder) RegisterClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterClient", reflect.TypeOf((*MockConsumerGroupTopicOption)(nil).RegisterClient))
}

// Size mocks base method.
func (m *MockConsumerGroupTopicOption) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockConsumerGroupTopicOptionMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockConsumerGroupTopicOption)(nil).Size))
}

// SubscriptionMode mocks base method.
func (m *MockConsumerGroupTopicOption) SubscriptionMode() pb.Subscription_SubscriptionItem_SubscriptionMode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionMode")
	ret0, _ := ret[0].(pb.Subscription_SubscriptionItem_SubscriptionMode)
	return ret0
}

// SubscriptionMode indicates an expected call of SubscriptionMode.
func (mr *MockConsumerGroupTopicOptionMockRecorder) SubscriptionMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionMode", reflect.TypeOf((*MockConsumerGroupTopicOption)(nil).SubscriptionMode))
}

// Topic mocks base method.
func (m *MockConsumerGroupTopicOption) Topic() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Topic")
	ret0, _ := ret[0].(string)
	return ret0
}

// Topic indicates an expected call of Topic.
func (mr *MockConsumerGroupTopicOptionMockRecorder) Topic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Topic", reflect.TypeOf((*MockConsumerGroupTopicOption)(nil).Topic))
}