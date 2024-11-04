package udp

import (
	"context"
	"github.com/ascenmmo/udp-server/internal/connection"
	"github.com/ascenmmo/udp-server/internal/service"
	memoryDB "github.com/ascenmmo/udp-server/internal/storage"
	"github.com/ascenmmo/udp-server/internal/utils"
	"github.com/rs/zerolog"
	"net"
	"runtime"
	"time"
)

const (
	bufferSize = 4096
)

type WorkerUDP struct {
	service   service.Service
	conn      *net.UDPConn
	chMsg     []chan ChanUDPMessage
	rateLimit utils.RateLimit
	logger    zerolog.Logger
	count     int
}

type ChanUDPMessage struct {
	client  connection.DataSender
	request []byte
}

func (w *WorkerUDP) Listener(ctx context.Context) error {
	defer w.conn.Close()

	buffer := make([]byte, bufferSize)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			n, clientAddr, err := w.conn.ReadFromUDP(buffer)
			if err != nil {
				w.logger.Error().Err(err).Msg("Listener ReadFromUDP")
				continue
			}

			if w.rateLimit.IsLimited(clientAddr.String()) {
				continue
			}

			if n == 0 {
				continue
			}

			err = w.handleConnection(clientAddr, buffer[:n])
			if err != nil {
				w.logger.Error().Err(err).Msg("Listener handleConnection")
				continue
			}
		}
	}
}

func (w *WorkerUDP) Sender(ctx context.Context) {
	go w.printer()
	for _, v := range w.chMsg {
		go w.sendWorker(ctx, v)
	}
}

func (w *WorkerUDP) handleConnection(clientAddr *net.UDPAddr, buf []byte) error {
	defer func() {
		if r := recover(); r != nil {
			w.logger.Error().Msgf("recover: %v", r)
		}
	}()

	ds := connection.DataSender(&connection.UDPConnection{
		ClientAddr: clientAddr,
		Conn:       w.conn,
	})

	select {
	case w.chMsg[w.counterChen()] <- ChanUDPMessage{
		client:  ds,
		request: buf,
	}:
	default:
		return nil //errors.New("counterChen is full")
	}

	return nil
}

func (w *WorkerUDP) sendWorker(ctx context.Context, ch chan ChanUDPMessage) {
	for {
		select {
		case <-ctx.Done():
			return
		case chMsg := <-ch:
			users, msg, err := w.service.GetUsersAndMessage(chMsg.client, chMsg.request)
			if err != nil {
				w.logger.Warn().Err(err).Msg("senderWorker GetUsersAndMessage")
				continue
			}
			if len(msg) == 0 || len(users) == 0 {
				continue
			}
			for _, user := range users {
				err = user.Connection.Write(msg)
				if err != nil {
					w.logger.Warn().Err(err).Interface("senderWorker WriteToUDP", user.ID)
					err := w.service.RemoveUser(chMsg.client, user.ID)
					if err != nil {
						w.logger.Warn().Err(err).Interface("senderWorker  RemoveUser", user.ID)
					}
				}
			}
		}
	}
}

func (w *WorkerUDP) counterChen() int {
	w.count = (w.count + 1) % len(w.chMsg)
	return w.count
}

func (w *WorkerUDP) printer() {
	ticker := time.NewTicker(time.Second * 1)
	counter := 0
	for range ticker.C {
		for _, v := range w.chMsg {
			counter += len(v)
		}
		w.logger.Info().Interface("msges in chan", counter)
		w.logger.Info().Interface("gorutins", runtime.NumGoroutine())
		counter = 0
	}
}

func NewWorkerUDP(addr string, service service.Service, rateLimit int, storage memoryDB.IMemoryDB, logger zerolog.Logger) (*WorkerUDP, error) {
	chs := make([]chan ChanUDPMessage, runtime.NumCPU())
	for i := range chs {
		chs[i] = make(chan ChanUDPMessage, rateLimit)
	}
	w := &WorkerUDP{
		service:   service,
		logger:    logger,
		chMsg:     chs,
		rateLimit: utils.NewRateLimit(rateLimit, storage),
	}

	addr = ":" + addr
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		w.logger.Error().Err(err).Msg("Listener ResolveUDPAddr")
		return nil, err
	}
	w.logger.Info().Msgf("Listener UDP Started on %s", addr)

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		w.logger.Error().Err(err).Msg("Listener ListenUDP")
		return nil, err
	}
	w.conn = conn

	return w, nil
}
