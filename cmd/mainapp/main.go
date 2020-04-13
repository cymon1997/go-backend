package main

import (
	"fmt"
	"net/http"

	"github.com/cymon1997/go-backend/internal/log"
	"github.com/cymon1997/go-backend/provider"
)

const mainTag = "Main"

func main() {
	consumers := provider.GetArticleConsumers()
	for _, c := range consumers {
		c.Consume()
	}
	cfg := provider.GetAppConfig()
	log.Fatalf(mainTag, "Aborting...",
		http.ListenAndServe(fmt.Sprintf(":%d", cfg.AppConfig.Port),
			provider.GetHandlers()))
}
