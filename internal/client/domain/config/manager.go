package config

import (
	"context"
)

/* __________________________________________________ */

type Manager interface {
	GetCurrentUser(
		ctx context.Context, path string,
	) (string, error)

	SetCurrentUser(
		ctx context.Context, path string, user string,
	) error

	SetCurrentUserAction(
		ctx context.Context, user string,
	) ConfigAction

	AddUser(
		ctx context.Context, path string, user UserObject,
	) error

	AddUserAction(
		ctx context.Context, user UserObject,
	) ConfigAction

	Within(
		ctx context.Context, path string, actions ...ConfigAction,
	) error
}

type ManagerService struct {
	reader Reader
	writer Writer
}

func ProvideManagerService(reader Reader, writer Writer) *ManagerService {
	return &ManagerService{
		reader: reader,
		writer: writer,
	}
}

func (m *ManagerService) GetCurrentUser(
	ctx context.Context, path string,
) (string, error) {
	obj, err := m.reader.Read(path)
	if err != nil {
		return "", err
	}
	return obj.Current, nil
}

func (m *ManagerService) SetCurrentUser(
	ctx context.Context, path string, user string,
) error {
	return m.Within(ctx, path, m.SetCurrentUserAction(ctx, user))
}

func (m *ManagerService) SetCurrentUserAction(
	ctx context.Context, user string,
) ConfigAction {
	return func(origin ConfigObject) ConfigObject {
		op := NewConfigObjectOperations(origin)
		op.UpdateContext(user)
		return op.Get()
	}
}

func (m *ManagerService) AddUser(
	ctx context.Context, path string, user UserObject,
) error {
	return m.Within(ctx, path, m.AddUserAction(ctx, user))
}

func (m *ManagerService) AddUserAction(
	ctx context.Context, user UserObject,
) ConfigAction {
	return func(origin ConfigObject) ConfigObject {
		op := NewConfigObjectOperations(origin)
		op.UpdateUser(user)
		return op.Get()
	}
}

func (m *ManagerService) Within(
	ctx context.Context, path string, actions ...ConfigAction,
) error {

	origin, err := m.reader.Read(path)
	if err != nil {
		return err
	}

	target := origin
	for _, action := range actions {
		target = action(target)
	}

	err = m.writer.Write(target, path)
	if err != nil {
		return err
	}

	return nil

}

/* __________________________________________________ */
