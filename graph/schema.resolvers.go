package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"koriebruh/graphql-product/domain"
	"koriebruh/graphql-product/graph/model"
	"strconv"
)

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, name string, description *string, price float64, categoryID string) (*model.Product, error) {
	id, err := strconv.ParseUint(categoryID, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid categoryID: %v", err)
	}

	newProduct := domain.Product{
		Name: name,
		Description: func() string {
			if description != nil {
				return *description
			}
			return ""
		}(),
		Price:      price,
		CategoryID: uint(id),
	}

	// validation category
	var existingCategory domain.Category
	if err := r.DB.WithContext(ctx).First(&existingCategory, id).Error; err != nil {
		return nil, fmt.Errorf("category with ID %s does not exist", categoryID)
	}

	if err := r.DB.WithContext(ctx).Create(&newProduct).Error; err != nil {
		return nil, fmt.Errorf("failed to create product: %v", err)
	}

	return &model.Product{
		ID:          fmt.Sprint(newProduct.ID), // Konversi uint ke string
		Name:        newProduct.Name,
		Description: description,
		Price:       newProduct.Price,
		Category:    &model.Category{ID: fmt.Sprint(newProduct.CategoryID)},
	}, nil

}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, name string) (*model.Category, error) {
	newCategory := domain.Category{
		Name: name,
	}

	if err := r.DB.WithContext(ctx).Create(&newCategory).Error; err != nil {
		return nil, fmt.Errorf("failed to create new category %v", err)
	}

	return &model.Category{
		ID:   strconv.Itoa(int(newCategory.ID)),
		Name: newCategory.Name,
	}, nil

}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	var dataProduct []domain.Product
	if err := r.DB.WithContext(ctx).Find(&dataProduct).Error; err != nil {
		return nil, fmt.Errorf("failed to get field product")
	}

	// mapping data from db
	var products []*model.Product
	for _, p := range dataProduct {
		products = append(products, &model.Product{
			ID:          fmt.Sprint(p.ID),
			Name:        p.Name,
			Description: &p.Description,
			Price:       p.Price,
			Category:    &model.Category{ID: fmt.Sprint(p.CategoryID)},
		})

	}

	return products, nil

}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	var product *domain.Product

	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid product ID: %v", err)
	}

	if err = r.DB.WithContext(ctx).Find(&product, uint(productID)).Error; err != nil {
		// check if data no exist
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("product not found")
		}
		return nil, fmt.Errorf("failed to fetch product: %v", err)
	}

	return &model.Product{
		ID:          fmt.Sprint(product.ID),
		Name:        product.Name,
		Description: &product.Description,
		Price:       product.Price,
		Category:    &model.Category{ID: fmt.Sprint(product.CategoryID)},
	}, nil

}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	var categoriesDB []*domain.Category

	if err := r.DB.WithContext(ctx).Find(&categoriesDB).Error; err != nil {
		return nil, fmt.Errorf("failed to get field category")
	}

	var categories []*model.Category
	for _, category := range categoriesDB {
		categories = append(categories, &model.Category{
			ID:   fmt.Sprint(category.ID),
			Name: category.Name,
		})
	}

	return categories, nil

}

// Category is the resolver for the category field.
func (r *queryResolver) Category(ctx context.Context, id string) (*model.Category, error) {
	var category domain.Category

	categoryID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid product ID: %v", err)
	}

	if err = r.DB.WithContext(ctx).Find(&category, uint(categoryID)).Error; err != nil {
		// check if data no exist
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("product not found")
		}
		return nil, fmt.Errorf("failed to fetch product: %v", err)
	}

	return &model.Category{
		ID:   fmt.Sprint(category.ID),
		Name: category.Name,
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}
*/
