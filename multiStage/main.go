package main
// external libraries
import (
    "encoding/json"
    "net/http"
)
import (
    "app/models"
    "app/repos"
)

func mainResponse(rw http.ResponseWriter, req *http.Request) {
    message := models.Message{Text: "Hello, World!"}
    messageAsJson, err := json.Marshal(message)
    if err != nil {
        http.Error(rw, "Failed to marshal JSON", http.StatusInternalServerError)
        return
    }

    // Set the Content-Type header
    rw.Header().Set("Content-Type", "application/json")

    // Write the JSON response
    _, err = rw.Write(messageAsJson)
    if err != nil {
        http.Error(rw, "Failed to write response", http.StatusInternalServerError)
        return
    }
}
func responseJugadores(rw http.ResponseWriter, req *http.Request) {

    jugadores, err := repos.ReadJugadoresFromDB()
    if err != nil {
        http.Error(rw,err.Error(), http.StatusInternalServerError)
        return
    }

    jugadoresAsJson, err := json.Marshal(jugadores)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }

    // Set the Content-Type header
    rw.Header().Set("Content-Type", "application/json")

    // Write the JSON response
    _, err = rw.Write(jugadoresAsJson)
    if err != nil {
        http.Error(rw, "Failed to write response", http.StatusInternalServerError)
        return
    }
}
func main() {
    // Define routes
    http.HandleFunc("/", mainResponse)
    http.HandleFunc("/jugadores", responseJugadores)
    println("Server started at http://localhost:8080")

    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
