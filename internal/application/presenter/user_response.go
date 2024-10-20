package presenter

import (
	domain "samsamoohooh-go-api/internal/application/domain"
	"time"
)

type UserCreateResponse struct {
	ID         int                 `json:"id"`
	Name       string              `json:"name"`
	Resolution string              `json:"resolution"`
	Role       domain.UserRoleType `json:"role"`
	CreatedAt  time.Time           `json:"createdAt"`
	UpdatedAt  time.Time           `json:"updatedAt"`
}

func NewUserCreateResponse(user *domain.User) *UserCreateResponse {
	return &UserCreateResponse{
		ID:         user.ID,
		Name:       user.Name,
		Resolution: user.Resolution,
		Role:       user.Role,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}

type UserListResponse struct {
	ID         int                   `json:"id"`
	Name       string                `json:"name"`
	Resolution string                `json:"resolution"`
	Role       domain.UserRoleType   `json:"role"`
	Social     domain.UserSocialType `json:"social"`
	SocialSub  string                `json:"socialSub"`
	CreatedAt  time.Time             `json:"createdAt"`
	UpdatedAt  time.Time             `json:"updatedAt"`
}

func NewListUserResponse(users []*domain.User) []*UserListResponse {
	var listUser []*UserListResponse
	for _, user := range users {
		listUser = append(listUser, &UserListResponse{
			ID:         user.ID,
			Name:       user.Name,
			Resolution: user.Resolution,
			Role:       user.Role,
			Social:     user.Social,
			SocialSub:  user.SocialSub,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		})
	}

	return listUser
}

type UserGetByMeResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Resolution string `json:"resolution"`
}

func NewUserGetByMeResponse(user *domain.User) *UserGetByMeResponse {
	return &UserGetByMeResponse{
		ID:         user.ID,
		Name:       user.Name,
		Resolution: user.Resolution,
	}
}

type UserGetGroupsByMeResponse struct {
	ID          int    `json:"id"`
	BookTitle   string `json:"bookTitle"`
	Author      string `json:"author"`
	MaxPage     int    `json:"maxPage"`
	Publisher   string `json:"publisher"`
	Description string `json:"description"`
	Bookmark    int    `json:"bookmark"`
}

func NewUserGetGroupsByMeResponse(groups []*domain.Group) []*UserGetGroupsByMeResponse {
	var listGroups []*UserGetGroupsByMeResponse
	for _, group := range groups {
		listGroups = append(listGroups, &UserGetGroupsByMeResponse{
			ID:          group.ID,
			BookTitle:   group.BookTitle,
			Author:      group.Author,
			MaxPage:     group.MaxPage,
			Publisher:   group.Publisher,
			Description: group.Description,
			Bookmark:    group.Bookmark,
		})
	}

	return listGroups
}

type UserUpdateMeResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Resolution string `json:"resolution"`
}

func NewUserUpdateMeResponse(user *domain.User) *UserUpdateMeResponse {
	return &UserUpdateMeResponse{
		ID:         user.ID,
		Name:       user.Name,
		Resolution: user.Resolution,
	}
}

type UserGetGroupsByIDResponse struct {
	ID          int       `json:"id"`
	BookTitle   string    `json:"bookTitle"`
	Author      string    `json:"author"`
	MaxPage     int       `json:"maxPage"`
	Publisher   string    `json:"publisher"`
	Description string    `json:"description"`
	Bookmark    int       `json:"bookmark"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewUserGetGroupsByIDResponse(groups []*domain.Group) []*UserGetGroupsByIDResponse {
	var listGroup []*UserGetGroupsByIDResponse
	for _, group := range groups {
		listGroup = append(listGroup, &UserGetGroupsByIDResponse{
			ID:          group.ID,
			BookTitle:   group.BookTitle,
			Author:      group.Author,
			MaxPage:     group.MaxPage,
			Publisher:   group.Publisher,
			Description: group.Description,
			Bookmark:    group.Bookmark,
			CreatedAt:   group.CreatedAt,
			UpdatedAt:   group.UpdatedAt,
		})
	}

	return listGroup
}

type UserUpdateResponse struct {
	ID         int                   `json:"id"`
	Name       string                `json:"name"`
	Resolution string                `json:"resolution"`
	Role       domain.UserRoleType   `json:"role"`
	Social     domain.UserSocialType `json:"social"`
	SocialSub  string                `json:"socialSub"`
	CreatedAt  time.Time             `json:"createdAt"`
	UpdatedAt  time.Time             `json:"updatedAt"`
}

func NewUserUpdateResponse(user *domain.User) *UserUpdateResponse {
	return &UserUpdateResponse{
		ID:         user.ID,
		Name:       user.Name,
		Resolution: user.Resolution,
		Role:       user.Role,
		Social:     user.Social,
		SocialSub:  user.SocialSub,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}
