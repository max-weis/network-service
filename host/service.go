package host

import (
	"github.com/Ullaakut/nmap"
	ns "gitlab.com/M4xwell/network-service"
	"gitlab.com/M4xwell/network-service/pkg/logger"
)

type Service struct {
	Repository ns.Repository
	Logger     logger.Logger
}

func (s *Service) ListHosts() ([]nmap.Host, error) {
	hosts, err := s.Repository.ListHosts()
	if err != nil {
		s.Logger.Errorf("could not list hosts: %v", err)
		return nil, err
	}

	return hosts, nil
}

func (s *Service) FindHostByName(name string) (nmap.Host, error) {
	host, err := s.Repository.FindHostByName(name)
	if err != nil {
		s.Logger.Errorf("could not find host: %v", err)
		return nmap.Host{}, err
	}

	return host, nil
}
