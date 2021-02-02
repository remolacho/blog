// pkg/post/post.go
package post

import (
	"gorm.io/gorm"
)

// Post created by a user.
type Post struct {
	gorm.Model
	Body   string `json:"body,omitempty"`
	UserID uint   `json:"user_id,omitempty"`
}
