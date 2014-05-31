package server
import (
	"clojure/java/io"
	"ring/adapter/jetty"
	"compojure/route"
	compojure "compojure/core"
        fgo "funcgo/main"
)

// Map of Funcgo, indexed by ID
const fgoResult = ref({})

func parse(filename, fgoStr) {
	try {
		fgo.CompileString(filename, fgoStr)
	} catch Exception e {
		str(e->getMessage())
	}
}

compojure.defroutes(app,
	compojure.GET("/",      [], io.resource("public/index.html")),
	compojure.PUT("/:id/fgo/:filename", request, dosync({
		const(
			{{id: ID, filename: FILENAME}: ROUTE_PARAMS, body: BODY} = request
			bodyStr = slurp(io.reader(body))
		)
		alter(fgoResult, 
			func(rs){
				rs += { id: parse(filename, bodyStr) }
			}
		)
	})),
	compojure.GET("/:id/clj", [id], (*fgoResult)(id)),
	route.resources("/"),
	route.notFound("<h1>Page not found</h1>")
)

func _main(args...) {
	app  jetty.runJetty  {PORT: 3000}
}

