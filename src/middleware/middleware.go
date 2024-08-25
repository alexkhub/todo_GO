package middleware

import ("fmt" 
		"time"
		"net/http"
)


func ConsoleLogMiddleware(w http.ResponseWriter, r *http.Request){
	defer func(start time.Time){
		fmt.Printf("[%s] %s ", time.Since(start), r.URL.Path, )
	}(time.Now())
}

