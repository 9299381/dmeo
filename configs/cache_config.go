package configs

type CacheConfig struct {
	Size string `json:"size"`
}

func (it *CacheConfig) Load() *CacheConfig {

	config := &CacheConfig{
		Size: Env("CACHE_SIZE", "1048576"),
	}
	return config
}
