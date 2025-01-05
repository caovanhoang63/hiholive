package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Config struct {
	Domain          string `json:"domain"`
	Ratio           int    `json:"ratio"`
	CumulativeRatio int    `json:"-"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")

	var configs []Config
	fileData, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	err = json.Unmarshal(fileData, &configs)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	totalRatio := 0
	for i := range configs {
		totalRatio += configs[i].Ratio
		configs[i].CumulativeRatio = totalRatio
	}
	if totalRatio != 100 {
		log.Fatalln(fmt.Errorf("total ratio does not match 100%%"))
	}
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		ran := rand.Intn(100)
		for i := range configs {
			if ran < configs[i].CumulativeRatio {
				http.Redirect(w, r, configs[i].Domain+r.RequestURI, http.StatusFound)
				break
			}
		}
	})

	http.HandleFunc("POST /config", func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("api-key")
		if key != apiKey {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		r.Body = http.MaxBytesReader(w, r.Body, 1024*1024)
		var configBody []Config
		err = json.NewDecoder(r.Body).Decode(&configBody)
		if err != nil {
			log.Printf("Error decoding JSON: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		totalRatio = 0
		for i := range configs {
			totalRatio += configBody[i].Ratio
			configBody[i].CumulativeRatio = totalRatio
		}
		if totalRatio != 100 {
			log.Printf("total ratio does not match 100%%")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		byteData, _ := json.Marshal(configBody)
		ioutil.WriteFile("config.json", byteData, fs.FileMode(0777))
		configs = configBody

		log.Println(configBody)
	})

	http.HandleFunc("GET /config", func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("api-key")
		if key != apiKey {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		data, err := json.Marshal(configs)
		if err != nil {
			log.Printf("Error encoding JSON: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(data)
		if err != nil {
			log.Printf("Error writing response: %v", err)
			return
		}
	})
	fmt.Println("Server is running on port 443...")
	if err = http.ListenAndServeTLS(":443", "./cert.pem", "./key.pem", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
		return
	}
}
