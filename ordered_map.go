package main

import (
	"encoding/json"

	"github.com/emirpasic/gods/maps/linkedhashmap"
)

func IsValidJSON(s string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(s), &js) == nil
}

type SortedMap struct {
	Map *linkedhashmap.Map
}

func NewSortedMap() *SortedMap {
	return &SortedMap{
		Map: linkedhashmap.New(),
	}
}

func (sm *SortedMap) FromJSON(json string) *SortedMap {
	if IsValidJSON(json) {
		sm.Map.FromJSON([]byte(json))
	}

	return sm
}

func (sm *SortedMap) ToJSON() string {
	if !sm.Map.Empty() {
		toJson, _ := sm.Map.ToJSON()
		return string(toJson)
	}
	return ""
}

func (sm *SortedMap) GetMap() *linkedhashmap.Map {
	return sm.Map
}

func (sm *SortedMap) AddPair(key string, value interface{}) *SortedMap {
	sm.Map.Put(key, value)
	return sm
}

func (sm *SortedMap) RemovePair(key string) *SortedMap {
	sm.Map.Remove(key)
	return sm
}

func (sm *SortedMap) AddPairs(data map[string]interface{}) *SortedMap {
	for key, val := range data {
		sm.Map.Put(key, val)
	}
	return sm
}

func (sm *SortedMap) ForEach(action func(key string, value interface{})) *SortedMap {
	sm.Map.Each(func(key, value interface{}) {
		action(key.(string), value)
	})
	return sm
}

func (sm *SortedMap) Clear() *SortedMap {
	sm.Map.Clear()

	return sm
}
