package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GIANTDATABASE struct {
	Msg string;
}

func main() {
  lebufur := GIANTDATABASE { Msg: "Server Start!" };

	http.HandleFunc("/donate", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(lebufur.Msg));
    err := json.NewDecoder(r.Body).Decode(&lebufur); //Decode request
    if (err != nil) {
      log.Panic("Could not read input!", err);
    }
	})

	var port string = ":13131"
	fmt.Println("Server started on port" + port)

	log.Fatal(http.ListenAndServe(port, nil))
}
