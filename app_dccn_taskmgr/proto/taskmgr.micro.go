// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/taskmgr.proto

/*
Package taskmgr is a generated protocol buffer package.

It is generated from these files:
	proto/taskmgr.proto

It has these top-level messages:
	Event
	AddTaskRequest
	AddTaskResponse
	TaskListRequest
	TaskInfo
	TaskListResponse
	CancelTaskRequest
	CancelTaskResponse
	UpdateTaskRequest
	UpdateTaskResponse
	TaskDetailRequest
	TaskDetailResponse
	Task
	TaskResult
*/
package taskmgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for TaskMgr service

type TaskMgrService interface {
	AddTask(ctx context.Context, in *AddTaskRequest, opts ...client.CallOption) (*AddTaskResponse, error)
	CancelTask(ctx context.Context, in *CancelTaskRequest, opts ...client.CallOption) (*CancelTaskResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...client.CallOption) (*UpdateTaskResponse, error)
	TaskList(ctx context.Context, in *TaskListRequest, opts ...client.CallOption) (*TaskListResponse, error)
	TaskDetail(ctx context.Context, in *TaskDetailRequest, opts ...client.CallOption) (*TaskDetailResponse, error)
}

type taskMgrService struct {
	c    client.Client
	name string
}

func NewTaskMgrService(name string, c client.Client) TaskMgrService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "taskmgr"
	}
	return &taskMgrService{
		c:    c,
		name: name,
	}
}

func (c *taskMgrService) AddTask(ctx context.Context, in *AddTaskRequest, opts ...client.CallOption) (*AddTaskResponse, error) {
	req := c.c.NewRequest(c.name, "TaskMgr.AddTask", in)
	out := new(AddTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrService) CancelTask(ctx context.Context, in *CancelTaskRequest, opts ...client.CallOption) (*CancelTaskResponse, error) {
	req := c.c.NewRequest(c.name, "TaskMgr.CancelTask", in)
	out := new(CancelTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrService) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...client.CallOption) (*UpdateTaskResponse, error) {
	req := c.c.NewRequest(c.name, "TaskMgr.UpdateTask", in)
	out := new(UpdateTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrService) TaskList(ctx context.Context, in *TaskListRequest, opts ...client.CallOption) (*TaskListResponse, error) {
	req := c.c.NewRequest(c.name, "TaskMgr.TaskList", in)
	out := new(TaskListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrService) TaskDetail(ctx context.Context, in *TaskDetailRequest, opts ...client.CallOption) (*TaskDetailResponse, error) {
	req := c.c.NewRequest(c.name, "TaskMgr.TaskDetail", in)
	out := new(TaskDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskMgr service

type TaskMgrHandler interface {
	AddTask(context.Context, *AddTaskRequest, *AddTaskResponse) error
	CancelTask(context.Context, *CancelTaskRequest, *CancelTaskResponse) error
	UpdateTask(context.Context, *UpdateTaskRequest, *UpdateTaskResponse) error
	TaskList(context.Context, *TaskListRequest, *TaskListResponse) error
	TaskDetail(context.Context, *TaskDetailRequest, *TaskDetailResponse) error
}

func RegisterTaskMgrHandler(s server.Server, hdlr TaskMgrHandler, opts ...server.HandlerOption) error {
	type taskMgr interface {
		AddTask(ctx context.Context, in *AddTaskRequest, out *AddTaskResponse) error
		CancelTask(ctx context.Context, in *CancelTaskRequest, out *CancelTaskResponse) error
		UpdateTask(ctx context.Context, in *UpdateTaskRequest, out *UpdateTaskResponse) error
		TaskList(ctx context.Context, in *TaskListRequest, out *TaskListResponse) error
		TaskDetail(ctx context.Context, in *TaskDetailRequest, out *TaskDetailResponse) error
	}
	type TaskMgr struct {
		taskMgr
	}
	h := &taskMgrHandler{hdlr}
	return s.Handle(s.NewHandler(&TaskMgr{h}, opts...))
}

type taskMgrHandler struct {
	TaskMgrHandler
}

func (h *taskMgrHandler) AddTask(ctx context.Context, in *AddTaskRequest, out *AddTaskResponse) error {
	return h.TaskMgrHandler.AddTask(ctx, in, out)
}

func (h *taskMgrHandler) CancelTask(ctx context.Context, in *CancelTaskRequest, out *CancelTaskResponse) error {
	return h.TaskMgrHandler.CancelTask(ctx, in, out)
}

func (h *taskMgrHandler) UpdateTask(ctx context.Context, in *UpdateTaskRequest, out *UpdateTaskResponse) error {
	return h.TaskMgrHandler.UpdateTask(ctx, in, out)
}

func (h *taskMgrHandler) TaskList(ctx context.Context, in *TaskListRequest, out *TaskListResponse) error {
	return h.TaskMgrHandler.TaskList(ctx, in, out)
}

func (h *taskMgrHandler) TaskDetail(ctx context.Context, in *TaskDetailRequest, out *TaskDetailResponse) error {
	return h.TaskMgrHandler.TaskDetail(ctx, in, out)
}
