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

var EmptyIP RemoteIP = RemoteIP{
	IpAddr: "",
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

	// Check if any of these headers are contained in the request.
	headers := c.Request.Header

	if len(headers) > 0 {
		checklist := []string{
			"x-client-ip",         // Standard headers used by Amazon EC2, Heroku, and others.
			"cf-connecting-ip",    // Cloudflare
			"fastly-client-ip",    // Fastly and Firebase 
			"true-client-ip",      // Akamai and Cloudflare 
			"x-real-ip",           // Default nginx proxy/fcgi
			"x-cluster-client-ip", // (Rackspace LB and Riverbed's Stingray)
			"x-forwarded",
			"forwarded-for",
			"forwarded",
		}

		for _, h := range checklist {
			if ip := c.Request.Header.Get(h); net.ParseIP(ip) != nil {
				return RemoteIP{IpAddr: ip,}, nil
			}
		}
	}

	// if IP is not in headers, take from logs
	remoteIp, _, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err != nil {
		return EmptyIP, err
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

	return EmptyIP, errors.New("IP not found")
}
