package pstsql

import (
	"1_env"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

//// 创建盐值
//func CreatSalt() string {
//	rand.Seed(time.Now().Unix())
//	salt := make([]byte, commonConst.SaltLength)
//	for i := 0; i < commonConst.SaltLength; i++ {
//		salt[i] = commonConst.SaltPool[rand.Intn(len(commonConst.SaltPool))]
//	}
//	return string(salt)
//}
//
//// 对 password 进行加盐哈希，其中盐值为 salt
//// 返回十六进制哈希值(string)
//func SaltHash(password string, salt string) (string, error) {
//	// 加盐
//	firstLayPassword := password + salt
//
//	// 开始哈希
//	h := sha1.New()
//	if _, err := h.Write([]byte(firstLayPassword)); err != nil {
//		return "", err
//	}
//
//	return fmt.Sprintf("%x", h.Sum(nil)), nil
//}
//
//func NewDefaultUai(uid int, email, password string) *UserAccountInformation {
//	var saltPassword string
//	var err error
//	salt := CreatSalt()
//	if saltPassword, err = SaltHash(password, salt); err != nil {
//		return nil
//	}
//	// 插入账户信息
//	uai := &UserAccountInformation{
//		UserId:            uid, // 让数据库根据serial获取
//		UserEmail:         email,
//		UserLastLoginTime: int(time.Now().Unix()),
//		UserRegisterTime:  int(time.Now().Unix()),
//		UserPassword:      saltPassword,
//		UserType:          commonConst.SmallUser,
//		Salt:              salt,
//		Reserved2:         " ",
//	}
//	return uai
//}
//func NewDefaultUpi(uid int, contactEmail string) *UserPersonalInformation {
//	return &UserPersonalInformation{
//		uid,
//		commonConst.DefaultPhotoUrl,
//		commonConst.DefaultUserName,
//		commonConst.DefaultUserSex,
//		commonConst.DefaultUserPhone,
//		contactEmail,
//		int(time.Now().Unix()),
//		" ",
//		" ",
//	}
//}

type Pgsql struct {
	*Dao
}

func (p *Pgsql) Init() error {
	var db *sql.DB
	// 开始连接数据库
	var err error
	if db, err = sql.Open(env.Conf.Postgresql.DriverName, fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=%s`,
		env.Conf.Postgresql.Host,
		env.Conf.Postgresql.Port,
		env.Conf.Postgresql.User,
		env.Conf.Postgresql.Password,
		env.Conf.Postgresql.Dbname,
		env.Conf.Postgresql.Sslmode,
	)); err != nil {
		return err
	}

	// 设置数据库连接
	db.SetMaxIdleConns(env.Conf.Postgresql.MaxIdleConns)
	db.SetMaxOpenConns(env.Conf.Postgresql.MaxOpenConns)

	// ping 数据库
	if err := db.Ping(); err != nil {
		return err
	}
	p.Dao = &Dao{db}
	return nil
}

func (p *Pgsql) Close() error {
	return p.db.Close()
}

//// 插入单个uai
//func (p *Pgsql) InsertUai(uai *UserAccountInformation) error {
//	return p.dao.Insert(uai)
//}
//
//// 插入单个upi
//func (p *Pgsql) InsertUpi(upi *UserPersonalInformation) error {
//	return p.dao.Insert(upi)
//}
//
//// 这个函数通过 email 或 userId 来获取用户信息 (根据identification的类型判断)
//func (p *Pgsql) GetUai(email string) (*UserAccountInformation, error) {
//	uai := &UserAccountInformation{}
//	results, err := p.dao.Select(uai, "where userEmail=$1", email)
//	if err != nil {
//		return nil, err
//	}
//	if len(results) == 0 {
//		return nil, nil
//	}
//	uai = results[0].(*UserAccountInformation)
//	return uai, nil
//}
//
//func (p *Pgsql) GetUid(email string) (uid int, err error) {
//	uai, err := p.GetUai(email)
//	if err != nil {
//		return 0, err
//	}
//	if uai == nil {
//		return 0, nil
//	}
//	return uai.UserId, nil
//}
//
//// 产生新用户
//func (p *Pgsql) GenerateNewUser(email, password string) error {
//
//	// 插入账户信息 uai
//	uai := NewDefaultUai(0, email, password)
//	var err error
//	if err = p.InsertUai(uai); err != nil {
//		return err
//	}
//
//	// 获取用户Id
//	var uid int
//
//	if uid, err = p.GetUid(email); err != nil {
//		return err
//	}
//	// 插入用户个人信息 upi
//	upi := NewDefaultUpi(uid, email)
//	return p.InsertUpi(upi)
//}
//
//
//// 更新用户登录时间
//func (p *Pgsql) UpdateLastLoginTime(userId int) error {
//	return p.dao.Update(
//		&UserAccountInformation{UserLastLoginTime: int(time.Now().Unix())},
//		"where userId=$1",
//		userId,
//	)
//}
//
//// 修改用户密码
//func (p *Pgsql) UpdatePassword(userId int, newPassword string) error {
//	// 这里的newPassword指的是加盐哈希后的
//	return p.dao.Update(
//		&UserAccountInformation{UserPassword: newPassword},
//		"where userId=$1",
//		userId,
//	)
//}
//
//// 更新用户名
//func (p *Pgsql) UpdateUserName(userId int, newUserName string) error {
//	return p.dao.Update(
//		&UserPersonalInformation{UserName: newUserName},
//		"where userId=$1",
//		userId,
//	)
//}
//
//// 更新用户头像链接
//func (p *Pgsql) UpdatePhotoUrl(userId int, newPhotoUrl string) error {
//	return p.dao.Update(
//		&UserPersonalInformation{UserPhotoUrl: newPhotoUrl},
//		"where userId=$1",
//		userId,
//	)
//}
//
//
