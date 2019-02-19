package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

/*
	7.6. sort.Interface接口

	Go语言的sort.Sort函数不会对具体的序列和它的元素做任何假设。相反，它使用了一个接口类型
	sort.Interface来指定通用的排序算法和可能被排序到的序列类型之间的约定。这个接口的实现
	由序列的具体表示和它希望排序的元素决定，序列的表示经常是一个切片。

	type Interface interface {
		// 一个内置的排序算法需要知道三个东
		Len() int	// 序列的长度
		Less(i, j int) bool // 表示两个元素比较的结果
		Swap(i, j int) // 一种交换两个元素的方式
	}
*/

// demo01
// StringSlice 类型的底层数据结构是切片
type StringSlice []string

// StringSlice 类型实现了sort.Interface 接口
func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// demo02

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

/*
	下面的变量tracks包含了一个播放列表。每个元素都不是Track本身而是指向它的指针。
	尽管我们在下面的代码中直接存储Tracks也可以工作，sort函数会交换很多对元素，所
	以如果每个元素都是指针而不是Track类型会更快，指针是一个机器字码长度而Track类型可能是八个或更多。
*/
var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

/*
	printTracks函数将播放列表打印成一个表格。
*/
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	// *tabwriter.Writer是满足io.Writer接口的。它会收集每一片写向它的数据；它的Flush方法会格式化整个表格并且将它写向os.Stdout
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

/*
	按照字段排序
*/
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//
/*
	demo03 多字段排序
*/
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	// demo01 -----------------------------------------------------------------------------
	names := [...]string{"justin", "emma", "mark", "lucas", "harry"}
	sort.Sort(StringSlice(names[0:]))

	for _, name := range names {
		fmt.Println(name)
	}

	// sort包 提供的排序
	cities := []string{"beijing", "shanghai", "guangzhou", "jiujiang", "fengcheng"}
	sort.Strings(cities)

	for _, city := range cities {
		fmt.Println(city)
	}
	fmt.Println()

	// demo02 -----------------------------------------------------------------------------
	sort.Sort(byArtist(tracks))
	printTracks(tracks)
	// sort包中提供了Reverse函数将排序顺序转换成逆序。sort.Reverse函数值得进行更近一步的学习，
	// 因为它使用了（§6.3）章中的组合，这是一个重要的思路。
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)
	fmt.Println()

	// demo03 -----------------------------------------------------------------------------
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)
	fmt.Println()

	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values)) // "false"
	sort.Ints(values)
	fmt.Println(values)                     // "[1 1 3 4]"
	fmt.Println(sort.IntsAreSorted(values)) // "true"
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)                     // "[4 3 1 1]"
	fmt.Println(sort.IntsAreSorted(values)) // "false"
}
