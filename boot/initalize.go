package boot

import (
	"gopkg.in/loremipsum.v1"
	"math/rand"
	"service-catalog/common"
	"service-catalog/config"
	"service-catalog/internal"
	"time"
)

func populateData() {
	for i := 0; i < 10; i++ {
		id, _ := common.UniqueId()
		loremIpsumGenerator := loremipsum.NewWithSeed(int64(rand.Intn(200)))
		desc := loremIpsumGenerator.Words(20)
		internal.DataBase["Services"] = append(internal.DataBase["Services"].([]internal.Service),
			internal.Service{
				ID:            id,
				Name:          "Service " + id,
				Description:   desc,
				LatestVersion: "1.0.1",
				Versions:      []string{"1.0.0", "1.0.1"},
				CreatedAt:     time.Now().Unix(),
				UpdatedAt:     time.Now().Unix(),
			})
	}
}

func Initialize() {
	populateData()

	env := common.GetEnv("DB_HOST", "default")

	conf, err := config.LoadConfig(env)
	if err != nil {
		panic(err)
	}
	GlobalConfig = *conf
}
