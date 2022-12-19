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
	return ConstructorMTS(data, 1, data.CreatedAt+data.NextIntervalTime)
}
func ConstructorMTS(data *Data, sendingNumOfData int, timeToSend int64) *MessageToSend {
	return &MessageToSend{
		Data:             data,
		SendingNumOfData: sendingNumOfData,
		TimeToSend:       timeToSend,
	}
}
func ConstructorData(id int, chatID int64, messageID int) *Data {
	return &Data{
		Id:        id,
		ChatID:    chatID,
		MessageID: messageID,
	}
}
