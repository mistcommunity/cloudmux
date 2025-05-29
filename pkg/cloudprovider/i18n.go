package cloudprovider

import (
	"golang.org/x/text/language"
)

type SModelI18nEntry struct {
	Value     string
	valueI18n map[language.Tag]string
}

func NewSModelI18nEntry(value string) *SModelI18nEntry {
	vn := make(map[language.Tag]string, 0)
	return &SModelI18nEntry{Value: value, valueI18n: vn}
}

func (self *SModelI18nEntry) GetKeyValue() string {
	return self.Value
}

func (self *SModelI18nEntry) Lookup(tag language.Tag) string {
	if v, ok := self.valueI18n[tag]; ok {
		return v
	}

	return self.Value
}

func (self *SModelI18nEntry) CN(v string) *SModelI18nEntry {
	self.valueI18n[language.Chinese] = v
	return self
}

func (self *SModelI18nEntry) EN(v string) *SModelI18nEntry {
	self.valueI18n[language.English] = v
	return self
}

type SModelI18nTable map[string]*SModelI18nEntry
