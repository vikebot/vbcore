package vbcore

import (
	"reflect"
	"testing"

	"github.com/mohae/deepcopy"
)

func TestUser_SafeUser(t *testing.T) {
	type fields struct {
		ID               int
		Permission       int
		PermissionString string
		Username         string
		Name             string
		Emails           []Email
		Bio              string
		Location         string
		Web              []string
		Company          string
		Social           map[string]string
		OAuth            map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   *SafeUser
	}{
		{"Simple", fields{
			ID:               100000,
			Permission:       PermissionAdmin,
			PermissionString: PermissionAdminString,
			Username:         "sample_admin",
			Name:             "Samle Admin Account",
			Emails: []Email{
				Email{
					Email:  "1@x.com",
					Status: EmailLinked,
					Public: false,
				},
			},
			Bio:      "This is the bio of the sample admin account",
			Location: "Somewher on the Serverfarm",
			Web: []string{
				"admin.vikebot.com",
				"adminuniverse.com",
			},
			Company: "@vikebot",
			Social: map[string]string{
				"github":   "https://github.com/admin",
				"facebook": "https://facebook.com/admin",
			},
			OAuth: map[string]string{
				"github":   "10000000",
				"facebook": "20000000",
			},
		}, &SafeUser{
			ID:               100000,
			Permission:       PermissionAdmin,
			PermissionString: PermissionAdminString,
			Username:         "sample_admin",
			Name:             "Samle Admin Account",
			Emails: []Email{
				Email{
					Email:  "1@x.com",
					Status: EmailLinked,
					Public: false,
				},
			},
			Bio:      "This is the bio of the sample admin account",
			Location: "Somewher on the Serverfarm",
			Web: []string{
				"admin.vikebot.com",
				"adminuniverse.com",
			},
			Company: "@vikebot",
			Social: map[string]string{
				"github":   "https://github.com/admin",
				"facebook": "https://facebook.com/admin",
			},
			OAuth: map[string]string{
				"github":   "10000000",
				"facebook": "20000000",
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:               &tt.fields.ID,
				Permission:       &tt.fields.Permission,
				PermissionString: &tt.fields.PermissionString,
				Username:         &tt.fields.Username,
				Name:             &tt.fields.Name,
				Emails:           tt.fields.Emails,
				Bio:              &tt.fields.Bio,
				Location:         &tt.fields.Location,
				Web:              tt.fields.Web,
				Company:          &tt.fields.Company,
				Social:           tt.fields.Social,
				OAuth:            tt.fields.OAuth,
			}

			su := u.SafeUser()
			if su == nil {
				t.Errorf("User.SafeUser() = nil")
				return
			}

			// Compare all fields
			if *u.ID != su.ID {
				t.Errorf("User.SafeUser().ID = %v, want %v", su.ID, u.ID)
			}
			if *u.Permission != su.Permission {
				t.Errorf("User.SafeUser().Permission = %v, want %v", su.Permission, u.Permission)
			}
			if *u.PermissionString != su.PermissionString {
				t.Errorf("User.SafeUser().PermissionString = %v, want %v", su.PermissionString, u.PermissionString)
			}
			if *u.Username != su.Username {
				t.Errorf("User.SafeUser().Username = %v, want %v", su.Username, u.Username)
			}
			if *u.Name != su.Name {
				t.Errorf("User.SafeUser().Name = %v, want %v", su.Name, u.Name)
			}
			if !reflect.DeepEqual(u.Emails, su.Emails) {
				t.Errorf("User.SafeUser().Emails = %v, want %v", su.Emails, u.Emails)
			}
			if *u.Bio != su.Bio {
				t.Errorf("User.SafeUser().Bio = %v, want %v", su.Bio, u.Bio)
			}
			if *u.Location != su.Location {
				t.Errorf("User.SafeUser().Location = %v, want %v", su.Location, u.Location)
			}
			if !reflect.DeepEqual(u.Web, su.Web) {
				t.Errorf("User.SafeUser().Web = %v, want %v", su.Web, u.Web)
			}
			if *u.Company != su.Company {
				t.Errorf("User.SafeUser().Company = %v, want %v", su.Company, u.Company)
			}
			if !reflect.DeepEqual(u.Social, su.Social) {
				t.Errorf("User.SafeUser().Social = %v, want %v", su.Social, u.Social)
			}
			if !reflect.DeepEqual(u.OAuth, su.OAuth) {
				t.Errorf("User.SafeUser().OAuth = %v, want %v", su.OAuth, u.OAuth)
			}
		})
	}
}

