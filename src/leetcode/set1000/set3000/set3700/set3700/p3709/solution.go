package p3709

import "sort"

type Exam struct {
	Time     int
	ScoreSum int
}

type ExamTracker struct {
	exams []Exam
}

func Constructor() ExamTracker {
	exams := []Exam{{0, 0}}
	return ExamTracker{
		exams: exams,
	}
}

func (this *ExamTracker) lastExam() Exam {
	return this.exams[len(this.exams)-1]
}

func (this *ExamTracker) Record(time int, score int) {
	cur := Exam{Time: time, ScoreSum: score + this.lastExam().ScoreSum}
	this.exams = append(this.exams, cur)
}

func (this *ExamTracker) TotalScore(startTime int, endTime int) int64 {
	l := sort.Search(len(this.exams), func(i int) bool {
		return this.exams[i].Time >= startTime
	})
	r := sort.Search(len(this.exams), func(i int) bool {
		return this.exams[i].Time > endTime
	})
	r--
	res := int64(this.exams[r].ScoreSum)
	if l > 0 {
		res -= int64(this.exams[l-1].ScoreSum)
	}
	return res
}

/**
 * Your ExamTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Record(time,score);
 * param_2 := obj.TotalScore(startTime,endTime);
 */
