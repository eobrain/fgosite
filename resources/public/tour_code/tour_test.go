package tour_test

import (
	test "midje/sweet"
	m1 "tour1/main"
	m3 "tour3/main"
	m4 "tour4/main"
	m6 "tour6/main"
	m7 "tour7/main"
	m8 "tour8/main"
	m9 "tour9/main"
	m10 "tour10/main"
	m15 "tour15/main"
	m16 "tour16/main"
	m17 "tour17/main"
	m23 "tour23/main"
	m25 "tour25/main"
	m26 "tour26/main"
	m31 "tour31/main"
	m40 "tour40/main"
	m47 "tour47/main"
	// m52 "tour52/main"
)


test.fact("tour1", withOutStr(m1.main()), =>, `Hello, 世界
`)
test.fact("tour3", withOutStr(m3.main()), =>, /Welcome to the playground!\nThe time is .*\nAnd if you try to open a file:\n.*No such file or directory.*\nOr access the network:\n.*\n/)
test.fact("tour4", withOutStr(m4.main()), =>, /My favorite number is [0-9]\n/)
test.fact("tour6", withOutStr(m6.main()), =>, `3.141592653589793
`)
test.fact("tour7", withOutStr(m7.main()), =>, `55
`)
test.fact("tour8", withOutStr(m8.main()), =>, `55
`)
test.fact("tour9", withOutStr(m9.main()), =>, `world hello
`)
test.fact("tour10", withOutStr(m10.main()), =>, `[7 10]
`)
test.fact("tour15", withOutStr(m15.main()), =>, `3 4 5
`)
test.fact("tour16", withOutStr(m16.main()), =>, `Hello 世界
Happy 3.14 Day
Go rules? true
`)
test.fact("tour17", withOutStr(m17.main()), =>, `21
0.2
`)
test.fact("tour23", withOutStr(m23.main()), =>, `9 20
`)
test.fact("tour25", withOutStr(m25.main()), =>, `10.000000000000007
`)
test.fact("tour26", withOutStr(m26.main()), =>, `{1 2}
`)
test.fact("tour31", withOutStr(m31.main()), =>, `Hello World
[Hello World]
`)
test.fact("tour40", withOutStr(m40.main()), =>, `map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
`)
test.fact("tour47", withOutStr(m47.main()), =>, `Go runs on clojure.`)

