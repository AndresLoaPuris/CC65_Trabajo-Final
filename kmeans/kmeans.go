package kmeans
/*
import (
	"math"
	"sync"
)

type K_Means struct {
	k, max_iter int
	tol         float64
}

var wg sync.WaitGroup
var centroids map[int][]float64
var classifications map[int][][]float64

func ReturnClassifications() map[int][][]float64 {
	return classifications;
}

func min_Pos(slice []float64) int {
	minValue := slice[0]
	minPos := 0
	for i, v := range slice {
		if v < minValue {
			minValue = v
			minPos = i
		}
	}

	return minPos
}

func average(slice [][]float64) []float64 {

	temp := make([]float64, len(slice[0]))

	for _, v := range slice {
		for i := 0; i < len(v); i++ {
			temp[i] = temp[i] + v[i]
		}

	}

	for i := 0; i < len(slice[0]); i++ {
		temp[i] = temp[i] / float64(len(slice))
	}

	return temp

}

func norm(slice []float64) float64 {
	sum := 0.0
	for _, v := range slice {
		sum += math.Pow(v, 2)
	}
	return math.Sqrt(sum)
}

func restar(firstSlice, secondSlice []float64) []float64 {
	temp := make([]float64, len(firstSlice))
	for i := 0; i < len(firstSlice); i++ {
		temp[i] = firstSlice[i] - secondSlice[i]
	}
	return temp
}

func sum(firstSlice, secondSlice []float64) float64 {

	temp := make([]float64, len(firstSlice))

	for i := 0; i < len(firstSlice); i++ {
		temp[i] = firstSlice[i] - secondSlice[i]
	}

	for i := 0; i < len(firstSlice); i++ {
		temp[i] = temp[i] / (secondSlice[i] * 100.0)
	}

	sum := 0.0
	for _, v := range temp {
		sum += v * 10000.0
	}
	return sum
}

func Team(v []float64) {

	defer wg.Done()
	distances := make([]float64, len(centroids))
	for c := 0; c < len(centroids); c++ {
		//sleep := rand.Int63n(1000)
		//time.Sleep(time.Duration(sleep) * time.Millisecond)
		distances[c] = norm(restar(v, centroids[c]))
	}
	classification := min_Pos(distances)
	classifications[classification] = append(classifications[classification], v)

}

func CreateKMeans(kteam int) *K_Means {
	return &(K_Means{k:kteam , tol: 0.001, max_iter: 10})
}

func (KMeans *K_Means) Fit(data [][]float64) {

	centroids = make(map[int][]float64)

	for i := 0; i < KMeans.k; i++ {

		centroids[i] = data[i]
	}

	for i := 0; i < KMeans.max_iter; i++ {

		classifications = make(map[int][][]float64)
		for r := 0; r < KMeans.k; r++ {

			classifications[r] = make([][]float64, 0)
		}

		wg.Add(len(data))

		for _, v := range data {
			go Team(v)
		}

		wg.Wait()

		prev_centroids := make(map[int][]float64)

		for key, value := range centroids {
			prev_centroids[key] = value
		}

		for c := 0; c < len(classifications); c++ {
			centroids[c] = average(classifications[c])
		}

		optimized := true

		for c := 0; c < len(centroids); c++ {
			original_centroid := prev_centroids[c]
			current_centroid := centroids[c]

			if sum(current_centroid, original_centroid) > KMeans.tol {
				optimized = false
			}
		}

		if optimized == true {
			break
		}
	}

}
*/



import (
	"math/rand"
	"time"
)

// Node represents an observation of floating point values
type Node []float64

// Train takes an array of Nodes (observations), and produces as many centroids as specified by
// clusterCount. It will stop adjusting centroids after maxRounds is reached. If there are less
// observations than the number of centroids requested, then Train will return (false, nil).
func Train(Nodes []Node, clusterCount int, maxRounds int) (bool, []Node) {
	if int(len(Nodes)) < clusterCount {
		return false, nil
	}

	// Check to make sure everything is consistent, dimension-wise
	stdLen := 0
	for i, Node := range Nodes {
		curLen := len(Node)

		if i > 0 && len(Node) != stdLen {
			return false, nil
		}

		stdLen = curLen
	}

	centroids := make([]Node, clusterCount)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Pick centroid starting points from Nodes
	for i := 0; i < clusterCount; i++ {
		srcIndex := r.Intn(len(Nodes))
		srcLen := len(Nodes[srcIndex])
		centroids[i] = make(Node, srcLen)
		copy(centroids[i], Nodes[r.Intn(len(Nodes))])
	}

	return Train2(Nodes, clusterCount, maxRounds, centroids)
}

// Provide initial centroids
func Train2(Nodes []Node, clusterCount int, maxRounds int, centroids []Node) (bool, []Node) {
	// Train centroids
	movement := true
	for i := 0; i < maxRounds && movement; i++ {
		movement = false

		groups := make(map[int][]Node)

		for _, Node := range Nodes {
			near := Nearest(Node, centroids)
			groups[near] = append(groups[near], Node)
		}

		for key, group := range groups {
			newNode := meanNode(group)

			if !equal(centroids[key], newNode) {
				centroids[key] = newNode
				movement = true
			}
		}
	}

	return true, centroids
}

// equal determines if two nodes have the same values.
func equal(node1, node2 Node) bool {
	if len(node1) != len(node2) {
		return false
	}

	for i, v := range node1 {
		if v != node2[i] {
			return false
		}
	}

	return true
}

// Nearest return the index of the closest centroid from nodes
func Nearest(in Node, nodes []Node) int {
	count := len(nodes)

	results := make(Node, count)
	cnt := make(chan int)
	for i, node := range nodes {
		go func(i int, node, cl Node) {
			results[i] = distance(in, node)
			cnt <- 1
		}(i, node, in)
	}

	wait(cnt, results)

	mindex := 0
	curdist := results[0]

	for i, dist := range results {
		if dist < curdist {
			curdist = dist
			mindex = i
		}
	}

	return mindex
}

// Distance determines the square Euclidean distance between two nodes
func distance(node1 Node, node2 Node) float64 {
	length := len(node1)
	squares := make(Node, length, length)

	cnt := make(chan int)

	for i, _ := range node1 {
		go func(i int) {
			diff := node1[i] - node2[i]
			squares[i] = diff * diff
			cnt <- 1
		}(i)
	}

	wait(cnt, squares)

	sum := 0.0
	for _, val := range squares {
		sum += val
	}

	return sum
}

// meanNode takes an array of Nodes and returns a node which represents the average
// value for the provided nodes. This is used to center the centroids within their cluster.
func meanNode(values []Node) Node {
	newNode := make(Node, len(values[0]))

	for _, value := range values {
		for j := 0; j < len(newNode); j++ {
			newNode[j] += value[j]
		}
	}

	for i, value := range newNode {
		newNode[i] = value / float64(len(values))
	}

	return newNode
}

// wait stops a function from continuing until the provided channel has processed as
// many items as there are dimensions in the provided Node.
func wait(c chan int, values Node) {
	count := len(values)

	<-c
	for respCnt := 1; respCnt < count; respCnt++ {
		<-c
	}
}