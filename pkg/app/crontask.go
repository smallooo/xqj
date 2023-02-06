package app

import (
	"github.com/robfig/cron"
	"log"
	"time"
	"xqj/models"
)

func RunTask() {
	log.Println("Starting cron task ...")

	c := cron.New()
	c.AddFunc("0 */1 * * *", func() {
		log.Println("RunTask models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("0 */1 * * *", func() {
		log.Println("RunTask models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	//c.AddFunc("0 */1 * * *", func() {
	//	log.Println("RunTask models.CleanAllArticle...")
	//	_, err := exec.Command("sh", "-c", "echo '"+"zaq123"+"' | sudo -S pkill -SIGINT main").Output()
	//
	//	if err != nil {
	//		log.Println("Process killed")
	//	} else {
	//		log.Println("Process killed failure...")
	//	}
	//})

	//c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
