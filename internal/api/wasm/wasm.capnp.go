// Code generated by capnpc-go. DO NOT EDIT.

package wasm

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	fc "capnproto.org/go/capnp/v3/flowcontrol"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	context "context"
	fmt "fmt"
	iostream "github.com/wetware/ww/internal/api/iostream"
	proc "github.com/wetware/ww/internal/api/proc"
)

type Runtime capnp.Client

// Runtime_TypeID is the unique identifier for the type Runtime.
const Runtime_TypeID = 0xff6bb7b1b05afb0e

func (c Runtime) Exec(ctx context.Context, params func(proc.Executor_exec_Params) error) (proc.Executor_exec_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xe8bb307fa2f406fb,
			MethodID:      0,
			InterfaceName: "proc.capnp:Executor",
			MethodName:    "exec",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(proc.Executor_exec_Params(s)) }
	}
	ans, release := capnp.Client(c).SendCall(ctx, s)
	return proc.Executor_exec_Results_Future{Future: ans.Future()}, release
}

// String returns a string that identifies this capability for debugging
// purposes.  Its format should not be depended on: in particular, it
// should not be used to compare clients.  Use IsSame to compare clients
// for equality.
func (c Runtime) String() string {
	return fmt.Sprintf("%T(%v)", c, capnp.Client(c))
}

// AddRef creates a new Client that refers to the same capability as c.
// If c is nil or has resolved to null, then AddRef returns nil.
func (c Runtime) AddRef() Runtime {
	return Runtime(capnp.Client(c).AddRef())
}

// Release releases a capability reference.  If this is the last
// reference to the capability, then the underlying resources associated
// with the capability will be released.
//
// Release will panic if c has already been released, but not if c is
// nil or resolved to null.
func (c Runtime) Release() {
	capnp.Client(c).Release()
}

// Resolve blocks until the capability is fully resolved or the Context
// expires.
func (c Runtime) Resolve(ctx context.Context) error {
	return capnp.Client(c).Resolve(ctx)
}

func (c Runtime) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Client(c).EncodeAsPtr(seg)
}

func (Runtime) DecodeFromPtr(p capnp.Ptr) Runtime {
	return Runtime(capnp.Client{}.DecodeFromPtr(p))
}

// IsValid reports whether c is a valid reference to a capability.
// A reference is invalid if it is nil, has resolved to null, or has
// been released.
func (c Runtime) IsValid() bool {
	return capnp.Client(c).IsValid()
}

// IsSame reports whether c and other refer to a capability created by the
// same call to NewClient.  This can return false negatives if c or other
// are not fully resolved: use Resolve if this is an issue.  If either
// c or other are released, then IsSame panics.
func (c Runtime) IsSame(other Runtime) bool {
	return capnp.Client(c).IsSame(capnp.Client(other))
}

// Update the flowcontrol.FlowLimiter used to manage flow control for
// this client. This affects all future calls, but not calls already
// waiting to send. Passing nil sets the value to flowcontrol.NopLimiter,
// which is also the default.
func (c Runtime) SetFlowLimiter(lim fc.FlowLimiter) {
	capnp.Client(c).SetFlowLimiter(lim)
}

// Get the current flowcontrol.FlowLimiter used to manage flow control
// for this client.
func (c Runtime) GetFlowLimiter() fc.FlowLimiter {
	return capnp.Client(c).GetFlowLimiter()
} // A Runtime_Server is a Runtime with a local implementation.
type Runtime_Server interface {
	Exec(context.Context, proc.Executor_exec) error
}

// Runtime_NewServer creates a new Server from an implementation of Runtime_Server.
func Runtime_NewServer(s Runtime_Server) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Runtime_Methods(nil, s), s, c)
}

// Runtime_ServerToClient creates a new Client from an implementation of Runtime_Server.
// The caller is responsible for calling Release on the returned Client.
func Runtime_ServerToClient(s Runtime_Server) Runtime {
	return Runtime(capnp.NewClient(Runtime_NewServer(s)))
}

// Runtime_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Runtime_Methods(methods []server.Method, s Runtime_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe8bb307fa2f406fb,
			MethodID:      0,
			InterfaceName: "proc.capnp:Executor",
			MethodName:    "exec",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Exec(ctx, proc.Executor_exec{call})
		},
	})

	return methods
}

// Runtime_List is a list of Runtime.
type Runtime_List = capnp.CapList[Runtime]

// NewRuntime creates a new list of Runtime.
func NewRuntime_List(s *capnp.Segment, sz int32) (Runtime_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[Runtime](l), err
}

