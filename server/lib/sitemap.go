package lib

import (
	"context"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/db"
)

func SitemapData(ctx context.Context, queries *db.Queries) map[string]interface{} {
	result := map[string]interface{}{
		"suppliers":  []string{},
		"spareParts": []string{},
		"blogPosts":  []string{},
	}

	// Get all supplier slugs
	suppliers, err := queries.GetAllSupplierSlugs(ctx)
	if err == nil {
		supplierPaths := make([]string, len(suppliers))
		for i, slug := range suppliers {
			supplierPaths[i] = fmt.Sprintf("/suppliers/%s", slug)
		}
		result["suppliers"] = supplierPaths
	}

	// Get all spare parts slugs
	spareParts, err := queries.GetAllSparePartsSlugs(ctx)
	if err == nil {
		sparePartsPaths := make([]string, len(spareParts))
		for i, slug := range spareParts {
			sparePartsPaths[i] = fmt.Sprintf("/spareparts/%s", slug)
		}
		result["spareParts"] = sparePartsPaths
	}

	// Get all blog slugs
	blogs, err := queries.GetAllBlogSlugs(ctx)
	if err == nil {
		blogPaths := make([]string, len(blogs))
		for i, slug := range blogs {
			blogPaths[i] = fmt.Sprintf("/blogs/%s", slug)
		}
		result["blogPosts"] = blogPaths
	}

	return result
}
