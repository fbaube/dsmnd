package dsmnd

// NSPath is pretty generic and can represent just about anything.
// Examples:
//   - Imported file: batch nr + abs/rel.path
//   - Disk file: PC FQN + abs.pah
//   - Facet value: facet FQN + hiercal path
//
// .
type NSPath struct {
	NS   string
	Path string
}
