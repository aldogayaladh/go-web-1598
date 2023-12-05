package eureka

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"codnect.io/chrono"
)

const (
	statusStarting = "STARTING"
)

type AppRegistrationBody struct {
	Instance InstanceDetails `json:"instance"`
}

type InstanceDetails struct {
	InstanceID       string         `json:"instanceId"`
	HostName         string         `json:"hostName"`
	App              string         `json:"app"`
	IPAddr           string         `json:"ipAddr"`
	VipAddress       string         `json:"vipAddress"`
	SecureVipAddress string         `json:"secureVipAddress"`
	Status           string         `json:"status"`
	Port             Port           `json:"port"`
	SecurePort       Port           `json:"securePort"`
	HealthCheckUrl   string         `json:"healthCheckUrl"`
	StatusPageUrl    string         `json:"statusPageUrl"`
	HomePageUrl      string         `json:"homePageUrl"`
	DataCenterInfo   DataCenterInfo `json:"dataCenterInfo"`
}

type DataCenterInfo struct {
	Class string `json:"@class"`
	Name  string `json:"name"`
}

type Port struct {
	Number  int    `json:"$"`
	Enabled string `json:"@enabled"`
}

func RegisterApp(appId, appName string) {
	fmt.Println("RegisterApp")
	body := buildBody(appId, appName, statusStarting)

	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(body)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8761/eureka/apps/"+appName, "application/json", &buffer)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respBody))
}

func UpdateStatus(appId, appName, status string) {
	fmt.Println("UpdateStatus")
	body := buildBody(appId, appName, status)

	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.Post("http://localhost:8761/eureka/apps/"+appName, "application/json", &buffer)

	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	respBody, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respBody))
}

func DeleteApp(appID, appName string) {
	fmt.Println("DeleteApp")
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8761/eureka/apps/"+appID+"/"+appName+"", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respBody))
}

func SendHeartbeat(appID, appName string) {
	fmt.Println("SendHeartbeat")
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8761/eureka/apps/"+appName+"/"+appID+"", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respBody))
}

func buildBody(appId, appName, status string) *AppRegistrationBody {
	hostName := "192.168.1.2"
	httpPort := 9090

	homePageUrl := fmt.Sprintf("http://%s:%d/", hostName, httpPort)
	statusPageUrl := fmt.Sprintf("http://%s:%d/status", hostName, httpPort)
	healthCheckUrl := fmt.Sprintf("http://%s:%d/healthcheck", hostName, httpPort)
	dataCenterInfo := DataCenterInfo{Class: "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo", Name: "MyOwn"}
	port := Port{Number: 9090, Enabled: "false"}
	securePort := Port{Number: 9090, Enabled: "false"}
	instance := InstanceDetails{
		InstanceID:       appId,
		HostName:         hostName,
		App:              appName,
		VipAddress:       appName,
		SecureVipAddress: appName,
		IPAddr:           "192.168.1.2",
		Status:           status,
		Port:             port,
		SecurePort:       securePort,
		HealthCheckUrl:   healthCheckUrl,
		StatusPageUrl:    statusPageUrl,
		HomePageUrl:      homePageUrl,
		DataCenterInfo:   dataCenterInfo,
	}

	return &AppRegistrationBody{Instance: instance}
}

func ScheduleHeartbeat(appId, appName string) chrono.ScheduledTask {
	taskSheduler := chrono.NewDefaultTaskScheduler()
	task, err := taskSheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
		fmt.Println("From ScheduleHeartbeat")
		SendHeartbeat(appId, appName)
	}, 25*time.Second)

	if err != nil {
		log.Fatal(err)
	}

	return task

}
