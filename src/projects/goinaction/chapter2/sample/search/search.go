package search

import (
	"log"
	"sync"
)

// 包内访问权限，
// 以大写字母开头的变量从包中公开出去，小写字母开头则不公开
// 当方法返回一个不公开的类型的变量值，接收者可以使用该值
// map 需要使用make来构造
var matchers = make(map[string]Matcher)

// var 和 :=
// 如果需要声明初始值为0值的变量，使用var
// 提供了确切的非0值或函数返回值创建的变量，使用:=
func Run(searchTerm string) {
	// := 声明，并赋值
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// Create an unbuffered channel to receive match results to display.
	results := make(chan *Result)

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds.
	waitGroup.Add(len(feeds))

	// _ 下划线表示占位符，如果函数返回多个值, 可以使用下划线忽略掉
	for _, feed := range feeds {
		// 查找map里的key时，可以返回一个值，也可以返回两个值
		// 一个值时返回的时value
		// 两个值时返回的时value 和 该key是否存在
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 匿名函数，支持闭包
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// Launch a goroutine to monitor when all the work is done.
	go func() {
		// Wait for everything to be processed.
		waitGroup.Wait()

		// Close the channel to signal to the Display
		// function that we can exit the program.
		close(results)
	}()

	// Start displaying results as they are available and
	// return after the final result is displayed.
	Display(results)
}

// Register is called to register a matcher for use by the program.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
