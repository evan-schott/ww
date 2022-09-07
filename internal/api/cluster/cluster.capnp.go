// Code generated by capnpc-go. DO NOT EDIT.

package cluster

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	fc "capnproto.org/go/capnp/v3/flowcontrol"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	context "context"
	fmt "fmt"
	pubsub "github.com/wetware/ww/internal/api/pubsub"
)

type Host capnp.Client

// Host_TypeID is the unique identifier for the type Host.
const Host_TypeID = 0x957cbefc645fd307

func (c Host) View(ctx context.Context, params func(Host_view_Params) error) (Host_view_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0x957cbefc645fd307,
			MethodID:      0,
			InterfaceName: "cluster.capnp:Host",
			MethodName:    "view",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Host_view_Params(s)) }
	}
	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Host_view_Results_Future{Future: ans.Future()}, release
}
func (c Host) PubSub(ctx context.Context, params func(Host_pubSub_Params) error) (Host_pubSub_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0x957cbefc645fd307,
			MethodID:      1,
			InterfaceName: "cluster.capnp:Host",
			MethodName:    "pubSub",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Host_pubSub_Params(s)) }
	}
	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Host_pubSub_Results_Future{Future: ans.Future()}, release
}

// String returns a string that identifies this capability for debugging
// purposes.  Its format should not be depended on: in particular, it
// should not be used to compare clients.  Use IsSame to compare clients
// for equality.
func (c Host) String() string {
	return fmt.Sprintf("%T(%v)", c, capnp.Client(c))
}

// AddRef creates a new Client that refers to the same capability as c.
// If c is nil or has resolved to null, then AddRef returns nil.
func (c Host) AddRef() Host {
	return Host(capnp.Client(c).AddRef())
}

// Release releases a capability reference.  If this is the last
// reference to the capability, then the underlying resources associated
// with the capability will be released.
//
// Release will panic if c has already been released, but not if c is
// nil or resolved to null.
func (c Host) Release() {
	capnp.Client(c).Release()
}

// Resolve blocks until the capability is fully resolved or the Context
// expires.
func (c Host) Resolve(ctx context.Context) error {
	return capnp.Client(c).Resolve(ctx)
}

func (c Host) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Client(c).EncodeAsPtr(seg)
}

func (Host) DecodeFromPtr(p capnp.Ptr) Host {
	return Host(capnp.Client{}.DecodeFromPtr(p))
}

// IsValid reports whether c is a valid reference to a capability.
// A reference is invalid if it is nil, has resolved to null, or has
// been released.
func (c Host) IsValid() bool {
	return capnp.Client(c).IsValid()
}

// IsSame reports whether c and other refer to a capability created by the
// same call to NewClient.  This can return false negatives if c or other
// are not fully resolved: use Resolve if this is an issue.  If either
// c or other are released, then IsSame panics.
func (c Host) IsSame(other Host) bool {
	return capnp.Client(c).IsSame(capnp.Client(other))
}

// Update the flowcontrol.FlowLimiter used to manage flow control for
// this client. This affects all future calls, but not calls already
// waiting to send. Passing nil sets the value to flowcontrol.NopLimiter,
// which is also the default.
func (c Host) SetFlowLimiter(lim fc.FlowLimiter) {
	capnp.Client(c).SetFlowLimiter(lim)
}

// Get the current flowcontrol.FlowLimiter used to manage flow control
// for this client.
func (c Host) GetFlowLimiter() fc.FlowLimiter {
	return capnp.Client(c).GetFlowLimiter()
} // A Host_Server is a Host with a local implementation.
type Host_Server interface {
	View(context.Context, Host_view) error

	PubSub(context.Context, Host_pubSub) error
}

// Host_NewServer creates a new Server from an implementation of Host_Server.
func Host_NewServer(s Host_Server) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Host_Methods(nil, s), s, c)
}

// Host_ServerToClient creates a new Client from an implementation of Host_Server.
// The caller is responsible for calling Release on the returned Client.
func Host_ServerToClient(s Host_Server) Host {
	return Host(capnp.NewClient(Host_NewServer(s)))
}

