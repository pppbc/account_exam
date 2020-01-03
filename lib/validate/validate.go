package validate

import (
	"reflect"
	"regexp"

	en "github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

var (
	Default      = validator.New()
	DefaultTrans = func() ut.Translator {
		var found bool
		en := en.New()
		trans, found := ut.New(en, en).GetTranslator("en_US")
		if !found {
			panic(found)
		}
		err := en_translations.RegisterDefaultTranslations(Default, trans)
		if err != nil {
			panic(err)
		}
		return trans
	}()

	mobileRegex   = regexp.MustCompile(`^[1]([3-9])[0-9]{9}$`)
	phoneRegex    = regexp.MustCompile(`^[0-9-]{7,20}$`)
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,18}[a-zA-Z0-9]{1}$`)
	codeRegex     = regexp.MustCompile(`^[a-z][a-z0-9-]{1,18}[a-z0-9]$`)
	emailRegex    = regexp.MustCompile(`^([A-Za-z0-9])+@([A-Za-z0-9])+\.([A-Za-z]{2,4})$`)
)

func init() {
	Default.RegisterValidation("mobile", validateMobile)
	Default.RegisterValidation("username", validateUsername)
	Default.RegisterValidation("code", validateCode)
	Default.RegisterValidation("phone", validatePhone)
	Default.RegisterValidation("email", validateEmail)
	Default.RegisterTranslation("username", DefaultTrans,
		func(ut ut.Translator) (err error) {
			if err = ut.Add("username", "{0} is a invalid field", false); err != nil {
				return
			}
			return
		},
		func(ut ut.Translator, fe validator.FieldError) (k string, p []string, t string) {
			k = fe.Tag()
			p = []string{fe.Field()}
			t, err := ut.T(fe.Tag(), fe.Field())
			if err != nil {
				return k, p, fe.(error).Error()
			}

			return k, p, t
		})
}

func validateMobile(fl validator.FieldLevel) bool {
	return mobileRegex.MatchString(fl.Field().String())
}

func validatePhone(fl validator.FieldLevel) bool {
	return phoneRegex.MatchString(fl.Field().String())
}

func validateUsername(fl validator.FieldLevel) bool {
	return usernameRegex.MatchString(fl.Field().String())
}

func validateCode(fl validator.FieldLevel) bool {
	return codeRegex.MatchString(fl.Field().String())
}

func validateEmail(fl validator.FieldLevel) bool {
	return emailRegex.MatchString(fl.Field().String())
}

type valider interface {
	Valid() bool
}

func ValidFunc(field reflect.Value) interface{} {
	if tar, ok := field.Interface().(valider); ok && tar.Valid() {
		return field.String()
	}
	return nil
}
