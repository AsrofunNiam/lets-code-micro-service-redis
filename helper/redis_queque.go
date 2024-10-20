package helper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/AsrofunNiam/lets-code-micro-service_redis/model/domain"
	"github.com/redis/go-redis/v9"
)

func TakeRedisQueue(client *redis.Client, key string) {
	ctx := context.Background()

	//  running redis forever
	for {
		jobData, err := client.RPop(ctx, key).Result()
		if err == redis.Nil {
			fmt.Println("No jobs in queue, waiting...")
			time.Sleep(30 * time.Second)
			continue
		} else if err != nil {
			log.Fatalf("Error fetching job from queue: %v", err)
		}

		// Unmarshal job data
		var job domain.JobQueue
		err = json.Unmarshal([]byte(jobData), &job)
		if err != nil {
			log.Fatalf("Error unmarshalling job data: %v", err)
		}

		err = ProcessJobQueue(job)
		if err != nil {
			log.Printf("Error processing job: %v", err)

			err = client.LPush(ctx, key+"_failed", jobData).Err()
			if err != nil {
				log.Fatalf("Error pushing job back to queue: %v", err)
			}
		} else {
			log.Printf("Job processed successfully: %s", jobData)
		}
	}

}

func ProcessJobQueue(jobQueue domain.JobQueue) error {
	fmt.Printf("Processing job: %v\n", jobQueue)

	// Prepare the HTTP request
	req, err := http.NewRequest(jobQueue.Method, jobQueue.URL, bytes.NewBuffer([]byte(jobQueue.Payload)))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{
		Timeout: 2 * time.Hour,
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}
	fmt.Printf("Response status: %s, body: %s\n", resp.Status, string(body))

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		return nil
	} else {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}
