package configuration

import (
	"sync"

	"github.com/Tony-Lucas/gopportunities/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}

type single struct {
	*gorm.DB
}

var singleInstance *single

func DbSingleInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {

			dsn := "root:containers-us-west-155.railway.app@tcp(containers-us-west-155.railway.app:7867)/railway?charset=utf8mb4&parseTime=True&loc=Local"
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				panic(err)
			} else {
				err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.DeliverAdress{}, &models.Contact{}, &models.Admin{})
				if err != nil {
					panic(err)
				} else {
					singleInstance = &single{
						db,
					}
				}

			}
		}
	}

	return singleInstance
}
