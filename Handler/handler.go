package Handler

import (
	"encoding/json"
	"glutton/GStruct"
	"math/rand"
	"net/http"
)

//HomeHandler is Default Landing Handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	randID := GStruct.RANDOMID{}
	randID.ID = rand.Intn(9999)
	slcB, _ := json.Marshal(randID)
	w.Write(slcB)

}
