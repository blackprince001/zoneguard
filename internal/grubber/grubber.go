package grubber

import (
	"errors"
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

type RemoteIP struct {
	Addr string
}

func GetRemoteIP(c *gin.Context) (RemoteIP, error) {
	ips := c.Request.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		netIP := net.ParseIP(splitIps[len(splitIps)-1]) // last if from proxy
		if netIP != nil {
			return RemoteIP{
				Addr: netIP.String(),
			}, nil
		}
	}

	none := RemoteIP{
		Addr: "",
	}

	remoteIp, _, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err != nil {
		return none, err
	}

	netIP := net.ParseIP(remoteIp)
	if netIP != nil {
		if remoteIp == "::1" {
			return RemoteIP{
				Addr: "127.0.0.1",
			}, nil
		}
		return RemoteIP{
			Addr: remoteIp,
		}, nil
	}

	return none, errors.New("IP not found")
}
