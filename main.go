package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

type SpeedTestResult struct {
	DownloadSpeed float64
	UploadSpeed   float64
	Ping          float64
	TestTime      time.Time
}

var result SpeedTestResult

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/test", handleTest)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	client := speedtest.New()

	serverList, err := client.FetchServers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	targets, err := serverList.FindServer([]int{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, s := range targets {
		err := s.PingTest(func(latency time.Duration) {})
		if err != nil {
			continue
		}
		err = s.DownloadTest()
		if err != nil {
			continue
		}
		err = s.UploadTest()
		if err != nil {
			continue
		}

		result = SpeedTestResult{
			DownloadSpeed: s.DLSpeed.Mbps(), // Convert to Mbps as float64
			UploadSpeed:   s.ULSpeed.Mbps(), // Convert to Mbps as float64
			Ping:          float64(s.Latency.Milliseconds()),
			TestTime:      time.Now(),
		}
		break
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
