package rest

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

var managers []SessionManager = []SessionManager{
	NewHazelcastEvaluationSessionManager("192.168.99.100", "5701"),
	NewHazelcastEvaluationSessionManager("192.168.99.100", "5702"),
}

func getRandomManager() SessionManager {
	return managers[rand.Intn(len(managers))]
}

func TestCrudFunctions(t *testing.T) {
	status, err := getRandomManager().PersistXtracSession("key1", "221963312296771475874091257270746538101")
	assert.NoError(t, err)
	assert.Equal(t, 200, status)

	session, status, err := getRandomManager().RetrieveXtracSession("key1")
	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, "221963312296771475874091257270746538101", session)

	status, err = getRandomManager().UpdateXtracSession("key1", "221963312296771475874091257270746538102")
	assert.NoError(t, err)
	assert.Equal(t, 200, status)

	session, status, err = getRandomManager().RetrieveXtracSession("key1")
	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, "221963312296771475874091257270746538102", session)

	session, status, err = getRandomManager().RetrieveXtracSession("key2")
	assert.NoError(t, err)
	assert.Equal(t, 204, status)
	assert.Empty(t, session)

	status, err = getRandomManager().DeleteXtracSession("key1")
	assert.NoError(t, err)
	assert.Equal(t, 200, status)

	session, status, err = getRandomManager().RetrieveXtracSession("key1")
	assert.NoError(t, err)
	assert.Equal(t, 204, status)
	assert.Empty(t, session)
}

func BenchmarkCrud(b *testing.B) {
	for i := 0; i < b.N; i++ {
		status, err := getRandomManager().PersistXtracSession("key1", "221963312296771475874091257270746538101")
		assert.NoError(b, err)
		assert.Equal(b, 200, status)

		session, status, err := getRandomManager().RetrieveXtracSession("key1")
		assert.NoError(b, err)
		assert.Equal(b, 200, status)
		assert.Equal(b, "221963312296771475874091257270746538101", session)

		status, err = getRandomManager().UpdateXtracSession("key1", "221963312296771475874091257270746538102")
		assert.NoError(b, err)
		assert.Equal(b, 200, status)

		session, status, err = getRandomManager().RetrieveXtracSession("key1")
		assert.NoError(b, err)
		assert.Equal(b, 200, status)
		assert.Equal(b, "221963312296771475874091257270746538102", session)

		session, status, err = getRandomManager().RetrieveXtracSession("key2")
		assert.NoError(b, err)
		assert.Equal(b, 204, status)
		assert.Empty(b, session)

		status, err = getRandomManager().DeleteXtracSession("key1")
		assert.NoError(b, err)
		assert.Equal(b, 200, status)

		session, status, err = getRandomManager().RetrieveXtracSession("key1")
		assert.NoError(b, err)
		assert.Equal(b, 204, status)
		assert.Empty(b, session)
	}
}

func BenchmarkCrudParallel(b *testing.B) {
	status, err := getRandomManager().PersistXtracSession("key1", "221963312296771475874091257270746538101")
	assert.NoError(b, err)
	assert.Equal(b, 200, status)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			session, status, err := getRandomManager().RetrieveXtracSession("key1")
			assert.NoError(b, err)
			assert.Equal(b, 200, status)
			assert.Equal(b, "221963312296771475874091257270746538101", session)

			status, err = getRandomManager().UpdateXtracSession("key1", "221963312296771475874091257270746538101")
			assert.NoError(b, err)
			assert.Equal(b, 200, status)

			session, status, err = getRandomManager().RetrieveXtracSession("key1")
			assert.NoError(b, err)
			assert.Equal(b, 200, status)
			assert.Equal(b, "221963312296771475874091257270746538101", session)

			session, status, err = getRandomManager().RetrieveXtracSession("key2")
			assert.NoError(b, err)
			assert.Equal(b, 204, status)
			assert.Empty(b, session)
		}
	})

	status, err = getRandomManager().DeleteXtracSession("key1")
	assert.NoError(b, err)
	assert.Equal(b, 200, status)
}
