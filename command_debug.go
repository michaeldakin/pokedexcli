package main

import (
	"fmt"

	"github.com/michaeldakin/pokedevcli/internal/pokecache"
)

func cacheDebug(cfg *config) error {
    cache := &pokecache.Cache{}
    fmt.Printf("cache items: %d\n", cache.Len())
    cache.GetAll()

    return nil
}
