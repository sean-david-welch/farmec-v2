package resources

import (
	"context"
	"database/sql"
	"embed"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
)

type Resources struct {
	// Services
	AuthService         services.AuthService
	BlogService         services.BlogService
	CarouselService     services.CarouselService
	CheckoutService     services.CheckoutService
	ContactService      services.ContactService
	EmployeeService     services.EmployeeService
	LineItemService     services.LineItemService
	MachineService      services.MachineService
	PartsService        services.PartsService
	PdfService          services.PdfService
	ProductService      services.ProductService
	RegistrationService services.RegistrationService
	SupplierService     services.SupplierService
	TimelineService     services.TimelineService
	VideoService        services.VideoService
	WarrantyService     services.WarrantyService

	// Handlers
	AuthHandler         *handlers.AuthHandler
	AboutHandler        *handlers.AboutHandler
	ViewHandler         *handlers.ViewHandler
	BlogHandler         *handlers.BlogHandler
	CarouselHandler     *handlers.CarouselHandler
	CheckoutHandler     *handlers.CheckoutHandler
	LineItemHandler     *handlers.LineItemHandler
	MachineHandler      *handlers.MachineHandler
	PartsHandler        *handlers.PartsHandler
	PdfHandler          *handlers.PdfHandler
	ProductHandler      *handlers.ProductHandler
	RegistrationHandler *handlers.RegistrationHandler
	SupplierHandler     *handlers.SupplierHandler
	VideoHandler        *handlers.VideoHandler
	WarrantyHandler     *handlers.WarrantyHandler
}

func NewResources(
	db *sql.DB, secrets *lib.Secrets, firebase *lib.Firebase, files embed.FS,
	s3Client lib.S3Client, emailClient *lib.EmailClientImpl, authMiddleware *middleware.AuthMiddlewareImpl, supplierCache *middleware.SupplierCache,
) *Resources {
	// Initialize repositories
	blogRepo := repository.NewBlogRepo(db)
	carouselRepo := repository.NewCarouselRepo(db)
	employeeRepo := repository.NewEmployeeRepo(db)
	lineItemRepo := repository.NewLineItemRepo(db)
	machineRepo := repository.NewMachineRepo(db)
	partsRepo := repository.NewPartsRepo(db)
	productRepo := repository.NewProductRepo(db)
	registrationRepo := repository.NewRegistrationRepo(db)
	supplierRepo := repository.NewSupplierRepo(db)
	timelineRepo := repository.NewTimelineRepo(db)
	videoRepo := repository.NewVideoRepo(db)
	warrantyRepo := repository.NewWarrantyRepo(db)

	// Initialize services
	authService := services.NewAuthService(firebase)
	blogService := services.NewBlogService(blogRepo, s3Client, "Blogs")
	carouselService := services.NewCarouselService(carouselRepo, s3Client, "Carousels")
	checkoutService := services.NewCheckoutService(secrets, lineItemRepo)
	contactService := services.NewContactService(emailClient)
	employeeService := services.NewEmployeeService(employeeRepo, s3Client, "Employees")
	lineItemService := services.NewLineItemService(lineItemRepo, s3Client, "LineItems")
	machineService := services.NewMachineService(machineRepo, s3Client, "Machines")
	partsService := services.NewPartsService(partsRepo, s3Client, "Parts")
	pdfService := services.NewPdfService(files)
	productService := services.NewProductService(productRepo, s3Client, "Products")
	registrationService := services.NewRegistrationService(registrationRepo, emailClient)
	supplierService := services.NewSupplierService(supplierRepo, s3Client, "Suppliers")
	timelineService := services.NewTimelineService(timelineRepo)
	yt, err := youtube.NewService(context.Background(), option.WithAPIKey(secrets.YoutubeApiKey))
	if err != nil {
		log.Fatal("error calling YouTube API: ", err)
	}
	videoService := services.NewVideoService(videoRepo, yt)
	warrantyService := services.NewWarrantyService(warrantyRepo, emailClient)

	return &Resources{
		// Services
		AuthService:         authService,
		BlogService:         blogService,
		CarouselService:     carouselService,
		CheckoutService:     checkoutService,
		ContactService:      contactService,
		EmployeeService:     employeeService,
		LineItemService:     lineItemService,
		MachineService:      machineService,
		PartsService:        partsService,
		PdfService:          pdfService,
		ProductService:      productService,
		RegistrationService: registrationService,
		SupplierService:     supplierService,
		TimelineService:     timelineService,
		VideoService:        videoService,
		WarrantyService:     warrantyService,

		// Handlers
		AuthHandler:         handlers.NewAuthHandler(authService),
		AboutHandler:        handlers.NewAboutHandler(employeeService, timelineService, authMiddleware, supplierCache),
		ViewHandler:         handlers.NewViewHandler(carouselService, contactService, authMiddleware, supplierCache),
		BlogHandler:         handlers.NewBlogHandler(blogService, authMiddleware, supplierCache),
		CarouselHandler:     handlers.NewCarouselHandler(carouselService, authMiddleware, supplierCache),
		CheckoutHandler:     handlers.NewCheckoutHandler(checkoutService),
		LineItemHandler:     handlers.NewLineItemHandler(lineItemService),
		MachineHandler:      handlers.NewMachineHandler(machineService, authMiddleware, supplierCache),
		PartsHandler:        handlers.NewPartsHandler(partsService),
		PdfHandler:          handlers.NewPdfHandler(pdfService),
		ProductHandler:      handlers.NewProductHandler(productService),
		RegistrationHandler: handlers.NewRegistrationHandler(registrationService),
		SupplierHandler:     handlers.NewSupplierHandler(supplierService, authMiddleware, supplierCache),
		VideoHandler:        handlers.NewVideoHandler(videoService),
		WarrantyHandler:     handlers.NewWarrantyHandler(warrantyService),
	}
}
