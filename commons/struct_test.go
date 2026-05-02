package commons

import (
  "fmt"
  "testing"

  "go.mongodb.org/mongo-driver/v2/bson"
)

func TestBsonDFromSlice(t *testing.T) {
  in := [][]any{
    {"first", -1},
    {"second", 1},
  }

  d := BsonDFromSlice(in)

  expected := bson.D{
    {Key: "first", Value: -1},
    {Key: "second", Value: 1},
  }

  AssertDeepEqual(t, d, expected)
}

func TestStructHasProperty(t *testing.T) {
  type O struct {
    Name string
    Age int
  }

  o := &O{
    Name: "Nick",
    Age: 15,
  }

  if !StructHasProperty(o, "Name") {
    t.Fatal("Should have property name")
  }
}

func TestStructSetValue(t *testing.T) {
  type O struct {
    Name string
    Age int
    Active bool
  }

  o := &O{
    Name: "Nick",
    Age: 15,
    Active: true,
  }

  StructSetValue(o, "Name", "George")

  if o.Name != "George" {
    t.Fatalf("Name should have been George not %s", o.Name)
  }
}

func TestStructToStruct(t *testing.T) {
  type O struct {
    Name string
    Age int
    Active bool
  }

  o := &O{
    Name: "Nick",
    Age: 15,
    Active: true,
  }

  var oo O

  StructToStruct(o, &oo)

  if oo.Name != "Nick" || oo.Age != 15 || !oo.Active {
    t.Fatalf("failed to copy %v", oo)
  }
}

func TestStructToMapRecursive(t *testing.T) {
  type O struct {
    Name string
    Age int
    Active bool
  }

  o := &O{
    Name: "Nick",
    Age: 15,
    Active: true,
  }

  mAny := StructToMapRecursive(o)

  switch mAny.(type) {
  case map[string]any:
      m := mAny.(map[string]any)
      if v, ok := m["Name"]; !ok || v != "Nick" {
        t.Fatalf("Unexpected map %v", v)
      }

      if v, ok := m["Age"]; !ok || v != 15 {
        t.Fatalf("Unexpected map %v", v)
      }

      if v, ok := m["Active"]; !ok || v != true {
        t.Fatalf("Unexpected map %v", v)
      }
    default:
      t.Fatal("is not map")
  }
}

func TestMapToStruct(t *testing.T) {
  o := map[string]any {
    "name": "name1",
    "age": 22,
    "active": true,
  }

  type O struct {
    Name string
    Age int
    Active bool
  }

  var oo O

  if err := MapToStruct(o, &oo); err != nil {
    t.Fatalf("1 Failed to convert %v", err)
  }

  if oo.Name != "name1" || oo.Age != 22 || !oo.Active {
    t.Fatalf("2 Failed to convert %v", oo)
  }
}

func TestMapIsSubset(t *testing.T) {
  super := map[string]any {
    "name": "name1",
    "age": 22,
    "active": "2",
  }

  sub := map[string]any {
    "name": "a",
  }

  if !MapIsSubset(sub, super) {
    t.Fatal("Should have indicated that map is a subset")
  }
}

func TestMapIsSubset_Nosubset(t *testing.T) {
  super := map[string]any {
    "name": "name1",
    "age": 22,
    "active": "2",
  }

  sub := map[string]any {
    "foo": "a",
  }

  if MapIsSubset(sub, super) {
    t.Fatal("Should have indicated that map is a NOT subset")
  }
}

func TestMapIsSubsetOfStruct_NoSubset(t *testing.T) {
  type O struct {
    Name string
    Age int `json:"age"`
    Notes string
  }

  oo := map[string]any{
    "name": "A Name",
    "age": 1,
    "notes": "dsafda",
    "other": "dd",
  }

  if MapIsSubsetOfStruct(oo, &O{}) {
    t.Fatal("Should have indicated that map is a NOT subset")
  }
}

func TestMapIsSubsetOfStruct(t *testing.T) {
  type O struct {
    Name string
    Age int `json:"age"`
    Notes string
  }

  oo := map[string]any{
    "Name": "A Name",
  }

  if !MapIsSubsetOfStruct(oo, &O{}) {
    t.Fatal("Should have indicated that map is a subset")
  }
}

func TestMapIsSubsetOfStruct_SomeFields(t *testing.T) {
  type O struct {
    Name string
    Age int `json:"age"`
    Notes string
  }

  oo := map[string]any{
    "Name": "A Name",
    "Foo": "bar",
  }

  if MapIsSubsetOfStruct(oo, &O{}) {
    t.Fatal("Should have indicated that map is a NOT subset")
  }
}

func TestStructHasJsonName_True(t *testing.T) {
  type O struct {
    Name string
    Age int `json:"age"`
    Notes string
  }

  o := &O{
  }

  if !StructHasJsonName(o, "age") {
    t.Fatal("False negative for json tag")
  }
}

func TestStructHasJsonName_False(t *testing.T) {
  type O struct {
    Name string
    Age int `json:"age"`
    Notes string
  }

  o := &O{
  }

  if StructHasJsonName(o, "foo") {
    t.Fatal("False positive for json tag")
  }
}

func TestStructHasBsonName_True(t *testing.T) {
  type O struct {
    Name string
    Age int `bson:"age"`
    Notes string
  }

  o := &O{
  }

  if !StructHasBsonName(o, "age") {
    t.Fatal("False negative for bson tag")
  }
}

func TestStructHasBsonName_False(t *testing.T) {
  type O struct {
    Name string
    Age int `bson:"age"`
    Notes string
  }

  o := &O{
  }

  if StructHasBsonName(o, "foo") {
    t.Fatal("False positive for bson tag")
  }
}

func TestCopyMatching(t *testing.T) {
  type O struct {
    Name string
    Age int
    Notes string
  }

  type OO struct {
    Name string
    Age int
    Active bool
  }

  from := &O{
    Name: "Nick",
    Age: 15,
    Notes: "Some notes about nick",
  }

  var to OO

  StructCopyMatching(from, &to)

  if to.Name != "Nick" && to.Age != 15 && to.Active != true {
    t.Fatalf("Failed to copy matching fields to %v", to)
  }
}

func TestMapMerge(t *testing.T) {
  m1 := map[string]any{
    "a": "va",
    "b": "vb",
  }
  m2 := map[string]any{
    "c": "vc",
    "d": "vd",
  }

  m := MapMerge(m1, m2)

  if len(m) != 4 {
    t.Fatalf("Merged map should have lenght 4 not %d", len(m))
  }
}

func TestBsonClone(t *testing.T) {
  orig := map[string]any{
    "name": "Kolias",
    "age": int32(56),
    "salary": int64(100000000),
    "meta": map[string]any{
      "num1": int32(110),
    },
  }

  clone := BsonClone(orig)

  if "int32" != fmt.Sprintf("%T", clone["age"]) {
    t.Fatalf("int32 should map to int32")
  }

  if "int64" != fmt.Sprintf("%T", clone["salary"]) {
    t.Fatalf("int64 should map to int64")
  }

  meta := clone["meta"].(map[string]any)

  if "int32" != fmt.Sprintf("%T", meta["num1"]) {
    t.Fatalf("netsted type should have been int32")
  }
}