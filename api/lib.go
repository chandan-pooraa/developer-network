package api

import "log"

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
func Logif(err error, message string){
	if err!=nil{
		log.Fatal(message)
	}
}

