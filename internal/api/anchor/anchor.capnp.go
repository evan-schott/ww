// Code generated by capnpc-go. DO NOT EDIT.

package anchor

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	fc "capnproto.org/go/capnp/v3/flowcontrol"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	context "context"
	fmt "fmt"
)

type Anchor capnp.Client

// Anchor_TypeID is the unique identifier for the type Anchor.
const Anchor_TypeID = 0xe41237e4098ed922

func (c Anchor) Ls(ctx context.Context, params func(Anchor_ls_Params) error) (Anchor_ls_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xe41237e4098ed922,
			MethodID:      0,
			InterfaceName: "anchor.capnp:Anchor",
			MethodName:    "ls",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Anchor_ls_Params(s)) }
	}
	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Anchor_ls_Results_Future{Future: ans.Future()}, release
}
func (c Anchor) Walk(ctx context.Context, params func(Anchor_walk_Params) error) (Anchor_walk_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xe41237e4098ed922,
			MethodID:      1,
			InterfaceName: "anchor.capnp:Anchor",
			MethodName:    "walk",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Anchor_walk_Params(s)) }
	}
	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Anchor_walk_Results_Future{Future: ans.Future()}, release
}

// String returns a string that identifies this capability for debugging
// purposes.  Its format should not be depended on: in particular, it
// should not be used to compare clients.  Use IsSame to compare clients
// for equality.
func (c Anchor) String() string {
	return fmt.Sprintf("%T(%v)", c, capnp.Client(c))
}

// AddRef creates a new Client that refers to the same capability as c.
// If c is nil or has resolved to null, then AddRef returns nil.
func (c Anchor) AddRef() Anchor {
	return Anchor(capnp.Client(c).AddRef())
}

// Release releases a capability reference.  If this is the last
// reference to the capability, then the underlying resources associated
// with the capability will be released.
//
// Release will panic if c has already been released, but not if c is
// nil or resolved to null.
func (c Anchor) Release() {
	capnp.Client(c).Release()
}

// Resolve blocks until the capability is fully resolved or the Context
// expires.
func (c Anchor) Resolve(ctx context.Context) error {
	return capnp.Client(c).Resolve(ctx)
}

func (c Anchor) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Client(c).EncodeAsPtr(seg)
}

func (Anchor) DecodeFromPtr(p capnp.Ptr) Anchor {
	return Anchor(capnp.Client{}.DecodeFromPtr(p))
}

// IsValid reports whether c is a valid reference to a capability.
// A reference is invalid if it is nil, has resolved to null, or has
// been released.
func (c Anchor) IsValid() bool {
	return capnp.Client(c).IsValid()
}

// IsSame reports whether c and other refer to a capability created by the
// same call to NewClient.  This can return false negatives if c or other
// are not fully resolved: use Resolve if this is an issue.  If either
// c or other are released, then IsSame panics.
func (c Anchor) IsSame(other Anchor) bool {
	return capnp.Client(c).IsSame(capnp.Client(other))
}

// Update the flowcontrol.FlowLimiter used to manage flow control for
// this client. This affects all future calls, but not calls already
// waiting to send. Passing nil sets the value to flowcontrol.NopLimiter,
// which is also the default.
func (c Anchor) SetFlowLimiter(lim fc.FlowLimiter) {
	capnp.Client(c).SetFlowLimiter(lim)
}

// Get the current flowcontrol.FlowLimiter used to manage flow control
// for this client.
func (c Anchor) GetFlowLimiter() fc.FlowLimiter {
	return capnp.Client(c).GetFlowLimiter()
} // A Anchor_Server is a Anchor with a local implementation.
type Anchor_Server interface {
	Ls(context.Context, Anchor_ls) error

	Walk(context.Context, Anchor_walk) error
}

// Anchor_NewServer creates a new Server from an implementation of Anchor_Server.
func Anchor_NewServer(s Anchor_Server) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Anchor_Methods(nil, s), s, c)
}

// Anchor_ServerToClient creates a new Client from an implementation of Anchor_Server.
// The caller is responsible for calling Release on the returned Client.
func Anchor_ServerToClient(s Anchor_Server) Anchor {
	return Anchor(capnp.NewClient(Anchor_NewServer(s)))
}

// Anchor_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Anchor_Methods(methods []server.Method, s Anchor_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe41237e4098ed922,
			MethodID:      0,
			InterfaceName: "anchor.capnp:Anchor",
			MethodName:    "ls",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Ls(ctx, Anchor_ls{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe41237e4098ed922,
			MethodID:      1,
			InterfaceName: "anchor.capnp:Anchor",
			MethodName:    "walk",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Walk(ctx, Anchor_walk{call})
		},
	})

	return methods
}

