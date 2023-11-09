package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	puerto = ":8080"
)

// Persona es una estructura que define ....
type Persona struct {
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Edad      int    `json:"edad"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
	Activo    bool   `json:"activo"`
}

func main() {

	// Punto 1.
	jsonPersona := `{
		"nombre":"Juan",
		"apellido":"Perez",
		"edad":25,
		"direccion":"Av. Siempre Viva",
		"telefono":"123434234",
		"activo":true
	}`

	var persona Persona
	err := json.Unmarshal([]byte(jsonPersona), &persona)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(persona)

	// Punto 2.
	// Instancia default gin
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"mensaje": "pong",
		})
	})

	personaResponse := Persona{
		Nombre:    "Pedro",
		Apellido:  "Pascal",
		Edad:      45,
		Direccion: "Av. Paz",
		Telefono:  "234234234",
		Activo:    true,
	}

	router.GET("/persona", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": personaResponse,
		})

	})

	if err := router.Run(puerto); err != nil {
		log.Fatal(err)
	}

}
