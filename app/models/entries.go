package models

import "strings"

type Entry struct {
	ID      int64
	Name    string
	Abbrs   []string
	Caption string
	CTime   string
	ATime   string
}

type Entries struct {
}

const (
	SQL_UPDATE_ENTRY   = `UPDATE entries SET name=?,abbrs=?,caption=?,atime = datetime('now','localtime') WHERE id = ?;`
	SQL_INSERT_ENTRY   = `INSERT INTO entries (name,abbrs,caption,ctime,atime) VALUES (?,?,?,datetime('now','localtime'),datetime('now','localtime'));`
	SQL_SELECT_ENTRIES = `SELECT id,name,ctime,atime,abbrs,caption FROM entries;`
)

func (self *Entry) Save() (err error) {

	// update
	if self.ID != 0 {

		_, err = db.Exec(SQL_UPDATE_ENTRY, self.Name, self.Abbrs, self.Caption, self.ID)
		return
	}

	// insert
	result, err := db.Exec(SQL_INSERT_ENTRY, self.Name, self.Abbrs, self.Caption)

	if err != nil {
		return
	}

	self.ID, err = result.LastInsertId()

	return
}

func (self *Entries) Search(offset, count int, words string) (entries []Entry, err error) {

	rows, err := db.Query(SQL_SELECT_ENTRIES)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {

		var entry Entry

		err = rows.Scan(&entry.ID, &entry.Name, &entry.CTime, &entry.ATime, &entry.Abbrs, &entry.Caption)

		if err != nil {
			return
		}

		if !strings.Contains(entry.Name+entry.Abbrs+entry.Caption, words) {
			continue
		}

		offset = offset - 1

		if offset > 0 {
			continue
		}

		entries = append(entries, entry)

		count = count - 1

		if count <= 0 {
			return
		}
	}

	err = rows.Err()

	return
}
