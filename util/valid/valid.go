package valid

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

//自定义验证规则
const (
	NOT_EMPTY      = "NotEmpty" //字符串不能为空
	INT_MAX        = "max"      //int最大值
	INT_MIN        = "min"      //int最小值
	TYPE           = "type"     //类型
	STR_MAX_LENGTH = "max-len"  //字符串最大长度
	STR_MIN_LENGTH = "min-len"  //字符串最小长度
	STR_LENGTH     = "len"      //字符串长度
	RANGE          = "range"    //元素必须在合适的范围内 例:1-100
)

//对外暴露结构体验证函数
func StructValidate(bean interface{}) error {
	fields := reflect.ValueOf(bean).Elem()
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Type().Field(i)
		valid := field.Tag.Get("valid")
		if valid == "" {
			continue
		}
		value := fields.FieldByName(field.Name)
		err := fieldValidate(field.Name, valid, value)
		if err != nil {
			return err
		}
	}
	return nil
}

//属性验证
func fieldValidate(fieldName, valid string, value reflect.Value) error {
	valids := strings.Split(valid, " ")
	for _, valid := range valids {

		if strings.Index(valid, TYPE) != -1 {
			v := value.Type().Name()
			split := strings.Split(valid, "=")
			t := split[1]
			if v != t {
				return errors.New(fieldName + " 类型必须为 " + t)
			}
		}

		if strings.Index(valid, NOT_EMPTY) != -1 {
			str := value.String()
			if str == "" {
				return errors.New(fieldName + " 不能为空")
			}
		}
		if strings.Index(valid, INT_MIN) != -1 {
			v := value.Int()
			split := strings.Split(valid, "=")
			rule, err := strconv.Atoi(split[1])
			if err != nil {
				return errors.New(fieldName + ":验证规则有误")
			}
			if int(v) < rule {
				return errors.New(fieldName + " 不能小于 " + strconv.Itoa(rule))
			}
		}

		if strings.Index(valid, INT_MAX) != -1 {
			v := value.Int()
			split := strings.Split(valid, "=")
			rule, err := strconv.Atoi(split[1])
			if err != nil {
				return errors.New(fieldName + ":验证规则有误")
			}
			if int(v) > rule {

				return errors.New(fieldName + " 不能大于 " + strconv.Itoa(rule))
			}
		}
		//字符串特殊处理
		if value.Type().Name() == "string" {
			if strings.Index(valid, STR_LENGTH) != -1 {
				v := value.String()
				split := strings.Split(valid, "=")
				lenStr := split[1]
				length, err := strconv.Atoi(lenStr)
				if err != nil {
					return errors.New(fieldName + " " + STR_LENGTH + " rule is error")
				}
				if len(v) != length {
					return errors.New(fieldName + " 长度必须为 " + lenStr)
				}
			}
			if strings.Index(valid, STR_MAX_LENGTH) != -1 {
				v := value.String()
				split := strings.Split(valid, "=")
				lenStr := split[1]
				length, err := strconv.Atoi(lenStr)
				if err != nil {
					return errors.New(fieldName + " " + STR_LENGTH + " rule is error")
				}
				if len(v) > length {
					return errors.New(fieldName + " 长度不能超过 " + lenStr)
				}
			}

			if strings.Index(valid, STR_MIN_LENGTH) != -1 {
				v := value.String()
				split := strings.Split(valid, "=")
				lenStr := split[1]
				length, err := strconv.Atoi(lenStr)
				if err != nil {
					return errors.New(fieldName + " " + STR_LENGTH + " rule is error")
				}
				if len(v) < length {
					return errors.New(fieldName + " 长度不能小于 " + lenStr)
				}
			}
		}
	}
	return nil
}



