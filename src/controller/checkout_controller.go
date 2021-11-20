package controller

import (
	"context"
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/OrlandoRomo/go-ambassador/src/database"
	"github.com/OrlandoRomo/go-ambassador/src/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

type MailHog struct {
	Host       string
	From       string
	AdminEmail string
}

var MailHogClient MailHog

func GetLinksByCode(c *fiber.Ctx) error {
	code := c.Params("code")
	if code == "" {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("the value %s is not valid", code),
		})
	}

	link := model.Link{
		Code: code,
	}

	tcx := database.DB.Preload("User").Preload("Products").First(&link)
	if tcx.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("there is not a link with code %s\n", code),
		})
	}
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}

	return c.JSON(link)
}

func CreateCheckoutOrders(c *fiber.Ctx) error {
	var request model.OrderRequest
	if err := c.BodyParser(&request); err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	link := model.Link{
		Code: request.Code,
	}
	tcx := database.DB.Preload("User").First(&link)
	if tcx.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("there is not a link with code %s\n", request.Code),
		})
	}
	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}
	order := model.Order{
		Code:            link.Code,
		UserID:          link.UserID,
		AmbassadorEmail: link.User.Email,
		FirstName:       request.FirstName,
		LastName:        request.LastName,
		Email:           request.Email,
		Address:         request.Address,
		City:            request.Country,
		Zip:             request.Zip,
	}

	tcx = database.DB.Begin()
	err := tcx.Create(&order).Error

	if err != nil {
		tcx.Rollback()
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}

	lineItems := make([]*stripe.CheckoutSessionLineItemParams, 0)

	for _, p := range request.Products {
		product := model.Product{
			ID: uint(p["product_id"]),
		}

		tcxProduct := database.DB.First(&product)
		if tcxProduct.RowsAffected == 0 {
			c.Status(http.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": fmt.Sprintf("there is not a product with %d\n", p["product_id"]),
			})
		}
		if tcxProduct.Error != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": tcxProduct.Error.Error(),
			})
		}

		total := product.Price * float64(p["quantity"])
		item := model.OrderItem{
			OrderID:           order.ID,
			ProductTitle:      product.Title,
			Price:             product.Price,
			Quantity:          uint(p["quantity"]),
			AmbassadorRevenue: 0.1 * total,
			AdminRevenue:      0.9 * total,
		}
		err = tcx.Create(&item).Error
		if err != nil {
			tcx.Rollback()
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": tcx.Error.Error(),
			})
		}

		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			Name:        &product.Title,
			Description: &product.Description,
			Images:      stripe.StringSlice([]string{product.Image}),
			Amount:      stripe.Int64(100 * int64(product.Price)),
			Currency:    stripe.String("usd"),
			Quantity:    stripe.Int64(int64(p["quantity"])),
		})
	}

	params := stripe.CheckoutSessionParams{
		// replace with dynamic values (this is the url for the frontend-side)
		SuccessURL:         stripe.String("http://localhost:5000/success?source={CHECKOUT_SESSION_ID}"),
		CancelURL:          stripe.String("http://localhost:5000/failure"),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems:          lineItems,
	}
	source, err := session.New(&params)
	if err != nil {
		tcx.Rollback()
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	order.TransactionID = source.ID

	if err := tcx.Save(&order).Error; err != nil {
		tcx.Rollback()
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}

	if err := tcx.Commit().Error; err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(source)
}

func ConfirmCheckoutOrders(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	order := model.Order{}

	tcx := database.DB.Preload("OrderItems").First(&order, model.Order{
		TransactionID: data["source"],
	})

	if tcx.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("there is not order's transaction id with %s\n", data["source"]),
		})
	}

	if tcx.Error != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}

	if !order.IsCompleted {
		order.IsCompleted = true
	}

	order.IsCompleted = true
	if err := tcx.Save(&order).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": tcx.Error.Error(),
		})
	}

	ambassadorRevenue := 0.0
	adminRevenue := 0.0
	for _, item := range order.OrderItems {
		ambassadorRevenue += item.AmbassadorRevenue
		adminRevenue += item.AdminRevenue
	}
	user := model.User{
		ID: order.UserID,
	}
	database.DB.First(&user)
	database.Cache.ZIncrBy(context.Background(), "rankings", ambassadorRevenue, user.GetFullName())

	ambassadorMessage := []byte(fmt.Sprintf("You earned %.2f from the link #%s", ambassadorRevenue, order.Code))

	smtp.SendMail(MailHogClient.Host, nil, MailHogClient.From, []string{order.AmbassadorEmail}, ambassadorMessage)

	adminMessage := []byte(fmt.Sprintf("You earned %.2f from the link #%s", ambassadorRevenue, order.Code))

	smtp.SendMail(MailHogClient.Host, nil, MailHogClient.From, []string{MailHogClient.AdminEmail}, adminMessage)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
