package models

type Comment struct {
	ID      int64
	Entry   int64
	User    string
	Email   string
	URL     string
	Title   string
	Content string
	CTime   string
	ATime   string
}

type Comments struct {
}

func (self *Comment) Save() (err error) {

}

func (self *Comments) Filter(offset, count int, entry int64) (comments []Comment, err error) {

}
