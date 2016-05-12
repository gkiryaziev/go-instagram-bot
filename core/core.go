package core

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/jinzhu/gorm"

	// postgresql driver
	_ "github.com/lib/pq"

	"github.com/gkiryaziev/go-instagram-bot/conf"
	"github.com/gkiryaziev/go-instagram-bot/db"
	"github.com/gkiryaziev/go-instagram-bot/models"

	instagram "github.com/gkiryaziev/go-instagram"
)

// TextField struct
type TextField struct {
	userName string
	textType string
	comment  string
}

// Service struct
type Service struct {
	mdl []interface{}
	cfg *conf.Config
	db  *gorm.DB
}

// NewService return new CoreService object
func NewService() (*Service, error) {

	// read config file
	config, err := conf.NewConfig("config.yaml").Load()
	if err != nil {
		return nil, err
	}

	// open database connection
	database, err := db.Connect(
		config.Db.DbUser,
		config.Db.DbPassword,
		config.Db.DbHost,
		config.Db.DbPort,
		config.Db.DbName,
	)
	if err != nil {
		return nil, err
	}

	// models list
	modelsList := []interface{}{
		&models.Activity{}, // tbl_activity table
	}

	// create service object
	service := &Service{
		mdl: modelsList,
		cfg: config,
		db:  database,
	}

	return service, nil
}

// DropTables drop tables
func (cs *Service) DropTables() error {

	defer cs.db.Close()

	err := cs.db.DropTableIfExists(cs.mdl...).Error
	if err != nil {
		return err
	}

	log.Println("Drop Tables: success.")

	return nil
}

// Migrate migrate tables
func (cs *Service) Migrate() error {

	defer cs.db.Close()

	err := cs.db.AutoMigrate(cs.mdl...).Error
	if err != nil {
		return err
	}

	log.Println("Migration: success.")

	return nil
}

// Run service
func (cs *Service) Run() error {

	defer cs.db.Close()

	rand.Seed(time.Now().Unix())

	// interrupt
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// instagram connection pool
	var apiPool []*instagram.Instagram

	// open connection and append connections pool
	for _, user := range cs.cfg.Instagram.Users {
		api, err := instagram.NewInstagram(
			user.Username,
			user.Password,
		)
		if err != nil {
			return err
		}
		apiPool = append(apiPool, api)
	}

	// run goroutine
	for _, api := range apiPool {

		// random timeout
		randomTimeout := generateTimeout(
			cs.cfg.Instagram.TimeoutMin,
			cs.cfg.Instagram.TimeoutMax,
		)

		go cs.getActivity(api, randomTimeout)

		time.Sleep(500 * time.Millisecond)
	}

	// wait for terminating
	for {
		select {
		case <-interrupt:
			log.Println("Cleanup and terminating...")
			os.Exit(0)
		}
	}
}

// getActivity return activity
func (cs *Service) getActivity(api *instagram.Instagram, randomTimeout int) {

	// little log
	if cs.cfg.Debug {
		log.Println("Start getting with timeout:", randomTimeout, "ms.")
	}

	for {
		// get recent activity
		ract, err := api.GetRecentActivity()
		if err != nil {
			log.Fatal(err)
		}

		// fetch old stories
		for _, story := range ract.OldStories {
			cs.fetch(story)
		}

		// fetch new stories
		for _, story := range ract.NewStories {
			cs.fetch(story)
		}

		// sleep
		time.Sleep(time.Duration(randomTimeout) * time.Millisecond)
	}
}

// fetch data and fill database model
func (cs *Service) fetch(stories instagram.RecentActivityStories) {

	var activity models.Activity

	// parse text field
	txt := parseText(stories.Args.Text)

	act := &models.Activity{
		Pk:           stories.Pk, // instagram's post primary key from json
		UserID:       stories.Args.ProfileID,
		UserImageURL: stories.Args.ProfileImage,
		UserName:     txt.userName,
		Type:         txt.textType,
		Comment:      txt.comment,
	}

	// check if Args.Media have items
	if len(stories.Args.Media) > 0 {
		act.MediaID = stories.Args.Media[0].ID
		act.MediaURL = stories.Args.Media[0].Image
	}

	// write activity to DB
	if ok := cs.db.NewRecord(act); ok {

		// check by pk if record exist
		if cs.db.Where(&models.Activity{Pk: act.Pk}).First(&activity).RecordNotFound() {
			// create new record
			err := cs.db.Create(&act).Error
			if err != nil {
				log.Fatal(err)
			} else {
				// little log
				if cs.cfg.Debug {
					log.Println("Add row:", act.Pk)
				}
			}
		}
	}
}

// parseText parse Args.Text field
func parseText(text string) *TextField {

	txt := &TextField{
		userName: strings.Fields(text)[0],
	}

	switch {
	case strings.Contains(text, "liked your photo"):
		txt.textType = "liked_photo"
	case strings.Contains(text, "started following you"):
		txt.textType = "start_following"
	case strings.Contains(text, "took a photo of you"):
		txt.textType = "took_photo"
	case strings.Contains(text, "mentioned you in a comment:"):
		txt.textType = "mentioned"
		txt.comment = strings.Split(text, "mentioned you in a comment: ")[1]
	case strings.Contains(text, "commented:"):
		txt.textType = "commented"
		txt.comment = strings.Split(text, "commented: ")[1]
	}

	return txt
}

// generateTimeout return random timeout from 'min' to 'max' value
func generateTimeout(min, max int) int {
	return min + rand.Intn(max-min)
}
