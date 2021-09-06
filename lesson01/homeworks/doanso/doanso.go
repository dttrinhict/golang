package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main()  {
reader := bufio.NewReader(os.Stdin)
fmt.Printf("Hãy nhập vào con số may mắn của bạn: ")
str, _ := reader.ReadString('\n')
so, _ := ChuyenChuoiSangInt(str)
XoSoMinhKhai(int(so))
}



/* Dữ đoán kết quả xổ số
Kết quả xổ số là số tự nhiên dương sinh ngẫu nhiên trong khoảng từ 0 đến 100 gọi là jackpot
Người chơi nhập số mình mong muốn và chờ so sánh với jackpot
Nếu số người chơi nhập vào:
- nhỏ hơn jackpot: Thiếu một chút, chọn số to hơn đi bạn
- lớn hơn jackpot: Bạn chơi hơi mạnh lên thừa một chút
- bằng jackpot: Bạn là người kém may mắn vì nếu mua xổ số thì bạn đã có thể được cầm tiền về, ĐEN
*/
func XoSoMinhKhai(so int) {
	// Tạo rand source thay thổi liên tục theo thời gian,
	// để đảm bảo mỗi số ngẫu nhiên được sinh ra theo một rand source khác nhau và cho số ngẫu nhiên khác nhau
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	jackpot := r1.Intn(100)
	if so < jackpot {
		fmt.Printf("Số bạn chọn là %v jackpot là %v, Chọn số to hơn đi bạn\n", so, jackpot)
	}else if so > jackpot {
		fmt.Printf("Số bạn chọn là %v jackpot là %v, Chọn số nhỏ hơn đi bạn\n", so, jackpot)
	}else{
		fmt.Printf("Số bạn chọn là %v jackpot là %v, Bạn là người kém may mắn vì nếu mua xổ số thì bạn đã có thể được cầm tiền về, ĐEN\n", so, jackpot)
	}
}


/* Chuyển dữ liệu chuỗi đầu vào sang int64
 */
func ChuyenChuoiSangInt(str string) (int64, error) {
	return strconv.ParseInt(strings.Trim(str,"\n"), 10, 8)
}