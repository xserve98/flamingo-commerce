// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import cart "flamingo.me/flamingo-commerce/v3/cart/domain/cart"
import context "context"
import domain "flamingo.me/flamingo/v3/core/oauth/domain"
import mock "github.com/stretchr/testify/mock"

// CustomerCartService is an autogenerated mock type for the CustomerCartService type
type CustomerCartService struct {
	mock.Mock
}

// GetCart provides a mock function with given fields: ctx, auth, cartID
func (_m *CustomerCartService) GetCart(ctx context.Context, auth domain.Auth, cartID string) (*cart.Cart, error) {
	ret := _m.Called(ctx, auth, cartID)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, domain.Auth, string) *cart.Cart); ok {
		r0 = rf(ctx, auth, cartID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Auth, string) error); ok {
		r1 = rf(ctx, auth, cartID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetModifyBehaviour provides a mock function with given fields: _a0, _a1
func (_m *CustomerCartService) GetModifyBehaviour(_a0 context.Context, _a1 domain.Auth) (cart.ModifyBehaviour, error) {
	ret := _m.Called(_a0, _a1)

	var r0 cart.ModifyBehaviour
	if rf, ok := ret.Get(0).(func(context.Context, domain.Auth) cart.ModifyBehaviour); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cart.ModifyBehaviour)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Auth) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestoreCart provides a mock function with given fields: ctx, auth, _a2
func (_m *CustomerCartService) RestoreCart(ctx context.Context, auth domain.Auth, _a2 cart.Cart) (*cart.Cart, error) {
	ret := _m.Called(ctx, auth, _a2)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, domain.Auth, cart.Cart) *cart.Cart); ok {
		r0 = rf(ctx, auth, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Auth, cart.Cart) error); ok {
		r1 = rf(ctx, auth, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
