package main

import (
	"fmt"
	vegeta "github.com/tsenart/vegeta/lib"
	"math/rand"
	"time"
)

func main() {
	url := "http://localhost:8080"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var targets []vegeta.Target
	for i := 0; i < 1000; i++ {
		tagId := r.Int63n(10) + 1
		featureId := r.Int63n(1000) + 1
		targets = append(targets, vegeta.Target{
			Method: "GET",
			URL:    fmt.Sprintf("%s/user_banner?tag_id=%d&feature_id=%d", url, tagId, featureId),
			Header: map[string][]string{
				"token": {"admin_token"},
			},
		})

		targets = append(targets, vegeta.Target{

			Method: "GET",
			URL:    fmt.Sprintf("%s/banner?tag_id=%d", url, tagId),
			Header: map[string][]string{
				"token": {"admin_token"},
			},
		})

		targets = append(targets, vegeta.Target{
			Method: "GET",
			URL:    fmt.Sprintf("%s/banner?feature_id=%d", url, featureId),
			Header: map[string][]string{
				"token": {"admin_token"},
			},
		})
	}
	targeter := vegeta.NewStaticTargeter(targets...)
	attacker := vegeta.NewAttacker()

	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 1 * time.Minute
	var metrics vegeta.Metrics

	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	//fmt.Printf("Results\n: %v\n", metrics)
	fmt.Printf("Results:\n")
	fmt.Printf("Total Requests: %d\n", metrics.Requests)
	fmt.Printf("Success Rate: %.2f%%\n", metrics.Success*100)
	fmt.Printf("Request Rate: %.2f requests/second\n", float64(metrics.Requests)/metrics.Duration.Seconds())
	fmt.Printf("Average Response Time: %s\n", metrics.Latencies.Mean)
	fmt.Printf("Max Response Time: %s\n", metrics.Latencies.Max)
	fmt.Printf("Throughput: %.2f requests/second\n", metrics.Throughput)
	fmt.Printf("Bytes Out: Total: %d Avg: %.2f\n", metrics.BytesOut.Total, float64(metrics.BytesOut.Total)/float64(metrics.Requests))
	fmt.Printf("Bytes In: Total: %d Avg: %.2f\n", metrics.BytesIn.Total, float64(metrics.BytesIn.Total)/float64(metrics.Requests))
	fmt.Println("Status Codes:")
	for code, count := range metrics.StatusCodes {
		fmt.Printf("  %s: %d\n", code, count)
	}
	fmt.Println("Errors:")
	for _, err := range metrics.Errors {
		fmt.Printf("  %s\n", err)
	}
}
