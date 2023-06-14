package subscriber

import (
	"context"
	"gorm.io/gorm"
	"log"
	"social-photo/common"
	"social-photo/common/asyncjob"
	"social-photo/pubsub"
)

type subJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type pbEngine struct {
	db *gorm.DB
	ps pubsub.PubSub
}

func NewEngine(db *gorm.DB, ps pubsub.PubSub) *pbEngine {
	return &pbEngine{db: db, ps: ps}
}

func (engine *pbEngine) Start() error {
	engine.startSubTopic(common.TopicUserLikedPost, true,
		IncreaseLikeCount(engine.db),
	)

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *pbEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, jobs ...subJob) error {
	ps := engine.ps

	c, _ := ps.Subscribe(context.Background(), topic)

	for _, item := range jobs {
		log.Println("Setup subscriber for:", item.Title)
	}

	getJobHandler := func(job *subJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("running job for ", job.Title, ". Value: ", message.Data())
			return job.Hld(ctx, message)
		}
	}

	go func() {
		for {
			msg := <-c

			jobHdlArr := make([]asyncjob.Job, len(jobs))

			for i := range jobs {
				jobHdl := getJobHandler(&jobs[i], msg)
				jobHdlArr[i] = asyncjob.NewJob(jobHdl, asyncjob.WithName(jobs[i].Title))
			}

			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
