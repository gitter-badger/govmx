// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package vmx

import (
	"fmt"
	"testing"
)

var data = `.encoding = "UTF-8"
annotation = "Terraform VMWARE VIX test"
bios.bootorder = "hdd,CDROM"
checkpoint.vmstate = ""
cleanshutdown = "TRUE"
config.version = "8"
cpuid.corespersocket = "1"
displayname = "core01"
ehci.pcislotnumber = "-1"
ehci.present = "FALSE"
ethernet1.addressType = "generated"
ethernet1.connectionType = "bridged"
ethernet1.linkStatePropagation.enable = "true"
ethernet1.present = "TRUE"
ethernet1.startConnected = "true"
ethernet1.virtualDev = "e1000"
ethernet1.wakeOnPcktRcv = "false"
ethernet2.address = "00:50:56:aa:bb:cc"
ethernet2.addressType = "static"
ethernet2.connectionType = "nat"
ethernet2.present = "TRUE"
ethernet2.startConnected = "true"
ethernet2.virtualDev = "e1000"
ethernet2.wakeOnPcktRcv = "false"
ethernet3.addressType = "generated"
ethernet3.connectionType = "hostonly"
ethernet3.present = "TRUE"
ethernet3.startConnected = "true"
ethernet3.virtualDev = "e1000"
ethernet3.wakeOnPcktRcv = "false"
extendedconfigfile = "core01.vmxf"
floppy0.present = "FALSE"
guestos = "other3xlinux-64"
gui.fullscreenatpoweron = "FALSE"
gui.viewmodeatpoweron = "windowed"
hgfs.linkrootshare = "TRUE"
hgfs.maprootshare = "TRUE"
ide1:0.devicetype = "cdrom-image"
ide1:0.filename = "/Users/camilo/Dropbox/Development/cloudescape/dobby-boxes/coreos/packer_cache/e159f7e70f4ccc346ee76b2a32cbdf059549a7ca82e91edbeed5d747bcdd50f9.iso"
ide1:0.present = "TRUE"
ide1:1.devicetype = "cdrom-image"
ide1:1.filename = "coreos.iso"
ide1:1.present = "TRUE"
isolation.tools.hgfs.disable = "FALSE"
memsize = "1024"
monitor.phys_bits_used = "40"
msg.autoanswer = "true"
numvcpus = "1"
nvram = "core01.nvram"
pcibridge0.pcislotnumber = "17"
pcibridge0.present = "TRUE"
pcibridge4.functions = "8"
pcibridge4.pcislotnumber = "21"
pcibridge4.present = "TRUE"
pcibridge4.virtualdev = "pcieRootPort"
pcibridge5.functions = "8"
pcibridge5.pcislotnumber = "22"
pcibridge5.present = "TRUE"
pcibridge5.virtualdev = "pcieRootPort"
pcibridge6.functions = "8"
pcibridge6.pcislotnumber = "23"
pcibridge6.present = "TRUE"
pcibridge6.virtualdev = "pcieRootPort"
pcibridge7.functions = "8"
pcibridge7.pcislotnumber = "24"
pcibridge7.present = "TRUE"
pcibridge7.virtualdev = "pcieRootPort"
policy.vm.mvmtid = ""
powertype.poweroff = "soft"
powertype.poweron = "soft"
powertype.reset = "soft"
powertype.suspend = "soft"
proxyapps.publishtohost = "FALSE"
remotedisplay.vnc.enabled = "TRUE"
remotedisplay.vnc.port = "5919"
replay.filename = ""
replay.supported = "FALSE"
scsi0.pcislotnumber = "16"
scsi0.present = "TRUE"
scsi0.virtualdev = "lsilogic"
scsi0:0.filename = "disk-cl1.vmdk"
scsi0:0.present = "TRUE"
scsi0:0.redo = ""
scsi0:1.present = "TRUE"
scsi0:1.autodetect = "TRUE"
scsi0:1.deviceType = "cdrom-raw"
softPowerOff = "FALSE"
sound.startconnected = "FALSE"
tools.synctime = "TRUE"
tools.upgrade.policy = "upgradeAtPowerCycle"
usb:1.speed = "2"
usb:1.present = "TRUE"
usb:1.deviceType = "hub"
usb:1.port = "1"
usb:1.parent = "-1"
usb:0.present = "TRUE"
usb:0.deviceType = "hid"
usb:0.port = "0"
usb:0.parent = "-1"
uuid.action = "create"
uuid.bios = "56 4d 59 1a 1a 9b 5f d8-29 6c 70 d0 bf 20 41 99"
uuid.location = "56 4d 59 1a 1a 9b 5f d8-29 6c 70 d0 bf 20 41 99"
vc.uuid = ""
virtualhw.productcompatibility = "hosted"
virtualhw.version = "9"
vmci0.id = "1861462627"
vmci0.pcislotnumber = "35"
vmci0.present = "TRUE"
vmotion.checkpointfbsize = "67108864"
ethernet1.pciSlotNumber = "32"
ethernet2.pciSlotNumber = "33"
ethernet3.pciSlotNumber = "34"
ethernet1.generatedAddress = "00:0c:29:20:41:a3"
ethernet1.generatedAddressOffset = "10"
ethernet3.generatedAddress = "00:0c:29:20:41:b7"
ethernet3.generatedAddressOffset = "30"
`

