package grubber

import (
	"errors"
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

type RemoteIP struct {
	IpAddr string `json:"ip_address"`
}

func GetRemoteIP(c *gin.Context) (RemoteIP, error) {
	ips := c.Request.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		netIP := net.ParseIP(splitIps[len(splitIps)-1]) // last if from proxy
		if netIP != nil {
			return RemoteIP{
				IpAddr: netIP.String(),
			}, nil
		}
	}

	none := RemoteIP{
		IpAddr: "",
	}

	remoteIp, _, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err != nil {
		return none, err
	}

	netIP := net.ParseIP(remoteIp)
	if netIP != nil {
		if remoteIp == "::1" {
			return RemoteIP{
				IpAddr: "127.0.0.1",
			}, nil
		}
		return RemoteIP{
			IpAddr: remoteIp,
		}, nil
	}

	return none, errors.New("IP not found")
}
