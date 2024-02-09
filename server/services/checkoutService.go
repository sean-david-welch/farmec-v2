package services

import (
	"log"

	"github.com/sean-david-welch/farmec-v2/server/config"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
)

type CheckoutService interface {
	CreateCheckoutSession(id string) (*stripe.CheckoutSession, error)
	RetrieveCheckoutSession(sessionId string) (*stripe.CheckoutSession, error)
}

type CheckoutServiceImpl struct {
	secrets    *config.Secrets
	repository repository.LineItemRepository
}

func NewCheckoutService(secrets *config.Secrets, repository repository.LineItemRepository) *CheckoutServiceImpl {
	return &CheckoutServiceImpl{secrets: secrets, repository: repository}
}
func (service *CheckoutServiceImpl) CreateCheckoutSession(id string) (*stripe.CheckoutSession, error) {
	stripe.Key = service.secrets.StripeSecretKeyTest

	log.Printf("Creating checkout session for product ID: %s", id)

	product, err := service.repository.GetLineItemById(id)
	if err != nil {
		log.Printf("Error retrieving product by ID: %v", err)
		return nil, err
	}

	log.Printf("Product retrieved: %+v", product)

	params := &stripe.CheckoutSessionParams{
		UIMode:             stripe.String("embedded"),
		ReturnURL:          stripe.String("http://localhost:5173/return?session_id={CHECKOUT_SESSION_ID}"),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("eur"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name:   stripe.String(product.Name),
						Images: stripe.StringSlice([]string{product.Image}),
					},
					UnitAmount: stripe.Int64(int64(product.Price * 100)),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
	}

	sess, err := session.New(params)
	if err != nil {
		log.Printf("Error creating checkout session: %v", err)
		return nil, err
	}

	log.Printf("Checkout session created successfully: %s", sess.ID)

	return sess, nil
}

func (service *CheckoutServiceImpl) RetrieveCheckoutSession(sessionId string) (*stripe.CheckoutSession, error) {
	stripe.Key = service.secrets.StripeSecretKeyTest

	sess, err := session.Get(sessionId, nil)
	if err != nil {
		return nil, err
	}

	if sess != nil {
		log.Printf("Session ID: %s, Status: %s, Customer Email: %s, Amount Total: %d, v Currency: %s",
			sess.ID,
			sess.Status,
			sess.CustomerDetails.Email,
			sess.AmountTotal,

			sess.Currency)
	}

	return sess, err
}
