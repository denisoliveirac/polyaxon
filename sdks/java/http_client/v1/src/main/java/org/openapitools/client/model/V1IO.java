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
import java.util.ArrayList;
import java.util.List;

/**
 * V1IO
 */

public class V1IO {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_IOTYPE = "iotype";
  @SerializedName(SERIALIZED_NAME_IOTYPE)
  private String iotype;

  public static final String SERIALIZED_NAME_VALUE = "value";
  @SerializedName(SERIALIZED_NAME_VALUE)
  private Object value;

  public static final String SERIALIZED_NAME_IS_OPTIONAL = "is_optional";
  @SerializedName(SERIALIZED_NAME_IS_OPTIONAL)
  private Boolean isOptional;

  public static final String SERIALIZED_NAME_IS_LIST = "is_list";
  @SerializedName(SERIALIZED_NAME_IS_LIST)
  private Boolean isList;

  public static final String SERIALIZED_NAME_IS_FLAG = "is_flag";
  @SerializedName(SERIALIZED_NAME_IS_FLAG)
  private Boolean isFlag;

  public static final String SERIALIZED_NAME_DELAY_VALIDATION = "delay_validation";
  @SerializedName(SERIALIZED_NAME_DELAY_VALIDATION)
  private Boolean delayValidation;

  public static final String SERIALIZED_NAME_OPTIONS = "options";
  @SerializedName(SERIALIZED_NAME_OPTIONS)
  private List<Object> options = null;


  public V1IO name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public V1IO description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public V1IO iotype(String iotype) {
    
    this.iotype = iotype;
    return this;
  }

   /**
   * Get iotype
   * @return iotype
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getIotype() {
    return iotype;
  }


  public void setIotype(String iotype) {
    this.iotype = iotype;
  }


  public V1IO value(Object value) {
    
    this.value = value;
    return this;
  }

   /**
   * Get value
   * @return value
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Object getValue() {
    return value;
  }


  public void setValue(Object value) {
    this.value = value;
  }


  public V1IO isOptional(Boolean isOptional) {
    
    this.isOptional = isOptional;
    return this;
  }

   /**
   * Get isOptional
   * @return isOptional
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIsOptional() {
    return isOptional;
  }


  public void setIsOptional(Boolean isOptional) {
    this.isOptional = isOptional;
  }


  public V1IO isList(Boolean isList) {
    
    this.isList = isList;
    return this;
  }

   /**
   * Get isList
   * @return isList
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIsList() {
    return isList;
  }


  public void setIsList(Boolean isList) {
    this.isList = isList;
  }


  public V1IO isFlag(Boolean isFlag) {
    
    this.isFlag = isFlag;
    return this;
  }

   /**
   * Get isFlag
   * @return isFlag
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIsFlag() {
    return isFlag;
  }


  public void setIsFlag(Boolean isFlag) {
    this.isFlag = isFlag;
  }


  public V1IO delayValidation(Boolean delayValidation) {
    
    this.delayValidation = delayValidation;
    return this;
  }

   /**
   * Get delayValidation
   * @return delayValidation
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDelayValidation() {
    return delayValidation;
  }


  public void setDelayValidation(Boolean delayValidation) {
    this.delayValidation = delayValidation;
  }


  public V1IO options(List<Object> options) {
    
    this.options = options;
    return this;
  }

  public V1IO addOptionsItem(Object optionsItem) {
    if (this.options == null) {
      this.options = new ArrayList<Object>();
    }
    this.options.add(optionsItem);
    return this;
  }

   /**
   * Get options
   * @return options
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getOptions() {
    return options;
  }


  public void setOptions(List<Object> options) {
    this.options = options;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1IO v1IO = (V1IO) o;
    return Objects.equals(this.name, v1IO.name) &&
        Objects.equals(this.description, v1IO.description) &&
        Objects.equals(this.iotype, v1IO.iotype) &&
        Objects.equals(this.value, v1IO.value) &&
        Objects.equals(this.isOptional, v1IO.isOptional) &&
        Objects.equals(this.isList, v1IO.isList) &&
        Objects.equals(this.isFlag, v1IO.isFlag) &&
        Objects.equals(this.delayValidation, v1IO.delayValidation) &&
        Objects.equals(this.options, v1IO.options);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, description, iotype, value, isOptional, isList, isFlag, delayValidation, options);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1IO {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    iotype: ").append(toIndentedString(iotype)).append("\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    isOptional: ").append(toIndentedString(isOptional)).append("\n");
    sb.append("    isList: ").append(toIndentedString(isList)).append("\n");
    sb.append("    isFlag: ").append(toIndentedString(isFlag)).append("\n");
    sb.append("    delayValidation: ").append(toIndentedString(delayValidation)).append("\n");
    sb.append("    options: ").append(toIndentedString(options)).append("\n");
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

