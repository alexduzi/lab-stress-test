package service

type RequestsDto struct {
	Url                string
	AmountRequests     int
	ConcurrentRequests int
}

func NewRequestsDto(url string, requests, concurrent int) RequestsDto {
	return RequestsDto{
		Url:                url,
		AmountRequests:     requests,
		ConcurrentRequests: concurrent,
	}
}

type RequestsService struct {
}

func NewRequestsService() *RequestsService {
	return &RequestsService{}
}

func (r *RequestsService) Call(req RequestsDto) error {
	return nil
}
