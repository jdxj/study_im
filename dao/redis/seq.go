package redis

import "fmt"

// 用于保证群聊天顺序

const (
	seqGroupChat = "SEQ:GROUP"
	seqGroupSlot = 1 << 14
)

func seqGroupChatKey(groupID uint32) string {
	key := fmt.Sprintf("%s:%d", seqGroupChat, groupID%seqGroupSlot)
	return key
}

func NextGroupChatSeq(groupID uint32) (int64, error) {
	key := seqGroupChatKey(groupID)
	field := fmt.Sprintf("%d", groupID)
	return client.HIncrBy(ctx, key, field, 1).Result()
}
