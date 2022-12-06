package routes

import (

    "fmt"
	"net/http"
	"encoding/xml"


	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

type Day struct {
    id int
    Name  string
}


type Month struct {
    id int
    Name  string
}


func GetMonthOfYear(w http.ResponseWriter, r* http.Request) {
	params := mux.Vars(r)
    var month *Month
	switch params["id"] {
    case "1":
       month= &Month{1, "January"}
    case "2":
        month= &Month{2, "February"}
    case "3":
        month= &Month{3, "March"}
	case "4":
        month= &Month{4, "April"}
	case "5":
        month= &Month{5, "May"}
	case "6":
        month= &Month{6, "June"}
	case "7":
        month= &Month{7, "July"}
	case "8":
        month= &Month{8, "August"}
    case "9":
        month= &Month{9, "September"}
    case "10":
        month= &Month{10, "October"}
    case "11":
        month= &Month{11, "November"}
    case "12":
        month= &Month{12, "December"}


	default:
		http.Error(w, "Month Not Found",  http.StatusInternalServerError)
        return
			
    }
   

	x, err := xml.MarshalIndent(month, "", "  ")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }


    w.Header().Set("Content-Type", "application/xml")
    // Write
    w.Write(x)



	
}










func GetDayOfWeek(w http.ResponseWriter, r* http.Request) {
	params := mux.Vars(r)

    var day *Day

	switch params["id"] {
    case "1":
       day= &Day{1, "Monday"}
    case "2":
        day= &Day{2, "Tuesday"}
    case "3":
        day= &Day{3, "Wednesday"}
	case "4":
        day= &Day{4, "Thursday"}
	case "5":
        day= &Day{5, "Friday"}
	case "6":
        day= &Day{6, "Saturday"}
	case "7":
        day= &Day{7, "Sunday"}
	default:
		http.Error(w, "Day Not Found",http.StatusInternalServerError)
        return
			
    }
   

	x, err := xml.MarshalIndent(day, "", "  ")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }


    w.Header().Set("Content-Type", "application/xml")
    // Write
    w.Write(x)



	
}





func RegisterAPIRoutes(r* mux.Router) {
	r.HandleFunc("/api/dayOfWeek/{id}", GetDayOfWeek).Methods("GET")
	r.HandleFunc("/api/monthOfYear/{id}", GetMonthOfYear).Methods("GET")

	
}