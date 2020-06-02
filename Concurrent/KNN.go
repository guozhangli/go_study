package Concurrent

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

type BankMarketing struct {
	age                         float64
	jobAdmin                    float64
	jobBlueCollar               float64
	jobEntrepreneur             float64
	jobHousemaid                float64
	jobManagement               float64
	jobRetired                  float64
	jobSelfEmployed             float64
	jobServices                 float64
	jobStudent                  float64
	jobTechnician               float64
	jobUnemployed               float64
	jobUnknown                  float64
	maritalDivorced             float64
	maritalMarried              float64
	maritalSingle               float64
	maritalUnknown              float64
	educationBasic4y            float64
	educationBasic6y            float64
	educationBasic9y            float64
	educationHighSchool         float64
	educationIlliterate         float64
	educationProfessionalCourse float64
	educationUniversityDegree   float64
	educationUnknown            float64
	creditNo                    float64
	creditYes                   float64
	creditUnknown               float64
	housingNo                   float64
	housingYes                  float64
	housingUnknown              float64
	loanNo                      float64
	loanYes                     float64
	loanUnknown                 float64
	contactCellular             float64
	contactTelephone            float64
	contactJan                  float64
	contactFeb                  float64
	contactMar                  float64
	contactApr                  float64
	contactMay                  float64
	contactJun                  float64
	contactJul                  float64
	contactAug                  float64
	contactSep                  float64
	contactOct                  float64
	contactNov                  float64
	contactDec                  float64
	contactMon                  float64
	contactTue                  float64
	contactWed                  float64
	contactThu                  float64
	contactFri                  float64
	duration                    int
	campaign                    float64
	pdays                       int
	pdaysNever                  float64
	previous                    float64
	poutcomeFailure             float64
	poutcomeNonexistent         float64
	poutcomeSuccess             float64
	empVarRate                  float64
	consPriceIdx                float64
	consConfIdx                 float64
	euribor3m                   float64
	nrEmployed                  float64
	target                      string
}

type Distance struct {
	index    int
	distance float64
}

func NewBankMarketing() *BankMarketing {
	return new(BankMarketing)
}

func (b *BankMarketing) setData(data []string) {
	if len(data) != 67 {
		panic("Wrong data length: " + string(len(data)))
	}
	b.age = bytes2Float64(data[0])
	b.jobAdmin = bytes2Float64(data[1])
	b.jobBlueCollar = bytes2Float64(data[2])
	b.jobEntrepreneur = bytes2Float64(data[3])
	b.jobHousemaid = bytes2Float64(data[4])
	b.jobManagement = bytes2Float64(data[5])
	b.jobRetired = bytes2Float64(data[6])
	b.jobSelfEmployed = bytes2Float64(data[7])
	b.jobServices = bytes2Float64(data[8])
	b.jobStudent = bytes2Float64(data[9])
	b.jobTechnician = bytes2Float64(data[10])
	b.jobUnemployed = bytes2Float64(data[11])
	b.jobUnknown = bytes2Float64(data[12])
	b.maritalDivorced = bytes2Float64(data[13])
	b.maritalMarried = bytes2Float64(data[14])
	b.maritalSingle = bytes2Float64(data[15])
	b.maritalUnknown = bytes2Float64(data[16])
	b.educationBasic4y = bytes2Float64(data[17])
	b.educationBasic6y = bytes2Float64(data[18])
	b.educationBasic9y = bytes2Float64(data[19])
	b.educationHighSchool = bytes2Float64(data[20])
	b.educationIlliterate = bytes2Float64(data[21])
	b.educationProfessionalCourse = bytes2Float64(data[22])
	b.educationUniversityDegree = bytes2Float64(data[23])
	b.educationUnknown = bytes2Float64(data[24])
	b.creditNo = bytes2Float64(data[25])
	b.creditYes = bytes2Float64(data[26])
	b.creditUnknown = bytes2Float64(data[27])
	b.housingNo = bytes2Float64(data[28])
	b.housingYes = bytes2Float64(data[29])
	b.housingUnknown = bytes2Float64(data[30])
	b.loanNo = bytes2Float64(data[31])
	b.loanYes = bytes2Float64(data[32])
	b.loanUnknown = bytes2Float64(data[33])
	b.contactCellular = bytes2Float64(data[34])
	b.contactTelephone = bytes2Float64(data[35])
	b.contactJan = bytes2Float64(data[36])
	b.contactFeb = bytes2Float64(data[37])
	b.contactMar = bytes2Float64(data[38])
	b.contactApr = bytes2Float64(data[39])
	b.contactMay = bytes2Float64(data[40])
	b.contactJun = bytes2Float64(data[41])
	b.contactJul = bytes2Float64(data[42])
	b.contactAug = bytes2Float64(data[43])
	b.contactSep = bytes2Float64(data[44])
	b.contactOct = bytes2Float64(data[45])
	b.contactNov = bytes2Float64(data[46])
	b.contactDec = bytes2Float64(data[47])
	b.contactMon = bytes2Float64(data[48])
	b.contactTue = bytes2Float64(data[49])
	b.contactWed = bytes2Float64(data[50])
	b.contactThu = bytes2Float64(data[51])
	b.contactFri = bytes2Float64(data[52])
	b.duration, _ = strconv.Atoi(data[53])
	b.campaign = bytes2Float64(data[54])
	b.pdays, _ = strconv.Atoi(data[55])
	b.pdaysNever = bytes2Float64(data[56])
	b.previous = bytes2Float64(data[57])
	b.poutcomeFailure = bytes2Float64(data[58])
	b.poutcomeNonexistent = bytes2Float64(data[59])
	b.poutcomeSuccess = bytes2Float64(data[60])
	b.empVarRate, _ = strconv.ParseFloat(data[61], 64)
	b.consPriceIdx, _ = strconv.ParseFloat(data[62], 64)
	b.consConfIdx, _ = strconv.ParseFloat(data[63], 64)
	b.euribor3m, _ = strconv.ParseFloat(data[64], 64)
	b.nrEmployed, _ = strconv.ParseFloat(data[65], 64)
	b.target = data[66]
}

