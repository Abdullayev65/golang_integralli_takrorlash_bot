package entity

import "time"

type Data struct {
	Id                    int
	ChatID                int64
	MessageID             int
	CreatedAt             int64
	NextIntervalTime      int64
	IncreasingCoefficient float32
	Active                bool
}

type MessageToSend struct {
	*Data
	SendingNumOfData int
	TimeToSend       int64
}

const day = 24 * int64(time.Hour)
const defaultInterval = day
const defaultIncreasingCoefficient = 1.8

func NewData(chatID int64, messageID int) *Data {
	return &Data{
		ChatID:                chatID,
		MessageID:             messageID,
		IncreasingCoefficient: defaultIncreasingCoefficient,
		CreatedAt:             time.Now().UnixNano(),
		NextIntervalTime:      defaultInterval,
		Active:                true,
	}
}
func NewMessageToSend(data *Data) *MessageToSend {
	return &MessageToSend{
		Data:             data,
		SendingNumOfData: 1,
		TimeToSend:       data.CreatedAt + data.NextIntervalTime,
	}
}
