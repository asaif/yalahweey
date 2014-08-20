package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"
	"os/user"
	"time"
)

type PowerGrid struct {
	Status string `json`
}

func Notify(status string) {
	usr, _ := user.Current()
	icon := usr.HomeDir + "/.icons/yalahweey.png"
	n := exec.Command("notify-send", "-i", icon, "Yalahweey", status)
	n.Run()
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func FetchState() (status string) {
	url := "http://power-grid-status.mos3abof.com/status"
	res, err := http.Get(url)
	HandleError(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	HandleError(err)
	var power PowerGrid
	err = json.Unmarshal(body, &power)
	HandleError(err)
	status = power.Status
	return status
}

func BindState(status string) (message string) {

	Hstate := map[string]string{
		"Danger":  "Dude we are doomed || Danger Zone",
		"Warning": "Hurry up unplug some cords || Warning Zone",
		"Safe":    "Ya7alawa ya wlad || Safe Zone",
		"Unknown": "I don't known what the current state",
	}

	message = Hstate[status]
	return message
}

func main() {
	Notify(BindState(FetchState()))
	for _ = range time.Tick(15 * time.Minute) {
		Notify(BindState(FetchState()))

	}
}
