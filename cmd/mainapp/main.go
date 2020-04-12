package main

import (
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
	log.Fatalf(mainTag, "Aborting...", http.ListenAndServe(":8000", provider.GetHandlers()))
}
