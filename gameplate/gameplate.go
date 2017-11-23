package gameplate

// Direction 表示移动方向
type Direction int

// Gameplate 游戏盘面
type Gameplate interface {
	// Clone 复制
	Clone() Gameplate

	// Move 移动
	Move(d Direction) bool

	// Rules 获取玩法 规则
	Rules() string

	// AvailableMoves 获取可以进行的移动
	AvailableMoves() map[rune]Direction

	// Score 统计得分
	Score() int

	// Print 打印盘面分数第几轮etc
	Print() string

	// IsGameOver 是否结束
	// true - 结束了 动弹不得
	// false - 没解说 我还能抢救一下
	IsGameOver() bool

	// GenerateNewCells 生成新的格子
	// 返回生成格子的数量
	GenerateNewCells() int

	// NewGame 新一局
	NewGame()

	// clear 清空盘面
	clear()
}
