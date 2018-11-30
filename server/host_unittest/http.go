package main

import (
    "fmt"
	"time"
    "log"
    "os/exec"
    "net/http"
    "encoding/json"

)

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type currentViewIno  struct {
	Channel string `json:"text"`
	ProgramTitle string `json:"text"`
}

var switchState string = "ON"
var remoteState string = "ON"
var currentViewChannel string =""
var currentViewProgram string = ""
var guideChannel string=""
var guideProgram string=""
var guideTitle string=""
var switchBool int = 1
var duration string = ""
var ProgStartTime string = ""
var name string = ""
var mac string = ""
var ip string = ""
var state string = ""   

func switchHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	
	if ( r.Method == "GET") {
	
		fmt.Fprintf(w, "{ \"status\" : \"%s\" }", switchState )
		
	}	
	if ( r.Method == "POST" ) {
		newState := r.URL.Query().Get("state")
		if ( ( newState == "ON" ) || ( newState == "OFF" ) ){
			switchState = newState;
		} else {
			log.Printf("Invalid state passed for HTTP PORT on %s",r.RequestURI)
		}
		fmt.Fprintf(w, "{ \"status\" : \"%s\" }", switchState )
	}
	log.Printf(
            "%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            time.Since(start),
			switchState,
        )
}


func remotehandler(w http.ResponseWriter, r *http.Request) {
	k := r.URL.Query().Get("remotekey")

        //NEW CODE
    fmt.Println("URL.String is %s",r.URL.String())

    fmt.Printf("\nparsed key :  %s",k)
    ip := r.URL.Query().Get("stbip")
    fmt.Printf("\nstbip :  %s , key : %s, keylen = %d ",ip, k, len(k))
	

	
    cmd := fmt.Sprintf("cmd2000 %s \"osdiag %s\"",ip, k)
    //Handle key including + symbol, its dropped by HTTP URL, append it again
    if ( (k == "c ") || (k == "v ") || (k == "p ") || ( k == "d ") ) {
        cmd = fmt.Sprintf("cmd2000 %s \"osdiag %s+\" ",ip, k[0:len(k)-1])
		fmt.Printf("Updated command for plus/minus is %s", cmd)
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
        }
	}

func stbinfohandler(w http.ResponseWriter, r *http.Request) {
	//start := time.Now()
	
   // fmt.Fprintf(w, "{ \"status\" : \"%s\" }", "OK" )
	//log.Printf(
    //        "%s\t%s\t%s\t%s",
    //        r.Method,
     //       r.RequestURI,
     //       time.Since(start),
	//		switchState,
    //    )
	if ( r.Method == "GET" ) {
		fmt.Fprintf(w, "{ \"DeviceType\" : \"%s\", \"DeviceMac\" : \"%s\",\"IpAddress\" : \"%s\" , \"DeviceState\" : \"%s\" }" , name, mac,ip,state )
	
		log.Printf(
            "%s\t%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            name,
			mac,
			ip,
			state,
		)
	}
	if( r.Method == "POST" ) {
		name = r.URL.Query().Get("DeviceType")
		mac = r.URL.Query().Get("DeviceMac")
		ip = r.URL.Query().Get("IpAddress")
		state = r.URL.Query().Get("DeviceState")
		fmt.Fprintf(w, "{ \"status\" : \"%s\" }", "OK" )
		//fmt.Fprintf(w, "{ \"status\" : \"%s\" }", "OK" )
		log.Printf(
            "%s\t%s\t%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            name,
			mac,
			ip,
			state,
		)
	}
}

func currentViewinghandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if ( r.Method == "GET" ) {
		fmt.Fprintf(w, "{ \"channel\" : \"%s\", \"progTitle\" : \"%s\" , \"programDuration\" : \"%s\" , \"ProgStartTime\" : \"%s\" }" , currentViewChannel, currentViewProgram,duration, ProgStartTime)
	
		log.Printf(
            "%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            time.Since(start),
		)
	}
	if( r.Method == "POST" ) {
		currentViewChannel = r.URL.Query().Get("channel")
		currentViewProgram = r.URL.Query().Get("programTitle")
	    duration = r.URL.Query().Get("programDuration")
		ProgStartTime = r.URL.Query().Get("ProgStartTime")
		fmt.Fprintf(w, "{ \"status\" : \"%s\" }", "OK" )
		log.Printf(
            "%s\t%s\t%s\t%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            time.Since(start),
			currentViewChannel,
			currentViewProgram,
			duration,
			ProgStartTime,
		)
	}
}

func guideHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if ( r.Method == "GET" ) {
	fmt.Fprintf(w, "{ \"guideChannel\" : \"%s\", \"guideTitle\" : \"%s\", \"guideProgram\" : \"%s\" }" , guideChannel, guideTitle, guideProgram )
	
		log.Printf(
            "%s\t%s\t%s %s%s%s",
            r.Method,
            r.RequestURI,
            time.Since(start),guideChannel, guideTitle, guideProgram,
		)
		}
		
		if( r.Method == "POST" ) {
		guideChannel = r.URL.Query().Get("channelNumber")
		guideTitle = r.URL.Query().Get("title")
		guideProgram = r.URL.Query().Get("programDesc")
		fmt.Fprintf(w, "{ \"status\" : \"%s\" }", "OK" )
		log.Printf(
            "%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            time.Since(start),
			"Chan="+guideChannel+"PRogramTitle="+guideTitle+"PRogramDesc="+guideProgram,
		)
	}
}

func main() {
    http.HandleFunc("/remote", remotehandler)
	http.HandleFunc("/switch", switchHandler)
	http.HandleFunc("/stbinfo", stbinfohandler)
	http.HandleFunc("/currentViewing", currentViewinghandler)
	http.HandleFunc("/guideNavigation", guideHandler)
	http.HandleFunc("/askalexa", stbinfohandler)
    log.Fatal(http.ListenAndServe(":8082", nil))
}