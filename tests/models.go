package tests

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	contractsorm "github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/support/carbon"
)

type contextKey int

const testContextKey contextKey = 0

type Model struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Timestamps
}

type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

type Timestamps struct {
	CreatedAt carbon.DateTime `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	UpdatedAt carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
}

type User struct {
	Model
	SoftDeletes
	Name    string
	Bio     *string
	Avatar  string
	Address *Address
	Books   []*Book
	House   *House   `gorm:"polymorphic:Houseable"`
	Phones  []*Phone `gorm:"polymorphic:Phoneable"`
	Roles   []*Role  `gorm:"many2many:role_user"`
	age     int
}

func (u *User) DispatchesEvents() map[contractsorm.EventType]func(contractsorm.Event) error {
	return map[contractsorm.EventType]func(contractsorm.Event) error{
		contractsorm.EventCreating: func(event contractsorm.Event) error {
			name := event.GetAttribute("name")
			if name != nil {
				if name.(string) == "event_creating_name" {
					event.SetAttribute("avatar", "event_creating_avatar")
				}
				if name.(string) == "event_creating_by_map_name" {
					event.SetAttribute("avatar", "event_creating_by_map_avatar1")
				}
				if name.(string) == "event_creating_FirstOrCreate_name" {
					event.SetAttribute("avatar", "event_creating_FirstOrCreate_avatar")
				}
				if name.(string) == "event_creating_omit_create_name" {
					event.SetAttribute("avatar", "event_creating_omit_create_avatar")
				}
				if name.(string) == "event_creating_select_create_name" {
					event.SetAttribute("avatar", "event_creating_select_create_avatar")
				}
				if name.(string) == "event_creating_save_name" {
					event.SetAttribute("avatar", "event_creating_save_avatar")
				}
				if name.(string) == "event_creating_IsDirty_name" {
					if event.IsDirty("name") {
						event.SetAttribute("avatar", "event_creating_IsDirty_avatar")
					}
				}
				if name.(string) == "event_context" {
					val := event.Context().Value(testContextKey)
					event.SetAttribute("avatar", val.(string))
				}
				if name.(string) == "event_query" {
					_ = event.Query().Create(&User{Name: "event_query1"})
				}
			}

			return nil
		},
		contractsorm.EventCreated: func(event contractsorm.Event) error {
			name := event.GetAttribute("name")
			if name != nil {
				if name.(string) == "event_created_name" {
					id := event.GetAttribute("ID")
					event.SetAttribute("avatar", fmt.Sprintf("event_created_avatar_%d", id))
				}
				if name.(string) == "event_created_by_map_name" {
					event.SetAttribute("avatar", "event_created_by_map_avatar1")
				}
				if name.(string) == "event_created_FirstOrCreate_name" {
					id := event.GetAttribute("ID")
					event.SetAttribute("avatar", fmt.Sprintf("event_created_FirstOrCreate_avatar_%d", id))
				}
				if name.(string) == "event_created_omit_create_name" {
					id := event.GetAttribute("ID")
					event.SetAttribute("avatar", fmt.Sprintf("event_created_omit_create_avatar_%d", id))
				}
				if name.(string) == "event_created_select_create_name" {
					id := event.GetAttribute("ID")
					event.SetAttribute("avatar", fmt.Sprintf("event_created_select_create_avatar_%d", id))
				}
				if name.(string) == "event_created_save_name" {
					id := event.GetAttribute("ID")
					event.SetAttribute("avatar", fmt.Sprintf("event_created_save_avatar_%d", id))
				}
			}

			return nil
		},
		contractsorm.EventSaving: func(event contractsorm.Event) error {
			name := event.GetAttribute("name")
			switch name.(type) {
			case string:
				if name == "event_saving_create_name" {
					event.SetAttribute("avatar", "event_saving_create_avatar")
				}
				if name == "event_saving_create_by_map_name" {
					event.SetAttribute("avatar", "event_saving_create_by_map_avatar1")
				}
				if name == "event_saving_save_name" {
					event.SetAttribute("avatar", "event_saving_save_avatar")
				}
				if name == "event_saving_FirstOrCreate_name" {
					event.SetAttribute("avatar", "event_saving_FirstOrCreate_avatar")
				}
				if name == "event_saving_omit_create_name" {
					event.SetAttribute("avatar", "event_saving_omit_create_avatar")
				}
				if name == "event_saving_select_create_name" {
					event.SetAttribute("avatar", "event_saving_select_create_avatar")
				}
				if name == "event_save_without_name" {
					event.SetAttribute("avatar", "event_save_without_avatar")
				}
				if name == "event_save_quietly_name" {
					event.SetAttribute("avatar", "event_save_quietly_avatar")
				}
				if name == "event_saving_IsDirty_name" {
					if event.IsDirty("name") {
						event.SetAttribute("avatar", "event_saving_IsDirty_avatar")
					}
				}
			}

			avatar := event.GetAttribute("avatar")
			if avatar != nil && avatar.(string) == "event_saving_single_update_avatar" {
				event.SetAttribute("avatar", "event_saving_single_update_avatar1")
			}

			return nil
		},
		contractsorm.EventSaved: func(event contractsorm.Event) error {
			name := event.GetAttribute("name")
			switch name.(type) {
			case string:
				if name == "event_saved_create_name" {
					event.SetAttribute("avatar", "event_saved_create_avatar")
				}
				if name == "event_saved_create_by_map_name" {
					event.SetAttribute("avatar", "event_saved_create_by_map_avatar1")
				}
				if name == "event_saved_omit_create_name" {
					event.SetAttribute("avatar", "event_saved_omit_create_avatar")
				}
				if name == "event_saved_select_create_name" {
					event.SetAttribute("avatar", "event_saved_select_create_avatar")
				}
				if name == "event_saved_save_name" {
					event.SetAttribute("avatar", "event_saved_save_avatar")
				}
				if name == "event_saved_FirstOrCreate_name" {
					event.SetAttribute("avatar", "event_saved_FirstOrCreate_avatar")
				}
				if name == "event_save_without_name" {
					event.SetAttribute("avatar", "event_saved_without_avatar")
				}
				if name == "event_save_quietly_name" {
					event.SetAttribute("avatar", "event_saved_quietly_avatar")
				}
			}

			avatar := event.GetAttribute("avatar")
			if avatar != nil && avatar.(string) == "event_saved_map_update_avatar" {
				event.SetAttribute("avatar", "event_saved_map_update_avatar1")
			}

			return nil
		},
		contractsorm.EventUpdating: func(event contractsorm.Event) error {
			name := event.GetAttribute("name")
			switch name.(type) {
			case string:
				if name == "event_updating_create_name" {
					event.SetAttribute("avatar", "event_updating_create_avatar")
				}
				if name == "event_updating_save_name" {
					event.SetAttribute("avatar", "event_updating_save_avatar")
				}
				if name == "event_updating_single_update_IsDirty_name1" {
					if event.IsDirty("name") {
						name := event.GetAttribute("name")
						if name != "event_updating_single_update_IsDirty_name1" {
							return errors.New("error")
						}
						event.SetAttribute("avatar", "event_updating_single_update_IsDirty_avatar")
					}
				}
				if name == "event_updating_map_update_IsDirty_name1" {
					if event.IsDirty("name") {
						name := event.GetAttribute("name")
						if name != "event_updating_map_update_IsDirty_name1" {
							return errors.New("error")
						}
						event.SetAttribute("avatar", "event_updating_map_update_IsDirty_avatar")
					}
				}
				if name == "event_updating_model_update_IsDirty_name1" {
					if event.IsDirty("name") {
						name := event.GetAttribute("name")
						if name != "event_updating_model_update_IsDirty_name1" {
							return errors.New("error")
						}
						event.SetAttribute("avatar", "event_updating_model_update_IsDirty_avatar")
					}
				}
			}

			avatar := event.GetAttribute("avatar")
			if avatar != nil {
				if avatar.(string) == "event_updating_save_avatar" {
					event.SetAttribute("avatar", "event_updating_save_avatar1")
				}
				if avatar.(string) == "event_updating_model_update_avatar" {
					id := event.GetOriginal("ID")
					event.SetAttribute("avatar", fmt.Sprintf("event_updating_model_update_avatar_%d", id))
				}
			}

			return nil
		},
		contractsorm.EventUpdated: func(event contractsorm.Event) error {
			name := event.GetAttribute("name")
			switch name.(type) {
			case string:
				if name == "event_updated_create_name" {
					event.SetAttribute("avatar", "event_updated_create_avatar")
				}
				if name == "event_updated_save_name" {
					event.SetAttribute("avatar", "event_updated_save_avatar")
				}
			}

			avatar := event.GetAttribute("avatar")
			if avatar != nil {
				if avatar.(string) == "event_updated_save_avatar" {
					event.SetAttribute("avatar", "event_updated_save_avatar1")
				}
				if avatar.(string) == "event_updated_model_update_avatar" {
					event.SetAttribute("avatar", "event_updated_model_update_avatar1")
				}
			}

			return nil
		},
		contractsorm.EventDeleting: func(event contractsorm.Event) error {
			name := event.GetAttribute("name")
			if name != nil && name.(string) == "event_deleting_name" {
				return errors.New("deleting error")
			}

			return nil
		},
		contractsorm.EventDeleted: func(event contractsorm.Event) error {
			name := event.GetAttribute("name")
			if name != nil && name.(string) == "event_deleted_name" {
				return errors.New("deleted error")
			}

			return nil
		},
		contractsorm.EventForceDeleting: func(event contractsorm.Event) error {
			name := event.GetAttribute("name")
			if name != nil && name.(string) == "event_force_deleting_name" {
				return errors.New("force deleting error")
			}

			return nil
		},
		contractsorm.EventForceDeleted: func(event contractsorm.Event) error {
			name := event.GetAttribute("name")
			if name != nil && name.(string) == "event_force_deleted_name" {
				return errors.New("force deleted error")
			}

			return nil
		},
		contractsorm.EventRetrieved: func(event contractsorm.Event) error {
			name := event.GetAttribute("name")
			if name != nil && name.(string) == "event_retrieved_name" {
				event.SetAttribute("name", "event_retrieved_name1")
			}

			return nil
		},
		contractsorm.EventRestored: func(event contractsorm.Event) error {
			name := event.GetAttribute("name")
			if name != nil && name.(string) == "event_restored_name" {
				event.SetAttribute("name", "event_restored_name1")
			}

			return nil
		},
		contractsorm.EventRestoring: func(event contractsorm.Event) error {
			name := event.GetAttribute("name")
			if name != nil && name.(string) == "event_restoring_name" {
				event.SetAttribute("name", "event_restoring_name1")
			}

			return nil
		},
	}
}

type UserObserver struct{}

func (u *UserObserver) Retrieved(event contractsorm.Event) error {
	return errors.New("retrieved")
}

func (u *UserObserver) Creating(event contractsorm.Event) error {
	return errors.New("creating")
}

func (u *UserObserver) Created(event contractsorm.Event) error {
	return errors.New("created")
}

func (u *UserObserver) Updating(event contractsorm.Event) error {
	return errors.New("updating")
}

func (u *UserObserver) Updated(event contractsorm.Event) error {
	return errors.New("updated")
}

func (u *UserObserver) Saving(event contractsorm.Event) error {
	return errors.New("saving")
}

func (u *UserObserver) Saved(event contractsorm.Event) error {
	return errors.New("saved")
}

func (u *UserObserver) Deleting(event contractsorm.Event) error {
	return errors.New("deleting")
}

func (u *UserObserver) Deleted(event contractsorm.Event) error {
	return errors.New("deleted")
}

func (u *UserObserver) ForceDeleting(event contractsorm.Event) error {
	return errors.New("forceDeleting")
}

func (u *UserObserver) ForceDeleted(event contractsorm.Event) error {
	return errors.New("forceDeleted")
}

type Role struct {
	Model
	Name  string
	Users []*User `gorm:"many2many:role_user"`
}

type Address struct {
	Model
	UserID   uint
	Name     string
	Province string
	User     *User
}

type Book struct {
	Model
	UserID uint
	Name   string
	User   *User
	Author *Author
}

type Author struct {
	Model
	BookID uint
	Name   string
}

type House struct {
	Model
	Name          string
	HouseableID   uint
	HouseableType string
}

func (h *House) Factory() string {
	return "house"
}

type Phone struct {
	Model
	Name          string
	PhoneableID   uint
	PhoneableType string
}

type Product struct {
	Model
	SoftDeletes
	Name string
}

func (p *Product) Connection() string {
	return "sqlite"
}

type Review struct {
	Model
	SoftDeletes
	Body string
}

func (r *Review) Connection() string {
	return ""
}

type People struct {
	Model
	SoftDeletes
	Body string
}

func (p *People) Connection() string {
	return "dummy"
}

type Person struct {
	Model
	SoftDeletes
	Name string
}

func (p *Person) Connection() string {
	return "dummy"
}

type Box struct {
	Model
	SoftDeletes
	Name string
}

func (p *Box) Connection() string {
	return "postgres"
}

type Schema struct {
	Model
	Name string
}

func (r *Schema) TableName() string {
	return "goravel.schemas"
}
