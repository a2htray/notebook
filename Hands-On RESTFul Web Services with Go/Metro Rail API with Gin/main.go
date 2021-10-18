package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.GET("/v1/stations/:stationID", GetStation)
	r.POST("/v1/stations", CreateStation)
	r.DELETE("/v1/stations/:stationID", RemoveStation)
}

func GetStation(c *gin.Context) {
	var train StationResponse
	stationID := c.Param("stationID")

	row := globalDB.Get().QueryRow("select id, name, cast(opening_time as char), cast(closing_time as char) from station where id=?", stationID)
	if err := row.Scan(&train.ID, &train.Name, &train.OpeningTime, &train.ClosingTime); err != nil {
		log.Printf("sqlite3: fetch a station record error - %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": train,
	})
}

func CreateStation(c *gin.Context) {
	var station StationResponse
	if err := c.BindJSON(&station); err != nil {
		log.Printf("request: parse request body error - %s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		if stmt, err := globalDB.Get().Prepare("insert into station (name, opening_time, closing_time) values (?, ?, ?)"); err != nil {
			log.Printf("sqlite3: prepare the create station statement error - %s\n", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			if result, err := stmt.Exec(station.Name, station.OpeningTime, station.ClosingTime); err != nil {
				log.Printf("sqlite3: execute the create station statement error - %s\n", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			} else {
				if stationID, err := result.LastInsertId(); err != nil {
					log.Printf("sqlite3: get the last insert ID error - %s\n", err.Error())
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
					})
					return
				} else {
					station.ID = int(stationID)
					c.JSON(http.StatusOK, gin.H{
						"result": station,
					})
				}
			}
		}
	}
}

func RemoveStation(c *gin.Context) {
	stationID := c.Param("station_id")

	if stmt, err := globalDB.Get().Prepare("delete from station where id=?"); err != nil {
		log.Printf("sqlite3: prepare the delete statement error - %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		if _, err := stmt.Exec(stationID); err != nil {
			log.Printf("sqlite3: execute the delete statement error - %s\n", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": true,
			})
		}
	}
}
