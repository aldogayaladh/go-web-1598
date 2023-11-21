package middleware

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Esto es extra//
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

/////////////////

// Esto es extra//
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

/////////////////

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtenemos el verbo, la url y el tiempo
		verbo := ctx.Request.Method
		time := time.Now()
		url := ctx.Request.URL
		var size int

		//Esto es extra//
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		/////////////////
		ctx.Next()

		if ctx.Writer != nil {
			size = ctx.Writer.Size()

		}

		fmt.Printf("\nPath:%s\nVerbo:%s\nTiempo:%v\nTamaño consulta:%d\n", url, verbo, time, size)
		//Esto es extra//
		fmt.Printf("Body:%s\n", blw.body.String())
		/////////////////

	}

}