// Host_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Host_Methods(methods []server.Method, s Host_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x957cbefc645fd307,
			MethodID:      0,
			InterfaceName: "cluster.capnp:Host",
			MethodName:    "view",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.View(ctx, Host_view{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x957cbefc645fd307,
			MethodID:      1,
			InterfaceName: "cluster.capnp:Host",
			MethodName:    "pubSub",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.PubSub(ctx, Host_pubSub{call})
		},
	})

	return methods
}

// Host_view holds the state for a server call to Host.view.
// See server.Call for documentation.
type Host_view struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Host_view) Args() Host_view_Params {
	return Host_view_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Host_view) AllocResults() (Host_view_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_view_Results(r), err
}

// Host_pubSub holds the state for a server call to Host.pubSub.
// See server.Call for documentation.
type Host_pubSub struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Host_pubSub) Args() Host_pubSub_Params {
	return Host_pubSub_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Host_pubSub) AllocResults() (Host_pubSub_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_pubSub_Results(r), err
}

// Host_List is a list of Host.
type Host_List = capnp.CapList[Host]

// NewHost creates a new list of Host.
func NewHost_List(s *capnp.Segment, sz int32) (Host_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[Host](l), err
}

type Host_view_Params capnp.Struct

// Host_view_Params_TypeID is the unique identifier for the type Host_view_Params.
const Host_view_Params_TypeID = 0xa404c24b5375b9e4

func NewHost_view_Params(s *capnp.Segment) (Host_view_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Host_view_Params(st), err
}

func NewRootHost_view_Params(s *capnp.Segment) (Host_view_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Host_view_Params(st), err
}

func ReadRootHost_view_Params(msg *capnp.Message) (Host_view_Params, error) {
	root, err := msg.Root()
	return Host_view_Params(root.Struct()), err
}

func (s Host_view_Params) String() string {
	str, _ := text.Marshal(0xa404c24b5375b9e4, capnp.Struct(s))
	return str
}

func (s Host_view_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Host_view_Params) DecodeFromPtr(p capnp.Ptr) Host_view_Params {
	return Host_view_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Host_view_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Host_view_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Host_view_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Host_view_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// Host_view_Params_List is a list of Host_view_Params.
type Host_view_Params_List = capnp.StructList[Host_view_Params]

// NewHost_view_Params creates a new list of Host_view_Params.
func NewHost_view_Params_List(s *capnp.Segment, sz int32) (Host_view_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Host_view_Params](l), err
}

// Host_view_Params_Future is a wrapper for a Host_view_Params promised by a client call.
type Host_view_Params_Future struct{ *capnp.Future }

func (p Host_view_Params_Future) Struct() (Host_view_Params, error) {
	s, err := p.Future.Struct()
	return Host_view_Params(s), err
}

type Host_view_Results capnp.Struct

// Host_view_Results_TypeID is the unique identifier for the type Host_view_Results.
const Host_view_Results_TypeID = 0x8f58928e854cd4f5

func NewHost_view_Results(s *capnp.Segment) (Host_view_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_view_Results(st), err
}

func NewRootHost_view_Results(s *capnp.Segment) (Host_view_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_view_Results(st), err
}

func ReadRootHost_view_Results(msg *capnp.Message) (Host_view_Results, error) {
	root, err := msg.Root()
	return Host_view_Results(root.Struct()), err
}

func (s Host_view_Results) String() string {
	str, _ := text.Marshal(0x8f58928e854cd4f5, capnp.Struct(s))
	return str
}

func (s Host_view_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Host_view_Results) DecodeFromPtr(p capnp.Ptr) Host_view_Results {
	return Host_view_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Host_view_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Host_view_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Host_view_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Host_view_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Host_view_Results) View() capnp.Client {
	p, _ := capnp.Struct(s).Ptr(0)
	return p.Interface().Client()
}

