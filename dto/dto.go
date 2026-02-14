package dto

import (
	"fmt"
	"time"
)

type RequestsDto struct {
	Url                string
	AmountRequests     int
	ConcurrentRequests int
}

func NewRequestsDto(url string, requests, concurrency int) (*RequestsDto, error) {
	if concurrency <= 0 {
		return nil, fmt.Errorf("amount of concurrency cannot be less than zero or equal to zero.")
	}

	if requests <= 0 {
		return nil, fmt.Errorf("amount of requests cannot be less than zero or equal to zero.")
	}

	if concurrency > requests {
		return nil, fmt.Errorf("amount of concurrency is higher than the amount of requests.")
	}

	return &RequestsDto{
		Url:                url,
		AmountRequests:     requests,
		ConcurrentRequests: concurrency,
	}, nil
}

type ResponseDto struct {
	Start        time.Time
	Url          string
	Method       string
	Elapsed      int64
	Status       int
	ErrorRequest bool
	ErrorMessage string
}

func NewResponseDto(url string) ResponseDto {
	return ResponseDto{
		Start: time.Now(),
		Url:   url,
	}
}

type Report struct {
	TimeSpent        time.Duration
	AmountRequests   int
	AmountHttpStatus map[int]int
	AmountOfErrors   int
}
