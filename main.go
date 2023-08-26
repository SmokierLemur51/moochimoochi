package main

import (
	"net/http"
	"log"
	"fmt"

	"github.com/go-chi/chi/v5"

	"github.com/SmokierLemur51/moochimoochi/routes"
)

func calc(monthlyRevenue float32) float32 {
	var month float32
	var flat float32 = 350.00
	var percent float32 = .02
	var total float32 = 0.00

	for month = 1.00; month < 24.00; month++ {
		total += (monthlyRevenue*percent)
		if total > flat {
			return month
		} else {
			fmt.Printf("Month: %f\r\nTotal: %f\r\n\n", month, total)
		}
	}
	return flat
}




func main() {

	x := calc(1000)
	fmt.Println(x)

	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	routes.ConfigureRoutes(r)

	log.Println("Starting server on :5000")
	http.ListenAndServe(":5000", r)
}