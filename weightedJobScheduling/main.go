package main

type Job struct {
	Name        string
	StartTime   int
	EndTime     int
	ProfitValue int
}

type TimeSpan struct {
	Start int
	End   int
}

func isOverLapping(t1 int, timeSpan *TimeSpan) bool {
	overLap := false

	for n := timeSpan.Start; n < (timeSpan.Start-timeSpan.End)+1; n++ {
		if t1 == n {
			overLap = true
			break
		}
	}
	return overLap
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func main() {
	jobs := make([]Job, 0)

	job1 := Job{Name: "a", StartTime: 1, EndTime: 4, ProfitValue: 3}
	job2 := Job{Name: "b", StartTime: 2, EndTime: 6, ProfitValue: 5}
	job3 := Job{Name: "c", StartTime: 4, EndTime: 7, ProfitValue: 2}
	job4 := Job{Name: "d", StartTime: 6, EndTime: 8, ProfitValue: 6}
	job5 := Job{Name: "e", StartTime: 5, EndTime: 9, ProfitValue: 4}
	job6 := Job{Name: "f", StartTime: 7, EndTime: 10, ProfitValue: 8}

	jobs = append(jobs, job1)
	jobs = append(jobs, job2)
	jobs = append(jobs, job3)
	jobs = append(jobs, job4)
	jobs = append(jobs, job5)
	jobs = append(jobs, job6)

	optProfitValues := make([]int, 6)
	optProfitValues = append(optProfitValues, job1.EndTime)
	optProfitValues = append(optProfitValues, job2.EndTime)
	optProfitValues = append(optProfitValues, job3.EndTime)
	optProfitValues = append(optProfitValues, job4.EndTime)
	optProfitValues = append(optProfitValues, job5.EndTime)
	optProfitValues = append(optProfitValues, job6.EndTime)

	for i := 1; i <= len(jobs); i++ {
		for j := 0; j < i; j++ {
			timeSpan := &TimeSpan{Start: jobs[i].StartTime, End: jobs[i].EndTime}
			if !isOverLapping(jobs[j].EndTime, timeSpan) {
				totalProfit := jobs[j].EndTime + jobs[i].EndTime
				if totalProfit != optProfitValues[i] {
					optProfitValues[i] = max(totalProfit, optProfitValues[i])
				}
			}
		}
	}
}
