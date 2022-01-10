package dhcp

import (
	"errors"
	"fmt"
	"sync"

	"github.com/bugfan/to"
	"github.com/bugfan/wireguard-auth/env"
	"github.com/bugfan/wireguard-auth/models"
	"github.com/bugfan/wireguard-auth/utils"
)

var Default *dhcp

func init() {
	Default = NewDHCP(env.Get("dhcp_default_group"), env.Get("dhcp_default_ip"), env.GetInt64("dhcp_default_num"))
}

func NewDHCP(group, ipv4 string, count int64, args ...string) *dhcp {
	// cidr suffix
	suffix := "32"
	if len(args) > 0 {
		suffix = args[0]
	}
	// ip is a number
	ipn := utils.IPv4ToNumber(ipv4)
	if ipn < 0 {
		ipn = 167772161 //10.0.0.1
	}
	dhcp := &dhcp{
		count:  count,
		suffix: suffix,
		ip:     ipv4,
		group:  group,
	}
	models.InitSetting(map[string]string{
		group:                  to.String(0),
		dhcp.getCurrentIpKey(): ipv4,
		dhcp.getCurrentKey():   to.String(ipn),
		dhcp.getCountKey():     to.String(count),
		dhcp.getStartIpKey():   ipv4,
		dhcp.getStartKey():     to.String(ipn),
	})
	return dhcp
}

type dhcp struct {
	group, ip, suffix string
	count             int64
	sync.Mutex
}

func (s *dhcp) getKey() string {
	return s.group
}
func (s *dhcp) getCurrentIpKey() string {
	return s.group + "_current_ip"
}
func (s *dhcp) getCountKey() string {
	return s.group + "_count"
}
func (s *dhcp) getStartKey() string {
	return s.group + "_start"
}
func (s *dhcp) getStartIpKey() string {
	return s.group + "_start_ip"
}
func (s *dhcp) getCurrentKey() string {
	return s.group + "_current"
}
func (s *dhcp) GetSuffix() string {
	return s.suffix
}
func (s *dhcp) Get() (string, error) {
	s.Lock()
	defer s.Unlock()
	value := to.Int64(models.GetValue(s.group))
	if value > s.count {
		return "", errors.New("no ip can use")
	}
	currentKey := s.getCurrentKey()
	current := to.Int64(models.GetValue(currentKey))
	current++
	if err := models.SetValue(currentKey, to.String(current)); err != nil {
		return "", err
	}

	if err := models.SetValue(s.getCurrentIpKey(), utils.InetNtoA(current)); err != nil {
		return "", err
	}

	key := s.getKey()
	dhcpn := to.Int64(models.GetValue(key))
	dhcpn++
	if err := models.SetValue(key, to.String(dhcpn)); err != nil {
		return "", err
	}

	ipv4 := utils.InetNtoA(current)

	return ipv4, nil
}
func GetCIDR() string {
	ip := GetIP()
	if ip == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", ip, Default.GetSuffix())
}
func GetIP() string {
	ip, err := Default.Get()
	if err != nil {
		return ""
	}
	return ip
}
