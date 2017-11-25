package gameplate

import (
	crand "crypto/rand"
	"fmt"
	big "math/big"
)

// Grid44 4*4格子 存放着数字 标准的2048盘面
//    j0  ->  3
// i0
// |
// v
// 3
// 131072 应该是4*4的盘面中会出现的数字最大值 1 << 17
// 相应的 一个格子代表的分数是 数字 << 1 (4或以上有效)
// 实现GamePlate接口
type Grid44 struct {
	Data       [4][4]int // Data 保存数字
	TurnNumber int       // TurnNumber 第几轮
}

// 方向常量
const (
	UP    Direction = iota // UP 上
	DOWN                   // DOWN 下
	LEFT                   // LEFT 左
	RIGHT                  // RIGHT 右
	NONE                   // NONE 不移动 备选

	UGLY_UI = `
* * * * * * * * * * * * * * * * * * * * * * * *
    Score = %d
    Turns = %d
* * * * * * * * * * * * * * * * * * * * * * * *
    + - - - - + - - - - + - - - - + - - - - +
    |         |         |         |         |
    | %7d | %7d | %7d | %7d |
    |         |         |         |         |
    + - - - - + - - - - + - - - - + - - - - +
    |         |         |         |         |
    | %7d | %7d | %7d | %7d |
    |         |         |         |         |
    + - - - - + - - - - + - - - - + - - - - +
    |         |         |         |         |
    | %7d | %7d | %7d | %7d |
    |         |         |         |         |
    + - - - - + - - - - + - - - - + - - - - +
    |         |         |         |         |
    | %7d | %7d | %7d | %7d |
    |         |         |         |         |
    + - - - - + - - - - + - - - - + - - - - +
`
)

// 验证接口实现
var _ Gameplate = (*Grid44)(nil)

// Clone 复制
func (g *Grid44) Clone() Gameplate {
	var cloneGp = &Grid44{}
	cloneGp.TurnNumber = g.TurnNumber
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			cloneGp.Data[i][j] = g.Data[i][j]
		}
	}
	return cloneGp
}

// Move 移动
func (g *Grid44) Move(d Direction) bool {
	data := &g.Data
	var tmpGrid *Grid44
	var ok bool
	if tmpGrid, ok = g.Clone().(*Grid44); !ok {
		panic("Not support game mode!")
	}
	switch d {
	case UP:
		for y := 0; y < 4; y++ {
			for x := 0; x < 3; x++ {
				for nx := x + 1; nx <= 3; nx++ {
					if data[nx][y] > 0 {
						if data[x][y] <= 0 {
							data[x][y] = data[nx][y]
							data[nx][y] = 0
							x--
						} else if data[x][y] == data[nx][y] {
							data[x][y] += data[nx][y]
							data[nx][y] = 0
						}
						break
					}
				}
			}
		}
	case DOWN:
		for y := 0; y < 4; y++ {
			for x := 3; x > 0; x-- {
				for nx := x - 1; nx >= 0; nx-- {
					if data[nx][y] > 0 {
						if data[x][y] <= 0 {
							data[x][y] = data[nx][y]
							data[nx][y] = 0
							x++
						} else if data[x][y] == data[nx][y] {
							data[x][y] += data[nx][y]
							data[nx][y] = 0
						}
						break
					}
				}
			}
		}
	case LEFT:
		for x := 0; x < 4; x++ {
			for y := 0; y < 3; y++ {
				for ny := y + 1; ny <= 3; ny++ {
					if data[x][ny] > 0 {
						if data[x][y] <= 0 {
							data[x][y] = data[x][ny]
							data[x][ny] = 0
							y--
						} else if data[x][y] == data[x][ny] {
							data[x][y] += data[x][ny]
							data[x][ny] = 0
						}
						break
					}
				}
			}
		}
	case RIGHT:
		for x := 0; x < 4; x++ {
			for y := 3; y > 0; y-- {
				for ny := y - 1; ny >= 0; ny-- {
					if data[x][ny] > 0 {
						if data[x][y] <= 0 {
							data[x][y] = data[x][ny]
							data[x][ny] = 0
							y++
						} else if data[x][y] == data[x][ny] {
							data[x][y] += data[x][ny]
							data[x][ny] = 0
						}
						break
					}
				}
			}
		}
	}
	if g.diff(tmpGrid) {
		g.TurnNumber++
		return true
	}
	return false
}

// Rules 获取玩法 规则
func (g *Grid44) Rules() string {
	return `
W, w move up
S, s move down
A, a move left
D, d move right
`
}

// AvailableMoves 获取可以进行的移动
// 可用按键 - Direction 的map
func (g *Grid44) AvailableMoves() map[rune]Direction {
	return map[rune]Direction{
		'w': UP,
		'W': UP,
		's': DOWN,
		'S': DOWN,
		'a': LEFT,
		'A': LEFT,
		'd': RIGHT,
		'D': RIGHT,
	}
}

// Score 统计得分
func (g *Grid44) Score() int {
	score := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.Data[i][j] > 2 {
				score += g.Data[i][j] << 1
			}
		}
	}
	return score
}

// Print 打印盘面
func (g *Grid44) Print() string {
	return fmt.Sprintf(UGLY_UI, g.Score(), g.TurnNumber,
		g.Data[0][0], g.Data[0][1], g.Data[0][2], g.Data[0][3],
		g.Data[1][0], g.Data[1][1], g.Data[1][2], g.Data[1][3],
		g.Data[2][0], g.Data[2][1], g.Data[2][2], g.Data[2][3],
		g.Data[3][0], g.Data[3][1], g.Data[3][2], g.Data[3][3])
}

// IsGameOver 是否结束
// true - 结束了 动弹不得
// false - 没结束 我还能抢救一下
func (g *Grid44) IsGameOver() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.Data[i][j] == 0 {
				return false
			}
			// 上面有没有?
			if i-1 >= 0 && g.Data[i-1][j] == g.Data[i][j] {
				return false
			}
			// 下面有没有?
			if i+1 < 4 && g.Data[i+1][j] == g.Data[i][j] {
				return false
			}
			// 左面有没有?
			if j-1 >= 0 && g.Data[i][j-1] == g.Data[i][j] {
				return false
			}
			// 右面有没有?
			if j+1 < 4 && g.Data[i][j+1] == g.Data[i][j] {
				return false
			}
		}
	}
	return true
}

// GenerateNewCells 生成新的格子
// 返回生成格子的数量
func (g *Grid44) GenerateNewCells() int {
	availableGrids := []*int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.Data[i][j] == 0 {
				availableGrids = append(availableGrids, &g.Data[i][j])
			}
		}
	}

	// 为了真随机...
	var v, per *big.Int
	var err error
	bond := &big.Int{}
	bond.SetUint64(uint64(len(availableGrids)))
	v, err = crand.Int(crand.Reader, bond)
	if err != nil {
		panic(err)
	}
	bond.SetUint64(10)
	per, err = crand.Int(crand.Reader, bond)
	if err != nil {
		panic(err)
	}

	if per.Uint64() == 9 {
		*availableGrids[v.Uint64()] = 4
	} else {
		*availableGrids[v.Uint64()] = 2
	}
	return 1
}

// NewGame 新一局
func (g *Grid44) NewGame() {
	g.clear()
	var _ = g.GenerateNewCells()
	var _ = g.GenerateNewCells()
	return
}

func (g *Grid44) clear() {
	g.Data = [4][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	g.TurnNumber = 0
	return
}

func (g *Grid44) diff(b *Grid44) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.Data[i][j] != b.Data[i][j] {
				return true
			}
		}
	}
	return g.TurnNumber != b.TurnNumber
}
