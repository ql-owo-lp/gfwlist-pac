package gfwlistpac

import (
	"html/template"
	"net/http"
	"time"
	"fmt"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

// [START greeting_struct]
type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}

// [END greeting_struct]

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
	http.HandleFunc("/pac", genProxy)
}

// guestbookKey returns the key used for all guestbook entries.
func guestbookKey(c appengine.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}

// [START func_root]
func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	// Ancestor queries, as shown here, are strongly consistent with the High
	// Replication Datastore. Queries that span entity groups are eventually
	// consistent. If we omitted the .Ancestor from this query there would be
	// a slight chance that Greeting that had just been written would not
	// show up in a query.
	// [START query]
	q := datastore.NewQuery("Greeting").Ancestor(guestbookKey(c)).Order("-Date").Limit(10)
	// [END query]
	// [START getall]
	greetings := make([]Greeting, 0, 10)
	if _, err := q.GetAll(c, &greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// [END getall]
	if err := guestbookTemplate.Execute(w, greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// [END func_root]

func genProxy(w http.ResponseWriter, r *http.Request) {
	proxy := []string {
		"SOCKS5 127.0.0.1:8088",
	}

	addtnlList := map[string]HostEntry {
		"google.com" : HostEntry {
			httpEnabled : true,
			httpsEnabled : true,
			httpDisabled : false,
			httpsDisabled : false,
			keywordEnabled : false,
		},
		"appspot.com" : HostEntry {
			httpEnabled : true,
			httpsEnabled : true,
			httpDisabled : false,
			httpsDisabled : false,
			keywordEnabled : false,
		},
		"twitter.com" : HostEntry {
			httpEnabled : true,
			httpsEnabled : true,
			httpDisabled : false,
			httpsDisabled : false,
			keywordEnabled : false,
		},
		"facebook.com" : HostEntry {
			httpEnabled : true,
			httpsEnabled : true,
			httpDisabled : false,
			httpsDisabled : false,
			keywordEnabled : false,
		},
	}

	pac := GeneratePac(proxy, addtnlList)

	w.Header().Set("Content-Type", "application/x-ns-proxy-autoconfig")

	fmt.Fprintf(w, pac)
}

var guestbookTemplate = template.Must(template.New("book").Parse(`
<html>
  <head>
    <title>Go Guestbook</title>
  </head>
  <body>
    {{range .}}
      {{with .Author}}
        <p><b>{{.}}</b> wrote:</p>
      {{else}}
        <p>An anonymous person wrote:</p>
      {{end}}
      <pre>{{.Content}}</pre>
    {{end}}
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`))

// [START func_sign]
func sign(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	g := Greeting{
		Content: r.FormValue("content"),
		Date:    time.Now(),
	}
	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}
	// We set the same parent key on every Greeting entity to ensure each Greeting
	// is in the same entity group. Queries across the single entity group
	// will be consistent. However, the write rate to a single entity group
	// should be limited to ~1/second.
	key := datastore.NewIncompleteKey(c, "Greeting", guestbookKey(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// [END func_sign]
