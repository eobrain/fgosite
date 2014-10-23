package memcache
//import type com.google.appengine.api.memcache.MemcacheServiceFactory
import type net.sf.jsr107cache.CacheManager

//cache := MemcacheServiceFactory::getMemcacheService()
cacheFactory := CacheManager::getInstance()->getCacheFactory()
cache := cacheFactory->createCache({})



// Return a memoized version of a function of one argument
func Memoized(f) {
	func(key) {
		cached := cache->get(key)
		if cached == null {
			newValue := f(key)
			cache->put(key, newValue)
			newValue
		} else {
			cached
		}
	}
}
