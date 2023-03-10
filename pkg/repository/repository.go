package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/dotEnv"
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/entity"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

var db *sql.DB
var err error

func init() {
	connStr := dotEnv.EnvMap["DB_Conn"]

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Now we are connected to POSTGRESQL DATABASE.")
	b, e := os.ReadFile("database.sql")
	if e == nil {
		db.Exec(string(b))
	}
}

func SaveData(data *entity.Data) {
	query := fmt.Sprintf("INSERT INTO e_data(chat_id, message_id, created_at, next_interval_time, increasing_coefficient, active) VALUES(%d, %d, %d , %d,  %.3f, %t) RETURNING id",
		data.ChatID, data.MessageID, data.CreatedAt,
		data.NextIntervalTime, data.IncreasingCoefficient, data.Active)
	row, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	row.Next()
	row.Scan(&data.Id)
}

func SaveMessageToSend(messageToSend *entity.MessageToSend) {
	query := fmt.Sprintf("INSERT INTO message_to_send(data_id, sending_num_of_data , time_to_send) VALUES ($1, $2, $3 )")
	_, err := db.Exec(query,
		messageToSend.Data.Id, messageToSend.SendingNumOfData, messageToSend.TimeToSend)
	if err != nil {
		log.Println(err)
		return
	}
}

func GetSliceOfMTS(from, to int64) []entity.MessageToSend {
	query := fmt.Sprintf(
		"SELECT time_to_send, data_id, chat_id, message_id, next_interval_time, sending_num_of_data FROM message_to_send AS m JOIN e_data AS d on d.id = m.data_id AND d.active  WHERE m.time_to_send BETWEEN %d AND %d ", from, to)
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
	query := "UPDATE e_data SET next_interval_time = next_interval_time * increasing_coefficient WHERE id = $1"
	db.Exec(query, data.Id)
}
func GetIdOfData(messageId int, chatID int64) (int, error) {
	query := fmt.Sprintf("SELECT id FROM e_data WHERE chat_id = %d AND message_id = %d;", chatID, messageId)
	rows, err := db.Query(query)
	if err != nil {
		return 0, err
	}
	if !rows.Next() {
		return 0, errors.New("NOT FOUNDED")
	}
	id := 0
	rows.Scan(&id)
	return id, nil
}
func UpdateK(dataID int, newK float64) error {
	query := "UPDATE e_data SET increasing_coefficient = $1 WHERE id = " + strconv.Itoa(dataID)
	_, err = db.Exec(query, newK)
	return err
}

// private
func scanAndMap(rows *sql.Rows) *entity.MessageToSend {
	var timeToSend, dataID, chatID, messageID, nextIntervalTime, sendingNum interface{}
	rows.Scan(&timeToSend, &dataID, &chatID, &messageID, &nextIntervalTime, &sendingNum)
	data := entity.ConstructorData(int(parseInt(dataID)), parseInt(chatID), int(parseInt(messageID)), parseInt(nextIntervalTime))
	return entity.ConstructorMTS(data, int(parseInt(sendingNum)), parseInt(timeToSend))
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
