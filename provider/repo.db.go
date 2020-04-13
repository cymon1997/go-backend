package provider

import (
	"sync"

	"github.com/cymon1997/go-backend/internal/base/repo"
	article "github.com/cymon1997/go-backend/module/article/repo"
)

var (
	baseDBRepo     repo.DBRepo
	syncBaseDBRepo sync.Once

	articleDBRepo     article.ArticleDBRepo
	syncArticleDBRepo sync.Once
)

func GetBaseDBRepo() repo.DBRepo {
	syncBaseDBRepo.Do(func() {
		baseDBRepo = repo.NewBaseDBRepo(GetDBClient())
	})
	return baseDBRepo
}

func GetArticleDBRepo() article.ArticleDBRepo {
	syncArticleDBRepo.Do(func() {
		articleDBRepo = article.NewArticleDBRepo(GetBaseDBRepo())
	})
	return articleDBRepo
}
