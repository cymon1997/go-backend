package repo

const (
	insertArticleQuery = `
		INSERT INTO article (title, description, content, create_by, update_by)
		VALUES (:title, :description, :content, :create_by, :update_by)
		RETURNING id`

	getArticleQuery = `
		SELECT id, title, description, content, create_time, create_by, update_time, update_by
		FROM article
		WHERE id = $1`
)
