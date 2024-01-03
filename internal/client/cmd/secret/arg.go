package secret

/* __________________________________________________ */

type CreateSecretArg struct {
	Name string
}

type CreateBinarySecretArg struct {
	CreateSecretArg
	Path string
}

/* __________________________________________________ */

type GetSecretArg struct {
	Name string
}

/* __________________________________________________ */
