package vc


type IEmailVrcChecker interface {
	IEmailVrcStorage
	SendVrc(email,vrc string,expiredTime int) error
	VrcIsRight(email,vrc string,whenPassIsNeedToDelete bool)  (bool,error)
}



