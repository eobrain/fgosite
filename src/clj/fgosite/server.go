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
	compojure.GET("/:id/clj", [id], clj(id)),
	compojure.GET("/:id/eval", [id], {
		try{
			const main = loadString(clj(id))
			withOutStr(main())
		} catch Exception e {
			str(e->getMessage())
		}
	}),
	route.resources("/"),
	route.notFound("<h1>Page not found</h1>")
)

func _main(args...) {
	app  jetty.runJetty  {PORT: 3000}
}

