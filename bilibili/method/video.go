package method

type Video struct {
	Id int 			`json:"id"`
	Title string 	`json:"title" binding:"required,min=2,max=16"`
	Image string 	`json:"image"`
	Author string 	`json:"author"`
	Info string 	`json:"info" binding:"required,min=0,max=200"`
	Views int 		`json:"views" `
	Likes int 		`json:"likes"`
	Coins int 		`json:"coins"`
	Collections int `json:"collections"`
}

//点赞
func (v *Video)GiveLike() {
	v.Likes++
}

//投币
func (v *Video)GiveCoin() {
	v.Coins++
}

//收藏
func (v *Video)Collect() {
	v.Collections++
}

//一键三连
func (v *Video)LongPress() {
	v.Likes++
	v.Collections++
	v.Coins++
}