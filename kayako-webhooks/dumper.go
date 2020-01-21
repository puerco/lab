package kayakodumper

import (
	"io/ioutil"
	"log"
	"net/http"
)

// DumpWebHookData Dump de datos a log de func
func DumpWebHookData(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	log.Printf("Recibi datos de webohook:\n%s", body)
	log.Println(string(body))
	//return nil
}
