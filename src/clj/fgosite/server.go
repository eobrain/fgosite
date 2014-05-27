package server
import (
	"clojure/java/io"
	"ring/adapter/jetty"
	"compojure/route"
	compojure "compojure/core"
)


compojure.defroutes(app,
	compojure.GET("/fgo/:id", [id], str("id='", id, "'")),
	compojure.GET("/",      [], io.resource("public/index.html")),
	route.resources("/"),
	route.notFound("<h1>Page not found</h1>")
)

func _main(args...) {
	app  jetty.runJetty  {PORT: 3000}
}

