//go:build go1.26

package handshake

import "github.com/metacubex/tls"

const quicErrorEvent tls.QUICEventKind = tls.QUICErrorEvent

func extractQUICEventError(ev tls.QUICEvent) error {
	return ev.Err
}
