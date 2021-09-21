package pattern

type INhanVien interface {
	Clone() INhanVien
	TinhLuong() (luong int)
	ChangeId(id int, empids []int) ([]int, error)
}