func bytes2Float64(bytes string) float64 {
	f, _ := strconv.ParseFloat(bytes, 64)
	return f
}

func (b *BankMarketing) getData() []float64 {
	return []float64{
		b.age,
		b.jobAdmin,
		b.jobBlueCollar,
		b.jobEntrepreneur,
		b.jobHousemaid,
		b.jobManagement,
		b.jobRetired,
		b.jobSelfEmployed,
		b.jobServices,
		b.jobStudent,
		b.jobTechnician,
		b.jobUnemployed,
		b.jobUnknown,
		b.maritalDivorced,
		b.maritalMarried,
		b.maritalSingle,
		b.maritalUnknown,
		b.educationBasic4y,
		b.educationBasic6y,
		b.educationBasic9y,
		b.educationHighSchool,
		b.educationIlliterate,
		b.educationProfessionalCourse,
		b.educationUniversityDegree,
		b.educationUnknown,
		b.creditNo,
		b.creditYes,
		b.creditUnknown,
		b.housingNo,
		b.housingYes,
		b.housingUnknown,
		b.loanNo,
		b.loanYes,
		b.loanUnknown,
		b.contactCellular,
		b.contactTelephone,
		b.contactJan,
		b.contactFeb,
		b.contactMar,
		b.contactApr,
		b.contactMay,
		b.contactJun,
		b.contactJul,
		b.contactAug,
		b.contactSep,
		b.contactOct,
		b.contactNov,
		b.contactDec,
		b.contactMon,
		b.contactTue,
		b.contactWed,
		b.contactThu,
		b.contactFri,
		float64(b.duration),
		b.campaign,
		float64(b.pdays),
		b.pdaysNever,
		b.previous,
		b.poutcomeFailure,
		b.poutcomeNonexistent,
		b.poutcomeSuccess,
		b.empVarRate,
		b.consPriceIdx,
		b.consConfIdx,
		b.euribor3m,
		b.nrEmployed,
	}
}

func Load(filePath string) ([]*BankMarketing, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("read file error")
	}
	datas := strings.Split(string(data), "\r\n")
	len := len(datas)
	list := make([]*BankMarketing, len-1)
	for i, line := range datas[:len-1] {
		ds := strings.Split(line, ";")
		bankMarketing := NewBankMarketing()
		bankMarketing.setData(ds)
		list[i] = bankMarketing
	}
	return list, nil
}

func NewDistance() *Distance {
	return &Distance{
		index:    0,
		distance: 0,
	}
}

func calculate(data, test *BankMarketing) float64 {
	dataF64 := data.getData()
	testF64 := test.getData()
	len1 := len(dataF64)
	len2 := len(testF64)
	if len1 != len2 {
		panic("Vector doesn't have the same length")
	}
	var res float64
	for i := 0; i < len1; i++ {
		res += math.Pow(dataF64[i]-testF64[i], 2)
	}
	return math.Sqrt(res)
}

//k-最近邻算法串行版本
func KnnSerial() {
	train, _ := Load("data/bank.data")
	fmt.Printf("Train: %d\n", len(train))
	test, _ := Load("data/bank.test")
	fmt.Printf("Test: %d\n", len(test))
	k := 10
	start := time.Now().UnixNano()
	var success, mistakes int
	len := len(test)
	for i := 0; i < len; i++ {
		t := test[i]
		tag := KnnClassifier(train, t, k)
		if tag == t.target {
			success++
		} else {
			mistakes++
		}
	}
	end := time.Now().UnixNano()
	fmt.Println("******************************************")
	fmt.Printf("Serial Classifier - K: %d\n", k)
	fmt.Printf("Success: %d\n", success)
	fmt.Printf("Mistakes: %d\n", mistakes)
	fmt.Printf("Execution Time: %d seconds.\n", ((end - start) / 1000000000))
	fmt.Println("******************************************")
}

