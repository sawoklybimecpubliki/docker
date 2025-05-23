package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func ans(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	year, month, day := now.Date()
	str := strconv.Itoa(year) + " " + month.String() + " " + strconv.Itoa(day) + " " + strconv.Itoa(now.Hour()) + ":" +
		strconv.Itoa(now.Minute())
	answer, err := json.Marshal(str)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(answer)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", ans)
	log.Println("server start listening on :8089")
	err := http.ListenAndServe(":8089", router)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("server stop")

}
