package keyword

import "testing"

func TestTdIdf(t *testing.T) {
	var doc1 = []string{"人工", "智能", "成为", "互联网", "大会", "焦点"}
	var doc2 = []string{"谷歌", "推出", "开源", "人工", "智能", "系统", "工具"}
	var doc3 = []string{"互联网", "的", "未来", "在", "人工", "智能"}
	var doc4 = []string{"谷歌", "开源", "机器", "学习", "工具"}
	var docs = [][]string{doc1, doc2, doc3, doc4}
	t.Log(td(doc2, "谷歌"))          //词频
	t.Log(df(docs, "谷歌"))          //文档频率
	t.Log(idf(docs, "谷歌"))         //逆文档频率
	t.Log(tdidf(doc2, docs, "谷歌")) //权重
}
