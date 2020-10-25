package db

import (
	"log"
	"time"

	"github.com/mkhstar/inmemdb/platform"
)

func keyCleaner() {
	for {
		for key := range inMemDB.keysToExpire {
			if expired := isExpired(key); expired {
				log.Printf("Key %s deleted from memory%s", key, platform.LineBreak)
			}

		}
		time.Sleep(5 * time.Second)
	}
}
