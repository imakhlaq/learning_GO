package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	if err := http.ListenAndServe(":8081", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		data := struct {
			id   string
			name string
		}{
			"1233",
			"Akhlaq",
		}

		/*
			One important note about this version is that it uses a byte slice in jData, possibly unnecessarily.
			Data can be of arbitrary size, depending on the data being marshalled, so this could be a non-trivial memory waster.
			After marshalling, we copy from memory to the ResponseWriter stream. The answer that uses json.NewEncoder() etc.
			would write the marshalled JSON straight into the ResponseWriter ( into its stream )
		*/

		/*
			jData, _ := json.Marshal(data)

			fmt.Println(r)

			w.WriteHeader(http.StatusOK)

			fmt.Println("starting writing")
			_, err1 := w.Write(j)

			fmt.Println("done writing")
			if err1 != nil {
				return
			}*/

		//==============================================================================


		// var data struct {
		// 	URL string `json:"url"`
		// }
		// err := json.NewDecoder(r.Body).Decode(&data)

		// better approach because you're directly writing it to response writer
		// no copy

		w.Header().Set("Content-Type", "application/json")

		//
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			fmt.Println("encoding failed")
			return
		}
		fmt.Println("written")

		return

	})); err != nil {
		fmt.Println("error while creating server", err.Error())
	}
}
