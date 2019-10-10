package vector

import (
    "math"
    "math/rand"
    "testing"
)

const (
    e = 1e-9
    L1 = int(1e2)
    L2 = int(1e3)
    L3 = int(1e4)
)

func cmp(a float64, b float64) bool {
    return math.Abs(a - b) > e
}

func generateRandomVector(l int) Vector {
    v := make(Vector, l, l)
    for i, _ := range v {
        v[i] = rand.Float64()
    }
    return v
}

func verifyVectorsEquality(t *testing.T, a Vector, b Vector) {
    if len(a) != len(b) {
        t.Fatalf("Vectors lengths are not equal: %d != %d.", len(a), len(b))
    }

    for i, _ := range a {
        if cmp(a[i], b[i]) {
            t.Errorf("Vectors are not equal at %d: %f != %f.", i, a[i], b[i])
        }
    }
}

func TestThatVectorsAreSuccessfullyAdded(t *testing.T) {
    a := Vector{1.0, 2.1, 3.2}
    b := Vector{2.0, 1.5, -10.5}

    result, err := Add(a, b)
    if err != nil {
        t.Fatalf("Add failed: %s.", err)
    }
    verifyVectorsEquality(t, result, Vector{1.0 + 2.0, 2.1 + 1.5, 3.2 - 10.5})

    result, err = Perform("add", a, b)
    if err != nil {
        t.Fatalf("Add failed: %s.", err)
    }
    verifyVectorsEquality(t, result, Vector{1.0 + 2.0, 2.1 + 1.5, 3.2 - 10.5})
}

func TestThatVectorsAreSuccessfullySubtracted(t *testing.T) {
    a := Vector{1.0, 2.1, 3.2}
    b := Vector{2.0, 1.5, -10.5}

    result, err := Sub(a, b)
    if err != nil {
        t.Fatalf("Sub failed: %s.", err)
    }
    verifyVectorsEquality(t, result, Vector{1.0 - 2.0, 2.1 - 1.5, 3.2 + 10.5})

    result, err = Perform("sub", a, b)
    if err != nil {
        t.Fatalf("Sub failed: %s.", err)
    }
    verifyVectorsEquality(t, result, Vector{1.0 - 2.0, 2.1 - 1.5, 3.2 + 10.5})
}

func TestThatVectorsAreSuccessfullyMultiplied(t *testing.T) {
    a := Vector{1.0, 2.1, 3.2}
    b := Vector{2.0, 1.5, -10.5}

    result1, err := Dot(a, b)
    if err != nil {
        t.Fatalf("Dot failed: %s.", err)
    }

    result2, err := Perform("dot", a, b)
    if err != nil {
        t.Fatalf("Dot failed: %s.", err)
    }
    verifyVectorsEquality(t, result1, result2)

    if len(result1) != 1 {
        t.Errorf("Wrong dot result: must be a vector with one component.")
    }

    expected := 1.0 * 2.0 + 2.1 * 1.5 + 3.2 * -10.5
    if cmp(result1[0], expected) {
        t.Errorf("Wrong dot result: %f != %f.", result1[0], expected)
    }
}

func TestThatPerformReturnsErrorForUnknownOperation(t *testing.T) {
    a := Vector{1.0, 2.1, 3.2}
    b := Vector{2.0, 1.5, -10.5}

    result, err := Perform("mul", a, b)
    if err == nil {
        t.Fatalf("Perform has not returned an error.")
    }
    if result != nil {
        t.Fatalf("Result vector must be nil in case of error.")
    }
}

func TestThatAddFailsForVectorsWithDifferentLength(t *testing.T) {
    a := Vector{1.0, 2.1, 3.2}
    b := Vector{2.0, 1.5}

    result, err := Add(a, b)
    if err == nil {
        t.Fatalf("Add has not returned an error.")
    }
    if result != nil {
        t.Fatalf("Add result must be nil in case of error.")
    }
}

func TestThatSubFailsForVectorsWithDifferentLength(t *testing.T) {
    a := Vector{1.0, 2.1, 3.2}
    b := Vector{2.0, 1.5}

    result, err := Sub(a, b)
    if err == nil {
        t.Fatalf("Sub has not returned an error.")
    }
    if result != nil {
        t.Fatalf("Sub result must be nil in case of error.")
    }
}

func TestThatDotFailsForVectorsWithDifferentLength(t *testing.T) {
    a := Vector{1.0, 2.1, 3.2}
    b := Vector{2.0, 1.5}

    result, err := Dot(a, b)
    if err == nil {
        t.Fatalf("Dot has not returned an error.")
    }
    if result != nil {
        t.Fatalf("Dot result must be nil in case of error.")
    }
}

func benchmarkOperation(op string, l int, b *testing.B) {
    v1 := generateRandomVector(l)
    v2 := generateRandomVector(l)
    b.ResetTimer()

    for n := 0; n < b.N; n++ {
        _, err := Perform(op, v1, v2)
        if err != nil {
            b.Fatalf("Add has returned an error: %s.", err)
        }
    }
}

func BenchmarkAddL1(b *testing.B) {
    benchmarkOperation("add", L1, b)
}

func BenchmarkAddL2(b *testing.B) {
    benchmarkOperation("add", L2, b)
}

func BenchmarkAddL3(b *testing.B) {
    benchmarkOperation("add", L3, b)
}

func BenchmarkSubL1(b *testing.B) {
    benchmarkOperation("sub", L1, b)
}

func BenchmarkSubL2(b *testing.B) {
    benchmarkOperation("sub", L2, b)
}

func BenchmarkSubL3(b *testing.B) {
    benchmarkOperation("sub", L3, b)
}

func BenchmarkDotL1(b *testing.B) {
    benchmarkOperation("dot", L1, b)
}

func BenchmarkDotL2(b *testing.B) {
    benchmarkOperation("dot", L2, b)
}

func BenchmarkDotL3(b *testing.B) {
    benchmarkOperation("dot", L3, b)
}
