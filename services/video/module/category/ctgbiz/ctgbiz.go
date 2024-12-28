package ctgbiz

import (
	"errors"
	"github.com/caovanhoang63/hiholive/services/video/module/category/ctgmodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type CategoryBiz interface {
	CreateCategory(ctx context.Context, requester core.Requester, create *ctgmodel.CategoryCreate) error
	UpdateCategory(ctx context.Context, requester core.Requester, id int, update *ctgmodel.CategoryUpdate) error
	DeleteCategory(ctx context.Context, requester core.Requester, id int) error
	FindCategory(ctx context.Context, id int) (*ctgmodel.Category, error)
	FindCategories(ctx context.Context, filter *ctgmodel.CategoryFilter, paging *core.Paging) ([]ctgmodel.Category, error)
	IncreaseTotalContent(ctx context.Context, id int) error
}

type CategoryRepo interface {
	IncreaseTotalContent(ctx context.Context, id int) error
	CreateCategory(ctx context.Context, create *ctgmodel.CategoryCreate) error
	UpdateCategory(ctx context.Context, id int, update *ctgmodel.CategoryUpdate) error
	DeleteCategory(ctx context.Context, id int) error
	FindCategory(ctx context.Context, id int) (*ctgmodel.Category, error)
	FindCategories(ctx context.Context, filter *ctgmodel.CategoryFilter, paging *core.Paging) ([]ctgmodel.Category, error)
}

type categoryBiz struct {
	ctgRepo CategoryRepo
}

func NewCategoryBiz(ctgRepo CategoryRepo) *categoryBiz { return &categoryBiz{ctgRepo: ctgRepo} }

func (c *categoryBiz) IncreaseTotalContent(ctx context.Context, id int) error {
	if err := c.ctgRepo.IncreaseTotalContent(ctx, id); err != nil {
		log.Println(err)
		return core.ErrInternalServerError.WithWrap(err)
	}
	return nil
}

func (c *categoryBiz) CreateCategory(ctx context.Context, requester core.Requester, create *ctgmodel.CategoryCreate) error {
	if requester.GetRole() != "admin" {
		return core.ErrForbidden
	}

	if field, err := core.Validator.ValidateField(create); err != nil {
		return core.ErrInvalidInput(field)
	}

	// Call the repository layer to create the category
	if err := c.ctgRepo.CreateCategory(ctx, create); err != nil {
		return core.ErrInternalServerError.WithWrap(err)
	}

	return nil
}

func (c *categoryBiz) UpdateCategory(ctx context.Context, requester core.Requester, id int, update *ctgmodel.CategoryUpdate) error {

	if requester.GetRole() != "admin" {
		return core.ErrForbidden
	}
	if field, err := core.Validator.ValidateField(update); err != nil {
		return core.ErrInvalidInput(field)
	}

	old, err := c.ctgRepo.FindCategory(ctx, id)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return core.ErrNotFound
		}
		return core.ErrInternalServerError.WithWrap(err)
	}

	if old.Status == 0 {
		return core.ErrBadRequest.WithError("category deleted")
	}

	if id <= 0 {
		return core.ErrInvalidInput("id")
	}

	if err = c.ctgRepo.UpdateCategory(ctx, id, update); err != nil {
		return core.ErrInternalServerError.WithWrap(err)
	}

	return nil
}

func (c *categoryBiz) DeleteCategory(ctx context.Context, requester core.Requester, id int) error {
	if requester.GetRole() != "admin" {
		return core.ErrForbidden
	}

	old, err := c.ctgRepo.FindCategory(ctx, id)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return core.ErrNotFound
		}
		return core.ErrInternalServerError.WithWrap(err)
	}

	if old.Status == 0 {
		return core.ErrBadRequest.WithError("category deleted")
	}

	if id <= 0 {
		return core.ErrInvalidInput("id")
	}

	if requester.GetRole() != "admin" {
		return core.ErrForbidden
	}

	if err = c.ctgRepo.DeleteCategory(ctx, id); err != nil {
		return core.ErrInternalServerError.WithWrap(err)
	}

	return nil
}

func (c *categoryBiz) FindCategory(ctx context.Context, id int) (*ctgmodel.Category, error) {
	// Validate the input
	if id <= 0 {
		return nil, core.ErrInvalidInput("id")
	}

	category, err := c.ctgRepo.FindCategory(ctx, id)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, core.ErrNotFound
		}
		return nil, core.ErrInternalServerError.WithWrap(err)
	}

	return category, nil
}

func (c *categoryBiz) FindCategories(ctx context.Context, filter *ctgmodel.CategoryFilter, paging *core.Paging) ([]ctgmodel.Category, error) {
	// Call the repository layer to find categories
	categories, err := c.ctgRepo.FindCategories(ctx, filter, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.WithWrap(err)
	}

	return categories, nil
}
