package server
import (
	"ring/adapter/jetty"
	"ring/middleware/resource"
	"ring/util/response"
)


func renderApp() {
	{
		STATUS:  200,
		HEADERS: {"Content-Type": "text/html"},
		BODY: 	 `<!DOCTYPE html>
<html>
  <head>
    <link rel="stylesheet" href="css/page.css" />
  </head>
  <body>
    <div>
      <p id="clickable">Click me!</p>
    </div>
    <div id="insert-here"/>
    <script src="js/cljs.js"></script>
  </body>
</html>
`
	}
}

func handler(request) {
	if "/" == URI(request) {
		response.redirect("/index.html")
	} else {
		renderApp()
	}
}

var app = resource.wrapResource(handler, "public")

func _main(args...) {
	app  jetty.runJetty  {PORT: 3000}
}

