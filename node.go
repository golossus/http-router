package routing

import (
	"net/http"
)

type nodeInterface interface {
	find(p string) nodeInterface
	merge(n nodeInterface) nodeInterface
	hasParameters() bool
	setParent(parent nodeInterface)
	addChild(child nodeInterface) nodeInterface
	sibling() nodeInterface
	setSibling(sibling nodeInterface)
	child() nodeInterface
	handler() http.HandlerFunc
	setHandler(handler http.HandlerFunc)
	getPrefix() string
	getParent() nodeInterface
	getWeight() int
}

func createNodeFromChunk(c chunk, h http.HandlerFunc) nodeInterface {
	switch c.t {
	case tChunkStatic:
		return newNodeStatic(c.v, h)
	case tChunkDynamic:
		return newNodeDynamic(c.v, c.exp, h)
	default:
		return nil
	}
}
