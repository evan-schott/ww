package anchor

import (
	"context"

	"capnproto.org/go/capnp/v3"

	casm "github.com/wetware/casm/pkg"
	api "github.com/wetware/ww/internal/api/anchor"
	"github.com/wetware/ww/pkg/internal/bounded"
)

type Anchor api.Anchor

func (a Anchor) AddRef() Anchor {
	return Anchor(api.Anchor(a).AddRef())
}

func (a Anchor) Release() {
	capnp.Client(a).Release()
}

func (a Anchor) Ls(ctx context.Context) (Iterator, capnp.ReleaseFunc) {
	f, release := api.Anchor(a).Ls(ctx, nil)

	h := &sequence{
		Future: casm.Future(f),
		pos:    -1,
	}

	return Iterator{
		Seq:    h,
		Future: h,
	}, release
}

// Walk to the register located at path.
func (a Anchor) Walk(ctx context.Context, path string) (Anchor, capnp.ReleaseFunc) {
	p := NewPath(path)

	if p.IsRoot() {
		return Anchor(a), a.AddRef().Release
	}

	f, release := api.Anchor(a).Walk(ctx, destination(p))
	return Anchor(f.Anchor()), release
}

type Child interface {
	String() string
	Anchor() Anchor
}

type Iterator casm.Iterator[Child]

func (it Iterator) Next() Child {
	r, _ := it.Seq.Next()
	return r
}

type sequence struct {
	casm.Future
	err error
	pos int
}

func (seq *sequence) Err() error {
	if seq.err == nil {
		select {
		case <-seq.Future.Done():
			_, seq.err = seq.Struct()
		default:
		}
	}

	return seq.err
}

func (seq *sequence) Next() (Child, bool) {
	if ok := seq.advance(); ok {
		return seq, true
	}

	return nil, false
}

func (seq *sequence) String() string {
	names, err := seq.results().Names()
	if err != nil {
		panic(err) // already validated; should never fail
	}

	name, err := names.At(seq.pos)
	if err != nil {
		panic(err)
	}

	return name
}

func (seq *sequence) Anchor() Anchor {
	children, err := seq.results().Children()
	if err != nil {
		panic(err) // already validated; should never fail
	}

	anchor, err := children.At(seq.pos)
	if err != nil {
		panic(err)
	}

	return Anchor(anchor)
}

func (seq *sequence) advance() (ok bool) {
	if ok = seq.Err() == nil; ok {
		seq.pos++
		ok = seq.validate()
	}

	return
}

func (seq *sequence) validate() bool {
	var s capnp.Struct
	if s, seq.err = seq.Struct(); seq.err != nil {
		return false
	}

	res := api.Anchor_ls_Results(s)
	if !(res.HasNames() && res.HasChildren()) {
		return false
	}

	return seq.validateName(res) && seq.validateChildren(res)
}

func (seq *sequence) validateName(res api.Anchor_ls_Results) (ok bool) {
	if ok = res.HasNames(); ok {
		names, err := res.Names()
		if ok = err == nil; ok {
			_, err = names.At(seq.pos)
			ok = err == nil
		}
	}

	return
}

func (seq *sequence) validateChildren(res api.Anchor_ls_Results) (ok bool) {
	if ok = res.HasChildren(); ok {
		children, err := res.Children()
		if ok = err == nil; ok {
			_, err = children.At(seq.pos)
			ok = err == nil
		}
	}

	return
}

func (seq *sequence) results() api.Anchor_ls_Results {
	res, err := seq.Struct()
	if err != nil {
		panic(err)
	}

	return api.Anchor_ls_Results(res)
}

func destination(path Path) func(api.Anchor_walk_Params) error {
	return func(ps api.Anchor_walk_Params) error {
		return path.bind(func(s string) bounded.Type[string] {
			err := ps.SetPath(trimmed(s))
			return bounded.Failure[string](err) // can be nil
		}).Err()
	}
}
