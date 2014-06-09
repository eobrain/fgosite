(defproject fgosite "0.1.1"
  :description "FIXME: write this!"
  :url "http://example.com/FIXME"
  :dependencies [[org.clojure/clojure "1.5.1"]
                 [org.eamonn.funcgo/funcgo-lein-plugin "0.2.4-SNAPSHOT"]
                 [org.eamonn.funcgo/fgolib "0.2.2-SNAPSHOT"]
                 [org.clojure/clojurescript "0.0-2156"]
                 [ring "1.2.1"]
                 [compojure "1.1.8"]
                 [hiccups "0.3.0"]
                 [midje "1.5.1" :scope "test"]]
  :test-paths ["resources/public/tour_code"]
  :plugins [[org.eamonn.funcgo/funcgo-lein-plugin "0.2.4-SNAPSHOT"]
            [lein-cljsbuild "1.0.2"]
            [lein-ring "0.8.10"]
            [lein-midje "3.1.1"]]
  :source-paths ["src/clj"]
  :cljsbuild {
    :builds {
      :main {
             :source-paths ["src/cljs"]
        :compiler {:output-to "resources/public/js/cljs.js"
                   :optimizations :simple
                   :pretty-print true}
        :jar true}}}
  :main fgosite.server
  :ring {:handler fgosite.server/app})
