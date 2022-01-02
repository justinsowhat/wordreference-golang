package tests

import (
	"os"
	"testing"

	utils "github.com/justinsowhat/wordreference-golang/utils"
)

// Test Parser
func TestParseHappyPath(t *testing.T) {

	file, _ := os.Open("./testData/subir.html")

	actual := utils.Parse(file)

	expectedIPA := "[sybiʀ]"

	if actual.IPA == "" || actual.IPA != expectedIPA {
		t.Fatalf("The actual IPA is not missing or incorrect: %s", actual.IPA)
	}

	if actual.PrincipalTranslations == nil {
		t.Fatal("The actual principal translations are missing")
	}

	if actual.AdditionalTranslations == nil {
		t.Fatal("The actual additional translations are missing")
	}

	if actual.CompoundForms == nil {
		t.Fatal("The actual compound forms are missing")
	}

}

func TestParseNilResponseBody(t *testing.T) {

	actual := utils.Parse(nil)

	if actual.IPA != "" {
		t.Fatalf("The actual result should not have an IPA!")
	}

	if actual.PrincipalTranslations != nil {
		t.Fatal("The actual result should not have principal translations!")
	}

	if actual.AdditionalTranslations != nil {
		t.Fatal("The actual result should not have additional translations!")
	}

	if actual.CompoundForms != nil {
		t.Fatal("The actual result should not have compound forms!")
	}

}

func TestParseEmptyHtml(t *testing.T) {

	file, _ := os.Open("./testData/empty.html")

	actual := utils.Parse(file)

	if actual.IPA != "" {
		t.Fatalf("The actual result should not have an IPA!")
	}

	if actual.PrincipalTranslations != nil {
		t.Fatal("The actual result should not have principal translations!")
	}

	if actual.AdditionalTranslations != nil {
		t.Fatal("The actual result should not have additional translations!")
	}

	if actual.CompoundForms != nil {
		t.Fatal("The actual result should not have compound forms!")
	}

}
