package commonFunction

import (
	"0_common/commonConst"
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
