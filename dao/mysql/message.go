package mysql

import (
	"fmt"
	"time"
)

type MessageSend struct {
	ID       int64
	FromID   uint32
	ToID     uint32
	Seq      uint32
	Content  []byte
	SendTime time.Time
	SendType int
}

// TableName 如果考虑分表, 使用 from 字段
func (ms *MessageSend) TableName() string {
	return "message_send"
}

func (ms *MessageSend) Insert() error {
	query := fmt.Sprintf(`insert into %s (from_id,to_id,seq,content,send_time,send_type) values (?,?,?,?,?,?)`,
		ms.TableName())
	result, err := db.Exec(query, ms.FromID, ms.ToID, ms.Seq, ms.Content, ms.SendTime, ms.SendType)
	if err != nil {
		return err
	}

	ms.ID, err = result.LastInsertId()
	return err
}

type MessageReceive struct {
	ID        int64
	FromID    uint32
	ToID      uint32
	MessageID int64
	Flag      int
}

func (mr *MessageReceive) TableName() string {
	return "message_receive"
}

func (mr *MessageReceive) Insert() error {
	query := fmt.Sprintf(`insert into %s (from_id,to_id,message_id,flag) values (?,?,?,?)`, mr.TableName())
	result, err := db.Exec(query, mr.FromID, mr.ToID, mr.MessageID, 1)
	if err != nil {
		return err
	}

	mr.ID, err = result.LastInsertId()
	return err
}

func (mr *MessageReceive) SetRead() error {
	query := fmt.Sprintf(`update %s set flag=? where to_id=? and message_id=?`, mr.TableName())
	_, err := db.Exec(query, 2, mr.ToID, mr.MessageID)
	return err
}
