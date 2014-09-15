package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/user"
	"time"

	"github.com/guelfey/go.dbus"
)

type PowerGrid struct {
	Status string `json`
}

func Notify(status string) {
	var (
		usr, _ = user.Current()
		icon   = usr.HomeDir + "/.icons/yalahweey.png"
		app    = "Yalahweey"
		text   = status
		title  = "Yalahweey"
		dst    = "org.freedesktop.Notifications"
		path   = "/org/freedesktop/Notifications"
		method = "org.freedesktop.Notifications.Notify"
		sound  = "/usr/share/sounds/freedesktop/stereo/message.oga"
		hint   = map[string]dbus.Variant{
			"sound-file": dbus.MakeVariant(sound)}
	)

	conn, err := dbus.SessionBus()
	HandleError(err)
	obj := conn.Object(dst, dbus.ObjectPath(path))
	call := obj.Call(method, 0, app, uint32(0),
		icon, title, text, []string{},
		hint, int32(5000))
	HandleError(call.Err)
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func FetchState() (status string) {
	url := "http://www.gridstatusnow.com/status"
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
