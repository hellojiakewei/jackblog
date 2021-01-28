package utils

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 定义一个全局翻译器T
var Trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		//v.RegisterStructValidation(SignUpParamStructLevelValidation, SignUpParam{})

		// 在校验器注册自定义的校验方法
		if err := v.RegisterValidation("checkDate", customFunc); err != nil {
			return err
		}

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}
		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		}
		if err != nil {
			return err
		}
		// 注意！因为要使用到trans实例
		// 所以这一步注册要放到trans初始化的后面
		if err := v.RegisterTranslation(
			"checkDate",
			Trans,
			registerTranslator("checkDate", "{0}必须要晚于当前日期"),
			translate,
		); err != nil {
			return err
		}
		return
	}
	return
}

//type SignUpParam struct {
//	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
//	Name       string `json:"name" binding:"required"`
//	Email      string `json:"email" binding:"required,email"`
//	Password   string `json:"password" binding:"required"`
//	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
//	// 需要使用自定义校验方法checkDate做参数校验的字段Date
//	Date string `json:"date" binding:"required,datetime=2006-01-02,checkDate"`
//}

// customFunc 自定义校验方法
func customFunc(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	if date.Before(time.Now()) {
		return false
	}
	return true
}

// SignUpParamStructLevelValidation 自定义SignUpParam结构体校验函数
//func SignUpParamStructLevelValidation(sl validator.StructLevel) {
//	su := sl.Current().Interface().(SignUpParam)
//
//	if su.Password != su.RePassword {
//		// 输出错误提示信息
//		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
//	}
//}

// removeTopStruct 去除字段名中的结构体名称标识
// refer from:https://github.com/go-playground/validator/issues/633#issuecomment-654382345
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	fmt.Println(fe.Tag())
	fmt.Println(fe.Field())
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}
