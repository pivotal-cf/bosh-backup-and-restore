package bosh

import (
	"fmt"

	"github.com/cloudfoundry/bosh-cli/director"
	"github.com/pivotal-cf/pcf-backup-and-restore/backuper"
)

type DeployedInstance struct {
	director.Deployment
	JobName  string
	JobIndex string
	SSHConnection
	Logger
}

//go:generate counterfeiter -o fakes/fake_ssh_connection.go . SSHConnection
type SSHConnection interface {
	Run(cmd string) ([]byte, []byte, int, error)
	Cleanup() error
	Username() string
}

func NewBoshInstance(jobName, jobIndex string, connection SSHConnection, deployment director.Deployment, logger Logger) backuper.Instance {
	return DeployedInstance{
		JobIndex:      jobIndex,
		JobName:       jobName,
		SSHConnection: connection,
		Deployment:    deployment,
		Logger:        logger,
	}
}

func (d DeployedInstance) IsBackupable() (bool, error) {
	d.Logger.Debug("", "Checking instance %s %s has backup scripts", d.JobName, d.JobIndex)
	stdin, stdout, exitCode, err := d.Run("ls /var/vcap/jobs/*/bin/backup")

	d.Logger.Debug("", "Stdin: %s", string(stdin))
	d.Logger.Debug("", "Stdout: %s", string(stdout))

	if err != nil {
		d.Logger.Debug("", "Error checking instance has backup scripts. Exit code %d, error %s", exitCode, err.Error())
	}

	return exitCode == 0, err
}

func (d DeployedInstance) Backup() error {
	d.Logger.Debug("", "Running all backup scripts on instance %s %s", d.JobName, d.JobIndex)
	stdin, stdout, exitCode, err := d.Run("sudo mkdir -p /var/vcap/store/backup && ls /var/vcap/jobs/*/bin/backup | xargs -IN sudo sh -c N")

	d.Logger.Debug("", "Stdin: %s", string(stdin))
	d.Logger.Debug("", "Stdout: %s", string(stdout))

	if err != nil {
		d.Logger.Debug("", "Error running instance backup scripts. Exit code %d, error %s", exitCode, err.Error())
	}

	if exitCode != 0 {
		return fmt.Errorf("Instance backup scripts returned %d. Error: %s", exitCode, stdout)
	}

	return err
}

func (d DeployedInstance) Cleanup() error {
	d.Logger.Debug("", "Cleaning up SSH connection on instance %s %s", d.JobName, d.JobIndex)
	return d.CleanUpSSH(director.NewAllOrPoolOrInstanceSlug(d.JobName, d.JobIndex), director.SSHOpts{Username: d.SSHConnection.Username()})
}
