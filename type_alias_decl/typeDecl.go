package type_alias_decl

/*
类型声明
*/
type Node struct {
}
type (
	Point struct{ x, y float64 } // Point 和struct{x,y float64} 是两个不同的类型
	polar Point                  // polar 和Point 是两个不同的类型
)

type TreeNode struct {
	left, right *TreeNode
	value       any
}

// 声明一个接口
type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}

/*
类型方法继承
*/
type Mutex struct {
	// Mutex fields
}

func (m *Mutex) Lock() {
	// Lock implementation
}
func (m *Mutex) Unlock() {
	// Unlock implementation
}

// NewMutex 不会继承方法method
type NewMutex Mutex

// PtrMutex 指针也不会继承方法method
type PtrMutex *Mutex

// 复合型的类型会继承方法 , 可以使用 PrintableMutex.Lock()
type PrintableMutex struct {
	Mutex
}

// 接口类型的也会继承, MyBlock.BlockSize() 有效
type MyBlock Block
