package keyword

import "testing"

/*
   	ml.put("Hello", 1.0);
       ml.put("css", 2.0);
       ml.put("Lucene", 3.0);
  m2.put("Hello", 1.0);
        m2.put("Word", 2.0);
        m2.put("Hadoop",3.0);
        m2.put("java", 4.0);
        m2.put("html", 1.0);
        m2.put("css", 2.0);
*/
func TestVsm(t *testing.T) {
	m1 := make(map[string]float64)
	m1["Hello"] = 1.0
	m1["css"] = 2.0
	m1["Lucene"] = 3.0
	m2 := make(map[string]float64)
	m2["Hello"] = 1.0
	m2["Word"] = 2.0
	m2["Hadoop"] = 3.0
	m2["java"] = 4.0
	m2["html"] = 1.0
	m2["css"] = 2.0
	calCosSim(m1, m2)
}
