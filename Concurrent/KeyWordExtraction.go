package Concurrent

/**
抽取关键字
*/

type word struct {
	word string
	tf   int //术语频次，是指一个单词在一个文档中出现的次数
	df   int //文档频次，是含有某个单词的文档的数量 idf,逆文档频次，用于度量单词所提供的使某个文档区别于其他文档的信息。如果一个单词很常用，它的idf值就低。但是如果该单词
	//仅在少数几个文档中出现，那么它的idf值就很高
	tfIdf float64 //tf*idf=F(td)*log(N/n)  F(td) 是单词t在文档d中出现的次数，N是集合中文档的树目，n是含有单词t的文档的树目
}
