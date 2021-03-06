# For deployment

WAR=target/fgosite-$(VERSION)-standalone.war

AE_VERSION=`echo $(VERSION) | tr .A-Z -a-z`

appengine-local: $(WAR)
	: Run on App Engine local development server
	unzip -q $(WAR) -d /tmp/exploded$$$$; dev_appserver.sh /tmp/exploded$$$$

appengine: $(WAR)
	: Deploy to App Engine
	unzip -q $(WAR) -d /tmp/exploded$$$$; appcfg.sh --oauth2 --version=$(AE_VERSION) update /tmp/exploded$$$$

heroku: $(WAR)
	: Deploy on Heroku
	heroku deploy:war --war $(WAR)

$(WAR):
	lein do  clean,  fgoc --force,  cljsbuild once,  ring uberwar

local:
	: Run locally
	lein do  clean,  fgoc --force,  cljsbuild once,  ring server-headless
