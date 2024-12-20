package no15

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	IsWall  bool
	IsBox   bool
	IsRobot bool
}

type EnlargedNode struct {
	IsWall     bool
	IsBoxLeft  bool
	IsBoxRight bool
	IsRobot    bool
}

func (e *EnlargedNode) IsBox() bool {
	return e.IsBoxLeft || e.IsBoxRight
}

type Warehouse struct {
	Nodes [][]*Node
}

type EnlargedWarehouse struct {
	Nodes [][]*EnlargedNode
}

func NewWarehouse(data []string) *Warehouse {
	nodes := [][]*Node{}
	for _, line := range data {
		row := []*Node{}
		for _, char := range strings.Split(line, "") {
			switch char {
			case "#":
				row = append(row, &Node{IsWall: true})
			case "O":
				row = append(row, &Node{IsBox: true})
			case "@":
				row = append(row, &Node{IsRobot: true})
			default:
				row = append(row, &Node{})
			}
		}
		nodes = append(nodes, row)
	}

	return &Warehouse{Nodes: nodes}
}

func (w *Warehouse) Enlarge() *EnlargedWarehouse {
	nodes := [][]*EnlargedNode{}
	for _, row := range w.Nodes {
		newRow := []*EnlargedNode{}
		for _, node := range row {
			if node.IsWall {
				newRow = append(newRow, []*EnlargedNode{{IsWall: true}, {IsWall: true}}...)
			} else if node.IsBox {
				newRow = append(newRow, []*EnlargedNode{{IsBoxLeft: true}, {IsBoxRight: true}}...)
			} else if node.IsRobot {
				newRow = append(newRow, []*EnlargedNode{{IsRobot: true}, {}}...)
			} else {
				newRow = append(newRow, []*EnlargedNode{{}, {}}...)
			}
		}
		nodes = append(nodes, newRow)
	}

	return &EnlargedWarehouse{Nodes: nodes}
}

