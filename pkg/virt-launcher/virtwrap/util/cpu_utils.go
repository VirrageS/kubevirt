/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2017, 2018 Red Hat, Inc.
 *
 */

package util

import (
	"bufio"
	"fmt"
	"os"

	"kubevirt.io/kubevirt/pkg/util"
	"kubevirt.io/kubevirt/pkg/virt-handler/cgroup"

	"k8s.io/utils/cpuset"
)

func GetPodCPUSet() (cpuset.CPUSet, error) {
	var cpusetLine string
	file, err := os.Open(cgroup.GetGlobalCpuSetPath())
	if err != nil {
		return cpuset.New(), err
	}
	defer util.CloseIOAndCheckErr(file, nil)
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		cpusetLine = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return cpuset.New(), err
	}
	cpuSet, err := cpuset.Parse(cpusetLine)
	if err != nil {
		return cpuSet, fmt.Errorf("failed to parse cpuset file: %v", err)
	}
	return cpuSet, nil
}