func TestUnmarshal(t *testing.T) {
	type Vhardware struct {
		Version int    `vmx:"version"`
		Compat  string `vmx:"productCompatibility"`
	}

	type Ethernet struct {
		StartConnected       bool   `vmx:"startConnected"`
		Present              bool   `vmx:"present"`
		ConnectionType       string `vmx:"connectionType"`
		VirtualDev           string `vmx:"virtualDev"`
		WakeOnPcktRcv        bool   `vmx:"wakeOnPcktRcv"`
		AddressType          string `vmx:"addressType"`
		LinkStatePropagation bool   `vmx:"linkStatePropagation.enable,omitempty"`
	}

	type SATADevice struct {
		Present  bool   `vmx:"present,omitempty"`
		Type     string `vmx:"devicetype,omitempty"`
		Filename string `vmxl:"filename,omitempty"`
	}

	type SCSIDevice struct {
		Present    bool   `vmx:"present"`
		PCISlot    uint   `vmx:"pcislotnumber,omitempty"`
		VirtualDev string `vmx:"virtualdev,omitempty"`
		Type       string `vmx:"devicetype"`
		Filename   string `vmxl:"filename,omitempty"`
	}

	type IDEDevice struct {
		Present  bool   `vmx:"present,omitempty"`
		Type     string `vmx:"devicetype,omitempty"`
		Filename string `vmxl:"filename,omitempty"`
	}

	type USBDevice struct {
		Present bool   `vmx:"present,omitempty"`
		Speed   uint   `vmx:"speed,omitempty"`
		Type    string `vmx:"devicetype,omitempty"`
		Port    uint   `vmx:"port,omitempty"`
		Parent  string `vmx:"parent,omitmepty"`
	}

	type PowerType struct {
		PowerOff string `vmx:"poweroff,omitempty"`
		PowerOn  string `vmx:"poweron,omitempty"`
		Reset    string `vmx:"reset,omitempty"`
		Suspend  string `vmx:"suspend,omitempty"`
	}

	type VM struct {
		Encoding    string       `vmx:".encoding"`
		PowerType   PowerType    `vmx:"powertype"`
		Annotation  string       `vmx:"annotation"`
		Vhardware   Vhardware    `vmx:"virtualhw"`
		Memsize     uint         `vmx:"memsize"`
		Numvcpus    uint         `vmx:"numvcpus"`
		MemHotAdd   bool         `vmx:"mem.hotadd"`
		DisplayName string       `vmx:"displayName"`
		GuestOS     string       `vmx:"guestOS"`
		Autoanswer  bool         `vmx:"msg.autoAnswer"`
		Ethernet    []Ethernet   `vmx:"ethernet"`
		IDEDevices  []IDEDevice  `vmx:"ide"`
		SCSIDevices []SCSIDevice `vmx:"scsi"`
		SATADevices []SATADevice `vmx:"sata"`
		USBDevices  []USBDevice  `vmx:"usb"`
	}

	vm := new(VM)
	err := Unmarshal([]byte(data), vm)
	ok(t, err)
	//fmt.Printf("%+v\n", vm.PowerType)
	assert(t, vm.Vhardware.Version == 9, "vhwversion should be 9")
	assert(t, len(vm.Ethernet) == 3, "there should be 3 ethernet devices")
	assert(t, vm.Numvcpus == 1, "there should be 1 vcpu")
	// fmt.Printf("%+v\n", vm.IDEDevices)
	// fmt.Printf("%+v\n", vm.SCSIDevices)
	//fmt.Printf("%+v\n", vm.USBDevices)
	assert(t, len(vm.IDEDevices) == 2, fmt.Sprintf("there should be 2 IDE devices, found %d", len(vm.IDEDevices)))
	assert(t, len(vm.SCSIDevices) == 3, fmt.Sprintf("there should be 3 SCSI devices, found %d", len(vm.SCSIDevices)))
	assert(t, len(vm.SATADevices) == 0, fmt.Sprintf("there should be 0 SATA controller, found %d", len(vm.SATADevices)))
	assert(t, len(vm.USBDevices) == 2, fmt.Sprintf("there should be 2 USB devices, found %d", len(vm.USBDevices)))
}
