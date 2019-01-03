
package main

import (
        "strings"
        "log"
        "net"
        "net/http"
        "os"
        "github.com/apex/gateway"
        "github.com/gin-gonic/gin"
        "github.com/gin-contrib/cors"
)

func GetOutboundIP(h string, p string) net.IP {
    conn, err := net.Dial("tcp", h+":"+p)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    RemoteAddr := conn.RemoteAddr().(*net.TCPAddr)

    return RemoteAddr.IP
}


func defaultHandler(c *gin.Context) {
    for k, v := range c.Request.Header {
        log.Printf("%s -> %s\n", k, v)
    }
    
    t := GetOutboundIP(c.GetHeader("Host"), "https")
    log.Print(t)
    ip := strings.Split(c.GetHeader("X-Forwarded-For"), ",")
     c.JSON(http.StatusOK, gin.H{
        "IP": ip,
        "RemoteIP":t,
    })
}

func routerEngine() *gin.Engine {
    gin.SetMode(gin.DebugMode)
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"*"},
        AllowHeaders:     []string{"Origin"},
    }))
    r.GET("/ips", defaultHandler)
    r.POST("/ips", defaultHandler)
    return r
}

func main() {
    addr := ":" + os.Getenv("PORT")
    log.Fatal(gateway.ListenAndServe(addr, routerEngine()))
}

/*
package main

import (
		"strings"
		"log"
        "net/http"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
        "github.com/aws/aws-lambda-go/events"
        "encoding/json"
)

type Response struct {
    IP string `json:"IP"`
}

func HandleRequest(ctx context.Context,  request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
    for k, v := range c.Request.Header {
        log.Printf("%s -> %s\n", k, v)
    }

	log.Print(request)

	s := strings.Split(request.Headers["X-Forwarded-For"], ",")

	r := &Response{ IP: s[0]}

	j, err := json.Marshal(r)

	if err != nil {
        return events.APIGatewayProxyResponse{}, err
    }

	return events.APIGatewayProxyResponse{Body: string(j), StatusCode: 200}, nil
}

func main() {
        lambda.Start(HandleRequest)
}
*/