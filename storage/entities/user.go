package entities

type User struct {
	Id             string                 `bson:"_id"`
	Email          string                 `bson:"email"`
	Name           string                 `bson:"name"`
	Family         string                 `bson:"family"`
	NormalizedName string                 `bson:"normalized_name"`
	Phone          string                 `bson:"phone"`
	changes        map[string]interface{} // changed fields during update
}

func (u *User) addUpdatedField(f string, v interface{}) {
	if u.changes == nil {
		u.changes = make(map[string]interface{})
	}
	u.changes[f] = v
}

func (u *User) SetName(name string) {
	if len(name) != 0 {
		u.Name = name
	}
	u.addUpdatedField("name", name)
}

func (u *User) SetFamily(family string) {
	if len(family) != 0 {
		u.Family = family
	}
	u.addUpdatedField("family", family)
}
func (u *User) SetEmail(email string) {
	if len(email) != 0 {
		u.Email = email
	}
	u.addUpdatedField("email", email)
}

func (u *User) SetPhone(phone string) {
	if len(phone) != 0 {
		u.Phone = phone
	}
	u.addUpdatedField("phone", phone)
}

func (u *User) GetChanges() map[string]interface{} {
	return u.changes
}
