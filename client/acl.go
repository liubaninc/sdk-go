package client

import ()

type SetRoleBuilder struct {
	SetRoleRequest
}

func (s *SetRoleBuilder) SetCreator(creator string) {
	s.Base.Creator = creator
}

func (s *SetRoleBuilder) SetFees(fees string) {
	s.Base.Fees = fees
}

func (s *SetRoleBuilder) SetMemo(memo string) {
	s.Base.Memo = memo
}

func (s *SetRoleBuilder) SetLock(lock uint64) {
	s.Base.Lock = lock
}

func (s *SetRoleBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	s.Base.TimeoutHeight = timeoutHeight
}

func (s *SetRoleBuilder) SetAddress(address string) {
	s.Address = address
}

func (s *SetRoleBuilder) SetRoles(roles []string) {
	s.Roles = roles
}

func (s *SetRoleBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxSetRole"
}

func (s *SetRoleBuilder) Bytes() []byte {
	bts, err := s.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}
