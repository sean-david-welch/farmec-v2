package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
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

	if time.Since(sc.lastFetch) > sc.ttl {
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
	return func(c *gin.Context) {
		// Try to get suppliers from cache
		if suppliers, ok := cache.Get(); ok {
			c.Set("suppliers", suppliers)
			c.Next()
			return
		}

		// If not in cache, fetch from database
		ctx := c.Request.Context()
		suppliers, err := service.GetSuppliers(ctx)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch suppliers"})
			c.Abort()
			return
		}

		// Update cache
		cache.Set(suppliers)
		c.Set("suppliers", suppliers)
		c.Next()
	}
}

func GetSuppliersFromContext(c *gin.Context) []types.Supplier {
	if suppliers, exists := c.Get("suppliers"); exists {
		if s, ok := suppliers.([]types.Supplier); ok {
			return s
		}
	}
	return []types.Supplier{}
}
