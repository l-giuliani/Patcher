package service

import (
	dto "patcher/gui/dto"
	dns "patcher/dns"
)

func GuiGeneratePatch(requestType string, generatePayload dto.GeneratePayload) bool{
	params := make([]string, 4)

	params[0] = generatePayload.PrecVersion
	params[1] = generatePayload.NewVersion
	params[2] = generatePayload.Patch
	params[3] = generatePayload.Crc32

	service, exist := dns.ServiceMap[requestType]
    if (exist) {
        res := service(params)
        if (res) {
            return true
        } else {
            return false
        }
    }

	return false
}