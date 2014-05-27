package server
import (
	"clojure/java/io"
	"ring/adapter/jetty"
	"compojure/route"
	compojure "compojure/core"
	rest "liberator/core"
)

rest.defresource(helloWorld,
	AVAILABLE_MEDIA_TYPES, ["text/plain"],
	HANDLE_OK,"Hello, world! -- 2"
)

compojure.defroutes(app,
	compojure.GET("/",      [], io.resource("public/index.html")),
	compojure.GET("/hello", [], helloWorld),
	route.resources("/"),
	route.notFound("<h1>Page not found</h1>")
)

func _main(args...) {
	app  jetty.runJetty  {PORT: 3000}
}

