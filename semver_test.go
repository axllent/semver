// Semver tests
package semver

import (
	"testing"
)

var validTests = []struct {
	in  string
	out string
}{
	{"bad", ""},
	{"v1-alpha.beta.gamma", ""},
	{"v1-pre", ""},
	{"v1+meta", ""},
	{"v1-pre+meta", ""},
	{"v1.2-pre", ""},
	{"v1.2+meta", ""},
	{"v1.2-pre+meta", ""},
	{"v1.0.0-alpha", "v1.0.0-alpha"},
	{"v1.0.0-alpha.1", "v1.0.0-alpha.1"},
	{"v1.0.0-alpha.beta", "v1.0.0-alpha.beta"},
	{"v1.0.0-beta", "v1.0.0-beta"},
	{"v1.0.0-beta.2", "v1.0.0-beta.2"},
	{"v1.0.0-beta.11", "v1.0.0-beta.11"},
	{"v1.0.0-rc.1", "v1.0.0-rc.1"},
	{"v1", "v1.0.0"},
	{"v1.0", "v1.0.0"},
	{"v1.0.0", "v1.0.0"},
	{"v1.2", "v1.2.0"},
	{"v1.2.0", "v1.2.0"},
	{"v1.2.3-456", "v1.2.3-456"},
	{"v1.2.3-456.789", "v1.2.3-456.789"},
	{"v1.2.3-456-789", "v1.2.3-456-789"},
	{"v1.2.3-456a", "v1.2.3-456a"},
	{"v1.2.3-pre", "v1.2.3-pre"},
	{"v1.2.3-pre+meta", "v1.2.3-pre"},
	{"v1.2.3-pre.1", "v1.2.3-pre.1"},
	{"v1.2.3-zzz", "v1.2.3-zzz"},
	{"v1.2.3", "v1.2.3"},
	{"v1.2.3+meta", "v1.2.3"},
	{"v1.2.3+meta-pre", "v1.2.3"},
}

func TestIsValid(t *testing.T) {
	for _, vt := range validTests {
		ok := IsValid(vt.in)
		if ok != (vt.out != "") {
			t.Errorf("IsValid(%q) = %v, want %v", vt.in, ok, !ok)
		}
	}
}

var versionCompare = []struct {
	in  string
	out int
}{
	{"bad", -1},
	{"0.5.0", -1},
	{"v0.0.0", -1},
	{"v1.0.0-beta", -1},
	{"v1.0.0-beta2", -1},
	{"1.0.1-beta2", 1},
	{"0.0.1", -1},
	{"1.0.1-beta", 1},
	{"v1.0.1-beta", 1},
	{"v2", 1},
	{"2.0", 1},
	{"2.0.0-beta", 1},
}

func TestVersionCompare(t *testing.T) {
	testVersions := []string{"1.0.0", "1.0", "1", "v1.0.0", "v1.0", "v1", "1.0.0-beta3"}

	for _, testVer := range testVersions {
		// compare equal versions
		equals := Compare(testVer, testVer)
		if equals != 0 {
			t.Errorf("TestVersionCompare(%q) = %v, want 0", testVer, equals)
		}
		// compare with list
		for _, v := range versionCompare {
			result := Compare(v.in, testVer)
			if result != v.out {
				t.Errorf("TestVersionCompare(%q) = %v, want %v", v.in, result, v.out)
			}
		}
		// compare with list reversed
		for _, v := range versionCompare {
			result := Compare(testVer, v.in)
			if result == v.out {
				t.Errorf("TestVersionCompare(%q) = %v, did not want %v", v.in, result, v.out)
			}
		}
	}
}

var prereleaseCompare = []struct {
	in  string
	ver string
	out int
}{
	{"v1.0.0-beta", "v1.0.0-beta1", -1},
	{"v1.0.0-beta2", "v1.0.0-beta3", -1},
	{"1.0.1-beta2", "1.0.1-beta", 1},
	{"1.0.1-beta-2", "1.0.1-beta-1", 1},
	{"7.0.1-beta-3", "1.0.1-beta-2", 1},
	{"v1.0.0", "1.0.1-beta", -1},
	{"v1.0.1", "1.0.1-beta", 1},
}

