package ip

import (
	"log"
	"net"
)

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() net.IP {
	// addrs, err := net.InterfaceAddrs()
	// if err != nil {
	// 	return ""
	// }
	// for _, address := range addrs {
	// 	// check the address type and if it is not a loopback the display it
	// 	if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	// 		if ipnet.IP.To4() != nil {
	// 			return ipnet.IP.String()
	// 		}
	// 	}
	// }
	// return ""

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
