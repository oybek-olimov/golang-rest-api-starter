package service_test

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	"go-crud-api/internal/model"
// 	"go-crud-api/internal/service"
// )

// // 🔧 Fake store (interfeysga mos)
// type fakeUserStore struct {
// 	users map[string]*model.User
// }

// func (f *fakeUserStore) Create(user *model.User) error {
// 	if _, exists := f.users[user.Email]; exists {
// 		return errors.New("user already exists")
// 	}
// 	f.users[user.Email] = user
// 	return nil
// }

// func (f *fakeUserStore) GetByEmail(email string) (*model.User, error) {
// 	user, ok := f.users[email]
// 	if !ok {
// 		return nil, errors.New("user not found")
// 	}
// 	return user, nil
// }

// //  Dummy hasher (interfeysga mos)
// type dummyHasher struct{}

// func (d *dummyHasher) Hash(password string) (string, error) {
// 	return "hashed:" + password, nil
// }

// func (d *dummyHasher) Compare(hash, password string) bool {
// 	return hash == "hashed:"+password
// }

// //  Test 1: Foydalanuvchi muvaffaqiyatli ro‘yxatdan o‘tadi
// func TestRegister_Success(t *testing.T) {
// 	store := &fakeUserStore{users: make(map[string]*model.User)}
// 	hasher := &dummyHasher{}
// 	svc := service.NewUserService(store, hasher)

// 	ctx := context.Background()
// 	username := "qwerty"
// 	email := "test@example.com"
// 	password := "123456"

// 	user, err := svc.Register(ctx, username, email, password)
// 	if err != nil {
// 		t.Fatalf("expected no error, got: %v", err)
// 	}
// 	if user.Email != email {
// 		t.Errorf("expected email %s, got %s", email, user.Email)
// 	}
// }

// //  Test 2: Foydalanuvchi allaqachon mavjud bo‘lsa
// func TestRegister_UserAlreadyExists(t *testing.T) {
// 	store := &fakeUserStore{
// 		users: map[string]*model.User{
// 			"test@example.com": {Email: "test@example.com"},
// 		},
// 	}
// 	hasher := &dummyHasher{}
// 	svc := service.NewUserService(store, hasher)

// 	ctx := context.Background()
// 	username := "existing"
// 	email := "test@example.com"
// 	password := "123456"

// 	_, err := svc.Register(ctx, username, email, password)
// 	if err == nil {
// 		t.Fatal("expected error, got nil")
// 	}
// }
