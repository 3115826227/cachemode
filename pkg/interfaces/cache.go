package interfaces

type Cache interface {
	Conn() (err error)
	Write(key CacheKey, value []byte) (err error)
	Read(key CacheKey) (value []byte, err error)
	Close() (err error)
}

type CacheKey interface {
	GetCacheKey()
}
