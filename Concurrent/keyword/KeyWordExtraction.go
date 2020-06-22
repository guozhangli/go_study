package keyword

import (
	"errors"
	"fmt"
	"go_study/Concurrent/textindexing"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

/**
抽取关键字
*/

type word struct {
	word string
	tf   int //术语频次，(term frequency)是指一个单词在一个文档中出现的次数
	df   int //文档频次，(document frequency)是含有某个单词的文档的数量 idf,逆文档频次，用于度量单词所提供的使某个文档区别于其他文档的信息。如果一个单词很常用，它的idf值就低。但是如果该单词
	//仅在少数几个文档中出现，那么它的idf值就很高
	tfIdf float64 //tf*idf=F(td)*log(N/n)  F(td) 是单词t在文档d中出现的次数，N是集合中文档的树目，n是含有单词t的文档的树目
}

func NewWord(w string) *word {
	return &word{
		word: w,
		df:   1,
	}
}

func (w *word) setTfIdf(dNum int) {
	w.tfIdf = float64(w.tf) * math.Log(float64(dNum/w.df))
}

func (w *word) addTf() *word {
	w.tf++
	return w
}

func (w *word) marge(other *word) {
	if w.word == other.word {
		w.df += other.df
		w.tf += other.tf
	}
}

type wordList []*word

func (w wordList) Len() int {
	return len(w)
}
func (w wordList) Less(i, j int) bool {
	return w[i].tfIdf > w[j].tfIdf
}

func (w wordList) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

type keyword struct {
	word string
	df   int
}

type keywordList []*keyword

func (kl keywordList) Len() int {
	return len(kl)
}

func (kl keywordList) Less(i, j int) bool {
	return kl[i].df > kl[j].df
}

func (kl keywordList) Swap(i, j int) {
	kl[i], kl[j] = kl[j], kl[i]
}

func NewKeyword(word string, df int) *keyword {
	return &keyword{
		word: word,
		df:   df,
	}
}

type document struct {
	fileName string
	voc      map[string]*word
}

func NewDocument(fileName string) *document {
	return &document{
		fileName: fileName,
		voc:      make(map[string]*word),
	}
}

func (d *document) addWord(key string) {
	if _, ok := d.voc[key]; !ok {
		d.voc[key] = NewWord(key)
	} else {
		d.voc[key] = d.voc[key].addTf()
	}
}

func KeywordSerial() {
	files := textindexing.ReadDir()
	var globalVoc = make(map[string]*word)
	var globalKeywords = make(map[string]int)
	var numDocuments int
	var orderedGlobalKeywords []*keyword
	start := time.Now().UnixNano()
	// Phase 1: Parse all the documents
	phase1(&files, &globalVoc, &numDocuments)
	// Phase 2: Update the df of the voc of the Documents
	phase2(&files, &globalVoc, &globalKeywords, &numDocuments)
	// Phase 3: Get a list of a better keywords
	phase3(&orderedGlobalKeywords, &globalKeywords)

	end := time.Now().UnixNano()
	fmt.Printf("Execution Time: %d\n", (end-start)/1000000)
	fmt.Printf("Vocabulary Size: %d\n", len(globalVoc))
	fmt.Printf("Keyword Size: %d\n", len(globalKeywords))
	fmt.Printf("Number of Documents: %d\n", numDocuments)
}

var reg = regexp.MustCompile(`[^[\x00-\x7F]]`)
var regsplit = regexp.MustCompile(`\W+`)

func parse(f *os.FileInfo) (*document, error) {
	b, err := ioutil.ReadFile(filepath.Join(textindexing.PATH, (*f).Name()))
	if err != nil {
		return nil, errors.New("read file error")
	}
	doc := NewDocument((*f).Name())
	lines := strings.Split(string(b), "\r\n")
	for _, line := range lines {
		parseLine(line, doc)
	}
	return doc, nil
}

func parseLine(line string, doc *document) {
	line = reg.ReplaceAllString(line, "")
	line = strings.ToLower(line)

	for _, l := range regsplit.Split(line, -1) {
		if l != "" {
			doc.addWord(l)
		}
	}
}

func phase1(files *[]os.FileInfo, globalVoc *map[string]*word, numDocuments *int) {
	for _, f := range *files {
		if strings.HasSuffix(f.Name(), ".txt") {
			document, err := parse(&f)
			if err != nil {
				continue
			}
			for _, v := range document.voc {
				if _, ok := (*globalVoc)[v.word]; !ok {
					(*globalVoc)[v.word] = v
				} else {
					(*globalVoc)[v.word].marge(v)
				}
			}
			*numDocuments++
		}
	}
}

func phase2(files *[]os.FileInfo, globalVoc *map[string]*word, globalKeywords *map[string]int, numDocuments *int) {
	for _, f := range *files {
		if strings.HasSuffix(f.Name(), ".txt") {
			document, err := parse(&f)
			if err != nil {
				continue
			}
			var keywords []*word
			for _, v := range document.voc {
				w := (*globalVoc)[v.word]
				w.setTfIdf(*numDocuments)
				keywords = append(keywords, w)
			}
			sort.Sort(wordList(keywords))
			if len(keywords) > 10 {
				keywords = keywords[:10]
			}
			for _, v := range keywords {
				if _, ok := (*globalKeywords)[v.word]; !ok {
					(*globalKeywords)[v.word] = 1
				} else {
					(*globalKeywords)[v.word] += 1
				}
			}
		}
	}
}
func phase3(orderedGlobalKeywords *[]*keyword, globalKeywords *map[string]int) {
	for k, v := range *globalKeywords {
		kw := NewKeyword(k, v)
		*orderedGlobalKeywords = append(*orderedGlobalKeywords, kw)
	}
	sort.Sort(keywordList(*orderedGlobalKeywords))
	*orderedGlobalKeywords = (*orderedGlobalKeywords)[0:100]
	for _, v := range *orderedGlobalKeywords {
		fmt.Printf("%s:%d\n", v.word, v.df)
	}
}
