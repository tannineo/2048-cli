package gameplate

// Grid44 4*4格子 存放着数字 标准的2048盘面
// 实现GamePlate接口
type Grid44 struct {
	Data [4][4]int `json:"data"`
}

// 方向常量
const (
	UP    Direction = iota // UP 上
	DOWN                   // DOWN 下
	LEFT                   // LEFT 左
	RIGHT                  // RIGHT 右
	NONE                   // NONE 不移动 备选
)

// 验证接口实现
var _ Gameplate = (*Grid44)(nil)

// Clone 复制
func (*Grid44) Clone() *Gameplate {
	return nil
}

// Move 移动
func (*Grid44) Move(d Direction) {
	return
}

// Rules 获取玩法 规则
func (*Grid44) Rules() string {
	return ""
}

// AvailableMoves 获取可以进行的移动
// 可用按键 - Direction 的map
func (*Grid44) AvailableMoves() map[string]Direction {
	return map[string]Direction{
		"w": UP,
		"W": UP,
		"s": DOWN,
		"S": DOWN,
		"a": LEFT,
		"A": LEFT,
		"d": RIGHT,
		"D": RIGHT,
	}
}

// Score 统计得分
func (*Grid44) Score() ScoreRecord {
	return 0
}

// Print 打印盘面
func (*Grid44) Print() string {
	return ""
}

// IsGameOver 是否结束
// true - 结束了 动弹不得
// false - 没解说 我还能抢救一下
func (*Grid44) IsGameOver() bool {
	return false
}

// GenerateNewCells 生成新的格子
// 返回生成格子的数量
func (*Grid44) GenerateNewCells() int {
	return 0
}

// Ranklist 获取高分榜
func (*Grid44) Ranklist() string {
	return ""
}
