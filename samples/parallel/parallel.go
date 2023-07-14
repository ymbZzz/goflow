package parallel

import (
	"fmt"
	flow "github.com/s8sg/goflow/flow/v1"
	"time"
)

// Workload function
func node1(data []byte, option map[string][]string) ([]byte, error) {
	result := fmt.Sprintf("1-%s", string(data))
	time.Sleep(time.Second * 5)
	fmt.Println(result)
	return []byte(result), nil
}

// Workload function
func node2(data []byte, option map[string][]string) ([]byte, error) {
	result := fmt.Sprintf("2-%s", string(data))
	time.Sleep(time.Second * 5)
	fmt.Println(result)
	return []byte(result), nil
}

// Workload function
func node3(data []byte, option map[string][]string) ([]byte, error) {
	result := fmt.Sprintf("3-%s", string(data))
	time.Sleep(time.Second * 5)
	fmt.Println(result)
	return []byte(result), nil
}

// Workload function
func node4(data []byte, option map[string][]string) ([]byte, error) {
	result := fmt.Sprintf("4-%s", string(data))
	time.Sleep(time.Second * 5)
	fmt.Println(result)
	return []byte(result), nil
}

// Aggregator can be used to collect and map response from multiple in-degree as a request to node
// here node4 uses the result from node2 and node3
func node4Aggregator(data map[string][]byte) ([]byte, error) {
	aggregatedResult := fmt.Sprintf("(node2: %s, node3: %s)", string(data["node2"]), string(data["node3"]))
	return []byte(aggregatedResult), nil
}

// DefineWorkflow Define provide definition of the workflow
func DefineWorkflow(workflow *flow.Workflow, context *flow.Context) error {
	dag := workflow.Dag()
	dag.Node("node1", node1)
	dag.Node("node2", node2)
	dag.Node("node3", node3)
	dag.Node("node4", node4, flow.Aggregator(node4Aggregator))
	dag.Edge("node1", "node2")
	dag.Edge("node1", "node3")
	dag.Edge("node3", "node4")
	dag.Edge("node2", "node4")
	return nil
}