func TestSafeUser_User(t *testing.T) {
	type fields struct {
		ID               int
		Permission       int
		PermissionString string
		Username         string
		Name             string
		Emails           []Email
		Bio              string
		Location         string
		Web              []string
		Company          string
		Social           map[string]string
		OAuth            map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		safe   *SafeUser
	}{
		{"Simple", fields{
			ID:               100000,
			Permission:       PermissionAdmin,
			PermissionString: PermissionAdminString,
			Username:         "sample_admin",
			Name:             "Samle Admin Account",
			Emails: []Email{
				Email{
					Email:  "1@x.com",
					Status: EmailLinked,
					Public: false,
				},
			},
			Bio:      "This is the bio of the sample admin account",
			Location: "Somewher on the Serverfarm",
			Web: []string{
				"admin.vikebot.com",
				"adminuniverse.com",
			},
			Company: "@vikebot",
			Social: map[string]string{
				"github":   "https://github.com/admin",
				"facebook": "https://facebook.com/admin",
			},
			OAuth: map[string]string{
				"github":   "10000000",
				"facebook": "20000000",
			},
		}, &SafeUser{
			ID:               100000,
			Permission:       PermissionAdmin,
			PermissionString: PermissionAdminString,
			Username:         "sample_admin",
			Name:             "Samle Admin Account",
			Emails: []Email{
				Email{
					Email:  "1@x.com",
					Status: EmailLinked,
					Public: false,
				},
			},
			Bio:      "This is the bio of the sample admin account",
			Location: "Somewher on the Serverfarm",
			Web: []string{
				"admin.vikebot.com",
				"adminuniverse.com",
			},
			Company: "@vikebot",
			Social: map[string]string{
				"github":   "https://github.com/admin",
				"facebook": "https://facebook.com/admin",
			},
			OAuth: map[string]string{
				"github":   "10000000",
				"facebook": "20000000",
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := &User{
				ID:               &tt.fields.ID,
				Permission:       &tt.fields.Permission,
				PermissionString: &tt.fields.PermissionString,
				Username:         &tt.fields.Username,
				Name:             &tt.fields.Name,
				Emails:           tt.fields.Emails,
				Bio:              &tt.fields.Bio,
				Location:         &tt.fields.Location,
				Web:              tt.fields.Web,
				Company:          &tt.fields.Company,
				Social:           tt.fields.Social,
				OAuth:            tt.fields.OAuth,
			}

			u := tt.safe.User()
			if u == nil {
				t.Errorf("SafeUser.User() = nil")
				return
			}

			if !reflect.DeepEqual(want, u) {
				t.Errorf("SafeUser().User() = %v, want %v", u, want)
			}
		})
	}
}

func TestSafeUser_MakePublic(t *testing.T) {
	type fields struct {
		ID               int
		Permission       int
		PermissionString string
		Username         string
		Name             string
		Emails           []Email
		Bio              string
		Location         string
		Web              []string
		Company          string
		Social           map[string]string
		OAuth            map[string]string
	}
	tests := []struct {
		name      string
		fields    fields
		wantMails []Email
	}{
		{"Extensive with data keep", fields{
			ID:               100000,
			Permission:       PermissionAdmin,
			PermissionString: PermissionAdminString,
			Username:         "sample_admin",
			Name:             "Samle Admin Account",
			Emails: []Email{
				Email{
					Email:  "1@x.com",
					Status: EmailLinked,
					Public: false,
				},
				Email{
					Email:  "2@x.com",
					Status: EmailLinked,
					Public: true,
				},
				Email{
					Email:  "3@x.com",
					Status: EmailVerified,
					Public: false,
				},
				Email{
					Email:  "4@x.com",
					Status: EmailVerified,
					Public: true,
				},
				Email{
					Email:  "5@x.com",
					Status: EmailPrimary,
					Public: false,
				},
				Email{
					Email:  "6@x.com",
					Status: EmailPrimary,
					Public: true,
				},
			},
			Bio:      "This is the bio of the sample admin account",
			Location: "Somewher on the Serverfarm",
			Web: []string{
				"admin.vikebot.com",
				"adminuniverse.com",
			},
			Company: "@vikebot",
			Social: map[string]string{
				"github":   "https://github.com/admin",
				"facebook": "https://facebook.com/admin",
			},
			OAuth: map[string]string{
				"github":   "10000000",
				"facebook": "20000000",
			},
		}, []Email{
			Email{
				Email:  "4@x.com",
				Status: EmailVerified,
				Public: true,
			},
			Email{
				Email:  "6@x.com",
				Status: EmailPrimary,
				Public: true,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := &SafeUser{
				ID:               tt.fields.ID,
				Permission:       tt.fields.Permission,
				PermissionString: tt.fields.PermissionString,
				Username:         tt.fields.Username,
				Name:             tt.fields.Name,
				Emails:           tt.fields.Emails,
				Bio:              tt.fields.Bio,
				Location:         tt.fields.Location,
				Web:              tt.fields.Web,
				Company:          tt.fields.Company,
				Social:           tt.fields.Social,
				OAuth:            tt.fields.OAuth,
			}

			copy := deepcopy.Copy(su).(*SafeUser)
			copy.MakePublic()

			// Copy email for comparison later
			copyEmails := copy.Emails

			// Remove emails from objects so reflec.DeepEqual works correctly
			copy.Emails = nil
			su.Emails = nil

			// Check if SafeUser is preserved
			if !reflect.DeepEqual(su, copy) {
				t.Errorf("SafeUser.MakePublic() = %v, want %v", copy, su)
			}

			// Check if emails are correct
			if !reflect.DeepEqual(tt.wantMails, copyEmails) {
				t.Errorf("SafeUser.MakePublic().Emails = %v, want %v", copyEmails, tt.wantMails)
			}
		})
	}
}

func TestSafeUser_MakePublicRemoveEmails(t *testing.T) {
	tests := []struct {
		name   string
		emails []Email
		want   []string
	}{
		{"Only private linked", []Email{
			Email{
				Email:  "test@x.com",
				Status: EmailLinked,
				Public: false,
			},
		}, []string{}},
		{"Only public linked", []Email{
			Email{
				Email:  "test@x.com",
				Status: EmailLinked,
				Public: true,
			},
		}, []string{}},
		{"Only private verified", []Email{
			Email{
				Email:  "test@x.com",
				Status: EmailVerified,
				Public: false,
			},
		}, []string{}},
		{"Only public verified", []Email{
			Email{
				Email:  "test@x.com",
				Status: EmailVerified,
				Public: true,
			},
		}, []string{"test@x.com"}},
		{"Only private primary", []Email{
			Email{
				Email:  "test@x.com",
				Status: EmailPrimary,
				Public: false,
			},
		}, []string{}},
		{"Only public primary", []Email{
			Email{
				Email:  "test@x.com",
				Status: EmailPrimary,
				Public: true,
			},
		}, []string{"test@x.com"}},
		{"Mixed", []Email{
			Email{
				Email:  "1@x.com",
				Status: EmailLinked,
				Public: false,
			},
			Email{
				Email:  "2@x.com",
				Status: EmailLinked,
				Public: true,
			},
			Email{
				Email:  "3@x.com",
				Status: EmailVerified,
				Public: false,
			},
			Email{
				Email:  "4@x.com",
				Status: EmailVerified,
				Public: true,
			},
			Email{
				Email:  "5@x.com",
				Status: EmailPrimary,
				Public: false,
			},
			Email{
				Email:  "6@x.com",
				Status: EmailPrimary,
				Public: true,
			},
		}, []string{"4@x.com", "6@x.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := &SafeUser{
				Emails: tt.emails,
			}
			su.MakePublic()

			addresses := []string{}
			for _, e := range su.Emails {
				if !e.Public {
					t.Errorf("SafeUser.MakePublic() contains private email!")
				}
				if e.IsLinked() {
					t.Errorf("SafeUser.MakePublic() contains linked address")
				}
				addresses = append(addresses, e.Email)
			}

			if !reflect.DeepEqual(addresses, tt.want) {
				t.Errorf("SafeUser.MakePublic().Emails = %v, want %v", addresses, tt.emails)
			}
		})
	}
}

func TestSafeUser_PrimaryEmail(t *testing.T) {
	tests := []struct {
		name   string
		Emails []Email
		want   string
	}{
		{"Only primary", []Email{
			Email{
				Email:  "test@x.com",
				Status: EmailPrimary,
			},
		}, "test@x.com"},
		{"Primary after Linked", []Email{
			Email{
				Email:  "false@x.com",
				Status: EmailLinked,
			},
			Email{
				Email:  "test@x.com",
				Status: EmailPrimary,
			},
		}, "test@x.com"},
		{"Primary after Verified", []Email{
			Email{
				Email:  "false@x.com",
				Status: EmailVerified,
			},
			Email{
				Email:  "test@x.com",
				Status: EmailPrimary,
			},
		}, "test@x.com"},
		{"Primary after Linked and Verified", []Email{
			Email{
				Email:  "false@x.com",
				Status: EmailLinked,
			},
			Email{
				Email:  "false@x.com",
				Status: EmailVerified,
			},
			Email{
				Email:  "test@x.com",
				Status: EmailPrimary,
			},
		}, "test@x.com"},
		{"Primary after Linked and Verified", []Email{
			Email{
				Email:  "false@x.com",
				Status: EmailVerified,
			},
			Email{
				Email:  "false@x.com",
				Status: EmailLinked,
			},
			Email{
				Email:  "test@x.com",
				Status: EmailPrimary,
			},
		}, "test@x.com"},
		{"Primary at second - linked first", []Email{
			Email{
				Email:  "false@x.com",
				Status: EmailLinked,
			},
			Email{
				Email:  "test@x.com",
				Status: EmailPrimary,
			},
			Email{
				Email:  "false@x.com",
				Status: EmailVerified,
			},
		}, "test@x.com"},
		{"Primary at second - verified first", []Email{
			Email{
				Email:  "false@x.com",
				Status: EmailVerified,
			},
			Email{
				Email:  "test@x.com",
				Status: EmailPrimary,
			},
			Email{
				Email:  "false@x.com",
				Status: EmailLinked,
			},
		}, "test@x.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := &SafeUser{
				Emails: tt.Emails,
			}
			if got := su.PrimaryEmail(); !reflect.DeepEqual(got.Email, tt.want) {
				t.Errorf("SafeUser.PrimaryEmail().Email = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeUser_PrimaryEmailNil(t *testing.T) {
	tests := []struct {
		name   string
		Emails []Email
	}{
		{"Only linked", []Email{
			Email{
				Email:  "test@x.com",
				Status: EmailLinked,
			},
		}},
		{"Only verified", []Email{
			Email{
				Email:  "test@x.com",
				Status: EmailVerified,
			},
		}},
		{"Linked and verified", []Email{
			Email{
				Email:  "test@x.com",
				Status: EmailLinked,
			},
			Email{
				Email:  "test@x.com",
				Status: EmailVerified,
			},
		}},
		{"Verified and Linked", []Email{
			Email{
				Email:  "test@x.com",
				Status: EmailVerified,
			},
			Email{
				Email:  "test@x.com",
				Status: EmailLinked,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := &SafeUser{
				Emails: tt.Emails,
			}
			if got := su.PrimaryEmail(); got != nil {
				t.Errorf("SafeUser.PrimaryEmail() = %v, want nil", got)
			}
		})
	}
}
