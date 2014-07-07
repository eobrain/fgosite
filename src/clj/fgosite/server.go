package server
import (
	"clojure/java/io"
	"compojure/route"
	compojure "compojure/core"
        fgo "funcgo/main"
	hiccup "hiccup/core"
	"fgosite/tag"
)

// Map of Funcgo, indexed by ID
fgoResult := ref({})

html := [HTML,
	[HEAD,
		[META, {CHARSET: "utf-8"}],
		[META, {HTTP_EQUIV: "X-UA-Compatible", CONTENT: "IE=edge"}],
		[META, {NAME: "viewport", CONTENT: "width=device-width, initial-scale=1"}],
		tag.Css("//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css"),
		tag.Css("//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap-theme.min.css"),
		tag.Css("css/page.css")
	],
	[BODY,
		[DIV#INSERT_HERE],
		tag.Js("https://ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"),
		tag.Js("//netdna.bootstrapcdn.com/bootstrap/3.1.1/js/bootstrap.min.js"),
		tag.Js("//netdna.bootstrapcdn.com/bootstrap/3.1.1/js/bootstrap.min.js"),
		tag.Js("js/md5.js"),
		tag.Js("js/cljs.js")
	]
]


func clj(id) {
	(*fgoResult)(id)
}

func parse(filename, fgoStr) {
	try {
		fgo.CompileString(filename, fgoStr)
	} catch Exception e {
		str("ERROR in fgo compiling:\n", e->getMessage())
	}
}

compojure.defroutes(app,
	compojure.GET("/",      [], {
		STATUS: 200,
		HEADERS: {"Content-Type": "text/html; charset=utf-8"},
		BODY: hiccup.html(html)
	}),
 	compojure.PUT("/:id/fgo/:filename", request, dosync({
		{{id: ID, filename: FILENAME}: ROUTE_PARAMS, body: BODY} := request
		bodyStr := slurp(io.reader(body))

		alter(fgoResult,
			func(rs){
				rs += { id: parse(filename, bodyStr) }
			}
		)
	})),
	compojure.GET("/:id/clj", [id], clj(id)),
	compojure.GET("/:id/eval", [id], {
		try{
			main := loadString(clj(id))
			withOutStr(main())
		} catch Throwable e {
			str(e)
		}
	}),
	route.resources("/"),
	route.notFound("<h1>Page not found</h1>")
)

var App = app
