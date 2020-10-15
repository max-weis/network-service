package host

import (
	"context"
	"github.com/Ullaakut/nmap"
	ns "gitlab.com/M4xwell/network-service"
	"gitlab.com/M4xwell/network-service/pkg/logger"
	"time"
)

type Job struct {
	Repository ns.Repository
	Logger     logger.Logger
	Target     string
}

func (j *Job) FindHosts() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	j.Logger.Info("Scanning hosts")

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(j.Target),
		nmap.WithContext(ctx),
	)
	if err != nil {
		j.Logger.Errorf("unable to create nmap scanner: %v", err)
	}

	result, warnings, err := scanner.Run()
	if err != nil {
		j.Logger.Warnf("unable to run nmap scan: %v", err)
	}

	if warnings != nil {
		j.Logger.Errorf("Warnings: %v", warnings)
	}

	for _, host := range result.Hosts {
		err := j.Repository.SaveHosts(host)
		if err != nil {
			j.Logger.Errorf("could not save host: %v", err)
			return err
		}
	}

	j.Logger.Infof("saved: %d hosts", len(result.Hosts))

	return nil
}
