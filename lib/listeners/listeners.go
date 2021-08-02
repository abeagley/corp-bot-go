package listeners

import "cloud.google.com/go/firestore"

type DiscordListener struct {
	Db *firestore.Client
}
