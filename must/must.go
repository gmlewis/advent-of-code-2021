// Package must simplifies writing advent-of-code scripts by providing
// methods that cause fatal errors instead of returning them.
package must

import "log"

// fatal is used for unit testing.
var fatal = log.Fatal
