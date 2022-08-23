package cluster

import (
	"context"
	"sync"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/wetware/casm/pkg/cluster/routing"
	"github.com/wetware/ww/internal/api/cluster"
	"github.com/wetware/ww/pkg/vat"
	"github.com/wetware/ww/pkg/vat/cap/anchor"
)

var HostCapability = vat.BasicCap{
	"host/packed",
	"host"}

// RoutingTable provides a global view of namespace peers.
type RoutingTable interface {
	Iter() routing.Iterator
	Lookup(peer.ID) (routing.Record, bool)
}

/*----------------------------*
|                             |
|    Client Implementations   |
|                             |
*-----------------------------*/
type Dialer interface {
	Dial(context.Context, peer.AddrInfo) (*rpc.Conn, error)
}

type Host struct {
	once   sync.Once
	Client capnp.Client
	Info   peer.AddrInfo
}

func (h *Host) Ls(ctx context.Context, d Dialer) (*anchor.Iterator, capnp.ReleaseFunc) {
	return anchor.Anchor(h.resolve(ctx, d)).Ls(ctx)
}

// Walk to the register located at path.  Panics if len(path) == 0.
func (h *Host) Walk(ctx context.Context, d Dialer, path anchor.Path) (anchor.Anchor, capnp.ReleaseFunc) {
	return anchor.Anchor(h.resolve(ctx, d)).Walk(ctx, path)
}

func (h *Host) resolve(ctx context.Context, d Dialer) cluster.Host {
	h.once.Do(func() {
		if h.Client == (capnp.Client{}) {
			if conn, err := d.Dial(ctx, h.Info); err != nil {
				h.Client = capnp.ErrorClient(err)
			} else {
				h.Client = conn.Bootstrap(ctx) // TODO:  wrap Client & call conn.Close() on Shutdown() hook?
			}
		}
	})

	return cluster.Host(h.Client)
}

/*---------------------------*
|                            |
|    Server Implementation   |
|                            |
*----------------------------*/

// HostServer represents a host instance on the network. It provides
// the Anchor and Joiner capabilities.
//
// The zero-value HostServer is ready to use.
type HostServer struct {
	once sync.Once

	// RoutingTable provides a global view of namespace peers.
	// Callers MUST set this value before the first call to Client()
	RoutingTable

	// The root anchor for the HostServer.  Users SHOULD NOT
	// set this field; it will be populated automatically on
	// the first call to Client().
	anchor.AnchorServer
}

func (h *HostServer) Client() capnp.Client {
	h.once.Do(func() {
		if h.AnchorServer.Path().IsZero() {
			h.AnchorServer = anchor.Root(h)
		}
	})

	return capnp.Client(cluster.Host_ServerToClient(h))
}
