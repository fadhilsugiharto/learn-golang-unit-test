package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"learn-golang-unit-test/entity"
	"learn-golang-unit-test/repository"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// test scenario where category is null in db
	// make a mock program to return nil when id is 1
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetFound(t *testing.T) {
	// test scenario where category value exist
	category := entity.Category{
		Id:   "2",
		Name: "KunK",
	}
	// make a mock program to return category when id is 2
	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, category.Id, result.Id, "Id is not the same")
	assert.Equal(t, category.Name, result.Name, "Name is not the same")
}
