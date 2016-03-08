package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	SESSION_MAP_NAME = "xtracApiSessionMap"
)

type SessionManager interface {
	RetrieveXtracSession(key string) (string, int, error)
	PersistXtracSession(key, session string) (int, error)
	DeleteXtracSession(key string) (int, error)
	UpdateXtracSession(key, session string) (int, error)
}

type hazelcastEvaluationSessionManager struct {
	client http.Client
	host   string
	port   string
}

func NewHazelcastEvaluationSessionManager(host, port string) SessionManager {
	return hazelcastEvaluationSessionManager{
		client: http.Client{},
		host:   host,
		port:   port,
	}
}

func (manager hazelcastEvaluationSessionManager) PersistXtracSession(key, session string) (int, error) {
	resp, err := manager.client.Post(fmt.Sprintf("http://%s:%s/hazelcast/rest/maps/%s/%s", manager.host, manager.port, SESSION_MAP_NAME, key), "text/plain", strings.NewReader(session))
	if err != nil {
		return 0, err
	} else {
		ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		return resp.StatusCode, nil
	}
}

func (manager hazelcastEvaluationSessionManager) RetrieveXtracSession(key string) (string, int, error) {
	resp, err := manager.client.Get(fmt.Sprintf("http://%s:%s/hazelcast/rest/maps/%s/%s", manager.host, manager.port, SESSION_MAP_NAME, key))
	if err != nil {
		return "", 0, err
	} else {
		value, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", 0, err
		}
		defer resp.Body.Close()
		//fmt.Printf("RetrieveXtracSession: statusCode: %s, key '%s', value '%s'.", resp.Status, key, string(value))
		return string(value), resp.StatusCode, nil
	}
}

func (manager hazelcastEvaluationSessionManager) UpdateXtracSession(key, session string) (int, error) {
	req, _ := http.NewRequest("PUT", fmt.Sprintf("http://%s:%s/hazelcast/rest/maps/%s/%s", manager.host, manager.port, SESSION_MAP_NAME, key), strings.NewReader(session))
	resp, err := manager.client.Do(req)
	if err != nil {
		return 0, err
	} else {
		ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		return resp.StatusCode, nil
	}
}

func (manager hazelcastEvaluationSessionManager) DeleteXtracSession(key string) (int, error) {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("http://%s:%s/hazelcast/rest/maps/%s/%s", manager.host, manager.port, SESSION_MAP_NAME, key), nil)
	resp, err := manager.client.Do(req)
	if err != nil {
		return 0, err
	} else {
		ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		return resp.StatusCode, nil
	}
}