type Runtime_Context capnp.Struct

// Runtime_Context_TypeID is the unique identifier for the type Runtime_Context.
const Runtime_Context_TypeID = 0xf0d086c2f7dd91fa

func NewRuntime_Context(s *capnp.Segment) (Runtime_Context, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return Runtime_Context(st), err
}

func NewRootRuntime_Context(s *capnp.Segment) (Runtime_Context, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return Runtime_Context(st), err
}

func ReadRootRuntime_Context(msg *capnp.Message) (Runtime_Context, error) {
	root, err := msg.Root()
	return Runtime_Context(root.Struct()), err
}

func (s Runtime_Context) String() string {
	str, _ := text.Marshal(0xf0d086c2f7dd91fa, capnp.Struct(s))
	return str
}

func (s Runtime_Context) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Runtime_Context) DecodeFromPtr(p capnp.Ptr) Runtime_Context {
	return Runtime_Context(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Runtime_Context) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Runtime_Context) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Runtime_Context) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Runtime_Context) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Runtime_Context) Src() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return []byte(p.Data()), err
}

func (s Runtime_Context) HasSrc() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Runtime_Context) SetSrc(v []byte) error {
	return capnp.Struct(s).SetData(0, v)
}

func (s Runtime_Context) Env() (Runtime_Context_Field_List, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return Runtime_Context_Field_List(p.List()), err
}

func (s Runtime_Context) HasEnv() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Runtime_Context) SetEnv(v Runtime_Context_Field_List) error {
	return capnp.Struct(s).SetPtr(1, v.ToPtr())
}

// NewEnv sets the env field to a newly
// allocated Runtime_Context_Field_List, preferring placement in s's segment.
func (s Runtime_Context) NewEnv(n int32) (Runtime_Context_Field_List, error) {
	l, err := NewRuntime_Context_Field_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Runtime_Context_Field_List{}, err
	}
	err = capnp.Struct(s).SetPtr(1, l.ToPtr())
	return l, err
}
func (s Runtime_Context) Stdin() iostream.Provider {
	p, _ := capnp.Struct(s).Ptr(2)
	return iostream.Provider(p.Interface().Client())
}

func (s Runtime_Context) HasStdin() bool {
	return capnp.Struct(s).HasPtr(2)
}

func (s Runtime_Context) SetStdin(v iostream.Provider) error {
	if !v.IsValid() {
		return capnp.Struct(s).SetPtr(2, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(capnp.Client(v)))
	return capnp.Struct(s).SetPtr(2, in.ToPtr())
}

func (s Runtime_Context) Stdout() iostream.Stream {
	p, _ := capnp.Struct(s).Ptr(3)
	return iostream.Stream(p.Interface().Client())
}

func (s Runtime_Context) HasStdout() bool {
	return capnp.Struct(s).HasPtr(3)
}

func (s Runtime_Context) SetStdout(v iostream.Stream) error {
	if !v.IsValid() {
		return capnp.Struct(s).SetPtr(3, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(capnp.Client(v)))
	return capnp.Struct(s).SetPtr(3, in.ToPtr())
}

func (s Runtime_Context) Stderr() iostream.Stream {
	p, _ := capnp.Struct(s).Ptr(4)
	return iostream.Stream(p.Interface().Client())
}

func (s Runtime_Context) HasStderr() bool {
	return capnp.Struct(s).HasPtr(4)
}

func (s Runtime_Context) SetStderr(v iostream.Stream) error {
	if !v.IsValid() {
		return capnp.Struct(s).SetPtr(4, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(capnp.Client(v)))
	return capnp.Struct(s).SetPtr(4, in.ToPtr())
}

func (s Runtime_Context) RandSeed() uint64 {
	return capnp.Struct(s).Uint64(0)
}

func (s Runtime_Context) SetRandSeed(v uint64) {
	capnp.Struct(s).SetUint64(0, v)
}

// Runtime_Context_List is a list of Runtime_Context.
type Runtime_Context_List = capnp.StructList[Runtime_Context]

// NewRuntime_Context creates a new list of Runtime_Context.
func NewRuntime_Context_List(s *capnp.Segment, sz int32) (Runtime_Context_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5}, sz)
	return capnp.StructList[Runtime_Context](l), err
}

// Runtime_Context_Future is a wrapper for a Runtime_Context promised by a client call.
type Runtime_Context_Future struct{ *capnp.Future }

func (f Runtime_Context_Future) Struct() (Runtime_Context, error) {
	p, err := f.Future.Ptr()
	return Runtime_Context(p.Struct()), err
}
func (p Runtime_Context_Future) Stdin() iostream.Provider {
	return iostream.Provider(p.Future.Field(2, nil).Client())
}

func (p Runtime_Context_Future) Stdout() iostream.Stream {
	return iostream.Stream(p.Future.Field(3, nil).Client())
}

func (p Runtime_Context_Future) Stderr() iostream.Stream {
	return iostream.Stream(p.Future.Field(4, nil).Client())
}

type Runtime_Context_Field capnp.Struct

// Runtime_Context_Field_TypeID is the unique identifier for the type Runtime_Context_Field.
const Runtime_Context_Field_TypeID = 0xee2036e1e97a5dce

func NewRuntime_Context_Field(s *capnp.Segment) (Runtime_Context_Field, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Runtime_Context_Field(st), err
}

func NewRootRuntime_Context_Field(s *capnp.Segment) (Runtime_Context_Field, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Runtime_Context_Field(st), err
}

func ReadRootRuntime_Context_Field(msg *capnp.Message) (Runtime_Context_Field, error) {
	root, err := msg.Root()
	return Runtime_Context_Field(root.Struct()), err
}

func (s Runtime_Context_Field) String() string {
	str, _ := text.Marshal(0xee2036e1e97a5dce, capnp.Struct(s))
	return str
}

func (s Runtime_Context_Field) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Runtime_Context_Field) DecodeFromPtr(p capnp.Ptr) Runtime_Context_Field {
	return Runtime_Context_Field(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Runtime_Context_Field) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Runtime_Context_Field) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Runtime_Context_Field) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Runtime_Context_Field) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Runtime_Context_Field) Key() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Runtime_Context_Field) HasKey() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Runtime_Context_Field) KeyBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Runtime_Context_Field) SetKey(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Runtime_Context_Field) Value() (string, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.Text(), err
}

