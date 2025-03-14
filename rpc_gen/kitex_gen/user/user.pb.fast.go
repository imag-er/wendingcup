// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package user

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *RegisterRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RegisterRequest[number], err)
}

func (x *RegisterRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Teamname, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RegisterRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Player
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Players = append(x.Players, &v)
	return offset, nil
}

func (x *RegisterResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RegisterResponse[number], err)
}

func (x *RegisterResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Code, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *RegisterResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Msg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RegisterResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v TeamInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Teaminfo = &v
	return offset, nil
}

func (x *LoginRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_LoginRequest[number], err)
}

func (x *LoginRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TeamId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_LoginResponse[number], err)
}

func (x *LoginResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Code, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *LoginResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Msg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v TeamInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Teaminfo = &v
	return offset, nil
}

func (x *Player) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Player[number], err)
}

func (x *Player) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Player) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Phonenumber, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Player) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.StudentId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *TeamInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_TeamInfo[number], err)
}

func (x *TeamInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TeamId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *TeamInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Teamname, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *TeamInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v Player
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Players = append(x.Players, &v)
	return offset, nil
}

func (x *GetTeamInfoRequst) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetTeamInfoRequst[number], err)
}

func (x *GetTeamInfoRequst) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TeamId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RegisterRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RegisterRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Teamname == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTeamname())
	return offset
}

func (x *RegisterRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Players == nil {
		return offset
	}
	for i := range x.GetPlayers() {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.GetPlayers()[i])
	}
	return offset
}

func (x *RegisterResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *RegisterResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Code == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetCode())
	return offset
}

func (x *RegisterResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Msg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMsg())
	return offset
}

func (x *RegisterResponse) fastWriteField3(buf []byte) (offset int) {
	if x.Teaminfo == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetTeaminfo())
	return offset
}

func (x *LoginRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *LoginRequest) fastWriteField1(buf []byte) (offset int) {
	if x.TeamId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTeamId())
	return offset
}

func (x *LoginResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *LoginResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Code == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetCode())
	return offset
}

func (x *LoginResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Msg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMsg())
	return offset
}

func (x *LoginResponse) fastWriteField3(buf []byte) (offset int) {
	if x.Teaminfo == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetTeaminfo())
	return offset
}

func (x *Player) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *Player) fastWriteField1(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetName())
	return offset
}

func (x *Player) fastWriteField2(buf []byte) (offset int) {
	if x.Phonenumber == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPhonenumber())
	return offset
}

func (x *Player) fastWriteField3(buf []byte) (offset int) {
	if x.StudentId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetStudentId())
	return offset
}

func (x *TeamInfo) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *TeamInfo) fastWriteField1(buf []byte) (offset int) {
	if x.TeamId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTeamId())
	return offset
}

func (x *TeamInfo) fastWriteField2(buf []byte) (offset int) {
	if x.Teamname == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetTeamname())
	return offset
}

func (x *TeamInfo) fastWriteField3(buf []byte) (offset int) {
	if x.Players == nil {
		return offset
	}
	for i := range x.GetPlayers() {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.GetPlayers()[i])
	}
	return offset
}

func (x *GetTeamInfoRequst) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetTeamInfoRequst) fastWriteField1(buf []byte) (offset int) {
	if x.TeamId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTeamId())
	return offset
}

func (x *RegisterRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RegisterRequest) sizeField1() (n int) {
	if x.Teamname == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTeamname())
	return n
}

func (x *RegisterRequest) sizeField2() (n int) {
	if x.Players == nil {
		return n
	}
	for i := range x.GetPlayers() {
		n += fastpb.SizeMessage(2, x.GetPlayers()[i])
	}
	return n
}

func (x *RegisterResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *RegisterResponse) sizeField1() (n int) {
	if x.Code == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetCode())
	return n
}

func (x *RegisterResponse) sizeField2() (n int) {
	if x.Msg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMsg())
	return n
}

func (x *RegisterResponse) sizeField3() (n int) {
	if x.Teaminfo == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetTeaminfo())
	return n
}

func (x *LoginRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *LoginRequest) sizeField1() (n int) {
	if x.TeamId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTeamId())
	return n
}

func (x *LoginResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *LoginResponse) sizeField1() (n int) {
	if x.Code == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetCode())
	return n
}

func (x *LoginResponse) sizeField2() (n int) {
	if x.Msg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMsg())
	return n
}

func (x *LoginResponse) sizeField3() (n int) {
	if x.Teaminfo == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetTeaminfo())
	return n
}

func (x *Player) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *Player) sizeField1() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetName())
	return n
}

func (x *Player) sizeField2() (n int) {
	if x.Phonenumber == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPhonenumber())
	return n
}

func (x *Player) sizeField3() (n int) {
	if x.StudentId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetStudentId())
	return n
}

func (x *TeamInfo) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *TeamInfo) sizeField1() (n int) {
	if x.TeamId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTeamId())
	return n
}

func (x *TeamInfo) sizeField2() (n int) {
	if x.Teamname == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetTeamname())
	return n
}

func (x *TeamInfo) sizeField3() (n int) {
	if x.Players == nil {
		return n
	}
	for i := range x.GetPlayers() {
		n += fastpb.SizeMessage(3, x.GetPlayers()[i])
	}
	return n
}

func (x *GetTeamInfoRequst) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetTeamInfoRequst) sizeField1() (n int) {
	if x.TeamId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTeamId())
	return n
}

var fieldIDToName_RegisterRequest = map[int32]string{
	1: "Teamname",
	2: "Players",
}

var fieldIDToName_RegisterResponse = map[int32]string{
	1: "Code",
	2: "Msg",
	3: "Teaminfo",
}

var fieldIDToName_LoginRequest = map[int32]string{
	1: "TeamId",
}

var fieldIDToName_LoginResponse = map[int32]string{
	1: "Code",
	2: "Msg",
	3: "Teaminfo",
}

var fieldIDToName_Player = map[int32]string{
	1: "Name",
	2: "Phonenumber",
	3: "StudentId",
}

var fieldIDToName_TeamInfo = map[int32]string{
	1: "TeamId",
	2: "Teamname",
	3: "Players",
}

var fieldIDToName_GetTeamInfoRequst = map[int32]string{
	1: "TeamId",
}
