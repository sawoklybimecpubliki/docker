package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func answer(w http.ResponseWriter, r *http.Request) {
	answer, err := json.Marshal("Vse rabotaet, vot otvet")
	if err != nil {
		log.Fatal(err)
	}

	w.Write(answer)

	/*ans, err := http.Get("http://time_app:8089")
	if err != nil {
		log.Fatal(err)
	}
	answer, err = json.Marshal(ans)
	w.Write(answer)*/
}

func kek(w http.ResponseWriter, r *http.Request) {
	answer, err := json.Marshal("AAAAAAA chto?")
	if err != nil {
		log.Fatal(err)
	}

	w.Write(answer)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", answer)
	router.HandleFunc("GET /kek", kek)

	log.Println("server start listening on :8088")
	err := http.ListenAndServe(":80", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	log.Println("server stop")
}
