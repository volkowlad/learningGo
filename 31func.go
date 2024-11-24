package main

import (
	"errors"
	"fmt"
)

type MergeJob struct {
	Dicts  []map[string]string
	Merged map[string]string
	IsFin  bool
}

var (
	NotDicts = errors.New("at least 2")
	NilDic   = errors.New("nil error")
)

func (j *MergeJob) setFin() {
	j.IsFin = true
}

func Execute(job *MergeJob) (*MergeJob, error) {
	defer job.setFin()
	if len(job.Dicts) < 2 {
		return job, NotDicts
	}

	job.Merged = make(map[string]string)
	for _, dic := range job.Dicts {
		if dic == nil {
			return job, NilDic
		}
		for i, j := range dic {
			job.Merged[i] = j
		}
	}
	return job, nil
}

func main() {
	fmt.Println(Execute(&MergeJob{Dicts: []map[string]string{{"a": "b"}, nil}}))
}
