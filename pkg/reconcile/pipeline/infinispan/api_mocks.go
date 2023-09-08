// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/reconcile/pipeline/infinispan/api.go

// Package infinispan is a generated GoMock package.
package infinispan

import (
	context "context"
	reflect "reflect"
	time "time"

	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/infinispan/infinispan-operator/api/v1"
	api "github.com/infinispan/infinispan-operator/pkg/infinispan/client/api"
	version "github.com/infinispan/infinispan-operator/pkg/infinispan/version"
	kubernetes "github.com/infinispan/infinispan-operator/pkg/kubernetes"
	v10 "k8s.io/api/core/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	record "k8s.io/client-go/tools/record"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockPipeline is a mock of Pipeline interface.
type MockPipeline struct {
	ctrl     *gomock.Controller
	recorder *MockPipelineMockRecorder
}

// MockPipelineMockRecorder is the mock recorder for MockPipeline.
type MockPipelineMockRecorder struct {
	mock *MockPipeline
}

// NewMockPipeline creates a new mock instance.
func NewMockPipeline(ctrl *gomock.Controller) *MockPipeline {
	mock := &MockPipeline{ctrl: ctrl}
	mock.recorder = &MockPipelineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipeline) EXPECT() *MockPipelineMockRecorder {
	return m.recorder
}

// Process mocks base method.
func (m *MockPipeline) Process(ctx context.Context) (bool, time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(time.Duration)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Process indicates an expected call of Process.
func (mr *MockPipelineMockRecorder) Process(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockPipeline)(nil).Process), ctx)
}

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockHandler) Handle(i *v1.Infinispan, ctx Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handle", i, ctx)
}

// Handle indicates an expected call of Handle.
func (mr *MockHandlerMockRecorder) Handle(i, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockHandler)(nil).Handle), i, ctx)
}

// MockContextProvider is a mock of ContextProvider interface.
type MockContextProvider struct {
	ctrl     *gomock.Controller
	recorder *MockContextProviderMockRecorder
}

// MockContextProviderMockRecorder is the mock recorder for MockContextProvider.
type MockContextProviderMockRecorder struct {
	mock *MockContextProvider
}

// NewMockContextProvider creates a new mock instance.
func NewMockContextProvider(ctrl *gomock.Controller) *MockContextProvider {
	mock := &MockContextProvider{ctrl: ctrl}
	mock.recorder = &MockContextProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContextProvider) EXPECT() *MockContextProviderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockContextProvider) Get(ctx context.Context, config *ContextProviderConfig) (Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, config)
	ret0, _ := ret[0].(Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockContextProviderMockRecorder) Get(ctx, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockContextProvider)(nil).Get), ctx, config)
}

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// ConfigFiles mocks base method.
func (m *MockContext) ConfigFiles() *ConfigFiles {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigFiles")
	ret0, _ := ret[0].(*ConfigFiles)
	return ret0
}

// ConfigFiles indicates an expected call of ConfigFiles.
func (mr *MockContextMockRecorder) ConfigFiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigFiles", reflect.TypeOf((*MockContext)(nil).ConfigFiles))
}

// Ctx mocks base method.
func (m *MockContext) Ctx() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ctx")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Ctx indicates an expected call of Ctx.
func (mr *MockContextMockRecorder) Ctx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ctx", reflect.TypeOf((*MockContext)(nil).Ctx))
}

// DefaultAnnotations mocks base method.
func (m *MockContext) DefaultAnnotations() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultAnnotations")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// DefaultAnnotations indicates an expected call of DefaultAnnotations.
func (mr *MockContextMockRecorder) DefaultAnnotations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultAnnotations", reflect.TypeOf((*MockContext)(nil).DefaultAnnotations))
}

// DefaultLabels mocks base method.
func (m *MockContext) DefaultLabels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultLabels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// DefaultLabels indicates an expected call of DefaultLabels.
func (mr *MockContextMockRecorder) DefaultLabels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultLabels", reflect.TypeOf((*MockContext)(nil).DefaultLabels))
}

// EventRecorder mocks base method.
func (m *MockContext) EventRecorder() record.EventRecorder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventRecorder")
	ret0, _ := ret[0].(record.EventRecorder)
	return ret0
}

// EventRecorder indicates an expected call of EventRecorder.
func (mr *MockContextMockRecorder) EventRecorder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventRecorder", reflect.TypeOf((*MockContext)(nil).EventRecorder))
}

