package main

import (
	"encoding/json"

	logging "github.com/Roma004/go-logger"
)

type Abc struct {
	Abc string `json:"field_one"`
	Asd string `json:"field_two"`
}

func main() {

	log := logging.MyLog{}

	log.Info("Start of func main")

	a := Abc{
		Abc: "hello from Golang",
		Asd: "I hate coding in Go ;(",
	}

	b, err := json.Marshal(a)
	if err != nil {
		log.Error("JSON Marshal error", err)
	}
	log.LargeDebug(b, "\n\n", string(b))
	log.Info("struct Abc marshaled, result:", string(b))

	b = append(b, 18)
	log.Warning("Try to fail json.Unmarshal")

	var data Abc
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Error("Unmarshalling data failure", err)
		log.FailureStatus("Unmarshalling data", "FAILURE")
	}

	log.Debug(data)

	log.Info("End of func main")
}
