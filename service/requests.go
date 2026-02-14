package service

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/alexduzi/labstresstest/dto"
)

type RequestsService struct {
	client *http.Client
}

func NewRequestsService() *RequestsService {
	return &RequestsService{
		client: http.DefaultClient,
	}
}

func (r *RequestsService) ProcessRequests(request *dto.RequestsDto) {
	start := time.Now()
	resChan := make(chan dto.ResponseDto, request.AmountRequests)
	wg := sync.WaitGroup{}

	reqPerThread := request.AmountRequests / request.ConcurrentRequests
	remainder := request.AmountRequests % request.ConcurrentRequests

	wg.Add(request.ConcurrentRequests)

	for i := range request.ConcurrentRequests {
		count := reqPerThread
		if i < remainder {
			count++
		}
		go func(url string, count int) {
			for range count {
				r.call(url, resChan)
			}
			wg.Done()
		}(request.Url, count)
	}

	wg.Wait()
	close(resChan)

	end := time.Since(start)

	report := r.prepareReport(resChan)

	report.TimeSpent = end

	r.reportRequests(report)
}

func (r *RequestsService) call(url string, resChan chan dto.ResponseDto) error {
	response := dto.NewResponseDto(url)
	defer func() {
		resChan <- response
	}()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		response.ErrorRequest = true
		response.ErrorMessage = err.Error()
		return err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		response.ErrorRequest = true
		response.ErrorMessage = err.Error()
		return err
	}
	defer resp.Body.Close()

	response.Elapsed = time.Since(response.Start).Milliseconds()
	response.Method = "GET"
	response.Status = resp.StatusCode

	return nil
}

func (r *RequestsService) prepareReport(resChan chan dto.ResponseDto) dto.Report {
	var report dto.Report
	report.AmountHttpStatus = make(map[int]int)
	for response := range resChan {
		report.AmountRequests++
		if response.ErrorRequest {
			report.AmountOfErrors++
			continue
		}
		if value, ok := report.AmountHttpStatus[response.Status]; ok {
			report.AmountHttpStatus[response.Status] = value + 1
		} else {
			report.AmountHttpStatus[response.Status] = 1
		}
	}
	return report
}

func (r *RequestsService) reportRequests(report dto.Report) {
	fmt.Printf("Tempo total gasto na execução: %d ms\n", report.TimeSpent.Milliseconds())
	fmt.Printf("Quantidade total de requests realizados: %d\n", report.AmountRequests)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", report.AmountHttpStatus[http.StatusOK])
	for key, value := range report.AmountHttpStatus {
		if key != http.StatusOK {
			fmt.Printf("Quantidade de requests com status HTTP %d: %d\n", key, value)
		}
	}
	fmt.Printf("Quantidade total de requests que deram erro: %d\n", report.AmountOfErrors)
}
