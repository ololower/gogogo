package Route

import "testing"

// TestMatcherCanServeUrlWithoutParams Проверяем что матчер может обрабатывать обычные строки
func TestMatcherCanServeUrlWithoutParams(t *testing.T) {
	var routePath = "pages"
	var path = "pages"
	var urlMatcher = NewUrlMatcher(routePath, path)

	if !urlMatcher.Match() {
		t.Fatal(routePath + " should be same as  " + path)
	}
}

func TestMatcherCantServeDifferentUrl(t *testing.T) {
	var routePath = "pages"
	var path = "other"
	var urlMatcher = NewUrlMatcher(routePath, path)

	if urlMatcher.Match() {
		t.Fatal(routePath + " shouldn't be same as  " + path)
	}
}

func TestMatcherCantServeLittleDifferentUrl(t *testing.T) {
	var routePath = "pages"
	var path = "pages/other"
	var urlMatcher = NewUrlMatcher(routePath, path)

	if urlMatcher.Match() {
		t.Fatal(routePath + " shouldn't be same as  " + path)
	}
}

func TestMatcherCanServeUrlWithParams(t *testing.T) {
	var routePath = "pages/{category}"
	var path = "pages/news"
	var urlMatcher = NewUrlMatcher(routePath, path)

	if !urlMatcher.Match() {
		t.Fatal("Route path \"pages\" should be same as path \"pages\"")
	}
}

func TestMatcherCanServeUrlWithFewParams(t *testing.T) {
	var routePath = "pages/{category}/{article}"
	var path = "pages/news/one_story_about_my_live"
	var urlMatcher = NewUrlMatcher(routePath, path)

	if !urlMatcher.Match() {
		t.Fatal("Route path \"pages\" should be same as path \"pages\"")
	}
}

func TestMatcherCanServeUrlWithManyParams(t *testing.T) {
	var routePath = "pages/{category}/2020/{article}"
	var path = "pages/news/2020/one_story_about_my_live"
	var urlMatcher = NewUrlMatcher(routePath, path)

	if !urlMatcher.Match() {
		t.Fatal("Route path \"pages\" should be same as path \"pages\"")
	}
}

func TestMatcherCanServeUrlWithManyConcatParams(t *testing.T) {
	var routePath = "pages/{category}-{article}"
	var path = "pages/news-one_story_about_my_live"
	var urlMatcher = NewUrlMatcher(routePath, path)

	if !urlMatcher.Match() {
		t.Fatal("Route path \"pages\" should be same as path \"pages\"")
	}
}

