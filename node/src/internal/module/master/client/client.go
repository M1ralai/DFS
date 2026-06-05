package client

import (
	"context"
	"time"

	"github.com/M1ralai/DFS/node/src/internal/module/master/dto"
	"github.com/M1ralai/DFS/node/src/utils/config"
	"github.com/go-resty/resty/v2"
)

type IMasterClient interface {
	Register(dto.RegisterRequest) error
	Ack(dto.AckRequest) error
}

type MasterClient struct {
	cfg    config.MasterConfig
	client *resty.Client
}

func NewMasterClient(cfg config.MasterConfig) IMasterClient {
	ret := &MasterClient{
		cfg:    cfg,
		client: resty.New(),
	}
	// TODO geri kalan implemente edildikten sonra hearthbeti burada baclat
	return ret
}

func (c *MasterClient) Register(req dto.RegisterRequest) error {
	_, err := c.client.R().
		SetBody(req).
		Post(c.cfg.MasterURL + "api/node")

	if err != nil {
		return err
	}
	return nil
}

func (c *MasterClient) Ack(req dto.AckRequest) error {
	_, err := c.client.R().
		SetBody(req).
		Post(c.cfg.MasterURL + "api/acknowledgement")
	if err != nil {
		return err
	}
	return nil
}

func (c *MasterClient) heartBeat(req dto.HeartBeatRequest, ctx context.Context) error {
	go func() {
		ticker := time.NewTicker(5 * time.Second) // TODO env dan gercek zamanlamayi cek
		for {
			select {
			case <-ticker.C:

				_, err := c.client.R().
					SetBody(req).
					Post(c.cfg.MasterURL + "api/heartbeat")
				if err != nil {
					return
				}

			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}
