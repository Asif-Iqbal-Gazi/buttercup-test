package ble

import (
	"github.com/buttercup-test/gatt"
)

var defaultBLEClientOptions = []gatt.Option{
	gatt.MacDeviceRole(gatt.CentralManager),
}

/*

var defaultBLEServerOptions = []gatt.Option{
	gatt.MacDeviceRole(gatt.PeripheralManager),
}

*/
