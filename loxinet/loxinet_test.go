/*
 * Copyright (c) 2022 NetLOX Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package loxinet

import (
	"fmt"
	"net"
	"testing"
)

type Tk struct {
}

func TestMain(t *testing.T) {

	fmt.Printf("LoxiLB Unit-Test \n")
	loxiNetInit()

	ifmac := [6]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6}
	_, err := mh.zr.Ports.PortAdd("hs0", 12, PORT_REAL, ROOT_ZONE,
		PortHwInfo{ifmac, true, true, 1500, "", "", 0},
		PortLayer2Info{false, 10})

	if err != nil {
		t.Errorf("failed to add port %s:%s", "hs0", err)
	}

	p := mh.zr.Ports.PortFindByName("hs0")
	if p == nil {
		t.Errorf("failed to add port %s", "hs0")
	}

	ifmac = [6]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x7}
	_, err = mh.zr.Ports.PortAdd("bond1", 15, PORT_BOND, ROOT_ZONE,
		PortHwInfo{ifmac, true, true, 1500, "", "", 0},
		PortLayer2Info{false, 10})

	if err != nil {
		t.Errorf("failed to add port %s", "bond1")
	}

	p = mh.zr.Ports.PortFindByName("bond1")
	if p == nil {
		t.Errorf("failed to add port %s", "bond1")
	}

	_, err = mh.zr.Ports.PortAdd("hs1", 15, PORT_REAL, ROOT_ZONE,
		PortHwInfo{ifmac, true, true, 1500, "", "", 0},
		PortLayer2Info{false, 10})
	if err != nil {
		t.Errorf("failed to add port hs1")
	}

	_, err = mh.zr.Ports.PortAdd("hs1", 15, PORT_BONDSIF, ROOT_ZONE,
		PortHwInfo{ifmac, true, true, 1500, "bond1", "", 0},
		PortLayer2Info{false, 10})
	if err != nil {
		t.Errorf("failed to add port hs1 to bond1")
	}

	ifmac = [6]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x8}
	_, err = mh.zr.Ports.PortAdd("hs2", 100, PORT_REAL, ROOT_ZONE,
		PortHwInfo{ifmac, true, true, 1500, "", "", 0},
		PortLayer2Info{false, 10})
	if err != nil {
		t.Errorf("failed to add port %s", "hs2")
	}

	ifmac = [6]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0xa}
	_, err = mh.zr.Vlans.VlanAdd(100, "vlan100", ROOT_ZONE, 124,
		PortHwInfo{ifmac, true, true, 1500, "", "", 0})
	if err != nil {
		t.Errorf("failed to add port %s", "vlan100")
	}

	p = mh.zr.Ports.PortFindByName("vlan100")
	if p == nil {
		t.Errorf("failed to add port %s", "vlan100")
	}

	_, err = mh.zr.Vlans.VlanPortAdd(100, "hs0", false)
	if err != nil {
		t.Errorf("failed to add port %s to vlan %d", "hs0", 100)
	}

	_, err = mh.zr.Vlans.VlanPortAdd(100, "hs0", true)
	if err != nil {
		t.Errorf("failed to add tagged port %s to vlan %d", "hs0", 100)
	}

	_, err = mh.zr.L3.IfaAdd("vlan100", "21.21.21.1/24")
	if err != nil {
		t.Errorf("failed to add l3 ifa to vlan%d", 100)
	}

	_, err = mh.zr.L3.IfaAdd("hs0", "11.11.11.1/32")
	if err != nil {
		t.Errorf("failed to add l3 ifa to hs0")
	}
	fmt.Printf("#### Interface List ####\n")
	mh.zr.Ports.Ports2String(&mh)
	fmt.Printf("#### IFA List ####\n")
	mh.zr.L3.Ifas2String(&mh)

	_, err = mh.zr.L3.IfaDelete("vlan100", "21.21.21.1/24")
	if err != nil {
		t.Errorf("failed to delete l3 ifa from vlan100")
	}

	fmt.Printf("#### IFA List ####\n")
	mh.zr.L3.Ifas2String(&mh)

	fmt.Printf("#### Vlan List ####\n")
	mh.zr.Vlans.Vlans2String(&mh)

	_, err = mh.zr.Vlans.VlanPortDelete(100, "hs0", false)
	if err != nil {
		t.Errorf("failed to delete hs0 from from vlan100")
	}
	_, err = mh.zr.Vlans.VlanPortDelete(100, "hs0", true)
	if err != nil {
		t.Errorf("failed to delete tagged hs0 from from vlan100")
	}
	_, err = mh.zr.Vlans.VlanDelete(100)
	if err != nil {
		t.Errorf("failed to delete vlan100")
	}

	ifmac = [6]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0xa}
	_, err = mh.zr.Vlans.VlanAdd(100, "vlan100", ROOT_ZONE, 124,
		PortHwInfo{ifmac, true, true, 1500, "", "", 0})
	if err != nil {
		t.Errorf("failed to add port %s", "vlan100")
	}

	fdbKey := FdbKey{[6]byte{0x05, 0x04, 0x03, 0x3, 0x1, 0x0}, 100}
	fdbAttr := FdbAttr{"hs0", net.ParseIP("0.0.0.0"), FDB_VLAN}

	_, err = mh.zr.L2.L2FdbAdd(fdbKey, fdbAttr)
	if err != nil {
		t.Errorf("failed to add fdb hs0:vlan100")
	}
	_, err = mh.zr.L2.L2FdbAdd(fdbKey, fdbAttr)
	if err == nil {
		t.Errorf("added duplicate fdb vlan100")
	}

	fdbKey1 := FdbKey{[6]byte{0xb, 0xa, 0x9, 0x8, 0x7, 0x6}, 100}
	fdbAttr1 := FdbAttr{"hs2", net.ParseIP("0.0.0.0"), FDB_VLAN}

	_, err = mh.zr.L2.L2FdbAdd(fdbKey1, fdbAttr1)
	if err != nil {
		t.Errorf("failed to add fdb hs2:vlan100")
	}

	_, err = mh.zr.L2.L2FdbDel(fdbKey1)
	if err != nil {
		t.Errorf("failed to del fdb hs2:vlan100")
	}

	_, err = mh.zr.L2.L2FdbDel(fdbKey1)
	if err == nil {
		t.Errorf("deleted non-existing fdb hs2:vlan100")
	}

	hwmac, _ := net.ParseMAC("00:00:00:00:00:01")
	_, err = mh.zr.Nh.NeighAdd(net.IPv4(8, 8, 8, 8), "default", NeighAttr{12, 1, hwmac})
	if err != nil {
		t.Errorf("NHAdd fail 8.8.8.8")
	}

	hwmac1, _ := net.ParseMAC("00:00:00:00:00:00")
	_, err = mh.zr.Nh.NeighAdd(net.IPv4(10, 10, 10, 10), "default", NeighAttr{12, 1, hwmac1})
	if err != nil {
		t.Errorf("NHAdd fail 10.10.10.10")
	}

	route := net.IPv4(1, 1, 1, 1)
	mask := net.CIDRMask(24, 32)
	route = route.Mask(mask)
	ipnet := net.IPNet{IP: route, Mask: mask}
	ra := RtAttr{0, 0, false}
	na := []RtNhAttr{{net.IPv4(8, 8, 8, 8), 12}}
	_, err = mh.zr.Rt.RtAdd(ipnet, "default", ra, na)
	if err != nil {
		t.Errorf("NHAdd fail 1.1.1.1/24 via 8.8.8.8")
	}

	_, err = mh.zr.Nh.NeighDelete(net.IPv4(8, 8, 8, 8), "default")
	if err != nil {
		t.Errorf("NHAdd fail 8.8.8.8")
	}

	_, err = mh.zr.Rt.RtDelete(ipnet, "default")
	if err != nil {
		t.Errorf("RouteDel fail 1.1.1.1/24 via 8.8.8.8")
	}

	fmt.Printf("#### Route-List ####\n")
	mh.zr.Rt.Rts2String(&mh)

	fmt.Printf("#### NH-List ####\n")
	mh.zr.Nh.Neighs2String(&mh)

	fmt.Printf("#### Trie-List1 ####\n")
	mh.zr.Rt.Trie4.Trie2String(mh.zr.Rt)

}