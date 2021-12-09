package client

import (
	"github.com/lthibault/log"
	"github.com/wetware/ww/pkg/cap"
)

type Option func(*Dialer)

func WithNamespace(ns string) Option {
	if ns == "" {
		ns = "ww"
	}

	return func(d *Dialer) {
		d.ns = ns
	}
}

func WithLogger(l log.Logger) Option {
	if l == nil {
		l = log.New()
	}

	return func(d *Dialer) {
		d.log = l
	}
}

func WithHost(h HostFactory) Option {
	if h == nil {
		h = &RoutedHostFactory{}
	}

	return func(d *Dialer) {
		d.host = h
	}
}

func WithRouting(r RoutingFactory) Option {
	if r == nil {
		r = defaultRoutingFactory{}
	}

	return func(d *Dialer) {
		d.routing = r
	}
}

func WithPubSub(p PubSubFactory) Option {
	if p == nil {
		p = defaultPubSubFactory{}
	}

	return func(d *Dialer) {
		d.pubsub = p
	}
}

func WithCapability(c cap.Dialer) Option {
	return func(d *Dialer) {
		if c == nil {
			c = BasicCapDialer{NS: d.ns}
		}

		d.cap = c
	}
}

func withDefault(opt []Option) []Option {
	return append([]Option{
		WithNamespace(""),
		WithLogger(nil),
		WithHost(nil),
		WithRouting(nil),
		WithPubSub(nil),
		WithCapability(nil),
	}, opt...)
}
