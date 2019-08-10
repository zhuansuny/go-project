package process

import (
	"encoding/json"
	"fmt"
	"go_code/chapter16_chatRoom/common/message"
)

func outputGrounp(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("序列化失败 err=", err)
		return
	}
	info := fmt.Sprintf("%d :\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)

}
