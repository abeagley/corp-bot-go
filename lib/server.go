package lib

import (
	"cloud.google.com/go/firestore"
	"corp_bot/lib/commands"
	"corp_bot/lib/listeners"
	dg "github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
)

// Server is our base struct for discordgo functionality
type Server struct {
	AppId            string
	FirebaseCredPath string
	Token            string

	db        *firestore.Client
	handlers  *commands.CommandHandler
	listeners *listeners.DiscordListener
}

// connectToDB attempts to connect to firebase through a
// credential file
func (s *Server) connectToDB() *firestore.Client {
	db, dbErr := newFirebaseDB(s.FirebaseCredPath)
	if dbErr != nil {
		log.Fatalf("Unable to establish connection to Firebase: %v", dbErr)
	}
	return db
}

// connectToSession opens our discordgo session and returns it
func (s *Server) connectToSession() *dg.Session {
	sess, err := dg.New("Bot " + s.Token)
	if err != nil {
		log.Fatalf("Unable to start discord session: %v", err)
	}

	sess.AddHandler(s.onSessionConnect)

	conErr := sess.Open()
	if conErr != nil {
		log.Fatalf("Cannot start CorpBot: %v", conErr)
	}

	return sess
}

// createTestCommands creates our slash commands on a specific guild server
// useful for testing as we don't have to wait for an hour when updating
// global slash commands
func (s *Server) createTestCommands(sess *dg.Session) {
	for _, c := range commands.Commands {
		_, err := sess.ApplicationCommandCreate(s.AppId, "797218711438098502", c)
		if err != nil {
			log.Printf("Unable to register application command: %v\n", err)
		}
	}
}

func (s *Server) onSessionCommand(sess *dg.Session, i *dg.InteractionCreate) {
	cmd := i.Interaction.Data.Name
	handlers := s.handlers.GetMap()
	if h, ok := handlers[cmd]; ok {
		h(sess, i)
	}
}

func (s *Server) onSessionConnect(sess *dg.Session, i *dg.Ready) {
	log.Println("Bot is up and running")
}

// StartBot attempts to initialize and start a discordgo session, this will block
func (s *Server) StartBot() {
	if s.AppId == "" || s.FirebaseCredPath == "" || s.Token == "" {
		log.Fatalf("Unable to start Corp Bot without required props")
	}

	s.db = s.connectToDB()
	sess := s.connectToSession()
	s.handlers = &commands.CommandHandler{Db: s.db}
	s.listeners = &listeners.DiscordListener{Db: s.db}

	// Use the below function to add commands to a specific guild for testing
	s.createTestCommands(sess) // uses guild commands for instant updates vs 1 hour on global

	// Start listening for slash commands
	sess.AddHandler(s.onSessionCommand)

	defer sess.Close()
	defer s.db.Close()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("\n\nShutting down...")
}
