package pubsub_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	inproc "github.com/lthibault/go-libp2p-inproc-transport"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"

	pscap "github.com/wetware/ww/pkg/pubsub"
)

func TestPubSub(t *testing.T) {
	t.Parallel()

	const topic = "test"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gs, release := newGossipSub(ctx)
	defer release()

	ps := (&pscap.Router{TopicJoiner: gs}).PubSub()
	defer ps.Release()

	const nmsg = 10
	g, ctx := errgroup.WithContext(ctx)

	// writer
	g.Go(func() error {
		topic, release := ps.Join(ctx, topic)
		defer release()

		// Publish 10 messages
		g, ctx := errgroup.WithContext(ctx)
		for i := 0; i < nmsg; i++ {
			g.Go(func() error {
				return topic.Publish(ctx, []byte("test"))
			})
		}

		return annotate("writer", g.Wait())
	})

	// reader
	g.Go(func() error {
		topic, release := ps.Join(ctx, topic)
		defer release()

		sub, cancel := topic.Subscribe(ctx)
		defer cancel()

		for got := sub.Next(); got != nil; got = sub.Next() {
			if string(got) != "test" {
				return fmt.Errorf("reader: unexpected message: %s", got)
			}
		}

		return nil
	})

	assert.NoError(t, g.Wait())
}

func annotate(prefix string, err error) error {
	if err != nil {
		err = fmt.Errorf("%s: %w", prefix, err)
	}

	return err
}

func newGossipSub(ctx context.Context) (*pubsub.PubSub, func()) {
	h := newTestHost()

	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}

	return ps, func() { h.Close() }
}

func newTestHost() host.Host {
	h, err := libp2p.New(
		libp2p.NoListenAddrs,
		libp2p.NoTransports,
		libp2p.Transport(inproc.New()),
		libp2p.ListenAddrStrings("/inproc/~"))
	if err != nil {
		panic(err)
	}

	return h
}
