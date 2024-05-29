package main

type Config struct {
	DB_USER        string
	DB_PASSWORD    string
	DB_HOST        string
	DB_PORT        int
	EMAIL_HOST     string
	EMAIL_PORT     int
	EMAIL_USER     string
	EMAIL_PASSWORD string
	SEND_INTERVAL  int
}