// FlowStatus mocks base method.
func (m *MockContext) FlowStatus() FlowStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlowStatus")
	ret0, _ := ret[0].(FlowStatus)
	return ret0
}

// FlowStatus indicates an expected call of FlowStatus.
func (mr *MockContextMockRecorder) FlowStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlowStatus", reflect.TypeOf((*MockContext)(nil).FlowStatus))
}

// InfinispanClient mocks base method.
func (m *MockContext) InfinispanClient() (api.Infinispan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InfinispanClient")
	ret0, _ := ret[0].(api.Infinispan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InfinispanClient indicates an expected call of InfinispanClient.
func (mr *MockContextMockRecorder) InfinispanClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InfinispanClient", reflect.TypeOf((*MockContext)(nil).InfinispanClient))
}

// InfinispanClientForPod mocks base method.
func (m *MockContext) InfinispanClientForPod(podName string) api.Infinispan {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InfinispanClientForPod", podName)
	ret0, _ := ret[0].(api.Infinispan)
	return ret0
}

// InfinispanClientForPod indicates an expected call of InfinispanClientForPod.
func (mr *MockContextMockRecorder) InfinispanClientForPod(podName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InfinispanClientForPod", reflect.TypeOf((*MockContext)(nil).InfinispanClientForPod), podName)
}

// InfinispanPods mocks base method.
func (m *MockContext) InfinispanPods() (*v10.PodList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InfinispanPods")
	ret0, _ := ret[0].(*v10.PodList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InfinispanPods indicates an expected call of InfinispanPods.
func (mr *MockContextMockRecorder) InfinispanPods() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InfinispanPods", reflect.TypeOf((*MockContext)(nil).InfinispanPods))
}

// IsTypeSupported mocks base method.
func (m *MockContext) IsTypeSupported(gvk schema.GroupVersionKind) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTypeSupported", gvk)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTypeSupported indicates an expected call of IsTypeSupported.
func (mr *MockContextMockRecorder) IsTypeSupported(gvk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTypeSupported", reflect.TypeOf((*MockContext)(nil).IsTypeSupported), gvk)
}

// Kubernetes mocks base method.
func (m *MockContext) Kubernetes() *kubernetes.Kubernetes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kubernetes")
	ret0, _ := ret[0].(*kubernetes.Kubernetes)
	return ret0
}

// Kubernetes indicates an expected call of Kubernetes.
func (mr *MockContextMockRecorder) Kubernetes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kubernetes", reflect.TypeOf((*MockContext)(nil).Kubernetes))
}

// Log mocks base method.
func (m *MockContext) Log() logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Log")
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// Log indicates an expected call of Log.
func (mr *MockContextMockRecorder) Log() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockContext)(nil).Log))
}

// Operand mocks base method.
func (m *MockContext) Operand() version.Operand {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Operand")
	ret0, _ := ret[0].(version.Operand)
	return ret0
}

// Operand indicates an expected call of Operand.
func (mr *MockContextMockRecorder) Operand() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Operand", reflect.TypeOf((*MockContext)(nil).Operand))
}

// Operands mocks base method.
func (m *MockContext) Operands() *version.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Operands")
	ret0, _ := ret[0].(*version.Manager)
	return ret0
}

// Operands indicates an expected call of Operands.
func (mr *MockContextMockRecorder) Operands() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Operands", reflect.TypeOf((*MockContext)(nil).Operands))
}

// Requeue mocks base method.
func (m *MockContext) Requeue(reason error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Requeue", reason)
}

// Requeue indicates an expected call of Requeue.
func (mr *MockContextMockRecorder) Requeue(reason interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Requeue", reflect.TypeOf((*MockContext)(nil).Requeue), reason)
}

// RequeueAfter mocks base method.
func (m *MockContext) RequeueAfter(delay time.Duration, reason error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RequeueAfter", delay, reason)
}

// RequeueAfter indicates an expected call of RequeueAfter.
func (mr *MockContextMockRecorder) RequeueAfter(delay, reason interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequeueAfter", reflect.TypeOf((*MockContext)(nil).RequeueAfter), delay, reason)
}

// RequeueEventually mocks base method.
func (m *MockContext) RequeueEventually(delay time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RequeueEventually", delay)
}

// RequeueEventually indicates an expected call of RequeueEventually.
func (mr *MockContextMockRecorder) RequeueEventually(delay interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequeueEventually", reflect.TypeOf((*MockContext)(nil).RequeueEventually), delay)
}

// Resources mocks base method.
func (m *MockContext) Resources() Resources {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resources")
	ret0, _ := ret[0].(Resources)
	return ret0
}

