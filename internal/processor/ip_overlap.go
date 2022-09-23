// Package processor contains processor to get CIDRs network's relation.
package processor

import (
	"errors"
	"fmt"
	"net/netip"
)

// Network possible relations.
const (
	SubsetRelation    NetworkRelation = "subset"
	SupersetRelation  NetworkRelation = "superset"
	DifferentRelation NetworkRelation = "different"
	SameRelation      NetworkRelation = "same"
)

// ErrNotIPv4 when given IP is not an IPv4.
var ErrNotIPv4 = errors.New("not an IPv4 IP")

type (
	// NetworkRelation for network relations.
	NetworkRelation string

	// Processor check the relation between two CIDRs.
	Processor struct{}
)

// NewProcessor returns a new instance of Processor.
func NewProcessor() *Processor {
	return &Processor{}
}

// GetOverlapRelation check the relation between two CIDRs
// The relation can be:
// 	subset: if the network of the second address is included in the first one
// 	superset: if the network of the second address includes the first one
// 	different: if the two networks are not overlapping
// 	same: if both address are in the same network
// Returns relation or an error if given IP is not an IPv4.
func (p *Processor) GetOverlapRelation(cidrA, cidrB string) (NetworkRelation, error) {
	var (
		networkA, networkB *netip.Prefix
		err                error
	)
	// Check if fist CIDR is correct IP
	networkA, err = getIPv4Network(cidrA)
	if err != nil {
		return "", fmt.Errorf("first CIDR invalid: %w", err)
	}

	// Check if second CIDR is correct IP
	networkB, err = getIPv4Network(cidrB)
	if err != nil {
		return "", fmt.Errorf("second CIDR invalid: %w", err)
	}

	// Check CIDRs relation.
	bInA := networkA.Contains(networkB.Addr())
	aInB := networkB.Contains(networkA.Addr())

	if aInB && bInA {
		return SameRelation, nil
	}
	if aInB && !bInA {
		return SupersetRelation, nil
	}
	if bInA && !aInB {
		return SubsetRelation, nil
	}

	return DifferentRelation, nil
}

// getIPv4Network check if givent CIDR is a correct IPv4 and
// returns netip.Prefix from given IP
func getIPv4Network(cidr string) (*netip.Prefix, error) {
	// Check if CIDR is correct IP.
	network, err := netip.ParsePrefix(cidr)
	if err != nil {
		return nil, fmt.Errorf("unable to parse CIDR: %w", err)
	}
	// Check if IPv4.
	if !network.Addr().Is4() {
		return nil, ErrNotIPv4
	}

	return &network, nil
}
