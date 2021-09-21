package pattern

import (
	"errors"
	"fmt"
)

var empIds = []int{}

/* Nhân viên hợp đồng
*/
type Contract struct {
	empId  int
	basicpay int
}

func NewContrat() INhanVien {
	return &Contract{
		empId: 1,
		basicpay: 1500000,
	}
}

/* Clone Nhân viên hợp đồng
*/
func (c *Contract) Clone() INhanVien {
	return &Contract{
		empId: c.empId,
		basicpay: c.basicpay,
	}
}
/* Tính lương
 */
func (c *Contract) TinhLuong() (luong int) {
	return c.basicpay
}

/* Change empId
 */
func (c *Contract) ChangeId(id int, empids []int) ([]int, error) {
	for _, v := range empids {
		if v == id {
			return empids, errors.New("Mã nhân viên đã tồn tai trên hệ thống, vui lòng dùng mã khác")
		}
	}
	c.empId = id
	empids = append(empids, id)
	return empids, nil
}

/* Nhân viên chính thức
*/
type Permanent struct {
	contract Contract
	pf       int
}

func NewPermanent() INhanVien {
	return &Permanent{
		contract: Contract{
			empId: 1,
			basicpay: 1000000,
		},
		pf: 500000,
	}
}

/* Clone nhân viên chính thức
*/
func (p *Permanent) Clone() INhanVien {
	return &Permanent{
		contract: Contract{
			empId: p.contract.empId,
			basicpay: p.contract.basicpay,
		},
		pf: p.pf,
	}
}

/* Tính lương
 */
func (p *Permanent) TinhLuong() (luong int) {
	return p.contract.basicpay + p.pf
}

/* Change empId
 */
func (p *Permanent) ChangeId(id int, empids []int) ([]int, error) {
	for _, v := range empids {
		if v == id {
			return empids, errors.New("Mã nhân viên đã tồn tai trên hệ thống, vui lòng dùng mã khác")
		}
	}
	p.contract.empId = id
	empids = append(empids, id)
	return empids, nil
}

func Demo()  {
	var err error
	contract := NewContrat()
	permanent := NewPermanent()

	var nhanviens []interface{}
	nv1 := permanent.Clone()
	empIds, err = nv1.ChangeId(1, empIds)
	if err != nil {
		fmt.Println(err.Error())
	}else{
		nv1.TinhLuong()
		nhanviens = append(nhanviens, nv1)
	}

	nv2 := permanent.Clone()
	empIds, err = nv2.ChangeId(2, empIds)
	if err != nil {
		fmt.Println(err.Error())
	}else{
		nv2.TinhLuong()
		nhanviens = append(nhanviens, nv2)
	}

	nv3 := permanent.Clone()
	empIds, err = nv3.ChangeId(2, empIds)
	if err != nil {
		fmt.Println(err.Error())
	}else{
		nv2.TinhLuong()
		nhanviens = append(nhanviens, nv3)
	}

	nv4 := contract.Clone()
	empIds, err = nv4.ChangeId(3, empIds)
	if err != nil {
		fmt.Println(err.Error())
	}else{
		nv4.TinhLuong()
		nhanviens = append(nhanviens, nv4)
	}

	nv5 := contract.Clone()
	empIds, err = nv5.ChangeId(3, empIds)
	if err != nil {
		fmt.Println(err.Error())
	}else{
		nv5.TinhLuong()
		nhanviens = append(nhanviens, &nv5)
	}

	/*
	Clone method trả về con tro nên khi kiển tra kiểu của interface ta phải dùng con trỏ
	*/
	var tongSoTien int = 0
	for _, v := range nhanviens {
		if c, okContract := v.(*Contract); okContract { //Golang if assignment statement
			tongSoTien = tongSoTien + c.TinhLuong()
		}
		if p, okPermanent := v.(*Permanent); okPermanent {
			tongSoTien = tongSoTien + p.TinhLuong()
		}
	}
	fmt.Printf("Tổng số tiền phải thanh toán là: %d", tongSoTien)
}