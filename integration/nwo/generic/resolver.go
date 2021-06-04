/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package generic

import (
	"fmt"

	registry2 "github.com/hyperledger-labs/fabric-smart-client/integration/nwo/common/registry"
)

const (
	ListenPort registry2.PortName = "Listen"
	ViewPort   registry2.PortName = "View"
	P2PPort    registry2.PortName = "P2P"
)

type ResolverIdentity struct {
	ID   string
	Path string
}

type Resolver struct {
	Name      string
	Domain    string
	Identity  ResolverIdentity
	Addresses map[registry2.PortName]string
	Port      int
}

func (p *platform) GenerateResolverMap() {
	p.Resolvers = []*Resolver{}
	for _, peer := range p.Peers {
		org := p.Organization(peer.Organization)

		addresses := map[registry2.PortName]string{
			ViewPort:   fmt.Sprintf("127.0.0.1:%d", p.Registry.PortsByPeerID[peer.Name][ListenPort]),
			ListenPort: fmt.Sprintf("127.0.0.1:%d", p.Registry.PortsByPeerID[peer.Name][ListenPort]),
			P2PPort:    fmt.Sprintf("127.0.0.1:%d", p.Registry.PortsByPeerID[peer.Name][P2PPort]),
		}

		p.Resolvers = append(p.Resolvers, &Resolver{
			Name: peer.Name,
			Identity: ResolverIdentity{
				ID:   peer.Name,
				Path: p.PeerLocalMSPIdentityCert(peer),
			},
			Domain:    org.Domain,
			Addresses: addresses,
		})
	}
}