func (s Runtime_Context_Field) HasValue() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Runtime_Context_Field) ValueBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.TextBytes(), err
}

func (s Runtime_Context_Field) SetValue(v string) error {
	return capnp.Struct(s).SetText(1, v)
}

// Runtime_Context_Field_List is a list of Runtime_Context_Field.
type Runtime_Context_Field_List = capnp.StructList[Runtime_Context_Field]

// NewRuntime_Context_Field creates a new list of Runtime_Context_Field.
func NewRuntime_Context_Field_List(s *capnp.Segment, sz int32) (Runtime_Context_Field_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return capnp.StructList[Runtime_Context_Field](l), err
}

// Runtime_Context_Field_Future is a wrapper for a Runtime_Context_Field promised by a client call.
type Runtime_Context_Field_Future struct{ *capnp.Future }

func (f Runtime_Context_Field_Future) Struct() (Runtime_Context_Field, error) {
	p, err := f.Future.Ptr()
	return Runtime_Context_Field(p.Struct()), err
}

type Runtime_Module capnp.Client

// Runtime_Module_TypeID is the unique identifier for the type Runtime_Module.
const Runtime_Module_TypeID = 0x84c67f001342cf99

func (c Runtime_Module) Run(ctx context.Context, params func(Runtime_Module_run_Params) error) (Runtime_Module_run_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0x84c67f001342cf99,
			MethodID:      0,
			InterfaceName: "wasm.capnp:Runtime.Module",
			MethodName:    "run",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Runtime_Module_run_Params(s)) }
	}
	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Runtime_Module_run_Results_Future{Future: ans.Future()}, release
}
func (c Runtime_Module) Close(ctx context.Context, params func(Runtime_Module_close_Params) error) (Runtime_Module_close_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0x84c67f001342cf99,
			MethodID:      1,
			InterfaceName: "wasm.capnp:Runtime.Module",
			MethodName:    "close",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Runtime_Module_close_Params(s)) }
	}
	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Runtime_Module_close_Results_Future{Future: ans.Future()}, release
}
func (c Runtime_Module) Wait(ctx context.Context, params func(proc.Waiter_wait_Params) error) (proc.Waiter_wait_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xc66c9bda04b0f29e,
			MethodID:      0,
			InterfaceName: "proc.capnp:Waiter",
			MethodName:    "wait",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(proc.Waiter_wait_Params(s)) }
	}
	ans, release := capnp.Client(c).SendCall(ctx, s)
	return proc.Waiter_wait_Results_Future{Future: ans.Future()}, release
}

