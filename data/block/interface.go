package block

import "github.com/ThotaGopichandThota/gn-core/data"

// EmptyBlockCreator is able to create empty block instances
type EmptyBlockCreator interface {
	CreateNewHeader() data.HeaderHandler
	IsInterfaceNil() bool
}