func TestPrereleaseVersionCompare(t *testing.T) {

	for _, v := range prereleaseCompare {
		// compare equal versions
		equals := Compare(v.in, v.ver)
		if equals != v.out {
			t.Errorf("TestVersionCompare(%q, %q) = %v, want %v", v.in, v.ver, v.out, equals)
		}
	}
}

var mmppCompare = []struct {
	in         string
	major      string
	minor      string
	patch      string
	prerelease string
}{
	{"v1.2.3-beta", "1", "2", "3", "beta"},
	{"1.2.3-beta", "1", "2", "3", "beta"},
	{"1.2.3-beta-3", "1", "2", "3", "beta-3"},
	{"1.2.3_beta-3", "", "", "", ""}, // invalid prerelease
	{"1.2", "1", "2", "0", ""},
}

func TestMajorMinorPatchPrerelease(t *testing.T) {
	for _, v := range mmppCompare {
		major := Major(v.in)
		if major != v.major {
			t.Errorf("TestMajorMinorPatchPrerelease(%q) Major = %v, want %v", v.in, major, v.major)
		}
		minor := Minor(v.in)
		if minor != v.minor {
			t.Errorf("TestMajorMinorPatchPrerelease(%q) Minor = %v, want %v", v.in, minor, v.minor)
		}
		patch := Patch(v.in)
		if patch != v.patch {
			t.Errorf("TestMajorMinorPatchPrerelease(%q) Patch = %v, want %v", v.in, patch, v.patch)
		}
		prerelease := Prerelease(v.in)
		if prerelease != v.prerelease {
			t.Errorf("TestMajorMinorPatchPrerelease(%q) Prerelease = %v, want %v", v.in, prerelease, v.prerelease)
		}
	}
}
func TestReverseSort(t *testing.T) {

	versions := []string{"z", "beta", "5.0.0", "v5.1.0", "5.1.0-beta", "1.2.3", "del"}

	sorted := SortMin(versions)

	validSorted := []string{"1.2.3", "5.0.0", "5.1.0-beta", "v5.1.0"}

	failed := false

	for i, v := range sorted {
		if v != validSorted[i] {
			failed = true
		}
	}

	if failed {
		t.Errorf("TestReverseSort() %v, want %v", sorted, validSorted)
	}
}

var maxCompare = []struct {
	first   string
	second  string
	highest string
}{
	{"1.3.4", "1.2.3", "1.3.4"},
	{"v1.0.0-beta", "v1.0.0-beta1", "v1.0.0-beta1"},
	{"v1.0.0-beta2", "v1.0.0-beta3", "v1.0.0-beta3"},
	{"1.0.1-beta2", "1.0.1-beta", "1.0.1-beta2"},
	{"1.0.1-beta2", "1.0.1-beta1", "1.0.1-beta2"},
	{"7.0.1-beta-3", "1.0.1-beta-2", "7.0.1-beta-3"},
	{"v1.0.0", "1.0.1-beta", "1.0.1-beta"},
	{"1.0.1", "1.0.1-alpha", "1.0.1"},
	{"v2.3.4-beta2", "v2.3.4-beta1", "v2.3.4-beta2"},
}

func TestMax(t *testing.T) {

	for _, v := range maxCompare {
		max := Max(v.first, v.second)
		if max != v.highest {
			t.Errorf("TestMax(%s, %s) got %s, want %s", v.first, v.second, max, v.highest)
		}
		// swap order around
		maxr := Max(v.second, v.first)
		if maxr != v.highest {
			t.Errorf("TestMax(%s, %s) got %s, want %s", v.second, v.first, maxr, v.highest)
		}
	}
}

var (
	v1 = "v1.0.0+metadata-dash"
	v2 = "v1.0.0+metadata-dash1"
)

func BenchmarkCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if Compare(v1, v2) != 0 {
			b.Fatalf("bad compare")
		}
	}
}
