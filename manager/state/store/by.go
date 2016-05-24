package store

import "github.com/docker/swarm-v2/api"

// By is an interface type passed to Find methods. Implementations must be
// defined in this package.
type By interface {
	// isBy allows this interface to only be satisfied by certain internal
	// types.
	isBy()
}

type byAll struct{}

func (a byAll) isBy() {
}

// All is an argument that can be passed to find to list all items in the
// set.
var All byAll

type byName string

func (b byName) isBy() {
}

// ByName creates an object to pass to Find to select by name.
func ByName(name string) By {
	return byName(name)
}

type byCN string

func (b byCN) isBy() {
}

// ByCN creates an object to pass to Find to select by CN.
func ByCN(name string) By {
	return byCN(name)
}

type byService string

func (b byService) isBy() {
}

// ByServiceID creates an object to pass to Find to select by service.
func ByServiceID(serviceID string) By {
	return byService(serviceID)
}

type byNode string

func (b byNode) isBy() {
}

// ByNodeID creates an object to pass to Find to select by node.
func ByNodeID(nodeID string) By {
	return byNode(nodeID)
}

type byQuery string

func (b byQuery) isBy() {
}

// ByQuery creates an object to pass to Find to select by query.
func ByQuery(query string) By {
	return byQuery(query)
}

type byInstance struct {
	serviceID string
	instance  uint64
}

func (b byInstance) isBy() {
}

// ByInstance creates an object to pass to Find to select by instance number.
func ByInstance(serviceID string, instance uint64) By {
	return byInstance{serviceID: serviceID, instance: instance}
}

type byIssuanceState api.IssuanceState

func (b byIssuanceState) isBy() {
}

// ByIssuanceState creates an object to pass to Find to select by instance number.
func ByIssuanceState(state api.IssuanceState) By {
	return byIssuanceState(state)
}
