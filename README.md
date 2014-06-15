# Functional Go Website

This is the code for the [Functional Go Web Site][1], and is itself
written in Functional Go.  Client-side Functional Go generates
JavaScript which powers the user interface, while server-side
Functional Go runs on the server.

Feel free to use this project as a model for your own Functional Go
web application, but be aware that in one respect it is not typical,
it uses the Functional Go compiler to enable the interactive coding
interface, something that is not needed in a typical web application.

## Build

Follow the [install instructions for Leiningen][lein] and go to the
top level directory of this project:

### Running Locally

To run locally on your laptop do the following from the command line:
```sh
lein do  fgoc,  cljsbuild once,  ring server-headless
```
And then visit http://localhost:3000/

### Deploying
You can deploy to any server environment that supports standard WARs.

From the command line do:
```sh
lein do  fgoc,  cljsbuild once,  ring uberwar
```
which generates a file WAR file in the `target` directory, which you
can then deploy.

## Architecture

The client-code is a single-page app that uses AJAX to talk to a REST
server.

### Server

The server component is `src/clj/fgosite/server.go`.  It implements a
simple REST API:

1. `/` GET the single static HTML/CSS/JS page
2. `/:id/fgo/:filename` PUT the Functional Go code with this ID
3. `/:id/clj` GET the Clojure code resulting from compiling the Functional Go with this ID
4. `/:id/eval` GET the result of evaluating the Clojure code with this ID

The state is stored in the instance memory in a dictionary accessed
via the Clojure transactional memory system.  (This means it does not
scale well to multiple servers -- See Improvements section below for
how we might fix this.)

### Client

The main client code is `src/cljs/fgosite.client.gos`, and is an
example of seamlessly mixing Clojure and JavaScript technology.  It
uses the Clojure Hiccups library to do client-side templating, and the
JavaScript JQuery library for everything else.

The styling is done using Boostrap.

The client code does the following:

1. Uses a template to build HTML from Functional Go lists and
dictionaries.
2. Gets some sample code from the server and displays it, allowing the
user to modify it.
3. Using the MD5 hash of the code as an ID, gets the corresponding
Clojure and evaluated results from the server.
4. If the MD5 ID returned no result then upload the code to the server
and retry step 3.

## Improvements

Currently if we scale to multiple load-balanced servers the client
code will still work, but it may go through the retry loop multiple
times until it eventually does a GET from one of the servers it
previously did a PUT from.  This would get worse as the number of
servers increases.

A more scalable solution is to store the state in a Memcache server.
We will still occasionally have to go around the retry loop because
Memcache is not durable and items may be evicted from memory, but it
should happen quite rarely.

[1]: http://www.funcgo.org/
