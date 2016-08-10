/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"errors"
	"github.com/acksin/strum/shared"
	"github.com/acksin/strum/stats"
)

// Agent runs a STRUM Cloud agent.
type agent struct {
	Config shared.Config

	configFile string
}

func (a *agent) Synopsis() string {
	return "Run a STRUM Cloud agent."
}

func (a *agent) Help() string {
	return "Run a STRUM Cloud agent."
}

func (a *agent) parseConfig() error {
	b, err := ioutil.ReadFile(a.configFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &a.Config)
	if err != nil {
		return err
	}

	if a.Config.APIKey == "" {
		return errors.New("Set the `APIKey`. For STRUM Cloud this can be found at https://www.acksin.com/console/credentials")
	}

	return nil
}

func (a *agent) post() error {
	s := stats.New([]int{})
	jsonStr, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
		return err
	}

	req, err := http.NewRequest("POST", a.Config.URL, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Acksin-API-Key", a.Config.APIKey)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var respForm struct {
		ID string
	}

	if err = json.Unmarshal(body, &respForm); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (a *agent) Run(args []string) int {
	log.Println("Starting STRUM Agent...")

	if len(args) == 0 {
		log.Println("need to pass a config file")
		return -1
	}

	a.configFile = args[0]

	if err := a.parseConfig(); err != nil {
		log.Println(err)
		return -1
	}

	for {
		if err := a.post(); err != nil {
			return -1
		}

		select {
		case <-time.After(1 * time.Hour):
			log.Println("Ping")
		}
	}
}