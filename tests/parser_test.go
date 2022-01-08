package tests

import (
	"os"
	"testing"

	utils "github.com/justinsowhat/wordreference-golang/utils"
)

// Test Parser
func TestParseHappyPath(t *testing.T) {

	file, _ := os.Open("./testData/subir.html")

	parser := utils.Parser{}
	actual := parser.Parse(file)

	expectedIPA := "[sybiʀ]"

	if actual.IPA == "" || actual.IPA != expectedIPA {
		t.Fatalf("The actual IPA is not missing or incorrect: %s", actual.IPA)
	}

	if len(actual.TranslationGroups) == 0 {
		t.Fatal("There's no translation groups are missing")
	}

	for _, group := range actual.TranslationGroups {
		if group.Title == utils.PRINCIPAL_TRANSLATIONS {
			if len(group.Translations) == 0 {
				t.Fatal("The actual principal translations are missing")
			}
		}

		if group.Title == utils.ADDITIONAL_TRANSLATION {
			if len(group.Translations) == 0 {
				t.Fatal("The actual additional translations are missing")
			}
		}

		if group.Title == utils.COMPOUND_FORMS {
			if len(group.Translations) == 0 {
				t.Fatal("The actual compound forms are missing")
			}
		}
	}

}

func TestParseNilResponseBody(t *testing.T) {

	parser := utils.Parser{}
	actual := parser.Parse(nil)

	if actual.IPA != "" {
		t.Fatalf("The actual result should not have an IPA!")
	}

	if len(actual.TranslationGroups) != 0 {
		t.Fatal("The actual result should not have translations!")
	}

}

func TestParseEmptyHtml(t *testing.T) {

	file, _ := os.Open("./testData/empty.html")

	parser := utils.Parser{}
	actual := parser.Parse(file)

	if actual.IPA != "" {
		t.Fatalf("The actual result should not have an IPA!")
	}

	if len(actual.TranslationGroups) != 0 {
		t.Fatal("The actual result should not have translations!")
	}

}
