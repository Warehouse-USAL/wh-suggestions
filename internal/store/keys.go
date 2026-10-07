// Package store defines the Redis data model (RFC §5.3).
package store

// HotGlobal is the sorted set of globally trending SKUs.
const HotGlobal = "hot:global"

// UserSKUs is the sorted set of SKUs a user has ordered.
func UserSKUs(userID string) string { return "user:" + userID + ":skus" }

// Cooccur is the sorted set of SKUs ordered together with sku.
func Cooccur(sku string) string { return "cooccur:" + sku }
