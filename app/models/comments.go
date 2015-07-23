package models

import "time"

type Comment struct {
	ID      int64
	Entry   int64
	User    string
	Email   string
	URL     string
	Title   string
	Content string
	CTime   time.Time
	ATime   time.Time
}

type Comments struct {
}

const (
	SQL_UPDATE_COMMENT  = `UPDATE comments SET eid=?,user=?,email=?,url=?,title=?,content=?,atime = datetime('now','localtime') WHERE id = ?;`
	SQL_INSERT_COMMENT  = `INSERT INTO entries (eid,user,email,url,title,content,ctime,atime) VALUES (?,?,?,?,?,?,datetime('now','localtime'),datetime('now','localtime'));`
	SQL_SELECT_COMMENTS = `SELECT id,eid,ctime,atime,user,email,url,title,content FROM comments WHERE eid=? ORDER BY id;`
)

func (self *Comment) Save() (err error) {

	// update
	if self.ID != 0 {

		_, err = db.Exec(SQL_UPDATE_COMMENT, self.Entry, self.User, self.Email, self.URL, self.Title, self.Content, self.ID)
		return
	}

	// insert
	result, err := db.Exec(SQL_INSERT_COMMENT, self.Entry, self.User, self.Email, self.URL, self.Title, self.Content)

	if err != nil {
		return
	}

	self.ID, err = result.LastInsertId()

	return
}

func (self *Comments) Filter(offset, count int, entry int64) (comments []Comment, err error) {

	rows, err := db.Query(SQL_SELECT_COMMENTS, entry)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {

		offset = offset - 1

		if offset > 0 {
			continue
		}

		var comment Comment

		err = rows.Scan(&comment.ID, &comment.Entry, &comment.CTime, &comment.ATime, &comment.User, &comment.Email, &comment.URL, &comment.Title, &comment.Content)

		if err != nil {
			return
		}

		comments = append(comments, comment)

		count = count - 1

		if count <= 0 {
			return
		}
	}

	err = rows.Err()

	return
}
