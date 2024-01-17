package services

import (
	"log"

	"github.com/sean-david-welch/farmec-v2/server/config"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
)

type CheckoutService struct {
	secrets *config.Secrets
	repository *repository.LineItemRepository
}

func NewCheckoutService(secrets *config.Secrets, repository *repository.LineItemRepository) *CheckoutService {
	return &CheckoutService{secrets: secrets, repository: repository}
}

func(service *CheckoutService) CreateCheckoutSession(id string) (*stripe.CheckoutSession, error) {
	stripe.Key = service.secrets.StripeSecretKey

	product, err := service.repository.GetLineItemById(id); if err != nil {
		return nil, err
	}

	params := &stripe.CheckoutSessionParams{
		UIMode: stripe.String("embedded"),
		ReturnURL: stripe.String(service.secrets.Domain + "/return?session_id={CHECKOUT_SESSION_ID}"),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("eur"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(product.Name),
					},
					UnitAmount: stripe.Int64(int64(product.Price * 100)),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		// SuccessURL: stripe.String(service.secrets.Domain + "payments/success"),
		// CancelURL: stripe.String(service.secrets.Domain + "payments/failure"),
		AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{Enabled: stripe.Bool(true)},
	}

	sess, err := session.New(params); if err != nil {
		log.Printf("session.Ne1: %v", err)
	}

	return sess, nil
}

func(service *CheckoutService) RetrieveCheckoutSession(sessionId string) (*stripe.CheckoutSession, error) {
	stripe.Key = service.secrets.StripeSecretKey

	sess, err := session.Get(sessionId, nil); if err != nil {
		return nil, err
	}

	return sess, err
}