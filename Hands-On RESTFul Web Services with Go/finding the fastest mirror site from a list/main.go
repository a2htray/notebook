package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var mirrorList = []string{
	"http://ftp.am.debian.org/debian/",
	"http://ftp.au.debian.org/debian/",
	"http://ftp.at.debian.org/debian/",
	"http://ftp.by.debian.org/debian/",
	"http://ftp.be.debian.org/debian/",
	"http://ftp.br.debian.org/debian/",
	"http://ftp.bg.debian.org/debian/",
	"http://ftp.ca.debian.org/debian/",
	"http://ftp.cl.debian.org/debian/",
	"http://ftp2.cn.debian.org/debian/",
	"http://ftp.cn.debian.org/debian/",
	"http://ftp.hr.debian.org/debian/",
	"http://ftp.cz.debian.org/debian/",
	"http://ftp.dk.debian.org/debian/",
	"http://ftp.sv.debian.org/debian/",
	"http://ftp.ee.debian.org/debian/",
	"http://ftp.fr.debian.org/debian/",
	"http://ftp2.de.debian.org/debian/",
	"http://ftp.de.debian.org/debian/",
	"http://ftp.gr.debian.org/debian/",
	"http://ftp.hk.debian.org/debian/",
	"http://ftp.hu.debian.org/debian/",
	"http://ftp.is.debian.org/debian/",
	"http://ftp.it.debian.org/debian/",
	"http://ftp.jp.debian.org/debian/",
	"http://ftp.kr.debian.org/debian/",
	"http://ftp.lt.debian.org/debian/",
	"http://ftp.mx.debian.org/debian/",
	"http://ftp.md.debian.org/debian/",
	"http://ftp.nl.debian.org/debian/",
	"http://ftp.nc.debian.org/debian/",
	"http://ftp.nz.debian.org/debian/",
	"http://ftp.no.debian.org/debian/",
	"http://ftp.pl.debian.org/debian/",
	"http://ftp.pt.debian.org/debian/",
	"http://ftp.ro.debian.org/debian/",
	"http://ftp.ru.debian.org/debian/",
	"http://ftp.sg.debian.org/debian/",
	"http://ftp.sk.debian.org/debian/",
	"http://ftp.si.debian.org/debian/",
	"http://ftp.es.debian.org/debian/",
	"http://ftp.fi.debian.org/debian/",
	"http://ftp.se.debian.org/debian/",
	"http://ftp.ch.debian.org/debian/",
	"http://ftp.tw.debian.org/debian/",
	"http://ftp.tr.debian.org/debian/",
	"http://ftp.uk.debian.org/debian/",
	"http://ftp.us.debian.org/debian/",
}

// response 返回给前端的数据结构
type response struct {
	FastestUrl string `json:"fastestUrl"`
	Latency time.Duration `json:"latency"`
}

func findFastest(urls []string) response {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)

	for _, url := range urls {
		go func() {
			start := time.Now()
			if resp, err := http.Get(url); err == nil {
				latency := time.Now().Sub(start) / time.Millisecond
				fmt.Println(fmt.Sprintf("url: %s, response status: %s", url, resp.Status))

				urlChan <- url
				latencyChan <- latency
			}
		}()
	}

	return response{
		<-urlChan,
		<-latencyChan,
	}
}

func myFindFastest(urls []string) response {
	respChan := make(chan response)

	for _, url := range urls {
		go func() {
			start := time.Now()
			_, err := http.Get(url)
			latency := time.Now().Sub(start) / time.Millisecond

			if err == nil {
				respChan <- response{
					FastestUrl: url,
					Latency: latency,
				}
			}

		}()
	}

	return <-respChan
}

func main() {
	http.HandleFunc("/fastest-mirror", func(w http.ResponseWriter, r *http.Request) {
		response := myFindFastest(mirrorList)
		respJSON, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(respJSON)
	})

	port := ":8000"
	server := &http.Server{
		Addr: port,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println(fmt.Sprintf("Starting server on port %s", port))

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}