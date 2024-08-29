package middleware

import ("fmt" 
		"time"
		"net/http"
)


func ConsoleLogMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		func(start time.Time){
			fmt.Printf("[%s] | %s |  %s%s \n", start.Format(time.RFC3339), r.Method, r.Host,r.URL.Path)
		}(time.Now())
		next.ServeHTTP(w, r)
	})	
}

