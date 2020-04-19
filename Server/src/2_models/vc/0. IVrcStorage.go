package vc


type IEmailVrcStorage interface {
	SetVrc(prefix string,email string,vrc string,expiredTime int) error
	GetVrc(prefix string,receiver string) (string,error)
	DelVrc(prefix string, receiver string) error
}



