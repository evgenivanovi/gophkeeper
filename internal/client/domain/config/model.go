package config

import (
	"github.com/evgenivanovi/gophkeeper/pkg/time"
	"golang.org/x/exp/maps"
)

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type ConfigObject struct {
	Current string       `yaml:"current" json:"current"`
	Users   []UserObject `yaml:"users" json:"users"`
}

//goland:noinspection GoNameStartsWithPackageName
type UserObject struct {
	User    string        `yaml:"user" json:"user"`
	Session string        `yaml:"session" json:"session"`
	Secrets SecretsObject `yaml:"secrets" json:"secrets"`
}

//goland:noinspection GoNameStartsWithPackageName
type SecretsObject struct {
	Access  SecretObject `yaml:"access" json:"access"`
	Refresh SecretObject `yaml:"refresh" json:"refresh"`
}

//goland:noinspection GoNameStartsWithPackageName
type SecretObject struct {
	Data       string       `yaml:"data" json:"data"`
	Expiration time.Instant `yaml:"expiration" json:"expiration"`
}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type ConfigAction func(ConfigObject) ConfigObject

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type ConfigObjectOperations struct {
	origin ConfigObject
}

func NewConfigObjectOperations(obj ConfigObject) ConfigObjectOperations {
	return ConfigObjectOperations{
		origin: obj,
	}
}

func (op *ConfigObjectOperations) Get() ConfigObject {
	return op.origin
}

func (op *ConfigObjectOperations) Users() map[string]UserObject {
	users := make(map[string]UserObject)
	for _, obj := range op.origin.Users {
		users[obj.User] = obj
	}
	return users
}

func (op *ConfigObjectOperations) GetUser(user string) (UserObject, bool) {
	if obj, ok := op.Users()[user]; ok {
		return obj, true
	}
	return UserObject{}, false
}

func (op *ConfigObjectOperations) UpdateContext(context string) {
	op.origin.Current = context
}

func (op *ConfigObjectOperations) UpdateUser(user UserObject) {
	users := op.Users()
	users[user.User] = user

	op.refreshUsers(maps.Values(users))
}

func (op *ConfigObjectOperations) RemoveUser(user string) {
	users := op.Users()
	delete(users, user)

	op.refreshUsers(maps.Values(users))
}

func (op *ConfigObjectOperations) refreshUsers(users []UserObject) {
	op.origin.Users = users
}

func (op *ConfigObjectOperations) GetSecrets(user string) (SecretsObject, bool) {
	users := op.Users()
	if obj, ok := users[user]; ok {
		return obj.Secrets, true
	}
	return SecretsObject{}, false
}

func (op *ConfigObjectOperations) UpdateAccessSecret(user string, secret SecretObject) {
	users := op.Users()
	if obj, ok := users[user]; ok {
		obj.Secrets.Access = secret
		op.refreshUsers(maps.Values(users))
	}
}

func (op *ConfigObjectOperations) UpdateRefreshSecret(user string, secret SecretObject) {
	users := op.Users()
	if obj, ok := users[user]; ok {
		obj.Secrets.Refresh = secret
		op.refreshUsers(maps.Values(users))
	}
}

/* __________________________________________________ */
