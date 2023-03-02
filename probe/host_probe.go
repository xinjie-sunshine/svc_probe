package probe

import (
	"net"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Host_probe() {
	// Load configuration from file
	viper.Reset()
	viper.SetConfigName("host_cnf")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to read configuration file")
	}

	// Initialize logger
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// logrus.SetOutput(os.Stdout)
	logFile, err := os.OpenFile("probe.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to open log file")
	}
	logrus.SetOutput(logFile)
	logrus.SetLevel(logrus.InfoLevel)

	hostsAndPorts := viper.GetStringSlice("hosts_and_ports")
	// log.Print(hostsAndPorts)
	timeout := time.Duration(viper.GetInt("timeout")) * time.Second
	// timeout := time.Duration(5) * time.Second

	for _, hostAndPort := range hostsAndPorts {
		conn, err := net.DialTimeout("tcp", hostAndPort, timeout)
		if err != nil {
			logrus.WithError(err).WithField("host", hostAndPort).Error("Connection failed")
		} else {
			conn.Close()
			logrus.WithField("host", hostAndPort).Info("Connection succeeded")
		}
	}

}
