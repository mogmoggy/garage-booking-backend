// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"time"
)

type Booking struct {
	ID                 int64          `json:"id"`
	UserID             int64          `json:"user_id"`
	CarID              int64          `json:"car_id"`
	BookingDate        time.Time      `json:"booking_date"`
	ProblemDescription sql.NullString `json:"problem_description"`
	CreatedAt          time.Time      `json:"created_at"`
}

type Car struct {
	ID              int64     `json:"id"`
	UserID          int64     `json:"user_id"`
	CarRegistration string    `json:"car_registration"`
	Make            string    `json:"make"`
	Model           string    `json:"model"`
	YearManufacture int32     `json:"year_manufacture"`
	CreatedAt       time.Time `json:"created_at"`
}

type User struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}
