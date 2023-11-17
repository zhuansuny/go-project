package main  //二维数组测试
import (
	"fmt"
)
func main(){
	var scores [3][5]float64
	
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%d的班的第%d个学生的成绩\n",i+1,j+1)
			fmt.Scanln(&scores[i][j])
		}
	}

	total := 0.0
	for i := 0; i < len(scores); i++ {
		sum := 0.0
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
		}
		fmt.Printf("第%d的班总成绩%v 平均分%v\n",i+1,sum,sum/ float64(len(scores[0])))
		total +=sum
	}
	fmt.Printf("所有班总成绩%v 平均分%v\n",total,total/float64((len(scores[0])*len(scores))))

}