// String returns a string that identifies this capability for debugging
// purposes.  Its format should not be depended on: in particular, it
// should not be used to compare clients.  Use IsSame to compare clients
// for equality.
func (c Runtime_Module) String() string {
	return fmt.Sprintf("%T(%v)", c, capnp.Client(c))
}

// AddRef creates a new Client that refers to the same capability as c.
// If c is nil or has resolved to null, then AddRef returns nil.
func (c Runtime_Module) AddRef() Runtime_Module {
	return Runtime_Module(capnp.Client(c).AddRef())
}

// Release releases a capability reference.  If this is the last
// reference to the capability, then the underlying resources associated
// with the capability will be released.
//
// Release will panic if c has already been released, but not if c is
// nil or resolved to null.
func (c Runtime_Module) Release() {
	capnp.Client(c).Release()
}

// Resolve blocks until the capability is fully resolved or the Context
// expires.
func (c Runtime_Module) Resolve(ctx context.Context) error {
	return capnp.Client(c).Resolve(ctx)
}

func (c Runtime_Module) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Client(c).EncodeAsPtr(seg)
}

func (Runtime_Module) DecodeFromPtr(p capnp.Ptr) Runtime_Module {
	return Runtime_Module(capnp.Client{}.DecodeFromPtr(p))
}

// IsValid reports whether c is a valid reference to a capability.
// A reference is invalid if it is nil, has resolved to null, or has
// been released.
func (c Runtime_Module) IsValid() bool {
	return capnp.Client(c).IsValid()
}

// IsSame reports whether c and other refer to a capability created by the
// same call to NewClient.  This can return false negatives if c or other
// are not fully resolved: use Resolve if this is an issue.  If either
// c or other are released, then IsSame panics.
func (c Runtime_Module) IsSame(other Runtime_Module) bool {
	return capnp.Client(c).IsSame(capnp.Client(other))
}

// Update the flowcontrol.FlowLimiter used to manage flow control for
// this client. This affects all future calls, but not calls already
// waiting to send. Passing nil sets the value to flowcontrol.NopLimiter,
// which is also the default.
func (c Runtime_Module) SetFlowLimiter(lim fc.FlowLimiter) {
	capnp.Client(c).SetFlowLimiter(lim)
}

// Get the current flowcontrol.FlowLimiter used to manage flow control
// for this client.
func (c Runtime_Module) GetFlowLimiter() fc.FlowLimiter {
	return capnp.Client(c).GetFlowLimiter()
} // A Runtime_Module_Server is a Runtime_Module with a local implementation.
type Runtime_Module_Server interface {
	Run(context.Context, Runtime_Module_run) error

	Close(context.Context, Runtime_Module_close) error

	Wait(context.Context, proc.Waiter_wait) error
}

// Runtime_Module_NewServer creates a new Server from an implementation of Runtime_Module_Server.
func Runtime_Module_NewServer(s Runtime_Module_Server) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Runtime_Module_Methods(nil, s), s, c)
}

// Runtime_Module_ServerToClient creates a new Client from an implementation of Runtime_Module_Server.
// The caller is responsible for calling Release on the returned Client.
func Runtime_Module_ServerToClient(s Runtime_Module_Server) Runtime_Module {
	return Runtime_Module(capnp.NewClient(Runtime_Module_NewServer(s)))
}

// Runtime_Module_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Runtime_Module_Methods(methods []server.Method, s Runtime_Module_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 3)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x84c67f001342cf99,
			MethodID:      0,
			InterfaceName: "wasm.capnp:Runtime.Module",
			MethodName:    "run",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Run(ctx, Runtime_Module_run{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x84c67f001342cf99,
			MethodID:      1,
			InterfaceName: "wasm.capnp:Runtime.Module",
			MethodName:    "close",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Close(ctx, Runtime_Module_close{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc66c9bda04b0f29e,
			MethodID:      0,
			InterfaceName: "proc.capnp:Waiter",
			MethodName:    "wait",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Wait(ctx, proc.Waiter_wait{call})
		},
	})

	return methods
}

// Runtime_Module_run holds the state for a server call to Runtime_Module.run.
// See server.Call for documentation.
type Runtime_Module_run struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Runtime_Module_run) Args() Runtime_Module_run_Params {
	return Runtime_Module_run_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Runtime_Module_run) AllocResults() (Runtime_Module_run_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Runtime_Module_run_Results(r), err
}