// Anchor_ls holds the state for a server call to Anchor.ls.
// See server.Call for documentation.
type Anchor_ls struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Anchor_ls) Args() Anchor_ls_Params {
	return Anchor_ls_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Anchor_ls) AllocResults() (Anchor_ls_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Anchor_ls_Results(r), err
}

// Anchor_walk holds the state for a server call to Anchor.walk.
// See server.Call for documentation.
type Anchor_walk struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Anchor_walk) Args() Anchor_walk_Params {
	return Anchor_walk_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Anchor_walk) AllocResults() (Anchor_walk_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Anchor_walk_Results(r), err
}

// Anchor_List is a list of Anchor.
type Anchor_List = capnp.CapList[Anchor]

// NewAnchor creates a new list of Anchor.
func NewAnchor_List(s *capnp.Segment, sz int32) (Anchor_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[Anchor](l), err
}

type Anchor_ls_Params capnp.Struct

// Anchor_ls_Params_TypeID is the unique identifier for the type Anchor_ls_Params.
const Anchor_ls_Params_TypeID = 0xc105d085735711e1

func NewAnchor_ls_Params(s *capnp.Segment) (Anchor_ls_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Anchor_ls_Params(st), err
}

func NewRootAnchor_ls_Params(s *capnp.Segment) (Anchor_ls_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Anchor_ls_Params(st), err
}

func ReadRootAnchor_ls_Params(msg *capnp.Message) (Anchor_ls_Params, error) {
	root, err := msg.Root()
	return Anchor_ls_Params(root.Struct()), err
}

func (s Anchor_ls_Params) String() string {
	str, _ := text.Marshal(0xc105d085735711e1, capnp.Struct(s))
	return str
}

func (s Anchor_ls_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Anchor_ls_Params) DecodeFromPtr(p capnp.Ptr) Anchor_ls_Params {
	return Anchor_ls_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Anchor_ls_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Anchor_ls_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Anchor_ls_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Anchor_ls_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// Anchor_ls_Params_List is a list of Anchor_ls_Params.
type Anchor_ls_Params_List = capnp.StructList[Anchor_ls_Params]

// NewAnchor_ls_Params creates a new list of Anchor_ls_Params.
func NewAnchor_ls_Params_List(s *capnp.Segment, sz int32) (Anchor_ls_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Anchor_ls_Params](l), err
}

// Anchor_ls_Params_Future is a wrapper for a Anchor_ls_Params promised by a client call.
type Anchor_ls_Params_Future struct{ *capnp.Future }

func (f Anchor_ls_Params_Future) Struct() (Anchor_ls_Params, error) {
	p, err := f.Future.Ptr()
	return Anchor_ls_Params(p.Struct()), err
}

type Anchor_ls_Results capnp.Struct

// Anchor_ls_Results_TypeID is the unique identifier for the type Anchor_ls_Results.
const Anchor_ls_Results_TypeID = 0xe325af947f127758

func NewAnchor_ls_Results(s *capnp.Segment) (Anchor_ls_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Anchor_ls_Results(st), err
}

func NewRootAnchor_ls_Results(s *capnp.Segment) (Anchor_ls_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Anchor_ls_Results(st), err
}

func ReadRootAnchor_ls_Results(msg *capnp.Message) (Anchor_ls_Results, error) {
	root, err := msg.Root()
	return Anchor_ls_Results(root.Struct()), err
}

func (s Anchor_ls_Results) String() string {
	str, _ := text.Marshal(0xe325af947f127758, capnp.Struct(s))
	return str
}

func (s Anchor_ls_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Anchor_ls_Results) DecodeFromPtr(p capnp.Ptr) Anchor_ls_Results {
	return Anchor_ls_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Anchor_ls_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Anchor_ls_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Anchor_ls_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Anchor_ls_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Anchor_ls_Results) Names() (capnp.TextList, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return capnp.TextList(p.List()), err
}

func (s Anchor_ls_Results) HasNames() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Anchor_ls_Results) SetNames(v capnp.TextList) error {
	return capnp.Struct(s).SetPtr(0, v.ToPtr())
}

// NewNames sets the names field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Anchor_ls_Results) NewNames(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(capnp.Struct(s).Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = capnp.Struct(s).SetPtr(0, l.ToPtr())
	return l, err
}
func (s Anchor_ls_Results) Children() (Anchor_List, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return Anchor_List(p.List()), err
}

