package main

import (
	"testing"
	"time"

	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.uber.org/zap/zaptest"
)

// Test generated using Keploy
func TestPutURL_ValidURL_321(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Run("should insert new URL and return shortened link", func(mt *mtest.T) {
		col = mt.Coll
		logger = zaptest.NewLogger(t)
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/put", strings.NewReader(`{"url":"http://example.com"}`))
		c.Request.Header.Set("Content-Type", "application/json")

		putURL(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "http://localhost:8080/")
	})
}

// Test generated using Keploy

// Test generated using Keploy

func TestGetURL_ValidHash_789(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Run("should retrieve URL and redirect", func(mt *mtest.T) {
		col = mt.Coll
		logger = zaptest.NewLogger(t)
		expectedURL := URL{
			ID:      "test-id",
			Created: time.Now(),
			Updated: time.Now(),
			URL:     "http://example.com",
		}
		first := mtest.CreateCursorResponse(
			1,
			"testdb.testcoll",
			mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: expectedURL.ID},
				{Key: "created", Value: expectedURL.Created},
				{Key: "updated", Value: expectedURL.Updated},
				{Key: "url", Value: expectedURL.URL},
			},
		)
		second := mtest.CreateCursorResponse(0, "testdb.testcoll", mtest.NextBatch)
		mt.AddMockResponses(first, second)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "param", Value: "test-id"}}
		c.Request = httptest.NewRequest(http.MethodGet, "/get/test-id", nil)

		getURL(c)

		assert.Equal(t, http.StatusSeeOther, w.Code)
		assert.Equal(t, expectedURL.URL, w.Header().Get("Location"))
	})
}

// Test generated using Keploy

func TestNew_ValidHostAndDB_852(t *testing.T) {
	host := "localhost"
	db := "testdb"

	client, err := New(host, db)
	require.NoError(t, err)
	assert.NotNil(t, client)
}

// Test generated using Keploy

func TestPutURL_MissingURL_654(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Run("should return bad request for missing URL parameter", func(mt *mtest.T) {
		col = mt.Coll
		logger = zaptest.NewLogger(t)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/put", strings.NewReader(`{"invalid_key":"value"}`))
		c.Request.Header.Set("Content-Type", "application/json")

		putURL(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "missing url param")
	})
}

// Test generated using Keploy
