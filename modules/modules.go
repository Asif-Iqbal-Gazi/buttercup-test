package modules

import (
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/any_proxy"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/api_rest"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/arp_spoof"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/ble"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/caplets"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/dhcp6_spoof"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/dns_spoof"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/events_stream"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/gps"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/hid"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/http_proxy"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/http_server"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/https_proxy"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/https_server"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/mac_changer"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/mdns_server"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/mysql_server"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/net_probe"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/net_recon"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/net_sniff"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/packet_proxy"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/syn_scan"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/tcp_proxy"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/ticker"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/ui"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/update"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/wifi"
	"github.com/Asif-Iqbal-Gazi/buttercup-test/modules/wol"

	"github.com/Asif-Iqbal-Gazi/buttercup-test/session"
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