func (s Anchor_ls_Results) HasChildren() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Anchor_ls_Results) SetChildren(v Anchor_List) error {
	return capnp.Struct(s).SetPtr(1, v.ToPtr())
}

// NewChildren sets the children field to a newly
// allocated Anchor_List, preferring placement in s's segment.
func (s Anchor_ls_Results) NewChildren(n int32) (Anchor_List, error) {
	l, err := NewAnchor_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Anchor_List{}, err
	}
	err = capnp.Struct(s).SetPtr(1, l.ToPtr())
	return l, err
}

// Anchor_ls_Results_List is a list of Anchor_ls_Results.
type Anchor_ls_Results_List = capnp.StructList[Anchor_ls_Results]

// NewAnchor_ls_Results creates a new list of Anchor_ls_Results.
func NewAnchor_ls_Results_List(s *capnp.Segment, sz int32) (Anchor_ls_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return capnp.StructList[Anchor_ls_Results](l), err
}

// Anchor_ls_Results_Future is a wrapper for a Anchor_ls_Results promised by a client call.
type Anchor_ls_Results_Future struct{ *capnp.Future }

func (f Anchor_ls_Results_Future) Struct() (Anchor_ls_Results, error) {
	p, err := f.Future.Ptr()
	return Anchor_ls_Results(p.Struct()), err
}

type Anchor_walk_Params capnp.Struct

// Anchor_walk_Params_TypeID is the unique identifier for the type Anchor_walk_Params.
const Anchor_walk_Params_TypeID = 0xb90ffa2761585171

func NewAnchor_walk_Params(s *capnp.Segment) (Anchor_walk_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Anchor_walk_Params(st), err
}

func NewRootAnchor_walk_Params(s *capnp.Segment) (Anchor_walk_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Anchor_walk_Params(st), err
}

func ReadRootAnchor_walk_Params(msg *capnp.Message) (Anchor_walk_Params, error) {
	root, err := msg.Root()
	return Anchor_walk_Params(root.Struct()), err
}

func (s Anchor_walk_Params) String() string {
	str, _ := text.Marshal(0xb90ffa2761585171, capnp.Struct(s))
	return str
}

func (s Anchor_walk_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Anchor_walk_Params) DecodeFromPtr(p capnp.Ptr) Anchor_walk_Params {
	return Anchor_walk_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Anchor_walk_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Anchor_walk_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Anchor_walk_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Anchor_walk_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Anchor_walk_Params) Path() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Anchor_walk_Params) HasPath() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Anchor_walk_Params) PathBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Anchor_walk_Params) SetPath(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

// Anchor_walk_Params_List is a list of Anchor_walk_Params.
type Anchor_walk_Params_List = capnp.StructList[Anchor_walk_Params]

// NewAnchor_walk_Params creates a new list of Anchor_walk_Params.
func NewAnchor_walk_Params_List(s *capnp.Segment, sz int32) (Anchor_walk_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Anchor_walk_Params](l), err
}

// Anchor_walk_Params_Future is a wrapper for a Anchor_walk_Params promised by a client call.
type Anchor_walk_Params_Future struct{ *capnp.Future }

func (f Anchor_walk_Params_Future) Struct() (Anchor_walk_Params, error) {
	p, err := f.Future.Ptr()
	return Anchor_walk_Params(p.Struct()), err
}

type Anchor_walk_Results capnp.Struct

// Anchor_walk_Results_TypeID is the unique identifier for the type Anchor_walk_Results.
const Anchor_walk_Results_TypeID = 0xaec21d58779cc86c

func NewAnchor_walk_Results(s *capnp.Segment) (Anchor_walk_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Anchor_walk_Results(st), err
}

func NewRootAnchor_walk_Results(s *capnp.Segment) (Anchor_walk_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Anchor_walk_Results(st), err
}

func ReadRootAnchor_walk_Results(msg *capnp.Message) (Anchor_walk_Results, error) {
	root, err := msg.Root()
	return Anchor_walk_Results(root.Struct()), err
}

func (s Anchor_walk_Results) String() string {
	str, _ := text.Marshal(0xaec21d58779cc86c, capnp.Struct(s))
	return str
}

func (s Anchor_walk_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Anchor_walk_Results) DecodeFromPtr(p capnp.Ptr) Anchor_walk_Results {
	return Anchor_walk_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Anchor_walk_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Anchor_walk_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Anchor_walk_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Anchor_walk_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Anchor_walk_Results) Anchor() Anchor {
	p, _ := capnp.Struct(s).Ptr(0)
	return Anchor(p.Interface().Client())
}

