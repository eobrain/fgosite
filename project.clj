(defproject fgosite "0.1.7"
  :description "Web site for Funcgo, written in Funcgo"
  :url "https://github.com/eobrain/funcgo"
  :dependencies [[org.clojure/clojure "1.6.0"]
                 [org.clojure/core.async "0.1.303.0-886421-alpha"]
                 [org.eamonn.funcgo/funcgo-lein-plugin "0.5.1"]
                 [org.eamonn.funcgo/fgolib "0.2.6"]
                 [org.clojure/clojurescript "0.0-2156"]
                 [ring "1.2.1"]
                 [compojure "1.1.8"]
                 [hiccup "1.0.5"]
                 [hiccups "0.3.0"]
                 [javax.cache/cache-api "1.0.0-RC1" :scope "test"]
                 ;; [net.sf.jsr107cache/jsr107cache "1.0" :scope "test"]
                 [midje "1.6.3"                     :scope "test"]]
  :min-lein-version "2.0.0"
  :test-paths ["resources/public/tour_code"]
  :plugins [[org.eamonn.funcgo/funcgo-lein-plugin "0.5.1"]
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
