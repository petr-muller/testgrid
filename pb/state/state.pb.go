/*
Copyright The TestGrid Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: state.proto

package state

import (
	fmt "fmt"
	config "github.com/GoogleCloudPlatform/testgrid/pb/config"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Row_Result int32

const (
	Row_NO_RESULT         Row_Result = 0
	Row_PASS              Row_Result = 1
	Row_PASS_WITH_ERRORS  Row_Result = 2
	Row_PASS_WITH_SKIPS   Row_Result = 3
	Row_RUNNING           Row_Result = 4
	Row_CATEGORIZED_ABORT Row_Result = 5
	Row_UNKNOWN           Row_Result = 6
	Row_CANCEL            Row_Result = 7
	Row_BLOCKED           Row_Result = 8
	Row_TIMED_OUT         Row_Result = 9
	Row_CATEGORIZED_FAIL  Row_Result = 10
	Row_BUILD_FAIL        Row_Result = 11
	Row_FAIL              Row_Result = 12
	Row_FLAKY             Row_Result = 13
	Row_TOOL_FAIL         Row_Result = 14
	Row_BUILD_PASSED      Row_Result = 15
)

var Row_Result_name = map[int32]string{
	0:  "NO_RESULT",
	1:  "PASS",
	2:  "PASS_WITH_ERRORS",
	3:  "PASS_WITH_SKIPS",
	4:  "RUNNING",
	5:  "CATEGORIZED_ABORT",
	6:  "UNKNOWN",
	7:  "CANCEL",
	8:  "BLOCKED",
	9:  "TIMED_OUT",
	10: "CATEGORIZED_FAIL",
	11: "BUILD_FAIL",
	12: "FAIL",
	13: "FLAKY",
	14: "TOOL_FAIL",
	15: "BUILD_PASSED",
}

var Row_Result_value = map[string]int32{
	"NO_RESULT":         0,
	"PASS":              1,
	"PASS_WITH_ERRORS":  2,
	"PASS_WITH_SKIPS":   3,
	"RUNNING":           4,
	"CATEGORIZED_ABORT": 5,
	"UNKNOWN":           6,
	"CANCEL":            7,
	"BLOCKED":           8,
	"TIMED_OUT":         9,
	"CATEGORIZED_FAIL":  10,
	"BUILD_FAIL":        11,
	"FAIL":              12,
	"FLAKY":             13,
	"TOOL_FAIL":         14,
	"BUILD_PASSED":      15,
}

func (x Row_Result) String() string {
	return proto.EnumName(Row_Result_name, int32(x))
}

func (Row_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{6, 0}
}

// A metric and its values for each test cycle.
type Metric struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Sparse encoding of values. Indices is a list of pairs of <index, count> that
	// details columns with metric values. So given:
	//   Indices: [0, 2, 6, 4]
	//   Values: [0.1,0.2,6.1,6.2,6.3,6.4]
	// Decoded 12-value equivalent is:
	// [0.1, 0.2, nil, nil, nil, nil, nil, 6.1, 6.2, 6.3, 6.4, nil, nil, ...]
	Indices              []int32   `protobuf:"varint,2,rep,packed,name=indices,proto3" json:"indices,omitempty"`
	Values               []float64 `protobuf:"fixed64,3,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{0}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metric) GetIndices() []int32 {
	if m != nil {
		return m.Indices
	}
	return nil
}

func (m *Metric) GetValues() []float64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type UpdatePhaseData struct {
	// The name for a part of the update cycle.
	PhaseName string `protobuf:"bytes,1,opt,name=phase_name,json=phaseName,proto3" json:"phase_name,omitempty"`
	// Time taken for a part of the update cycle, in seconds.
	PhaseSeconds         float64  `protobuf:"fixed64,2,opt,name=phase_seconds,json=phaseSeconds,proto3" json:"phase_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePhaseData) Reset()         { *m = UpdatePhaseData{} }
func (m *UpdatePhaseData) String() string { return proto.CompactTextString(m) }
func (*UpdatePhaseData) ProtoMessage()    {}
func (*UpdatePhaseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{1}
}

func (m *UpdatePhaseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePhaseData.Unmarshal(m, b)
}
func (m *UpdatePhaseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePhaseData.Marshal(b, m, deterministic)
}
func (m *UpdatePhaseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePhaseData.Merge(m, src)
}
func (m *UpdatePhaseData) XXX_Size() int {
	return xxx_messageInfo_UpdatePhaseData.Size(m)
}
func (m *UpdatePhaseData) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePhaseData.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePhaseData proto.InternalMessageInfo

func (m *UpdatePhaseData) GetPhaseName() string {
	if m != nil {
		return m.PhaseName
	}
	return ""
}

func (m *UpdatePhaseData) GetPhaseSeconds() float64 {
	if m != nil {
		return m.PhaseSeconds
	}
	return 0
}

// Info on time taken to update test results during the last update cycle.
type UpdateInfo struct {
	// Metrics for how long parts of the update cycle take.
	UpdatePhaseData      []*UpdatePhaseData `protobuf:"bytes,1,rep,name=update_phase_data,json=updatePhaseData,proto3" json:"update_phase_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateInfo) Reset()         { *m = UpdateInfo{} }
func (m *UpdateInfo) String() string { return proto.CompactTextString(m) }
func (*UpdateInfo) ProtoMessage()    {}
func (*UpdateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{2}
}

func (m *UpdateInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateInfo.Unmarshal(m, b)
}
func (m *UpdateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateInfo.Marshal(b, m, deterministic)
}
func (m *UpdateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateInfo.Merge(m, src)
}
func (m *UpdateInfo) XXX_Size() int {
	return xxx_messageInfo_UpdateInfo.Size(m)
}
func (m *UpdateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateInfo proto.InternalMessageInfo

func (m *UpdateInfo) GetUpdatePhaseData() []*UpdatePhaseData {
	if m != nil {
		return m.UpdatePhaseData
	}
	return nil
}

// Info on a failing test row about the failure.
type AlertInfo struct {
	// Number of results that have failed.
	FailCount int32 `protobuf:"varint,1,opt,name=fail_count,json=failCount,proto3" json:"fail_count,omitempty"`
	// The build ID the test first failed at.
	FailBuildId string `protobuf:"bytes,2,opt,name=fail_build_id,json=failBuildId,proto3" json:"fail_build_id,omitempty"`
	// The time the test first failed at.
	FailTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=fail_time,json=failTime,proto3" json:"fail_time,omitempty"`
	// The test ID for the first test failure.
	FailTestId string `protobuf:"bytes,4,opt,name=fail_test_id,json=failTestId,proto3" json:"fail_test_id,omitempty"`
	// The build ID the test last passed at.
	PassBuildId string `protobuf:"bytes,5,opt,name=pass_build_id,json=passBuildId,proto3" json:"pass_build_id,omitempty"`
	// The time the test last passed at.
	PassTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=pass_time,json=passTime,proto3" json:"pass_time,omitempty"`
	// A snippet explaining the failure.
	FailureMessage string `protobuf:"bytes,7,opt,name=failure_message,json=failureMessage,proto3" json:"failure_message,omitempty"`
	// Link to search for build changes, internally a code-search link.
	BuildLink string `protobuf:"bytes,8,opt,name=build_link,json=buildLink,proto3" json:"build_link,omitempty"`
	// Text for option to search for build changes.
	BuildLinkText string `protobuf:"bytes,9,opt,name=build_link_text,json=buildLinkText,proto3" json:"build_link_text,omitempty"`
	// Text to display for link to search for build changes.
	BuildUrlText         string   `protobuf:"bytes,10,opt,name=build_url_text,json=buildUrlText,proto3" json:"build_url_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlertInfo) Reset()         { *m = AlertInfo{} }
func (m *AlertInfo) String() string { return proto.CompactTextString(m) }
func (*AlertInfo) ProtoMessage()    {}
func (*AlertInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{3}
}

func (m *AlertInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertInfo.Unmarshal(m, b)
}
func (m *AlertInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertInfo.Marshal(b, m, deterministic)
}
func (m *AlertInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertInfo.Merge(m, src)
}
func (m *AlertInfo) XXX_Size() int {
	return xxx_messageInfo_AlertInfo.Size(m)
}
func (m *AlertInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AlertInfo proto.InternalMessageInfo

func (m *AlertInfo) GetFailCount() int32 {
	if m != nil {
		return m.FailCount
	}
	return 0
}

func (m *AlertInfo) GetFailBuildId() string {
	if m != nil {
		return m.FailBuildId
	}
	return ""
}

func (m *AlertInfo) GetFailTime() *timestamp.Timestamp {
	if m != nil {
		return m.FailTime
	}
	return nil
}

func (m *AlertInfo) GetFailTestId() string {
	if m != nil {
		return m.FailTestId
	}
	return ""
}

func (m *AlertInfo) GetPassBuildId() string {
	if m != nil {
		return m.PassBuildId
	}
	return ""
}

func (m *AlertInfo) GetPassTime() *timestamp.Timestamp {
	if m != nil {
		return m.PassTime
	}
	return nil
}

func (m *AlertInfo) GetFailureMessage() string {
	if m != nil {
		return m.FailureMessage
	}
	return ""
}

func (m *AlertInfo) GetBuildLink() string {
	if m != nil {
		return m.BuildLink
	}
	return ""
}

func (m *AlertInfo) GetBuildLinkText() string {
	if m != nil {
		return m.BuildLinkText
	}
	return ""
}

func (m *AlertInfo) GetBuildUrlText() string {
	if m != nil {
		return m.BuildUrlText
	}
	return ""
}

// Info on default test metadata for a dashboard tab.
type TestMetadata struct {
	// Name of the test with associated test metadata.
	TestName string `protobuf:"bytes,1,opt,name=test_name,json=testName,proto3" json:"test_name,omitempty"`
	// Default bug component.
	BugComponent int32 `protobuf:"varint,2,opt,name=bug_component,json=bugComponent,proto3" json:"bug_component,omitempty"`
	// Default owner.
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	// Default list of cc's.
	Cc []string `protobuf:"bytes,4,rep,name=cc,proto3" json:"cc,omitempty"`
	// When present, only file a bug for failed tests with same error type.
	// Otherwise, always file a bug.
	ErrorType            string   `protobuf:"bytes,5,opt,name=error_type,json=errorType,proto3" json:"error_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestMetadata) Reset()         { *m = TestMetadata{} }
func (m *TestMetadata) String() string { return proto.CompactTextString(m) }
func (*TestMetadata) ProtoMessage()    {}
func (*TestMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{4}
}

func (m *TestMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMetadata.Unmarshal(m, b)
}
func (m *TestMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMetadata.Marshal(b, m, deterministic)
}
func (m *TestMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMetadata.Merge(m, src)
}
func (m *TestMetadata) XXX_Size() int {
	return xxx_messageInfo_TestMetadata.Size(m)
}
func (m *TestMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TestMetadata proto.InternalMessageInfo

func (m *TestMetadata) GetTestName() string {
	if m != nil {
		return m.TestName
	}
	return ""
}

func (m *TestMetadata) GetBugComponent() int32 {
	if m != nil {
		return m.BugComponent
	}
	return 0
}

func (m *TestMetadata) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *TestMetadata) GetCc() []string {
	if m != nil {
		return m.Cc
	}
	return nil
}

func (m *TestMetadata) GetErrorType() string {
	if m != nil {
		return m.ErrorType
	}
	return ""
}

// TestGrid columns (also known as TestCycle).
type Column struct {
	// Build ID.
	Build string `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty"`
	// Name associated with the column (such as the run/invocation ID).No two
	// columns should have the same build_id and name. The name field allows the
	// display of multiple columns with the same build_id.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Timestamp for the column, typically the earliest timestamp of any test
	// result in the column.
	Started float64 `protobuf:"fixed64,3,opt,name=started,proto3" json:"started,omitempty"`
	// Additional custom column labels, displayed below the standard column
	// headers of build_id, date, and time.
	Extra []string `protobuf:"bytes,4,rep,name=extra,proto3" json:"extra,omitempty"`
	// Custom hotlist ids.
	HotlistIds           string   `protobuf:"bytes,5,opt,name=hotlist_ids,json=hotlistIds,proto3" json:"hotlist_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Column) Reset()         { *m = Column{} }
func (m *Column) String() string { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()    {}
func (*Column) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{5}
}

func (m *Column) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Column.Unmarshal(m, b)
}
func (m *Column) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Column.Marshal(b, m, deterministic)
}
func (m *Column) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Column.Merge(m, src)
}
func (m *Column) XXX_Size() int {
	return xxx_messageInfo_Column.Size(m)
}
func (m *Column) XXX_DiscardUnknown() {
	xxx_messageInfo_Column.DiscardUnknown(m)
}

var xxx_messageInfo_Column proto.InternalMessageInfo

func (m *Column) GetBuild() string {
	if m != nil {
		return m.Build
	}
	return ""
}

func (m *Column) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Column) GetStarted() float64 {
	if m != nil {
		return m.Started
	}
	return 0
}

func (m *Column) GetExtra() []string {
	if m != nil {
		return m.Extra
	}
	return nil
}

func (m *Column) GetHotlistIds() string {
	if m != nil {
		return m.HotlistIds
	}
	return ""
}

// TestGrid rows (also known as TestRow)
type Row struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Results for this row, run-length encoded to reduce size/improve performance.
	// Thus (encoded -> decoded equivalent):
	//   [0, 3, 5, 4] -> [0, 0, 0, 5, 5, 5, 5]
	//   [5, 1] -> [5]
	//   [1, 5] -> [1, 1, 1, 1, 1]
	// The decoded values are Result enums
	Results []int32 `protobuf:"varint,3,rep,packed,name=results,proto3" json:"results,omitempty"`
	// Test IDs for each test result in this test case. If used, there is one
	// test ID per cycle with a non-empty status (where status is not NO_RESULT).
	CellIds []string `protobuf:"bytes,4,rep,name=cell_ids,json=cellIds,proto3" json:"cell_ids,omitempty"`
	// Text messages for each test result in this test case. There is one text
	// message per cycle with a non-empty status (where status is not NO_RESULT).
	Messages []string `protobuf:"bytes,5,rep,name=messages,proto3" json:"messages,omitempty"`
	// Names of metrics associated with this test case. Stored separate from
	// metric info (which may be omitted).
	Metric  []string  `protobuf:"bytes,7,rep,name=metric,proto3" json:"metric,omitempty"`
	Metrics []*Metric `protobuf:"bytes,8,rep,name=metrics,proto3" json:"metrics,omitempty"`
	Icons   []string  `protobuf:"bytes,9,rep,name=icons,proto3" json:"icons,omitempty"`
	// IDs for bugs associated with results in this test case.
	BugId []string `protobuf:"bytes,10,rep,name=bug_id,json=bugId,proto3" json:"bug_id,omitempty"`
	// An alert for the failure if there's a recent failure for this test case.
	AlertInfo            *AlertInfo `protobuf:"bytes,11,opt,name=alert_info,json=alertInfo,proto3" json:"alert_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}
func (*Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{6}
}

func (m *Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Row.Unmarshal(m, b)
}
func (m *Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Row.Marshal(b, m, deterministic)
}
func (m *Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Row.Merge(m, src)
}
func (m *Row) XXX_Size() int {
	return xxx_messageInfo_Row.Size(m)
}
func (m *Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Row proto.InternalMessageInfo

func (m *Row) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Row) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Row) GetResults() []int32 {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *Row) GetCellIds() []string {
	if m != nil {
		return m.CellIds
	}
	return nil
}

func (m *Row) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *Row) GetMetric() []string {
	if m != nil {
		return m.Metric
	}
	return nil
}

func (m *Row) GetMetrics() []*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *Row) GetIcons() []string {
	if m != nil {
		return m.Icons
	}
	return nil
}

func (m *Row) GetBugId() []string {
	if m != nil {
		return m.BugId
	}
	return nil
}

func (m *Row) GetAlertInfo() *AlertInfo {
	if m != nil {
		return m.AlertInfo
	}
	return nil
}

// A single table of test results backing a dashboard tab.
type Grid struct {
	// A cycle of test results, not including the results. In the TestGrid client,
	// the cycles define the columns.
	Columns []*Column `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	// A test case with test results. In the TestGrid client, the cases define the
	// rows (and the results define the individual cells).
	Rows []*Row `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	// Seconds since epoch when last mail alert was sent.
	LastAlertMailTime float64 `protobuf:"fixed64,3,opt,name=last_alert_mail_time,json=lastAlertMailTime,proto3" json:"last_alert_mail_time,omitempty"`
	// The latest configuration used to generate this test group.
	Config *config.TestGroup `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	// Seconds since epoch for last time this cycle was updated.
	LastTimeUpdated float64 `protobuf:"fixed64,6,opt,name=last_time_updated,json=lastTimeUpdated,proto3" json:"last_time_updated,omitempty"`
	// Stored info on previous timing for parts of the update cycle.
	UpdateInfo []*UpdateInfo `protobuf:"bytes,8,rep,name=update_info,json=updateInfo,proto3" json:"update_info,omitempty"`
	// Stored info on default test metadata.
	TestMetadata []*TestMetadata `protobuf:"bytes,9,rep,name=test_metadata,json=testMetadata,proto3" json:"test_metadata,omitempty"`
	// Clusters of failures for a TestResultTable instance.
	Cluster []*Cluster `protobuf:"bytes,10,rep,name=cluster,proto3" json:"cluster,omitempty"`
	// Most recent timestamp that clusters have processed.
	MostRecentClusterTimestamp float64  `protobuf:"fixed64,11,opt,name=most_recent_cluster_timestamp,json=mostRecentClusterTimestamp,proto3" json:"most_recent_cluster_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *Grid) Reset()         { *m = Grid{} }
func (m *Grid) String() string { return proto.CompactTextString(m) }
func (*Grid) ProtoMessage()    {}
func (*Grid) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{7}
}

func (m *Grid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Grid.Unmarshal(m, b)
}
func (m *Grid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Grid.Marshal(b, m, deterministic)
}
func (m *Grid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grid.Merge(m, src)
}
func (m *Grid) XXX_Size() int {
	return xxx_messageInfo_Grid.Size(m)
}
func (m *Grid) XXX_DiscardUnknown() {
	xxx_messageInfo_Grid.DiscardUnknown(m)
}

var xxx_messageInfo_Grid proto.InternalMessageInfo

func (m *Grid) GetColumns() []*Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *Grid) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *Grid) GetLastAlertMailTime() float64 {
	if m != nil {
		return m.LastAlertMailTime
	}
	return 0
}

func (m *Grid) GetConfig() *config.TestGroup {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Grid) GetLastTimeUpdated() float64 {
	if m != nil {
		return m.LastTimeUpdated
	}
	return 0
}

func (m *Grid) GetUpdateInfo() []*UpdateInfo {
	if m != nil {
		return m.UpdateInfo
	}
	return nil
}

func (m *Grid) GetTestMetadata() []*TestMetadata {
	if m != nil {
		return m.TestMetadata
	}
	return nil
}

func (m *Grid) GetCluster() []*Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *Grid) GetMostRecentClusterTimestamp() float64 {
	if m != nil {
		return m.MostRecentClusterTimestamp
	}
	return 0
}

// A cluster of failures grouped by test status and message for a test results
// table.
type Cluster struct {
	// Test status cluster grouped by.
	TestStatus int32 `protobuf:"varint,1,opt,name=test_status,json=testStatus,proto3" json:"test_status,omitempty"`
	// Error message or testFailureClassification string cluster grouped by.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// ClusterRows that belong to this cluster.
	ClusterRow           []*ClusterRow `protobuf:"bytes,3,rep,name=cluster_row,json=clusterRow,proto3" json:"cluster_row,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{8}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetTestStatus() int32 {
	if m != nil {
		return m.TestStatus
	}
	return 0
}

func (m *Cluster) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Cluster) GetClusterRow() []*ClusterRow {
	if m != nil {
		return m.ClusterRow
	}
	return nil
}

// Cells in a TestRow that belong to a specific Cluster.
type ClusterRow struct {
	// Name of TestRow.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Index within row that belongs to Cluster (refer to columns of the row).
	Index                []int32  `protobuf:"varint,2,rep,packed,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterRow) Reset()         { *m = ClusterRow{} }
func (m *ClusterRow) String() string { return proto.CompactTextString(m) }
func (*ClusterRow) ProtoMessage()    {}
func (*ClusterRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{9}
}

func (m *ClusterRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterRow.Unmarshal(m, b)
}
func (m *ClusterRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterRow.Marshal(b, m, deterministic)
}
func (m *ClusterRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterRow.Merge(m, src)
}
func (m *ClusterRow) XXX_Size() int {
	return xxx_messageInfo_ClusterRow.Size(m)
}
func (m *ClusterRow) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterRow.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterRow proto.InternalMessageInfo

func (m *ClusterRow) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ClusterRow) GetIndex() []int32 {
	if m != nil {
		return m.Index
	}
	return nil
}

func init() {
	proto.RegisterEnum("Row_Result", Row_Result_name, Row_Result_value)
	proto.RegisterType((*Metric)(nil), "Metric")
	proto.RegisterType((*UpdatePhaseData)(nil), "UpdatePhaseData")
	proto.RegisterType((*UpdateInfo)(nil), "UpdateInfo")
	proto.RegisterType((*AlertInfo)(nil), "AlertInfo")
	proto.RegisterType((*TestMetadata)(nil), "TestMetadata")
	proto.RegisterType((*Column)(nil), "Column")
	proto.RegisterType((*Row)(nil), "Row")
	proto.RegisterType((*Grid)(nil), "Grid")
	proto.RegisterType((*Cluster)(nil), "Cluster")
	proto.RegisterType((*ClusterRow)(nil), "ClusterRow")
}

func init() { proto.RegisterFile("state.proto", fileDescriptor_a888679467bb7853) }

var fileDescriptor_a888679467bb7853 = []byte{
	// 1150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdf, 0x8f, 0xdb, 0x44,
	0x10, 0xc6, 0xf9, 0x65, 0x7b, 0x9c, 0x5c, 0xdc, 0xa5, 0xad, 0xcc, 0xa1, 0xaa, 0xa9, 0x41, 0x10,
	0x10, 0xca, 0x49, 0xc7, 0x03, 0x2f, 0xbc, 0xe4, 0x72, 0xe9, 0xe1, 0xbb, 0x5c, 0x52, 0x6d, 0x12,
	0x55, 0xf0, 0x62, 0xf9, 0xec, 0xbd, 0xd4, 0xaa, 0x63, 0x47, 0xf6, 0x9a, 0xbb, 0x3e, 0x22, 0xfe,
	0x06, 0x84, 0x90, 0xf8, 0x63, 0xd1, 0xcc, 0xda, 0x49, 0x8a, 0x90, 0x78, 0xca, 0x7e, 0xdf, 0x4c,
	0x66, 0xc6, 0x3b, 0x33, 0xdf, 0x82, 0x55, 0xc8, 0x40, 0x8a, 0xd1, 0x2e, 0xcf, 0x64, 0x76, 0xfa,
	0x72, 0x93, 0x65, 0x9b, 0x44, 0x9c, 0x11, 0xba, 0x2b, 0xef, 0xcf, 0x64, 0xbc, 0x15, 0x85, 0x0c,
	0xb6, 0xbb, 0xca, 0xe1, 0xf9, 0xee, 0xee, 0x2c, 0xcc, 0xd2, 0xfb, 0x78, 0x53, 0xfd, 0x28, 0xde,
	0x9d, 0x43, 0xe7, 0x56, 0xc8, 0x3c, 0x0e, 0x19, 0x83, 0x56, 0x1a, 0x6c, 0x85, 0xa3, 0x0d, 0xb4,
	0xa1, 0xc9, 0xe9, 0xcc, 0x1c, 0xd0, 0xe3, 0x34, 0x8a, 0x43, 0x51, 0x38, 0x8d, 0x41, 0x73, 0xd8,
	0xe6, 0x35, 0x64, 0xcf, 0xa1, 0xf3, 0x6b, 0x90, 0x94, 0xa2, 0x70, 0x9a, 0x83, 0xe6, 0x50, 0xe3,
	0x15, 0x72, 0xd7, 0xd0, 0x5f, 0xef, 0xa2, 0x40, 0x8a, 0x37, 0xef, 0x82, 0x42, 0x5c, 0x06, 0x32,
	0x60, 0x2f, 0x00, 0x76, 0x08, 0xfc, 0xa3, 0xf0, 0x26, 0x31, 0x73, 0xcc, 0xf1, 0x05, 0xf4, 0x94,
	0xb9, 0x10, 0x61, 0x96, 0x46, 0x98, 0x49, 0x1b, 0x6a, 0xbc, 0x4b, 0xe4, 0x52, 0x71, 0xee, 0x35,
	0x80, 0x0a, 0xeb, 0xa5, 0xf7, 0x19, 0xfb, 0x11, 0x9e, 0x94, 0x84, 0x7c, 0xf5, 0xcf, 0x28, 0x90,
	0x81, 0xa3, 0x0d, 0x9a, 0x43, 0xeb, 0xdc, 0x1e, 0xfd, 0x2b, 0x3d, 0xef, 0x97, 0x1f, 0x13, 0xee,
	0xdf, 0x4d, 0x30, 0xc7, 0x89, 0xc8, 0x25, 0xc5, 0x7a, 0x01, 0x70, 0x1f, 0xc4, 0x89, 0x1f, 0x66,
	0x65, 0x2a, 0xa9, 0xba, 0x36, 0x37, 0x91, 0x99, 0x20, 0xc1, 0x5c, 0xe8, 0x91, 0xf9, 0xae, 0x8c,
	0x93, 0xc8, 0x8f, 0x23, 0xaa, 0xce, 0xe4, 0x16, 0x92, 0x17, 0xc8, 0x79, 0x11, 0xfb, 0x01, 0xe8,
	0x0f, 0x3e, 0xde, 0xb9, 0xd3, 0x1c, 0x68, 0x43, 0xeb, 0xfc, 0x74, 0xa4, 0x1a, 0x32, 0xaa, 0x1b,
	0x32, 0x5a, 0xd5, 0x0d, 0xe1, 0x06, 0x3a, 0x23, 0x64, 0x03, 0xe8, 0xaa, 0x3f, 0x8a, 0x42, 0x62,
	0xec, 0x16, 0xc5, 0xa6, 0x7a, 0x56, 0xa2, 0x90, 0x5e, 0x84, 0xe9, 0x77, 0x41, 0x51, 0x1c, 0xd2,
	0xb7, 0x55, 0x7a, 0x24, 0x8f, 0xd2, 0x93, 0x0f, 0xa5, 0xef, 0xfc, 0x7f, 0x7a, 0x74, 0xa6, 0xf4,
	0x5f, 0x43, 0x1f, 0x53, 0x95, 0xb9, 0xf0, 0xb7, 0xa2, 0x28, 0x82, 0x8d, 0x70, 0x74, 0x0a, 0x7f,
	0x52, 0xd1, 0xb7, 0x8a, 0xc5, 0x3b, 0x52, 0x05, 0x24, 0x71, 0xfa, 0xde, 0x31, 0x54, 0x07, 0x89,
	0x99, 0xc5, 0xe9, 0x7b, 0xf6, 0x15, 0xf4, 0x0f, 0x66, 0x5f, 0x8a, 0x47, 0xe9, 0x98, 0xe4, 0xd3,
	0xdb, 0xfb, 0xac, 0xc4, 0xa3, 0x64, 0x5f, 0xc2, 0x89, 0xf2, 0x2b, 0xf3, 0x44, 0xb9, 0x01, 0xb9,
	0x75, 0x89, 0x5d, 0xe7, 0x09, 0x7a, 0xb9, 0x7f, 0x68, 0xd0, 0xc5, 0xaf, 0xbf, 0x15, 0x32, 0xc0,
	0xc6, 0xb2, 0xcf, 0xc1, 0xa4, 0x0b, 0x3a, 0x1a, 0x1f, 0x03, 0x89, 0x7a, 0x7a, 0xee, 0xca, 0x8d,
	0x1f, 0x66, 0xdb, 0x5d, 0x96, 0x8a, 0x54, 0x52, 0x7f, 0xda, 0x18, 0x72, 0x33, 0xa9, 0x39, 0xf6,
	0x14, 0xda, 0xd9, 0x43, 0x2a, 0x72, 0x6a, 0x8e, 0xc9, 0x15, 0x60, 0x27, 0xd0, 0x08, 0x43, 0xa7,
	0x35, 0x68, 0x0e, 0x4d, 0xde, 0x08, 0x43, 0xfc, 0x4a, 0x91, 0xe7, 0x59, 0xee, 0xcb, 0x0f, 0x3b,
	0x51, 0x5d, 0xb4, 0x49, 0xcc, 0xea, 0xc3, 0x4e, 0xb8, 0xbf, 0x6b, 0xd0, 0x99, 0x64, 0x49, 0xb9,
	0x4d, 0x31, 0x1e, 0x95, 0x5c, 0x55, 0xa3, 0xc0, 0x7e, 0x81, 0x1a, 0x1f, 0x2f, 0x50, 0x21, 0x83,
	0x5c, 0x8a, 0x88, 0x72, 0x6b, 0xbc, 0x86, 0x18, 0x43, 0x3c, 0xca, 0x3c, 0xa8, 0x0a, 0x50, 0x80,
	0xbd, 0x04, 0xeb, 0x5d, 0x26, 0x93, 0x98, 0xe6, 0xa1, 0xa8, 0x8a, 0x80, 0x8a, 0xf2, 0xa2, 0xc2,
	0xfd, 0xb3, 0x05, 0x4d, 0x9e, 0x3d, 0xfc, 0xe7, 0xb6, 0x9e, 0x40, 0x63, 0x3f, 0xa0, 0x8d, 0x38,
	0xc2, 0xe4, 0xb9, 0x28, 0xca, 0x44, 0xaa, 0x25, 0x6d, 0xf3, 0x1a, 0xb2, 0xcf, 0xc0, 0x08, 0x45,
	0x92, 0x50, 0x0e, 0x95, 0x5f, 0x47, 0xec, 0x45, 0x05, 0x3b, 0x05, 0xa3, 0x1a, 0x06, 0x4c, 0x8f,
	0xa6, 0x3d, 0xc6, 0xa5, 0xdf, 0x92, 0x58, 0x38, 0x3a, 0x59, 0x2a, 0xc4, 0x5e, 0x81, 0xae, 0x4e,
	0x85, 0x63, 0xd0, 0x16, 0xea, 0x23, 0x25, 0x2a, 0xbc, 0xe6, 0xf1, 0x73, 0xe3, 0x30, 0x4b, 0x0b,
	0xc7, 0x54, 0x9f, 0x4b, 0x80, 0x3d, 0x83, 0x0e, 0x76, 0x2f, 0x8e, 0x1c, 0x50, 0xf4, 0x5d, 0xb9,
	0xf1, 0x22, 0xf6, 0x0d, 0x40, 0x80, 0x0b, 0xea, 0xc7, 0xe9, 0x7d, 0xe6, 0x58, 0x34, 0xd2, 0x30,
	0xda, 0xef, 0x2c, 0x37, 0x83, 0xfa, 0xe8, 0xfe, 0xd6, 0x80, 0x0e, 0xa7, 0xaf, 0x62, 0x3d, 0x30,
	0xe7, 0x0b, 0x9f, 0x4f, 0x97, 0xeb, 0xd9, 0xca, 0xfe, 0x84, 0x19, 0xd0, 0x7a, 0x33, 0x5e, 0x2e,
	0x6d, 0x8d, 0x3d, 0x05, 0x1b, 0x4f, 0xfe, 0x5b, 0x6f, 0xf5, 0x93, 0x3f, 0xe5, 0x7c, 0xc1, 0x97,
	0x76, 0x83, 0x7d, 0x0a, 0xfd, 0x03, 0xbb, 0xbc, 0xf1, 0xde, 0x2c, 0xed, 0x26, 0xb3, 0x40, 0xe7,
	0xeb, 0xf9, 0xdc, 0x9b, 0x5f, 0xd9, 0x2d, 0xf6, 0x0c, 0x9e, 0x4c, 0xc6, 0xab, 0xe9, 0xd5, 0x82,
	0x7b, 0xbf, 0x4c, 0x2f, 0xfd, 0xf1, 0xc5, 0x82, 0xaf, 0xec, 0x36, 0xfa, 0xac, 0xe7, 0x37, 0xf3,
	0xc5, 0xdb, 0xb9, 0xdd, 0x61, 0x00, 0x9d, 0xc9, 0x78, 0x3e, 0x99, 0xce, 0x6c, 0x1d, 0x0d, 0x17,
	0xb3, 0xc5, 0xe4, 0x66, 0x7a, 0x69, 0x1b, 0x58, 0xcd, 0xca, 0xbb, 0x9d, 0x5e, 0xfa, 0x8b, 0xf5,
	0xca, 0x36, 0xb1, 0x86, 0xe3, 0x58, 0xaf, 0xc7, 0xde, 0xcc, 0x06, 0x76, 0x02, 0x70, 0xb1, 0xf6,
	0x66, 0x15, 0xb6, 0xb0, 0x66, 0x3a, 0x75, 0x99, 0x09, 0xed, 0xd7, 0xb3, 0xf1, 0xcd, 0xcf, 0x76,
	0x8f, 0x22, 0x2d, 0x16, 0x33, 0xe5, 0x73, 0xc2, 0x6c, 0xe8, 0xaa, 0xff, 0x60, 0xf5, 0xd3, 0x4b,
	0xbb, 0x7f, 0xdd, 0x32, 0x3a, 0xb6, 0xee, 0xfe, 0xd5, 0x84, 0xd6, 0x55, 0x1e, 0x47, 0xd8, 0x8d,
	0x90, 0xe6, 0xb4, 0xa8, 0x34, 0x51, 0x1f, 0xa9, 0xb9, 0xe5, 0x35, 0xcf, 0x1c, 0x68, 0xe5, 0xd9,
	0x83, 0x12, 0x75, 0xeb, 0xbc, 0x35, 0xe2, 0xd9, 0x03, 0x27, 0x86, 0x9d, 0xc1, 0xd3, 0x24, 0x28,
	0xa4, 0xaf, 0xee, 0x7f, 0xfb, 0x91, 0xac, 0x69, 0xfc, 0x09, 0xda, 0xa8, 0x0f, 0xb7, 0xb5, 0x86,
	0xb9, 0xd0, 0x51, 0x0f, 0x0a, 0xa9, 0x17, 0xf6, 0x09, 0x97, 0xf7, 0x2a, 0xcf, 0xca, 0x1d, 0xaf,
	0x2c, 0xec, 0x5b, 0xa0, 0x3f, 0x52, 0x24, 0x5f, 0xc9, 0x71, 0x44, 0x4a, 0xa5, 0xf1, 0x3e, 0x1a,
	0x30, 0x90, 0x92, 0xed, 0x88, 0x7d, 0x07, 0x56, 0xa5, 0xed, 0xd4, 0x7c, 0x35, 0x4f, 0xd6, 0xe8,
	0xa0, 0xfe, 0x1c, 0xca, 0xc3, 0x4b, 0x70, 0x0e, 0x3d, 0xd2, 0x86, 0x6d, 0x25, 0x16, 0x34, 0x5e,
	0xd6, 0x79, 0x6f, 0x74, 0xac, 0x20, 0xbc, 0x2b, 0x8f, 0xf5, 0xc4, 0x05, 0x3d, 0x4c, 0xca, 0x42,
	0x8a, 0x9c, 0xa6, 0xce, 0x3a, 0x37, 0x46, 0x13, 0x85, 0x79, 0x6d, 0x60, 0x63, 0x78, 0xb1, 0xcd,
	0x0a, 0xe9, 0xe7, 0x22, 0x14, 0xa9, 0xf4, 0x2b, 0xda, 0xdf, 0xbf, 0xaa, 0x34, 0x94, 0x1a, 0x3f,
	0x45, 0x27, 0x4e, 0x3e, 0x55, 0x88, 0xbd, 0xce, 0x5e, 0xb7, 0x8c, 0xb6, 0xdd, 0xb9, 0x6e, 0x19,
	0xba, 0x6d, 0xb8, 0x39, 0xe8, 0x95, 0x1d, 0x37, 0x9c, 0x2a, 0xc6, 0xd7, 0xbb, 0x2c, 0xaa, 0x07,
	0x07, 0x90, 0x5a, 0x12, 0x83, 0x5b, 0x5b, 0xab, 0xb1, 0x5a, 0xe5, 0x1a, 0xe2, 0xd5, 0xd4, 0x85,
	0xe4, 0xd9, 0x03, 0xed, 0x34, 0x5e, 0x4d, 0x5d, 0x7c, 0xf6, 0xc0, 0x21, 0xdc, 0x9f, 0xdd, 0x29,
	0xc0, 0xc1, 0xc2, 0x5e, 0x41, 0x37, 0x8a, 0x8b, 0x5d, 0x12, 0x7c, 0x38, 0xd6, 0x51, 0xab, 0xe2,
	0x48, 0x4a, 0x71, 0x45, 0xd3, 0x48, 0x3c, 0x56, 0x4f, 0xbd, 0x02, 0x77, 0x1d, 0x7a, 0x42, 0xbe,
	0xff, 0x27, 0x00, 0x00, 0xff, 0xff, 0x74, 0xb5, 0x6a, 0x2e, 0x6f, 0x08, 0x00, 0x00,
}
