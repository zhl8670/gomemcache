package memcache

// Error represents an error returned in a command reply.
type Error string

func (err Error) Error() string { return string(err) }

// Item is an item to be got or stored in a memcached server.
type Reply struct {
	// Key is the Item's key (250 bytes maximum).
	Key string

	// Value is the Item's value.
	Value []byte

	// Flags are server-opaque flags whose semantics are entirely
	// up to the app.
	Flags uint32

	// Expiration is the cache expiration time, in seconds: either a relative
	// time from now (up to 1 month), or an absolute Unix epoch time.
	// Zero means the Item has no expiration time.
	Expiration int32

	// Compare and swap ID.
	Cas uint64
}

// Conn represents a connection to a Memcache server.
type Conn interface {
	// Close closes the connection.
	Close() error

	// Err returns a non-nil value if the connection is broken. The returned
	// value is either the first non-nil value returned from the underlying
	// network connection or a protocol parsing error. Applications should
	// close broken connections.
	Err() error

	// Store sends a command to the server for store data.
	Store(cmd, key string, value []byte, flags uint32, timeout int32, cas uint64) (err error)

	// Store sends a command to the server for gets data.
	Get(cmd string, keys ...string) (replies []*Reply, err error)

	// Store sends a command to the server for delete data.
	Delete(keys ...string) (err error)

	// IncrDecr sends a command to the server for incr/decr data.
	IncrDecr(cmd string, key string, delta uint64) (uint64, error)
}
