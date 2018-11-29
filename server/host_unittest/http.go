package main

import (
    "fmt"
	"time"
    "log"
    "net/http"
)
var switchState string = "ON"
var remoteState string = "ON"
var switchBool int = 1

func switchHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if  switchState == "ON" {
		switchState = "OFF"
	} else {
		switchState = "ON"
	}
	
	if ( r.Method == "GET") {
	
		fmt.Fprintf(w, "{ \"status\" : \"%s\" }", switchState )
		log.Printf(
            "%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            time.Since(start),
			switchState,
        )
	}	else {
		log.Printf("Method not supported")
	}
}


func remotehandler(w http.ResponseWriter, r *http.Request) {
	k := r.URL.Query().Get("remotekey")

        //NEW CODE
    fmt.Println("URL.String is %s",r.URL.String())

        fmt.Printf("\nparsed key :  %s",k)
        ip := r.URL.Query().Get("stbip")
        fmt.Printf("\nstbip :  %s , key : %s ",ip, k)
        cmd := fmt.Sprintf("cmd2000 %s \"osdiag %s\"",ip, k)
        //Handle key including + symbol, its dropped by HTTP URL, append it again
        if ( (k == "c") || (k == "v") || (k == "p") || ( k == "d") ) {
                cmd = fmt.Sprintf("cmd2000 %s \"osdiag %s+\"",ip, k)
        }
        fmt.Printf("\nCommand is  %s : ",cmd)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(jsonErr{Code : http.StatusOK , Text : "Success"} ); err != nil {
                panic(err)
        }

    out, err :=  exec.Command("sh","-c",cmd ).Output()
    if(err != nil){
                fmt.Printf("%s",err)
        }else{
                fmt.Printf("%s",out);
        }}

func stbinfohandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	
    fmt.Fprintf(w, "{ \"status\" : \"%s\" }", "OK" )
	log.Printf(
            "%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            time.Since(start),
			switchState,
        )
}

func currentViewinghandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	
    fmt.Fprintf(w, "{ \"status\" : \"%s\" }", "OK" )
	log.Printf(
            "%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            time.Since(start),
			switchState,
        )
}

func main() {
    http.HandleFunc("/remote", remotehandler)
	http.HandleFunc("/switch", switchHandler)
	http.HandleFunc("/stbinfo", stbinfohandler)
	http.HandleFunc("/currentViewing ", currentViewinghandler)
    log.Fatal(http.ListenAndServe(":8082", nil))
}