package subscriber

import (
	"context"
	"gorm.io/gorm"
	"social-photo/modules/post/storage"
	"social-photo/pubsub"
)

//func DecreaseLikeCount(ctx context.Context, db *gorm.DB, ps pubsub.PubSub) {
//	c, _ := ps.Subscribe(ctx, common.TopicUserLikedPost)
//
//	go func() {
//		// defer common.Recovery()
//		for msg := range c {
//			data := msg.Data().(HasItemId)
//
//			if err := storage.NewSQLStore(db).DecreaseLikeCount(ctx, data.GetItemId()); err != nil {
//				log.Println(err)
//			}
//		}
//	}()
//}

func DecreaseLikeCount(db *gorm.DB) subJob {
	return subJob{
		Title: "Decrease like count after user unlikes item",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			data := message.Data().(HasItemId)

			return storage.NewSQLStore(db).DecreaseLikeCount(ctx, data.GetItemId())
		},
	}
}
