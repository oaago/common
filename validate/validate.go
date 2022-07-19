package validate

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

type Interfaces struct {
	Url       string       `json:"url,omitempty" validate:"required"`
	Method    string       `json:"method,omitempty" validate:"required,oneof=post get options delete update"`
	Ttl       int64        `json:"ttl,omitempty" validate:"required"`
	Desc      string       `json:"desc,omitempty" validate:"required"`
	ChildNode []Interfaces `json:"child_node,omitempty"`
}

func Input(inter []Interfaces) {
	en := en.New() //英文翻译器
	zh := zh.New() //中文翻译器
	uni := ut.New(en, zh)
	trans, _ := uni.GetTranslator("zh") //获取需要的语言
	validate := validator.New()
	zhtrans.RegisterDefaultTranslations(validate, trans)
	errList := map[string]string{}
	for _, interfaces := range inter {
		err := validate.Struct(interfaces)
		if err != nil {
			fmt.Println("=== error msg ====")
			errs := err.(validator.ValidationErrors)
			list := errs.Translate(trans)
			for field, err := range list {
				errList[field[strings.Index(field, ".")+1:]] = err
			}
		}
	}
	if len(errList) > 0 {
		fmt.Println("数据验证失败：", errList)
	}
}
