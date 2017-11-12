/*
 * Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"). You may
 * not use this file except in compliance with the License. A copy of the
 * License is located at
 *
 *     http://aws.amazon.com/apache2.0/
 *
 * or in the "LICENSE" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */
package com.amazonaws.blox.frontend.controllers;

import com.amazonaws.blox.frontend.models.Environment;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

@Api
@RestController
@RequestMapping(
  path = "/environments",
  produces = "application/json",
  consumes = "application/json"
)
public class EnvironmentController {
  @RequestMapping(path = "/{name}", method = RequestMethod.GET, consumes = "*/*")
  @ApiOperation(value = "Describe Environment by name")
  public Environment describeEnvironment(@PathVariable("name") String name) {
    return new Environment(name);
  }
}
