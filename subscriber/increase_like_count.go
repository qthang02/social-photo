package subscriber

import (
	"context"
	"gorm.io/gorm"
	"log"
	"social-photo/common"
	"social-photo/modules/post/storage"
	"social-photo/pubsub"
)

type HasItemId interface {
	GetItemId() int
}

func IncreaseLikeCount(ctx context.Context, db *gorm.DB, ps pubsub.PubSub) {
	c, _ := ps.Subscribe(ctx, common.TopicUserLikedPost)

	go func() {
		// defer common.Recovery()
		for msg := range c {
			data := msg.Data().(HasItemId)

			if err := storage.NewSQLStore(db).IncreaseLikeCount(ctx, data.GetItemId()); err != nil {
				log.Println(err)
			}
		}
	}()
}
