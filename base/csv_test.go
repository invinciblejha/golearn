package base

import "testing"

func TestParseCSVGetRows(testEnv *testing.T) {
	lineCount := ParseCSVGetRows("../examples/datasets/iris.csv")
	if lineCount != 150 {
		testEnv.Error("Should have %d lines, has %d", 150, lineCount)
	}

	lineCount = ParseCSVGetRows("../examples/datasets/iris_headers.csv")
	if lineCount != 151 {
		testEnv.Error("Should have %d lines, has %d", 151, lineCount)
	}

}

func TestParseCCSVGetAttributes(testEnv *testing.T) {
	attrs := ParseCSVGetAttributes("../examples/datasets/iris_headers.csv", true)
	if attrs[0].GetType() != Float64Type {
		testEnv.Error("First attribute should be a float, %s", attrs[0])
	}
	if attrs[0].GetName() != "Sepal length" {
		testEnv.Error(attrs[0].GetName())
	}

	if attrs[4].GetType() != CategoricalType {
		testEnv.Error("Final attribute should be categorical, %s", attrs[4])
	}
	if attrs[4].GetName() != "Species" {
		testEnv.Error(attrs[4])
	}
}

func TestParseCsvSniffAttributeTypes(testEnv *testing.T) {
	attrs := ParseCSVSniffAttributeTypes("../examples/datasets/iris_headers.csv", true)
	if attrs[0].GetType() != Float64Type {
		testEnv.Error("First attribute should be a float, %s", attrs[0])
	}
	if attrs[1].GetType() != Float64Type {
		testEnv.Error("Second attribute should be a float, %s", attrs[1])
	}
	if attrs[2].GetType() != Float64Type {
		testEnv.Error("Third attribute should be a float, %s", attrs[2])
	}
	if attrs[3].GetType() != Float64Type {
		testEnv.Error("Fourth attribute should be a float, %s", attrs[3])
	}
	if attrs[4].GetType() != CategoricalType {
		testEnv.Error("Final attribute should be categorical, %s", attrs[4])
	}
}

func TestParseCSVSniffAttributeNamesWithHeaders(testEnv *testing.T) {
	attrs := ParseCSVSniffAttributeNames("../examples/datasets/iris_headers.csv", true)
	if attrs[0] != "Sepal length" {
		testEnv.Error(attrs[0])
	}
	if attrs[1] != "Sepal width" {
		testEnv.Error(attrs[1])
	}
	if attrs[2] != "Petal length" {
		testEnv.Error(attrs[2])
	}
	if attrs[3] != "Petal width" {
		testEnv.Error(attrs[3])
	}
	if attrs[4] != "Species" {
		testEnv.Error(attrs[4])
	}
}

func TestReadInstances(testEnv *testing.T) {
	inst, err := ParseCSVToInstances("../examples/datasets/iris_headers.csv", true)
	if err != nil {
		testEnv.Error(err)
		return
	}
	row1 := inst.RowStr(0)
	row2 := inst.RowStr(50)
	row3 := inst.RowStr(100)

	if row1 != "5.10 3.50 1.40 0.20 Iris-setosa" {
		testEnv.Error(row1)
	}
	if row2 != "7.00 3.20 4.70 1.40 Iris-versicolor" {
		testEnv.Error(row2)
	}
	if row3 != "6.30 3.30 6.00 2.50 Iris-virginica" {
		testEnv.Error(row3)
	}
}

func TestReadAwkwardInsatnces(testEnv *testing.T) {
	inst, err := ParseCSVToInstances("../examples/datasets/chim.csv", true)
	if err != nil {
		testEnv.Error(err)
		return
	}
	if inst.GetAttr(0).GetType() != Float64Type {
		testEnv.Error("Should be float!")
	}
	if inst.GetAttr(1).GetType() != CategoricalType {
		testEnv.Error("Should be discrete!")
	}
}
