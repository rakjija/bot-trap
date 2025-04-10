package models

import (
	"database/sql"
	"log"
	"time"
)

type LogEntry struct {
	IPAddress string    `json:"ip"`
	UserAgent string    `json:"user_agent"`
	Path      string    `json:"path"`
	Timestamp time.Time `json:"timestamp"`
}

var DB *sql.DB

func InsertLog(entry LogEntry) error {
	_, err := DB.Exec("INSERT INTO logs (ip_address, user_agent, path, timestamp) VALUES ($1, $2, $3, $4)",
		entry.IPAddress, entry.UserAgent, entry.Path, entry.Timestamp)
	return err
}

func FetchLogs() ([]LogEntry, error) {
	rows, err := DB.Query("SELECT ip_address, user_agent, path, timestamp FROM logs ORDER BY timestamp DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []LogEntry
	for rows.Next() {
		var entry LogEntry
		if err := rows.Scan(&entry.IPAddress, &entry.UserAgent, &entry.Path, &entry.Timestamp); err != nil {
			log.Println(err)
			continue
		}
		logs = append(logs, entry)
	}
	return logs, nil
}
