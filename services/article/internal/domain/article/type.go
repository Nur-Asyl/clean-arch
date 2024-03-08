package article

type Article struct {
	ID   int
	Name string
	Text string
}

func NewArticle(name string, text string) (*Article, error) {
	article := &Article{
		Name: name,
		Text: text,
	}

	return article, nil
}
