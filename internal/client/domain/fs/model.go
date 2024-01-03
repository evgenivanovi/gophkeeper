package fs

/* __________________________________________________ */

type CreationSettings struct {
	Path    string
	Options CreationOptions
}

func NewCreationSettings(
	path string, ops ...CreationOp,
) CreationSettings {
	options := NewCreationOptions(ops...)
	return CreationSettings{
		Path:    path,
		Options: options,
	}
}

/* __________________________________________________ */

type CreationOptions struct {
	Force bool
}

func NewCreationOptions(ops ...CreationOp) CreationOptions {
	options := new(CreationOptions)
	for _, op := range ops {
		op(options)
	}
	return *options
}

type CreationOp func(*CreationOptions)

func WithForceCreation(force bool) CreationOp {
	return func(options *CreationOptions) {
		options.Force = force
	}
}

/* __________________________________________________ */
