package cart

import (
	"context"
	"encoding/json"

	"flamingo.me/flamingo/v3/core/oauth/domain"
	"flamingo.me/flamingo/v3/framework/flamingo"

	"github.com/pkg/errors"
)

type (
	// GuestCartService interface - Secondary PORT
	GuestCartService interface {
		// GetModifyBehaviour gets the behaviour for the guest cart service
		GetModifyBehaviour(context.Context) (ModifyBehaviour, error)
		GetCart(ctx context.Context, cartID string) (*Cart, error)
		// GetNewCart - should return a new guest cart (including the id of the cart)
		GetNewCart(ctx context.Context) (*Cart, error)
		// RestoreCart restores a previously used guest cart with all its content.
		// Depending on the used adapter this can lead to a new Cart.ID
		// Deprecated: Implement CompleteBehaviour instead
		RestoreCart(ctx context.Context, cart Cart) (*Cart, error)
	}

	// CustomerCartService interface - Secondary PORT
	CustomerCartService interface {
		// GetModifyBehaviour gets the behaviour for the customer cart service
		GetModifyBehaviour(context.Context, domain.Auth) (ModifyBehaviour, error)
		GetCart(ctx context.Context, auth domain.Auth, cartID string) (*Cart, error)
		// RestoreCart restores a previously used customer cart with all its content.
		// Depending on the used adapter this can lead to a new Cart.ID
		// Deprecated: Implement CompleteBehaviour instead
		RestoreCart(ctx context.Context, auth domain.Auth, cart Cart) (*Cart, error)
	}

	// DeferEvents represents events that should be dispatched after a cart modify call
	DeferEvents []flamingo.Event

	// ModifyBehaviour is a interface that can be implemented by other packages to provide cart actions
	// This port can not be registered directly but is provided by the registered "GuestCartService"
	ModifyBehaviour interface {
		DeleteItem(ctx context.Context, cart *Cart, itemID string, deliveryCode string) (*Cart, DeferEvents, error)
		UpdateItem(ctx context.Context, cart *Cart, itemUpdateCommand ItemUpdateCommand) (*Cart, DeferEvents, error)
		UpdateItems(ctx context.Context, cart *Cart, itemUpdateCommands []ItemUpdateCommand) (*Cart, DeferEvents, error)
		AddToCart(ctx context.Context, cart *Cart, deliveryCode string, addRequest AddRequest) (*Cart, DeferEvents, error)
		CleanCart(ctx context.Context, cart *Cart) (*Cart, DeferEvents, error)
		CleanDelivery(ctx context.Context, cart *Cart, deliveryCode string) (*Cart, DeferEvents, error)
		UpdatePurchaser(ctx context.Context, cart *Cart, purchaser *Person, additionalData *AdditionalData) (*Cart, DeferEvents, error)
		UpdateAdditionalData(ctx context.Context, cart *Cart, additionalData *AdditionalData) (*Cart, DeferEvents, error)
		UpdatePaymentSelection(ctx context.Context, cart *Cart, paymentSelection PaymentSelection) (*Cart, DeferEvents, error)
		UpdateDeliveryInfo(ctx context.Context, cart *Cart, deliveryCode string, deliveryInfo DeliveryInfoUpdateCommand) (*Cart, DeferEvents, error)
		UpdateBillingAddress(ctx context.Context, cart *Cart, billingAddress Address) (*Cart, DeferEvents, error)
		UpdateDeliveryInfoAdditionalData(ctx context.Context, cart *Cart, deliveryCode string, additionalData *AdditionalData) (*Cart, DeferEvents, error)
		ApplyVoucher(ctx context.Context, cart *Cart, couponCode string) (*Cart, DeferEvents, error)
		RemoveVoucher(ctx context.Context, cart *Cart, couponCode string) (*Cart, DeferEvents, error)
	}

	// CompleteBehaviour can be implemented by a cart service.
	// Complete is normally called before the cart is placed
	// This can for example be used to invalidate gift cards
	CompleteBehaviour interface {
		Complete(context.Context) (*Cart, DeferEvents, error)
		Restore(context.Context, Cart) (*Cart, DeferEvents, error)
	}

	//GiftCardBehaviour - additional interface that can be implemented to support GiftCard features
	GiftCardBehaviour interface {
		ApplyGiftCard(ctx context.Context, cart *Cart, giftCardCode string) (*Cart, DeferEvents, error)
		RemoveGiftCard(ctx context.Context, cart *Cart, giftCardCode string) (*Cart, DeferEvents, error)
	}

	//GiftCardAndVoucherBehaviour - additional interface that can be implemented to support generic code entry (which can either be voucher or giftcard)
	GiftCardAndVoucherBehaviour interface {
		ApplyAny(ctx context.Context, cart *Cart, anyCode string) (*Cart, DeferEvents, error)
	}

	// AddRequest defines add to cart request
	AddRequest struct {
		MarketplaceCode        string
		Qty                    int
		VariantMarketplaceCode string
	}

	// ItemUpdateCommand defines the update item command
	ItemUpdateCommand struct {
		// Source Id of where the items should be initial picked - This is set by the SourcingLogic
		SourceID       *string
		Qty            *int
		AdditionalData map[string]string
		// Mandatory field: ItemID is only for identifying the item.
		ItemID string
	}

	// DeliveryInfoUpdateCommand defines the update item command
	DeliveryInfoUpdateCommand struct {
		DeliveryInfo DeliveryInfo
		additional   map[string]json.RawMessage
	}
)

var (
	// ErrCartNotFound is used if a cart was not found
	ErrCartNotFound = errors.New("Cart not found")
	// ErrItemNotFound is used if a item on cart was not found
	ErrItemNotFound = errors.New("Item not found")
	// ErrDeliveryCodeNotFound is used if a delivery was not found
	ErrDeliveryCodeNotFound = errors.New("Delivery not found")
)

//CreateDeliveryInfoUpdateCommand - factory to get the update command based on the given deliveryInfos (which might come from cart)
func CreateDeliveryInfoUpdateCommand(info DeliveryInfo) DeliveryInfoUpdateCommand {
	uc := DeliveryInfoUpdateCommand{
		DeliveryInfo: info,
	}

	uc.SetAdditional(info.AdditionalDeliveryInfos)
	return uc
}

// AddAdditional adds additional delivery info data
func (d *DeliveryInfoUpdateCommand) AddAdditional(key string, val AdditionalDeliverInfo) (err error) {
	d.init()
	d.additional[key], err = val.Marshal()
	return err
}

// SetAdditional adds additional delivery info data
func (d *DeliveryInfoUpdateCommand) SetAdditional(val map[string]json.RawMessage) {
	d.init()

	if val == nil {
		return
	}

	d.additional = val
}

// Additional gets the additional data as war map from the delivery info update command
func (d *DeliveryInfoUpdateCommand) Additional() map[string]json.RawMessage {
	d.init()
	return d.additional
}

func (d *DeliveryInfoUpdateCommand) init() {
	if d.additional == nil {
		d.additional = make(map[string]json.RawMessage)
	}
}
