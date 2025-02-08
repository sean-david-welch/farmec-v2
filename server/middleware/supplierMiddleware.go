package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"net/http"
	"sync"
	"time"
)

type SupplierCache struct {
	suppliers []types.Supplier
	lastFetch time.Time
	mutex     sync.RWMutex
	ttl       time.Duration
}

func NewSupplierCache(ttl time.Duration) *SupplierCache {
	return &SupplierCache{
		ttl: ttl,
	}
}

func (sc *SupplierCache) Get() ([]types.Supplier, bool) {
	sc.mutex.RLock()
	defer sc.mutex.RUnlock()

	if expired := time.Since(sc.lastFetch) > sc.ttl; expired {
		return nil, false
	}
	return sc.suppliers, true
}

func (sc *SupplierCache) Set(suppliers []types.Supplier) {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()

	sc.suppliers = suppliers
	sc.lastFetch = time.Now()
}

func WithSupplierCache(cache *SupplierCache, service services.SupplierService) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Check cache first
		if suppliers, ok := cache.Get(); ok {
			context.Set("suppliers", suppliers)
			context.Next()
			return
		}

		// Fetch from database
		ctx := context.Request.Context()
		suppliers, err := service.GetSuppliers(ctx)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch suppliers"})
			context.Abort()
			return
		}

		// Update cache
		cache.Set(suppliers)
		context.Set("suppliers", suppliers)
		context.Next()
	}
}

func GetSuppliersFromContext(context *gin.Context) []types.Supplier {
	if suppliers, exists := context.Get("suppliers"); exists {
		if supplierList, ok := suppliers.([]types.Supplier); ok {
			return supplierList
		}
	}
	return []types.Supplier{}
}
