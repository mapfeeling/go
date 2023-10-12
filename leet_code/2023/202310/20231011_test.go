package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func topStudents(positiveFeedback []string, negativeFeedback []string, report []string, studentId []int, k int) []int {
	var (
		feedbackWords = map[string]int{}
		res           []int
	)

	for _, val := range positiveFeedback {
		feedbackWords[val] = 3
	}
	for _, val := range negativeFeedback {
		feedbackWords[val] = -1
	}

	type Pair struct {
		score, id int
	}

	var A = make([]Pair, len(report))
	for index, val := range report {
		count := 0
		for _, data := range strings.Split(val, " ") {
			count += feedbackWords[data]
		}
		A[index] = Pair{
			score: count, id: studentId[index],
		}
	}

	sort.Slice(A, func(i, j int) bool {
		a, b := A[i], A[j]
		return a.score > b.score || a.score == b.score && a.id < b.id
	})

	for _, v := range A[:k] {
		res = append(res, v.id)
	}

	return res
}

func Test20231011(t *testing.T) {
	var (
		positiveFeedback = []string{"smart", "brilliant", "studious"}
		negativeFeedback = []string{"not"}
		report           = []string{"this student is not studious", "the student is smart"}
		studentId        = []int{1, 2}
		k                = 2
	)
	fmt.Println(topStudents(positiveFeedback, negativeFeedback, report, studentId, k))
}

// 奖励最顶尖的k名学生
/*
给你两个字符串数组 positive_feedback 和 negative_feedback ，分别包含表示正面的和负面的词汇。不会 有单词同时是正面的和负面的
一开始，每位学生分数为 0 。每个正面的单词会给学生的分数 加 3 分，每个负面的词会给学生的分数 减  1 分
给你 n 个学生的评语，用一个下标从 0 开始的字符串数组 report 和一个下标从 0 开始的整数数组 student_id 表示，其中 student_id[i] 表示这名学生的 ID ，这名学生的评语是 report[i] 。每名学生的 ID 互不相同
给你一个整数 k ，请你返回按照得分 从高到低 最顶尖的 k 名学生。如果有多名学生分数相同，ID 越小排名越前
*/
