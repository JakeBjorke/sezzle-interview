package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Knetic/govaluate"
	"github.com/jakebjorke/sezzle-interview/models"
	"github.com/jakebjorke/sezzle-interview/websockets"
)

var broadcastChan chan models.Message

//SetBroadcastChan channel to broadcast good values to websockets
func SetBroadcastChan(c chan models.Message) {
	broadcastChan = c
}

//Expression is used to evaluate expressions
func Expression(w http.ResponseWriter, r *http.Request) {
	input := models.ExpressionRequest{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Println("unable to decode body:  ", err)
		//todo need to return bad request or something
		return
	}

	expression, err := govaluate.NewEvaluableExpression(input.Value)
	if err != nil {
		log.Println("unable to create valid expression:  ", err)
		//todo need to return bad request or something
		return
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		log.Println("unable to evaluate expression:  ", err)
		//todo need to return bad request or something
		return
	}

	response := models.ExpressionResponse{}
	response.Statement = fmt.Sprintf("%s = %v", input.Value, result)
	json.NewEncoder(w).Encode(response)

	if broadcastChan != nil {
		broadcastChan <- models.Message{Type: 2, Body: response.Statement}
	}
}

//ServeWebSockets is used to upgrade serve websockets over a connection
func ServeWebSockets(pool *websockets.Pool, w http.ResponseWriter, r *http.Request) {
	ws, err := websockets.Upgrade(w, r)
	if err != nil {
		//write the output error
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websockets.Client{
		Conn: ws,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}
