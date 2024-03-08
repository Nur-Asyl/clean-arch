package comment

type Comment struct {
	ID        int
	UserID    int
	ArticleID int
	Text      string
}

func NewComment(userID, articleID int, text string) *Comment {
	return &Comment{
		UserID:    userID,
		ArticleID: articleID,
		Text:      text,
	}
}
