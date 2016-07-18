package app

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

type ShortUrl struct {
	Value string
}

func shortHandler(w http.ResponseWriter, r *http.Request) {
	p := r.FormValue("p")
	ctx := appengine.NewContext(r)

	k := datastore.NewKey(ctx, "Entity", "stringID", 0, nil)
	e := new(ShortUrl)

	if p == "create" {
		dataCreate(w, r, k, e)
	}
	// if err := datastore.Get(ctx, k, e); err != nil {
	//	http.Error(w, err.Error(), 500)

	// }

	// old := e.Value
	// e.Value = "Hello World!"

	// if _, err := datastore.Put(ctx, k, e); err != nil {
	//	http.Error(w, err.Error(), 500)
	// }

	// w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// fmt.Fprintf(w, "old=%q\nnew=%q\np=%q\n", old, e.Value, p)

}

func dataCreate(w http.ResponseWriter, r *http.Request, k *datastore.Key, e interface{}) {

	ctx := appengine.NewContext(r)
	if _, err := datastore.Put(ctx, k, e); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func dataRead(w http.ResponseWriter, r *http.Request) {

}

func dataUpdate(w http.ResponseWriter, r *http.Request) {

}

func dataDelete(w http.ResponseWriter, r *http.Request) {

}
