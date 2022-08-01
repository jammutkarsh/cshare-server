package utils

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"os"
	"reflect"
	"runtime"
	"time"
)

const (
	ColorReset  = string("\033[0m")
	ColorRed    = string("\033[31m")
	ColorGreen  = string("\033[32m")
	ColorYellow = string("\033[33m")
	ColorBlue   = string("\033[34m")
	ColorPurple = string("\033[35m")
	ColorCyan   = string("\033[36m")
	ColorWhite  = string("\033[37m")
)

// LoadEnv loads the .env file	and returns the error if any.
func LoadEnv(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		ErrorReaderWriter(err, LoadEnv)
		log.Fatalf("error loading .env file: %s", err.Error())
	}
}

// ByteToData converts the byte array to the given type.
func ByteToData(data []byte) (v interface{}, err error) {
	err = json.Unmarshal(data, &v)
	return v, err
}

// DataToByte converts the given data to byte array.
func DataToByte(v interface{}) (data []byte, err error) {
	data, err = json.Marshal(v)
	return data, err
}

// GetCurrentFuncName returns the name of the function calling the function.
func GetCurrentFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// ErrorReaderWriter writes the error to the log file.
func ErrorReaderWriter(errMsg error, i interface{}) {
	file, err := os.OpenFile("error.logs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %s", err.Error())
	}
	errorMessage := time.Now().Format("2006-01-02 15:04:05") + "\t" + GetCurrentFuncName(i) + "\t" + errMsg.Error() + "\n"
	// ColorErrorMessage to log the error in color.
	//ColorErrorMessage := ColorWhite + time.Now().Format("2006-01-02 15:04:05") + ColorReset + "\t" +
	//	ColorYellow + GetCurrentFuncName(i) + ColorReset + "\t" +
	//	ColorRed + errMsg.Error() + ColorReset + "\n"
	_, err = file.WriteString(errorMessage)
	if err != nil {
		log.Fatalf("error writing to file: %s", err.Error())
	}
	defer func() {
		var f *os.File
		err := f.Close()
		if err != nil {
		}
	}()

}
