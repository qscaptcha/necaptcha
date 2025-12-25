package necaptcha

// Options configures the Pool struct.
type Options struct {
	// Initial creates an initial number of ready-to-use items in the pool.
	Initial uint32

	// Max represents the maximum number of items you can borrow at a time. This
	// prevents unbounded growth in the pool.
	//
	// Depending on the timing of Returns and Factory calls, the maximum number of
	// items in the pool can exceed Max by a small number for a short time.
	Max uint32

	// EnableCount, when set, enables the pool's Count function.
	//
	// NOTE: If you set this AND you need to set your own runtime Finalizer on the item,
	// wrap your item in another struct, with the Finalizer added to the inner item.
	EnableCount bool
}
