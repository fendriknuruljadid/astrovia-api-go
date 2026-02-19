package progress

import (
	"encoding/json"
	"log"

	"app/internal/packages/redis"
	ws "app/internal/packages/websocket"
	"app/internal/services/v1/astro-zenith/auto-clip/repository"
)

type ProgressEvent struct {
	VideoID string                 `json:"video_id"`
	Stage   string                 `json:"stage"`
	Percent int                    `json:"percent"`
	Message string                 `json:"message"`
	Meta    map[string]interface{} `json:"meta"`
}

func StartConsumer() {
	pubsub := redis.Rdb.PSubscribe(redis.Ctx, "progress:*")
	ch := pubsub.Channel()

	log.Println("[progress] redis consumer started")

	for msg := range ch {
		var evt ProgressEvent

		if err := json.Unmarshal([]byte(msg.Payload), &evt); err != nil {
			log.Println("invalid payload:", err)
			continue
		}

		if evt.VideoID == "" {
			continue
		}
		log.Println("RAW PAYLOAD:", msg.Payload)

		progress := repository.VideoProgress{
			Stage:   evt.Stage,
			Percent: evt.Percent,
			Message: evt.Message,
		}

		// if evt.Percent%5 == 0 || evt.Percent == 100 {
		go func() {
			if err := repository.UpdateVideoProgress(evt.VideoID, progress); err != nil {
				log.Println("update progress failed:", err)
			}
		}()
		// }
		// Update DB
		// if err := repository.UpdateVideoProgress(
		//     evt.VideoID,
		//     evt.Stage,
		//     evt.Percent,
		//     evt.Message,
		// ); err != nil {
		//     log.Println("update progress failed:", err)
		// }
		// log.Println("update progress:", "Update Progress Video")
		log.Println("progres: ", evt)
		// OPTIONAL: republish untuk WS gateway
		// redis.Rdb.Publish(
		//     redis.Ctx,
		//     "ws:"+evt.VideoID,
		//     msg.Payload,
		// )
		ws.ProgressHub.Broadcast(evt.VideoID, []byte(msg.Payload))
	}
}