// Runtime_Module_close holds the state for a server call to Runtime_Module.close.
// See server.Call for documentation.
type Runtime_Module_close struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Runtime_Module_close) Args() Runtime_Module_close_Params {
	return Runtime_Module_close_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Runtime_Module_close) AllocResults() (Runtime_Module_close_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Runtime_Module_close_Results(r), err
}

// Runtime_Module_List is a list of Runtime_Module.
type Runtime_Module_List = capnp.CapList[Runtime_Module]

// NewRuntime_Module creates a new list of Runtime_Module.
func NewRuntime_Module_List(s *capnp.Segment, sz int32) (Runtime_Module_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[Runtime_Module](l), err
}

type Runtime_Module_Status capnp.Struct

// Runtime_Module_Status_TypeID is the unique identifier for the type Runtime_Module_Status.
const Runtime_Module_Status_TypeID = 0xa0d2281eb34bc498

func NewRuntime_Module_Status(s *capnp.Segment) (Runtime_Module_Status, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Runtime_Module_Status(st), err
}

func NewRootRuntime_Module_Status(s *capnp.Segment) (Runtime_Module_Status, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Runtime_Module_Status(st), err
}

func ReadRootRuntime_Module_Status(msg *capnp.Message) (Runtime_Module_Status, error) {
	root, err := msg.Root()
	return Runtime_Module_Status(root.Struct()), err
}

func (s Runtime_Module_Status) String() string {
	str, _ := text.Marshal(0xa0d2281eb34bc498, capnp.Struct(s))
	return str
}

func (s Runtime_Module_Status) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Runtime_Module_Status) DecodeFromPtr(p capnp.Ptr) Runtime_Module_Status {
	return Runtime_Module_Status(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Runtime_Module_Status) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Runtime_Module_Status) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Runtime_Module_Status) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Runtime_Module_Status) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Runtime_Module_Status) StatusCode() uint32 {
	return capnp.Struct(s).Uint32(0)
}

func (s Runtime_Module_Status) SetStatusCode(v uint32) {
	capnp.Struct(s).SetUint32(0, v)
}

// Runtime_Module_Status_List is a list of Runtime_Module_Status.
type Runtime_Module_Status_List = capnp.StructList[Runtime_Module_Status]

// NewRuntime_Module_Status creates a new list of Runtime_Module_Status.
func NewRuntime_Module_Status_List(s *capnp.Segment, sz int32) (Runtime_Module_Status_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return capnp.StructList[Runtime_Module_Status](l), err
}

// Runtime_Module_Status_Future is a wrapper for a Runtime_Module_Status promised by a client call.
type Runtime_Module_Status_Future struct{ *capnp.Future }

func (f Runtime_Module_Status_Future) Struct() (Runtime_Module_Status, error) {
	p, err := f.Future.Ptr()
	return Runtime_Module_Status(p.Struct()), err
}

type Runtime_Module_run_Params capnp.Struct

// Runtime_Module_run_Params_TypeID is the unique identifier for the type Runtime_Module_run_Params.
const Runtime_Module_run_Params_TypeID = 0xe7389b7d0b5f32af

func NewRuntime_Module_run_Params(s *capnp.Segment) (Runtime_Module_run_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Runtime_Module_run_Params(st), err
}

func NewRootRuntime_Module_run_Params(s *capnp.Segment) (Runtime_Module_run_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Runtime_Module_run_Params(st), err
}

func ReadRootRuntime_Module_run_Params(msg *capnp.Message) (Runtime_Module_run_Params, error) {
	root, err := msg.Root()
	return Runtime_Module_run_Params(root.Struct()), err
}

func (s Runtime_Module_run_Params) String() string {
	str, _ := text.Marshal(0xe7389b7d0b5f32af, capnp.Struct(s))
	return str
}

func (s Runtime_Module_run_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Runtime_Module_run_Params) DecodeFromPtr(p capnp.Ptr) Runtime_Module_run_Params {
	return Runtime_Module_run_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Runtime_Module_run_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Runtime_Module_run_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Runtime_Module_run_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Runtime_Module_run_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// Runtime_Module_run_Params_List is a list of Runtime_Module_run_Params.
type Runtime_Module_run_Params_List = capnp.StructList[Runtime_Module_run_Params]

// NewRuntime_Module_run_Params creates a new list of Runtime_Module_run_Params.
func NewRuntime_Module_run_Params_List(s *capnp.Segment, sz int32) (Runtime_Module_run_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Runtime_Module_run_Params](l), err
}

