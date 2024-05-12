package deliverypartnerconnectionlib

type DeliveryPartnerConnectionLib struct {
	partnerCreateOrderAdaptor        map[string]OrderCreator
	partnerUpdateOrderAdaptor        map[string]OrderUpdator
	partnerDeleteOrderAdaptor        map[string]OrderDeleter
	partnerHookOrderAdaptor          map[string]OrderHook
	partnerCreateReceivedAdaptor     map[string]OrderReceived
	partnerCancelCreatedOrderAdaptor map[string]OrderCancelCreated
}

func New(
	partnerCreateOrderAdaptor map[string]OrderCreator,
	partnerUpdateOrderAdaptor map[string]OrderUpdator,
	partnerDeleteOrderAdaptor map[string]OrderDeleter,
	partnerHookOrderAdaptor map[string]OrderHook,
	partnerCreateReceivedAdaptor map[string]OrderReceived,
	partnerCancelCreatedOrderAdaptor map[string]OrderCancelCreated,
) *DeliveryPartnerConnectionLib {
	return &DeliveryPartnerConnectionLib{
		partnerCreateOrderAdaptor:        partnerCreateOrderAdaptor,
		partnerUpdateOrderAdaptor:        partnerUpdateOrderAdaptor,
		partnerDeleteOrderAdaptor:        partnerDeleteOrderAdaptor,
		partnerHookOrderAdaptor:          partnerHookOrderAdaptor,
		partnerCreateReceivedAdaptor:     partnerCreateReceivedAdaptor,
		partnerCancelCreatedOrderAdaptor: partnerCancelCreatedOrderAdaptor,
	}
}

func (c *DeliveryPartnerConnectionLib) CreateOrder(partner string, order Order) (map[string]interface{}, error) {
	orderRefID, err := c.partnerCreateOrderAdaptor[partner].CreateOrder(order)
	return orderRefID, err
}

func (c *DeliveryPartnerConnectionLib) UpdateOrder(partner, trackingNo string, order Order) error {
	err := c.partnerUpdateOrderAdaptor[partner].UpdateOrder(trackingNo, order)
	return err
}

func (c *DeliveryPartnerConnectionLib) DeleteOrder(partner, trackingNo string) error {
	err := c.partnerDeleteOrderAdaptor[partner].DeleteOrder(trackingNo)
	return err
}

func (c *DeliveryPartnerConnectionLib) CreateReceived(partner string, order Order) (map[string]interface{}, error) {
	res, err := c.partnerCreateReceivedAdaptor[partner].CreateReceived(order)
	return res, err
}

func (c *DeliveryPartnerConnectionLib) HookOrder(partner string, tracking_no_list []string) (map[string]interface{}, error) {
	res, err := c.partnerHookOrderAdaptor[partner].HookOrder(tracking_no_list)
	return res, err
}

func (c *DeliveryPartnerConnectionLib) CancelCreatedOrder(partner string, trackingNumber string) (map[string]interface{}, error) {
	res, err := c.partnerCancelCreatedOrderAdaptor[partner].CancelCreatedOrder(trackingNumber)
	return res, err
}
