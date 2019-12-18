// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.10
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.PolyaxonSdk);
  }
}(this, function(expect, PolyaxonSdk) {
  'use strict';

  var instance;

  describe('(package)', function() {
    describe('V1RunSettings', function() {
      beforeEach(function() {
        instance = new PolyaxonSdk.V1RunSettings();
      });

      it('should create an instance of V1RunSettings', function() {
        // TODO: update the code to test V1RunSettings
        expect(instance).to.be.a(PolyaxonSdk.V1RunSettings);
      });

      it('should have the property namespace (base name: "namespace")', function() {
        // TODO: update the code to test the property namespace
        expect(instance).to.have.property('namespace');
        // expect(instance.namespace).to.be(expectedValueLiteral);
      });

      it('should have the property agent (base name: "agent")', function() {
        // TODO: update the code to test the property agent
        expect(instance).to.have.property('agent');
        // expect(instance.agent).to.be(expectedValueLiteral);
      });

      it('should have the property queue (base name: "queue")', function() {
        // TODO: update the code to test the property queue
        expect(instance).to.have.property('queue');
        // expect(instance.queue).to.be(expectedValueLiteral);
      });

      it('should have the property logs_store (base name: "logs_store")', function() {
        // TODO: update the code to test the property logs_store
        expect(instance).to.have.property('logs_store');
        // expect(instance.logs_store).to.be(expectedValueLiteral);
      });

      it('should have the property outputs_store (base name: "outputs_store")', function() {
        // TODO: update the code to test the property outputs_store
        expect(instance).to.have.property('outputs_store');
        // expect(instance.outputs_store).to.be(expectedValueLiteral);
      });

      it('should have the property init_connections (base name: "init_connections")', function() {
        // TODO: update the code to test the property init_connections
        expect(instance).to.have.property('init_connections');
        // expect(instance.init_connections).to.be(expectedValueLiteral);
      });

      it('should have the property connections (base name: "connections")', function() {
        // TODO: update the code to test the property connections
        expect(instance).to.have.property('connections');
        // expect(instance.connections).to.be(expectedValueLiteral);
      });

      it('should have the property git_accesses (base name: "git_accesses")', function() {
        // TODO: update the code to test the property git_accesses
        expect(instance).to.have.property('git_accesses');
        // expect(instance.git_accesses).to.be(expectedValueLiteral);
      });

      it('should have the property registry_access (base name: "registry_access")', function() {
        // TODO: update the code to test the property registry_access
        expect(instance).to.have.property('registry_access');
        // expect(instance.registry_access).to.be(expectedValueLiteral);
      });

      it('should have the property config_resources (base name: "config_resources")', function() {
        // TODO: update the code to test the property config_resources
        expect(instance).to.have.property('config_resources');
        // expect(instance.config_resources).to.be(expectedValueLiteral);
      });

    });
  });

}));