// Runtime_Module_run_Params_Future is a wrapper for a Runtime_Module_run_Params promised by a client call.
type Runtime_Module_run_Params_Future struct{ *capnp.Future }

func (f Runtime_Module_run_Params_Future) Struct() (Runtime_Module_run_Params, error) {
	p, err := f.Future.Ptr()
	return Runtime_Module_run_Params(p.Struct()), err
}

type Runtime_Module_run_Results capnp.Struct

// Runtime_Module_run_Results_TypeID is the unique identifier for the type Runtime_Module_run_Results.
const Runtime_Module_run_Results_TypeID = 0xe058eb2275c7c09f

func NewRuntime_Module_run_Results(s *capnp.Segment) (Runtime_Module_run_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Runtime_Module_run_Results(st), err
}

func NewRootRuntime_Module_run_Results(s *capnp.Segment) (Runtime_Module_run_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Runtime_Module_run_Results(st), err
}

func ReadRootRuntime_Module_run_Results(msg *capnp.Message) (Runtime_Module_run_Results, error) {
	root, err := msg.Root()
	return Runtime_Module_run_Results(root.Struct()), err
}

func (s Runtime_Module_run_Results) String() string {
	str, _ := text.Marshal(0xe058eb2275c7c09f, capnp.Struct(s))
	return str
}

func (s Runtime_Module_run_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Runtime_Module_run_Results) DecodeFromPtr(p capnp.Ptr) Runtime_Module_run_Results {
	return Runtime_Module_run_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Runtime_Module_run_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Runtime_Module_run_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Runtime_Module_run_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Runtime_Module_run_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// Runtime_Module_run_Results_List is a list of Runtime_Module_run_Results.
type Runtime_Module_run_Results_List = capnp.StructList[Runtime_Module_run_Results]

// NewRuntime_Module_run_Results creates a new list of Runtime_Module_run_Results.
func NewRuntime_Module_run_Results_List(s *capnp.Segment, sz int32) (Runtime_Module_run_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Runtime_Module_run_Results](l), err
}

// Runtime_Module_run_Results_Future is a wrapper for a Runtime_Module_run_Results promised by a client call.
type Runtime_Module_run_Results_Future struct{ *capnp.Future }

func (f Runtime_Module_run_Results_Future) Struct() (Runtime_Module_run_Results, error) {
	p, err := f.Future.Ptr()
	return Runtime_Module_run_Results(p.Struct()), err
}

type Runtime_Module_close_Params capnp.Struct

// Runtime_Module_close_Params_TypeID is the unique identifier for the type Runtime_Module_close_Params.
const Runtime_Module_close_Params_TypeID = 0xe13d3cafc7c68823

func NewRuntime_Module_close_Params(s *capnp.Segment) (Runtime_Module_close_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Runtime_Module_close_Params(st), err
}

func NewRootRuntime_Module_close_Params(s *capnp.Segment) (Runtime_Module_close_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Runtime_Module_close_Params(st), err
}

func ReadRootRuntime_Module_close_Params(msg *capnp.Message) (Runtime_Module_close_Params, error) {
	root, err := msg.Root()
	return Runtime_Module_close_Params(root.Struct()), err
}

func (s Runtime_Module_close_Params) String() string {
	str, _ := text.Marshal(0xe13d3cafc7c68823, capnp.Struct(s))
	return str
}

func (s Runtime_Module_close_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Runtime_Module_close_Params) DecodeFromPtr(p capnp.Ptr) Runtime_Module_close_Params {
	return Runtime_Module_close_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Runtime_Module_close_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Runtime_Module_close_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Runtime_Module_close_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Runtime_Module_close_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Runtime_Module_close_Params) ExitCode() uint32 {
	return capnp.Struct(s).Uint32(0)
}

func (s Runtime_Module_close_Params) SetExitCode(v uint32) {
	capnp.Struct(s).SetUint32(0, v)
}

// Runtime_Module_close_Params_List is a list of Runtime_Module_close_Params.
type Runtime_Module_close_Params_List = capnp.StructList[Runtime_Module_close_Params]

// NewRuntime_Module_close_Params creates a new list of Runtime_Module_close_Params.
func NewRuntime_Module_close_Params_List(s *capnp.Segment, sz int32) (Runtime_Module_close_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return capnp.StructList[Runtime_Module_close_Params](l), err
}

// Runtime_Module_close_Params_Future is a wrapper for a Runtime_Module_close_Params promised by a client call.
type Runtime_Module_close_Params_Future struct{ *capnp.Future }

