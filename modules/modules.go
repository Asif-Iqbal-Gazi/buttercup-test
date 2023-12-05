package modules

import (
	"github.com/buttercup-test/bettercap/modules/any_proxy"
	"github.com/buttercup-test/bettercap/modules/api_rest"
	"github.com/buttercup-test/bettercap/modules/arp_spoof"
	"github.com/buttercup-test/bettercap/modules/ble"
	"github.com/buttercup-test/bettercap/modules/caplets"
	"github.com/buttercup-test/bettercap/modules/dhcp6_spoof"
	"github.com/buttercup-test/bettercap/modules/dns_spoof"
	"github.com/buttercup-test/bettercap/modules/events_stream"
	"github.com/buttercup-test/bettercap/modules/gps"
	"github.com/buttercup-test/bettercap/modules/hid"
	"github.com/buttercup-test/bettercap/modules/http_proxy"
	"github.com/buttercup-test/bettercap/modules/http_server"
	"github.com/buttercup-test/bettercap/modules/https_proxy"
	"github.com/buttercup-test/bettercap/modules/https_server"
	"github.com/buttercup-test/bettercap/modules/mac_changer"
	"github.com/buttercup-test/bettercap/modules/mdns_server"
	"github.com/buttercup-test/bettercap/modules/mysql_server"
	"github.com/buttercup-test/bettercap/modules/net_probe"
	"github.com/buttercup-test/bettercap/modules/net_recon"
	"github.com/buttercup-test/bettercap/modules/net_sniff"
	"github.com/buttercup-test/bettercap/modules/packet_proxy"
	"github.com/buttercup-test/bettercap/modules/syn_scan"
	"github.com/buttercup-test/bettercap/modules/tcp_proxy"
	"github.com/buttercup-test/bettercap/modules/ticker"
	"github.com/buttercup-test/bettercap/modules/ui"
	"github.com/buttercup-test/bettercap/modules/update"
	"github.com/buttercup-test/bettercap/modules/wifi"
	"github.com/buttercup-test/bettercap/modules/wol"

	"github.com/buttercup-test/bettercap/session"
)

func LoadModules(sess *session.Session) {
	sess.Register(any_proxy.NewAnyProxy(sess))
	sess.Register(arp_spoof.NewArpSpoofer(sess))
	sess.Register(api_rest.NewRestAPI(sess))
	sess.Register(ble.NewBLERecon(sess))
	sess.Register(dhcp6_spoof.NewDHCP6Spoofer(sess))
	sess.Register(net_recon.NewDiscovery(sess))
	sess.Register(dns_spoof.NewDNSSpoofer(sess))
	sess.Register(events_stream.NewEventsStream(sess))
	sess.Register(gps.NewGPS(sess))
	sess.Register(http_proxy.NewHttpProxy(sess))
	sess.Register(http_server.NewHttpServer(sess))
	sess.Register(https_proxy.NewHttpsProxy(sess))
	sess.Register(https_server.NewHttpsServer(sess))
	sess.Register(mac_changer.NewMacChanger(sess))
	sess.Register(mysql_server.NewMySQLServer(sess))
	sess.Register(mdns_server.NewMDNSServer(sess))
	sess.Register(net_sniff.NewSniffer(sess))
	sess.Register(packet_proxy.NewPacketProxy(sess))
	sess.Register(net_probe.NewProber(sess))
	sess.Register(syn_scan.NewSynScanner(sess))
	sess.Register(tcp_proxy.NewTcpProxy(sess))
	sess.Register(ticker.NewTicker(sess))
	sess.Register(wifi.NewWiFiModule(sess))
	sess.Register(wol.NewWOL(sess))
	sess.Register(hid.NewHIDRecon(sess))

	sess.Register(caplets.NewCapletsModule(sess))
	sess.Register(update.NewUpdateModule(sess))
	sess.Register(ui.NewUIModule(sess))
}
