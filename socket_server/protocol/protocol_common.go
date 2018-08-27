package protocol

import (
	"bytes"
	"strings"
)

const (
	PROTOCOL_MIN_LENGTH int    = 33
	PROTOCOL_START_FLAG string = "$HX"
	PROTOCOL_END_FLAG   string = "\r\n"
	PROTOCOL_SEP        string = ","

	PROTOCOL_ILLEGAL   string = "illegal"
	PROTOCOL_HALF_PACK string = "half_pack"
	PROTOCOL_UNKNOWN   string = "unknown"

	PROTOCOL_REPORT_LOGIN    string = "0001"
	PROTOCOL_REPORT_HEART    string = "0002"
	PROTOCOL_REPORT_TIME     string = "0004"
	PROTOCOL_REPORT_LOCATION string = "0005"
	PROTOCOL_REPORT_LBS      string = "0007"

	PROTOCOL_DISTRIBUTE_REQP      string = "1001"
	PROTOCOL_DISTRIBUTE_FRESET    string = "1002"
	PROTOCOL_DISTRIBUTE_RESP_TIME string = "1004"
	PROTOCOL_DISTRIBUTE_RESP_LBS  string = "1007"
	PROTOCOL_DISTRIBUTE_MESSAGE   string = "1011"
)

func Parse(buffer string) []string {
	return strings.Split(buffer, PROTOCOL_SEP)
}

func write_header(protocol_id string, imei string) string {
	cmd := PROTOCOL_START_FLAG + PROTOCOL_SEP
	cmd += protocol_id
	cmd += PROTOCOL_SEP
	cmd += imei
	cmd += PROTOCOL_SEP

	return cmd
}

func CheckProtocol(buffer *bytes.Buffer) (string, []string) {
	cmd := PROTOCOL_ILLEGAL
	var values []string
	bufferlen := buffer.Len()
	if bufferlen == 0 {
		return PROTOCOL_ILLEGAL, nil
	}
	p := string(buffer.Bytes())
	if string(p[0:3]) != PROTOCOL_START_FLAG {
		buffer.ReadByte()
		cmd, values = CheckProtocol(buffer)
	} else if bufferlen >= PROTOCOL_MIN_LENGTH {
		end_idx := strings.Index(p, PROTOCOL_END_FLAG)
		if end_idx == -1 {
			return PROTOCOL_HALF_PACK, nil
		} else {
			buf, _ := buffer.ReadString('\n')
			values = strings.Split(buf, PROTOCOL_SEP)
			return values[1], values
		}
	} else {
		return PROTOCOL_HALF_PACK, nil
	}

	return cmd, values
}