func (f Runtime_Module_close_Params_Future) Struct() (Runtime_Module_close_Params, error) {
	p, err := f.Future.Ptr()
	return Runtime_Module_close_Params(p.Struct()), err
}

type Runtime_Module_close_Results capnp.Struct

// Runtime_Module_close_Results_TypeID is the unique identifier for the type Runtime_Module_close_Results.
const Runtime_Module_close_Results_TypeID = 0xc9d95f3753a8ec42

func NewRuntime_Module_close_Results(s *capnp.Segment) (Runtime_Module_close_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Runtime_Module_close_Results(st), err
}

func NewRootRuntime_Module_close_Results(s *capnp.Segment) (Runtime_Module_close_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Runtime_Module_close_Results(st), err
}

func ReadRootRuntime_Module_close_Results(msg *capnp.Message) (Runtime_Module_close_Results, error) {
	root, err := msg.Root()
	return Runtime_Module_close_Results(root.Struct()), err
}

func (s Runtime_Module_close_Results) String() string {
	str, _ := text.Marshal(0xc9d95f3753a8ec42, capnp.Struct(s))
	return str
}

func (s Runtime_Module_close_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Runtime_Module_close_Results) DecodeFromPtr(p capnp.Ptr) Runtime_Module_close_Results {
	return Runtime_Module_close_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Runtime_Module_close_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Runtime_Module_close_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Runtime_Module_close_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Runtime_Module_close_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// Runtime_Module_close_Results_List is a list of Runtime_Module_close_Results.
type Runtime_Module_close_Results_List = capnp.StructList[Runtime_Module_close_Results]

// NewRuntime_Module_close_Results creates a new list of Runtime_Module_close_Results.
func NewRuntime_Module_close_Results_List(s *capnp.Segment, sz int32) (Runtime_Module_close_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Runtime_Module_close_Results](l), err
}

// Runtime_Module_close_Results_Future is a wrapper for a Runtime_Module_close_Results promised by a client call.
type Runtime_Module_close_Results_Future struct{ *capnp.Future }

func (f Runtime_Module_close_Results_Future) Struct() (Runtime_Module_close_Results, error) {
	p, err := f.Future.Ptr()
	return Runtime_Module_close_Results(p.Struct()), err
}

