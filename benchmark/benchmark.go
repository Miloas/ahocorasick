package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/anknown/ahocorasick"
	"github.com/cloudflare/ahocorasick"
	"github.com/iohub/Ahocorasick"
)

const CHN_DICT_FILE = "./cn/dictionary.txt"
const CHN_TEXT_FILE = "./cn/text.txt"
const ENG_DICT_FILE = "./en/dictionary.txt"
const ENG_TEXT_FILE = "./en/text.txt"

func ReadBytes(filename string) ([][]byte, error) {
	dict := [][]byte{}

	f, err := os.OpenFile(filename, os.O_RDONLY, 0660)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(f)
	for {
		l, err := r.ReadBytes('\n')
		if err != nil || err == io.EOF {
			break
		}
		l = bytes.TrimSpace(l)
		dict = append(dict, l)
	}

	return dict, nil
}

func ReadRunes(filename string) ([][]rune, error) {
	dict := [][]rune{}

	f, err := os.OpenFile(filename, os.O_RDONLY, 0660)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(f)
	for {
		l, err := r.ReadBytes('\n')
		if err != nil || err == io.EOF {
			break
		}
		l = bytes.TrimSpace(l)
		dict = append(dict, bytes.Runes(l))
	}

	return dict, nil
}

func TestAEnglish() {
	fmt.Println("** English Benchmark of cloudflare/ahocorasick **")
	fmt.Println("-------------------------------------------------")
	fmt.Println("=> Start to Load... ")
	start := time.Now()
	dict, err := ReadBytes(ENG_DICT_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := ioutil.ReadFile(ENG_TEXT_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}
	end := time.Now()
	fmt.Printf("load file cost:%d(ms)\n", (end.UnixNano()-start.UnixNano())/(1000*1000))

	fmt.Println("=> Start to Search... ")
	start = time.Now()
	m := ahocorasick.NewMatcher(dict)

	//res := m.Match(content)
	m.Match(content)
	end = time.Now()

	fmt.Printf("search cost:%d(ms)\n", (end.UnixNano()-start.UnixNano())/(1000*1000))
}

func TestAChinese() {
	fmt.Println("\n** Chinese Benchmark of cloudflare/ahocorasick **")
	fmt.Println("---------------------------------------------------")
	fmt.Println("=> Start to Load... ")
	start := time.Now()
	dict, err := ReadBytes(CHN_DICT_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := ioutil.ReadFile(CHN_TEXT_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}
	end := time.Now()
	fmt.Printf("load file cost:%d(ms)\n", (end.UnixNano()-start.UnixNano())/(1000*1000))

	fmt.Println("=> Start to Search... ")
	start = time.Now()
	m := ahocorasick.NewMatcher(dict)

	//res := m.Match(content)
	m.Match(content)
	end = time.Now()

	fmt.Printf("search cost:%d(ms)\n", (end.UnixNano()-start.UnixNano())/(1000*1000))
}

func TestBEnglish() {
	fmt.Println("\n** English Benchmark of anknown/ahocorasick **")
	fmt.Println("------------------------------------------------")
	fmt.Println("=> Start to Load... ")
	start := time.Now()
	dict, err := ReadRunes(ENG_DICT_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := ioutil.ReadFile(ENG_TEXT_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	contentRune := bytes.Runes([]byte(content))
	end := time.Now()
	fmt.Printf("load file cost:%d(ms)\n", (end.UnixNano()-start.UnixNano())/(1000*1000))

	fmt.Println("=> Start to Search... ")
	start = time.Now()
	m := new(goahocorasick.Machine)
	if err := m.Build(dict); err != nil {
		fmt.Println(err)
		return
	}
	//terms := m.Search(contentRune)
	m.MultiPatternSearch(contentRune, false)
	end = time.Now()
	fmt.Printf("search cost:%d(ms)\n", (end.UnixNano()-start.UnixNano())/(1000*1000))

}

func TestBChinese() {
	fmt.Println("\n** Chinese Benchmark of anknown/ahocorasick **")
	fmt.Println("------------------------------------------------")
	fmt.Println("=> Start to Load... ")
	start := time.Now()
	dict, err := ReadRunes(CHN_DICT_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := ioutil.ReadFile(CHN_TEXT_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	contentRune := bytes.Runes([]byte(content))
	end := time.Now()
	fmt.Printf("load file cost:%d(ms)\n", (end.UnixNano()-start.UnixNano())/(1000*1000))

	fmt.Println("=> Start to Search... ")
	start = time.Now()
	m := new(goahocorasick.Machine)
	if err := m.Build(dict); err != nil {
		fmt.Println(err)
		return
	}
	//terms := m.Search(contentRune)
	m.MultiPatternSearch(contentRune, false)
	end = time.Now()
	fmt.Printf("search cost:%d(ms)\n", (end.UnixNano()-start.UnixNano())/(1000*1000))

}

func TestCEnglish() {
	fmt.Println("** English Benchmark of iohub/Ahocorasick **")
	fmt.Println("-------------------------------------------------")
	fmt.Println("=> Start to Load... ")
	start := time.Now()
	dict, err := ReadBytes(ENG_DICT_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := ioutil.ReadFile(ENG_TEXT_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}
	end := time.Now()
	fmt.Printf("load file cost:%d(ms)\n", (end.UnixNano()-start.UnixNano())/(1000*1000))

	fmt.Println("=> Start to Search... ")
	start = time.Now()
	m := cedar.NewMatcher()
	for i, b := range dict {
		m.Insert(b, i)
	}
	m.Compile()
	m.Match(content)
	end = time.Now()

	fmt.Printf("search cost:%d(ms)\n", (end.UnixNano()-start.UnixNano())/(1000*1000))
}

func TestCChinese() {
	fmt.Println("\n** Chinese Benchmark of iohub/ahocorasick **")
	fmt.Println("------------------------------------------------")
	fmt.Println("=> Start to Load... ")
	start := time.Now()
	dict, err := ReadBytes(CHN_DICT_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := ioutil.ReadFile(CHN_TEXT_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	// contentRune := bytes.Runes([]byte(content))
	end := time.Now()
	fmt.Printf("load file cost:%d(ms)\n", (end.UnixNano()-start.UnixNano())/(1000*1000))

	fmt.Println("=> Start to Search... ")
	start = time.Now()
	m := cedar.NewMatcher()
	for i, b := range dict {
		m.Insert(b, i)
	}
	m.Compile()
	m.Match(content)
	end = time.Now()
	fmt.Printf("search cost:%d(ms)\n", (end.UnixNano()-start.UnixNano())/(1000*1000))

}

func main() {
	/*
		f, err := os.Create("benchmark.bin")
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	*/
	TestAEnglish()
	TestBEnglish()
	TestCEnglish()
	TestAChinese()
	TestBChinese()
	TestCChinese()
}
