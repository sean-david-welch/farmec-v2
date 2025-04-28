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
		var supplierPaths []string
		for _, slugNullable := range suppliers {
			if slugNullable.Valid {
				supplierPaths = append(supplierPaths, fmt.Sprintf("/suppliers/%s", slugNullable.String))
			}
		}
		result["suppliers"] = supplierPaths
	}

	// Get all spare parts slugs
	spareParts, err := queries.GetAllSparePartsSlugs(ctx)
	if err == nil {
		var sparePartsPaths []string
		for _, slugNullable := range spareParts {
			if slugNullable.Valid {
				sparePartsPaths = append(sparePartsPaths, fmt.Sprintf("/spareparts/%s", slugNullable.String))
			}
		}
		result["spareParts"] = sparePartsPaths
	}

	// Get all blog slugs
	blogs, err := queries.GetAllBlogSlugs(ctx)
	if err == nil {
		var blogPaths []string
		for _, slugNullable := range blogs {
			if slugNullable.Valid {
				blogPaths = append(blogPaths, fmt.Sprintf("/blogs/%s", slugNullable.String))
			}
		}
		result["blogPosts"] = blogPaths
	}

	return result
}
