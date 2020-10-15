package network_service

import (
	"github.com/Ullaakut/nmap"
)

type Service interface {
	ListHosts() ([]nmap.Host, error)
	FindHostByName(name string) (nmap.Host, error)
}

type Repository interface {
	ListHosts() ([]nmap.Host, error)
	FindHostByName(name string) (nmap.Host, error)
	SaveHosts(nmap.Host) error
}

type Job interface {
	FindHosts() error
}