// Resources indicates an expected call of Resources.
func (mr *MockContextMockRecorder) Resources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resources", reflect.TypeOf((*MockContext)(nil).Resources))
}

// Stop mocks base method.
func (m *MockContext) Stop(err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", err)
}

// Stop indicates an expected call of Stop.
func (mr *MockContextMockRecorder) Stop(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockContext)(nil).Stop), err)
}

// UpdateInfinispan mocks base method.
func (m *MockContext) UpdateInfinispan(arg0 func()) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInfinispan", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInfinispan indicates an expected call of UpdateInfinispan.
func (mr *MockContextMockRecorder) UpdateInfinispan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInfinispan", reflect.TypeOf((*MockContext)(nil).UpdateInfinispan), arg0)
}

// MockResources is a mock of Resources interface.
type MockResources struct {
	ctrl     *gomock.Controller
	recorder *MockResourcesMockRecorder
}

// MockResourcesMockRecorder is the mock recorder for MockResources.
type MockResourcesMockRecorder struct {
	mock *MockResources
}

// NewMockResources creates a new mock instance.
func NewMockResources(ctrl *gomock.Controller) *MockResources {
	mock := &MockResources{ctrl: ctrl}
	mock.recorder = &MockResourcesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResources) EXPECT() *MockResourcesMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockResources) Create(obj client.Object, setControllerRef bool, opts ...func(*ResourcesConfig)) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{obj, setControllerRef}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockResourcesMockRecorder) Create(obj, setControllerRef interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{obj, setControllerRef}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockResources)(nil).Create), varargs...)
}

// CreateOrPatch mocks base method.
func (m *MockResources) CreateOrPatch(obj client.Object, setControllerRef bool, mutate func() error, opts ...func(*ResourcesConfig)) (OperationResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{obj, setControllerRef, mutate}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrPatch", varargs...)
	ret0, _ := ret[0].(OperationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrPatch indicates an expected call of CreateOrPatch.
func (mr *MockResourcesMockRecorder) CreateOrPatch(obj, setControllerRef, mutate interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{obj, setControllerRef, mutate}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrPatch", reflect.TypeOf((*MockResources)(nil).CreateOrPatch), varargs...)
}

// CreateOrUpdate mocks base method.
func (m *MockResources) CreateOrUpdate(obj client.Object, setControllerRef bool, mutate func() error, opts ...func(*ResourcesConfig)) (OperationResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{obj, setControllerRef, mutate}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrUpdate", varargs...)
	ret0, _ := ret[0].(OperationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockResourcesMockRecorder) CreateOrUpdate(obj, setControllerRef, mutate interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{obj, setControllerRef, mutate}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockResources)(nil).CreateOrUpdate), varargs...)
}

// Delete mocks base method.
func (m *MockResources) Delete(name string, obj client.Object, opts ...func(*ResourcesConfig)) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockResourcesMockRecorder) Delete(name, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockResources)(nil).Delete), varargs...)
}

// List mocks base method.
func (m *MockResources) List(set map[string]string, list client.ObjectList, opts ...func(*ResourcesConfig)) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{set, list}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockResourcesMockRecorder) List(set, list interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{set, list}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockResources)(nil).List), varargs...)
}

// Load mocks base method.
func (m *MockResources) Load(name string, obj client.Object, opts ...func(*ResourcesConfig)) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Load", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockResourcesMockRecorder) Load(name, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockResources)(nil).Load), varargs...)
}

// LoadGlobal mocks base method.
func (m *MockResources) LoadGlobal(name string, obj client.Object, opts ...func(*ResourcesConfig)) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LoadGlobal", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadGlobal indicates an expected call of LoadGlobal.
func (mr *MockResourcesMockRecorder) LoadGlobal(name, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadGlobal", reflect.TypeOf((*MockResources)(nil).LoadGlobal), varargs...)
}

// SetControllerReference mocks base method.
func (m *MockResources) SetControllerReference(controlled v11.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetControllerReference", controlled)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetControllerReference indicates an expected call of SetControllerReference.
func (mr *MockResourcesMockRecorder) SetControllerReference(controlled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetControllerReference", reflect.TypeOf((*MockResources)(nil).SetControllerReference), controlled)
}

// Update mocks base method.
func (m *MockResources) Update(obj client.Object, opts ...func(*ResourcesConfig)) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockResourcesMockRecorder) Update(obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockResources)(nil).Update), varargs...)
}