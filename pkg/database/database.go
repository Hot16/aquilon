package database

import (
	"aquilon/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
	"time"
)

type Database struct {
	conn *pgx.Conn
}

func NewDatabase() (*Database, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	fmt.Println(connString)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	return &Database{
		conn: conn,
	}, nil
}

func (db *Database) Close() {
	db.conn.Close(context.Background())
}

func (db *Database) SaveBatch(clients []models.Clients, batchSize int) error {
	values := []interface{}{}
	query := "INSERT INTO clients (id, dt, client_id, type, submit_id, referer, os, lead_source, creative_name, country) VALUES "

	// Добавляем параметры для каждой записи
	for i, client := range clients {
		values = append(values, client.Id, time.Unix(int64(client.Dt), 0), client.ClientId, client.Type, client.SubmitId, client.Referer, client.Os, client.LeadSource, client.CreativeName, client.Country)
		query += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d), ", i*10+1, i*10+2, i*10+3, i*10+4, i*10+5, i*10+6, i*10+7, i*10+8, i*10+9, i*10+10)
	}

	query = query[:len(query)-2]

	query += " ON CONFLICT (client_id) DO UPDATE SET dt = excluded.dt"

	// Разбиваем батч на несколько частей, если необходимо
	for i := 0; i < len(values); i += batchSize * 10 {
		j := i + batchSize*10
		if j > len(values) {
			j = len(values)
		}

		_, err := db.conn.Exec(context.Background(), query, values[i:j]...)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *Database) SaveClients(clients []models.Clients, batchSize int) error {
	for i := 0; i < len(clients); i += batchSize {
		j := i + batchSize
		if j > len(clients) {
			j = len(clients)
		}

		err := db.SaveBatch(clients[i:j], batchSize)
		if err != nil {
			return err
		}
	}

	return nil
}