func KnnClassifier(dataset []*BankMarketing, bankMarketing *BankMarketing, k int) string {
	len := len(dataset)
	result := make([]*Distance, len)
	for i := 0; i < len; i++ {
		distance := NewDistance()
		distance.index = i
		distance.distance = calculate(dataset[i], bankMarketing)
		result[i] = distance
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].distance < result[j].distance
	})
	m := make(map[string]int)
	for _, v := range result[:k] {
		v1 := dataset[v.index].target
		if _, ok := m[v1]; !ok {
			m[v1] = 1
		} else {
			m[v1]++
		}
	}
	var t string
	var mx int
	for tag, max := range m {
		if max > mx {
			t = tag
			mx = max
		}
	}
	return t
}

//k-邻近算法并行版本（细粒度的并行）
func KnnParallelIndividual() {
	train, _ := Load("data/bank.data")
	fmt.Printf("Train: %d\n", len(train))
	test, _ := Load("data/bank.test")
	fmt.Printf("Test: %d\n", len(test))
	k := 10
	start := time.Now().UnixNano()
	ch := make(chan struct{}, len(train))
	var success, mistakes int
	len := len(test)
	for i := 0; i < len; i++ {
		t := test[i]
		tag := KnnClassifierParallel(train, t, k, ch)
		if tag == t.target {
			success++
		} else {
			mistakes++
		}
	}
	end := time.Now().UnixNano()
	fmt.Println("******************************************")
	fmt.Printf("Serial Classifier - K: %d\n", k)
	fmt.Printf("Success: %d\n", success)
	fmt.Printf("Mistakes: %d\n", mistakes)
	fmt.Printf("Execution Time: %d seconds.\n", ((end - start) / 1000000000))
	fmt.Println("******************************************")
}

func KnnClassifierParallel(dataset []*BankMarketing, bankMarketing *BankMarketing, k int, ch chan struct{}) string {
	len := len(dataset)
	result := make([]*Distance, len)
	for i := 0; i < len; i++ {
		go func(dataBM, testBM *BankMarketing, i int, ch chan struct{}) {
			distance := NewDistance()
			distance.index = i
			distance.distance = calculate(dataset[i], bankMarketing)
			result[i] = distance
			ch <- struct{}{}
		}(dataset[i], bankMarketing, i, ch)
	}
	for i := 0; i < len; i++ {
		<-ch
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].distance < result[j].distance
	})
	m := make(map[string]int)
	for _, v := range result[:k] {
		v1 := dataset[v.index].target
		if _, ok := m[v1]; !ok {
			m[v1] = 1
		} else {
			m[v1]++
		}
	}
	var t string
	var mx int
	for tag, max := range m {
		if max > mx {
			t = tag
			mx = max
		}
	}
	return t
}

//k-邻近算法并行版本（粗粒度的并行）
func KnnParallelGroup() {
	train, _ := Load("data/bank.data")
	fmt.Printf("Train: %d\n", len(train))
	test, _ := Load("data/bank.test")
	fmt.Printf("Test: %d\n", len(test))
	k := 10
	start := time.Now().UnixNano()
	ch := make(chan struct{}, len(train))
	var success, mistakes int
	len := len(test)
	for i := 0; i < len; i++ {
		t := test[i]
		tag := KnnClassifierParallelGroup(train, t, k, ch)
		if tag == t.target {
			success++
		} else {
			mistakes++
		}
	}
	end := time.Now().UnixNano()
	fmt.Println("******************************************")
	fmt.Printf("Serial Classifier - K: %d\n", k)
	fmt.Printf("Success: %d\n", success)
	fmt.Printf("Mistakes: %d\n", mistakes)
	fmt.Printf("Execution Time: %d seconds.\n", ((end - start) / 1000000000))
	fmt.Println("******************************************")
}

func KnnClassifierParallelGroup(dataset []*BankMarketing, bankMarketing *BankMarketing, k int, ch chan struct{}) string {
	len := len(dataset)
	result := make([]*Distance, len)
	var numThread = 10
	var startIndex = 0
	var step = len / numThread
	var endIndex = step
	for i := 0; i < numThread; i++ {
		go func(startIndex, endIndex int, dataset []*BankMarketing, bankMarketing *BankMarketing, ch chan struct{}) {
			for i := startIndex; i < endIndex; i++ {
				distance := NewDistance()
				distance.index = i
				distance.distance = calculate(dataset[i], bankMarketing)
				result[i] = distance
			}
			ch <- struct{}{}
		}(startIndex, endIndex, dataset, bankMarketing, ch)
		startIndex = endIndex
		if i < numThread-2 {
			endIndex = endIndex + step
		} else {
			endIndex = len
		}
	}
	for i := 0; i < numThread; i++ {
		<-ch
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].distance < result[j].distance
	})
	m := make(map[string]int)
	for _, v := range result[:k] {
		v1 := dataset[v.index].target
		if _, ok := m[v1]; !ok {
			m[v1] = 1
		} else {
			m[v1]++
		}
	}
	var t string
	var mx int
	for tag, max := range m {
		if max > mx {
			t = tag
			mx = max
		}
	}
	return t
}
