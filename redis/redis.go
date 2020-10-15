package redis

import (
	"encoding/json"
	"fmt"
	"github.com/Ullaakut/nmap"
	"github.com/go-redis/redis"
	"gitlab.com/M4xwell/network-service/pkg/logger"
)

type Repository struct {
	Logger logger.Logger
	Client redis.Client
}

func (r *Repository) ListHosts() ([]nmap.Host, error) {
	r.Logger.Info("find hosts")

	result, err := r.Client.Do( "keys", "*").Result()
	if err != nil {
		r.Logger.Errorf("could not get keys: %v", err)
		return nil, err
	}

	keys := result.([]interface{})

	var hosts []nmap.Host

	for _, key := range keys {
		var host nmap.Host

		keyString := fmt.Sprintf("%v", key)

		value, err := r.Client.Get(keyString).Result()
		if err != nil {
			r.Logger.Errorf("could not find host: %v", err)
			return nil, err
		}

		err = json.Unmarshal([]byte(value), &host)
		if err != nil {
			r.Logger.Errorf("could not decode host: %v", err)
			return nil, err
		}

		hosts = append(hosts, host)
	}

	return hosts, nil
}

func (r *Repository) FindHostByName(name string) (nmap.Host, error) {
	result, err := r.Client.Get(name).Result()
	if err != nil {
		r.Logger.Errorf("could not find host: %v", err)
		return nmap.Host{}, err
	}

	var host nmap.Host

	err = json.Unmarshal([]byte(result), &host)
	if err != nil {
		r.Logger.Errorf("could not decode host: %v", err)
		return nmap.Host{}, err
	}

	return host, nil
}

func (r *Repository) SaveHosts(host nmap.Host) error {
	hostname := host.Hostnames[0].Name

	payload, err := json.Marshal(host)
	if err != nil {
		r.Logger.Errorf("could not encode host: %v", err)
		return err
	}

	err = r.Client.Set(hostname, payload, 0).Err()
	if err != nil {
		r.Logger.Errorf("could not persist host: %v", err)
		return err
	}

	return nil
}
