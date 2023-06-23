package api

import (
	"bytes"
	"database/sql"
	mockdb "ecom/database/mock"
	db "ecom/database/sqlc"
	"ecom/database/util"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetUserAPI(t *testing.T) {
	account := randomUser()
	testCases := []struct {
		name          string
		accountID     int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		//{
		//	name:      "OK",
		//	accountID: account.ID,
		//	buildStubs: func(store *mockdb.MockStore) {
		//		store.EXPECT().
		//			GetUser(gomock.Any(), gomock.Eq(account.ID)).
		//			Times(1).
		//			Return(account, nil)
		//	},
		//	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		//		require.Equal(t, http.StatusOK, recorder.Code)
		//		requireBodyMatchUser(t, recorder.Body, account)
		//	},
		//},
		{
			name:      "NotFound",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "InternalError",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "InvalidID",
			accountID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/users/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, account db.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.User
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}

func randomUser() db.User {
	return db.User{
		ID:          util.RandomInt(1, 999999),
		FullName:    util.RandomFullName(),
		Email:       util.RandomEmail(),
		Gender:      util.RandomGender(),
		DateOfBirth: time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CountryCode: util.RandomString(3),
	}
}
