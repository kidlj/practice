package regex

import "testing"

func TestValidNginxSize(t *testing.T) {
	sizes := []string{
		"0", "123", "10k", "20K", "0k", "0K", "20m", "30M",
	}
	for _, size := range sizes {
		ok := ValidNginxSize(size)
		if !ok {
			t.Errorf("test failed for %s", size)
		}
	}

	badSizes := []string{
		"", "-3", " 333", "10G", "m",
	}

	for _, size := range badSizes {
		ok := ValidNginxSize(size)
		if ok {
			t.Errorf("test failed for %s", size)
		}
	}
}

func TestValidNginxPrefix(t *testing.T) {
	prefixes := []string{
		"/", "/123", "/10k", "/3.20K", "/a_0k", "/中国",
	}
	for _, size := range prefixes {
		ok := ValidNginxPrefix(size)
		if !ok {
			t.Errorf("test failed for %s", size)
		}
	}

	badPrefixes := []string{
		"", "aaa", " 333", "/ 10G", "/m hello", "中国",
	}

	for _, size := range badPrefixes {
		ok := ValidNginxPrefix(size)
		if ok {
			t.Errorf("test failed for %s", size)
		}
	}
}
