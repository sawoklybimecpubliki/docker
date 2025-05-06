package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func answer(w http.ResponseWriter, r *http.Request) {
	answer, err := json.Marshal("Vse rabotaet, vot otvet")
	if err != nil {
		log.Fatal(err)
	}

	w.Write(answer)

	ans, err := http.Get("http://time_app:8089")
	if err != nil {
		log.Fatal(err)
	}

	out, _ := io.ReadAll(ans.Body)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(out)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", answer)

	log.Println("server start listening on :8088")
	err := http.ListenAndServe(":80", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	log.Println("server stop")
}
