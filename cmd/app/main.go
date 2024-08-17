package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"

	_ "github.com/mattn/go-sqlite3"
	monitorEntity "github.com/scostadavid/websteady/internal/app/monitorable/entity"
	monitorRepo "github.com/scostadavid/websteady/internal/app/monitorable/repository"
	monitorServ "github.com/scostadavid/websteady/internal/app/monitorable/service"
)

type MonitorableResponse struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	URL  string `db:"url"`
	Up   bool   `db:"up"`
}

func checkStatus(m *monitorEntity.Monitorable, ch chan<- *MonitorableResponse) {
	fmt.Println("Check for", m.URL)
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", m.URL, nil)
	if err != nil {
		ch <- &MonitorableResponse{ID: m.ID, Name: m.Name, URL: m.URL, Up: false}
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		ch <- &MonitorableResponse{ID: m.ID, Name: m.Name, URL: m.URL, Up: false}
		return
	}
	defer resp.Body.Close()

	ch <- &MonitorableResponse{ID: m.ID, Name: m.Name, URL: m.URL, Up: resp.StatusCode == 200}
}

func main() {
	db, err := sql.Open("sqlite3", "./databse.db")
	if err != nil {
		log.Fatal("err", err)
	}
	defer db.Close()

	monitorableRepository, err := monitorRepo.NewMonitorableRepository(db)
	if err != nil {
		log.Fatal("Erro ao criar o repositório:", err)
	}
	defer monitorableRepository.Close()

	monitorableService := monitorServ.NewMonitorableService(monitorableRepository)

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS monitorables (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100),
			url VARCHAR(100)
		)
	`)

	if err != nil {
		log.Fatal(err)
	}

	site := func(w http.ResponseWriter, r *http.Request) {

		tmpl, err := template.ParseFiles("./cmd/app/templates/site.html")
		if err != nil {
			http.Error(w, "Unable to parse template", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Unable to execute template", http.StatusInternalServerError)
			return
		}
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		res, err := monitorableService.GetAllMonitorables()

		if err != nil {
			http.Error(w, "Unable to get monitorables", http.StatusInternalServerError)
			return
		}

		ch := make(chan *MonitorableResponse, len(res))
		defer close(ch)

		for _, m := range res {
			checkStatus(m, ch) // vamos remover as routines por enquanto e depois nós paralelizamos
		}

		var monitorableResponses []*MonitorableResponse
		for i := 0; i < len(res); i++ {
			monitorableResponses = append(monitorableResponses, <-ch)
		}

		// isso não ta ordenando direito de jeito nenhum
		sort.Slice(monitorableResponses, func(i, j int) bool {
			return monitorableResponses[i].ID < monitorableResponses[j].ID
		})

		monitorables := map[string][]*MonitorableResponse{
			"Monitorables": monitorableResponses,
		}

		tmpl, err := template.ParseFiles("./cmd/app/templates/index.html")
		if err != nil {
			http.Error(w, "Unable to parse template", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, monitorables)
		if err != nil {
			http.Error(w, "Unable to execute template", http.StatusInternalServerError)
			return
		}
	}

	handler2 := func(w http.ResponseWriter, r *http.Request) {

		name := r.PostFormValue("name")
		url := r.PostFormValue("url")

		fmt.Println(name, " ", url)
		// res, err := monitorableService.GetAllMonitorables()

		// htmlStr := fmt.Sprintf(`
		// 		<div class='bg-white rounded-lg shadow p-4 flex flex-row justify-between'>
		// 				<div>
		// 						<div class='flex flex-row items-center gap-2'>
		// 								<div class='w-4 h-4 rounded-full {{if .Up}}bg-green-500{{else}}bg-red-500{{end}}'></div>
		// 								<div class='text-gray-700'>%s</div>
		// 						</div>
		// 						<div class='text-gray-500'>
		// 								<a href='%s' target='_blank' rel='noreferrer noopener'>%s</a>
		// 						</div>
		// 				</div>
		// 		</div>
		// `, name, url, url)
		// tmpl, _ := template.New("t").Parse(htmlStr)
		// tmpl.Execute(w, nil)

		// res, err := monitorableService.GetAllMonitorables()

		// if err != nil {
		// 	http.Error(w, "Unable to get monitorables", http.StatusInternalServerError)
		// 	return
		// }

		// ch := make(chan *MonitorableResponse, len(res))
		// defer close(ch)

		// for _, m := range res {
		// 	go checkStatus(m, ch)
		// }

		// var monitorableResponses []*MonitorableResponse
		// for i := 0; i < len(res); i++ {
		// 	monitorableResponses = append(monitorableResponses, <-ch)
		// }

		// // isso não ta ordenando direito de jeito nenhum
		// sort.Slice(monitorableResponses, func(i, j int) bool {
		// 	return monitorableResponses[i].ID < monitorableResponses[j].ID
		// })

		// monitorables := map[string][]*MonitorableResponse{
		// 	"Monitorables": monitorableResponses,
		// }

		// tmpl, err := template.ParseFiles("./cmd/app/templates/index.html")
		// if err != nil {
		// 	http.Error(w, "Unable to parse template", http.StatusInternalServerError)
		// 	return
		// }

		// err = tmpl.Execute(w, monitorables)
		// if err != nil {
		// 	http.Error(w, "Unable to execute template", http.StatusInternalServerError)
		// 	return
		// }
	}

	http.HandleFunc("/", site)
	http.HandleFunc("/app", handler)
	http.HandleFunc("/add-monitorable", handler2)
	fmt.Println("Server")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// monitorableRepository := monitorRepo.NewMonitorableRepository() // db
	// monitorableService := monitorServ.NewMonitorableService(monitorableRepository)

	// até aqui acho que é isso mesmo, agr de resto tenho que usar o repositório e tals

	// isso seria um see mas não sei se faz sentido
	// por default o loop de request vai ser de 60 segundos mas acho que vou deixar editável

	// fmt.Println("websteady")
	// init api handler
	// setup routes
}
