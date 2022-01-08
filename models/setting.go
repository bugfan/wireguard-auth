package models

import (
	"github.com/sirupsen/logrus"
)

type Setting struct {
	Key   string `xorm:"pk"`
	Value string `xorm:"varchar(450) 'value'"`
}

func InitSetting(m map[string]string) {
	for key, value := range m {
		s := new(Setting)
		has, err := x.ID(key).Get(s)
		if err != nil {
			logrus.Errorf("setting %s fail %s", key, err.Error())
		}
		if !has {
			s.Key = key
			s.Value = value
			x.Insert(s)
		}
	}
}

func SetValue(key string, value string) error {
	set := &Setting{
		Key:   key,
		Value: value,
	}
	_, err := x.ID(key).Update(set)
	return err
}
func GetValue(key string) string {
	set := &Setting{Key: key}
	_, err := x.Table(set).Get(set)
	if err != nil {
		logrus.Error(err)
		return ""
	}
	return set.Value
}