func (s Host_view_Results) HasView() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Host_view_Results) SetView(c capnp.Client) error {
	if !c.IsValid() {
		return capnp.Struct(s).SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(c))
	return capnp.Struct(s).SetPtr(0, in.ToPtr())
}

// Host_view_Results_List is a list of Host_view_Results.
type Host_view_Results_List = capnp.StructList[Host_view_Results]

// NewHost_view_Results creates a new list of Host_view_Results.
func NewHost_view_Results_List(s *capnp.Segment, sz int32) (Host_view_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Host_view_Results](l), err
}

// Host_view_Results_Future is a wrapper for a Host_view_Results promised by a client call.
type Host_view_Results_Future struct{ *capnp.Future }

func (p Host_view_Results_Future) Struct() (Host_view_Results, error) {
	s, err := p.Future.Struct()
	return Host_view_Results(s), err
}

func (p Host_view_Results_Future) View() *capnp.Future {
	return p.Future.Field(0, nil)
}

type Host_pubSub_Params capnp.Struct

// Host_pubSub_Params_TypeID is the unique identifier for the type Host_pubSub_Params.
const Host_pubSub_Params_TypeID = 0xe5b5227505fcaa99

func NewHost_pubSub_Params(s *capnp.Segment) (Host_pubSub_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Host_pubSub_Params(st), err
}

func NewRootHost_pubSub_Params(s *capnp.Segment) (Host_pubSub_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Host_pubSub_Params(st), err
}

func ReadRootHost_pubSub_Params(msg *capnp.Message) (Host_pubSub_Params, error) {
	root, err := msg.Root()
	return Host_pubSub_Params(root.Struct()), err
}

func (s Host_pubSub_Params) String() string {
	str, _ := text.Marshal(0xe5b5227505fcaa99, capnp.Struct(s))
	return str
}

func (s Host_pubSub_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Host_pubSub_Params) DecodeFromPtr(p capnp.Ptr) Host_pubSub_Params {
	return Host_pubSub_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Host_pubSub_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Host_pubSub_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Host_pubSub_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Host_pubSub_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// Host_pubSub_Params_List is a list of Host_pubSub_Params.
type Host_pubSub_Params_List = capnp.StructList[Host_pubSub_Params]

// NewHost_pubSub_Params creates a new list of Host_pubSub_Params.
func NewHost_pubSub_Params_List(s *capnp.Segment, sz int32) (Host_pubSub_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Host_pubSub_Params](l), err
}

// Host_pubSub_Params_Future is a wrapper for a Host_pubSub_Params promised by a client call.
type Host_pubSub_Params_Future struct{ *capnp.Future }

func (p Host_pubSub_Params_Future) Struct() (Host_pubSub_Params, error) {
	s, err := p.Future.Struct()
	return Host_pubSub_Params(s), err
}

type Host_pubSub_Results capnp.Struct

// Host_pubSub_Results_TypeID is the unique identifier for the type Host_pubSub_Results.
const Host_pubSub_Results_TypeID = 0xdc88f975f5090eee

func NewHost_pubSub_Results(s *capnp.Segment) (Host_pubSub_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_pubSub_Results(st), err
}

func NewRootHost_pubSub_Results(s *capnp.Segment) (Host_pubSub_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_pubSub_Results(st), err
}

func ReadRootHost_pubSub_Results(msg *capnp.Message) (Host_pubSub_Results, error) {
	root, err := msg.Root()
	return Host_pubSub_Results(root.Struct()), err
}

func (s Host_pubSub_Results) String() string {
	str, _ := text.Marshal(0xdc88f975f5090eee, capnp.Struct(s))
	return str
}

func (s Host_pubSub_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Host_pubSub_Results) DecodeFromPtr(p capnp.Ptr) Host_pubSub_Results {
	return Host_pubSub_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Host_pubSub_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Host_pubSub_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Host_pubSub_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Host_pubSub_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Host_pubSub_Results) PubSub() pubsub.PubSub {
	p, _ := capnp.Struct(s).Ptr(0)
	return pubsub.PubSub(p.Interface().Client())
}

