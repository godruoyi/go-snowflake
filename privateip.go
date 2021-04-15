package snowflake

import (
	"errors"
	"net"
)

// PrivateIPToMachineID convert private ip to machine id.
// From https://github.com/sony/sonyflake/blob/master/sonyflake.go
func PrivateIPToMachineID() uint16 {
	ip, err := lower16BitPrivateIP()
	if err != nil {
		return 0
	}

	return ip
}

//--------------------------------------------------------------------
// private function defined.
//--------------------------------------------------------------------

func privateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}

	return nil, errors.New("no private ip address")
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

func lower16BitPrivateIP() (uint16, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}

	// Snowflake macnineID max length is 10, max value is 1023
	// If ip[2] > 3, return ip[2] + ip[3], but:
	// 10.21.5.211 => return 5 + 211 = 216
	// 10.21.4.212 => return 4 + 212 = 216
	// need help if you can do this.
	if ip[2] > 3 {
		return uint16(ip[2] + ip[3]), nil
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}
