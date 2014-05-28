package server
import (
	"clojure/java/io"
	"ring/adapter/jetty"
	"compojure/route"
	compojure "compojure/core"
        fgo "funcgo/core"
)

// Map of Funcgo, indexed by ID
const fgoResult = ref({})

compojure.defroutes(app,
	compojure.GET("/",      [], io.resource("public/index.html")),
	compojure.PUT("/:id/fgo", request, dosync({
		const(
			{{id: ID}: ROUTE_PARAMS, body: BODY} = request
			bodyStr = slurp(io.reader(body))
		)
		println(id, bodyStr)
		alter(fgoResult, 
			func(rs){
				rs += { id: fgo.Parse("main.go", bodyStr) }
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

