package id

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mt1976/trnsl8r_service/app/support/date"
	"github.com/mt1976/trnsl8r_service/app/support/logger"
	"github.com/segmentio/ksuid"
	"golang.org/x/exp/rand"
)

const SEP = "."

func Encode(in string) string {

	out := in
	out = strings.Replace(out, " ", "", -1)
	out = strings.Replace(out, "_", "", -1)
	out = strings.Replace(out, "/", "", -1)
	out = strings.Replace(out, "\\", "", -1)
	out = strings.Replace(out, ":", "", -1)
	out = strings.Replace(out, ".", "", -1)
	out = strings.Replace(out, ",", "", -1)
	out = strings.Replace(out, ";", "", -1)
	out = strings.Replace(out, "'", "", -1)
	out = strings.Replace(out, "\"", "", -1)
	out = strings.Replace(out, "(", "", -1)
	out = strings.Replace(out, ")", "", -1)
	out = strings.Replace(out, "[", "", -1)
	out = strings.Replace(out, "]", "", -1)
	out = strings.Replace(out, "{", "", -1)
	out = strings.Replace(out, "}", "", -1)
	out = strings.Replace(out, "<", "", -1)
	out = strings.Replace(out, ">", "", -1)
	out = strings.Replace(out, "?", "", -1)
	out = strings.Replace(out, "!", "", -1)
	out = strings.Replace(out, "@", "", -1)
	out = strings.Replace(out, "#", "", -1)
	out = strings.Replace(out, "$", "", -1)
	out = strings.Replace(out, "%", "", -1)
	out = strings.Replace(out, "^", "", -1)
	out = strings.Replace(out, "&", "", -1)
	out = strings.Replace(out, "*", "", -1)
	out = strings.Replace(out, "+", "", -1)
	out = strings.Replace(out, "=", "", -1)
	out = strings.Replace(out, "~", "", -1)
	out = strings.Replace(out, "`", "", -1)
	out = strings.Replace(out, "|", "", -1)
	out = strings.Replace(out, ":", "", -1)
	out = strings.Replace(out, "%", "", -1)

	//working := out

	out = base64.StdEncoding.EncodeToString([]byte(out))
	// remove padding
	out = strings.Replace(out, "=", "-", -1)

	//msg := fmt.Sprintf("[ID] Encoding In=[%v] Working=[%v] Encoded=[%v]", in, working, out)
	//logger.InfoLogger.Println(msg)

	return out
}

func GetUUID() string {
	// Get a new UUID
	// Get TODAY and convert to string
	today := time.Now().Format("060102-150405.000000")
	today = today + ""
	today = strings.Replace(today, ".", "-", -1)
	//xx := shortuuid.New()
	uid := 000000
	if os.Getuid() > 0 {
		uid = os.Getuid()
	}

	//ip, _ := get_IP()
	//ip = strings.Replace(ip, ".", "", -1)
	xx := rand.Intn(100000)
	yy := fmt.Sprintf("%s-%06d-%06d", today, uid, xx)
	yy = strings.Replace(yy, ".", "", -1)
	yy = strings.Replace(yy, "-", "", -1)
	//yy = base64Encode(yy)

	//	logger.InfoLogger.Printf("[UUID] %v %v", yy, UUID2String(yy))

	return yy
}

func UUID2String(uuid string) string {
	// Convert UUID to string
	// 2407032122304271385011014720229731 convert to 240703\212230\427138\501\1014720229\731
	// 2407032122304271385011014720229731 convert to 240703.212230.427138.501.1014720229.731
	// 2407032122304271385011014720229731 convert to 240703-212230-427138-501-1014720229-731
	//logger.InfoLogger.Println("UID: UUID: ", uuid, len(uuid))
	fmtr := "%s" + SEP + "%s" + SEP + "%s" + SEP + "%s" + SEP + "%s"
	op := fmt.Sprintf(fmtr, uuid[0:6], uuid[6:12], uuid[12:18], uuid[18:24], uuid[24:])
	day, _ := time.Parse("060102150405", uuid[0:12])
	fmtr2 := "(Date=[%s]" + " " + "Time=[%s]" + " " + "ms=[%sms]" + " " + "uid=[%s]" + " " + "rnd=[%s])"
	op2 := fmt.Sprintf(fmtr2, date.FormatHuman(day), day.Format("15:04:05"), uuid[12:18], strings.TrimLeft(uuid[18:24], "0"), uuid[24:])
	//logger.InfoLogger.Println("UID: String:", op, len(op))
	return op + ", " + op2
}

func GetUUIDv2() string {
	return ksuid.New().String()
}

func GetUUIDv2WithPayload(payload string) (string, error) {
	// Ensure payload is 16 bytes

	if len(payload) > 16 {
		return "", fmt.Errorf("Payload must be 16 bytes or less")
	}
	if len(payload) < 16 {
		payload = fmt.Sprintf("%-16s", payload)
	}
	ksuid, err := ksuid.FromParts(time.Now(), []byte(payload))
	if err != nil {
		logger.ErrorLogger.Println("Error generating KSUID:", err)
		return "", err
	}
	return ksuid.String(), nil
}

func InspectUUIDv2(uuid string) string {
	ksuid, err := ksuid.Parse(uuid)
	if err != nil {
		logger.ErrorLogger.Println("Error parsing KSUID:", err, " got:", len(uuid), " uuid", uuid)
		return ""
	}
	payload := ksuid.Payload()
	return fmt.Sprintf("Time: %v, Payload: %v", ksuid.Time(), string(payload))
}
