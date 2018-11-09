package main

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	Nodes []*Node
}

func (t *Tree) addNode(node *Node, value int) {
	if value < node.Value {
		if node.Left.Value < -1 {
			t.addNode(node.Left, value)
		} else {
			node.Left.Value = value
		}
	} else if value > node.Value {
		if node.Right.Value < -1 {
			t.addNode(node.Left, value)
		} else {
			node.Right.Value = value
		}
	}
}

func (t *Tree) addParent(node Node) {

}

func (t *Tree) addRoot(value int) {
	t.Nodes = make([]*Node, 0)
	nRoot := &Node{}
	nRoot.Value = value
	t.Nodes = append(t.Nodes, nRoot)
}

func min(values []int) {
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
}

func createArrangements(arrangement []int, index int) int {
	root := arrangement[0]
	leaf := arrangement[1]
	return root
}

type IdxFrequency struct {
	Index     int
	Frequency int
}

type TreeNode struct {
	value int
	level int
	left bool
	right bool
}

type Tree struct {
	TreeNodes []TreeNode
}

func (t *Tree) buildTree(values, parent) *Tree {
	for i = 0; i < len(values); i++ {
		if value < root {
			treeNode := &TreeNode{value: value; level: level}
		}
	}
}

func 

func getIndexOfItem(value int) int {
	if value == 12 {
		return 0
	}
	if value == 15 {
		return 1
	}
	if value == 20 {
		return 2
	}
	if value == 25 {
		return 3
	}
}

func getSearchFrequencyOfItem(value int) int {
	if value == 12 {
		return 4
	}
	if value == 15 {
		return 3
	}
	if value == 20 {
		return 6
	}
	if value == 25 {
		return 2
	}
}

func main() {
	searchElements := []int{12, 15, 20, 25}
	searchFrequency := []int{4, 3, 6, 2}
	var memoizationMatrix [4][4]int

	elements := make(map[int]interface{})
	elements[12] = &IdxFrequency{Index: 0, Frequency: 4}
	elements[15] = &IdxFrequency{Index: 1, Frequency: 3}
	elements[20] = &IdxFrequency{Index: 2, Frequency: 6}
	elements[25] = &IdxFrequency{Index: 3, Frequency: 3}

	noOfElements := len(searchElements)
	numArrangements := noOfElements - 1

	for i := 0; i < noOfElements; i++ { // batch-size
		if i == 0 {
			memoizationMatrix[0][0] = searchFrequency[0]
			memoizationMatrix[1][1] = searchFrequency[1]
			memoizationMatrix[2][2] = searchFrequency[2]
			memoizationMatrix[3][3] = searchFrequency[3]
		} else {
			subProblemSize := i + 1 // 1 + 1
			extraCost := make([]int, 0)
			necessaryCost := 0
			for j := 0; j < subProblemSize; j++ { // subset
				subset = searchElements[j:subProblemSize]
				necessaryCost := calculateNecessaryCost(subset)
				extraCost := make([]int, 0)
				for k := 0; k < length(subset); k++ { //root
					if k == 0 {
						items := subset[k+1:length(subset)]
						if len(items) > 1 {
							// get the index of the range
							// get the range's extra cost from memoization
						} else {
							cost := getSearchFrequencyOfItem(items[0])
							extraCost = append(extraCost, cost)  
						}
					} else if k > 0 && k < length(subset) {
						items1 := subset[0:k]
						items2 := subset[k:length(subset)]
						
					} else {
						
					}
							
								
				necessaryCost := calculateNecessaryCost(subset)
				for k := j; k < subProblemSize; k++ {



					root := searchElements[k]
					left := 
				}



				




				root := subset[j]


				necessaryCost := necessaryCost + searchFrequency[j]
				
				
				if j < subProblemSize {
					cost := searchFrequency[j+1] * j

				} else {
					cost := searchFrequency[j-1] * j

				}
				necessaryCost := necessaryCost + searchFrequency[j]
			}	
		}	









			numArrangements := subProblemSize + 1 //
			for j := 0; j < numArrangements; j++ {
				arrangement := searchElements[j:i]
				for k := 1; k <= len(arrangement); k++ {
					root := arrangement[k]
					necessaryCost := searchFrequency[j] + searchFrequency[i]
				}

				root := createArrangements(searchElements[j:i], j)

				necessaryCost := searchFrequency[j] + searchFrequency[i]
				extraCost := make([]int, 0)
				for k := 1; k <= len(arrangement); k++ {
					c := searchFrequency[0] * k
					extraCost = append(extraCost, c)
				}

				optimalCost := necessaryCost + min(extraCost)
			}
		}
	}
}
