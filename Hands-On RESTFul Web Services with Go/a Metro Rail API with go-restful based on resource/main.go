package main

import (
	"database/sql"
	"encoding/json"
	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"sync"
	"time"
)

const sqlCreateTrain = `
create table if not exists train (
  id integer primary key autoincrement,
  driver_name varchar(64) not null,
  operating_status boolean
);`

type TrainResponse struct {
	ID              int    `json:"id"`
	DriverName      string `json:"driverName"`
	OperatingStatus bool   `json:"operatingStatus"`
}

func (t *TrainResponse) Register(ctn *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/v1/train").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{train_id}").To(t.GetTrain))
	ws.Route(ws.POST("").To(t.CreateTrain))
	ws.Route(ws.DELETE("/{train_id}").To(t.RemoveTrain))

	ctn.Add(ws)
}

func (t *TrainResponse) GetTrain(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("train_id")
	err := globalDB.Get().QueryRow("select id, driver_name, operating_status from train where id=?", id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Printf("sqlite3: fetch train record via [id=%s] error - %s\n", id, err.Error())
		resp.AddHeader("Content-Type", "text/plain")
		resp.WriteHeader(http.StatusNotFound)
		_, _ = resp.Write([]byte("train could not be found"))
	} else {
		_ = resp.WriteEntity(t)
	}
}

func (t *TrainResponse) CreateTrain(req *restful.Request, resp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	response := TrainResponse{}

	if err := decoder.Decode(&response); err != nil {
		log.Printf("decoder: decode request body error - %s\n", err.Error())
		resp.WriteHeader(http.StatusBadRequest)
		resp.AddHeader("Content-Type", "text/plain")
		_, _ = resp.Write([]byte("error: decode request body"))
	} else {
		if stmt, err := globalDB.Get().Prepare("insert into train (driver_name, operating_status) values (?, ?);"); err != nil {
			log.Printf("sqlite3: failed to prepare a create statement\n")
			resp.WriteHeader(http.StatusInternalServerError)
			resp.AddHeader("Content-Type", "text/plain")
			_, _ = resp.Write([]byte("sqlite3: encounter a server-side error"))
		} else {
			if res, err := stmt.Exec(response.DriverName, response.OperatingStatus); err != nil {
				log.Printf("sqlite3: failed to execute a create statement\n")
				resp.WriteHeader(http.StatusInternalServerError)
				resp.AddHeader("Content-Type", "text/plain")
				_, _ = resp.Write([]byte("sqlite3: encounter a server-side error"))
			} else {
				if id, err := res.LastInsertId(); err != nil {
					log.Printf("sqlite3: failed to get the last insert id\n")
					resp.WriteHeader(http.StatusInternalServerError)
					resp.AddHeader("Content-Type", "text/plain")
					_, _ = resp.Write([]byte("sqlite3: encounter a server-side error"))
				} else {
					response.ID = int(id)
					resp.AddHeader("Content-Type", "application/json")
					_ = resp.WriteEntity(&response)
				}
			}
		}
	}
}

func (t *TrainResponse) RemoveTrain(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("train_id")
	if stmt, err := globalDB.Get().Prepare("delete from train where id=?"); err != nil {
		log.Printf("sqlite3: failed to prepare a delete statement - %s\n", err.Error())
		resp.WriteHeader(http.StatusInternalServerError)
		resp.AddHeader("Content-Type", "text/plain")
		_, _ = resp.Write([]byte("error: decode request body"))
	} else {
		if _, err := stmt.Exec(id); err != nil {
			log.Printf("sqlite3: failed to execute the delete statement\n")
			resp.WriteHeader(http.StatusInternalServerError)
			resp.AddHeader("Content-Type", "text/plain")
			_, _ = resp.Write([]byte("sqlite3: encounter a server-side error"))
		} else {
			resp.WriteHeader(http.StatusOK)
		}
	}

}

const sqlCreateStation = `
create table if not exists station (
  id integer primary key autoincrement, 
  name varchar(64) not null,
  opening_time time null,
  closing_time time null
);`

type StationResponse struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

const sqlCreateSchedule = `
create table if not exists schedule (
  id integer primary key autoincrement,
  train_id int,
  station_id int,
  arrival_time time,
  foreign key (train_id) references train(id),
  foreign key (station_id) references station(id)
);
`

type ScheduleResponse struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time
}

type DB struct {
	*sql.DB
	once *sync.Once
}

func (d *DB) Get() *sql.DB {
	d.once.Do(func() {
		db, err := sql.Open("sqlite3", "./db/metro.db")
		if err != nil {
			log.Fatalf("sqlite3: connect to the server error - %s\n", err.Error())
		}

		d.DB = db
	})

	return d.DB
}

var globalDB = &DB{
	once: new(sync.Once),
}

func initialize(db *sql.DB) {
	var stmt *sql.Stmt
	var err error
	if stmt, err = db.Prepare(sqlCreateTrain); err != nil {
		log.Fatalf("sqlite3: initialize train table error - %s\n", err.Error())
	} else {
		if _, err = stmt.Exec(); err != nil {
			log.Fatalf("sqlite3: initialize train table error - %s\n", err.Error())
		}
	}

	if stmt, err = db.Prepare(sqlCreateStation); err != nil {
		log.Fatalf("sqlite3: initialize station table error - %s\n", err.Error())
	} else {
		if _, err = stmt.Exec(); err != nil {
			log.Fatalf("sqlite3: initialize station table error - %s\n", err.Error())
		}
	}

	if stmt, err = db.Prepare(sqlCreateSchedule); err != nil {
		log.Fatalf("sqlite3: initialize schedule table error - %s\n", err.Error())
	} else {
		if _, err = stmt.Exec(); err != nil {
			log.Fatalf("sqlite3: initialize schedule table error - %s\n", err.Error())
		}
	}
}

func main() {
	db := globalDB.Get()
	initialize(db)

	wsContainer := &restful.Container{
		ServeMux: http.DefaultServeMux,
	}
	wsContainer.Router(restful.CurlyRouter{})

	t := TrainResponse{}
	t.Register(wsContainer)

	server := &http.Server{
		Addr:    ":8000",
		Handler: wsContainer,
	}
	log.Fatal(server.ListenAndServe())
}
