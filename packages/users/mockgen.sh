mockgen -source=internal/domain/entities/user.go -destination=internal/domain/entities/mocks/user.go -package=mocks

mockgen -destination=mocks/user_mock.go -package=mocks . UserInterface
