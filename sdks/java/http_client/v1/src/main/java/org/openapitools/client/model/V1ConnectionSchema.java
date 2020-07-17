// Copyright 2018-2020 Polyaxon, Inc.
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
 * The version of the OpenAPI document: 1.1.4-rc8
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import org.openapitools.client.model.V1BucketConnection;
import org.openapitools.client.model.V1ClaimConnection;
import org.openapitools.client.model.V1GitConnection;
import org.openapitools.client.model.V1HostConnection;
import org.openapitools.client.model.V1HostPathConnection;

/**
 * V1ConnectionSchema
 */

public class V1ConnectionSchema {
  public static final String SERIALIZED_NAME_BUCKET_CONNECTION = "bucket_connection";
  @SerializedName(SERIALIZED_NAME_BUCKET_CONNECTION)
  private V1BucketConnection bucketConnection;

  public static final String SERIALIZED_NAME_HOST_PATH_CONNECTION = "host_path_connection";
  @SerializedName(SERIALIZED_NAME_HOST_PATH_CONNECTION)
  private V1HostPathConnection hostPathConnection;

  public static final String SERIALIZED_NAME_CLAIM_CONNECTION = "claim_connection";
  @SerializedName(SERIALIZED_NAME_CLAIM_CONNECTION)
  private V1ClaimConnection claimConnection;

  public static final String SERIALIZED_NAME_HOST_CONNECTION = "host_connection";
  @SerializedName(SERIALIZED_NAME_HOST_CONNECTION)
  private V1HostConnection hostConnection;

  public static final String SERIALIZED_NAME_GIT_CONNECTION = "git_connection";
  @SerializedName(SERIALIZED_NAME_GIT_CONNECTION)
  private V1GitConnection gitConnection;


  public V1ConnectionSchema bucketConnection(V1BucketConnection bucketConnection) {
    
    this.bucketConnection = bucketConnection;
    return this;
  }

   /**
   * Get bucketConnection
   * @return bucketConnection
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1BucketConnection getBucketConnection() {
    return bucketConnection;
  }


  public void setBucketConnection(V1BucketConnection bucketConnection) {
    this.bucketConnection = bucketConnection;
  }


  public V1ConnectionSchema hostPathConnection(V1HostPathConnection hostPathConnection) {
    
    this.hostPathConnection = hostPathConnection;
    return this;
  }

   /**
   * Get hostPathConnection
   * @return hostPathConnection
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1HostPathConnection getHostPathConnection() {
    return hostPathConnection;
  }


  public void setHostPathConnection(V1HostPathConnection hostPathConnection) {
    this.hostPathConnection = hostPathConnection;
  }


  public V1ConnectionSchema claimConnection(V1ClaimConnection claimConnection) {
    
    this.claimConnection = claimConnection;
    return this;
  }

   /**
   * Get claimConnection
   * @return claimConnection
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1ClaimConnection getClaimConnection() {
    return claimConnection;
  }


  public void setClaimConnection(V1ClaimConnection claimConnection) {
    this.claimConnection = claimConnection;
  }


  public V1ConnectionSchema hostConnection(V1HostConnection hostConnection) {
    
    this.hostConnection = hostConnection;
    return this;
  }

   /**
   * Get hostConnection
   * @return hostConnection
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1HostConnection getHostConnection() {
    return hostConnection;
  }


  public void setHostConnection(V1HostConnection hostConnection) {
    this.hostConnection = hostConnection;
  }


  public V1ConnectionSchema gitConnection(V1GitConnection gitConnection) {
    
    this.gitConnection = gitConnection;
    return this;
  }

   /**
   * Get gitConnection
   * @return gitConnection
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1GitConnection getGitConnection() {
    return gitConnection;
  }


  public void setGitConnection(V1GitConnection gitConnection) {
    this.gitConnection = gitConnection;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1ConnectionSchema v1ConnectionSchema = (V1ConnectionSchema) o;
    return Objects.equals(this.bucketConnection, v1ConnectionSchema.bucketConnection) &&
        Objects.equals(this.hostPathConnection, v1ConnectionSchema.hostPathConnection) &&
        Objects.equals(this.claimConnection, v1ConnectionSchema.claimConnection) &&
        Objects.equals(this.hostConnection, v1ConnectionSchema.hostConnection) &&
        Objects.equals(this.gitConnection, v1ConnectionSchema.gitConnection);
  }

  @Override
  public int hashCode() {
    return Objects.hash(bucketConnection, hostPathConnection, claimConnection, hostConnection, gitConnection);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1ConnectionSchema {\n");
    sb.append("    bucketConnection: ").append(toIndentedString(bucketConnection)).append("\n");
    sb.append("    hostPathConnection: ").append(toIndentedString(hostPathConnection)).append("\n");
    sb.append("    claimConnection: ").append(toIndentedString(claimConnection)).append("\n");
    sb.append("    hostConnection: ").append(toIndentedString(hostConnection)).append("\n");
    sb.append("    gitConnection: ").append(toIndentedString(gitConnection)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