func (s Host_pubSub_Results) HasPubSub() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Host_pubSub_Results) SetPubSub(v pubsub.PubSub) error {
	if !v.IsValid() {
		return capnp.Struct(s).SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(capnp.Client(v)))
	return capnp.Struct(s).SetPtr(0, in.ToPtr())
}

// Host_pubSub_Results_List is a list of Host_pubSub_Results.
type Host_pubSub_Results_List = capnp.StructList[Host_pubSub_Results]

// NewHost_pubSub_Results creates a new list of Host_pubSub_Results.
func NewHost_pubSub_Results_List(s *capnp.Segment, sz int32) (Host_pubSub_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Host_pubSub_Results](l), err
}

// Host_pubSub_Results_Future is a wrapper for a Host_pubSub_Results promised by a client call.
type Host_pubSub_Results_Future struct{ *capnp.Future }

func (p Host_pubSub_Results_Future) Struct() (Host_pubSub_Results, error) {
	s, err := p.Future.Struct()
	return Host_pubSub_Results(s), err
}

func (p Host_pubSub_Results_Future) PubSub() pubsub.PubSub {
	return pubsub.PubSub(p.Future.Field(0, nil).Client())
}

const schema_fcf6ac08e448a6ac = "x\xdat\x90=HrQ\x1c\x87\x7f\xff{\xcf\xf1\xea" +
	"\x0b\xfar\xbc.\xc2;\xbd\xe4\xe2 \x89\x9b\x8b\x8eB" +
	"\x0d\xde\\\xdaB\xcd!\xb0\xbax\xee\xa9\xa5\xb5\x8f\xa9" +
	" h\xa8\xa5%\x0a\xc2\xb5\xc1\xa1\x86\xe6\xb6\x1ak\x08" +
	"t\x0d\x1a\x0c\x0a\xe3\x84_\xdcK\xd4|\xefy\x9e\xdf" +
	"\xf3\x9f\xdd\xa1\"\xcbF/\"0\x1c\x97\x87t\xffa" +
	"~{\xffp\xf1\x00\xc2&\x80\x93\x05\xe4:\xcc \x90" +
	"}\xcd\x0a m\xdd/-\x0fn\xb6\x8e b\xa6n" +
	"\x9f\x97\xba\xe1\xf6\xdb\x00 \xfb\x89\x9d\xd8=f\x01\xf6" +
	"3\xdb\xb5\xb3\xdc\x02t\xb7\xa3*s\xb7\xeclL\x1b" +
	"}L\xf2\x0f0\xfd\x12\x8b\xf4\xd5\xfb\xdecP\xf3\xc9" +
	"\xe2C\x0d\xe7C\xcd\xf1\xe5\x80\xab\xffW=\xffe." +
	"\xc5\xff\x104H\xd7\x9bJz\x8dV\xc6\xacW\xdd5" +
	"7_Z\x97^fc\xa5\xb19\xb3\xd0\x90\xaa\xe9I" +
	"8\xccd\x00#@D\xd3\x80\x136\xc9I\x18\xf4w" +
	"\xf8\x13\xc5\x99\x09\xa2x\x80CS\x8e)\xbd2\x91\x13" +
	"6y`;MO\"\xb2i\x18\"e\x91\xbf\x8e\xa6" +
	"\x1d\"\x99\x87!\xa2\xd6HQ\xa4\x82\xabj\x15U+" +
	"R\x99|\x8d\xf1}n\xa1\\mUW\xe5\x8f=c" +
	"\xc0\xa4\x88d\xb0(\xef\x17M<$t\xac\x17\xfaw" +
	"\x9a\xb8{\x05\x88\xc4/7\x9a0\xc7V\xe0+\x00\x00" +
	"\xff\xff\xdb6\x94Q"

func init() {
	schemas.Register(schema_fcf6ac08e448a6ac,
		0x8f58928e854cd4f5,
		0x957cbefc645fd307,
		0xa404c24b5375b9e4,
		0xdc88f975f5090eee,
		0xe5b5227505fcaa99)
}
