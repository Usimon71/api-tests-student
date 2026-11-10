package advertisementSearch

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/assertions"
	advList "api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
)

const (
	existingAdTitle = "Клюшка хоккейная Bauer Vapor X3"
	searchBeginning = "Клюшка хоккейная"
	searchMiddle    = "хоккейная Bauer"
	searchEnd       = "Vapor X3"
)

type TestSuite struct {
	suite.Suite
	loginData models.LoginOkResponse
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()
	base.Precondition("Check that the search constants are set")
	require.NotEmpty(s.T(), existingAdTitle, "set existingAdTitle")
	require.NotEmpty(s.T(), searchBeginning, "set searchBeginning")
	require.NotEmpty(s.T(), searchMiddle, "set searchMiddle")
	require.NotEmpty(s.T(), searchEnd, "set searchEnd")

	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

func (s *TestSuite) TestSearchByTitle() {
	list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, searchMiddle)

	s.Run("The response contains at least one advertisement", func() {
		require.NotEmpty(s.T(), list.Items)
	})
	s.Run("The results include our advertisement", func() {
		assertions.ContainsAdvertisementByTitle(s.T(), list.Items, existingAdTitle)
	})
}

func (s *TestSuite) TestSearchByTitle_Substring() {
	cases := []struct {
		name   string
		search string
	}{
		{"beginning", searchBeginning},
		{"middle", searchMiddle},
		{"end", searchEnd},
		{"lowercase", strings.ToLower(searchMiddle)},
		{"uppercase", strings.ToUpper(searchMiddle)},
		{"mixed case", utils.MixCase(searchMiddle)},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, tc.search)
			assertions.ContainsAdvertisementByTitle(s.T(), list.Items, existingAdTitle)
		})
	}
}
