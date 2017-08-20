package loadtest

/**
 * Attempts to attack the single little docker container. Poor little gopher.
 */

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	vegeta "github.com/tsenart/vegeta/lib"
)

var attacksPerSecond = 1
var secondsToAttack = 4

func TestAttack(t *testing.T) {
	meanLatency := attack(attacksPerSecond, secondsToAttack)
	fmt.Printf("Mean latency: %s \n", meanLatency)
	assert.True(t, meanLatency.Seconds() < 5) // Make sure we're maxmium under 5 seconds
}

func randomStartAndGoal() (string, string) {
	randos := []string{"Lundomys", "Brazil", "Dog", "Cat", "Airplane", "Mexico", "Olmec", "Guatemala"}

	rand.Seed(time.Now().UTC().UnixNano())
	startI := rand.Int() % len(randos)

	rand.Seed(time.Now().UTC().UnixNano())
	endI := rand.Int() % len(randos)
	return randos[startI], randos[endI]
}

func randomTargets(numTargets int) []vegeta.Target {
	targets := []vegeta.Target{}
	for i := 0; i < numTargets; i++ {
		start, goal := randomStartAndGoal()

		targets = append(targets, vegeta.Target{
			Method: "GET",
			URL:    fmt.Sprintf("http://localhost:6060/search/%s/%s", start, goal),
		})
	}

	return targets
}

func attack(perSecond int, seconds int) time.Duration {
	rate := uint64(perSecond)
	duration := time.Duration(seconds) * time.Second
	targets := randomTargets(10)
	targeter := vegeta.NewStaticTargeter(targets...)
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration) {
		metrics.Add(res)
	}
	metrics.Close()

	return metrics.Latencies.Mean
}
