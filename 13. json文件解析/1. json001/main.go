package main

import (
	"encoding/json"
	"fmt"
)

type trafficRule struct {
	TrafficRuleID  string          `json:"trafficRuleId"` // 1-63 bytes
	Priority       int16           `json:"priority"`      // 1~255
	Action         string          `json:"action"`        //DROP, ALLOW
	TrafficFilters []trafficFilter `json:"trafficFilter"`
}

type trafficFilter struct {
	//trafficFilterID string
	IPAddressType string   `json:"ipAddressType,omitempty"` //IP_V4，IP_V6
	SrcAddress    []string   `json:"srcAddress,omitempty"`    //[“192.168.1.1/32”, “172.18.0.0/24”, “10.0.0.0/16”]
	SrcPort       []string `json:"srcPort,omitempty"`       //[“80”, “2000-3000”]
	DstAddress    []string `json:"dstAddress,omitempty"`    //[“192.168.1.1/32”, “172.18.0.0/24”, “10.0.0.0/16”]
	DstPort       []string `json:"dstPort,omitempty"`       //[“80”, “2000-3000”]
	Protocol      []string `json:"protocol,omitempty"`      // ["60", “76”]
	Qci           int16    `json:"qci,omitempty"`
	Dscp          int16    `json:"dscp,omitempty"`
	Tc            int16    `json:"tc,omitempty"`
}

func main() {
	    s := "{\"trafficRuleId\":\"test_ip\",\"priority\":5,\"action\":\"DROP\",\"trafficFilter\":[{\"ipAddressType\":\"IP_V4\",\"srcAddress\":[\"192.168.0.11\"],\"srcPort\":[\"80\"],\"dstAddress\":[\"192.168.0.120\"],\"dstPort\":[\"8080\"],\"protocol\":[\"5\"],\"dscp\":2}]}"
	    jsonBody := []byte(s)
        fmt.Println(s)
        fmt.Println(jsonBody)
	    trNew := &trafficRule{}
	    err := json.Unmarshal(jsonBody, trNew)
	    if err != nil {
		    fmt.Println(err)
	    }
	    fmt.Println(trNew)

}