package model
type Student struct{
	Name string
	Score float64

}
type stu struct{
	name string
	score float64
}

func NewStudent(n string,s float64) *stu{  //使用一个函数创建一个stu结构体实例，并将地址返回（工厂模式）
	return &stu{                        //等同于JAVA中的构造方法
		name : n ,
	    score : s,
	}

}

func (s *stu) GetScore() float64{   //score的get set方法
	return s.score
}

func (s *stu) SetScore(f float64) float64{   
	s.score = f
	return s.score
}

func (s *stu) GetName() string{    //name的get set方法
	return s.name
}

func (s *stu) SetName(str string) string{   
	s.name  = str
	return s.name
}