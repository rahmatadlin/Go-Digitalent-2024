// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/internal/model"
)

// CommentService is an autogenerated mock type for the CommentService type
type CommentService struct {
	mock.Mock
}

// DeleteComment provides a mock function with given fields: ctx, commentId
func (_m *CommentService) DeleteComment(ctx context.Context, commentId uint32) error {
	ret := _m.Called(ctx, commentId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteComment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) error); ok {
		r0 = rf(ctx, commentId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllCommentsByPhotoId provides a mock function with given fields: ctx, photoId
func (_m *CommentService) GetAllCommentsByPhotoId(ctx context.Context, photoId uint32) ([]model.CommentView, error) {
	ret := _m.Called(ctx, photoId)

	if len(ret) == 0 {
		panic("no return value specified for GetAllCommentsByPhotoId")
	}

	var r0 []model.CommentView
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) ([]model.CommentView, error)); ok {
		return rf(ctx, photoId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []model.CommentView); ok {
		r0 = rf(ctx, photoId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.CommentView)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, photoId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommentById provides a mock function with given fields: ctx, commentId
func (_m *CommentService) GetCommentById(ctx context.Context, commentId uint32) (*model.Comment, error) {
	ret := _m.Called(ctx, commentId)

	if len(ret) == 0 {
		panic("no return value specified for GetCommentById")
	}

	var r0 *model.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) (*model.Comment, error)); ok {
		return rf(ctx, commentId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) *model.Comment); ok {
		r0 = rf(ctx, commentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Comment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, commentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostComment provides a mock function with given fields: ctx, userId, newComment
func (_m *CommentService) PostComment(ctx context.Context, userId uint32, newComment model.CreateComment) (*model.CreateCommentRes, error) {
	ret := _m.Called(ctx, userId, newComment)

	if len(ret) == 0 {
		panic("no return value specified for PostComment")
	}

	var r0 *model.CreateCommentRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, model.CreateComment) (*model.CreateCommentRes, error)); ok {
		return rf(ctx, userId, newComment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, model.CreateComment) *model.CreateCommentRes); ok {
		r0 = rf(ctx, userId, newComment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CreateCommentRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, model.CreateComment) error); ok {
		r1 = rf(ctx, userId, newComment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateComment provides a mock function with given fields: ctx, comment
func (_m *CommentService) UpdateComment(ctx context.Context, comment model.Comment) (*model.UpdateCommentRes, error) {
	ret := _m.Called(ctx, comment)

	if len(ret) == 0 {
		panic("no return value specified for UpdateComment")
	}

	var r0 *model.UpdateCommentRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Comment) (*model.UpdateCommentRes, error)); ok {
		return rf(ctx, comment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Comment) *model.UpdateCommentRes); ok {
		r0 = rf(ctx, comment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UpdateCommentRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Comment) error); ok {
		r1 = rf(ctx, comment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCommentService creates a new instance of CommentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCommentService(t interface {
	mock.TestingT
	Cleanup(func())
}) *CommentService {
	mock := &CommentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}