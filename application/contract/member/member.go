/*
 *    Copyright 2019 Insolar Technologies
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package member

import (
	"fmt"
	"github.com/insolar/insolar/application/proxy/account"

	"github.com/insolar/insolar/application/contract/member/signer"
	"github.com/insolar/insolar/application/proxy/nodedomain"
	"github.com/insolar/insolar/application/proxy/rootdomain"
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type Member struct {
	foundation.BaseContract
	Name      string
	PublicKey string
}

func (m *Member) GetName() (string, error) {
	return m.Name, nil
}

var INSATTR_GetPublicKey_API = true

func (m *Member) GetPublicKey() (string, error) {
	return m.PublicKey, nil
}

func New(name string, key string) (*Member, error) {
	return &Member{
		Name:      name,
		PublicKey: key,
	}, nil
}

func (m *Member) VerifySig(method string, params []byte, seed []byte, sign []byte) error {
	args, err := core.MarshalArgs(m.GetReference(), method, params, seed)
	if err != nil {
		return fmt.Errorf("[ verifySig ] Can't MarshalArgs: %s", err.Error())
	}
	key, err := m.GetPublicKey()
	if err != nil {
		return fmt.Errorf("[ verifySig ]: %s", err.Error())
	}

	publicKey, err := foundation.ImportPublicKey(key)
	if err != nil {
		return fmt.Errorf("[ verifySig ] Invalid public key")
	}

	verified := foundation.Verify(args, sign, publicKey)
	if !verified {
		return fmt.Errorf("[ verifySig ] Incorrect signature")
	}
	return nil
}

var INSATTR_Call_API = true

// Call method for authorized calls
func (m *Member) Call(rootDomain core.RecordRef, method string, params []byte, seed []byte, sign []byte) (interface{}, error) {

	switch method {
	case "CreateMember":
		return m.createMemberCall(rootDomain, params)
	}

	if err := m.VerifySig(method, params, seed, sign); err != nil {
		return nil, fmt.Errorf("[ Call ]: %s", err.Error())
	}

	switch method {
	case "DumpUserInfo":
		return m.dumpUserInfoCall(rootDomain, params)
	case "DumpAllUsers":
		return m.dumpAllUsersCall(rootDomain)
	case "RegisterNode":
		return m.registerNodeCall(rootDomain, params)
	case "GetNodeRef":
		return m.getNodeRefCall(rootDomain, params)

	case "GetAccountRef":
		return m.getAccountRef()

		// ethStore methods
	case "SaveToMap":
		return m.saveEthTx(rootDomain, method, params, seed, sign)

		// account methods
	case "GetBalance", "Transfer", "SecretTransfer", "ApplySecret":
		return m.CallAccount(rootDomain, method, params, seed, sign)

	}
	return nil, &foundation.Error{S: "[ Member Call ] Unknown method"}
}

func (m *Member) createMemberCall(ref core.RecordRef, params []byte) (interface{}, error) {
	rootDomain := rootdomain.GetObject(ref)
	var name string
	var key string
	if err := signer.UnmarshalParams(params, &name, &key); err != nil {
		return nil, fmt.Errorf("[ createMemberCall ]: %s", err.Error())
	}
	return rootDomain.CreateMember(name, key)
}

//
//func (m *Member) getMyBalanceCall() (interface{}, error) {
//	w, err := wallet.GetImplementationFrom(m.GetReference())
//	if err != nil {
//		return 0, fmt.Errorf("[ getMyBalanceCall ]: %s", err.Error())
//	}
//
//	return w.GetBalance()
//}
//
//func (m *Member) getBalanceCall(params []byte) (interface{}, error) {
//	var member string
//	if err := signer.UnmarshalParams(params, &member); err != nil {
//		return nil, fmt.Errorf("[ getBalanceCall ] : %s", err.Error())
//	}
//	memberRef, err := core.NewRefFromBase58(member)
//	if err != nil {
//		return nil, fmt.Errorf("[ getBalanceCall ] : %s", err.Error())
//	}
//	w, err := wallet.GetImplementationFrom(*memberRef)
//	if err != nil {
//		return nil, fmt.Errorf("[ getBalanceCall ] : %s", err.Error())
//	}
//
//	return w.GetBalance()
//}

//func (m *Member) transferCall(params []byte) (interface{}, error) {
//	var amount uint
//	var toStr string
//	if err := signer.UnmarshalParams(params, &amount, &toStr); err != nil {
//		return nil, fmt.Errorf("[ transferCall ] Can't unmarshal params: %s", err.Error())
//	}
//	to, err := core.NewRefFromBase58(toStr)
//	if err != nil {
//		return nil, fmt.Errorf("[ transferCall ] Failed to parse 'to' param: %s", err.Error())
//	}
//	if m.GetReference() == *to {
//		return nil, fmt.Errorf("[ transferCall ] Recipient must be different from the sender")
//	}
//	w, err := wallet.GetImplementationFrom(m.GetReference())
//	if err != nil {
//		return nil, fmt.Errorf("[ transferCall ] Can't get implementation: %s", err.Error())
//	}
//
//	return nil, w.Transfer(amount, to)
//}

func (m *Member) dumpUserInfoCall(ref core.RecordRef, params []byte) (interface{}, error) {
	rootDomain := rootdomain.GetObject(ref)
	var user string
	if err := signer.UnmarshalParams(params, &user); err != nil {
		return nil, fmt.Errorf("[ dumpUserInfoCall ] Can't unmarshal params: %s", err.Error())
	}
	return rootDomain.DumpUserInfo(user)
}

func (m *Member) dumpAllUsersCall(ref core.RecordRef) (interface{}, error) {
	rootDomain := rootdomain.GetObject(ref)
	return rootDomain.DumpAllUsers()
}

func (m *Member) registerNodeCall(ref core.RecordRef, params []byte) (interface{}, error) {
	var publicKey string
	var role string
	if err := signer.UnmarshalParams(params, &publicKey, &role); err != nil {
		return nil, fmt.Errorf("[ registerNodeCall ] Can't unmarshal params: %s", err.Error())
	}

	rootDomain := rootdomain.GetObject(ref)
	nodeDomainRef, err := rootDomain.GetNodeDomainRef()
	if err != nil {
		return nil, fmt.Errorf("[ registerNodeCall ] %s", err.Error())
	}

	nd := nodedomain.GetObject(nodeDomainRef)
	cert, err := nd.RegisterNode(publicKey, role)
	if err != nil {
		return nil, fmt.Errorf("[ registerNodeCall ] Problems with RegisterNode: %s", err.Error())
	}

	return string(cert), nil
}

func (m *Member) getNodeRefCall(ref core.RecordRef, params []byte) (interface{}, error) {
	var publicKey string
	if err := signer.UnmarshalParams(params, &publicKey); err != nil {
		return nil, fmt.Errorf("[ getNodeRefCall ] Can't unmarshal params: %s", err.Error())
	}

	rootDomain := rootdomain.GetObject(ref)
	nodeDomainRef, err := rootDomain.GetNodeDomainRef()
	if err != nil {
		return nil, fmt.Errorf("[ getNodeRefCall ] Can't get nodeDmainRef: %s", err.Error())
	}

	nd := nodedomain.GetObject(nodeDomainRef)
	nodeRef, err := nd.GetNodeRefByPK(publicKey)
	if err != nil {
		return nil, fmt.Errorf("[ getNodeRefCall ] Node not found: %s", err.Error())
	}

	return nodeRef, nil
}

func (m *Member) getAccountRef() (string, error) {

	a, err := account.GetImplementationFrom(m.GetReference())
	if err != nil {
		return "", fmt.Errorf("[ getAccountRef ] Can't get implementation: %s", err.Error())
	}

	return a.GetReference().String(), nil
}

func (m *Member) saveEthTx(rootDomainRef core.RecordRef, method string, params []byte, seed []byte, sign []byte) (interface{}, error) {

	type inputRequest struct {
		OracleName string `oracleName`
		EthAddr    string `ethAddr`
		Balance    uint   `balance`
		EthTxHash  string `ethTxHash`
	}

	inputJSON := new(inputRequest)

	//verifySign that it is oracle

	if err := signer.UnmarshalParams(params, &inputJSON); err != nil {
		return nil, fmt.Errorf("[ saveEthTx ]: %s", err.Error())
	}

	rootDomain := rootdomain.GetObject(rootDomainRef)
	return rootDomain.SaveEthTx(inputJSON.EthAddr, inputJSON.Balance, inputJSON.EthTxHash, inputJSON.OracleName)
}

func (m *Member) CallAccount(rootDomain core.RecordRef, method string, params []byte, seed []byte, sign []byte) (interface{}, error) {

	account, err := account.GetImplementationFrom(m.GetReference())
	if err != nil {
		return "", fmt.Errorf("[ CallAccount ] Can't get implementation: %s", err.Error())
	}

	return account.Call(rootDomain, method, params, seed, sign)
}