const schema_9419a7a54f76d35b = "x\xda|T]h\x1cU\x14\xfe\xbe{gg&6" +
	"i{\x99\x854\x15\xba\x12\x13\xda\x0a\x09\xe9VTV" +
	"ec\x02>(\xe2\xde\x04Q*Z\x96\xecEB7" +
	"\xbbeg6\x8dB\x89\x0f\xa2\xbe\x08\x0a\x15ZK\xb4" +
	"\x08\xa5\x16\x94&\x98\xa2b\x0bU$\xb5\x0f\xfePA" +
	"\xb0 X_T,\xad\x88P\xb0\x0f#w\x92\xdd\x1d" +
	"L\xe8\xcb0s\xcew\xcf\xf7\x9d\xef\x9c;#\xa7\xc4" +
	"\xa8\xb3\xa7\xe7\x85n\x08\xfdf\xc6\xbdq\xec\xbb\xb1`" +
	"~\xe5\x15\x15\xc8x\xf3\xad}\x8bK\x9f\x1c\x88\x01\x06" +
	"}\xee\xe5`\xd0\xed\x05\x82!\xf7\xb5\xe0\xb0}\x8b\x8f" +
	"~\xf5\xf8\xc7;v]>\x01\xbd\x9dl\x1dt<`" +
	"\xef\x82{\x07\x83\x8f\\\x0f\x08N\xbbE0\x1e\xbb\xf6" +
	"\xc1\xe4\xfd\xfb\x7f\xba\x04\xb5\x9d@\x82\xf9\xc6\x1d#\x9c" +
	"\xf8\xbd\x0b\x17\x9b\xfd\x7f>\xf3K*s\xd6\xcd\xdb\xcc" +
	"\xdd\xaf\xaf\\<\xf3\xd0\xc3W\x93\xf2k\xa9\x05\xb7@" +
	"08\x99\x14=\x93\xdf\xbf\xe9\xf0\xf1\x07~K\x1d\xbd" +
	"\xe4\xdec\x8f~\xfb\xdcK\x7f\\\xbd\xef\xae\xebPw" +
	"2\xfe\xf7\xad\x9fo~\xf9\xea\xf7\x7f!#,d\xc9" +
	"j\xfb\"\xd1v\xde=\x84T^\x07d\xa7\xe9L\xc6" +
	"B\xfa\xbc+\xc1\xa0\xd7\x0b\xec\x1d\xf2\x9e&Ry\xb5" +
	"I\xc6\xcf\xfe0\xfb\xe4\xc9S}G\xacCo\xf8\xef" +
	"\x04o\xfb;m\xcb\xbe\x17\x9c\xf6{\x11cw|\xa8" +
	"\x1c\xce\x0cO\x95\x0f\x8a\xda\xc1\xc2D\xb3\x16M\xcf\x98" +
	"\xe1'\xea\x95f\x95F;d\xc7C\xc5Bq2*" +
	"G\xcdP\xfb2\x03\xb4\xbbc\xcb!\xb5\xa7\x1fB\x0d" +
	"zd\xdb\x1a\xb6|U}y\x08\xd5\xe3y\x8dfm" +
	"\x94\xb9\xa9j=4\xa3L\x18\xde\xfd{\xd1\xb9r\xbc" +
	"\xba\x82Q*\xe6\xb4#R!@\xb1\xd7\xa2\xc8\x92$" +
	"\xb7v\xe4\x80m\xe9\xf2\xff\xd2\xcd\xf0\xaaP\xa0Dj" +
	"G:\x80C@\xf5\xec\x03t\xb7\xa4\xde&\x18\x87\x09" +
	"d\xbc\x0eY1\xf4!\xe8\xa7J\xfa\xebK&\x92\x07" +
	"&L\xd8\xacF!Z\xc0\xf5\xb8F\xb360ar" +
	"\x09\xecv\x12W\xeb\x15K\xe5Fy&L\x8b|," +
	"%\xd2\xccMG\xe3\xf5\x8a\x01\xb0N\xa3\xdc\x98\xbbT" +
	"\xdeb+n\x08\x1b\xaf\xd7\"3\x17\x0d?:m\xaa" +
	"\x95Uw\xfc6\xf1\xee~@\x0fH\xea\x11AEf" +
	"\xedN\xab\xa1<\xa0wI\xea{\x05\xbd\x03\xe6Ev" +
	"C\xb0\x1b\xcc\xcd\x96\xabM\xd3\xfa\xdap\x8b\x1229" +
	"\x17%C\xee,<\xf3\xb9\x84^ok3\x1f\xb3\xcc" +
	"G$\xf5\x89\x14\xf3\x82\x0d\x1e\x95\xd4\xcb\x82J\x88," +
	"\x05\xa0\x96\xac\x9c\x0f%\xf5\xa7\x82J\xca,%\xa0\xce" +
	"\x16\x00\xbd(\xa9\xcf\x09*\xc7\xc9\xd2\x01\xd4g6\xb8" +
	",\xa9/\x082\x93e\x06P\xe7\xad\xb5\xe7$\xf5\xd7" +
	"\x82^\xd8\x98b\x0f\x04{@\xcf\xd4f\xb9\x19k+" +
	"\xd6\x96\x0a\xda`.\x8c*\xd35\xaa\x9bO=xc" +
	"\xfe\xf9\xfek \x15X\x0c\xa3J\xbd\x19Q\xc5\xbf>" +
	"\xb2\xfc\xe3\x8e\xeb[^F'a\x1a\x8d\xf5\x89\xb8Q" +
	"\xaeU&\x8d\xa9\xd8YvA\xb0+\xe5\x1b[\xbe\xc9" +
	"\x19\xa3}\xa6\xee\xbc\xea\x1ak\xff\xf22\x85\xf9\xb5\x09" +
	"\x16W\x07^\x92\x99\xc4\xdd[\xee?\xef\xcf\x8f|\xfe" +
	"{\xea\x0a\xb5C\xf6\x0a\xed\xb45\x93a\xdb\xe7\xd6\xa4" +
	"\xcd6\x83\x95\xd7\xa2\x00\xff\x0b\x00\x00\xff\xff\x09\xc1\x94" +
	"h"

func init() {
	schemas.Register(schema_9419a7a54f76d35b,
		0x84c67f001342cf99,
		0xa0d2281eb34bc498,
		0xc9d95f3753a8ec42,
		0xe058eb2275c7c09f,
		0xe13d3cafc7c68823,
		0xe7389b7d0b5f32af,
		0xee2036e1e97a5dce,
		0xf0d086c2f7dd91fa,
		0xff6bb7b1b05afb0e)
}
