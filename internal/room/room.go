package room

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

var rooms = make(map[string]*Room)

// message types
const (
	someoneJoined = "SOMEONE_JOINED"
	musicPlayed   = "MUSIC_PLAYED"
	roomShutdown  = "ROOM_SHUTDOWN"
)

type Connection struct {
	UserID string
	Conn   *websocket.Conn
}

type Room struct {
	ServerURL          string
	Id                 string
	Clients            map[*Connection]bool
	Broadcast          chan *Message
	mutex              sync.Mutex
	CurrentPlayedMusic string
}

type RoomInfo struct {
	Id                 string
	ClientsNo          int
	CurrentPlayedMusic string
	Clients            []string
}

type Message struct {
	Type string

	UserJoinedUsername string

	MusicUrl string
	ByWho    string
}

func (m *Message) String() string {
	msg := fmt.Sprintf("[%s]", m.Type)
	if m.Type == musicPlayed {
		msg += fmt.Sprintf(" music url: %s, played by: %s", m.MusicUrl, m.ByWho)
	}
	if m.Type == someoneJoined {
		msg += fmt.Sprintf(" user joined: %s", m.UserJoinedUsername)
	}
	return msg
}

func New() *Room {
	room := &Room{
		Id:        uuid.New().String(),
		Clients:   make(map[*Connection]bool),
		Broadcast: make(chan *Message),
	}
	rooms[room.Id] = room
	return room
}

func GetRoomInfo(roomId string) (*RoomInfo, error) {
	room, exist := rooms[roomId]
	if !exist {
		return nil, errors.New("room doesn't exist")
	}
	clientsUsernames := make([]string, 0)
	for client := range room.Clients {
		clientsUsernames = append(clientsUsernames, client.UserID)
	}

	return &RoomInfo{
		Id:                 roomId,
		ClientsNo:          len(room.Clients),
		CurrentPlayedMusic: room.CurrentPlayedMusic,
		Clients:            clientsUsernames,
	}, nil
}

// Shutdown the room and close all connections
func (room *Room) Shutdown() {
	room.mutex.Lock()

	// Close all connections
	for conn := range room.Clients {
		err := conn.Conn.Close()
		if err != nil {
			log.Error().Msgf("Error closing connection for user %s: %v", conn.UserID, err)
		}
		delete(room.Clients, conn)
	}

	// delete the room
	delete(rooms, room.Id)
}

// joins the music room
func (room *Room) Join(userId string) error {
	room.mutex.Lock()
	// check if user is already connected
	for conn := range room.Clients {
		if userId == conn.UserID {
			// user already in the room
			log.Info().Msgf("user %s already in the room %s\n", userId, room.Id)
			room.mutex.Unlock()
			return nil
		}
	}

	// create a new socket connection
	u, err := url.Parse(room.ServerURL)
	if err != nil {
		return err
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	roomConnection := &Connection{
		UserID: userId,
		Conn:   conn,
	}

	room.Clients[roomConnection] = true
	room.mutex.Unlock()

	room.Broadcast <- &Message{
		Type:               someoneJoined,
		UserJoinedUsername: userId, // TODO: fix this later
	}

	return nil
}

// play music will send a message with the music url to the broadcast channel that will be published to all room clients
func (room *Room) PlayMusic(musicUrl string) error {
	room.mutex.Lock()
	room.CurrentPlayedMusic = musicUrl
	room.mutex.Unlock()

	newMessage := &Message{
		Type:     musicPlayed,
		MusicUrl: musicUrl,
		ByWho:    "", // TODO: leave it for now, will handle it later
	}
	room.Broadcast <- newMessage
	return nil
}

// starts the room and
// runs infinte loop that will wait for messages over the broadcast channle and publish it to all room clients
func (room *Room) Start(ctx context.Context) {
	log.Printf("room %s started with %d clients", room.Id, len(room.Clients))
	for {
		select {
		case <-ctx.Done():
			log.Printf("Room %s shutting down\n", room.Id)
			delete(rooms, room.Id)
			return
		case msg := <-room.Broadcast:
			log.Printf("Received message: %v", msg)

			if msg.Type == roomShutdown {
				log.Printf("Shutting down room %s, bye!", room.Id)
				return
			}

			room.mutex.Lock()
			for conn, connected := range room.Clients {
				if !connected {
					continue
				}

				err := conn.Conn.WriteJSON(msg)
				if err != nil {
					// if there is an error while sending the message to the client connection, then close the connection and remove the client from the room
					conn.Conn.Close()
					delete(room.Clients, conn)
				}
			}
			room.mutex.Unlock()

		}

	}

}
