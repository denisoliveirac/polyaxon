// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/run.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Run kind enum
type RunKind int32

const (
	RunKind_job      RunKind = 0
	RunKind_service  RunKind = 1
	RunKind_dag      RunKind = 2
	RunKind_parallel RunKind = 3
)

var RunKind_name = map[int32]string{
	0: "job",
	1: "service",
	2: "dag",
	3: "parallel",
}

var RunKind_value = map[string]int32{
	"job":      0,
	"service":  1,
	"dag":      2,
	"parallel": 3,
}

func (x RunKind) String() string {
	return proto.EnumName(RunKind_name, int32(x))
}

func (RunKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{0}
}

// Run meta info specification
type RunMetaInfo struct {
	// Optional flag to tell if the run has a service
	Service bool `protobuf:"varint,1,opt,name=service,proto3" json:"service,omitempty"`
	// Optional an indicator if the run has a concurrency
	Concurrency int32 `protobuf:"varint,2,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	// Optional the parallel kind
	ParallelKind string `protobuf:"bytes,3,opt,name=parallel_kind,json=parallelKind,proto3" json:"parallel_kind,omitempty"`
	// Optional the run kind
	RunKind              string   `protobuf:"bytes,4,opt,name=run_kind,json=runKind,proto3" json:"run_kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunMetaInfo) Reset()         { *m = RunMetaInfo{} }
func (m *RunMetaInfo) String() string { return proto.CompactTextString(m) }
func (*RunMetaInfo) ProtoMessage()    {}
func (*RunMetaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{0}
}

func (m *RunMetaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunMetaInfo.Unmarshal(m, b)
}
func (m *RunMetaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunMetaInfo.Marshal(b, m, deterministic)
}
func (m *RunMetaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunMetaInfo.Merge(m, src)
}
func (m *RunMetaInfo) XXX_Size() int {
	return xxx_messageInfo_RunMetaInfo.Size(m)
}
func (m *RunMetaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RunMetaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RunMetaInfo proto.InternalMessageInfo

func (m *RunMetaInfo) GetService() bool {
	if m != nil {
		return m.Service
	}
	return false
}

func (m *RunMetaInfo) GetConcurrency() int32 {
	if m != nil {
		return m.Concurrency
	}
	return 0
}

func (m *RunMetaInfo) GetParallelKind() string {
	if m != nil {
		return m.ParallelKind
	}
	return ""
}

func (m *RunMetaInfo) GetRunKind() string {
	if m != nil {
		return m.RunKind
	}
	return ""
}

// Run specification
type Run struct {
	// UUID
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Optional name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Optional description
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Optional Tags of this entity
	Tags []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	// Optional if the entity has been deleted
	Deleted bool `protobuf:"varint,5,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// Required name of user started this entity
	User string `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	// Required name of owner of this entity
	Owner string `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
	// Required project name
	Project string `protobuf:"bytes,8,opt,name=project,proto3" json:"project,omitempty"`
	// Optional time when the entityt was created
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Optional last time the entity was updated
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Optional last time the entity was started
	StartedAt *timestamp.Timestamp `protobuf:"bytes,11,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// Optional last time the entity was started
	FinishedAt *timestamp.Timestamp `protobuf:"bytes,12,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	// Optional flag to tell if this entity is managed by the platform
	IsManaged string `protobuf:"bytes,13,opt,name=is_managed,json=isManaged,proto3" json:"is_managed,omitempty"`
	// Optional content of the entity's spec
	Content string `protobuf:"bytes,14,opt,name=content,proto3" json:"content,omitempty"`
	// Optional latest status of this entity
	Status string `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	// Optional a readme text describing this entity
	Readme string `protobuf:"bytes,16,opt,name=readme,proto3" json:"readme,omitempty"`
	// Optional if this entity was bookmarked
	Bookmarked bool `protobuf:"varint,17,opt,name=bookmarked,proto3" json:"bookmarked,omitempty"`
	// Optional run meta info
	MetaInfo *RunMetaInfo `protobuf:"bytes,18,opt,name=meta_info,json=metaInfo,proto3" json:"meta_info,omitempty"`
	// Optional kind to tell the nature of this run
	Kind RunKind `protobuf:"varint,19,opt,name=kind,proto3,enum=v1.RunKind" json:"kind,omitempty"`
	// Optional hub reference versioned component
	Hub string `protobuf:"bytes,20,opt,name=hub,proto3" json:"hub,omitempty"`
	// Optional inputs of this entity
	Inputs *_struct.Struct `protobuf:"bytes,21,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// Optional outputs of this entity
	Outputs *_struct.Struct `protobuf:"bytes,22,opt,name=outputs,proto3" json:"outputs,omitempty"`
	// Optional run environment tracked
	RunEnv *_struct.Struct `protobuf:"bytes,23,opt,name=run_env,json=runEnv,proto3" json:"run_env,omitempty"`
	// Is resume
	IsResume bool `protobuf:"varint,24,opt,name=is_resume,json=isResume,proto3" json:"is_resume,omitempty"`
	// Is clone
	IsClone bool `protobuf:"varint,25,opt,name=is_clone,json=isClone,proto3" json:"is_clone,omitempty"`
	// Optional if this run was restarted/copied/resumed/cached
	CloningStrategy string `protobuf:"bytes,26,opt,name=cloning_strategy,json=cloningStrategy,proto3" json:"cloning_strategy,omitempty"`
	// Optional uuid of the pipeline
	Pipeline string `protobuf:"bytes,27,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	// Optional uuid of the original run
	Original string `protobuf:"bytes,28,opt,name=original,proto3" json:"original,omitempty"`
	// Optional name of the pipeline
	PipelineName string `protobuf:"bytes,29,opt,name=pipeline_name,json=pipelineName,proto3" json:"pipeline_name,omitempty"`
	// Optional name of the original run
	OriginalName         string   `protobuf:"bytes,30,opt,name=original_name,json=originalName,proto3" json:"original_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Run) Reset()         { *m = Run{} }
func (m *Run) String() string { return proto.CompactTextString(m) }
func (*Run) ProtoMessage()    {}
func (*Run) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{1}
}

func (m *Run) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Run.Unmarshal(m, b)
}
func (m *Run) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Run.Marshal(b, m, deterministic)
}
func (m *Run) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Run.Merge(m, src)
}
func (m *Run) XXX_Size() int {
	return xxx_messageInfo_Run.Size(m)
}
func (m *Run) XXX_DiscardUnknown() {
	xxx_messageInfo_Run.DiscardUnknown(m)
}

var xxx_messageInfo_Run proto.InternalMessageInfo

func (m *Run) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Run) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Run) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Run) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Run) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *Run) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Run) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Run) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Run) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Run) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Run) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *Run) GetFinishedAt() *timestamp.Timestamp {
	if m != nil {
		return m.FinishedAt
	}
	return nil
}

func (m *Run) GetIsManaged() string {
	if m != nil {
		return m.IsManaged
	}
	return ""
}

func (m *Run) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Run) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Run) GetReadme() string {
	if m != nil {
		return m.Readme
	}
	return ""
}

func (m *Run) GetBookmarked() bool {
	if m != nil {
		return m.Bookmarked
	}
	return false
}

func (m *Run) GetMetaInfo() *RunMetaInfo {
	if m != nil {
		return m.MetaInfo
	}
	return nil
}

func (m *Run) GetKind() RunKind {
	if m != nil {
		return m.Kind
	}
	return RunKind_job
}

func (m *Run) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}

func (m *Run) GetInputs() *_struct.Struct {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Run) GetOutputs() *_struct.Struct {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *Run) GetRunEnv() *_struct.Struct {
	if m != nil {
		return m.RunEnv
	}
	return nil
}

func (m *Run) GetIsResume() bool {
	if m != nil {
		return m.IsResume
	}
	return false
}

func (m *Run) GetIsClone() bool {
	if m != nil {
		return m.IsClone
	}
	return false
}

func (m *Run) GetCloningStrategy() string {
	if m != nil {
		return m.CloningStrategy
	}
	return ""
}

func (m *Run) GetPipeline() string {
	if m != nil {
		return m.Pipeline
	}
	return ""
}

func (m *Run) GetOriginal() string {
	if m != nil {
		return m.Original
	}
	return ""
}

func (m *Run) GetPipelineName() string {
	if m != nil {
		return m.PipelineName
	}
	return ""
}

func (m *Run) GetOriginalName() string {
	if m != nil {
		return m.OriginalName
	}
	return ""
}

// Request data to create run
type RunBodyRequest struct {
	// Owner of the namespace
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Project where the experiement will be assigned
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Run object
	Run                  *Run     `protobuf:"bytes,3,opt,name=run,proto3" json:"run,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunBodyRequest) Reset()         { *m = RunBodyRequest{} }
func (m *RunBodyRequest) String() string { return proto.CompactTextString(m) }
func (*RunBodyRequest) ProtoMessage()    {}
func (*RunBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{2}
}

func (m *RunBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunBodyRequest.Unmarshal(m, b)
}
func (m *RunBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunBodyRequest.Marshal(b, m, deterministic)
}
func (m *RunBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunBodyRequest.Merge(m, src)
}
func (m *RunBodyRequest) XXX_Size() int {
	return xxx_messageInfo_RunBodyRequest.Size(m)
}
func (m *RunBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunBodyRequest proto.InternalMessageInfo

func (m *RunBodyRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RunBodyRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *RunBodyRequest) GetRun() *Run {
	if m != nil {
		return m.Run
	}
	return nil
}

// Request data to update run
type EntityRunBodyRequest struct {
	// Enitity Run
	Entity *ProjectEntityResourceRequest `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	// Run object
	Run                  *Run     `protobuf:"bytes,2,opt,name=run,proto3" json:"run,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityRunBodyRequest) Reset()         { *m = EntityRunBodyRequest{} }
func (m *EntityRunBodyRequest) String() string { return proto.CompactTextString(m) }
func (*EntityRunBodyRequest) ProtoMessage()    {}
func (*EntityRunBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{3}
}

func (m *EntityRunBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityRunBodyRequest.Unmarshal(m, b)
}
func (m *EntityRunBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityRunBodyRequest.Marshal(b, m, deterministic)
}
func (m *EntityRunBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityRunBodyRequest.Merge(m, src)
}
func (m *EntityRunBodyRequest) XXX_Size() int {
	return xxx_messageInfo_EntityRunBodyRequest.Size(m)
}
func (m *EntityRunBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityRunBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EntityRunBodyRequest proto.InternalMessageInfo

func (m *EntityRunBodyRequest) GetEntity() *ProjectEntityResourceRequest {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *EntityRunBodyRequest) GetRun() *Run {
	if m != nil {
		return m.Run
	}
	return nil
}

// Contains list runs
type ListRunsResponse struct {
	// Count of the entities
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// List of all entities
	Results []*Run `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	// Previous page
	Previous string `protobuf:"bytes,3,opt,name=previous,proto3" json:"previous,omitempty"`
	// Next page
	Next                 string   `protobuf:"bytes,4,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRunsResponse) Reset()         { *m = ListRunsResponse{} }
func (m *ListRunsResponse) String() string { return proto.CompactTextString(m) }
func (*ListRunsResponse) ProtoMessage()    {}
func (*ListRunsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{4}
}

func (m *ListRunsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRunsResponse.Unmarshal(m, b)
}
func (m *ListRunsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRunsResponse.Marshal(b, m, deterministic)
}
func (m *ListRunsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRunsResponse.Merge(m, src)
}
func (m *ListRunsResponse) XXX_Size() int {
	return xxx_messageInfo_ListRunsResponse.Size(m)
}
func (m *ListRunsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRunsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRunsResponse proto.InternalMessageInfo

func (m *ListRunsResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListRunsResponse) GetResults() []*Run {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ListRunsResponse) GetPrevious() string {
	if m != nil {
		return m.Previous
	}
	return ""
}

func (m *ListRunsResponse) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

// Run Settings specification
type RunSettingsCatalog struct {
	// Uuid
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Name
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunSettingsCatalog) Reset()         { *m = RunSettingsCatalog{} }
func (m *RunSettingsCatalog) String() string { return proto.CompactTextString(m) }
func (*RunSettingsCatalog) ProtoMessage()    {}
func (*RunSettingsCatalog) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{5}
}

func (m *RunSettingsCatalog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunSettingsCatalog.Unmarshal(m, b)
}
func (m *RunSettingsCatalog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunSettingsCatalog.Marshal(b, m, deterministic)
}
func (m *RunSettingsCatalog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunSettingsCatalog.Merge(m, src)
}
func (m *RunSettingsCatalog) XXX_Size() int {
	return xxx_messageInfo_RunSettingsCatalog.Size(m)
}
func (m *RunSettingsCatalog) XXX_DiscardUnknown() {
	xxx_messageInfo_RunSettingsCatalog.DiscardUnknown(m)
}

var xxx_messageInfo_RunSettingsCatalog proto.InternalMessageInfo

func (m *RunSettingsCatalog) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *RunSettingsCatalog) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RunSettings struct {
	// Namespace
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Agent
	Agent *RunSettingsCatalog `protobuf:"bytes,2,opt,name=agent,proto3" json:"agent,omitempty"`
	// Queue
	Queue *RunSettingsCatalog `protobuf:"bytes,3,opt,name=queue,proto3" json:"queue,omitempty"`
	// Logs Store
	LogsStore *RunSettingsCatalog `protobuf:"bytes,4,opt,name=logs_store,json=logsStore,proto3" json:"logs_store,omitempty"`
	// Outputs Store
	OutputsStore *RunSettingsCatalog `protobuf:"bytes,5,opt,name=outputs_store,json=outputsStore,proto3" json:"outputs_store,omitempty"`
	// Init connections
	InitConnections []*RunSettingsCatalog `protobuf:"bytes,6,rep,name=init_connections,json=initConnections,proto3" json:"init_connections,omitempty"`
	// Connections
	Connections []*RunSettingsCatalog `protobuf:"bytes,7,rep,name=connections,proto3" json:"connections,omitempty"`
	// Git Accesses
	GitAccesses []*RunSettingsCatalog `protobuf:"bytes,8,rep,name=git_accesses,json=gitAccesses,proto3" json:"git_accesses,omitempty"`
	// Registry Access
	RegistryAccess *RunSettingsCatalog `protobuf:"bytes,9,opt,name=registry_access,json=registryAccess,proto3" json:"registry_access,omitempty"`
	// K8S config maps
	ConfigResources      []*RunSettingsCatalog `protobuf:"bytes,10,rep,name=config_resources,json=configResources,proto3" json:"config_resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RunSettings) Reset()         { *m = RunSettings{} }
func (m *RunSettings) String() string { return proto.CompactTextString(m) }
func (*RunSettings) ProtoMessage()    {}
func (*RunSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{6}
}

func (m *RunSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunSettings.Unmarshal(m, b)
}
func (m *RunSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunSettings.Marshal(b, m, deterministic)
}
func (m *RunSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunSettings.Merge(m, src)
}
func (m *RunSettings) XXX_Size() int {
	return xxx_messageInfo_RunSettings.Size(m)
}
func (m *RunSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_RunSettings.DiscardUnknown(m)
}

var xxx_messageInfo_RunSettings proto.InternalMessageInfo

func (m *RunSettings) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *RunSettings) GetAgent() *RunSettingsCatalog {
	if m != nil {
		return m.Agent
	}
	return nil
}

func (m *RunSettings) GetQueue() *RunSettingsCatalog {
	if m != nil {
		return m.Queue
	}
	return nil
}

func (m *RunSettings) GetLogsStore() *RunSettingsCatalog {
	if m != nil {
		return m.LogsStore
	}
	return nil
}

func (m *RunSettings) GetOutputsStore() *RunSettingsCatalog {
	if m != nil {
		return m.OutputsStore
	}
	return nil
}

func (m *RunSettings) GetInitConnections() []*RunSettingsCatalog {
	if m != nil {
		return m.InitConnections
	}
	return nil
}

func (m *RunSettings) GetConnections() []*RunSettingsCatalog {
	if m != nil {
		return m.Connections
	}
	return nil
}

func (m *RunSettings) GetGitAccesses() []*RunSettingsCatalog {
	if m != nil {
		return m.GitAccesses
	}
	return nil
}

func (m *RunSettings) GetRegistryAccess() *RunSettingsCatalog {
	if m != nil {
		return m.RegistryAccess
	}
	return nil
}

func (m *RunSettings) GetConfigResources() []*RunSettingsCatalog {
	if m != nil {
		return m.ConfigResources
	}
	return nil
}

func init() {
	proto.RegisterEnum("v1.RunKind", RunKind_name, RunKind_value)
	proto.RegisterType((*RunMetaInfo)(nil), "v1.RunMetaInfo")
	proto.RegisterType((*Run)(nil), "v1.Run")
	proto.RegisterType((*RunBodyRequest)(nil), "v1.RunBodyRequest")
	proto.RegisterType((*EntityRunBodyRequest)(nil), "v1.EntityRunBodyRequest")
	proto.RegisterType((*ListRunsResponse)(nil), "v1.ListRunsResponse")
	proto.RegisterType((*RunSettingsCatalog)(nil), "v1.RunSettingsCatalog")
	proto.RegisterType((*RunSettings)(nil), "v1.RunSettings")
}

func init() { proto.RegisterFile("v1/run.proto", fileDescriptor_46ec9b83e76db539) }

var fileDescriptor_46ec9b83e76db539 = []byte{
	// 1045 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xcd, 0x72, 0x1b, 0x45,
	0x10, 0x46, 0x92, 0xf5, 0xb3, 0x2d, 0xd9, 0x16, 0x43, 0x48, 0xc6, 0x8e, 0x93, 0x08, 0xe5, 0x62,
	0x28, 0x47, 0xc2, 0xa6, 0xa0, 0x92, 0x0a, 0x55, 0x94, 0x71, 0xe5, 0x40, 0x41, 0x28, 0x6a, 0xcd,
	0x8d, 0xc3, 0xd6, 0x68, 0xb7, 0xbd, 0x99, 0x78, 0x35, 0xb3, 0x99, 0x1f, 0x05, 0x17, 0x4f, 0xc0,
	0x9d, 0x77, 0xe1, 0xf5, 0xa8, 0xf9, 0x59, 0x5b, 0x98, 0xd8, 0xe6, 0xa4, 0xe9, 0xaf, 0xbf, 0x6f,
	0x7a, 0x7b, 0x7a, 0xba, 0x47, 0x30, 0x5a, 0x1d, 0xce, 0x95, 0x15, 0xb3, 0x5a, 0x49, 0x23, 0x49,
	0x7b, 0x75, 0xb8, 0xbb, 0x57, 0x4a, 0x59, 0x56, 0x38, 0x67, 0x35, 0x9f, 0x33, 0x21, 0xa4, 0x61,
	0x86, 0x4b, 0xa1, 0x03, 0xe3, 0xd2, 0xeb, 0xad, 0x85, 0x3d, 0x9b, 0x6b, 0xa3, 0x6c, 0x6e, 0xa2,
	0xf7, 0xc9, 0x75, 0xaf, 0xe1, 0x4b, 0xd4, 0x86, 0x2d, 0xeb, 0x48, 0x38, 0xf0, 0x3f, 0xf9, 0xb3,
	0x12, 0xc5, 0x33, 0xfd, 0x9e, 0x95, 0x25, 0xaa, 0xb9, 0xac, 0x7d, 0x80, 0x0f, 0x04, 0xdb, 0x5c,
	0x1d, 0xce, 0x17, 0x4c, 0x63, 0x30, 0xa7, 0x7f, 0xb6, 0x60, 0x98, 0x5a, 0xf1, 0x1a, 0x0d, 0xfb,
	0x41, 0x9c, 0x49, 0x42, 0xa1, 0xaf, 0x51, 0xad, 0x78, 0x8e, 0xb4, 0x35, 0x69, 0xed, 0x0f, 0xd2,
	0xc6, 0x24, 0x13, 0x18, 0xe6, 0x52, 0xe4, 0x56, 0x29, 0x14, 0xf9, 0x05, 0x6d, 0x4f, 0x5a, 0xfb,
	0xdd, 0x74, 0x1d, 0x22, 0x4f, 0x61, 0xb3, 0x66, 0x8a, 0x55, 0x15, 0x56, 0xd9, 0x39, 0x17, 0x05,
	0xed, 0x4c, 0x5a, 0xfb, 0x49, 0x3a, 0x6a, 0xc0, 0x1f, 0xb9, 0x28, 0xc8, 0x0e, 0x0c, 0x94, 0x15,
	0xc1, 0xbf, 0xe1, 0xfd, 0x7d, 0x65, 0x85, 0x73, 0x4d, 0xff, 0x1a, 0x40, 0x27, 0xb5, 0x82, 0x10,
	0xd8, 0xb0, 0x96, 0x17, 0xfe, 0x03, 0x92, 0xd4, 0xaf, 0x1d, 0x26, 0xd8, 0x12, 0x7d, 0xd8, 0x24,
	0xf5, 0x6b, 0xf7, 0x45, 0x05, 0xea, 0x5c, 0x71, 0x9f, 0x6c, 0x8c, 0xb6, 0x0e, 0x39, 0x95, 0x61,
	0xa5, 0xa6, 0x1b, 0x93, 0x8e, 0x53, 0xb9, 0xb5, 0xcb, 0xb0, 0xc0, 0x0a, 0x0d, 0x16, 0xb4, 0x1b,
	0x32, 0x8c, 0xa6, 0x8f, 0xab, 0x51, 0xd1, 0x5e, 0x8c, 0xab, 0x51, 0x91, 0x7b, 0xd0, 0x95, 0xef,
	0x05, 0x2a, 0xda, 0xf7, 0x60, 0x30, 0xdc, 0x1e, 0xb5, 0x92, 0x6f, 0x31, 0x37, 0x74, 0x10, 0x72,
	0x88, 0x26, 0x79, 0x01, 0x90, 0x2b, 0x64, 0x06, 0x8b, 0x8c, 0x19, 0x9a, 0x4c, 0x5a, 0xfb, 0xc3,
	0xa3, 0xdd, 0x59, 0x28, 0xe1, 0xac, 0x29, 0xe1, 0xec, 0xd7, 0xa6, 0x84, 0x69, 0x12, 0xd9, 0xc7,
	0x5e, 0x6a, 0xeb, 0xa2, 0x91, 0xc2, 0xdd, 0xd2, 0xc8, 0x0e, 0x52, 0x6d, 0x98, 0x8a, 0xd2, 0xe1,
	0xdd, 0xd2, 0xc8, 0x3e, 0x36, 0xe4, 0x25, 0x0c, 0xcf, 0xb8, 0xe0, 0xfa, 0x4d, 0xd0, 0x8e, 0xee,
	0xd4, 0x42, 0x43, 0x3f, 0x36, 0xe4, 0x11, 0x00, 0xd7, 0xd9, 0x92, 0x09, 0x56, 0x62, 0x41, 0x37,
	0xfd, 0x51, 0x24, 0x5c, 0xbf, 0x0e, 0x80, 0x3b, 0xa6, 0x5c, 0x0a, 0x83, 0xc2, 0xd0, 0xad, 0x70,
	0x4c, 0xd1, 0x24, 0xf7, 0xa1, 0xa7, 0x0d, 0x33, 0x56, 0xd3, 0x6d, 0xef, 0x88, 0x96, 0xc3, 0x15,
	0xb2, 0x62, 0x89, 0x74, 0x1c, 0xf0, 0x60, 0x91, 0xc7, 0x00, 0x0b, 0x29, 0xcf, 0x97, 0x4c, 0x9d,
	0x63, 0x41, 0x3f, 0xf6, 0x75, 0x5b, 0x43, 0xc8, 0x01, 0x24, 0x4b, 0x34, 0x2c, 0xe3, 0xe2, 0x4c,
	0x52, 0xe2, 0x73, 0xd8, 0x9e, 0xad, 0x0e, 0x67, 0x6b, 0x57, 0x3b, 0x1d, 0x2c, 0x9b, 0x4b, 0xfe,
	0x04, 0x36, 0xfc, 0xfd, 0xfb, 0x64, 0xd2, 0xda, 0xdf, 0x3a, 0x1a, 0x46, 0xa2, 0xbb, 0x83, 0xa9,
	0x77, 0x90, 0x31, 0x74, 0xde, 0xd8, 0x05, 0xbd, 0xe7, 0xbf, 0xc1, 0x2d, 0xc9, 0x1c, 0x7a, 0x5c,
	0xd4, 0xd6, 0x68, 0xfa, 0xa9, 0xdf, 0xfd, 0xc1, 0x7f, 0x4e, 0xe8, 0xd4, 0x37, 0x6d, 0x1a, 0x69,
	0xe4, 0x10, 0xfa, 0xd2, 0x1a, 0xaf, 0xb8, 0x7f, 0xbb, 0xa2, 0xe1, 0x91, 0x2f, 0xc1, 0xb5, 0x42,
	0x86, 0x62, 0x45, 0x1f, 0xdc, 0x11, 0x44, 0x59, 0xf1, 0x4a, 0xac, 0xc8, 0x43, 0x48, 0xb8, 0xce,
	0x14, 0x6a, 0xbb, 0x44, 0x4a, 0xfd, 0xa9, 0x0c, 0xb8, 0x4e, 0xbd, 0xed, 0x3a, 0x8d, 0xeb, 0x2c,
	0xaf, 0xa4, 0x40, 0xba, 0x13, 0x6e, 0x3a, 0xd7, 0x27, 0xce, 0x24, 0x9f, 0xc3, 0xd8, 0xe1, 0x5c,
	0x94, 0x99, 0x36, 0x8a, 0x19, 0x2c, 0x2f, 0xe8, 0xae, 0x4f, 0x76, 0x3b, 0xe2, 0xa7, 0x11, 0x26,
	0xbb, 0x30, 0xa8, 0x79, 0x8d, 0x15, 0x17, 0x48, 0x1f, 0x7a, 0xca, 0xa5, 0xed, 0x7c, 0x52, 0xf1,
	0x92, 0x0b, 0x56, 0xd1, 0xbd, 0xe0, 0x6b, 0x6c, 0x3f, 0x0c, 0x22, 0x2f, 0xf3, 0x9d, 0xfb, 0x28,
	0x0e, 0x83, 0x08, 0xfe, 0xec, 0x3a, 0xf8, 0x29, 0x6c, 0x36, 0x82, 0x40, 0x7a, 0x1c, 0x48, 0x0d,
	0xe8, 0x48, 0xd3, 0xdf, 0x60, 0x2b, 0xb5, 0xe2, 0x7b, 0x59, 0x5c, 0xa4, 0xf8, 0xce, 0xa2, 0x36,
	0x57, 0x4d, 0xd9, 0xba, 0xa1, 0x29, 0xdb, 0xff, 0x6e, 0xca, 0x1d, 0xe8, 0x28, 0x1b, 0x06, 0xc4,
	0xf0, 0xa8, 0x1f, 0xcb, 0x9d, 0x3a, 0x6c, 0x7a, 0x0e, 0xf7, 0x5e, 0x09, 0xc3, 0xcd, 0xc5, 0xb5,
	0x10, 0xcf, 0xa1, 0x87, 0x1e, 0xf7, 0x31, 0x86, 0x47, 0x13, 0xa7, 0xfa, 0x25, 0xec, 0x17, 0x05,
	0xa8, 0xa5, 0x55, 0x39, 0x46, 0x45, 0x1a, 0xf9, 0x4d, 0xb0, 0xf6, 0x07, 0x82, 0xfd, 0x01, 0xe3,
	0x9f, 0xb8, 0x36, 0xa9, 0x15, 0xae, 0x46, 0xb5, 0x14, 0x1a, 0x5d, 0x2e, 0xb9, 0xb4, 0xc2, 0xf8,
	0x38, 0xdd, 0x34, 0x18, 0xe4, 0x33, 0xe8, 0xbb, 0xaa, 0x56, 0x46, 0xd3, 0xf6, 0xa4, 0xb3, 0xbe,
	0x51, 0x83, 0xfb, 0xc2, 0x28, 0x5c, 0x71, 0x69, 0x75, 0x1c, 0x7d, 0x97, 0xb6, 0x9f, 0x96, 0xf8,
	0xbb, 0x89, 0x03, 0xd6, 0xaf, 0xa7, 0xdf, 0x02, 0x49, 0xad, 0x38, 0x45, 0x63, 0xb8, 0x28, 0xf5,
	0x09, 0x33, 0xac, 0x92, 0xe5, 0xff, 0x9d, 0xb5, 0xd3, 0xbf, 0x37, 0xfc, 0x3b, 0xd1, 0xc8, 0xc9,
	0x1e, 0x24, 0x0e, 0xd7, 0x35, 0x8b, 0x2f, 0x45, 0x92, 0x5e, 0x01, 0xe4, 0x00, 0xba, 0xac, 0x74,
	0x6d, 0x1f, 0x4e, 0xe1, 0x7e, 0xfc, 0xf8, 0x6b, 0xc1, 0xd3, 0x40, 0x72, 0xec, 0x77, 0x16, 0x2d,
	0xc6, 0x02, 0xdd, 0xc8, 0xf6, 0x24, 0xf2, 0x35, 0x40, 0x25, 0x4b, 0x9d, 0x69, 0x23, 0x15, 0xfa,
	0x0c, 0x6f, 0x96, 0x24, 0x8e, 0x79, 0xea, 0x88, 0xe4, 0x25, 0x6c, 0xc6, 0x3e, 0x8b, 0xca, 0xee,
	0xad, 0xca, 0x51, 0x24, 0x07, 0xf1, 0x31, 0x8c, 0xb9, 0xe0, 0x26, 0xcb, 0xa5, 0x10, 0x98, 0xfb,
	0xe7, 0x94, 0xf6, 0x7c, 0x5d, 0x6e, 0xd2, 0x6f, 0x3b, 0xfe, 0xc9, 0x15, 0x9d, 0x3c, 0xf7, 0xcf,
	0xe7, 0xa5, 0xba, 0x7f, 0xab, 0x7a, 0x9d, 0x4a, 0x5e, 0xc0, 0xa8, 0xe4, 0x26, 0x63, 0x79, 0x8e,
	0x5a, 0xa3, 0xa6, 0x83, 0xdb, 0xa5, 0x25, 0x37, 0xc7, 0x91, 0x4a, 0xbe, 0x83, 0x6d, 0x85, 0x25,
	0xd7, 0x46, 0x5d, 0x44, 0x7d, 0x7c, 0x92, 0x6e, 0x52, 0x6f, 0x35, 0xf4, 0xb0, 0x85, 0x4b, 0x3c,
	0x97, 0xe2, 0x8c, 0x97, 0x6e, 0xc8, 0xf8, 0xeb, 0xae, 0x29, 0xdc, 0x9e, 0x78, 0xe0, 0x37, 0xdd,
	0xa1, 0xbf, 0xf8, 0x06, 0xfa, 0x71, 0xb8, 0x92, 0x3e, 0x74, 0xde, 0xca, 0xc5, 0xf8, 0x23, 0x32,
	0xbc, 0xfc, 0x97, 0x31, 0x6e, 0x39, 0xb4, 0x60, 0xe5, 0xb8, 0x4d, 0x46, 0x30, 0x68, 0xfe, 0x2a,
	0x8c, 0x3b, 0x8b, 0x9e, 0x1f, 0x7a, 0x5f, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x5b, 0x37,
	0x53, 0x4e, 0x09, 0x00, 0x00,
}
