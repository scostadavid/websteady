package service

import (
	"github.com/scostadavid/tiger/internal/app/monitor/repository"
)

// entidade é usada pelo repositório que é usada pelo serviço que é usada pela main, outros serviços só se conversam via broker
type MonitorableService struct {
	repository repository.ReaderWriter
}

func NewMonitorableService(repository repository.ReaderWriter) *MonitorableService {
	return &MonitorableService{
		repository: repository,
		// broker
	}
}

// test
// func (s *MonitorableService) GetM(id int) (*entity.Monitorable, error) {
// 	monitorable, _ := s.repository.GetByID(id)
// 	return monitorable, nil
// }

// func Monitor(ch chan *MonitorResponse, service Service, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	client := &http.Client{
// 		Timeout: 5 * time.Second,
// 	}

// 	req, err := http.NewRequest("GET", service.Url, nil)

// 	if err != nil {
// 		ch <- &MonitorResponse{Service: service, Up: false}
// 		return
// 	}

// 	// browser request mimic
// 	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

// 	resp, err := client.Do(req)

// 	if err != nil {
// 		ch <- &MonitorResponse{Service: service, Up: false}
// 		return
// 	}
// 	resp.Body.Close()

// 	ch <- &MonitorResponse{Service: service, Up: resp.StatusCode == 200}
// }

// func (m *Module) Loop() {
// 	services := []Service{
// 		{Name: "Google", Url: "https://www.google.com"},
// 		{Name: "GitHub", Url: "https://www.github.com"},
// 		{Name: "Netflix", Url: "https://www.netflix.com"},
// 		{Name: "Httpbin", Url: "https://httpbin.org/status/400"},
// 	}

// 	for {
// 		var wg sync.WaitGroup

// 		ch := make(chan *MonitorResponse)

// 		for _, service := range services {
// 			wg.Add(1)
// 			go Monitor(ch, service, &wg)
// 		}

// 		go func() {
// 			wg.Wait()
// 			close(ch)
// 		}()

// 		for value := range ch {
// 			fmt.Printf("Service: %s, Status: %v\n", value.Service.Name, value.Up)
// 		}

// 		fmt.Println("Loop")
// 		time.Sleep(5 * time.Second)
// 	}

// 	fmt.Println("Monitoring complete.")
// }
