package domain

type Message struct {
	FeedTitle        string
	FeedLink         string
	EntryTitle       string
	EntryLink        string
	EntryDescription string
	EntryPublishDate string
	EntryAuthors     []string
}
