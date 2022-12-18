package repository

import (
	"database/sql"
	"fmt"
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/entity"
	_ "github.com/lib/pq"
	"log"
	"net/http"
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

func dataRecord(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM newtable")

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	snbs := make([]sandbox, 0)

	for rows.Next() {
		snb := sandbox{}
		err := rows.Scan(&snb.id, &snb.Firstname, &snb.Lastname, &snb.Age)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		snbs = append(snbs, snb)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, snb := range snbs {
		fmt.Fprintf(w, "%d %s %s %d\n", snb.id, snb.Firstname, snb.Lastname, snb.Age)
	}
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
