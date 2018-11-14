package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/mongo"
)

const (
	LayoutYYYYMMDDTHHMMSSZ = "2006-01-02T15:04:05Z"
	CollectExercise3       = "exercise3"
	Host                   = "localhost"
	PortServer             = ":8081"
	Database               = "webvue_db"
	DBType                 = "mongo"
)

type Configuration struct {
	Connection dbox.IConnection
}

func New(configuration *Configuration) *Configuration {
	return &Configuration{configuration.Connection}
}

func CreateConn() dbox.IConnection {
	ci := dbox.ConnectionInfo{}
	ci.Host = Host
	ci.Password = ""
	ci.UserName = ""
	ci.Database = Database
	ci.Settings = nil
	conn, err := dbox.NewConnection(DBType, &ci)
	if err != nil {
		panic("Connect Failed") // Change with your error handling
	}
	err = conn.Connect()
	if err != nil {
		panic("Connect Failed") // Change with your error handling
	}
	return conn
}
func FormatDateToStr(layout string, input time.Time) string {
	//2006-01-02T15:04:05Z
	sNow := input.Format(layout)
	return sNow
}

func FormatStrToDate(layout string, input string) time.Time {
	//2006-01-02T15:04:05Z
	parse, err := time.Parse(layout, input)
	if err != nil {
		fmt.Println("Error FormatStrToDate : ", err)
	}
	fmt.Println(parse)
	return parse
}
func TimeAdd(parse time.Time, addDuration time.Duration, typeData string) time.Time {
	if typeData == "Hour" {
		ret := parse.Add(time.Duration(addDuration) * time.Hour) //tipe date
		fmt.Println(ret)
		return ret
	}
	return time.Time{}
}

func StringToMap(myString string) map[string]interface{} {

	data := make(map[string]interface{})
	err2 := json.Unmarshal([]byte(myString), &data)
	if err2 != nil {
		fmt.Println("Error!", err2.Error())
	}
	return data
}
func ArrayObjectToString(retFetch map[string]interface{}) string {

	var Buf bytes.Buffer
	b, _ := json.Marshal(retFetch["data"])
	Buf.Write(b)
	myString := Buf.String()
	Buf.Reset()
	return myString
}
func DiffYears(a, b time.Time) (year int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month := int(M2 - M1)
	day := int(d2 - d1)
	hour := int(h2 - h1)
	min := int(m2 - m1)
	sec := int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