func (w *Warehouse) Print() {
	for _, row := range w.Nodes {
		for _, node := range row {
			if node.IsWall {
				fmt.Print("#")
			} else if node.IsBox {
				fmt.Print("O")
			} else if node.IsRobot {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (w *EnlargedWarehouse) Print() {
	for _, row := range w.Nodes {
		for _, node := range row {
			if node.IsWall {
				fmt.Print("#")
			} else if node.IsBoxLeft {
				fmt.Print("[")
			} else if node.IsBoxRight {
				fmt.Print("]")
			} else if node.IsRobot {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (w *Warehouse) GetCurrentRobotLocation() (int, int) {
	for x, row := range w.Nodes {
		for y, node := range row {
			if node.IsRobot {
				return x, y
			}
		}
	}
	return -1, -1
}

func (w *EnlargedWarehouse) GetCurrentRobotLocation() (int, int) {
	for x, row := range w.Nodes {
		for y, node := range row {
			if node.IsRobot {
				return x, y
			}
		}
	}
	return -1, -1
}

func (w *Warehouse) CanMove(sx, sy, vx, vy int) bool {
	if w.Nodes[sx+vx][sy+vy].IsWall {
		return false
	}
	if w.Nodes[sx+vx][sy+vy].IsBox {
		for {
			if w.Nodes[sx+vx][sy+vy].IsWall {
				return false
			}
			if w.Nodes[sx+vx][sy+vy].IsBox {
				sx += vx
				sy += vy
				continue
			}
			break
		}
	}
	return true
}

func (w *Warehouse) Move(direction string) {
	robotX, robotY := w.GetCurrentRobotLocation()
	v := []int{}

	switch direction {
	case "<":
		v = []int{0, -1}
	case "^":
		v = []int{-1, 0}
	case ">":
		v = []int{0, 1}
	case "v":
		v = []int{1, 0}
	}

	if !w.CanMove(robotX, robotY, v[0], v[1]) {
		return
	}

	w.Nodes[robotX][robotY].IsRobot = false
	np := []int{robotX + v[0], robotY + v[1]}
	npc := []int{robotX + v[0], robotY + v[1]}
	w.Nodes[np[0]][np[1]].IsRobot = true
	if w.Nodes[np[0]][np[1]].IsBox {
		for {
			npc[0] += v[0]
			npc[1] += v[1]
			if w.Nodes[npc[0]][npc[1]].IsWall {
				panic("should not happen!")
			}
			if w.Nodes[npc[0]][npc[1]].IsBox {
				continue
			}
			break
		}
		w.Nodes[npc[0]][npc[1]].IsBox = true
		w.Nodes[np[0]][np[1]].IsBox = false
	}
}

func (w *Warehouse) CalculateGPS() int {
	res := 0
	for x, row := range w.Nodes {
		for y, node := range row {
			if node.IsBox {
				res = res + x*100 + y
			}
		}
	}
	return res
}

func (w *EnlargedWarehouse) CanMove(x, y, vx, vy int) bool {
	n := w.Nodes[x][y]
	if n.IsWall {
		return false
	}
	if n.IsRobot {
		return w.CanMove(x+vx, y+vy, vx, vy)
	}
	if n.IsBox() {
		if vx == 0 {
			return w.CanMove(x, y+vy, vx, vy)
		} else {
			if n.IsBoxLeft {
				return w.CanMove(x+vx, y, vx, vy) && w.CanMove(x+vx, y+1, vx, vy)
			} else {
				return w.CanMove(x+vx, y, vx, vy) && w.CanMove(x+vx, y-1, vx, vy)
			}
		}
	}
	return true
}

func (w *EnlargedWarehouse) DoMove(x, y, vx, vy int) {
	n := w.Nodes[x][y]
	np := w.Nodes[x+vx][y+vy]
	if np.IsWall || np.IsRobot {
		panic("should not happen!")
	}

	if np.IsBox() {
		if vx == 0 {
			w.DoMove(x, y+vy, vx, vy)
		} else {
			if np.IsBoxLeft {
				w.DoMove(x+vx, y+vy, vx, vy)
				w.DoMove(x+vx, y+vy+1, vx, vy)
			} else {
				w.DoMove(x+vx, y+vy, vx, vy)
				w.DoMove(x+vx, y+vy-1, vx, vy)
			}
		}
	}
	np.IsBoxLeft = n.IsBoxLeft
	np.IsBoxRight = n.IsBoxRight
	np.IsRobot = n.IsRobot
	np.IsWall = n.IsWall
	n.IsBoxLeft = false
	n.IsBoxRight = false
	n.IsWall = false
	n.IsRobot = false
}

func (w *EnlargedWarehouse) Move(direction string) {
	robotX, robotY := w.GetCurrentRobotLocation()
	v := []int{}

	switch direction {
	case "<":
		v = []int{0, -1}
	case "^":
		v = []int{-1, 0}
	case ">":
		v = []int{0, 1}
	case "v":
		v = []int{1, 0}
	}

	if !w.CanMove(robotX, robotY, v[0], v[1]) {
		return
	}

	w.DoMove(robotX, robotY, v[0], v[1])
}

func (w *EnlargedWarehouse) CalculateGPS() int {
	res := 0
	for x, row := range w.Nodes {
		for y, node := range row {
			if node.IsBoxLeft {
				res = res + x*100 + y
			}
		}
	}
	return res
}

func Solve() error {
	datab, err := os.ReadFile("15/input1.txt")
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}

	datas := string(datab)
	dataspl := strings.Split(datas, "\n")
	var sindex int
	for ind, v := range dataspl {
		if v == "" {
			sindex = ind
		}
	}

	w := NewWarehouse(dataspl[:sindex])
	moves := strings.Split(strings.Join(dataspl[sindex+1:], ""), "")
	for _, move := range moves {
		w.Move(move)
	}
	fmt.Println("Result 1:", w.CalculateGPS())

	we := NewWarehouse(dataspl[:sindex]).Enlarge()
	moves = strings.Split(strings.Join(dataspl[sindex+1:], ""), "")
	for _, move := range moves {
		we.Move(move)
	}
	fmt.Println("Result 2:", we.CalculateGPS())
	return nil
}
