package repository

import (
	"database/sql"
	"fmt"
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/entity"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB
var err error

type sandbox struct {
	id        int
	Firstname string
	Lastname  string
	Age       int
}

func init() {
	connStr := "postgres://postgres:1@localhost/postgres?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Now we are connected to POSTGRESQL DATABASE.")
}

func SaveData(data *entity.Data) {
	query := fmt.Sprintf("INSERT INTO e_data(chat_id, message_id, created_at, next_interval_time, increasing_coefficient, active) VALUES(%d, %d, %d , %d,  %.3f, %t)",
		data.ChatID, data.MessageID, data.CreatedAt,
		data.NextIntervalTime, data.IncreasingCoefficient, data.Active)
	_, err := db.Exec(query)
	if err != nil {
		log.Println(err)
		return
	}
	setDataID(data)
}

func SaveMessageToSend(messageToSend *entity.MessageToSend) {
	query := fmt.Sprintf("INSERT INTO message_to_send(data_id, sending_num_of_data , time_to_send) VALUES (%d, %d, %d )",
		messageToSend.Data.Id, messageToSend.SendingNumOfData, messageToSend.TimeToSend)
	_, err := db.Exec(query)
	if err != nil {
		log.Println(err)
		return
	}
}

func GetSliceOfMTS(from, to int64) []entity.MessageToSend {
	query := fmt.Sprintf(
		"SELECT time_to_send, data_id, chat_id, message_id FROM message_to_send AS m JOIN e_data AS d on d.id = m.data_id AND d.active  WHERE m.time_to_send BETWEEN %d AND %d ", from, to)
	rows, err := db.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	sliceMTS := make([]entity.MessageToSend, 0, 20)

	for rows.Next() {
		mts := scanAndMap(rows)
		if err != nil {
			log.Println(err)
			return nil
		}
		sliceMTS = append(sliceMTS, *mts)
	}
	return sliceMTS
}
func UpdateNextIntervalTime(data *entity.Data) {
	query := "UPDATE e_data SET next_interval_time = next_interval_time * increasing_coefficient WHERE id = " + string(data.Id)
	db.Exec(query)
}

// private
func scanAndMap(rows *sql.Rows) *entity.MessageToSend {
	var timeToSend, dataID, chatID, messageID interface{}
	rows.Scan(&timeToSend, &dataID, &chatID, &messageID)
	data := entity.ConstructorData(int(parseInt(dataID)), parseInt(chatID), int(parseInt(messageID)))
	return entity.ConstructorMTS(data, -1, parseInt(timeToSend))
}
func setDataID(data *entity.Data) error {
	query := fmt.Sprintf("SELECT id FROM e_data WHERE chat_id = %d AND message_id = %d",
		data.ChatID, data.MessageID)
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return err
	}
	rows.Next()
	return rows.Scan(&data.Id)
}

func parseInt(i interface{}) (n int64) {
	switch i.(type) {
	case int:
		n = int64(i.(int))
	case int8:
		n = int64(i.(int8))
	case int16:
		n = int64(i.(int16))
	case int32:
		n = int64(i.(int32))
	case int64:
		n = i.(int64)
	}
	return n
}
