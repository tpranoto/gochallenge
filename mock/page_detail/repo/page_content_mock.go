// Code generated by MockGen. DO NOT EDIT.
// Source: page_detail/repo/page_content.go

package mock_repo

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockPageContent is a mock of PageContent interface
type MockPageContent struct {
	ctrl     *gomock.Controller
	recorder *MockPageContentMockRecorder
}

// MockPageContentMockRecorder is the mock recorder for MockPageContent
type MockPageContentMockRecorder struct {
	mock *MockPageContent
}

// NewMockPageContent creates a new mock instance
func NewMockPageContent(ctrl *gomock.Controller) *MockPageContent {
	mock := &MockPageContent{ctrl: ctrl}
	mock.recorder = &MockPageContentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockPageContent) EXPECT() *MockPageContentMockRecorder {
	return _m.recorder
}

// GetContentFromURL mocks base method
func (_m *MockPageContent) GetContentFromURL(ctx context.Context, url string) (*http.Response, error) {
	ret := _m.ctrl.Call(_m, "GetContentFromURL", ctx, url)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContentFromURL indicates an expected call of GetContentFromURL
func (_mr *MockPageContentMockRecorder) GetContentFromURL(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetContentFromURL", reflect.TypeOf((*MockPageContent)(nil).GetContentFromURL), arg0, arg1)
}