// Code generated by "stringer -type=MessageType"; DO NOT EDIT.

package core

import "strconv"

const _MessageType_name = "TypeCallMethodTypeCallConstructorTypeExecutorResultsTypeValidateCaseBindTypeValidationResultsTypeRequestCallTypeGetCodeTypeGetObjectTypeGetDelegateTypeGetChildrenTypeUpdateObjectTypeRegisterChildTypeJetDropTypeSetRecordTypeValidateRecordTypeSetBlobTypeGetHistoryTypeBootstrapRequest"

var _MessageType_index = [...]uint16{0, 14, 33, 52, 72, 93, 108, 119, 132, 147, 162, 178, 195, 206, 219, 237, 248, 262, 282}

func (i MessageType) String() string {
	if i >= MessageType(len(_MessageType_index)-1) {
		return "MessageType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MessageType_name[_MessageType_index[i]:_MessageType_index[i+1]]
}
