package redis

import (
	"context"
	"fmt"

	"github.com/jdxj/study_im/utils"
)

const (
	loginSession = "LOGIN:SESSION"
)

type Session struct {
	NodeID   uint32
	ClientID int64
	UserID   uint32
}

func (s *Session) key() string {
	return fmt.Sprintf("%s:%d", loginSession, s.UserID)
}

func (s *Session) fields() []string {
	return []string{
		"NodeID", "ClientID", "UserID",
	}
}

func (s *Session) Set() error {
	m, err := utils.Struct2Map(s)
	if err != nil {
		return err
	}
	return client.HSet(context.Background(), s.key(), m).Err()
}

func (s *Session) Get() error {
	result, err := client.HMGet(ctx, s.key(), s.fields()...).Result()
	if err != nil {
		return err
	}

	for i, name := range s.fields() {
		err := utils.SetField(s, name, result[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Session) Del() error {
	return client.Del(ctx, s.key()).Err()
}