func (s Anchor_walk_Results) HasAnchor() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Anchor_walk_Results) SetAnchor(v Anchor) error {
	if !v.IsValid() {
		return capnp.Struct(s).SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(capnp.Client(v)))
	return capnp.Struct(s).SetPtr(0, in.ToPtr())
}

// Anchor_walk_Results_List is a list of Anchor_walk_Results.
type Anchor_walk_Results_List = capnp.StructList[Anchor_walk_Results]

// NewAnchor_walk_Results creates a new list of Anchor_walk_Results.
func NewAnchor_walk_Results_List(s *capnp.Segment, sz int32) (Anchor_walk_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Anchor_walk_Results](l), err
}

// Anchor_walk_Results_Future is a wrapper for a Anchor_walk_Results promised by a client call.
type Anchor_walk_Results_Future struct{ *capnp.Future }

func (f Anchor_walk_Results_Future) Struct() (Anchor_walk_Results, error) {
	p, err := f.Future.Ptr()
	return Anchor_walk_Results(p.Struct()), err
}
func (p Anchor_walk_Results_Future) Anchor() Anchor {
	return Anchor(p.Future.Field(0, nil).Client())
}

const schema_efb5a91f96d44de3 = "x\xda|\x921h\xd4P\x1c\xc6\xbf\xef\xbd\x9c9\xa5" +
	"\xb9\xf35\xb7(H@R\xa4E\x85\xd3A(\xca\xc5" +
	"U\x10\xf2\xba\x98\xf5q\x1eDL\xd33i\xc9T\x9c" +
	"\xc4\xcdI\x9ctPp\xe8\xa2\x93\x83c\x9d\x1c\x1d\\" +
	"\x04\x17iw\xc1\xb1K$\x89\xe9\xa5\x83\x1d\xdf\xfb\xbe" +
	"\xff\xef\xff\xbd\xf7\xff\x9f\xff\x19Xc\xc7\x93\x10\xda\xef" +
	"\x9d)\x93\xaf\xaf\x8b\xe8\xd2\x97\x0fP#\x02=\xda\xc0" +
	"\xcd\x1d\x9e#\xe8\xeer\x02\x96Otd\xae\x1c\x0d?" +
	"w\x0do(*\xc3\xdb\xda\xf0K=\xc8\x9f}\xeb\xed" +
	"7\x06\xcb\x06\xdc}\xfe\x81UF\xc5\xf2\xd3\x97\x1fW" +
	"\x0e\xfeU\x8aJy\xcf#\xd0\xddc\x01\x96\x97\x7f\xbc" +
	"8{xk\xf9\x10\xca\x91\xe5\xc1\xfd\xef\xaf\xbc\xbdO" +
	"\xbf\x01\xba\xab\xe2\x9d;\xae\xed\xd7\xc4swW\xd8\xb8" +
	"Z\x9at\x1aoe\xd7\xa7\xd2\xcc\xd3\xf9\xfa\xdd\xe6T" +
	"\x98\xe4\xb1\xbf1\xcbw\x92\xed\x1c\xd0\x96\xb4\x00\x8b\x80" +
	"r\xd6\x01\xdd\x97\xd4#\xc1ISJ\xb5h\x08R\x81" +
	"\xa70C\x93\x99\xcd\xfc\x04qmA\x1c\xce\xcdv\xcc" +
	"%\x08.u(\xa2KIr?\xf4j\xc8\x7f\xf5\x8d" +
	"I\x13\\\xf7\x8f\x9b\xac\xde\x00\xb4/\xa9\x03AE\x8e" +
	"X]\xde\xb9\x07\xe8\xdb\x92:\x12\xf4R\xb39\xcb9" +
	"\x00C\xc9:\xc1\x00,\xa7\xf1\xa3\xe4a6K\x01\xb4" +
	"\xd2\xc9\xc7\x0e:1\xd9\xc6\xb0\xe3\xad,$u_\xf6" +
	"\x80\xe3)\xb2\x9d\x9a\x1a_\x84P+6\x17+\xc0v" +
	"Y\xd4\x855\x08\xe5\xd82\xc9\x03\x0e\xab\x1f\x0b\x18\x92" +
	"\x7f\x03\x00\x00\xff\xff\x8fH\x99\x1a"

func init() {
	schemas.Register(schema_efb5a91f96d44de3,
		0xaec21d58779cc86c,
		0xb90ffa2761585171,
		0xc105d085735711e1,
		0xe325af947f127758,
		0xe41237e4098ed922)
}
