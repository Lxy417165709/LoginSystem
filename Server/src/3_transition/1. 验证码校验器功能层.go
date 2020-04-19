package transition

import (
	"0_common/commonConst"
	"2_models/vc"
	"math/rand"
	"time"
)

func CreatVrc() string{

	rand.Seed(time.Now().Unix())
	vrc := make([]byte, commonConst.VrcLength)
	for i := 0; i < commonConst.VrcLength; i++ {
		vrc[i] = commonConst.VrcPool[rand.Intn(len(commonConst.VrcPool))]
	}
	return string(vrc)

}
// 发送验证码
func SendVrc(checker vc.IEmailVrcChecker,email string,expiredTime int) error{
	vrc := CreatVrc()
	return checker.SendVrc(email,vrc,expiredTime)
}
// 判断验证码是否正确
func VrcIsRight(checker vc.IEmailVrcChecker,email,vrc string,whenPassIsNeedToDelete bool) (bool,error){
	return checker.VrcIsRight(email,vrc,whenPassIsNeedToDelete)
}
